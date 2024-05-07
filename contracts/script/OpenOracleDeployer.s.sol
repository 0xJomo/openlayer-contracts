// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import "@eigenlayer/contracts/permissions/PauserRegistry.sol";
import {IAVSDirectory} from "@eigenlayer/contracts/interfaces/IAVSDirectory.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IStrategyManager, IStrategy} from "@eigenlayer/contracts/interfaces/IStrategyManager.sol";
import {ISlasher} from "@eigenlayer/contracts/interfaces/ISlasher.sol";
import {StrategyBaseTVLLimits} from "@eigenlayer/contracts/strategies/StrategyBaseTVLLimits.sol";
import "@eigenlayer/test/mocks/EmptyContract.sol";

import "@eigenlayer-middleware/src/RegistryCoordinator.sol" as regcoord;
import {IBLSApkRegistry, IIndexRegistry, IStakeRegistry} from "@eigenlayer-middleware/src/RegistryCoordinator.sol";
import {BLSApkRegistry} from "@eigenlayer-middleware/src/BLSApkRegistry.sol";
import {IndexRegistry} from "@eigenlayer-middleware/src/IndexRegistry.sol";
import {StakeRegistry} from "@eigenlayer-middleware/src/StakeRegistry.sol";
import "@eigenlayer-middleware/src/OperatorStateRetriever.sol";

import {OpenOracleServiceManager, IServiceManager} from "../src/OpenOracleServiceManager.sol";
import {OpenOracleTaskManager} from "../src/OpenOracleTaskManager.sol";
import {OpenOraclePriceFeed} from "../src/OpenOraclePriceFeed.sol";
import {IOpenOraclePriceFeed} from "../src/IOpenOraclePriceFeed.sol";
import {IOpenOracleTaskManager} from "../src/IOpenOracleTaskManager.sol";
import "../src/ERC20Mock.sol";

import {Utils} from "./utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

// # To deploy and verify our contract
// forge script script/OpenOracleDeployer.s.sol:OpenOracleDeployer --rpc-url http://127.0.0.1:8545  --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --broadcast -vvvv
contract OpenOracleDeployer is Script, Utils {
    string public deployConfigPath = string(bytes("./script/config/devnet/testnet.config.json"));

    // DEPLOYMENT CONSTANTS
    uint256 public constant QUORUM_THRESHOLD_PERCENTAGE = 100;
    uint32 public constant TASK_RESPONSE_WINDOW_BLOCK = 30;
    uint32 public constant TASK_DURATION_BLOCKS = 0;
    // TODO: right now hardcoding these (this address is anvil's default address 9)
    address public constant AGGREGATOR_ADDR =
        0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;

    // Credible Squaring contracts
    ProxyAdmin public openOracleProxyAdmin;
    PauserRegistry public openOraclePauserReg;

    regcoord.RegistryCoordinator public registryCoordinator;
    regcoord.IRegistryCoordinator public registryCoordinatorImplementation;

    IBLSApkRegistry public blsApkRegistry;
    IBLSApkRegistry public blsApkRegistryImplementation;

    IIndexRegistry public indexRegistry;
    IIndexRegistry public indexRegistryImplementation;

    IStakeRegistry public stakeRegistry;
    IStakeRegistry public stakeRegistryImplementation;

    OperatorStateRetriever public operatorStateRetriever;

    OpenOracleServiceManager public openOracleServiceManager;
    IServiceManager public openOracleServiceManagerImplementation;

    OpenOracleTaskManager public openOracleTaskManager;
    IOpenOracleTaskManager public openOracleTaskManagerImplementation;

    OpenOraclePriceFeed public openOraclePriceFeed;
    IOpenOraclePriceFeed public openOraclePriceFeedImplementation;

    function run() external {
        // Eigenlayer contracts
        string memory eigenlayerDeployedContracts = readOutput(
            "eigenlayer_deployment_output"
        );
        IStrategyManager strategyManager = IStrategyManager(
            stdJson.readAddress(
                eigenlayerDeployedContracts,
                ".addresses.strategyManager"
            )
        );
        IDelegationManager delegationManager = IDelegationManager(
            stdJson.readAddress(
                eigenlayerDeployedContracts,
                ".addresses.delegationManager"
            )
        );
        IAVSDirectory avsDirectory = IAVSDirectory(
            stdJson.readAddress(
                eigenlayerDeployedContracts,
                ".addresses.avsDirectory"
            )
        );

        address openOracleCommunityMultisig = msg.sender;
        address openOraclePauser = msg.sender;

        vm.startBroadcast();

        _deployOpenOracleContracts(
            avsDirectory,
            delegationManager,
            openOracleCommunityMultisig,
            openOraclePauser
        );
        vm.stopBroadcast();
    }

    function _deployOpenOracleContracts(
        IAVSDirectory avsDirectory,
        IDelegationManager delegationManager,
        address openOracleCommunityMultisig,
        address openOraclePauser
    ) internal {
        // READ JSON CONFIG DATA
        string memory config_data = vm.readFile(deployConfigPath);

        // parse initalization params and permissions from config data
        (
            uint96[] memory minimumStakeForQuourm, 
            IStakeRegistry.StrategyParams[][] memory strategyAndWeightingMultipliers
        ) = _parseStakeRegistryParams(config_data);

        // deploy proxy admin for ability to upgrade proxy contracts
        openOracleProxyAdmin = new ProxyAdmin();

        // deploy pauser registry
        {
            address[] memory pausers = new address[](2);
            pausers[0] = openOraclePauser;
            pausers[1] = openOracleCommunityMultisig;
            openOraclePauserReg = new PauserRegistry(
                pausers,
                openOracleCommunityMultisig
            );
        }

        EmptyContract emptyContract = new EmptyContract();

        // hard-coded inputs

        /**
         * First, deploy upgradeable proxy contracts that **will point** to the implementations. Since the implementation contracts are
         * not yet deployed, we give these proxies an empty contract as the initial implementation, to act as if they have no code.
         */
        openOracleServiceManager = OpenOracleServiceManager(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(openOracleProxyAdmin),
                    ""
                )
            )
        );
        openOracleTaskManager = OpenOracleTaskManager(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(openOracleProxyAdmin),
                    ""
                )
            )
        );
        registryCoordinator = regcoord.RegistryCoordinator(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(openOracleProxyAdmin),
                    ""
                )
            )
        );
        blsApkRegistry = IBLSApkRegistry(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(openOracleProxyAdmin),
                    ""
                )
            )
        );
        indexRegistry = IIndexRegistry(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(openOracleProxyAdmin),
                    ""
                )
            )
        );
        stakeRegistry = IStakeRegistry(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(openOracleProxyAdmin),
                    ""
                )
            )
        );

        openOraclePriceFeed = OpenOraclePriceFeed(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(openOracleProxyAdmin),
                    ""
                )
            )
        );

        operatorStateRetriever = new OperatorStateRetriever();

        // openOraclePriceFeed = new OpenOraclePriceFeed(openOracleTaskManager, 4, 1, 0);

        // Second, deploy the *implementation* contracts, using the *proxy contracts* as inputs
        {
            stakeRegistryImplementation = new StakeRegistry(
                registryCoordinator,
                delegationManager
            );

            openOracleProxyAdmin.upgrade(
                TransparentUpgradeableProxy(payable(address(stakeRegistry))),
                address(stakeRegistryImplementation)
            );

            blsApkRegistryImplementation = new BLSApkRegistry(
                registryCoordinator
            );

            openOracleProxyAdmin.upgrade(
                TransparentUpgradeableProxy(payable(address(blsApkRegistry))),
                address(blsApkRegistryImplementation)
            );

            indexRegistryImplementation = new IndexRegistry(
                registryCoordinator
            );

            openOracleProxyAdmin.upgrade(
                TransparentUpgradeableProxy(payable(address(indexRegistry))),
                address(indexRegistryImplementation)
            );
        }

        registryCoordinatorImplementation = new regcoord.RegistryCoordinator(
            regcoord.IServiceManager(address(openOracleServiceManager)),
            regcoord.IStakeRegistry(address(stakeRegistry)),
            regcoord.IBLSApkRegistry(address(blsApkRegistry)),
            regcoord.IIndexRegistry(address(indexRegistry))
        );

        {
            uint numQuorums = 1;
            // for each quorum to setup, we need to define
            // QuorumOperatorSetParam, minimumStakeForQuorum, and strategyParams
            regcoord.IRegistryCoordinator.OperatorSetParam[]
                memory quorumsOperatorSetParams = new regcoord.IRegistryCoordinator.OperatorSetParam[](
                    numQuorums
                );
            for (uint i = 0; i < numQuorums; i++) {
                // hard code these for now
                quorumsOperatorSetParams[i] = regcoord
                    .IRegistryCoordinator
                    .OperatorSetParam({
                        maxOperatorCount: 200,
                        kickBIPsOfOperatorStake: 11000,
                        kickBIPsOfTotalStake: 50
                    });
            }
            // set to 0 for every quorum
            openOracleProxyAdmin.upgradeAndCall(
                TransparentUpgradeableProxy(
                    payable(address(registryCoordinator))
                ),
                address(registryCoordinatorImplementation),
                abi.encodeWithSelector(
                    regcoord.RegistryCoordinator.initialize.selector,
                    // we set churnApprover and ejector to communityMultisig because we don't need them
                    openOracleCommunityMultisig, // owner
                    openOracleCommunityMultisig, // churner
                    openOracleCommunityMultisig, // ejector
                    openOraclePauserReg,
                    0, // 0 initialPausedStatus means everything unpaused
                    quorumsOperatorSetParams,
                    minimumStakeForQuourm,
                    strategyAndWeightingMultipliers
                )
            );
        }

        openOracleServiceManagerImplementation = new OpenOracleServiceManager(
            avsDirectory,
            registryCoordinator,
            stakeRegistry,
            openOracleTaskManager
        );
        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        openOracleProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(
                payable(address(openOracleServiceManager))
            ),
            address(openOracleServiceManagerImplementation),
            abi.encodeWithSelector(
                openOracleServiceManager.initialize.selector,
                openOracleCommunityMultisig
            )
        );

        address[] memory whitelist = new address[](1);
        whitelist[0] = 0x860B6912C2d0337ef05bbC89b0C2CB6CbAEAB4A5;
        openOracleServiceManager.addOperatorToRegistryWhitelist(whitelist);

        openOracleTaskManagerImplementation = new OpenOracleTaskManager(
            registryCoordinator,
            TASK_RESPONSE_WINDOW_BLOCK
        );

        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        openOracleProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(
                payable(address(openOracleTaskManager))
            ),
            address(openOracleTaskManagerImplementation),
            abi.encodeWithSelector(
                openOracleTaskManager.initialize.selector,
                openOraclePauserReg,
                openOracleCommunityMultisig,
                AGGREGATOR_ADDR
            )
        );

        openOraclePriceFeedImplementation = new OpenOraclePriceFeed(openOracleTaskManager);

        openOracleProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(
                payable(address(openOraclePriceFeed))
            ),
            address(openOraclePriceFeedImplementation),
            abi.encodeWithSelector(
                openOraclePriceFeed.initialize.selector,
                openOracleCommunityMultisig
            )
        );

        IOpenOracleTaskManager taskManagerInterface = IOpenOracleTaskManager(address(openOracleTaskManager));
        address feedAddress = address(openOraclePriceFeed); // Specify the address you want to add
        taskManagerInterface.addToFeedlist(feedAddress);

        // WRITE JSON DATA
        string memory parent_object = "parent object";

        string memory deployed_addresses = "addresses";
        vm.serializeAddress(
            deployed_addresses,
            "proxyAdmin",
            address(openOracleProxyAdmin)
        );
        vm.serializeAddress(
            deployed_addresses,
            "openOracleServiceManager",
            address(openOracleServiceManager)
        );
        vm.serializeAddress(
            deployed_addresses,
            "openOracleServiceManagerImplementation",
            address(openOracleServiceManagerImplementation)
        );
        vm.serializeAddress(
            deployed_addresses,
            "openOracleTaskManager",
            address(openOracleTaskManager)
        );
        vm.serializeAddress(
            deployed_addresses,
            "openOracleTaskManagerImplementation",
            address(openOracleTaskManagerImplementation)
        );
        vm.serializeAddress(
            deployed_addresses,
            "blsApkRegistry",
            address(blsApkRegistry)
        );
        vm.serializeAddress(
            deployed_addresses,
            "blsApkRegistryImplementation",
            address(blsApkRegistryImplementation)
        );
        vm.serializeAddress(
            deployed_addresses,
            "stakeRegistry",
            address(stakeRegistry)
        );
        vm.serializeAddress(
            deployed_addresses,
            "stakeRegistryImplementation",
            address(stakeRegistryImplementation)
        );
        vm.serializeAddress(
            deployed_addresses,
            "indexRegistry",
            address(indexRegistry)
        );
        vm.serializeAddress(
            deployed_addresses,
            "indexRegistryImplementation",
            address(indexRegistryImplementation)
        );
        vm.serializeAddress(
            deployed_addresses,
            "registryCoordinator",
            address(registryCoordinator)
        );
        vm.serializeAddress(
            deployed_addresses,
            "registryCoordinatorImplementation",
            address(registryCoordinatorImplementation)
        );
        vm.serializeAddress(
            deployed_addresses,
            "openOraclePriceFeed",
            address(openOraclePriceFeed)
        );
        vm.serializeAddress(
            deployed_addresses, 
            "openOraclePriceFeedImplementation", 
            address(openOraclePriceFeedImplementation)
        );
        string memory deployed_addresses_output = vm.serializeAddress(
            deployed_addresses,
            "operatorStateRetriever",
            address(operatorStateRetriever)
        );

        // serialize all the data
        string memory finalJson = vm.serializeString(
            parent_object,
            deployed_addresses,
            deployed_addresses_output
        );

        writeOutput(finalJson, "open_oracle_avs_deployment_output");
    }

    function _parseStakeRegistryParams(string memory config_data) internal pure returns (uint96[] memory minimumStakeForQuourm, IStakeRegistry.StrategyParams[][] memory strategyAndWeightingMultipliers) {
        bytes memory stakesConfigsRaw = stdJson.parseRaw(config_data, ".minimumStakes");
        minimumStakeForQuourm = abi.decode(stakesConfigsRaw, (uint96[]));
        
        bytes memory strategyConfigsRaw = stdJson.parseRaw(config_data, ".strategyWeights");
        strategyAndWeightingMultipliers = abi.decode(strategyConfigsRaw, (IStakeRegistry.StrategyParams[][]));
    }
}

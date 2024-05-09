// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import "@eigenlayer/contracts/permissions/PauserRegistry.sol";
import "@eigenlayer/test/mocks/EmptyContract.sol";

import {BLSApkRegistry} from "@eigenlayer-middleware/src/BLSApkRegistry.sol";
import {StakeRegistry} from "@eigenlayer-middleware/src/StakeRegistry.sol";
import "@eigenlayer-middleware/src/OperatorStateRetriever.sol";

import {OpenOracleTaskManager} from "../src/OpenOracleTaskManager.sol";
import {IOpenOracleTaskManager} from "../src/IOpenOracleTaskManager.sol";
import {OpenOraclePriceFeed} from "../src/OpenOraclePriceFeed.sol";
import {OpenOracleServiceManager} from "../src/OpenOracleServiceManager.sol";

import "../src/ERC20Mock.sol";

import {Utils} from "./utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

// # To deploy and verify our contract
// forge script script/OpenOracleTaskManagerDeployer.s.sol:OpenOracleTaskManagerDeployer --rpc-url http://127.0.0.1:8545  --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --broadcast -vvvv
contract OpenOracleTaskManagerDeployer is Script, Utils {

    // DEPLOYMENT CONSTANTS
    uint32 public constant TASK_RESPONSE_WINDOW_BLOCK = 30;
    uint32 public constant TASK_DURATION_BLOCKS = 0;
    // TODO: right now hardcoding these (this address is anvil's default address 9)
    address public constant AGGREGATOR_ADDR =
        0x23d8245D9fB3bC7E890E152E5b7aA0CE9E28D93b;

    // Credible Squaring contracts
    ProxyAdmin public openOracleProxyAdmin;
    PauserRegistry public openOraclePauserReg;

    OperatorStateRetriever public operatorStateRetriever;

    OpenOracleTaskManager public openOracleTaskManager;
    OpenOraclePriceFeed public openOraclePriceFeed;
    IOpenOracleTaskManager public openOracleTaskManagerImplementation;

    // OpenOracleServiceManager public openOracleServiceManager;

    function run() external {
        // Eigenlayer contracts
        string memory openOracleBridgeeployedContracts = readOutput(
            "open_oracle_bridge_registry_deployment_output"
        );
        IStakeRegistry stakeRegistry = IStakeRegistry(
            stdJson.readAddress(
                openOracleBridgeeployedContracts,
                ".addresses.stakeRegistry"
            )
        );
        IBLSApkRegistry blsApkRegistry = IBLSApkRegistry(
            stdJson.readAddress(
                openOracleBridgeeployedContracts,
                ".addresses.blsApkRegistry"
            )
        );
        openOracleProxyAdmin = ProxyAdmin(
            stdJson.readAddress(
                openOracleBridgeeployedContracts,
                ".addresses.proxyAdmin"
            )
        );
        // openOracleServiceManager = OpenOracleServiceManager(
        //     stdJson.readAddress(
        //         openOracleBridgeeployedContracts,
        //         ".addresses.openOracleServiceManager"
        //     )
        // );


        address openOracleCommunityMultisig = msg.sender;
        address openOraclePauser = msg.sender;

        vm.startBroadcast();

        _deployOpenOracleContracts(
            stakeRegistry,
            blsApkRegistry,
            openOracleCommunityMultisig,
            openOraclePauser
        );
        vm.stopBroadcast();
    }

    function _deployOpenOracleContracts(
        IStakeRegistry stakeRegistry,
        IBLSApkRegistry blsApkRegistry,
        address openOracleCommunityMultisig,
        address openOraclePauser
    ) internal {
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
        openOracleTaskManager = OpenOracleTaskManager(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(openOracleProxyAdmin),
                    ""
                )
            )
        );

        operatorStateRetriever = new OperatorStateRetriever();

        // Upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        openOracleTaskManagerImplementation = new OpenOracleTaskManager(
            stakeRegistry,
            blsApkRegistry,
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

        IOpenOracleTaskManager taskManagerInterface = IOpenOracleTaskManager(address(openOracleTaskManager));

        openOraclePriceFeed = new OpenOraclePriceFeed(openOracleTaskManager, 4, 1, 0);
        address feedAddress = address(openOraclePriceFeed);
        taskManagerInterface.addToFeedlist(feedAddress);

        // Add deployed taskManager to service manager
        // openOracleServiceManager.addTaskManager("sepolia", address(taskManagerInterface));

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
            "openOracleTaskManager",
            address(openOracleTaskManager)
        );
        vm.serializeAddress(
            deployed_addresses,
            "openOracleTaskManagerImplementation",
            address(openOracleTaskManagerImplementation)
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

        writeOutput(finalJson, "open_oracle_avs_task_manager_deployment_output");
    }
}

// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import "@eigenlayer/contracts/permissions/PauserRegistry.sol";
import "@eigenlayer/test/mocks/EmptyContract.sol";

import {BLSApkRegistry} from "@eigenlayer-middleware/src/BLSApkRegistry.sol";
import {StakeRegistry} from "@eigenlayer-middleware/src/StakeRegistry.sol";
import {IBLSApkRegistry} from "@eigenlayer-middleware/src/interfaces/IBLSApkRegistry.sol";
import {IStakeRegistry} from "@eigenlayer-middleware/src/interfaces/IStakeRegistry.sol";

import {OpenOracleTaskManager} from "../../src/OpenOracleTaskManager.sol";
import {IOpenOracleTaskManager} from "../../src/IOpenOracleTaskManager.sol";
import {OpenOracleServiceManager} from "../../src/OpenOracleServiceManager.sol";

import "../../src/ERC20Mock.sol";

import {Utils} from "../utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

// # To deploy and verify our contract
contract OpenOracleTaskManagerDeployer is Script, Utils {
    string public deployConfigPath = string.concat("./script/config/", vm.toString(block.chainid), "/config.avs.json");

    ProxyAdmin public openOracleProxyAdmin;
    PauserRegistry public openOraclePauserReg;

    OpenOracleTaskManager public openOracleTaskManager;
    IOpenOracleTaskManager public openOracleTaskManagerImplementation;

    struct DeployParams {
        uint32 taskResponseWindowBlock;
        address aggregatorAddr;
    }

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
        // READ JSON CONFIG DATA
        string memory config_data = vm.readFile(deployConfigPath);
        
        // parse initalization params and permissions from config data
        DeployParams memory deployParams = _parseDeployParams(config_data);

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

        // Upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        openOracleTaskManagerImplementation = new OpenOracleTaskManager(
            stakeRegistry,
            blsApkRegistry,
            deployParams.taskResponseWindowBlock
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
                deployParams.aggregatorAddr
            )
        );

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
        string memory deployed_addresses_output = vm.serializeAddress(
            deployed_addresses,
            "openOracleTaskManagerImplementation",
            address(openOracleTaskManagerImplementation)
        );

        // serialize all the data
        string memory finalJson = vm.serializeString(
            parent_object,
            deployed_addresses,
            deployed_addresses_output
        );

        writeOutput(finalJson, "open_oracle_avs_task_manager_deployment_output");
    }

    function _parseDeployParams(string memory config_data) internal pure returns (DeployParams memory deployParams) {
        bytes memory taskResponseWindowBlockRaw = stdJson.parseRaw(config_data, ".taskResponseWindowBlock");
        deployParams.taskResponseWindowBlock = abi.decode(taskResponseWindowBlockRaw, (uint32));
        
        bytes memory aggregatorAddrRaw = stdJson.parseRaw(config_data, ".aggregatorAddr");
        deployParams.aggregatorAddr = abi.decode(aggregatorAddrRaw, (address));
    }
}

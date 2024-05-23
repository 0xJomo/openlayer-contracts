// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import "@eigenlayer/contracts/permissions/PauserRegistry.sol";

import {OpenOracleTaskManager, IOpenOracleTaskManager} from "../../src/OpenOracleTaskManager.sol";
import {IBLSApkRegistry, IStakeRegistry} from "@eigenlayer-middleware/src/RegistryCoordinator.sol";

import {Utils} from "../utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

// Upgrade TaskManager contract
contract UpgradeTaskManager is Script, Utils {
    string public deployConfigPath = string.concat("./script/config/", vm.toString(block.chainid), "/config.avs.json");

    // DEPLOYMENT CONSTANTS
    uint32 public constant TASK_RESPONSE_WINDOW_BLOCK = 30;

    IOpenOracleTaskManager public openOracleTaskManagerImplementation;

    struct DeployParams {
        uint32 taskResponseWindowBlock;
    }

    function run() external {
        // Eigenlayer contracts
        string memory openOracleDeployedContracts = readOutput(
            "open_oracle_avs_task_manager_deployment_output"
        );


        string memory openOracleBridgeDeployedContracts = readOutput(
            "open_oracle_bridge_registry_deployment_output"
        );

        OpenOracleTaskManager openOracleTaskManager = OpenOracleTaskManager(
            stdJson.readAddress(
                openOracleDeployedContracts,
                ".addresses.openOracleTaskManager"
            )
        );
        IStakeRegistry stakeRegistry = IStakeRegistry(
            stdJson.readAddress(
                openOracleBridgeDeployedContracts,
                ".addresses.stakeRegistry"
            )
        );
        IBLSApkRegistry blsApkRegistry = IBLSApkRegistry(
            stdJson.readAddress(
                openOracleBridgeDeployedContracts,
                ".addresses.blsApkRegistry"
            )
        );
        ProxyAdmin openOracleProxyAdmin = ProxyAdmin(
            stdJson.readAddress(
                openOracleDeployedContracts,
                ".addresses.proxyAdmin"
            )
        );


        vm.startBroadcast();

        _upgradeOpenOracleContracts(
            openOracleTaskManager,
            stakeRegistry,
            blsApkRegistry,
            openOracleProxyAdmin
        );
        vm.stopBroadcast();
    }

    function _upgradeOpenOracleContracts(
        OpenOracleTaskManager openOracleTaskManager,
        IStakeRegistry stakeRegistry,
        IBLSApkRegistry blsApkRegistry,
        ProxyAdmin openOracleProxyAdmin
    ) internal {
        // READ JSON CONFIG DATA
        string memory config_data = vm.readFile(deployConfigPath);

        // parse initalization params and permissions from config data
        DeployParams memory deployParams = _parseDeployParams(config_data);

        openOracleTaskManagerImplementation = new OpenOracleTaskManager(
            stakeRegistry,
            blsApkRegistry,
            deployParams.taskResponseWindowBlock
        );

        // Upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        openOracleProxyAdmin.upgrade(
            TransparentUpgradeableProxy(
                payable(address(openOracleTaskManager))
            ),
            address(openOracleTaskManagerImplementation)
        );

        // WRITE JSON DATA
        string memory parent_object = "parent object";

        string memory deployed_addresses = "addresses";
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

        writeOutput(finalJson, "open_oracle_avs_upgrade_taskmanager_output");
    }

    function _parseDeployParams(string memory config_data) internal pure returns (DeployParams memory deployParams) {
        bytes memory taskResponseWindowBlockRaw = stdJson.parseRaw(config_data, ".taskResponseWindowBlock");
        deployParams.taskResponseWindowBlock = abi.decode(taskResponseWindowBlockRaw, (uint32));
    }
}
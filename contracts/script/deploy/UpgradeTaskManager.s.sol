// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import "@eigenlayer-middleware/src/RegistryCoordinator.sol" as regcoord;
import "@eigenlayer/contracts/permissions/PauserRegistry.sol";

import {OpenOracleTaskManager, IOpenOracleTaskManager} from "../../src/OpenOracleTaskManager.sol";

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
            "open_oracle_avs_deployment_output"
        );

        OpenOracleTaskManager openOracleTaskManager = OpenOracleTaskManager(
            stdJson.readAddress(
                openOracleDeployedContracts,
                ".addresses.openOracleTaskManager"
            )
        );
        regcoord.RegistryCoordinator registryCoordinator = regcoord.RegistryCoordinator(
            stdJson.readAddress(
                openOracleDeployedContracts,
                ".addresses.registryCoordinator"
            )
        );
        ProxyAdmin openOracleProxyAdmin = ProxyAdmin(
            stdJson.readAddress(
                openOracleDeployedContracts,
                ".addresses.proxyAdmin"
            )
        );
        PauserRegistry openOraclePauserReg = PauserRegistry(
            stdJson.readAddress(
                openOracleDeployedContracts,
                ".addresses.pauserRegistry"
            )
        );

        address openOracleCommunityMultisig = msg.sender;
        address openOraclePauser = msg.sender;

        vm.startBroadcast();

        _upgradeOpenOracleContracts(
            openOracleTaskManager,
            registryCoordinator,
            openOracleProxyAdmin,
            openOraclePauserReg,
            openOracleCommunityMultisig,
            openOraclePauser
        );
        vm.stopBroadcast();
    }

    function _upgradeOpenOracleContracts(
        OpenOracleTaskManager openOracleTaskManager,
        regcoord.RegistryCoordinator registryCoordinator,
        ProxyAdmin openOracleProxyAdmin,
        PauserRegistry openOraclePauserReg,
        address openOracleCommunityMultisig,
        address openOraclePauser
    ) internal {
        // READ JSON CONFIG DATA
        string memory config_data = vm.readFile(deployConfigPath);

        // parse initalization params and permissions from config data
        DeployParams memory deployParams = _parseDeployParams(config_data);

        openOracleTaskManagerImplementation = new OpenOracleTaskManager(
            registryCoordinator,
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
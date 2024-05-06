// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import {IAVSDirectory} from "@eigenlayer/contracts/interfaces/IAVSDirectory.sol";
import "@eigenlayer-middleware/src/RegistryCoordinator.sol" as regcoord;
import {StakeRegistry} from "@eigenlayer-middleware/src/StakeRegistry.sol";

import {OpenOracleTaskManager, IOpenOracleTaskManager} from "../../src/OpenOracleTaskManager.sol";
import {OpenOracleServiceManager, IServiceManager} from "../../src/OpenOracleServiceManager.sol";

import {Utils} from "../utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

// Upgrade ServiceManager contract
contract UpgradeServiceManager is Script, Utils {
    IServiceManager public openOracleServiceManagerImplementation;

    function run() external {
        // Eigenlayer contracts
        string memory openOracleDeployedContracts = readOutput(
            "open_oracle_avs_deployment_output"
        );

        OpenOracleServiceManager openOracleServiceManager = OpenOracleServiceManager(
            stdJson.readAddress(
                openOracleDeployedContracts,
                ".addresses.openOracleServiceManager"
            )
        );
        regcoord.RegistryCoordinator registryCoordinator = regcoord.RegistryCoordinator(
            stdJson.readAddress(
                openOracleDeployedContracts,
                ".addresses.registryCoordinator"
            )
        );
        OpenOracleTaskManager openOracleTaskManager = OpenOracleTaskManager(
            stdJson.readAddress(
                openOracleDeployedContracts,
                ".addresses.openOracleTaskManager"
            )
        );
        StakeRegistry stakeRegistry = StakeRegistry(
            stdJson.readAddress(
                openOracleDeployedContracts,
                ".addresses.stakeRegistry"
            )
        );

        ProxyAdmin openOracleProxyAdmin = ProxyAdmin(
            stdJson.readAddress(
                openOracleDeployedContracts,
                ".addresses.proxyAdmin"
            )
        );

        // Eigenlayer contracts
        string memory eigenlayerDeployedContracts = readOutput(
            "eigenlayer_deployment_output"
        );
        IAVSDirectory avsDirectory = IAVSDirectory(
            stdJson.readAddress(
                eigenlayerDeployedContracts,
                ".addresses.avsDirectory"
            )
        );

        vm.startBroadcast();
        // Start upgrade

        openOracleServiceManagerImplementation = new OpenOracleServiceManager(
            avsDirectory,
            registryCoordinator,
            stakeRegistry,
            openOracleTaskManager
        );

        // Upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        openOracleProxyAdmin.upgrade(
            TransparentUpgradeableProxy(
                payable(address(openOracleServiceManager))
            ),
            address(openOracleServiceManagerImplementation)
        );

        // WRITE JSON DATA
        string memory parent_object = "parent object";

        string memory deployed_addresses = "addresses";
        vm.serializeAddress(
            deployed_addresses,
            "openOracleServiceManager",
            address(openOracleServiceManager)
        );
        string memory deployed_addresses_output = vm.serializeAddress(
            deployed_addresses,
            "openOracleServiceManagerImplementation",
            address(openOracleServiceManagerImplementation)
        );

        // serialize all the data
        string memory finalJson = vm.serializeString(
            parent_object,
            deployed_addresses,
            deployed_addresses_output
        );

        writeOutput(finalJson, "open_oracle_avs_upgrade_servicemanager_output");

        // End upgrade
        vm.stopBroadcast();
    }
}
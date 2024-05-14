// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import "@eigenlayer-middleware/src/RegistryCoordinator.sol";
import {Utils} from "../utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

// Upgrade ServiceManager contract
contract UpgradeBLSApkRegistry is Script, Utils {
    RegistryCoordinator public registryCoordinatorImplementation;

    function run() external {

        string memory openOracleDeployedContracts = readOutput(
            "open_oracle_avs_deployment_output"
        );

        RegistryCoordinator registryCoordinator = RegistryCoordinator(
            stdJson.readAddress(
                openOracleDeployedContracts,
                ".addresses.registryCoordinator"
            )
        );

        IServiceManager serviceManager = IServiceManager(
            stdJson.readAddress(
                openOracleDeployedContracts,
                ".addresses.openOracleServiceManager"
            )
        );

        IStakeRegistry stakeRegistry = IStakeRegistry(
            stdJson.readAddress(
                openOracleDeployedContracts,
                ".addresses.stakeRegistry"
            )
        );

        IBLSApkRegistry bLSApkRegistry = IBLSApkRegistry(
            stdJson.readAddress(
                openOracleDeployedContracts,
                ".addresses.blsApkRegistry"
            )
        );

        IIndexRegistry indexRegistry = IIndexRegistry(
            stdJson.readAddress(
                openOracleDeployedContracts,
                ".addresses.indexRegistry"
            )
        );

        ProxyAdmin openOracleProxyAdmin = ProxyAdmin(
            stdJson.readAddress(
                openOracleDeployedContracts,
                ".addresses.proxyAdmin"
            )
        );

        vm.startBroadcast();
        // Start upgrade

        registryCoordinatorImplementation = new RegistryCoordinator(
            serviceManager,
            stakeRegistry,
            bLSApkRegistry,
            indexRegistry
        );

        // Upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        openOracleProxyAdmin.upgrade(
            TransparentUpgradeableProxy(
                payable(address(registryCoordinator))
            ),
            address(registryCoordinatorImplementation)
        );

        // WRITE JSON DATA
        string memory parent_object = "parent object";

        string memory deployed_addresses = "addresses";
        vm.serializeAddress(
            deployed_addresses,
            "registryCoordinator",
            address(registryCoordinator)
        );
        string memory deployed_addresses_output = vm.serializeAddress(
            deployed_addresses,
            "registryCoordinatorImplementation",
            address(registryCoordinatorImplementation)
        );

        // serialize all the data
        string memory finalJson = vm.serializeString(
            parent_object,
            deployed_addresses,
            deployed_addresses_output
        );

        writeOutput(finalJson, "open_oracle_avs_upgrade_registry_coordinator_output");

        // End upgrade
        vm.stopBroadcast();
    }
}
// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import {IBLSApkRegistry} from "@eigenlayer-middleware/src/RegistryCoordinator.sol";

import {OpenOracleBridgeBLSApkRegistry} from "../../src/bridge/OpenOracleBridgeBLSApkRegistry.sol";
import {OpenOracleBridgeRegistryCoordinator} from "../../src/bridge/OpenOracleBridgeRegistryCoordinator.sol";
import {Utils} from "../utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

// Upgrade ServiceManager contract
contract UpgradeBLSApkRegistry is Script, Utils {
    IBLSApkRegistry public bLSApkRegistryImplementation;

    function run() external {

        string memory openOracleDeployedContracts = readOutput(
            "open_oracle_bridge_registry_deployment_output"
        );


        OpenOracleBridgeRegistryCoordinator registryCoordinator = OpenOracleBridgeRegistryCoordinator(
            stdJson.readAddress(
                openOracleDeployedContracts,
                ".addresses.registryCoordinator"
            )
        );

        OpenOracleBridgeBLSApkRegistry bLSApkRegistry = OpenOracleBridgeBLSApkRegistry(
            stdJson.readAddress(
                openOracleDeployedContracts,
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
        // Start upgrade

        bLSApkRegistryImplementation = new OpenOracleBridgeBLSApkRegistry(
            registryCoordinator
        );

        // Upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        openOracleProxyAdmin.upgrade(
            TransparentUpgradeableProxy(
                payable(address(bLSApkRegistry))
            ),
            address(bLSApkRegistryImplementation)
        );

        // WRITE JSON DATA
        string memory parent_object = "parent object";

        string memory deployed_addresses = "addresses";
        vm.serializeAddress(
            deployed_addresses,
            "bLSApkRegistry",
            address(bLSApkRegistry)
        );
        string memory deployed_addresses_output = vm.serializeAddress(
            deployed_addresses,
            "bLSApkRegistryImplementation",
            address(bLSApkRegistryImplementation)
        );

        // serialize all the data
        string memory finalJson = vm.serializeString(
            parent_object,
            deployed_addresses,
            deployed_addresses_output
        );

        writeOutput(finalJson, "open_oracle_avs_upgrade_bridge_bls_apk_registry_output");

        // End upgrade
        vm.stopBroadcast();
    }
}
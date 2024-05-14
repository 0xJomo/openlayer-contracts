// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import "@eigenlayer-middleware/src/RegistryCoordinator.sol" as regcoord;
import {IndexRegistry} from "@eigenlayer-middleware/src/IndexRegistry.sol";
import {Utils} from "../utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

// Upgrade ServiceManager contract
contract UpgradeIndexRegistry is Script, Utils {
    IndexRegistry public indexRegistryImplementation;

    function run() external {

        string memory openOracleDeployedContracts = readOutput(
            "open_oracle_avs_deployment_output"
        );


        regcoord.RegistryCoordinator registryCoordinator = regcoord.RegistryCoordinator(
            stdJson.readAddress(
                openOracleDeployedContracts,
                ".addresses.registryCoordinator"
            )
        );

        IndexRegistry indexRegistry = IndexRegistry(
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

        indexRegistryImplementation = new IndexRegistry(
            registryCoordinator
        );

        // Upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        openOracleProxyAdmin.upgrade(
            TransparentUpgradeableProxy(
                payable(address(indexRegistry))
            ),
            address(indexRegistryImplementation)
        );

        // WRITE JSON DATA
        string memory parent_object = "parent object";

        string memory deployed_addresses = "addresses";
        vm.serializeAddress(
            deployed_addresses,
            "indexRegistry",
            address(indexRegistry)
        );
        string memory deployed_addresses_output = vm.serializeAddress(
            deployed_addresses,
            "indexRegistryImplementation",
            address(indexRegistryImplementation)
        );

        // serialize all the data
        string memory finalJson = vm.serializeString(
            parent_object,
            deployed_addresses,
            deployed_addresses_output
        );

        writeOutput(finalJson, "open_oracle_avs_upgrade_index_registry_output");

        // End upgrade
        vm.stopBroadcast();
    }
}
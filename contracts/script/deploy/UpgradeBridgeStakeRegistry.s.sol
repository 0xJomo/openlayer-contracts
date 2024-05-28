// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import {IStakeRegistry} from "@eigenlayer-middleware/src/RegistryCoordinator.sol";

import {OpenOracleBridgeStakeRegistry} from "../../src/bridge/OpenOracleBridgeStakeRegistry.sol";
import {OpenOracleBridgeRegistryCoordinator} from "../../src/bridge/OpenOracleBridgeRegistryCoordinator.sol";
import {Utils} from "../utils/Utils.sol";
import {OpenOracleBridgeDelegationManager} from "../../src/bridge/OpenOracleBridgeDelegationManager.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

// Upgrade ServiceManager contract
contract UpgradeStakeRegistry is Script, Utils {
    IStakeRegistry public stakeRegistryImplementation;

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

        OpenOracleBridgeStakeRegistry stakeRegistry = OpenOracleBridgeStakeRegistry(
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

        OpenOracleBridgeDelegationManager delegationManager = OpenOracleBridgeDelegationManager(
            stdJson.readAddress(
                openOracleDeployedContracts,
                ".addresses.delegationManager"
            )
        );

        vm.startBroadcast();
        // Start upgrade

        stakeRegistryImplementation = new OpenOracleBridgeStakeRegistry(
            registryCoordinator,
            delegationManager
        );

        // Upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        openOracleProxyAdmin.upgrade(
            TransparentUpgradeableProxy(
                payable(address(stakeRegistry))
            ),
            address(stakeRegistryImplementation)
        );

        // WRITE JSON DATA
        string memory parent_object = "parent object";

        string memory deployed_addresses = "addresses";
        vm.serializeAddress(
            deployed_addresses,
            "stakeRegistry",
            address(stakeRegistry)
        );
        string memory deployed_addresses_output = vm.serializeAddress(
            deployed_addresses,
            "stakeRegistryImplementation",
            address(stakeRegistryImplementation)
        );

        // serialize all the data
        string memory finalJson = vm.serializeString(
            parent_object,
            deployed_addresses,
            deployed_addresses_output
        );

        writeOutput(finalJson, "open_oracle_avs_upgrade_bridge_stake_registry_output");

        // End upgrade
        vm.stopBroadcast();
    }
}
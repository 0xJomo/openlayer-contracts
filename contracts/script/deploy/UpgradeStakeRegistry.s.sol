// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import "@eigenlayer-middleware/src/RegistryCoordinator.sol" as regcoord;
import {StakeRegistry} from "@eigenlayer-middleware/src/StakeRegistry.sol";
import {Utils} from "../utils/Utils.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

// Upgrade ServiceManager contract
contract UpgradeStakeRegistry is Script, Utils {
    StakeRegistry public stakeRegistryImplementation;

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

        string memory eigenLayerDeployedContracts = readOutput(
            "eigenlayer_deployment_output"
        );

        IDelegationManager delegationManager = IDelegationManager(
            stdJson.readAddress(
                eigenLayerDeployedContracts,
                ".addresses.delegationManager"
            )
        );

        vm.startBroadcast();
        // Start upgrade

        stakeRegistryImplementation = new StakeRegistry(
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

        writeOutput(finalJson, "open_oracle_avs_upgrade_stake_registry_output");

        // End upgrade
        vm.stopBroadcast();
    }
}
// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import "@eigenlayer/test/mocks/EmptyContract.sol";

import "@eigenlayer-middleware/src/RegistryCoordinator.sol" as regcoord;
import {IBLSApkRegistry, IStakeRegistry} from "@eigenlayer-middleware/src/RegistryCoordinator.sol";

import {OpenOracleBridgeBLSApkRegistry} from "../../src/bridge/OpenOracleBridgeBLSApkRegistry.sol";
import {OpenOracleBridgeDelegationManager} from "../../src/bridge/OpenOracleBridgeDelegationManager.sol";
import {OpenOracleBridgeRegistryCoordinator} from "../../src/bridge/OpenOracleBridgeRegistryCoordinator.sol";
import {OpenOracleBridgeStakeRegistry} from "../../src/bridge/OpenOracleBridgeStakeRegistry.sol";

import "../../src/ERC20Mock.sol";

import {Utils} from "../utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

// # To deploy and verify our contract
contract OpenOracleBridgeDeployer is Script, Utils {
    ProxyAdmin public openOracleProxyAdmin;

    OpenOracleBridgeRegistryCoordinator public registryCoordinator;
    regcoord.IRegistryCoordinator public registryCoordinatorImplementation;

    OpenOracleBridgeBLSApkRegistry public blsApkRegistry;
    IBLSApkRegistry public blsApkRegistryImplementation;

    OpenOracleBridgeStakeRegistry public stakeRegistry;
    IStakeRegistry public stakeRegistryImplementation;
    
    OpenOracleBridgeDelegationManager public delegationManager;
    IDelegationManager public delegationManagerImplementation;

    function run() external {
        address openOracleCommunityMultisig = msg.sender;

        vm.startBroadcast();

        _deployOpenOracleContracts(
            openOracleCommunityMultisig
        );
        vm.stopBroadcast();
    }

    function _deployOpenOracleContracts(
        address openOracleCommunityMultisig
    ) internal {
        // deploy proxy admin for ability to upgrade proxy contracts
        openOracleProxyAdmin = new ProxyAdmin();

        EmptyContract emptyContract = new EmptyContract();

        // hard-coded inputs

        /**
         * First, deploy upgradeable proxy contracts that **will point** to the implementations. Since the implementation contracts are
         * not yet deployed, we give these proxies an empty contract as the initial implementation, to act as if they have no code.
         */
        registryCoordinator = OpenOracleBridgeRegistryCoordinator(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(openOracleProxyAdmin),
                    ""
                )
            )
        );
        blsApkRegistry = OpenOracleBridgeBLSApkRegistry(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(openOracleProxyAdmin),
                    ""
                )
            )
        );
        stakeRegistry = OpenOracleBridgeStakeRegistry(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(openOracleProxyAdmin),
                    ""
                )
            )
        );
        delegationManager = OpenOracleBridgeDelegationManager(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(openOracleProxyAdmin),
                    ""
                )
            )
        );

        // Second, deploy the *implementation* contracts, using the *proxy contracts* as inputs
        {
            stakeRegistryImplementation = new OpenOracleBridgeStakeRegistry(
                registryCoordinator,
                delegationManager
            );

            openOracleProxyAdmin.upgradeAndCall(
                TransparentUpgradeableProxy(payable(address(stakeRegistry))),
                address(stakeRegistryImplementation),
                abi.encodeWithSelector(
                    stakeRegistry.initialize.selector,
                    openOracleCommunityMultisig
                )
            );

            blsApkRegistryImplementation = new OpenOracleBridgeBLSApkRegistry(
                registryCoordinator
            );

            openOracleProxyAdmin.upgradeAndCall(
                TransparentUpgradeableProxy(payable(address(blsApkRegistry))),
                address(blsApkRegistryImplementation),
                abi.encodeWithSelector(
                    blsApkRegistry.initialize.selector,
                    openOracleCommunityMultisig
                )
            );

            delegationManagerImplementation = new OpenOracleBridgeDelegationManager();

            openOracleProxyAdmin.upgradeAndCall(
                TransparentUpgradeableProxy(payable(address(delegationManager))),
                address(delegationManagerImplementation),
                abi.encodeWithSelector(
                    delegationManager.initialize.selector,
                    openOracleCommunityMultisig
                )
            );

            registryCoordinatorImplementation = new OpenOracleBridgeRegistryCoordinator(
                stakeRegistry,
                blsApkRegistry
            );

            openOracleProxyAdmin.upgradeAndCall(
                TransparentUpgradeableProxy(payable(address(registryCoordinator))),
                address(registryCoordinatorImplementation),
                abi.encodeWithSelector(
                    registryCoordinator.initialize.selector,
                    openOracleCommunityMultisig
                )
            );
        }

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
            "blsApkRegistry",
            address(blsApkRegistry)
        );
        vm.serializeAddress(
            deployed_addresses,
            "blsApkRegistryImplementation",
            address(blsApkRegistryImplementation)
        );
        vm.serializeAddress(
            deployed_addresses,
            "stakeRegistry",
            address(stakeRegistry)
        );
        vm.serializeAddress(
            deployed_addresses,
            "stakeRegistryImplementation",
            address(stakeRegistryImplementation)
        );
        vm.serializeAddress(
            deployed_addresses,
            "registryCoordinator",
            address(registryCoordinator)
        );
        vm.serializeAddress(
            deployed_addresses,
            "registryCoordinatorImplementation",
            address(registryCoordinatorImplementation)
        );
        vm.serializeAddress(
            deployed_addresses,
            "delegationManager",
            address(delegationManager)
        );
        string memory deployed_addresses_output = vm.serializeAddress(
            deployed_addresses,
            "delegationManagerImplementation",
            address(delegationManagerImplementation)
        );

        // serialize all the data
        string memory finalJson = vm.serializeString(
            parent_object,
            deployed_addresses,
            deployed_addresses_output
        );

        writeOutput(finalJson, "open_oracle_bridge_registry_deployment_output");
    }
}

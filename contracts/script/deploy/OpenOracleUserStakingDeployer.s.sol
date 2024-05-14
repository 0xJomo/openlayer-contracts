// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "forge-std/Script.sol";
import "../../src/OpenOracleUserStaking.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {Utils} from "../utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

contract OpenOracleUserStakingDeployScript is Script, Utils {
    function run() external {

        string memory parent_object = "parent object";
        string memory deployed_addresses = "addresses";
        // string memory openOracleDeployedContracts = readOutput(
        //     "open_oracle_avs_deployment_output"
        // );

        ProxyAdmin proxyAdmin = new ProxyAdmin();
        // ProxyAdmin proxyAdmin = ProxyAdmin(
        //     stdJson.readAddress(
        //         openOracleDeployedContracts,
        //         ".addresses.proxyAdmin"
        //     )
        // );
        vm.startBroadcast();
        OpenOracleUserStaking stakeContract = new OpenOracleUserStaking();

        vm.serializeAddress(
            deployed_addresses,
            "proxyAdmin",
            address(proxyAdmin)
        );

        string memory deployed_addresses_impl_output = vm.serializeAddress(
            deployed_addresses,
            "openOracleUserStakingImplementation",
            address(stakeContract)
        );

        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(stakeContract), address(proxyAdmin), abi.encodeWithSelector(stakeContract.initialize.selector)
        );
        vm.serializeString(parent_object, deployed_addresses, deployed_addresses_impl_output);
        string memory deployed_addresses_output = vm.serializeAddress(
            deployed_addresses,
            "openOracleUserStaking",
            address(proxy)
        );

        vm.serializeString(parent_object, deployed_addresses, deployed_addresses_output);
        string memory chain_info = "chainInfo";
        vm.serializeUint(chain_info, "deploymentBlock", block.number);
        string memory chain_info_output = vm.serializeUint(chain_info, "chainId", block.chainid);
        string memory finalJson = vm.serializeString(parent_object, chain_info, chain_info_output);
        writeOutput(finalJson, "open_oracle_avs_stake_deployment_output");
        vm.stopBroadcast();
    }
}

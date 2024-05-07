// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "forge-std/Script.sol";
import "../src/OpenOracleStake.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract OpenOracleStakeScript is Script {

    function run() external {
        vm.startBroadcast();
        OpenOracleStake stakeContract = new OpenOracleStake();
        console2.log("OpenOracleStake implement deployed on %s", address(stakeContract));
        ProxyAdmin proxyAdmin = new ProxyAdmin();
        console2.log("ProxyAdmin deployed on %s", address(proxyAdmin));
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(address(stakeContract), address(proxyAdmin), abi.encodeWithSelector(stakeContract.initialize.selector));
        console2.log("OpenOracleStake deployed on %s", address(proxy));
        OpenOracleStake stakeProxyContract = OpenOracleStake(address(proxy));
        console2.log("view owner %s", stakeProxyContract.owner());
        vm.stopBroadcast();
    }
}

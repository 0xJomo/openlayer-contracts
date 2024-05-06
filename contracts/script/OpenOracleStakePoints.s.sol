// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "forge-std/Script.sol";
import "../src/OpenOracleStakePoints.sol";

contract OpenOracleStakePointsScript is Script {

    function run() external {
        vm.startBroadcast();
        OpenOracleStakePoints stakeContract = new OpenOracleStakePoints();
        console2.log("OpenOracleStakePoints deployed on %s", address(stakeContract));
        vm.stopBroadcast();
    }
}

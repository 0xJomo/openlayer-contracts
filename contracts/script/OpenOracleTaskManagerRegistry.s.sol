// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import {IOpenOracleTaskManager} from "../src/IOpenOracleTaskManager.sol";
import {OpenOracleServiceManager} from "../src/OpenOracleServiceManager.sol";


import {Utils} from "./utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

// # To deploy and verify our contract
// forge script script/OpenOracleTaskManagerRegistry.s.sol:OpenOracleTaskManagerRegistry --rpc-url http://127.0.0.1:8545  --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --broadcast -vvvv
contract OpenOracleTaskManagerRegistry is Script, Utils {

    OpenOracleServiceManager public openOracleServiceManager;

    function run() external {
        string memory openOracleTaskManagerdeployedContracts = readOutput(
            "open_oracle_avs_task_manager_deployment_output"
        );
        IOpenOracleTaskManager openOracleTaskManager = IOpenOracleTaskManager(
            stdJson.readAddress(
                openOracleTaskManagerdeployedContracts,
                ".addresses.openOracleTaskManager"
            )
        );
        string memory openOracleBAvsDeployedContracts = readOutput(
            "open_oracle_avs_deployment_output"
        );
        openOracleServiceManager = OpenOracleServiceManager(
            stdJson.readAddress(
                openOracleBAvsDeployedContracts,
                ".addresses.openOracleServiceManager"
            )
        );
        openOracleServiceManager.addTaskManager("sepolia", address(openOracleTaskManager));
    }
}

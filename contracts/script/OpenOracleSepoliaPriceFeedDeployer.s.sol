// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import {OpenOracleTaskManager} from "../src/OpenOracleTaskManager.sol";
import {OpenOraclePriceFeed} from "../src/OpenOraclePriceFeed.sol";

import {Utils} from "./utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

// # To deploy and verify our contract
// forge script script/OpenOracleDeployer.s.sol:OpenOracleDeployer --rpc-url http://127.0.0.1:8545  --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --broadcast -vvvv
contract OpenOracleSepoliaPriceFeedDeployer is Script, Utils {
    string public deployConfigPath = string(bytes("./script/config/sepolia/testnet.config.json"));

    uint8 taskTypeLower = 5;
    uint8 taskTypeUpper = 13;
    uint8 responderThreshold = 3;
    uint96 stakeThreshold = 10000000000000000;

    function toString(uint8 value) internal pure returns (string memory) {
        if (value == 0) {
            return "0";
        }
 
        uint8 temp = value;
        uint8 digits;
 
        while (temp != 0) {
            digits++;
            temp /= 10;
        }
 
        bytes memory buffer = new bytes(digits);
 
        while (value != 0) {
            digits--;
            buffer[digits] = bytes1(uint8(48 + (value % 10)));
            value /= 10;
        }
 
        return string(buffer);
    }

    function run() external {
        // Eigenlayer contracts
        string memory openOracleDeployedContracts = readOutput(
            "open_oracle_avs_task_manager_deployment_output"
        );

        OpenOracleTaskManager openOracleTaskManager = OpenOracleTaskManager(
            stdJson.readAddress(
                openOracleDeployedContracts,
                ".addresses.openOracleTaskManager"
            )
        );

        vm.startBroadcast();

        // WRITE JSON DATA
        string memory parent_object = "parent object";
        string memory deployed_addresses = "addresses";

        for (uint8 i = taskTypeLower; i <= taskTypeUpper; i++) {
            OpenOraclePriceFeed openOraclePriceFeed = new OpenOraclePriceFeed(openOracleTaskManager, i, responderThreshold, stakeThreshold);
            address feedAddress = address(openOraclePriceFeed); // Specify the address you want to add
            openOracleTaskManager.addToFeedlist(feedAddress);
            
            vm.serializeAddress(
                deployed_addresses,
                toString(i),
                address(openOraclePriceFeed)
            );
        }
        string memory deployed_addresses_output = vm.serializeAddress(
            deployed_addresses,
            "openOracleTaskManager",
            address(openOracleTaskManager)
        );

        // serialize all the data
        string memory finalJson = vm.serializeString(
            parent_object,
            deployed_addresses,
            deployed_addresses_output
        );

        writeOutput(finalJson, "open_oracle_avs_pricefeeds_output");

        vm.stopBroadcast();
    }
}
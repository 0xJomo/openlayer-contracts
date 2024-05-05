// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import {OpenOracleTaskManager} from "../../src/OpenOracleTaskManager.sol";
import {OpenOraclePriceFeed} from "../../src/OpenOraclePriceFeed.sol";

import {Utils} from "../utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

// Deploy pricefeeds
contract OpenOracleHoleskyPriceFeedsDeployer is Script, Utils {
    string public deployConfigPath = string.concat("./script/config/", vm.toString(block.chainid), "/config.pricefeeds.json");

    struct DeployParams {
        uint8 taskTypeLower;
        uint8 taskTypeUpper;
        uint8 responderThreshold;
        uint96 stakeThreshold;
    }

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
            "open_oracle_avs_deployment_output"
        );

        OpenOracleTaskManager openOracleTaskManager = OpenOracleTaskManager(
            stdJson.readAddress(
                openOracleDeployedContracts,
                ".addresses.openOracleTaskManager"
            )
        );

        // READ JSON CONFIG DATA
        string memory config_data = vm.readFile(deployConfigPath);
        
        // parse initalization params and permissions from config data
        DeployParams memory deployParams = _parseDeployParams(config_data);

        vm.startBroadcast();

        // WRITE JSON DATA
        string memory parent_object = "parent object";
        string memory deployed_addresses = "addresses";

        for (uint8 i = deployParams.taskTypeLower; i <= deployParams.taskTypeUpper; i++) {
            OpenOraclePriceFeed openOraclePriceFeed = new OpenOraclePriceFeed(openOracleTaskManager, i, deployParams.responderThreshold, deployParams.stakeThreshold);
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

        writeOutput(finalJson, "open_oracle_avs_pricefeeds_output_v2");

        vm.stopBroadcast();
    }

    function _parseDeployParams(string memory config_data) internal pure returns (DeployParams memory deployParams) {
        bytes memory taskTypeLowerRaw = stdJson.parseRaw(config_data, ".taskTypeLower");
        deployParams.taskTypeLower = abi.decode(taskTypeLowerRaw, (uint8));

        bytes memory taskTypeUpperRaw = stdJson.parseRaw(config_data, ".taskTypeUpper");
        deployParams.taskTypeUpper = abi.decode(taskTypeUpperRaw, (uint8));

        bytes memory responderThresholdRaw = stdJson.parseRaw(config_data, ".responderThreshold");
        deployParams.responderThreshold = abi.decode(responderThresholdRaw, (uint8));

        bytes memory stakeThresholdRaw = stdJson.parseRaw(config_data, ".stakeThreshold");
        deployParams.stakeThreshold = abi.decode(stakeThresholdRaw, (uint96));
    }
}
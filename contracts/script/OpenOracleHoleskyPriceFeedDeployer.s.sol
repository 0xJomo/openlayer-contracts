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
contract OpenOracleHoleskyPriceFeedDeployer is Script, Utils {
    string public deployConfigPath = string(bytes("./script/config/holesky/testnet.config.json"));

    uint8 taskType = 1;
    uint8 responderThreshold = 1;
    uint96 stakeThreshold = 0;

    OpenOraclePriceFeed public openOraclePriceFeed;

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

        vm.startBroadcast();

        openOraclePriceFeed = new OpenOraclePriceFeed(openOracleTaskManager, taskType, responderThreshold, stakeThreshold);
        address feedAddress = address(openOraclePriceFeed); // Specify the address you want to add
        openOracleTaskManager.addToFeedlist(feedAddress);

        // WRITE JSON DATA
        string memory parent_object = "parent object";

        string memory deployed_addresses = "addresses";
        vm.serializeAddress(
            deployed_addresses,
            "openOracleTaskManager",
            address(openOracleTaskManager)
        );
        string memory deployed_addresses_output = vm.serializeAddress(
            deployed_addresses,
            "openOraclePriceFeed",
            address(openOraclePriceFeed)
        );

        // serialize all the data
        string memory finalJson = vm.serializeString(
            parent_object,
            deployed_addresses,
            deployed_addresses_output
        );

        writeOutput(finalJson, "open_oracle_avs_pricefeed_output");

        vm.stopBroadcast();
    }
}
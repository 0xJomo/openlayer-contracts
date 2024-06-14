// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@eigenlayer/test/mocks/EmptyContract.sol";

import {OpenOracleCommonDataFeed} from "../../src/OpenOracleCommonDataFeed.sol";
import {IOpenOracleCommonDataFeed} from "../../src/IOpenOracleCommonDataFeed.sol";

import {OpenOracleTaskManager} from "../../src/OpenOracleTaskManager.sol";

import {Utils} from "../utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

// Deploy pricefeeds
contract OpenOracleCommonDataFeedsDeployer is Script, Utils {
    string public deployConfigPath = string.concat("./script/config/", vm.toString(block.chainid), "/config.pricefeeds.json");

    ProxyAdmin public openOracleProxyAdmin;

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

        vm.startBroadcast();

        _deployOpenOraclePriceFeeds(
            openOracleTaskManager,
            msg.sender
        );
        vm.stopBroadcast();
    }

    function _deployOpenOraclePriceFeeds(
        OpenOracleTaskManager openOracleTaskManager,
        address openOracleCommunityMultisig
    ) internal {
        // READ JSON CONFIG DATA
        string memory config_data = vm.readFile(deployConfigPath);

        // parse initalization params and permissions from config data
        DeployParams memory deployParams = _parseDeployParams(config_data);

        // deploy proxy admin for ability to upgrade proxy contracts
        openOracleProxyAdmin = new ProxyAdmin();

        // WRITE JSON DATA
        string memory parent_object = "parent object";
        string memory deployed_addresses = "addresses";
        string memory implementation_addresses = "impl_addresses";


        OpenOracleCommonDataFeed openOraclePriceFeedImplementation = new OpenOracleCommonDataFeed(
                openOracleTaskManager
            );


        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(openOraclePriceFeedImplementation), address(openOracleProxyAdmin), abi.encodeWithSelector(
                    openOraclePriceFeedImplementation.initialize.selector,
                    msg.sender,
                    deployParams.responderThreshold,
                    deployParams.stakeThreshold
                )
        );

        vm.serializeAddress(
            deployed_addresses,
            "openOracleDataFeed",
            address(proxy)
        );
        vm.serializeAddress(
            implementation_addresses,
            "openOracleDataFeedImplementation",
            address(openOraclePriceFeedImplementation)
        );

        openOracleTaskManager.addToFeedlist(address(proxy));


        string memory deployed_addresses_output = vm.serializeAddress(
            deployed_addresses,
            "openOracleTaskManager",
            address(openOracleTaskManager)
        );
        string memory implementation_addresses_output = vm.serializeAddress(
            implementation_addresses,
            "proxyAdmin",
            address(openOracleProxyAdmin)
        );

        // serialize all the data
        vm.serializeString(
            parent_object,
            implementation_addresses,
            implementation_addresses_output
        );
        string memory finalJson = vm.serializeString(
            parent_object,
            deployed_addresses,
            deployed_addresses_output
        );

        writeOutput(finalJson, "open_oracle_avs_data_feeds_output");
    }

    function _parseDeployParams(string memory config_data) internal pure returns (DeployParams memory deployParams) {
        bytes memory taskTypeLowerRaw = stdJson.parseRaw(config_data, ".taskTypeLower");
        deployParams.taskTypeLower = abi.decode(taskTypeLowerRaw, (uint8));
    }
}

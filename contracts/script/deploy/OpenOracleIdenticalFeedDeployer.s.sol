// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {OpenOracleIdenticalAnswerTaskManager} from "../../src/OpenOracleIdenticalAnswerTaskManager.sol";
import {OpenOracleVRFFeed} from "../../src/OpenOracleVRFFeed.sol";
import {BLSApkRegistry} from "@eigenlayer-middleware/src/BLSApkRegistry.sol";
import {StakeRegistry} from "@eigenlayer-middleware/src/StakeRegistry.sol";
import "@eigenlayer/contracts/permissions/PauserRegistry.sol";
import {BLSSignatureChecker,IRegistryCoordinator} from "@eigenlayer-middleware/src/BLSSignatureChecker.sol";

import {Utils} from "../utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

// Deploy pricefeeds
contract OpenOracleIdenticalFeedDeployer is Script, Utils {
    string public deployConfigPath = string.concat("./script/config/", vm.toString(block.chainid), "/config.pricefeeds.json");

    OpenOracleIdenticalAnswerTaskManager openOracleIdenticalAnswerTaskManager;

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

        ProxyAdmin proxyAdmin = ProxyAdmin(
            stdJson.readAddress(
                openOracleDeployedContracts,
                ".addresses.proxyAdmin"
            )
        );

        string memory openOracleDeployedIdenticalContracts = readOutput(
            "open_oracle_avs_identical_task_manager_output"
        );

        openOracleIdenticalAnswerTaskManager = OpenOracleIdenticalAnswerTaskManager(
            stdJson.readAddress(
                openOracleDeployedIdenticalContracts,
                ".addresses.openOracleIdenticalAnswerTaskManager"
            )
        );

        address openOracleCommunityMultisig = msg.sender;
        address openOraclePauser = msg.sender;

        vm.startBroadcast();

        _deployOpenOraclePriceFeeds(
            proxyAdmin
        );
        vm.stopBroadcast();
    }

    function _deployOpenOraclePriceFeeds(
        ProxyAdmin proxyAdmin
    ) internal {
        // READ JSON CONFIG DATA
        string memory config_data = vm.readFile(deployConfigPath);

        // parse initalization params and permissions from config data
        DeployParams memory deployParams = _parseDeployParams(config_data);




        // WRITE JSON DATA
        string memory parent_object = "parent object";
        string memory deployed_addresses = "addresses";
        string memory implementation_addresses = "impl_addresses";

        for (uint8 i = deployParams.taskTypeLower; i <= deployParams.taskTypeUpper; i++) {

            OpenOracleVRFFeed openOracleVRFFeedImplementation = new OpenOracleVRFFeed(
                openOracleIdenticalAnswerTaskManager
            );
            TransparentUpgradeableProxy openOracleVRFFeed = new TransparentUpgradeableProxy(
                address(openOracleVRFFeedImplementation), address(proxyAdmin), abi.encodeWithSelector(
                        openOracleVRFFeedImplementation.initialize.selector,
                        msg.sender,
                        i,
                        deployParams.responderThreshold,
                        deployParams.stakeThreshold
                    )
            );

            vm.serializeAddress(
                deployed_addresses,
                toString(i),
                address(openOracleVRFFeed)
            );
            vm.serializeAddress(
                implementation_addresses,
                toString(i),
                address(openOracleVRFFeedImplementation)
            );
            openOracleIdenticalAnswerTaskManager.addToFeedlist(address(openOracleVRFFeed));
        }

        string memory deployed_addresses_output = vm.serializeAddress(
            deployed_addresses,
            "openOracleIdenticalAnswerTaskManager",
                address(openOracleIdenticalAnswerTaskManager)
        );

        string memory implementation_addresses_output = vm.serializeAddress(
            implementation_addresses,
            "proxyAdmin",
            address(proxyAdmin)
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

        writeOutput(finalJson, "open_oracle_avs_identical_feed_output");
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


// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {OpenOracleIdenticalAnswerTaskManager} from "../../src/OpenOracleIdenticalAnswerTaskManager.sol";
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
contract OpenOracleIdenticalTaskManagerDeployer is Script, Utils {

    string public deployAVSConfigPath = string.concat("./script/config/", vm.toString(block.chainid), "/config.avs.json");

    struct DeployParams {
        uint8 taskTypeLower;
        uint8 taskTypeUpper;
        uint8 responderThreshold;
        uint96 stakeThreshold;
    }

    struct DeployAvsParams {
        uint32 taskResponseWindowBlock;
        address aggregatorAddr;
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

        StakeRegistry stakeRegistry = StakeRegistry(
            stdJson.readAddress(
                openOracleDeployedContracts,
                ".addresses.stakeRegistry"
            )
        );
        BLSApkRegistry blsApkRegistry = BLSApkRegistry(
            stdJson.readAddress(
                openOracleDeployedContracts,
                ".addresses.blsApkRegistry"
            )
        );

        IRegistryCoordinator registryCoordinator = IRegistryCoordinator(
            stdJson.readAddress(
                openOracleDeployedContracts,
                ".addresses.registryCoordinator"
            )
        );

        address openOracleCommunityMultisig = msg.sender;
        address openOraclePauser = msg.sender;

        vm.startBroadcast();

        _deployOpenOraclePriceFeeds(
            proxyAdmin,
                stakeRegistry,
                blsApkRegistry,
            msg.sender,
                registryCoordinator
        );
        vm.stopBroadcast();
    }

    function _deployOpenOraclePriceFeeds(
        ProxyAdmin proxyAdmin,
        StakeRegistry stakeRegistry,
        BLSApkRegistry blsApkRegistry,
        address openOracleCommunityMultisig,
        IRegistryCoordinator registryCoordinator
    ) internal {



        string memory config_avs_data = vm.readFile(deployAVSConfigPath);

        // parse initalization params and permissions from config data
        DeployAvsParams memory deployAvsParams = _parseAvsDeployParams(config_avs_data);

        // WRITE JSON DATA
        string memory parent_object = "parent object";
        string memory deployed_addresses = "addresses";
        string memory implementation_addresses = "impl_addresses";

        BLSSignatureChecker blsSignatureCheckerImpl = new BLSSignatureChecker(registryCoordinator);
        OpenOracleIdenticalAnswerTaskManager openOracleIdenticalAnswerTaskManagerImpl = new OpenOracleIdenticalAnswerTaskManager(
            stakeRegistry,
            blsApkRegistry,
                deployAvsParams.taskResponseWindowBlock,
                blsSignatureCheckerImpl
        );
        PauserRegistry openOraclePauserReg;
        {
            address[] memory pausers = new address[](2);
            pausers[0] = openOracleCommunityMultisig;
            pausers[1] = openOracleCommunityMultisig;
            openOraclePauserReg = new PauserRegistry(
                pausers,
                openOracleCommunityMultisig
            );
        }

        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(openOracleIdenticalAnswerTaskManagerImpl), address(proxyAdmin), abi.encodeWithSelector(
                    openOracleIdenticalAnswerTaskManagerImpl.initialize.selector,
                    openOraclePauserReg,
                    openOracleCommunityMultisig,
                        deployAvsParams.aggregatorAddr
                )
        );

        string memory deployed_addresses_output = vm.serializeAddress(
            deployed_addresses,
            "openOracleIdenticalAnswerTaskManager",
            address(proxy)
        );

        string memory implementation_addresses_output = vm.serializeAddress(
            implementation_addresses,
            "openOracleIdenticalAnswerTaskManagerImplementation",
            address(openOracleIdenticalAnswerTaskManagerImpl)
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

        writeOutput(finalJson, "open_oracle_avs_identical_task_manager_output");
    }

    function _parseAvsDeployParams(string memory config_data) internal pure returns (DeployAvsParams memory deployParams) {
        bytes memory taskResponseWindowBlockRaw = stdJson.parseRaw(config_data, ".taskResponseWindowBlock");
        deployParams.taskResponseWindowBlock = abi.decode(taskResponseWindowBlockRaw, (uint32));

        bytes memory aggregatorAddrRaw = stdJson.parseRaw(config_data, ".aggregatorAddr");
        deployParams.aggregatorAddr = abi.decode(aggregatorAddrRaw, (address));
    }
}


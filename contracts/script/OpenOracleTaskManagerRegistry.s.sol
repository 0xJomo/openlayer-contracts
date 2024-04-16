// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import "@eigenlayer/contracts/permissions/PauserRegistry.sol";
import "@eigenlayer/test/mocks/EmptyContract.sol";

import {BLSApkRegistry} from "@eigenlayer-middleware/src/BLSApkRegistry.sol";
import {StakeRegistry} from "@eigenlayer-middleware/src/StakeRegistry.sol";
import "@eigenlayer-middleware/src/OperatorStateRetriever.sol";

import {OpenOracleServiceManager} from "../src/OpenOracleServiceManager.sol";
import "../src/ERC20Mock.sol";

import {Utils} from "./utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

// # To deploy and verify our contract
// forge script script/OpenOracleTaskManagerRegistry.s.sol:OpenOracleTaskManagerRegistry --rpc-url http://127.0.0.1:8545  --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --broadcast -vvvv
contract OpenOracleTaskManagerRegistry is Script, Utils {
    string public deployConfigPath = string(bytes("./script/config/devnet/testnet.config.json"));

    // DEPLOYMENT CONSTANTS
    address public constant SERVICE_MANAGER_ADDR =
        0x67d269191c92Caf3cD7723F116c85e6E9bf55933;

    OpenOracleServiceManager public openOracleTaskManager;

    function run() external {
        // Eigenlayer contracts
        string memory openOracleAVSDeployedContracts = readOutput(
            "open_oracle_avs_deployment_output"
        );
        OpenOracleServiceManager openOracleServiceManager = OpenOracleServiceManager(
            stdJson.readAddress(
                openOracleAVSDeployedContracts,
                ".addresses.openOracleServiceManager"
            )
        );

        // openOracleServiceManager.addTaskManager("http://127.0.0.1:8545", address("0x67d269191c92Caf3cD7723F116c85e6E9bf5593"));
    }
}

// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import {OpenOraclePriceFeed} from "../../src/OpenOraclePriceFeed.sol";
import {OpenOracleTaskManager} from "../../src/OpenOracleTaskManager.sol";
import {Utils} from "../utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

// Upgrade ServiceManager contract
contract UpgradePriceFeed is Script, Utils {
    function run() external {

        string memory openOracleDeployedContracts = readOutput(
            "open_oracle_avs_pricefeeds_output_v2"
        );


        ProxyAdmin openOracleProxyAdmin = ProxyAdmin(
            stdJson.readAddress(
                openOracleDeployedContracts,
                ".impl_addresses.proxyAdmin"
            )
        );

        OpenOracleTaskManager openOracleTaskManager = OpenOracleTaskManager(
            stdJson.readAddress(
                openOracleDeployedContracts,
                ".addresses.openOracleTaskManager"
            )
        );

        //        "0": "0x3535894cA83F8490d67F637a2F82c79Ce0770A59",
        //    "1": "0x83664EF16F3C75364a317483E3fbfcE6FB54ca97",
        //    "10": "0x8e9d9e76cce57f89c38f76783Dc59C41e28ee893",
        //        "11": "0x86BB9371830F3b752e11c773DAB0b0B3db35F061",
        //        "12": "0xcC03A6ca0fE01bd199Ff9B194Fc92f87c82Fb49c",
        //        "13": "0x825a5A37181B80465E7fba42810E911FdABD16E6",
        //    "2": "0xA259D44aC7749CF95a76A739d2E920784727b829",
        //    "3": "0xF4419a82826B31E9C3fF99edc5c1D112E52ac128",
        //        "4": "0x98ff7FB8BbB95Bec3Ccb19AcACa7122f3471AeBB",
        //        "5": "0xB2ddf346E53DdAF5fF7bCD6EDdDcb3b3D370e102",
        //        "6": "0x389b5e92a024b066920156B6b747D35e04e3E4FA",
        //    "7": "0x0C440C4B79bbea118Ed8bBA1857d507Bd38F7e2d",
        //    "8": "0x02CFf54e42B8a8174877d02747f079057B183d7d",
        //        "9": "0xe8C4f1Fb39967240B31d9532AF5010D2552B3a20",
        // holesky
        address[14] memory deployParams = [
            0x3535894cA83F8490d67F637a2F82c79Ce0770A59,
            0x83664EF16F3C75364a317483E3fbfcE6FB54ca97,
            0x8e9d9e76cce57f89c38f76783Dc59C41e28ee893,
            0x86BB9371830F3b752e11c773DAB0b0B3db35F061,
            0xcC03A6ca0fE01bd199Ff9B194Fc92f87c82Fb49c,
            0x825a5A37181B80465E7fba42810E911FdABD16E6,
            0xA259D44aC7749CF95a76A739d2E920784727b829,
            0xF4419a82826B31E9C3fF99edc5c1D112E52ac128,
            0x98ff7FB8BbB95Bec3Ccb19AcACa7122f3471AeBB,
            0xB2ddf346E53DdAF5fF7bCD6EDdDcb3b3D370e102,
            0x389b5e92a024b066920156B6b747D35e04e3E4FA,
            0x0C440C4B79bbea118Ed8bBA1857d507Bd38F7e2d,
            0x02CFf54e42B8a8174877d02747f079057B183d7d,
            0xe8C4f1Fb39967240B31d9532AF5010D2552B3a20];

        // plumetest todo : first
//        address[1] memory deployParams = [0xa852961E8c22982489a6f401caf8e0a3040eD847];

        vm.startBroadcast();
        // Start upgrade

        OpenOraclePriceFeed openOraclePriceFeedImplementation = new OpenOraclePriceFeed(
            openOracleTaskManager
        );


        for (uint256 i = 0; i < deployParams.length; i++){
            openOracleProxyAdmin.upgrade(
                TransparentUpgradeableProxy(
                    payable(deployParams[i])
                ),
                address(openOraclePriceFeedImplementation)
            );
        }


        // End upgrade
        vm.stopBroadcast();
    }
}
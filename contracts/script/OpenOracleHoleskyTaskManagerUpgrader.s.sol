// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import "@eigenlayer/contracts/permissions/PauserRegistry.sol";
import {IAVSDirectory} from "@eigenlayer/contracts/interfaces/IAVSDirectory.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IStrategyManager, IStrategy} from "@eigenlayer/contracts/interfaces/IStrategyManager.sol";
import {ISlasher} from "@eigenlayer/contracts/interfaces/ISlasher.sol";
import {StrategyBaseTVLLimits} from "@eigenlayer/contracts/strategies/StrategyBaseTVLLimits.sol";
import "@eigenlayer/test/mocks/EmptyContract.sol";

import "@eigenlayer-middleware/src/RegistryCoordinator.sol" as regcoord;
import {IBLSApkRegistry, IIndexRegistry, IStakeRegistry} from "@eigenlayer-middleware/src/RegistryCoordinator.sol";
import {BLSApkRegistry} from "@eigenlayer-middleware/src/BLSApkRegistry.sol";
import {IndexRegistry} from "@eigenlayer-middleware/src/IndexRegistry.sol";
import {StakeRegistry} from "@eigenlayer-middleware/src/StakeRegistry.sol";
import "@eigenlayer-middleware/src/OperatorStateRetriever.sol";

import {OpenOracleServiceManager, IServiceManager} from "../src/OpenOracleServiceManager.sol";
import {OpenOracleTaskManager} from "../src/OpenOracleTaskManager.sol";
import {IOpenOracleTaskManager} from "../src/IOpenOracleTaskManager.sol";
import "../src/ERC20Mock.sol";

import {Utils} from "./utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

// # To deploy and verify our contract
// forge script script/OpenOracleDeployer.s.sol:OpenOracleDeployer --rpc-url http://127.0.0.1:8545  --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --broadcast -vvvv
contract OpenOracleHoleskyTaskManagerUpgrader is Script, Utils {
    string public deployConfigPath = string(bytes("./script/config/holesky/testnet.config.json"));

    // DEPLOYMENT CONSTANTS
    uint32 public constant TASK_RESPONSE_WINDOW_BLOCK = 30;
    // TODO: right now hardcoding these (this address is anvil's default address 9)
    address public constant AGGREGATOR_ADDR =
        0xa0Ee7A142d267C1f36714E4a8F75612F20a79720;

    IOpenOracleTaskManager public openOracleTaskManagerImplementation;

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
        regcoord.RegistryCoordinator registryCoordinator = regcoord.RegistryCoordinator(
            stdJson.readAddress(
                openOracleDeployedContracts,
                ".addresses.registryCoordinator"
            )
        );
        ProxyAdmin openOracleProxyAdmin = ProxyAdmin(
            stdJson.readAddress(
                openOracleDeployedContracts,
                ".addresses.proxyAdmin"
            )
        );
        PauserRegistry openOraclePauserReg = PauserRegistry(
            stdJson.readAddress(
                openOracleDeployedContracts,
                ".addresses.pauserRegistry"
            )
        );

        address openOracleCommunityMultisig = msg.sender;
        address openOraclePauser = msg.sender;

        vm.startBroadcast();

        _upgradeOpenOracleContracts(
            openOracleTaskManager,
            registryCoordinator,
            openOracleProxyAdmin,
            openOraclePauserReg,
            openOracleCommunityMultisig,
            openOraclePauser
        );
        vm.stopBroadcast();
    }

    function _upgradeOpenOracleContracts(
        OpenOracleTaskManager openOracleTaskManager,
        regcoord.RegistryCoordinator registryCoordinator,
        ProxyAdmin openOracleProxyAdmin,
        PauserRegistry openOraclePauserReg,
        address openOracleCommunityMultisig,
        address openOraclePauser
    ) internal {

        openOracleTaskManagerImplementation = new OpenOracleTaskManager(
            registryCoordinator,
            TASK_RESPONSE_WINDOW_BLOCK
        );

        // Upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        openOracleProxyAdmin.upgrade(
            TransparentUpgradeableProxy(
                payable(address(openOracleTaskManager))
            ),
            address(openOracleTaskManagerImplementation)
        );
        openOracleTaskManager.updateAggregator(AGGREGATOR_ADDR);

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
            "openOracleTaskManagerImplementation",
            address(openOracleTaskManagerImplementation)
        );

        // serialize all the data
        string memory finalJson = vm.serializeString(
            parent_object,
            deployed_addresses,
            deployed_addresses_output
        );

        writeOutput(finalJson, "open_oracle_avs_upgrade_taskmanager_output");
    }
}
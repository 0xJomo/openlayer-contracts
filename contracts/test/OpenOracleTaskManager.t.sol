// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import "../src/OpenOracleServiceManager.sol" as incsqsm;
import {OpenOracleTaskManager} from "../src/OpenOracleTaskManager.sol";
import {BLSMockAVSDeployer} from "@eigenlayer-middleware/test/utils/BLSMockAVSDeployer.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract OpenOracleTaskManagerTest is BLSMockAVSDeployer {
    incsqsm.OpenOracleServiceManager sm;
    incsqsm.OpenOracleServiceManager smImplementation;
    OpenOracleTaskManager tm;
    OpenOracleTaskManager tmImplementation;

    uint32 public constant TASK_RESPONSE_WINDOW_BLOCK = 30;
    address aggregator =
        address(uint160(uint256(keccak256(abi.encodePacked("aggregator")))));

    function setUp() public {
        _setUpBLSMockAVSDeployer();

        tmImplementation = new OpenOracleTaskManager(
            incsqsm.IRegistryCoordinator(address(registryCoordinator)),
            TASK_RESPONSE_WINDOW_BLOCK
        );

        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        tm = OpenOracleTaskManager(
            address(
                new TransparentUpgradeableProxy(
                    address(tmImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(
                        tm.initialize.selector,
                        pauserRegistry,
                        registryCoordinatorOwner,
                        aggregator
                    )
                )
            )
        );
    }

    function testCreateNewTask() public {
        bytes memory quorumNumbers = new bytes(0);
        tm.createNewTask(2, 100, quorumNumbers);
        assertEq(tm.latestTaskNum(), 1);
    }
}

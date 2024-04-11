// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "./OpenOracleTaskManager.sol";
import "./IOpenOraclePriceFeed.sol";

contract OpenOraclePriceFeed is IOpenOraclePriceFeed {

    OpenOracleTaskManager internal immutable _openOracleTaskManager;

    uint8 internal _taskType;
    uint8 internal _responderThreshold;
    uint96 internal _stakeThreshold;

    int256 _latestValue;

    modifier onlyTaskManager() {
        require(
            msg.sender == address(_openOracleTaskManager),
            "OpenOraclePriceFeed: caller is not the task manager"
        );
        _;
    } 

    constructor(
        OpenOracleTaskManager __openOracleTaskManager,
        uint8 __taskType,
        uint8 __responderThreshold,
        uint96 __stakeThreshold
    ) {
        _openOracleTaskManager = __openOracleTaskManager;
        _taskType = __taskType;
        _responderThreshold = __responderThreshold;
        _stakeThreshold = __stakeThreshold;
    }

    function requestNewReport() external {
        _openOracleTaskManager.createNewTask(_taskType, _responderThreshold, _stakeThreshold);
    }

    function saveLatestData(int256 value) external onlyTaskManager {
        _latestValue = value;
    }


    function latestRoundData() view external
    returns (int256) {
        return _latestValue;
    }
}
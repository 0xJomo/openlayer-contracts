// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "./IOpenOracleIdenticalAnswerTaskManager.sol";

interface IOpenOracleVRFFeed {
    event NewIdenticalAnswerReported(
        uint8 indexed taskType, 
        uint32 referenceTaskIndex,
        bytes result,
        uint256 timestamp, 
        uint32 createdBlock, 
        uint32 respondedBlock
    );

    function createNewTask(bytes calldata taskData) external;

    function createNewTask(bytes calldata taskData,
        uint8 responderThreshold,
        uint96 stakeThreshold) external;

    function createNewTaskWithCallback(bytes calldata taskData, uint256 requestId) external;

    function createNewTaskWithCallback(bytes calldata taskData,
        uint8 responderThreshold,
        uint96 stakeThreshold, uint256 requestId) external;

    /// @notice Saves the latest data from task manager in contract
    function saveLatestData(
        IOpenOracleIdenticalAnswerTaskManager.Task calldata task, 
        IOpenOracleIdenticalAnswerTaskManager.AggregatedTaskResponse calldata response, 
        IOpenOracleIdenticalAnswerTaskManager.TaskResponseMetadata calldata metadata
    ) external;

    /// @notice Returns the latest data
    function latestRoundData() view external returns (
        bytes calldata result,
        uint256 timestamp,
        uint32 startBlock,
        uint32 endBlock
    );

    /// @notice Returns the latest data
    function getRoundData(uint32 roundId) view external returns (
        bytes calldata result,
        uint256 timestamp,
        uint32 startBlock,
        uint32 endBlock
    );

    function setThresholds(uint8 responderThreshold, uint96 stakeThreshold) external;
}

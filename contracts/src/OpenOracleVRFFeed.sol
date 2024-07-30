// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "./OpenOracleIdenticalAnswerTaskManager.sol";
import "./IOpenOracleVRFFeed.sol";

import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";


contract OpenOracleVRFFeed is Initializable, OwnableUpgradeable,
    IOpenOracleVRFFeed  {

    OpenOracleIdenticalAnswerTaskManager internal immutable _openOracleIdenticalAnswerTaskManager;

    uint8 internal _taskType;
    uint8 internal _responderThreshold;
    uint96 internal _stakeThreshold;

    IOpenOracleIdenticalAnswerTaskManager.AggregatedTaskResponse _latestResponse;
    IOpenOracleIdenticalAnswerTaskManager.TaskResponseMetadata _latestMetadata;
    uint32 _latestCreatedBlock;

    struct RoundInfo{
        IOpenOracleIdenticalAnswerTaskManager.AggregatedTaskResponse _latestResponse;
        IOpenOracleIdenticalAnswerTaskManager.TaskResponseMetadata _latestMetadata;
        uint32 _latestCreatedBlock;
    }
    mapping(uint32 => RoundInfo) internal taskRoundInfo;

    modifier onlyTaskManager() {
        require(
            msg.sender == address(_openOracleIdenticalAnswerTaskManager),
            "OpenOraclePriceFeed: caller is not the task manager"
        );
        _;
    }

    constructor(
        OpenOracleIdenticalAnswerTaskManager __openOracleIdenticalAnswerTaskManager
    ) {
        _openOracleIdenticalAnswerTaskManager = __openOracleIdenticalAnswerTaskManager;
    }

    function initialize(
        address initialOwner,
        uint8 __taskType,
        uint8 __responderThreshold,
        uint96 __stakeThreshold
    ) public initializer {
        _transferOwnership(initialOwner);
        _taskType = __taskType;
        _responderThreshold = __responderThreshold;
        _stakeThreshold = __stakeThreshold;
    }

    function createNewTask(bytes calldata taskData) external {
        _openOracleIdenticalAnswerTaskManager.createNewTask(
            _taskType, taskData, _responderThreshold, _stakeThreshold);
    }

    function createNewTask(bytes calldata taskData,
        uint8 responderThreshold,
        uint96 stakeThreshold) external {
        _openOracleIdenticalAnswerTaskManager.createNewTask(
            _taskType, taskData, responderThreshold, stakeThreshold);
    }

    function saveLatestData(
        IOpenOracleIdenticalAnswerTaskManager.Task calldata task,
        IOpenOracleIdenticalAnswerTaskManager.AggregatedTaskResponse calldata response,
        IOpenOracleIdenticalAnswerTaskManager.TaskResponseMetadata calldata metadata
    ) external onlyTaskManager {
        bytes memory result = new bytes(32);
        bytes32 data = keccak256(abi.encode(response.aggregatedSignature));
        assembly {
            mstore(add(result, 32), data) // Store the bytes32 value at the start of the memory allocated for the bytes array
        }
        _latestResponse = response;
        _latestMetadata = metadata;
        _latestCreatedBlock = task.taskCreatedBlock;
        _latestResponse.msg.result = result;
        uint32 taskIndex = response.msg.referenceTaskIndex;
        taskRoundInfo[taskIndex]._latestResponse = _latestResponse;
        taskRoundInfo[taskIndex]._latestMetadata = _latestMetadata;
        taskRoundInfo[taskIndex]._latestCreatedBlock = _latestCreatedBlock;

        emit NewIdenticalAnswerReported(
            task.taskType,
                taskIndex,
                result,
            response.timestamp,
            task.taskCreatedBlock,
            metadata.taskResponsedBlock
        );
    }

    function latestRoundData() view external
    returns (
        bytes memory result,
        uint256 timestamp,
        uint32 startBlock,
        uint32 endBlock
    ) {
        return (
            _latestResponse.msg.result,
            _latestResponse.timestamp,
            _latestCreatedBlock,
            _latestMetadata.taskResponsedBlock
        );
    }

    function getRoundData(uint32 roundId) view external
    returns (
        bytes memory result,
        uint256 timestamp,
        uint32 startBlock,
        uint32 endBlock
    ) {
        return (
            taskRoundInfo[roundId]._latestResponse.msg.result,
            taskRoundInfo[roundId]._latestResponse.timestamp,
            taskRoundInfo[roundId]._latestCreatedBlock,
            taskRoundInfo[roundId]._latestMetadata.taskResponsedBlock
        );
    }

    function setThresholds(uint8 responderThreshold, uint96 stakeThreshold) external onlyOwner {
        _responderThreshold = responderThreshold;
        _stakeThreshold = stakeThreshold;
    }
}
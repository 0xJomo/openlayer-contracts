// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "./OpenOracleTaskManager.sol";
import "./IOpenOracleCommonDataFeed.sol";

import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";


contract OpenOracleCommonDataFeed is Initializable, OwnableUpgradeable, IOpenOracleCommonDataFeed  {

    OpenOracleTaskManager internal immutable _openOracleTaskManager;

    struct TaskInfo{
        uint8  _taskType;
        uint8  _responderThreshold;
        uint96  _stakeThreshold;
        IOpenOracleTaskManager.WeightedTaskResponse _latestResponse;
        IOpenOracleTaskManager.TaskResponseMetadata _latestMetadata;
        uint32 _latestCreatedBlock;
    }

    uint8  responderThreshold;
    uint96  stakeThreshold;

    mapping(uint8 => TaskInfo) internal taskInfo;

    struct RoundInfo{
        IOpenOracleTaskManager.WeightedTaskResponse _latestResponse;
        IOpenOracleTaskManager.TaskResponseMetadata _latestMetadata;
        uint32 _latestCreatedBlock;
    }
    mapping(uint32 => RoundInfo) internal taskRoundInfo;

    mapping(bytes32 => address) private requestClients;

    uint256 public callbackLimit;

    mapping(bytes32 => uint256) private requestRounds;

    modifier onlyTaskManager() {
        require(
            msg.sender == address(_openOracleTaskManager),
            "OpenOraclePriceFeed: caller is not the task manager"
        );
        _;
    } 

    constructor(
        OpenOracleTaskManager __openOracleTaskManager
    ) {
        _openOracleTaskManager = __openOracleTaskManager;
    }

    function initialize(
        address initialOwner,
        uint8 __responderThreshold,
        uint96 __stakeThreshold
    ) public initializer {
        _transferOwnership(initialOwner);
        responderThreshold = __responderThreshold;
        stakeThreshold = __stakeThreshold;
    }

    function requestNewReport(uint8 _taskType) override external {
        _createTaskInfo(_taskType);
        _openOracleTaskManager.createNewTask(_taskType, taskInfo[_taskType]._responderThreshold, taskInfo[_taskType]._stakeThreshold);

    }

    function requestNewReportWithData(uint8 _taskType,bytes calldata _taskData) override external {
        _createTaskInfo(_taskType);
        _openOracleTaskManager.createNewTaskWithData(_taskType, taskInfo[_taskType]._responderThreshold, taskInfo[_taskType]._stakeThreshold, _taskData);
    }

    function requestNewReportCallback(uint8 _taskType, uint256 requestId) override external {
        _createTaskInfo(_taskType);
        bytes32 taskHash = _openOracleTaskManager.createNewTask(_taskType, taskInfo[_taskType]._responderThreshold, taskInfo[_taskType]._stakeThreshold);
        requestClients[taskHash] = msg.sender;
        requestRounds[taskHash] = requestId;
    }

    function requestNewReportWithDataCallback(uint8 _taskType,bytes calldata _taskData, uint256 requestId) override external {
        _createTaskInfo(_taskType);
        bytes32 taskHash = _openOracleTaskManager.createNewTaskWithData(_taskType, taskInfo[_taskType]._responderThreshold, taskInfo[_taskType]._stakeThreshold, _taskData);
        requestClients[taskHash] = msg.sender;
        requestRounds[taskHash] = requestId;
    }

    function _createTaskInfo(uint8 _taskType) internal{
        if(taskInfo[_taskType]._latestCreatedBlock <= uint32(0)){
            IOpenOracleTaskManager.WeightedTaskResponse memory _latestResponse;
            IOpenOracleTaskManager.TaskResponseMetadata memory _latestMetadata;
            TaskInfo memory _taskInfo = TaskInfo(_taskType,responderThreshold,stakeThreshold,_latestResponse,_latestMetadata,uint32(block.number));
            taskInfo[_taskType] = _taskInfo;
        }
    }

    function saveLatestData(
        IOpenOracleTaskManager.Task calldata task,
        IOpenOracleTaskManager.WeightedTaskResponse calldata response, 
        IOpenOracleTaskManager.TaskResponseMetadata calldata metadata
    ) external onlyTaskManager {
        taskRoundInfo[response.referenceTaskIndex]._latestResponse = response;
        taskRoundInfo[response.referenceTaskIndex]._latestMetadata = metadata;
        taskRoundInfo[response.referenceTaskIndex]._latestCreatedBlock = task.taskCreatedBlock;
        taskInfo[task.taskType]._latestResponse = response;
        taskInfo[task.taskType]._latestMetadata = metadata;
        taskInfo[task.taskType]._latestCreatedBlock = task.taskCreatedBlock;
        emit NewPriceReported(
            task.taskType, 
            response.referenceTaskIndex, 
            response.result, 
            response.sd, 
            response.timestamp, 
            task.taskCreatedBlock, 
            metadata.taskResponsedBlock
        );

        bytes32 taskHash = keccak256(abi.encode(task));
        address _callback = requestClients[taskHash];
        if (_callback != address(0)) {
            (bool success, ) = _callback.call{gas: callbackLimit}(abi.encodeWithSignature("fulfillResult(uint256,bytes,bytes)",requestRounds[taskHash], response.result, 0x0));
            require(success, "fulfillResult callback failed");
        }
    }

    function latestRoundData(uint8 taskType) view external
    returns (
        bytes memory result,
        uint256 sd,
        uint256 timestamp,
        uint32 startBlock,
        uint32 endBlock
    ) {
        return (
        taskInfo[taskType]._latestResponse.result,
        taskInfo[taskType]._latestResponse.sd,
        taskInfo[taskType]._latestResponse.timestamp,
        taskInfo[taskType]._latestCreatedBlock,
        taskInfo[taskType]._latestMetadata.taskResponsedBlock
        );
    }

    function getRoundData(uint32 roundId) view external
    returns (
        bytes memory result,
        uint256 sd,
        uint256 timestamp,
        uint32 startBlock,
        uint32 endBlock
    ) {
        return (
        taskRoundInfo[roundId]._latestResponse.result,
        taskRoundInfo[roundId]._latestResponse.sd,
        taskRoundInfo[roundId]._latestResponse.timestamp,
        taskRoundInfo[roundId]._latestCreatedBlock,
        taskRoundInfo[roundId]._latestMetadata.taskResponsedBlock
        );
    }

    function setDefaultThresholds(uint8 _responderThreshold, uint96 _stakeThreshold) external onlyOwner {
        responderThreshold = _responderThreshold;
        stakeThreshold = _stakeThreshold;
    }

    function setThresholds(uint8 taskType, uint8 _responderThreshold, uint96 _stakeThreshold) external onlyOwner{
        require(taskInfo[taskType]._latestCreatedBlock > uint32(0), "OpenOraclePriceFeed: taskType not exists");
        taskInfo[taskType]._responderThreshold = _responderThreshold;
        taskInfo[taskType]._stakeThreshold = _stakeThreshold;
    }

    function setCallbackLimit(uint256 _callbackLimit) external onlyOwner {
        callbackLimit = _callbackLimit;
    }
}
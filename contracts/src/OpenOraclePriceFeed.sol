// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "./OpenOracleTaskManager.sol";
import "./IOpenOraclePriceFeed.sol";

import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";


contract OpenOraclePriceFeed is Initializable, OwnableUpgradeable,
    IOpenOraclePriceFeed  {

    OpenOracleTaskManager internal immutable _openOracleTaskManager;

    uint8 internal _taskType;
    uint8 internal _responderThreshold;
    uint96 internal _stakeThreshold;

    IOpenOracleTaskManager.WeightedTaskResponse _latestResponse;
    IOpenOracleTaskManager.TaskResponseMetadata _latestMetadata;
    uint32 _latestCreatedBlock;

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
        uint8 __taskType,
        uint8 __responderThreshold,
        uint96 __stakeThreshold
    ) public initializer {
        _transferOwnership(initialOwner);
        _taskType = __taskType;
        _responderThreshold = __responderThreshold;
        _stakeThreshold = __stakeThreshold;
    }

    function requestNewReport() override external {
        _openOracleTaskManager.createNewTask(_taskType, _responderThreshold, _stakeThreshold);
    }

    function requestNewReportWithData(bytes calldata _taskData) override external {
        _openOracleTaskManager.createNewTaskWithData(_taskType, _responderThreshold, _stakeThreshold, _taskData);
    }


    function saveLatestData(
        IOpenOracleTaskManager.Task calldata task,
        IOpenOracleTaskManager.WeightedTaskResponse calldata response, 
        IOpenOracleTaskManager.TaskResponseMetadata calldata metadata
    ) external onlyTaskManager {
        _latestResponse = response;
        _latestMetadata = metadata;
        _latestCreatedBlock = task.taskCreatedBlock;

        emit NewPriceReported(
            task.taskType, 
            response.referenceTaskIndex, 
            response.result, 
            response.sd, 
            response.timestamp, 
            task.taskCreatedBlock, 
            metadata.taskResponsedBlock
        );
    }

    function latestRoundData() view external
    returns (
        bytes memory result,
        uint256 sd,
        uint256 timestamp,
        uint32 startBlock,
        uint32 endBlock
    ) {
        return (
            _latestResponse.result, 
            _latestResponse.sd, 
            _latestResponse.timestamp, 
            _latestCreatedBlock, 
            _latestMetadata.taskResponsedBlock
        );
    }

    function setThresholds(uint8 responderThreshold, uint96 stakeThreshold) external onlyOwner {
        _responderThreshold = responderThreshold;
        _stakeThreshold = stakeThreshold;
    }
}
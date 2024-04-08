// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@eigenlayer-middleware/src/libraries/BN254.sol";

interface IOpenOracleTaskManager {
    // EVENTS
    event NewTaskCreated(uint32 indexed taskIndex, Task task);

    event TaskResponded(
        WeightedTaskResponse taskResponse,
        TaskResponseMetadata taskResponseMetadata
    );

    event TaskCompleted(uint32 indexed taskIndex);

    event FundsWithdrawn(uint taskNum, address creator, uint amount);

    // STRUCTS
    struct Task {
        uint8 taskType;
        uint32 taskCreatedBlock;
        uint8 responderNumber;
        uint96 stakeThreshold;
        address payable creator;
        uint creationFee;
    }

    // Task response is hashed and signed by operators.
    // these signatures are aggregated and sent to the contract as response.
    struct TaskResponse {
        // Can be obtained by the operator from the event NewTaskCreated.
        uint32 referenceTaskIndex;
        // This is just the response that the operator has to compute by itself.
        uint256 result;
        // This is just the response that the operator has to compute by itself.
        uint256 timeStamp;
    }

    struct OperatorResponse {
        address operator;
        TaskResponse response;
        bytes signature;
    }

    struct WeightedTaskResponse {
        // Can be obtained by the operator from the event NewTaskCreated.
        uint32 referenceTaskIndex;
        // Weighted result from operator responses
        uint256 result;
        // Standard deviation for weighted result
        uint256 sd;
        // Timestamp for result
        uint256 timeStamp;
    }

    // Extra information related to taskResponse, which is filled inside the contract.
    struct TaskResponseMetadata {
        uint32 taskResponsedBlock;
    }

    // FUNCTIONS
    // NOTE: this function creates new task.
    function createNewTask(
        uint8 taskType,
        uint8 responderNumber,
        uint96 stakeThreshold
    ) external payable;

    /// @notice Returns the current 'taskNumber' for the middleware
    function taskNumber() external view returns (uint32);

    /// @notice Returns the TASK_RESPONSE_WINDOW_BLOCK
    function getTaskResponseWindowBlock() external view returns (uint32);
}

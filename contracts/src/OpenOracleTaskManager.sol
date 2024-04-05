// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@eigenlayer/contracts/permissions/Pausable.sol";
import "@eigenlayer-middleware/src/interfaces/IServiceManager.sol";
import {BLSApkRegistry} from "@eigenlayer-middleware/src/BLSApkRegistry.sol";
import {RegistryCoordinator} from "@eigenlayer-middleware/src/RegistryCoordinator.sol";
import {BLSSignatureChecker, IRegistryCoordinator} from "@eigenlayer-middleware/src/BLSSignatureChecker.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {OperatorStateRetriever} from "@eigenlayer-middleware/src/OperatorStateRetriever.sol";
import "@eigenlayer-middleware/src/libraries/BN254.sol";
import "./IOpenOracleTaskManager.sol";

contract OpenOracleTaskManager is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    BLSSignatureChecker,
    OperatorStateRetriever,
    IOpenOracleTaskManager
{
    using BN254 for BN254.G1Point;
    using ECDSA for bytes32;

    /* CONSTANT */
    // The number of blocks from the task initialization within which the aggregator has to respond to
    uint32 public immutable TASK_RESPONSE_WINDOW_BLOCK;
    uint256 internal constant _THRESHOLD_DENOMINATOR = 100;
    uint8 private constant _TASKRESPONSE_QUORUM_NUMBER = 1;

    // Fee for creating a task
    uint public taskCreationFee = 0.0001 ether; 

    /* STORAGE */
    // The latest task index
    uint32 public latestTaskNum;

    // mapping of task indices to all tasks hashes
    // when a task is created, task hash is stored here,
    // and responses need to pass the actual task,
    // which is hashed onchain and checked against this mapping
    mapping(uint32 => bytes32) public allTaskHashes;

    // mapping of task indices to hash of abi.encode(taskResponse, taskResponseMetadata)
    mapping(uint32 => bytes32) public allTaskResponses;

    address public aggregator;

    /* MODIFIERS */
    modifier onlyAggregator() {
        require(msg.sender == aggregator, "Aggregator must be the caller");
        _;
    }

    modifier paysTaskCreationFee(uint totalRequests) {
        require(msg.value >= taskCreationFee * totalRequests, "Creating a task requires a fee.");
        _;
    }

    modifier onlyTaskCreator(uint32 taskNum, Task memory task) {
        require(
            taskNum <= latestTaskNum,
            "task number not exists"
        );
        require(
            keccak256(abi.encode(task)) ==
                allTaskHashes[taskNum],
            "supplied task does not match the one recorded in the contract"
        );
        require(task.creator == msg.sender, "Only the task creator can perform this operation.");
        _;
    }

    constructor(
        IRegistryCoordinator _registryCoordinator,
        uint32 _taskResponseWindowBlock
    ) BLSSignatureChecker(_registryCoordinator) {
        TASK_RESPONSE_WINDOW_BLOCK = _taskResponseWindowBlock;
    }

    function initialize(
        IPauserRegistry _pauserRegistry,
        address initialOwner,
        address _aggregator
    ) public initializer {
        _initializePauser(_pauserRegistry, UNPAUSE_ALL);
        _transferOwnership(initialOwner);
        aggregator = _aggregator;
    }

    /* FUNCTIONS */
    // NOTE: this function creates new task, assigns it a taskId
    function createNewTask(
        uint8 taskType,
        uint32 quorumThresholdPercentage,
        bytes calldata quorumNumbers
    ) external payable paysTaskCreationFee(1) {
        // create a new task struct
        Task memory newTask;
        newTask.taskType = taskType;
        newTask.taskCreatedBlock = uint32(block.number);
        newTask.quorumThresholdPercentage = quorumThresholdPercentage;
        newTask.quorumNumbers = quorumNumbers;
        newTask.creator = payable(msg.sender);
        newTask.creationFee = msg.value;

        // store hash of task onchain, emit event, and increase taskNum
        allTaskHashes[latestTaskNum] = keccak256(abi.encode(newTask));
        emit NewTaskCreated(latestTaskNum, newTask);
        latestTaskNum = latestTaskNum + 1;

        // Refund any excess payment
        if (msg.value > taskCreationFee) {
            payable(msg.sender).transfer(msg.value - taskCreationFee);
        }
    }

    // NOTE: this function responds to existing tasks.
    function respondToTask(
        Task calldata task,
        OperatorResponse[] calldata responses
    ) external onlyAggregator {
        uint32 taskCreatedBlock = task.taskCreatedBlock;

        require(
            uint32(block.number) <=
                taskCreatedBlock + TASK_RESPONSE_WINDOW_BLOCK,
            "Aggregator has responded to the task too late"
        );

        require(
            responses.length > 0,
            "Operator responses should not be empty"
        );

        uint32 referenceTaskIndex = responses[0].response.referenceTaskIndex;

        require(
            keccak256(abi.encode(task)) == allTaskHashes[referenceTaskIndex],
            "supplied task does not match the one recorded in the contract"
        );

        require(
            allTaskResponses[referenceTaskIndex] == bytes32(0),
            "Aggregator has already responded to the task"
        );

        uint256 totalResult = 0;
        uint256 totalTimestamp = 0;
        uint256 totalWeight = 0;

        // check that the task is valid, hasn't been responsed yet, and is being responsed in time
        for (uint i = 0; i < responses.length; i++) {
            // Verify ECDSA signature
            bytes32 messageHash = keccak256(abi.encode(responses[i].response));
            bytes32 ethSignedMessageHash = messageHash.toEthSignedMessageHash();

            require(
                ethSignedMessageHash.recover(responses[i].signature) == responses[i].operator,
                "Invalid signature or operator address"
            );

            require(
                responses[i].response.referenceTaskIndex == referenceTaskIndex,
                "Aggregator response task indices should be consistent"
            );

            uint256 weight = sqrt(stakeRegistry.weightOfOperatorForQuorum(_TASKRESPONSE_QUORUM_NUMBER, responses[i].operator));

            totalWeight += weight;
            totalResult += responses[i].response.result * weight;
            totalTimestamp += responses[i].response.timeStamp * weight;
        }

        require(
            totalWeight > 0,
            "Operator doesn't have sufficient stakes"
        );

        uint256 avgResult = totalResult / totalWeight;
        uint256 avgTimestamp = totalTimestamp / totalWeight;

        TaskResponse memory taskResponse = TaskResponse({
            referenceTaskIndex: referenceTaskIndex,
            result: avgResult,
            timeStamp: avgTimestamp
        });


        TaskResponseMetadata memory taskResponseMetadata = TaskResponseMetadata(
            uint32(block.number)
        );
        // updating the storage with weighted task responsea
        allTaskResponses[taskResponse.referenceTaskIndex] = keccak256(
            abi.encode(taskResponse, taskResponseMetadata)
        );

        // emitting event
        emit TaskResponded(taskResponse, taskResponseMetadata);
    }


    function withdrawTaskFunds(uint32 taskNum, Task memory task) external onlyTaskCreator(taskNum, task) {
        require(
            allTaskResponses[taskNum] != bytes32(0),
            "Cannot withdraw funds for a completed task."
        );
        require(task.creationFee > 0, "No funds to withdraw.");

        uint amountToWithdraw = task.creationFee;
        task.creationFee = 0; // Prevent re-entrancy
        allTaskHashes[taskNum] = keccak256(abi.encode(task));

        (bool success, ) = task.creator.call{value: amountToWithdraw}("");
        require(success, "Withdrawal failed.");

        emit FundsWithdrawn(taskNum, msg.sender, amountToWithdraw);
    }

    function sqrt(uint256 y) internal pure returns (uint256 z) {
        if (y > 3) {
            z = y;
            uint256 x = y / 2 + 1;
            while (x < z) {
                z = x;
                x = (y / x + x) / 2;
            }
        } else if (y != 0) {
            z = 1;
        }
        // else z = 0
    }

    // Function to update the task creation fee, restricted to owner
    function updateTaskCreationFee(uint _newFee) external onlyOwner {
        taskCreationFee = _newFee;
    }

    // Function to allow the contract owner to withdraw the balance
    function withdraw() external onlyOwner {
        payable(msg.sender).transfer(address(this).balance);
    }

    function taskNumber() external view returns (uint32) {
        return latestTaskNum;
    }

    function getTaskResponseWindowBlock() external view returns (uint32) {
        return TASK_RESPONSE_WINDOW_BLOCK;
    }
}

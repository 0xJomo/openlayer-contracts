// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@eigenlayer/contracts/permissions/Pausable.sol";
import "@eigenlayer-middleware/src/interfaces/IServiceManager.sol";
import {IBLSApkRegistry} from "@eigenlayer-middleware/src/interfaces/IBLSApkRegistry.sol";
import {IStakeRegistry} from "@eigenlayer-middleware/src/interfaces/IStakeRegistry.sol";
import {RegistryCoordinator} from "@eigenlayer-middleware/src/RegistryCoordinator.sol";
import {BLSSignatureChecker, IRegistryCoordinator} from "@eigenlayer-middleware/src/BLSSignatureChecker.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {OperatorStateRetriever} from "@eigenlayer-middleware/src/OperatorStateRetriever.sol";
import "@eigenlayer-middleware/src/libraries/BN254.sol";
import "./IOpenOracleIdenticalAnswerTaskManager.sol";
import "./IOpenOracleVRFFeed.sol";

contract OpenOracleIdenticalAnswerTaskManager is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    OperatorStateRetriever,
    IOpenOracleIdenticalAnswerTaskManager
{
    using BN254 for BN254.G1Point;
    using ECDSA for bytes32;

    IStakeRegistry public immutable stakeRegistry;
    IBLSApkRegistry public immutable blsApkRegistry;
    BLSSignatureChecker public immutable blsSignatureChecker;

    /* CONSTANT */
    // The number of blocks from the task initialization within which the aggregator has to respond to
    uint32 public immutable TASK_RESPONSE_WINDOW_BLOCK;
    uint256 internal constant _THRESHOLD_DENOMINATOR = 100;
    uint8 private constant _TASKRESPONSE_QUORUM_NUMBER = 0;

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

    // Mapping to keep track of feed addresses
    mapping(address => bool) public isFeed;

    address public aggregator;

    /* MODIFIERS */
    modifier onlyAggregator() {
        require(msg.sender == aggregator, "Aggregator must be the caller");
        _;
    }

    modifier onlyFeed() {
        require(isFeed[msg.sender], "Caller is not feed");
        _;
    }

    modifier onlyTaskCreator(uint32 taskNum, Task memory task) {
        require(taskNum <= latestTaskNum, "task number not exists");
        require(
            keccak256(abi.encode(task)) == allTaskHashes[taskNum],
            "supplied task does not match the one recorded in the contract"
        );
        require(
            task.creator == msg.sender,
            "Only the task creator can perform this operation."
        );
        _;
    }

    constructor(
        IStakeRegistry _stakeRegistry,
        IBLSApkRegistry _BLSApkRegistry,
        uint32 _taskResponseWindowBlock,
        BLSSignatureChecker _blsSignatureChecker
    ) {
        stakeRegistry = _stakeRegistry;
        blsApkRegistry = _BLSApkRegistry;
        TASK_RESPONSE_WINDOW_BLOCK = _taskResponseWindowBlock;
        blsSignatureChecker = _blsSignatureChecker;
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

    function updateAggregator(address _aggregator) public onlyOwner {
        aggregator = _aggregator;
        emit AggregatorUpdated(aggregator);
    }

    /* FUNCTIONS */
    // NOTE: this function creates new task, assigns it a taskId
    function createNewTask(
        uint8 taskType,
        bytes calldata taskData,
        uint8 responderThreshold,
        uint96 stakeThreshold
    ) external onlyFeed {
        // create a new task struct
        Task memory newTask;
        newTask.taskType = taskType;
        newTask.taskData = taskData;
        newTask.taskCreatedBlock = uint32(block.number);
        newTask.stakeThreshold = stakeThreshold;
        newTask.responderThreshold = responderThreshold;
        newTask.creator = payable(msg.sender);
        newTask.creationFee = 0; // msg.value;

        // store hash of task onchain, emit event, and increase taskNum
        allTaskHashes[latestTaskNum] = keccak256(abi.encode(newTask));
        emit NewIdenticalAnswerTaskCreated(latestTaskNum, newTask);

        latestTaskNum = latestTaskNum + 1;

        // // Refund any excess payment
        // if (msg.value > taskCreationFee) {
        //     payable(msg.sender).transfer(msg.value - taskCreationFee);
        // }
    }

    // NOTE: this function responds to existing tasks.
    function respondToTask(
        Task calldata task,
        address[] calldata operators,
        AggregatedTaskResponse calldata response
    ) external onlyAggregator {
        uint32 taskCreatedBlock = task.taskCreatedBlock;

        require(
            uint32(block.number) <=
                taskCreatedBlock + TASK_RESPONSE_WINDOW_BLOCK,
            "Aggregator has responded to the task too late"
        );

        require(
            keccak256(abi.encode(task)) ==
                allTaskHashes[response.msg.referenceTaskIndex],
            "supplied task does not match the one recorded in the contract"
        );

        require(
            allTaskResponses[response.msg.referenceTaskIndex] == bytes32(0),
            "Aggregator has already responded to the task"
        );

        require(
            stakeRegistry.weightOfOperatorForQuorum(0, msg.sender) >=
                task.stakeThreshold,
            "Operator stake not meet minimum threshold"
        );

        // TODO: verify the aggregate signature
        bytes32 messageHash = keccak256(abi.encode(response.msg));
        (BN254.G1Point memory p, ) = blsApkRegistry
            .getRegisteredPubkey(operators[0]);
        for (uint i = 1; i < operators.length; i++) {
            (BN254.G1Point memory pk, ) = blsApkRegistry
                .getRegisteredPubkey(operators[i]);
            p = p.plus(pk);
        }
        (bool pairingSuccessful, bool signatureIsValid) = blsSignatureChecker.trySignatureAndApkVerification(
                messageHash,
                p,
                response.apkG2,
                response.aggregatedSignature
            );
        require(pairingSuccessful, "BLSSignatureChecker.checkSignatures: pairing precompile call failed");
        require(signatureIsValid, "BLSSignatureChecker.checkSignatures: signature is invalid");

        TaskResponseMetadata memory taskResponseMetadata = TaskResponseMetadata(
            uint32(block.number)
        );
        // updating the storage with weighted task response
        allTaskResponses[response.msg.referenceTaskIndex] = keccak256(
            abi.encode(response, taskResponseMetadata)
        );

        require(
            isFeed[task.creator],
            "Feed needs to register with task manager"
        );
        IOpenOracleVRFFeed vrfFeed = IOpenOracleVRFFeed(task.creator);
        AggregatedTaskResponse memory responseCopy = response;
        vrfFeed.saveLatestData(task, responseCopy, taskResponseMetadata);
        // emitting event
        emit TaskResponded(task, responseCopy, taskResponseMetadata);
    }

    function withdrawTaskFunds(
        uint32 taskNum,
        Task memory task
    ) external onlyTaskCreator(taskNum, task) {
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

    // Function to update the task creation fee, restricted to owner
    function updateTaskCreationFee(uint _newFee) external onlyOwner {
        taskCreationFee = _newFee;
    }

    // Function to add an address to the feed list
    function addToFeedlist(address _address) external onlyOwner {
        isFeed[_address] = true;
        emit AddressAddedToFeedlist(_address);
    }

    // Function to remove an address from the feed list
    function removeFromFeed(address _address) external onlyOwner {
        isFeed[_address] = false;
        emit AddressRemovedFromFeedlist(_address);
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

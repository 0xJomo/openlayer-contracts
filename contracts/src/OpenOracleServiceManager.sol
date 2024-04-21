// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@eigenlayer/contracts/libraries/BytesLib.sol";
import "@eigenlayer-middleware/src/ServiceManagerBase.sol";

/**
 * @title Primary entrypoint for procuring services from OpenOracle.
 * @author Layr Labs, Inc.
 */
contract OpenOracleServiceManager is ServiceManagerBase {
    using BytesLib for bytes;

<<<<<<< HEAD
    event OperatorAddedToRegistryWhitelist(address operator);

    event OperatorRemovedFromRegistryWhitelist(address operator);

    IOpenOracleTaskManager public immutable openOracleTaskManager;

    mapping(address => bool) public operatorIsWhitelistedForRegister;
=======
>>>>>>> Update serviceManager
    struct TaskManagerEntry {
        string chainName;
        address taskManagerAddress;
        bool isActive;
    }

    // Mapping from unique identifiers to task managers
    mapping(uint256 => TaskManagerEntry) public taskManagers;
    uint256 public taskManagerCount;

<<<<<<< HEAD
    /// @notice when applied to a function, ensures that the function is only callable by the `registryCoordinator`.
    modifier onlyOpenOracleTaskManager() {
        require(
            msg.sender == address(openOracleTaskManager),
            "onlyOpenOracleTaskManager: not from open oracle task manager"
        );
        _;
    }

    modifier onlyWhiteListedOperator(address operator) {
        require(
            operatorIsWhitelistedForRegister[operator],
            "onlyWhiteListedOperator: operator not in whitelist"
        );
        _;
    }

=======
>>>>>>> Update serviceManager
    constructor(
        IAVSDirectory _avsDirectory,
        IRegistryCoordinator _registryCoordinator,
        IStakeRegistry _stakeRegistry
    ) ServiceManagerBase(_avsDirectory, _registryCoordinator, _stakeRegistry) {
    }

    function initialize(address initialOwner) public virtual initializer {
        __ServiceManagerBase_init(initialOwner);
    }

    /// @notice Adds a new Task Manager entry to the list.
    /// @param chainName The chain associated with the Task Manager.
    /// @param taskManagerAddress The address of the Task Manager contract.
    function addTaskManager(string memory chainName, address taskManagerAddress) external onlyOwner {
        uint256 id = taskManagerCount++;
        taskManagers[id] = TaskManagerEntry(chainName, taskManagerAddress, true);
        emit TaskManagerAdded(id, chainName, taskManagerAddress);
    }

    /// @notice Removes a Task Manager entry from the list.
    /// @param id The identifier of the Task Manager to remove.
    function removeTaskManager(uint256 id) external onlyOwner {
        require(taskManagers[id].isActive, "Task Manager does not exist or already removed");
        taskManagers[id].isActive = false;
        emit TaskManagerRemoved(id);
    }

    /// @notice Emitted when a new Task Manager is added.
    /// @param id The identifier of the new Task Manager.
    /// @param chainName The Chain associated with the new Task Manager Net.
    /// @param taskManagerAddress The address of the new Task Manager.
    event TaskManagerAdded(uint256 id, string chainName, address taskManagerAddress);

    /// @notice Emitted when a Task Manager is removed.
    /// @param id The identifier of the removed Task Manager.
    event TaskManagerRemoved(uint256 id);
<<<<<<< HEAD

    /// @notice Called in the event of challenge resolution, in order to forward a call to the Slasher, which 'freezes' the `operator`.
    /// @dev The Slasher contract is under active development and its interface expected to change.
    ///      We recommend writing slashing logic without integrating with the Slasher at this point in time.
    function freezeOperator(
        address operatorAddr
    ) external onlyOpenOracleTaskManager {
        // slasher.freezeOperator(operatorAddr);
    }

    function registerOperatorToAVS(
        address operator,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    )
        public
        override
        onlyRegistryCoordinator
        onlyWhiteListedOperator(operator)
    {
        _avsDirectory.registerOperatorToAVS(operator, operatorSignature);
    }

    function addOperatorToRegistryWhitelist(
        address[] calldata operatorsToWhitelist
    ) external onlyOwner {
        uint256 whitelistLength = operatorsToWhitelist.length;
        for (uint256 i = 0; i < whitelistLength; ) {
            if (!operatorIsWhitelistedForRegister[operatorsToWhitelist[i]]) {
                operatorIsWhitelistedForRegister[
                    operatorsToWhitelist[i]
                ] = true;
                emit OperatorAddedToRegistryWhitelist(operatorsToWhitelist[i]);
            }
            unchecked {
                ++i;
            }
        }
    }

    function removeOperatorFromRegistryWhitelist(
        address[] calldata operatorsToRemoveFromWhitelist
    ) external onlyOwner {
        uint256 whitelistLength = operatorsToRemoveFromWhitelist.length;
        for (uint256 i = 0; i < whitelistLength; ) {
            if (
                operatorIsWhitelistedForRegister[
                    operatorsToRemoveFromWhitelist[i]
                ]
            ) {
                operatorIsWhitelistedForRegister[
                    operatorsToRemoveFromWhitelist[i]
                ] = false;
                emit OperatorRemovedFromRegistryWhitelist(
                    operatorsToRemoveFromWhitelist[i]
                );
            }
            unchecked {
                ++i;
            }
        }
    }
=======
>>>>>>> Update serviceManager
}

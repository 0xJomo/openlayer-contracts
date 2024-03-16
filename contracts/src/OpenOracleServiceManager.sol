// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@eigenlayer/contracts/libraries/BytesLib.sol";
import "./IOpenOracleTaskManager.sol";
import "@eigenlayer-middleware/src/ServiceManagerBase.sol";

/**
 * @title Primary entrypoint for procuring services from OpenOracle.
 * @author Layr Labs, Inc.
 */
contract OpenOracleServiceManager is ServiceManagerBase {
    using BytesLib for bytes;

    IOpenOracleTaskManager public immutable openOracleTaskManager;

    /// @notice when applied to a function, ensures that the function is only callable by the `registryCoordinator`.
    modifier onlyOpenOracleTaskManager() {
        require(
            msg.sender == address(openOracleTaskManager),
            "onlyOpenOracleTaskManager: not from open oracle task manager"
        );
        _;
    }

    constructor(
        IDelegationManager _delegationManager,
        IRegistryCoordinator _registryCoordinator,
        IStakeRegistry _stakeRegistry,
        IOpenOracleTaskManager _openOracleTaskManager
    )
        ServiceManagerBase(
            _delegationManager,
            _registryCoordinator,
            _stakeRegistry
        )
    {
        openOracleTaskManager = _openOracleTaskManager;
    }

    /// @notice Called in the event of challenge resolution, in order to forward a call to the Slasher, which 'freezes' the `operator`.
    /// @dev The Slasher contract is under active development and its interface expected to change.
    ///      We recommend writing slashing logic without integrating with the Slasher at this point in time.
    function freezeOperator(
        address operatorAddr
    ) external onlyOpenOracleTaskManager {
        // slasher.freezeOperator(operatorAddr);
    }
}

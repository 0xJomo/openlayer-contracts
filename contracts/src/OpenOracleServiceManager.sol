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

    event OperatorAddedToRegistryWhitelist(address operator);

    event OperatorRemovedFromRegistryWhitelist(address operator);

    IOpenOracleTaskManager public immutable openOracleTaskManager;

    mapping(address => bool) public operatorIsWhitelistedForRegister;

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

    constructor(
        IAVSDirectory _avsDirectory,
        IRegistryCoordinator _registryCoordinator,
        IStakeRegistry _stakeRegistry,
        IOpenOracleTaskManager _openOracleTaskManager
    ) ServiceManagerBase(_avsDirectory, _registryCoordinator, _stakeRegistry) {
        openOracleTaskManager = _openOracleTaskManager;
    }

    function initialize(address initialOwner) public virtual initializer {
        __ServiceManagerBase_init(initialOwner);
    }

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
}

// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import {IPauserRegistry} from "eigenlayer-contracts/src/contracts/interfaces/IPauserRegistry.sol";
import {ISignatureUtils} from "eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol";
import {ISocketUpdater} from "@eigenlayer-middleware/src/interfaces/ISocketUpdater.sol";
import {IBLSApkRegistry} from "@eigenlayer-middleware/src/interfaces/IBLSApkRegistry.sol";
import {IStakeRegistry} from "@eigenlayer-middleware/src/interfaces/IStakeRegistry.sol";
import {IIndexRegistry} from "@eigenlayer-middleware/src/interfaces/IIndexRegistry.sol";
import {IServiceManager} from "@eigenlayer-middleware/src/interfaces/IServiceManager.sol";
import {IRegistryCoordinator} from "@eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol";

import {EIP1271SignatureUtils} from "eigenlayer-contracts/src/contracts/libraries/EIP1271SignatureUtils.sol";
import {BitmapUtils} from "@eigenlayer-middleware/src/libraries/BitmapUtils.sol";
import {BN254} from "@eigenlayer-middleware/src/libraries/BN254.sol";

import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import {Initializable} from "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import {EIP712} from "@openzeppelin/contracts/utils/cryptography/draft-EIP712.sol";

import {Pausable} from "eigenlayer-contracts/src/contracts/permissions/Pausable.sol";
import {OpenOracleBridgeRegistryCoordinatorStorage} from "./OpenOracleBridgeRegistryCoordinatorStorage.sol";

/**
 * @title A `RegistryCoordinator` that has three registries:
 *      1) a `StakeRegistry` that keeps track of operators' stakes
 *      2) a `BLSApkRegistry` that keeps track of operators' BLS public keys and aggregate BLS public keys for each quorum
 *      3) an `IndexRegistry` that keeps track of an ordered list of operators for each quorum
 * 
 * @author Layr Labs, Inc.
 */
contract OpenOracleBridgeRegistryCoordinator is
    EIP712,
    Initializable,
    OwnableUpgradeable,
    OpenOracleBridgeRegistryCoordinatorStorage,
    ISocketUpdater,
    ISignatureUtils
{
    using BitmapUtils for *;
    using BN254 for BN254.G1Point;

    constructor(
        IStakeRegistry _stakeRegistry,
        IBLSApkRegistry _blsApkRegistry
    )
        OpenOracleBridgeRegistryCoordinatorStorage(_stakeRegistry, _blsApkRegistry)
        EIP712("AVSRegistryCoordinator", "v0.0.1")
    {
        _disableInitializers();
    }

    /**
     * @param _initialOwner will hold the owner role
     */
    function initialize(
        address _initialOwner
    ) external initializer {
        // Initialize roles
        _transferOwnership(_initialOwner);
    }

    /*******************************************************************************
                            EXTERNAL FUNCTIONS 
    *******************************************************************************/

    /**
     * @notice Registers msg.sender as an operator for one or more quorums. If any quorum exceeds its maximum
     * operator capacity after the operator is registered, this method will fail.
     * @param quorumNumbers is an ordered byte array containing the quorum numbers being registered for
     * @param socket is the socket of the operator (typically an IP address)
     * @param params contains the G1 & G2 public keys of the operator, and a signature proving their ownership
     * @param operatorSignature is the signature of the operator used by the AVS to register the operator in the delegation manager
     * @dev `params` is ignored if the caller has previously registered a public key
     * @dev `operatorSignature` is ignored if the operator's status is already REGISTERED
     */
    function registerOperator(
        bytes calldata quorumNumbers,
        string calldata socket,
        IBLSApkRegistry.PubkeyRegistrationParams calldata params,
        SignatureWithSaltAndExpiry memory operatorSignature
    ) external {}

    /**
     * @notice Registers msg.sender as an operator for one or more quorums. If any quorum reaches its maximum operator
     * capacity, `operatorKickParams` is used to replace an old operator with the new one.
     * @param quorumNumbers is an ordered byte array containing the quorum numbers being registered for
     * @param params contains the G1 & G2 public keys of the operator, and a signature proving their ownership
     * @param operatorKickParams used to determine which operator is removed to maintain quorum capacity as the
     * operator registers for quorums
     * @param churnApproverSignature is the signature of the churnApprover over the `operatorKickParams`
     * @param operatorSignature is the signature of the operator used by the AVS to register the operator in the delegation manager
     * @dev `params` is ignored if the caller has previously registered a public key
     * @dev `operatorSignature` is ignored if the operator's status is already REGISTERED
     */
    function registerOperatorWithChurn(
        bytes calldata quorumNumbers,
        string calldata socket,
        IBLSApkRegistry.PubkeyRegistrationParams calldata params,
        OperatorKickParam[] calldata operatorKickParams,
        SignatureWithSaltAndExpiry memory churnApproverSignature,
        SignatureWithSaltAndExpiry memory operatorSignature
    ) external {}

    /**
     * @notice Deregisters the caller from one or more quorums
     * @param quorumNumbers is an ordered byte array containing the quorum numbers being deregistered from
     */
    function deregisterOperator(
        bytes calldata quorumNumbers
    ) external {}

    /**
     * @notice Updates the StakeRegistry's view of one or more operators' stakes. If any operator
     * is found to be below the minimum stake for the quorum, they are deregistered.
     * @dev stakes are queried from the Eigenlayer core DelegationManager contract
     * @param operators a list of operator addresses to update
     */
    function updateOperators(
        address[] calldata operators
    ) external {}

    /**
     * @notice For each quorum in `quorumNumbers`, updates the StakeRegistry's view of ALL its registered operators' stakes.
     * Each quorum's `quorumUpdateBlockNumber` is also updated, which tracks the most recent block number when ALL registered
     * operators were updated.
     * @dev stakes are queried from the Eigenlayer core DelegationManager contract
     * @param operatorsPerQuorum for each quorum in `quorumNumbers`, this has a corresponding list of operators to update.
     * @dev Each list of operator addresses MUST be sorted in ascending order
     * @dev Each list of operator addresses MUST represent the entire list of registered operators for the corresponding quorum
     * @param quorumNumbers is an ordered byte array containing the quorum numbers being updated
     * @dev invariant: Each list of `operatorsPerQuorum` MUST be a sorted version of `IndexRegistry.getOperatorListAtBlockNumber`
     * for the corresponding quorum.
     * @dev note on race condition: if an operator registers/deregisters for any quorum in `quorumNumbers` after a txn to
     * this method is broadcast (but before it is executed), the method will fail
     */
    function updateOperatorsForQuorum(
        address[][] calldata operatorsPerQuorum,
        bytes calldata quorumNumbers
    ) external {}

    /**
     * @notice Updates the socket of the msg.sender given they are a registered operator
     * @param socket is the new socket of the operator
     */
    function updateSocket(string memory socket) external {}

    function updateQuorumsUpdateBlockNumber(
        bytes calldata quorumNumbers
    ) external onlyOwner {
        for (uint256 i = 0; i < quorumNumbers.length; i++) {
            quorumUpdateBlockNumber[uint8(quorumNumbers[i])] = block.number;
        }
    }

    function updateOperatorBitmapHistory(
        bytes32 operatorId,
        uint192 quorumBitmap
    ) external onlyOwner {
        _updateOperatorBitmap(operatorId, quorumBitmap);
    }

    function initializeOperatorsBitmapHistories(
        bytes32[] calldata operatorIds,
        uint192[] calldata quorumBitmaps
    ) external onlyOwner {
        require(operatorIds.length == quorumBitmaps.length,
            "operatorId and quorumBitmaps input length not matching");
        for (uint256 i = 0; i < operatorIds.length; i++) {
            _updateOperatorBitmap(operatorIds[i], quorumBitmaps[i]);
        }
    }

    function updateQuorumCount(
        uint8 _quorumCount
    ) external onlyOwner {
        quorumCount = _quorumCount;
    }

    /*******************************************************************************
                            EXTERNAL FUNCTIONS - EJECTOR
    *******************************************************************************/

    /**
     * @notice Forcibly deregisters an operator from one or more quorums
     * @param operator the operator to eject
     * @param quorumNumbers the quorum numbers to eject the operator from
     */
    function ejectOperator(
        address operator,
        bytes calldata quorumNumbers
    ) external {}

    /*******************************************************************************
                            EXTERNAL FUNCTIONS - OWNER
    *******************************************************************************/

    /**
     * @notice Creates a quorum and initializes it in each registry contract
     * @param operatorSetParams configures the quorum's max operator count and churn parameters
     * @param minimumStake sets the minimum stake required for an operator to register or remain
     * registered
     * @param strategyParams a list of strategies and multipliers used by the StakeRegistry to
     * calculate an operator's stake weight for the quorum
     */
    function createQuorum(
        OperatorSetParam memory operatorSetParams,
        uint96 minimumStake,
        IStakeRegistry.StrategyParams[] memory strategyParams
    ) external virtual onlyOwner {}

    /**
     * @notice Updates an existing quorum's configuration with a new max operator count
     * and operator churn parameters
     * @param quorumNumber the quorum number to update
     * @param operatorSetParams the new config
     * @dev only callable by the owner
     */
    function setOperatorSetParams(
        uint8 quorumNumber,
        OperatorSetParam memory operatorSetParams
    ) external {}

    /**
     * @notice Sets the churnApprover, which approves operator registration with churn
     * (see `registerOperatorWithChurn`)
     * @param _churnApprover the new churn approver
     * @dev only callable by the owner
     */
    function setChurnApprover(address _churnApprover) external onlyOwner {}

    /**
     * @notice Sets the ejector, which can force-deregister operators from quorums
     * @param _ejector the new ejector
     * @dev only callable by the owner
     */
    function setEjector(address _ejector) external onlyOwner {}

    /*******************************************************************************
                            INTERNAL FUNCTIONS
    *******************************************************************************/

    struct RegisterResults {
        uint32[] numOperatorsPerQuorum;
        uint96[] operatorStakes;
        uint96[] totalStakes;
    }

    /**
     * @notice Record an update to an operator's quorum bitmap.
     * @param newBitmap is the most up-to-date set of bitmaps the operator is registered for
     */
    function _updateOperatorBitmap(bytes32 operatorId, uint192 newBitmap) internal {

        uint256 historyLength = _operatorBitmapHistory[operatorId].length;

        if (historyLength == 0) {
            // No prior bitmap history - push our first entry
            _operatorBitmapHistory[operatorId].push(QuorumBitmapUpdate({
                updateBlockNumber: uint32(block.number),
                nextUpdateBlockNumber: 0,
                quorumBitmap: newBitmap
            }));
        } else {
            // We have prior history - fetch our last-recorded update
            QuorumBitmapUpdate storage lastUpdate = _operatorBitmapHistory[operatorId][historyLength - 1];

            /**
             * If the last update was made in the current block, update the entry.
             * Otherwise, push a new entry and update the previous entry's "next" field
             */
            if (lastUpdate.updateBlockNumber == uint32(block.number)) {
                lastUpdate.quorumBitmap = newBitmap;
            } else {
                lastUpdate.nextUpdateBlockNumber = uint32(block.number);
                _operatorBitmapHistory[operatorId].push(QuorumBitmapUpdate({
                    updateBlockNumber: uint32(block.number),
                    nextUpdateBlockNumber: 0,
                    quorumBitmap: newBitmap
                }));
            }
        }
        emit OperatorQuorumBitmapUpdated(operatorId, newBitmap);
    }

    /*******************************************************************************
                            VIEW FUNCTIONS
    *******************************************************************************/

    /// @notice Returns the operator set params for the given `quorumNumber`
    function getOperatorSetParams(
        uint8 quorumNumber
    ) external view returns (OperatorSetParam memory) {}

    /// @notice Returns the operator struct for the given `operator`
    function getOperator(
        address operator
    ) external view returns (OperatorInfo memory) {}

    /// @notice Returns the operatorId for the given `operator`
    function getOperatorId(address operator) external view returns (bytes32) {}

    /// @notice Returns the operator address for the given `operatorId`
    function getOperatorFromId(
        bytes32 operatorId
    ) external view returns (address) {}

    /// @notice Returns the status for the given `operator`
    function getOperatorStatus(
        address operator
    ) external view returns (IRegistryCoordinator.OperatorStatus) {}

    /**
     * @notice Returns the indices of the quorumBitmaps for the provided `operatorIds` at the given `blockNumber`
     * @dev Reverts if any of the `operatorIds` was not (yet) registered at `blockNumber`
     * @dev This function is designed to find proper inputs to the `getQuorumBitmapAtBlockNumberByIndex` function
     */
    function getQuorumBitmapIndicesAtBlockNumber(
        uint32 blockNumber,
        bytes32[] memory operatorIds
    ) external view returns (uint32[] memory) {}

    /**
     * @notice Returns the quorum bitmap for the given `operatorId` at the given `blockNumber` via the `index`,
     * reverting if `index` is incorrect
     * @dev This function is meant to be used in concert with `getQuorumBitmapIndicesAtBlockNumber`, which
     * helps off-chain processes to fetch the correct `index` input
     */
    function getQuorumBitmapAtBlockNumberByIndex(
        bytes32 operatorId,
        uint32 blockNumber,
        uint256 index
    ) external view returns (uint192) {
        QuorumBitmapUpdate memory quorumBitmapUpdate = _operatorBitmapHistory[operatorId][index];
        
        /**
         * Validate that the update is valid for the given blockNumber:
         * - blockNumber should be >= the update block number
         * - the next update block number should be either 0 or strictly greater than blockNumber
         */
        require(
            blockNumber >= quorumBitmapUpdate.updateBlockNumber, 
            "RegistryCoordinator.getQuorumBitmapAtBlockNumberByIndex: quorumBitmapUpdate is from after blockNumber"
        );
        require(
            quorumBitmapUpdate.nextUpdateBlockNumber == 0 || blockNumber < quorumBitmapUpdate.nextUpdateBlockNumber,
            "RegistryCoordinator.getQuorumBitmapAtBlockNumberByIndex: quorumBitmapUpdate is from before blockNumber"
        );

        return quorumBitmapUpdate.quorumBitmap;
    }

    /// @notice Returns the `index`th entry in the operator with `operatorId`'s bitmap history
    function getQuorumBitmapUpdateByIndex(
        bytes32 operatorId,
        uint256 index
    ) external view returns (QuorumBitmapUpdate memory) {}

    /// @notice Returns the current quorum bitmap for the given `operatorId` or 0 if the operator is not registered for any quorum
    function getCurrentQuorumBitmap(
        bytes32 operatorId
    ) external view returns (uint192) {}

    /// @notice Returns the length of the quorum bitmap history for the given `operatorId`
    function getQuorumBitmapHistoryLength(
        bytes32 operatorId
    ) external view returns (uint256) {}

    /// @notice Returns the number of registries
    function numRegistries() external view returns (uint256) {}

    /**
     * @notice Public function for the the churnApprover signature hash calculation when operators are being kicked from quorums
     * @param registeringOperatorId The id of the registering operator
     * @param operatorKickParams The parameters needed to kick the operator from the quorums that have reached their caps
     * @param salt The salt to use for the churnApprover's signature
     * @param expiry The desired expiry time of the churnApprover's signature
     */
    function calculateOperatorChurnApprovalDigestHash(
        address registeringOperator,
        bytes32 registeringOperatorId,
        OperatorKickParam[] memory operatorKickParams,
        bytes32 salt,
        uint256 expiry
    ) public view returns (bytes32) {}

    /**
     * @notice Returns the message hash that an operator must sign to register their BLS public key.
     * @param operator is the address of the operator registering their BLS public key
     */
    function pubkeyRegistrationMessageHash(
        address operator
    ) public view returns (BN254.G1Point memory) {}

    /// @dev need to override function here since its defined in both these contracts
    function owner()
        public
        view
        override(OwnableUpgradeable, IRegistryCoordinator)
        returns (address)
    {}

    function indexRegistry() external view override returns (IIndexRegistry) {}
}

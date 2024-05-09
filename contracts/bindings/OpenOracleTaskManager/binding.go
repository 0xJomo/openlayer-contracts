// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractOpenOracleTaskManager

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSSignatureCheckerNonSignerStakesAndSignature is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerNonSignerStakesAndSignature struct {
	NonSignerQuorumBitmapIndices []uint32
	NonSignerPubkeys             []BN254G1Point
	QuorumApks                   []BN254G1Point
	ApkG2                        BN254G2Point
	Sigma                        BN254G1Point
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// IBLSSignatureCheckerQuorumStakeTotals is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerQuorumStakeTotals struct {
	SignedStakeForQuorum []*big.Int
	TotalStakeForQuorum  []*big.Int
}

// IOpenOracleTaskManagerOperatorResponse is an auto generated low-level Go binding around an user-defined struct.
type IOpenOracleTaskManagerOperatorResponse struct {
	Operator  common.Address
	Response  IOpenOracleTaskManagerTaskResponse
	Signature []byte
}

// IOpenOracleTaskManagerTask is an auto generated low-level Go binding around an user-defined struct.
type IOpenOracleTaskManagerTask struct {
	TaskType           uint8
	TaskCreatedBlock   uint32
	ResponderThreshold uint8
	StakeThreshold     *big.Int
	Creator            common.Address
	CreationFee        *big.Int
}

// IOpenOracleTaskManagerTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type IOpenOracleTaskManagerTaskResponse struct {
	ReferenceTaskIndex uint32
	Result             *big.Int
	Timestamp          *big.Int
}

// IOpenOracleTaskManagerTaskResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type IOpenOracleTaskManagerTaskResponseMetadata struct {
	TaskResponsedBlock uint32
}

// IOpenOracleTaskManagerWeightedTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type IOpenOracleTaskManagerWeightedTaskResponse struct {
	ReferenceTaskIndex uint32
	Result             *big.Int
	Sd                 *big.Int
	Timestamp          *big.Int
}

// OperatorStateRetrieverCheckSignaturesIndices is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverCheckSignaturesIndices struct {
	NonSignerQuorumBitmapIndices []uint32
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// OperatorStateRetrieverOperator is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverOperator struct {
	Operator   common.Address
	OperatorId [32]byte
	Stake      *big.Int
}

// ContractOpenOracleTaskManagerMetaData contains all meta data concerning the ContractOpenOracleTaskManager contract.
var ContractOpenOracleTaskManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addToFeedlist\",\"inputs\":[{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"taskType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"responderThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"stakeThreshold\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isFeed\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestTaskNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeFromFeed\",\"inputs\":[{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIOpenOracleTaskManager.Task\",\"components\":[{\"name\":\"taskType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"responderThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"stakeThreshold\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"creationFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"responses\",\"type\":\"tuple[]\",\"internalType\":\"structIOpenOracleTaskManager.OperatorResponse[]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"response\",\"type\":\"tuple\",\"internalType\":\"structIOpenOracleTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"weightedTaskResponse\",\"type\":\"tuple\",\"internalType\":\"structIOpenOracleTaskManager.WeightedTaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sd\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskCreationFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAggregator\",\"inputs\":[{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTaskCreationFee\",\"inputs\":[{\"name\":\"_newFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawTaskFunds\",\"inputs\":[{\"name\":\"taskNum\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIOpenOracleTaskManager.Task\",\"components\":[{\"name\":\"taskType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"responderThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"stakeThreshold\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"creationFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AddressAddedToFeedlist\",\"inputs\":[{\"name\":\"_address\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AddressRemovedFromFeedlist\",\"inputs\":[{\"name\":\"_address\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AggregatorUpdated\",\"inputs\":[{\"name\":\"aggregator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FundsWithdrawn\",\"inputs\":[{\"name\":\"taskNum\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"creator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOpenOracleTaskManager.Task\",\"components\":[{\"name\":\"taskType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"responderThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"stakeThreshold\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"creationFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOpenOracleTaskManager.Task\",\"components\":[{\"name\":\"taskType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"responderThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"stakeThreshold\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"creationFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOpenOracleTaskManager.WeightedTaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sd\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOpenOracleTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x610120604052655af3107a400060c9553480156200001c57600080fd5b5060405162006271380380620062718339810160408190526200003f9162000201565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000099573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000bf919062000248565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000117573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013d919062000248565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000197573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001bd919062000248565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff1661010052506200026f565b6001600160a01b0381168114620001fe57600080fd5b50565b600080604083850312156200021557600080fd5b82516200022281620001e8565b602084015190925063ffffffff811681146200023d57600080fd5b809150509250929050565b6000602082840312156200025b57600080fd5b81516200026881620001e8565b9392505050565b60805160a05160c05160e05161010051615f706200030160003960008181610272015281816105a101526106f50152600081816105470152612b7801526000818161041c0152612d5a0152600081816104430152818161099e01528181612f1601526130bf01526000818161046a015281816117030152818161285a015281816129f20152612c150152615f706000f3fe608060405234801561001057600080fd5b50600436106101d85760003560e01c8063033ff004146101dd5780630c0f5792146101f2578063103642111461020557806310d67a2f14610218578063136439dd1461022b578063171f1d5b1461023e5780631ad431891461026d5780632191297d146102a9578063245a7bfc146102bc5780632cb223d5146102dc5780632d89f6fc1461030a5780633563b0d11461032a5780633ccfd60b1461034a578063416c7e5e14610352578063459b58b1146103655780634d075ce7146103785780634f739f741461038157806352e9408b146103a1578063595c6a67146103b45780635ac86ab7146103bc5780635c155662146103ef5780635c975abb1461040f5780635df4594614610417578063683048351461043e5780636d14a987146104655780636efb46361461048c578063715018a6146104ad57806372d18e8d146104b5578063886f1195146104c35780638b00ce7c146104d65780638da5cb5b146104e65780639fe4ee47146104ee578063b98d090814610501578063c0c53b8b1461050e578063cefdc1d414610521578063df5cf72314610542578063e58fdd0414610569578063f2fde38b1461058c578063f5c9899d1461059f578063fabc1cbc146105c5575b600080fd5b6101f06101eb3660046149bd565b6105d8565b005b6101f06102003660046149bd565b61062c565b6101f0610213366004614a1e565b61067d565b6101f06102263660046149bd565b610dc1565b6101f0610239366004614a92565b610e74565b61025161024c366004614c32565b610fa1565b6040805192151583529015156020830152015b60405180910390f35b6102947f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff9091168152602001610264565b6101f06102b7366004614ca7565b611107565b60ce546102cf906001600160a01b031681565b6040516102649190614cf2565b6102fc6102ea366004614d23565b60cc6020526000908152604090205481565b604051908152602001610264565b6102fc610318366004614d23565b60cb6020526000908152604090205481565b61033d610338366004614d40565b61124e565b6040516102649190614e9b565b6101f06116cd565b6101f0610360366004614ebc565b611701565b6101f0610373366004614ed9565b611876565b6102fc60c95481565b61039461038f366004614fc5565b611b94565b604051610264919061509a565b6101f06103af366004614a92565b61224f565b6101f061225c565b6103df6103ca366004615155565b606654600160ff9092169190911b9081161490565b6040519015158152602001610264565b6104026103fd366004615195565b612316565b6040516102649190615246565b6066546102fc565b6102cf7f000000000000000000000000000000000000000000000000000000000000000081565b6102cf7f000000000000000000000000000000000000000000000000000000000000000081565b6102cf7f000000000000000000000000000000000000000000000000000000000000000081565b61049f61049a3660046153dc565b6124c7565b6040516102649291906155a8565b6101f0613370565b60ca5463ffffffff16610294565b6065546102cf906001600160a01b031681565b60ca546102949063ffffffff1681565b6102cf613384565b6101f06104fc3660046149bd565b613393565b6097546103df9060ff1681565b6101f061051c3660046155f1565b6133e7565b61053461052f366004615631565b613522565b604051610264929190615668565b6102cf7f000000000000000000000000000000000000000000000000000000000000000081565b6103df6105773660046149bd565b60cd6020526000908152604090205460ff1681565b6101f061059a3660046149bd565b6136b4565b7f0000000000000000000000000000000000000000000000000000000000000000610294565b6101f06105d3366004614a92565b61372a565b6105e0613881565b6001600160a01b038116600081815260cd6020526040808220805460ff19166001179055517fe0fb1d5933f4eca7f5928d3c3c41d075398c1d1a4c4b09a5771a1c5c67eb1d869190a250565b610634613881565b6001600160a01b038116600081815260cd6020526040808220805460ff19169055517f28d57cd2f66bc88bf9eeda7d803691bff6e77bfdbf4212b41841a0d1b7189c5f9190a250565b60ce546001600160a01b031633146106dc5760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c657200000060448201526064015b60405180910390fd5b60006106ee6040860160208701614d23565b905061071a7f00000000000000000000000000000000000000000000000000000000000000008261569f565b63ffffffff164363ffffffff16111561078b5760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320726573706f6e64656420746f207468652060448201526c7461736b20746f6f206c61746560981b60648201526084016106d3565b826107e75760405162461bcd60e51b815260206004820152602660248201527f4f70657261746f7220726573706f6e7365732073686f756c64206e6f7420626560448201526520656d70747960d01b60648201526084016106d3565b60cb60006107f86020850185614d23565b63ffffffff1663ffffffff16815260200190815260200160002054856040516020016108249190615749565b60405160208183030381529060405280519060200120146108575760405162461bcd60e51b81526004016106d390615757565b600060cc816108696020860186614d23565b63ffffffff1663ffffffff16815260200190815260200160002054146108e65760405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b60648201526084016106d3565b60005b83811015610c09576000858583818110610905576109056157b4565b905060200281019061091791906157ca565b60200160405160200161092a91906157ea565b604051602081830303815290604052805190602001209050600061099a826040517b0ca2ba3432b932bab69029b4b3b732b21026b2b9b9b0b3b29d05199960211b6020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b90507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663a5679f5e8888868181106109dd576109dd6157b4565b90506020028101906109ef91906157ca565b6109fd9060208101906149bd565b6040518263ffffffff1660e01b8152600401610a199190614cf2565b6020604051808303816000875af1158015610a38573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a5c919061581d565b6001600160a01b0316610ad7888886818110610a7a57610a7a6157b4565b9050602002810190610a8c91906157ca565b610a9a90608081019061583a565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525086939250506138e09050565b6001600160a01b031614610b3b5760405162461bcd60e51b815260206004820152602560248201527f496e76616c6964207369676e6174757265206f72206f70657261746f72206164604482015264647265737360d81b60648201526084016106d3565b610b486020860186614d23565b63ffffffff16878785818110610b6057610b606157b4565b9050602002810190610b7291906157ca565b610b83906040810190602001614d23565b63ffffffff1614610bf45760405162461bcd60e51b815260206004820152603560248201527f41676772656761746f7220726573706f6e7365207461736b20696e6469636573604482015274081cda1bdd5b190818994818dbdb9cda5cdd195b9d605a1b60648201526084016106d3565b50508080610c0190615880565b9150506108e9565b50604080516020808201835263ffffffff4316825291519091610c309185918491016158cd565b6040516020818303038152906040528051906020012060cc6000856000016020810190610c5d9190614d23565b63ffffffff168152602081019190915260400160009081209190915560cd90610c8c60a0890160808a016149bd565b6001600160a01b0316815260208101919091526040016000205460ff16610d065760405162461bcd60e51b815260206004820152602860248201527f46656564206e6565647320746f2072656769737465722077697468207461736b6044820152671036b0b730b3b2b960c11b60648201526084016106d3565b6000610d1860a08801608089016149bd565b60405163d34f4fff60e01b81529091506001600160a01b0382169063d34f4fff90610d4b908a90889087906004016158ef565b600060405180830381600087803b158015610d6557600080fd5b505af1158015610d79573d6000803e3d6000fd5b505050507f5d0b79c46a19832eb8bcd7612cf0231afd4412288c9010e3fae97a0d3375fa6f878584604051610db0939291906158ef565b60405180910390a150505050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e14573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e38919061581d565b6001600160a01b0316336001600160a01b031614610e685760405162461bcd60e51b81526004016106d390615921565b610e7181613906565b50565b60655460405163237dfb4760e11b81526001600160a01b03909116906346fbf68e90610ea4903390600401614cf2565b602060405180830381865afa158015610ec1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ee59190615959565b610f015760405162461bcd60e51b81526004016106d390615976565b60665481811614610f755760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d707420604482015277746f20756e70617573652066756e6374696f6e616c69747960401b60648201526084016106d3565b60668190556040518181523390600080516020615e9b833981519152906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000187876000015188602001518860000151600060028110610fe957610fe96157b4565b60200201518951600160200201518a6020015160006002811061100e5761100e6157b4565b60200201518b6020015160016002811061102a5761102a6157b4565b602090810291909101518c518d8301516040516110879a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c6110aa91906159ac565b90506110f96110c36110bc88846139fd565b8690613a88565b6110cb613b10565b6110ef6110e0856110da613bd0565b906139fd565b6110e98c613bf1565b90613a88565b886201d4c0613c75565b909890975095505050505050565b33600090815260cd602052604090205460ff1661115b5760405162461bcd60e51b815260206004820152601260248201527110d85b1b195c881a5cc81b9bdd081999595960721b60448201526064016106d3565b6040805160c081018252600060a082015260ff85811682524363ffffffff166020808401919091526001600160601b038516606084015290851682840152336080830152915190916111af918391016159ce565b60408051601f19818403018152828252805160209182012060ca805463ffffffff908116600090815260cb90945293909220555416907fdeb7c64f4d939c7b58f45fb9cd5b5118e0b5920d43390c9ba0d47032754c3121906112129084906159ce565b60405180910390a260ca5461122e9063ffffffff16600161569f565b60ca805463ffffffff191663ffffffff9290921691909117905550505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015611290573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112b4919061581d565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa1580156112f6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061131a919061581d565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa15801561135c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611380919061581d565b9050600086516001600160401b0381111561139d5761139d614aab565b6040519080825280602002602001820160405280156113d057816020015b60608152602001906001900390816113bb5790505b50905060005b87518110156116c15760008882815181106113f3576113f36157b4565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015611454573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261147c9190810190615a2d565b905080516001600160401b0381111561149757611497614aab565b6040519080825280602002602001820160405280156114e257816020015b60408051606081018252600080825260208083018290529282015282526000199092019101816114b55790505b508484815181106114f5576114f56157b4565b602002602001018190525060005b81518110156116ab576040518060600160405280876001600160a01b03166347b314e8858581518110611538576115386157b4565b60200260200101516040518263ffffffff1660e01b815260040161155e91815260200190565b602060405180830381865afa15801561157b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061159f919061581d565b6001600160a01b031681526020018383815181106115bf576115bf6157b4565b60200260200101518152602001896001600160a01b031663fa28c6278585815181106115ed576115ed6157b4565b6020026020010151878f6040518463ffffffff1660e01b815260040161161593929190615abd565b602060405180830381865afa158015611632573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116569190615adc565b6001600160601b0316815250858581518110611674576116746157b4565b6020026020010151828151811061168d5761168d6157b4565b602002602001018190525080806116a390615880565b915050611503565b50505080806116b990615880565b9150506113d6565b50979650505050505050565b6116d5613881565b60405133904780156108fc02916000818181858888f19350505050158015610e71573d6000803e3d6000fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561175f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611783919061581d565b6001600160a01b0316336001600160a01b03161461182e5760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527b391037b3103a3432903932b3b4b9ba393ca1b7b7b93234b730ba37b960211b608482015260a4016106d3565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc906020015b60405180910390a150565b60ca548290829063ffffffff90811690831611156118cf5760405162461bcd60e51b81526020600482015260166024820152757461736b206e756d626572206e6f742065786973747360501b60448201526064016106d3565b63ffffffff8216600090815260cb60209081526040918290205491516118f7918491016159ce565b604051602081830303815290604052805190602001201461192a5760405162461bcd60e51b81526004016106d390615757565b60808101516001600160a01b031633146119a05760405162461bcd60e51b815260206004820152603160248201527f4f6e6c7920746865207461736b2063726561746f722063616e20706572666f7260448201527036903a3434b99037b832b930ba34b7b71760791b60648201526084016106d3565b63ffffffff8416600090815260cc6020526040902054611a165760405162461bcd60e51b815260206004820152602b60248201527f43616e6e6f742077697468647261772066756e647320666f72206120636f6d7060448201526a3632ba32b2103a30b9b59760a91b60648201526084016106d3565b60008360a0015111611a625760405162461bcd60e51b8152602060048201526015602482015274273790333ab73239903a37903bb4ba34323930bb9760591b60448201526064016106d3565b60a0830180516000909152604051611a7e9085906020016159ce565b60408051601f19818403018152828252805160209182012063ffffffff8916600090815260cb90925291812091909155608086015190916001600160a01b039091169083908381818185875af1925050503d8060008114611afb576040519150601f19603f3d011682016040523d82523d6000602084013e611b00565b606091505b5050905080611b465760405162461bcd60e51b81526020600482015260126024820152712bb4ba34323930bbb0b6103330b4b632b21760711b60448201526064016106d3565b6040805163ffffffff881681523360208201529081018390527ff440aec6b52895984d061d622e6edeba6210f7c3e059be920663140c084560d79060600160405180910390a1505050505050565b611b9c614872565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015611bdc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c00919061581d565b9050611c0a614872565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e90611c3a908b9089908990600401615af9565b600060405180830381865afa158015611c57573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611c7f9190810190615b43565b81526040516340e03a8160e11b81526001600160a01b038316906381c0750290611cb1908b908b908b90600401615bfa565b600060405180830381865afa158015611cce573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611cf69190810190615b43565b6040820152856001600160401b03811115611d1357611d13614aab565b604051908082528060200260200182016040528015611d4657816020015b6060815260200190600190039081611d315790505b50606082015260005b60ff8116871115612160576000856001600160401b03811115611d7457611d74614aab565b604051908082528060200260200182016040528015611d9d578160200160208202803683370190505b5083606001518360ff1681518110611db757611db76157b4565b602002602001018190525060005b868110156120605760008c6001600160a01b03166304ec63518a8a85818110611df057611df06157b4565b905060200201358e88600001518681518110611e0e57611e0e6157b4565b60200260200101516040518463ffffffff1660e01b8152600401611e3493929190615c23565b602060405180830381865afa158015611e51573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e759190615c3f565b90506001600160c01b038116611f185760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527b3132903932b3b4b9ba32b932b21030ba10313637b1b5b73ab6b132b960211b608482015260a4016106d3565b8a8a8560ff16818110611f2d57611f2d6157b4565b6001600160c01b03841692013560f81c9190911c60019081161415905061204d57856001600160a01b031663dd9846b98a8a85818110611f6f57611f6f6157b4565b905060200201358d8d8860ff16818110611f8b57611f8b6157b4565b9050013560f81c60f81b60f81c8f6040518463ffffffff1660e01b8152600401611fb793929190615abd565b602060405180830381865afa158015611fd4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ff89190615c68565b85606001518560ff1681518110612011576120116157b4565b6020026020010151848151811061202a5761202a6157b4565b63ffffffff909216602092830291909101909101528261204981615880565b9350505b508061205881615880565b915050611dc5565b506000816001600160401b0381111561207b5761207b614aab565b6040519080825280602002602001820160405280156120a4578160200160208202803683370190505b50905060005b828110156121255784606001518460ff16815181106120cb576120cb6157b4565b602002602001015181815181106120e4576120e46157b4565b60200260200101518282815181106120fe576120fe6157b4565b63ffffffff909216602092830291909101909101528061211d81615880565b9150506120aa565b508084606001518460ff1681518110612140576121406157b4565b60200260200101819052505050808061215890615c85565b915050611d4f565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156121a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121c5919061581d565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c906121f8908b908b908e90600401615ca5565b600060405180830381865afa158015612215573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261223d9190810190615b43565b60208301525098975050505050505050565b612257613881565b60c955565b60655460405163237dfb4760e11b81526001600160a01b03909116906346fbf68e9061228c903390600401614cf2565b602060405180830381865afa1580156122a9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122cd9190615959565b6122e95760405162461bcd60e51b81526004016106d390615976565b60001960668190556040519081523390600080516020615e9b8339815191529060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b8152600401612348929190615ccf565b600060405180830381865afa158015612365573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261238d9190810190615b43565b9050600084516001600160401b038111156123aa576123aa614aab565b6040519080825280602002602001820160405280156123d3578160200160208202803683370190505b50905060005b85518110156124bd57866001600160a01b03166304ec6351878381518110612403576124036157b4565b60200260200101518786858151811061241e5761241e6157b4565b60200260200101516040518463ffffffff1660e01b815260040161244493929190615c23565b602060405180830381865afa158015612461573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124859190615c3f565b6001600160c01b03168282815181106124a0576124a06157b4565b6020908102919091010152806124b581615880565b9150506123d9565b5095945050505050565b6124cf61489a565b60008461252c5760405162461bcd60e51b81526020600482015260376024820152600080516020615f1b8339815191526044820152761c995cce88195b5c1d1e481c5d5bdc9d5b481a5b9c1d5d604a1b60648201526084016106d3565b60408301515185148015612544575060a08301515185145b8015612554575060c08301515185145b8015612564575060e08301515185145b6125ce5760405162461bcd60e51b81526020600482015260416024820152600080516020615f1b83398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a4016106d3565b825151602084015151146126465760405162461bcd60e51b815260206004820152604460248201819052600080516020615f1b833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a4016106d3565b4363ffffffff168463ffffffff16106126b45760405162461bcd60e51b815260206004820152603c6024820152600080516020615f1b83398151915260448201527b7265733a20696e76616c6964207265666572656e636520626c6f636b60201b60648201526084016106d3565b60408051808201909152600080825260208201526126d061489a565b866001600160401b038111156126e8576126e8614aab565b604051908082528060200260200182016040528015612711578160200160208202803683370190505b506020820152866001600160401b0381111561272f5761272f614aab565b604051908082528060200260200182016040528015612758578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b0381111561278c5761278c614aab565b6040519080825280602002602001820160405280156127b5578160200160208202803683370190505b5081526020860151516001600160401b038111156127d5576127d5614aab565b6040519080825280602002602001820160405280156127fe578160200160208202803683370190505b50816020018190525060006128d08a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa1580156128a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128cb9190615d23565b613e99565b905060005b876020015151811015612b545761291a886020015182815181106128fb576128fb6157b4565b6020026020010151805160009081526020918201519091526040902090565b83602001518281518110612930576129306157b4565b602090810291909101015280156129f0576020830151612951600183615d40565b81518110612961576129616157b4565b602002602001015160001c83602001518281518110612982576129826157b4565b602002602001015160001c116129f0576040805162461bcd60e51b8152602060048201526024810191909152600080516020615f1b83398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f7274656460648201526084016106d3565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec635184602001518381518110612a3557612a356157b4565b60200260200101518b8b600001518581518110612a5457612a546157b4565b60200260200101516040518463ffffffff1660e01b8152600401612a7a93929190615c23565b602060405180830381865afa158015612a97573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612abb9190615c3f565b6001600160c01b031683600001518281518110612ada57612ada6157b4565b602002602001018181525050612b406110bc612b148486600001518581518110612b0657612b066157b4565b602002602001015116613f18565b8a602001518481518110612b2a57612b2a6157b4565b6020026020010151613f4390919063ffffffff16565b945080612b4c81615880565b9150506128d5565b5050612b5f8361401b565b60975490935060ff16600081612b76576000612bf8565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa158015612bd4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612bf89190615d57565b905060005b8a811015613243578215612d58578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f86818110612c5457612c546157b4565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015612c94573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612cb89190615d57565b612cc29190615d70565b11612d585760405162461bcd60e51b81526020600482015260666024820152600080516020615f1b83398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c4016106d3565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d84818110612d9957612d996157b4565b9050013560f81c60f81b60f81c8c8c60a001518581518110612dbd57612dbd6157b4565b60200260200101516040518463ffffffff1660e01b8152600401612de393929190615d88565b602060405180830381865afa158015612e00573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e249190615da9565b6001600160401b031916612e478a6040015183815181106128fb576128fb6157b4565b6001600160401b03191614612ee25760405162461bcd60e51b81526020600482015260616024820152600080516020615f1b83398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c4016106d3565b612f1289604001518281518110612efb57612efb6157b4565b602002602001015187613a8890919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d84818110612f5557612f556157b4565b9050013560f81c60f81b60f81c8c8c60c001518581518110612f7957612f796157b4565b60200260200101516040518463ffffffff1660e01b8152600401612f9f93929190615d88565b602060405180830381865afa158015612fbc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612fe09190615adc565b85602001518281518110612ff657612ff66157b4565b6001600160601b03909216602092830291909101820152850151805182908110613022576130226157b4565b602002602001015185600001518281518110613040576130406157b4565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a602001515181101561322e576130b88660000151828151811061308a5761308a6157b4565b60200260200101518f8f868181106130a4576130a46157b4565b600192013560f81c9290921c811614919050565b1561321c577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f868181106130fe576130fe6157b4565b9050013560f81c60f81b60f81c8e89602001518581518110613122576131226157b4565b60200260200101518f60e001518881518110613140576131406157b4565b60200260200101518781518110613159576131596157b4565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa1580156131bd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131e19190615adc565b87518051859081106131f5576131f56157b4565b602002602001018181516132099190615dd3565b6001600160601b03169052506001909101905b8061322681615880565b915050613064565b5050808061323b90615880565b915050612bfd565b50505060008061325d8c868a606001518b60800151610fa1565b91509150816132ce5760405162461bcd60e51b81526020600482015260436024820152600080516020615f1b83398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a4016106d3565b8061332b5760405162461bcd60e51b81526020600482015260396024820152600080516020615f1b8339815191526044820152781c995cce881cda59db985d1d5c99481a5cc81a5b9d985b1a59603a1b60648201526084016106d3565b50506000878260200151604051602001613346929190615dfb565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b613378613881565b61338260006140aa565b565b6033546001600160a01b031690565b61339b613881565b60ce80546001600160a01b0319166001600160a01b0383169081179091556040517f602cec4b1583b07d071161da5eb9589444d2459201e2fab7753dc941e9351c219161186b91614cf2565b600054610100900460ff16158080156134075750600054600160ff909116105b806134215750303b158015613421575060005460ff166001145b6134845760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016106d3565b6000805460ff1916600117905580156134a7576000805461ff0019166101001790555b6134b28460006140fc565b6134bb836140aa565b60ce80546001600160a01b0319166001600160a01b038416179055801561351c576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b604080516001808252818301909252600091606091839160208083019080368337019050509050848160008151811061355d5761355d6157b4565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e906135999088908690600401615ccf565b600060405180830381865afa1580156135b6573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526135de9190810190615b43565b6000815181106135f0576135f06157b4565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa15801561365c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906136809190615c3f565b6001600160c01b031690506000613696826141d4565b9050816136a48a838a61124e565b9550955050505050935093915050565b6136bc613881565b6001600160a01b0381166137215760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016106d3565b610e71816140aa565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561377d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137a1919061581d565b6001600160a01b0316336001600160a01b0316146137d15760405162461bcd60e51b81526004016106d390615921565b60665419811960665419161461384a5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d706044820152777420746f2070617573652066756e6374696f6e616c69747960401b60648201526084016106d3565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610f96565b3361388a613384565b6001600160a01b0316146133825760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016106d3565b60008060006138ef85856142a0565b915091506138fc81614310565b5090505b92915050565b6001600160a01b0381166139945760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a4016106d3565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b613a056148b4565b613a0d6148ce565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa9050808015613a4057613a42565bfe5b5080613a805760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b60448201526064016106d3565b505092915050565b613a906148b4565b613a986148ec565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa9050808015613a40575080613a805760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b60448201526064016106d3565b613b1861490a565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b613bd86148b4565b5060408051808201909152600181526002602082015290565b613bf96148b4565b60008080613c15600080516020615ebb833981519152866159ac565b90505b613c21816144c6565b9093509150600080516020615ebb833981519152828309831415613c5b576040805180820190915290815260208101919091529392505050565b600080516020615ebb833981519152600182089050613c18565b604080518082018252868152602080820186905282518084019093528683528201849052600091829190613ca761492f565b60005b6002811015613e6c576000613cc0826006615e43565b9050848260028110613cd457613cd46157b4565b60200201515183613ce6836000615d70565b600c8110613cf657613cf66157b4565b6020020152848260028110613d0d57613d0d6157b4565b60200201516020015183826001613d249190615d70565b600c8110613d3457613d346157b4565b6020020152838260028110613d4b57613d4b6157b4565b6020020151515183613d5e836002615d70565b600c8110613d6e57613d6e6157b4565b6020020152838260028110613d8557613d856157b4565b6020020151516001602002015183613d9e836003615d70565b600c8110613dae57613dae6157b4565b6020020152838260028110613dc557613dc56157b4565b602002015160200151600060028110613de057613de06157b4565b602002015183613df1836004615d70565b600c8110613e0157613e016157b4565b6020020152838260028110613e1857613e186157b4565b602002015160200151600160028110613e3357613e336157b4565b602002015183613e44836005615d70565b600c8110613e5457613e546157b4565b60200201525080613e6481615880565b915050613caa565b50613e7561494e565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b600080613ea584614548565b9050808360ff166001901b11613f115760405162461bcd60e51b815260206004820152603f6024820152600080516020615efb83398151915260448201527f69746d61703a206269746d61702065786365656473206d61782076616c75650060648201526084016106d3565b9392505050565b6000805b821561390057613f2d600184615d40565b9092169180613f3b81615e62565b915050613f1c565b613f4b6148b4565b6102008261ffff1610613f935760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b60448201526064016106d3565b8161ffff1660011415613fa7575081613900565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff161061401057600161ffff871660ff83161c81161415613ff357613ff08484613a88565b93505b613ffd8384613a88565b92506201fffe600192831b169101613fc3565b509195945050505050565b6140236148b4565b815115801561403457506020820151155b15614052575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020615ebb833981519152846020015161408591906159ac565b61409d90600080516020615ebb833981519152615d40565b905292915050565b919050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6065546001600160a01b031615801561411d57506001600160a01b03821615155b61419f5760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a4016106d3565b60668190556040518181523390600080516020615e9b8339815191529060200160405180910390a26141d082613906565b5050565b60606000806141e284613f18565b61ffff166001600160401b038111156141fd576141fd614aab565b6040519080825280601f01601f191660200182016040528015614227576020820181803683370190505b5090506000805b82518210801561423f575061010081105b15614296576001811b935085841615614286578060f81b838381518110614268576142686157b4565b60200101906001600160f81b031916908160001a9053508160010191505b61428f81615880565b905061422e565b5090949350505050565b6000808251604114156142d75760208301516040840151606085015160001a6142cb878285856146b1565b94509450505050614309565b82516040141561430157602083015160408401516142f6868383614794565b935093505050614309565b506000905060025b9250929050565b600081600481111561432457614324615e84565b141561432d5750565b600181600481111561434157614341615e84565b141561438a5760405162461bcd60e51b815260206004820152601860248201527745434453413a20696e76616c6964207369676e617475726560401b60448201526064016106d3565b600281600481111561439e5761439e615e84565b14156143ec5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016106d3565b600381600481111561440057614400615e84565b14156144595760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016106d3565b600481600481111561446d5761446d615e84565b1415610e715760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b60648201526084016106d3565b60008080600080516020615ebb8339815191526003600080516020615ebb83398151915286600080516020615ebb83398151915288890909089050600061453c827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020615ebb8339815191526147cd565b91959194509092505050565b6000610100825111156145bf5760405162461bcd60e51b815260206004820152604460248201819052600080516020615efb833981519152908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a4016106d3565b81516145cd57506000919050565b600080836000815181106145e3576145e36157b4565b0160200151600160f89190911c81901b92505b84518110156146a857848181518110614611576146116157b4565b0160200151600160f89190911c1b91508282116146945760405162461bcd60e51b81526020600482015260476024820152600080516020615efb83398151915260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a4016106d3565b918117916146a181615880565b90506145f6565b50909392505050565b6000806fa2a8918ca85bafe22016d0b997e4df60600160ff1b038311156146de575060009050600361478b565b8460ff16601b141580156146f657508460ff16601c14155b15614707575060009050600461478b565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561475b573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166147845760006001925092505061478b565b9150600090505b94509492505050565b6000806001600160ff1b038316816147b160ff86901c601b615d70565b90506147bf878288856146b1565b935093505050935093915050565b6000806147d861494e565b6147e061496c565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa9250828015613a405750826148675760405162461bcd60e51b815260206004820152601a602482015279424e3235342e6578704d6f643a2063616c6c206661696c75726560301b60448201526064016106d3565b505195945050505050565b6040518060800160405280606081526020016060815260200160608152602001606081525090565b604051806040016040528060608152602001606081525090565b604051806040016040528060008152602001600081525090565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b604051806040016040528061491d61498a565b815260200161492a61498a565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b0381168114610e7157600080fd5b6000602082840312156149cf57600080fd5b8135613f11816149a8565b60008083601f8401126149ec57600080fd5b5081356001600160401b03811115614a0357600080fd5b6020830191508360208260051b850101111561430957600080fd5b600080600080848603610160811215614a3657600080fd5b60c0811215614a4457600080fd5b85945060c08501356001600160401b03811115614a6057600080fd5b614a6c888289016149da565b909550935050608060df1982011215614a8457600080fd5b5092959194509260e0019150565b600060208284031215614aa457600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715614ae357614ae3614aab565b60405290565b60405160c081016001600160401b0381118282101715614ae357614ae3614aab565b60405161010081016001600160401b0381118282101715614ae357614ae3614aab565b604051601f8201601f191681016001600160401b0381118282101715614b5657614b56614aab565b604052919050565b600060408284031215614b7057600080fd5b614b78614ac1565b9050813581526020820135602082015292915050565b600082601f830112614b9f57600080fd5b604080519081016001600160401b0381118282101715614bc157614bc1614aab565b8060405250806040840185811115614bd857600080fd5b845b81811015614010578035835260209283019201614bda565b600060808284031215614c0457600080fd5b614c0c614ac1565b9050614c188383614b8e565b8152614c278360408401614b8e565b602082015292915050565b6000806000806101208587031215614c4957600080fd5b84359350614c5a8660208701614b5e565b9250614c698660608701614bf2565b9150614c788660e08701614b5e565b905092959194509250565b60ff81168114610e7157600080fd5b6001600160601b0381168114610e7157600080fd5b600080600060608486031215614cbc57600080fd5b8335614cc781614c83565b92506020840135614cd781614c83565b91506040840135614ce781614c92565b809150509250925092565b6001600160a01b0391909116815260200190565b63ffffffff81168114610e7157600080fd5b80356140a581614d06565b600060208284031215614d3557600080fd5b8135613f1181614d06565b600080600060608486031215614d5557600080fd5b8335614d60816149a8565b92506020848101356001600160401b0380821115614d7d57600080fd5b818701915087601f830112614d9157600080fd5b813581811115614da357614da3614aab565b614db5601f8201601f19168501614b2e565b91508082528884828501011115614dcb57600080fd5b8084840185840137600084828401015250809450505050614dee60408501614d18565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015614e8d578385038a52825180518087529087019087870190845b81811015614e7857835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101614e34565b50509a87019a95505091850191600101614e16565b509298975050505050505050565b602081526000613f116020830184614df7565b8015158114610e7157600080fd5b600060208284031215614ece57600080fd5b8135613f1181614eae565b60008082840360e0811215614eed57600080fd5b8335614ef881614d06565b925060c0601f1982011215614f0c57600080fd5b50614f15614ae9565b6020840135614f2381614c83565b81526040840135614f3381614d06565b60208201526060840135614f4681614c83565b60408201526080840135614f5981614c92565b606082015260a0840135614f6c816149a8565b608082015260c0939093013560a08401525092909150565b60008083601f840112614f9657600080fd5b5081356001600160401b03811115614fad57600080fd5b60208301915083602082850101111561430957600080fd5b60008060008060008060808789031215614fde57600080fd5b8635614fe9816149a8565b95506020870135614ff981614d06565b945060408701356001600160401b038082111561501557600080fd5b6150218a838b01614f84565b9096509450606089013591508082111561503a57600080fd5b5061504789828a016149da565b979a9699509497509295939492505050565b600081518084526020808501945080840160005b8381101561508f57815163ffffffff168752958201959082019060010161506d565b509495945050505050565b6000602080835283516080828501526150b660a0850182615059565b905081850151601f19808684030160408701526150d38383615059565b925060408701519150808684030160608701526150f08383615059565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b828110156151475784878303018452615135828751615059565b9588019593880193915060010161511b565b509998505050505050505050565b60006020828403121561516757600080fd5b8135613f1181614c83565b60006001600160401b0382111561518b5761518b614aab565b5060051b60200190565b6000806000606084860312156151aa57600080fd5b83356151b5816149a8565b92506020848101356001600160401b038111156151d157600080fd5b8501601f810187136151e257600080fd5b80356151f56151f082615172565b614b2e565b81815260059190911b8201830190838101908983111561521457600080fd5b928401925b8284101561523257833582529284019290840190615219565b8096505050505050614dee60408501614d18565b6020808252825182820181905260009190848201906040850190845b8181101561527e57835183529284019291840191600101615262565b50909695505050505050565b600082601f83011261529b57600080fd5b813560206152ab6151f083615172565b82815260059290921b840181019181810190868411156152ca57600080fd5b8286015b848110156152ee5780356152e181614d06565b83529183019183016152ce565b509695505050505050565b600082601f83011261530a57600080fd5b8135602061531a6151f083615172565b82815260069290921b8401810191818101908684111561533957600080fd5b8286015b848110156152ee5761534f8882614b5e565b83529183019160400161533d565b600082601f83011261536e57600080fd5b8135602061537e6151f083615172565b82815260059290921b8401810191818101908684111561539d57600080fd5b8286015b848110156152ee5780356001600160401b038111156153c05760008081fd5b6153ce8986838b010161528a565b8452509183019183016153a1565b6000806000806000608086880312156153f457600080fd5b8535945060208601356001600160401b038082111561541257600080fd5b61541e89838a01614f84565b90965094506040880135915061543382614d06565b9092506060870135908082111561544957600080fd5b90870190610180828a03121561545e57600080fd5b615466614b0b565b82358281111561547557600080fd5b6154818b82860161528a565b82525060208301358281111561549657600080fd5b6154a28b8286016152f9565b6020830152506040830135828111156154ba57600080fd5b6154c68b8286016152f9565b6040830152506154d98a60608501614bf2565b60608201526154eb8a60e08501614b5e565b60808201526101208301358281111561550357600080fd5b61550f8b82860161528a565b60a0830152506101408301358281111561552857600080fd5b6155348b82860161528a565b60c0830152506101608301358281111561554d57600080fd5b6155598b82860161535d565b60e0830152508093505050509295509295909350565b600081518084526020808501945080840160005b8381101561508f5781516001600160601b031687529582019590820190600101615583565b60408152600083516040808401526155c3608084018261556f565b90506020850151603f198483030160608501526155e0828261556f565b925050508260208301529392505050565b60008060006060848603121561560657600080fd5b8335615611816149a8565b92506020840135615621816149a8565b91506040840135614ce7816149a8565b60008060006060848603121561564657600080fd5b8335615651816149a8565b9250602084013591506040840135614ce781614d06565b8281526040602082015260006156816040830184614df7565b949350505050565b634e487b7160e01b600052601160045260246000fd5b600063ffffffff8083168185168083038211156156be576156be615689565b01949350505050565b80356156d281614c83565b60ff16825260208101356156e581614d06565b63ffffffff16602083015260408101356156fe81614c83565b60ff166040830152606081013561571481614c92565b6001600160601b031660608301526080810135615730816149a8565b6001600160a01b0316608083015260a090810135910152565b60c0810161390082846156c7565b6020808252603d908201527f737570706c696564207461736b20646f6573206e6f74206d617463682074686560408201527f206f6e65207265636f7264656420696e2074686520636f6e7472616374000000606082015260800190565b634e487b7160e01b600052603260045260246000fd5b60008235609e198336030181126157e057600080fd5b9190910192915050565b6060810182356157f981614d06565b63ffffffff8116835250602083013560208301526040830135604083015292915050565b60006020828403121561582f57600080fd5b8151613f11816149a8565b6000808335601e1984360301811261585157600080fd5b8301803591506001600160401b0382111561586b57600080fd5b60200191503681900382131561430957600080fd5b600060001982141561589457615894615689565b5060010190565b80356158a681614d06565b63ffffffff1682526020818101359083015260408082013590830152606090810135910152565b60a081016158db828561589b565b63ffffffff83511660808301529392505050565b61016081016158fe82866156c7565b61590b60c083018561589b565b63ffffffff835116610140830152949350505050565b6020808252602a90820152600080516020615edb83398151915260408201526939903ab73830bab9b2b960b11b606082015260800190565b60006020828403121561596b57600080fd5b8151613f1181614eae565b6020808252602890820152600080516020615edb83398151915260408201526739903830bab9b2b960c11b606082015260800190565b6000826159c957634e487b7160e01b600052601260045260246000fd5b500690565b815160ff908116825260208084015163ffffffff1690830152604080840151909116908201526060808301516001600160601b0316908201526080808301516001600160a01b03169082015260a0918201519181019190915260c00190565b60006020808385031215615a4057600080fd5b82516001600160401b03811115615a5657600080fd5b8301601f81018513615a6757600080fd5b8051615a756151f082615172565b81815260059190911b82018301908381019087831115615a9457600080fd5b928401925b82841015615ab257835182529284019290840190615a99565b979650505050505050565b92835260ff91909116602083015263ffffffff16604082015260600190565b600060208284031215615aee57600080fd5b8151613f1181614c92565b63ffffffff84168152604060208201819052810182905260006001600160fb1b03831115615b2657600080fd5b8260051b8085606085013760009201606001918252509392505050565b60006020808385031215615b5657600080fd5b82516001600160401b03811115615b6c57600080fd5b8301601f81018513615b7d57600080fd5b8051615b8b6151f082615172565b81815260059190911b82018301908381019087831115615baa57600080fd5b928401925b82841015615ab2578351615bc281614d06565b82529284019290840190615baf565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff84168152604060208201526000615c1a604083018486615bd1565b95945050505050565b92835263ffffffff918216602084015216604082015260600190565b600060208284031215615c5157600080fd5b81516001600160c01b0381168114613f1157600080fd5b600060208284031215615c7a57600080fd5b8151613f1181614d06565b600060ff821660ff811415615c9c57615c9c615689565b60010192915050565b604081526000615cb9604083018587615bd1565b905063ffffffff83166020830152949350505050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b81811015615d1657845183529383019391830191600101615cfa565b5090979650505050505050565b600060208284031215615d3557600080fd5b8151613f1181614c83565b600082821015615d5257615d52615689565b500390565b600060208284031215615d6957600080fd5b5051919050565b60008219821115615d8357615d83615689565b500190565b60ff93909316835263ffffffff918216602084015216604082015260600190565b600060208284031215615dbb57600080fd5b81516001600160401b031981168114613f1157600080fd5b60006001600160601b0383811690831681811015615df357615df3615689565b039392505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b83811015615e3657815185529382019390820190600101615e1a565b5092979650505050505050565b6000816000190483118215151615615e5d57615e5d615689565b500290565b600061ffff80831681811415615e7a57615e7a615689565b6001019392505050565b634e487b7160e01b600052602160045260246000fdfeab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd476d73672e73656e646572206973206e6f74207065726d697373696f6e656420614269746d61705574696c732e6f72646572656442797465734172726179546f42424c535369676e6174757265436865636b65722e636865636b5369676e617475a26469706673582212205617325b5b5fd25f45c4af733e77b2da0d333744665d98d61c577a1b2b3d74cc64736f6c634300080c0033",
}

// ContractOpenOracleTaskManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractOpenOracleTaskManagerMetaData.ABI instead.
var ContractOpenOracleTaskManagerABI = ContractOpenOracleTaskManagerMetaData.ABI

// ContractOpenOracleTaskManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractOpenOracleTaskManagerMetaData.Bin instead.
var ContractOpenOracleTaskManagerBin = ContractOpenOracleTaskManagerMetaData.Bin

// DeployContractOpenOracleTaskManager deploys a new Ethereum contract, binding an instance of ContractOpenOracleTaskManager to it.
func DeployContractOpenOracleTaskManager(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address, _taskResponseWindowBlock uint32) (common.Address, *types.Transaction, *ContractOpenOracleTaskManager, error) {
	parsed, err := ContractOpenOracleTaskManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractOpenOracleTaskManagerBin), backend, _registryCoordinator, _taskResponseWindowBlock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractOpenOracleTaskManager{ContractOpenOracleTaskManagerCaller: ContractOpenOracleTaskManagerCaller{contract: contract}, ContractOpenOracleTaskManagerTransactor: ContractOpenOracleTaskManagerTransactor{contract: contract}, ContractOpenOracleTaskManagerFilterer: ContractOpenOracleTaskManagerFilterer{contract: contract}}, nil
}

// ContractOpenOracleTaskManager is an auto generated Go binding around an Ethereum contract.
type ContractOpenOracleTaskManager struct {
	ContractOpenOracleTaskManagerCaller     // Read-only binding to the contract
	ContractOpenOracleTaskManagerTransactor // Write-only binding to the contract
	ContractOpenOracleTaskManagerFilterer   // Log filterer for contract events
}

// ContractOpenOracleTaskManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractOpenOracleTaskManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOpenOracleTaskManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractOpenOracleTaskManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOpenOracleTaskManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractOpenOracleTaskManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOpenOracleTaskManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractOpenOracleTaskManagerSession struct {
	Contract     *ContractOpenOracleTaskManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ContractOpenOracleTaskManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractOpenOracleTaskManagerCallerSession struct {
	Contract *ContractOpenOracleTaskManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// ContractOpenOracleTaskManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractOpenOracleTaskManagerTransactorSession struct {
	Contract     *ContractOpenOracleTaskManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// ContractOpenOracleTaskManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractOpenOracleTaskManagerRaw struct {
	Contract *ContractOpenOracleTaskManager // Generic contract binding to access the raw methods on
}

// ContractOpenOracleTaskManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractOpenOracleTaskManagerCallerRaw struct {
	Contract *ContractOpenOracleTaskManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractOpenOracleTaskManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractOpenOracleTaskManagerTransactorRaw struct {
	Contract *ContractOpenOracleTaskManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractOpenOracleTaskManager creates a new instance of ContractOpenOracleTaskManager, bound to a specific deployed contract.
func NewContractOpenOracleTaskManager(address common.Address, backend bind.ContractBackend) (*ContractOpenOracleTaskManager, error) {
	contract, err := bindContractOpenOracleTaskManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleTaskManager{ContractOpenOracleTaskManagerCaller: ContractOpenOracleTaskManagerCaller{contract: contract}, ContractOpenOracleTaskManagerTransactor: ContractOpenOracleTaskManagerTransactor{contract: contract}, ContractOpenOracleTaskManagerFilterer: ContractOpenOracleTaskManagerFilterer{contract: contract}}, nil
}

// NewContractOpenOracleTaskManagerCaller creates a new read-only instance of ContractOpenOracleTaskManager, bound to a specific deployed contract.
func NewContractOpenOracleTaskManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractOpenOracleTaskManagerCaller, error) {
	contract, err := bindContractOpenOracleTaskManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleTaskManagerCaller{contract: contract}, nil
}

// NewContractOpenOracleTaskManagerTransactor creates a new write-only instance of ContractOpenOracleTaskManager, bound to a specific deployed contract.
func NewContractOpenOracleTaskManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractOpenOracleTaskManagerTransactor, error) {
	contract, err := bindContractOpenOracleTaskManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleTaskManagerTransactor{contract: contract}, nil
}

// NewContractOpenOracleTaskManagerFilterer creates a new log filterer instance of ContractOpenOracleTaskManager, bound to a specific deployed contract.
func NewContractOpenOracleTaskManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractOpenOracleTaskManagerFilterer, error) {
	contract, err := bindContractOpenOracleTaskManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleTaskManagerFilterer{contract: contract}, nil
}

// bindContractOpenOracleTaskManager binds a generic wrapper to an already deployed contract.
func bindContractOpenOracleTaskManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractOpenOracleTaskManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractOpenOracleTaskManager.Contract.ContractOpenOracleTaskManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.ContractOpenOracleTaskManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.ContractOpenOracleTaskManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractOpenOracleTaskManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.contract.Transact(opts, method, params...)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCaller) TASKRESPONSEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractOpenOracleTaskManager.contract.Call(opts, &out, "TASK_RESPONSE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractOpenOracleTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractOpenOracleTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCallerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractOpenOracleTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractOpenOracleTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOpenOracleTaskManager.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) Aggregator() (common.Address, error) {
	return _ContractOpenOracleTaskManager.Contract.Aggregator(&_ContractOpenOracleTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCallerSession) Aggregator() (common.Address, error) {
	return _ContractOpenOracleTaskManager.Contract.Aggregator(&_ContractOpenOracleTaskManager.CallOpts)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCaller) AllTaskHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractOpenOracleTaskManager.contract.Call(opts, &out, "allTaskHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractOpenOracleTaskManager.Contract.AllTaskHashes(&_ContractOpenOracleTaskManager.CallOpts, arg0)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCallerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractOpenOracleTaskManager.Contract.AllTaskHashes(&_ContractOpenOracleTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCaller) AllTaskResponses(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractOpenOracleTaskManager.contract.Call(opts, &out, "allTaskResponses", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractOpenOracleTaskManager.Contract.AllTaskResponses(&_ContractOpenOracleTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCallerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractOpenOracleTaskManager.Contract.AllTaskResponses(&_ContractOpenOracleTaskManager.CallOpts, arg0)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOpenOracleTaskManager.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractOpenOracleTaskManager.Contract.BlsApkRegistry(&_ContractOpenOracleTaskManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractOpenOracleTaskManager.Contract.BlsApkRegistry(&_ContractOpenOracleTaskManager.CallOpts)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractOpenOracleTaskManager.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)

	if err != nil {
		return *new(IBLSSignatureCheckerQuorumStakeTotals), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSSignatureCheckerQuorumStakeTotals)).(*IBLSSignatureCheckerQuorumStakeTotals)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractOpenOracleTaskManager.Contract.CheckSignatures(&_ContractOpenOracleTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractOpenOracleTaskManager.Contract.CheckSignatures(&_ContractOpenOracleTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOpenOracleTaskManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) Delegation() (common.Address, error) {
	return _ContractOpenOracleTaskManager.Contract.Delegation(&_ContractOpenOracleTaskManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCallerSession) Delegation() (common.Address, error) {
	return _ContractOpenOracleTaskManager.Contract.Delegation(&_ContractOpenOracleTaskManager.CallOpts)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractOpenOracleTaskManager.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractOpenOracleTaskManager.Contract.GetCheckSignaturesIndices(&_ContractOpenOracleTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractOpenOracleTaskManager.Contract.GetCheckSignaturesIndices(&_ContractOpenOracleTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractOpenOracleTaskManager.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractOpenOracleTaskManager.Contract.GetOperatorState(&_ContractOpenOracleTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractOpenOracleTaskManager.Contract.GetOperatorState(&_ContractOpenOracleTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractOpenOracleTaskManager.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

	if err != nil {
		return *new(*big.Int), *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, out1, err

}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractOpenOracleTaskManager.Contract.GetOperatorState0(&_ContractOpenOracleTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractOpenOracleTaskManager.Contract.GetOperatorState0(&_ContractOpenOracleTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCaller) GetQuorumBitmapsAtBlockNumber(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractOpenOracleTaskManager.contract.Call(opts, &out, "getQuorumBitmapsAtBlockNumber", registryCoordinator, operatorIds, blockNumber)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractOpenOracleTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractOpenOracleTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCallerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractOpenOracleTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractOpenOracleTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCaller) GetTaskResponseWindowBlock(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractOpenOracleTaskManager.contract.Call(opts, &out, "getTaskResponseWindowBlock")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractOpenOracleTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractOpenOracleTaskManager.CallOpts)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCallerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractOpenOracleTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractOpenOracleTaskManager.CallOpts)
}

// IsFeed is a free data retrieval call binding the contract method 0xe58fdd04.
//
// Solidity: function isFeed(address ) view returns(bool)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCaller) IsFeed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ContractOpenOracleTaskManager.contract.Call(opts, &out, "isFeed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFeed is a free data retrieval call binding the contract method 0xe58fdd04.
//
// Solidity: function isFeed(address ) view returns(bool)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) IsFeed(arg0 common.Address) (bool, error) {
	return _ContractOpenOracleTaskManager.Contract.IsFeed(&_ContractOpenOracleTaskManager.CallOpts, arg0)
}

// IsFeed is a free data retrieval call binding the contract method 0xe58fdd04.
//
// Solidity: function isFeed(address ) view returns(bool)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCallerSession) IsFeed(arg0 common.Address) (bool, error) {
	return _ContractOpenOracleTaskManager.Contract.IsFeed(&_ContractOpenOracleTaskManager.CallOpts, arg0)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCaller) LatestTaskNum(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractOpenOracleTaskManager.contract.Call(opts, &out, "latestTaskNum")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) LatestTaskNum() (uint32, error) {
	return _ContractOpenOracleTaskManager.Contract.LatestTaskNum(&_ContractOpenOracleTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCallerSession) LatestTaskNum() (uint32, error) {
	return _ContractOpenOracleTaskManager.Contract.LatestTaskNum(&_ContractOpenOracleTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOpenOracleTaskManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) Owner() (common.Address, error) {
	return _ContractOpenOracleTaskManager.Contract.Owner(&_ContractOpenOracleTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCallerSession) Owner() (common.Address, error) {
	return _ContractOpenOracleTaskManager.Contract.Owner(&_ContractOpenOracleTaskManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractOpenOracleTaskManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) Paused(index uint8) (bool, error) {
	return _ContractOpenOracleTaskManager.Contract.Paused(&_ContractOpenOracleTaskManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractOpenOracleTaskManager.Contract.Paused(&_ContractOpenOracleTaskManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractOpenOracleTaskManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) Paused0() (*big.Int, error) {
	return _ContractOpenOracleTaskManager.Contract.Paused0(&_ContractOpenOracleTaskManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractOpenOracleTaskManager.Contract.Paused0(&_ContractOpenOracleTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOpenOracleTaskManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractOpenOracleTaskManager.Contract.PauserRegistry(&_ContractOpenOracleTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractOpenOracleTaskManager.Contract.PauserRegistry(&_ContractOpenOracleTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOpenOracleTaskManager.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractOpenOracleTaskManager.Contract.RegistryCoordinator(&_ContractOpenOracleTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractOpenOracleTaskManager.Contract.RegistryCoordinator(&_ContractOpenOracleTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOpenOracleTaskManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) StakeRegistry() (common.Address, error) {
	return _ContractOpenOracleTaskManager.Contract.StakeRegistry(&_ContractOpenOracleTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractOpenOracleTaskManager.Contract.StakeRegistry(&_ContractOpenOracleTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractOpenOracleTaskManager.contract.Call(opts, &out, "staleStakesForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) StaleStakesForbidden() (bool, error) {
	return _ContractOpenOracleTaskManager.Contract.StaleStakesForbidden(&_ContractOpenOracleTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCallerSession) StaleStakesForbidden() (bool, error) {
	return _ContractOpenOracleTaskManager.Contract.StaleStakesForbidden(&_ContractOpenOracleTaskManager.CallOpts)
}

// TaskCreationFee is a free data retrieval call binding the contract method 0x4d075ce7.
//
// Solidity: function taskCreationFee() view returns(uint256)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCaller) TaskCreationFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractOpenOracleTaskManager.contract.Call(opts, &out, "taskCreationFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaskCreationFee is a free data retrieval call binding the contract method 0x4d075ce7.
//
// Solidity: function taskCreationFee() view returns(uint256)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) TaskCreationFee() (*big.Int, error) {
	return _ContractOpenOracleTaskManager.Contract.TaskCreationFee(&_ContractOpenOracleTaskManager.CallOpts)
}

// TaskCreationFee is a free data retrieval call binding the contract method 0x4d075ce7.
//
// Solidity: function taskCreationFee() view returns(uint256)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCallerSession) TaskCreationFee() (*big.Int, error) {
	return _ContractOpenOracleTaskManager.Contract.TaskCreationFee(&_ContractOpenOracleTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCaller) TaskNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractOpenOracleTaskManager.contract.Call(opts, &out, "taskNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) TaskNumber() (uint32, error) {
	return _ContractOpenOracleTaskManager.Contract.TaskNumber(&_ContractOpenOracleTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCallerSession) TaskNumber() (uint32, error) {
	return _ContractOpenOracleTaskManager.Contract.TaskNumber(&_ContractOpenOracleTaskManager.CallOpts)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _ContractOpenOracleTaskManager.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

	outstruct := new(struct {
		PairingSuccessful bool
		SiganatureIsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PairingSuccessful = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SiganatureIsValid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractOpenOracleTaskManager.Contract.TrySignatureAndApkVerification(&_ContractOpenOracleTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractOpenOracleTaskManager.Contract.TrySignatureAndApkVerification(&_ContractOpenOracleTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// AddToFeedlist is a paid mutator transaction binding the contract method 0x033ff004.
//
// Solidity: function addToFeedlist(address _address) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactor) AddToFeedlist(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.contract.Transact(opts, "addToFeedlist", _address)
}

// AddToFeedlist is a paid mutator transaction binding the contract method 0x033ff004.
//
// Solidity: function addToFeedlist(address _address) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) AddToFeedlist(_address common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.AddToFeedlist(&_ContractOpenOracleTaskManager.TransactOpts, _address)
}

// AddToFeedlist is a paid mutator transaction binding the contract method 0x033ff004.
//
// Solidity: function addToFeedlist(address _address) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactorSession) AddToFeedlist(_address common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.AddToFeedlist(&_ContractOpenOracleTaskManager.TransactOpts, _address)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x2191297d.
//
// Solidity: function createNewTask(uint8 taskType, uint8 responderThreshold, uint96 stakeThreshold) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactor) CreateNewTask(opts *bind.TransactOpts, taskType uint8, responderThreshold uint8, stakeThreshold *big.Int) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.contract.Transact(opts, "createNewTask", taskType, responderThreshold, stakeThreshold)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x2191297d.
//
// Solidity: function createNewTask(uint8 taskType, uint8 responderThreshold, uint96 stakeThreshold) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) CreateNewTask(taskType uint8, responderThreshold uint8, stakeThreshold *big.Int) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.CreateNewTask(&_ContractOpenOracleTaskManager.TransactOpts, taskType, responderThreshold, stakeThreshold)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x2191297d.
//
// Solidity: function createNewTask(uint8 taskType, uint8 responderThreshold, uint96 stakeThreshold) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactorSession) CreateNewTask(taskType uint8, responderThreshold uint8, stakeThreshold *big.Int) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.CreateNewTask(&_ContractOpenOracleTaskManager.TransactOpts, taskType, responderThreshold, stakeThreshold)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.contract.Transact(opts, "initialize", _pauserRegistry, initialOwner, _aggregator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.Initialize(&_ContractOpenOracleTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactorSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.Initialize(&_ContractOpenOracleTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.Pause(&_ContractOpenOracleTaskManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.Pause(&_ContractOpenOracleTaskManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.PauseAll(&_ContractOpenOracleTaskManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.PauseAll(&_ContractOpenOracleTaskManager.TransactOpts)
}

// RemoveFromFeed is a paid mutator transaction binding the contract method 0x0c0f5792.
//
// Solidity: function removeFromFeed(address _address) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactor) RemoveFromFeed(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.contract.Transact(opts, "removeFromFeed", _address)
}

// RemoveFromFeed is a paid mutator transaction binding the contract method 0x0c0f5792.
//
// Solidity: function removeFromFeed(address _address) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) RemoveFromFeed(_address common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.RemoveFromFeed(&_ContractOpenOracleTaskManager.TransactOpts, _address)
}

// RemoveFromFeed is a paid mutator transaction binding the contract method 0x0c0f5792.
//
// Solidity: function removeFromFeed(address _address) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactorSession) RemoveFromFeed(_address common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.RemoveFromFeed(&_ContractOpenOracleTaskManager.TransactOpts, _address)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.RenounceOwnership(&_ContractOpenOracleTaskManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.RenounceOwnership(&_ContractOpenOracleTaskManager.TransactOpts)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x10364211.
//
// Solidity: function respondToTask((uint8,uint32,uint8,uint96,address,uint256) task, (address,(uint32,uint256,uint256),bytes)[] responses, (uint32,uint256,uint256,uint256) weightedTaskResponse) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactor) RespondToTask(opts *bind.TransactOpts, task IOpenOracleTaskManagerTask, responses []IOpenOracleTaskManagerOperatorResponse, weightedTaskResponse IOpenOracleTaskManagerWeightedTaskResponse) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.contract.Transact(opts, "respondToTask", task, responses, weightedTaskResponse)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x10364211.
//
// Solidity: function respondToTask((uint8,uint32,uint8,uint96,address,uint256) task, (address,(uint32,uint256,uint256),bytes)[] responses, (uint32,uint256,uint256,uint256) weightedTaskResponse) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) RespondToTask(task IOpenOracleTaskManagerTask, responses []IOpenOracleTaskManagerOperatorResponse, weightedTaskResponse IOpenOracleTaskManagerWeightedTaskResponse) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.RespondToTask(&_ContractOpenOracleTaskManager.TransactOpts, task, responses, weightedTaskResponse)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x10364211.
//
// Solidity: function respondToTask((uint8,uint32,uint8,uint96,address,uint256) task, (address,(uint32,uint256,uint256),bytes)[] responses, (uint32,uint256,uint256,uint256) weightedTaskResponse) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactorSession) RespondToTask(task IOpenOracleTaskManagerTask, responses []IOpenOracleTaskManagerOperatorResponse, weightedTaskResponse IOpenOracleTaskManagerWeightedTaskResponse) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.RespondToTask(&_ContractOpenOracleTaskManager.TransactOpts, task, responses, weightedTaskResponse)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.SetPauserRegistry(&_ContractOpenOracleTaskManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.SetPauserRegistry(&_ContractOpenOracleTaskManager.TransactOpts, newPauserRegistry)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.SetStaleStakesForbidden(&_ContractOpenOracleTaskManager.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.SetStaleStakesForbidden(&_ContractOpenOracleTaskManager.TransactOpts, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.TransferOwnership(&_ContractOpenOracleTaskManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.TransferOwnership(&_ContractOpenOracleTaskManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.Unpause(&_ContractOpenOracleTaskManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.Unpause(&_ContractOpenOracleTaskManager.TransactOpts, newPausedStatus)
}

// UpdateAggregator is a paid mutator transaction binding the contract method 0x9fe4ee47.
//
// Solidity: function updateAggregator(address _aggregator) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactor) UpdateAggregator(opts *bind.TransactOpts, _aggregator common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.contract.Transact(opts, "updateAggregator", _aggregator)
}

// UpdateAggregator is a paid mutator transaction binding the contract method 0x9fe4ee47.
//
// Solidity: function updateAggregator(address _aggregator) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) UpdateAggregator(_aggregator common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.UpdateAggregator(&_ContractOpenOracleTaskManager.TransactOpts, _aggregator)
}

// UpdateAggregator is a paid mutator transaction binding the contract method 0x9fe4ee47.
//
// Solidity: function updateAggregator(address _aggregator) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactorSession) UpdateAggregator(_aggregator common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.UpdateAggregator(&_ContractOpenOracleTaskManager.TransactOpts, _aggregator)
}

// UpdateTaskCreationFee is a paid mutator transaction binding the contract method 0x52e9408b.
//
// Solidity: function updateTaskCreationFee(uint256 _newFee) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactor) UpdateTaskCreationFee(opts *bind.TransactOpts, _newFee *big.Int) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.contract.Transact(opts, "updateTaskCreationFee", _newFee)
}

// UpdateTaskCreationFee is a paid mutator transaction binding the contract method 0x52e9408b.
//
// Solidity: function updateTaskCreationFee(uint256 _newFee) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) UpdateTaskCreationFee(_newFee *big.Int) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.UpdateTaskCreationFee(&_ContractOpenOracleTaskManager.TransactOpts, _newFee)
}

// UpdateTaskCreationFee is a paid mutator transaction binding the contract method 0x52e9408b.
//
// Solidity: function updateTaskCreationFee(uint256 _newFee) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactorSession) UpdateTaskCreationFee(_newFee *big.Int) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.UpdateTaskCreationFee(&_ContractOpenOracleTaskManager.TransactOpts, _newFee)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) Withdraw() (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.Withdraw(&_ContractOpenOracleTaskManager.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactorSession) Withdraw() (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.Withdraw(&_ContractOpenOracleTaskManager.TransactOpts)
}

// WithdrawTaskFunds is a paid mutator transaction binding the contract method 0x459b58b1.
//
// Solidity: function withdrawTaskFunds(uint32 taskNum, (uint8,uint32,uint8,uint96,address,uint256) task) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactor) WithdrawTaskFunds(opts *bind.TransactOpts, taskNum uint32, task IOpenOracleTaskManagerTask) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.contract.Transact(opts, "withdrawTaskFunds", taskNum, task)
}

// WithdrawTaskFunds is a paid mutator transaction binding the contract method 0x459b58b1.
//
// Solidity: function withdrawTaskFunds(uint32 taskNum, (uint8,uint32,uint8,uint96,address,uint256) task) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) WithdrawTaskFunds(taskNum uint32, task IOpenOracleTaskManagerTask) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.WithdrawTaskFunds(&_ContractOpenOracleTaskManager.TransactOpts, taskNum, task)
}

// WithdrawTaskFunds is a paid mutator transaction binding the contract method 0x459b58b1.
//
// Solidity: function withdrawTaskFunds(uint32 taskNum, (uint8,uint32,uint8,uint96,address,uint256) task) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactorSession) WithdrawTaskFunds(taskNum uint32, task IOpenOracleTaskManagerTask) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.WithdrawTaskFunds(&_ContractOpenOracleTaskManager.TransactOpts, taskNum, task)
}

// ContractOpenOracleTaskManagerAddressAddedToFeedlistIterator is returned from FilterAddressAddedToFeedlist and is used to iterate over the raw logs and unpacked data for AddressAddedToFeedlist events raised by the ContractOpenOracleTaskManager contract.
type ContractOpenOracleTaskManagerAddressAddedToFeedlistIterator struct {
	Event *ContractOpenOracleTaskManagerAddressAddedToFeedlist // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOpenOracleTaskManagerAddressAddedToFeedlistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenOracleTaskManagerAddressAddedToFeedlist)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOpenOracleTaskManagerAddressAddedToFeedlist)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOpenOracleTaskManagerAddressAddedToFeedlistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenOracleTaskManagerAddressAddedToFeedlistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenOracleTaskManagerAddressAddedToFeedlist represents a AddressAddedToFeedlist event raised by the ContractOpenOracleTaskManager contract.
type ContractOpenOracleTaskManagerAddressAddedToFeedlist struct {
	Address common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddressAddedToFeedlist is a free log retrieval operation binding the contract event 0xe0fb1d5933f4eca7f5928d3c3c41d075398c1d1a4c4b09a5771a1c5c67eb1d86.
//
// Solidity: event AddressAddedToFeedlist(address indexed _address)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) FilterAddressAddedToFeedlist(opts *bind.FilterOpts, _address []common.Address) (*ContractOpenOracleTaskManagerAddressAddedToFeedlistIterator, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _ContractOpenOracleTaskManager.contract.FilterLogs(opts, "AddressAddedToFeedlist", _addressRule)
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleTaskManagerAddressAddedToFeedlistIterator{contract: _ContractOpenOracleTaskManager.contract, event: "AddressAddedToFeedlist", logs: logs, sub: sub}, nil
}

// WatchAddressAddedToFeedlist is a free log subscription operation binding the contract event 0xe0fb1d5933f4eca7f5928d3c3c41d075398c1d1a4c4b09a5771a1c5c67eb1d86.
//
// Solidity: event AddressAddedToFeedlist(address indexed _address)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) WatchAddressAddedToFeedlist(opts *bind.WatchOpts, sink chan<- *ContractOpenOracleTaskManagerAddressAddedToFeedlist, _address []common.Address) (event.Subscription, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _ContractOpenOracleTaskManager.contract.WatchLogs(opts, "AddressAddedToFeedlist", _addressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenOracleTaskManagerAddressAddedToFeedlist)
				if err := _ContractOpenOracleTaskManager.contract.UnpackLog(event, "AddressAddedToFeedlist", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddressAddedToFeedlist is a log parse operation binding the contract event 0xe0fb1d5933f4eca7f5928d3c3c41d075398c1d1a4c4b09a5771a1c5c67eb1d86.
//
// Solidity: event AddressAddedToFeedlist(address indexed _address)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) ParseAddressAddedToFeedlist(log types.Log) (*ContractOpenOracleTaskManagerAddressAddedToFeedlist, error) {
	event := new(ContractOpenOracleTaskManagerAddressAddedToFeedlist)
	if err := _ContractOpenOracleTaskManager.contract.UnpackLog(event, "AddressAddedToFeedlist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOpenOracleTaskManagerAddressRemovedFromFeedlistIterator is returned from FilterAddressRemovedFromFeedlist and is used to iterate over the raw logs and unpacked data for AddressRemovedFromFeedlist events raised by the ContractOpenOracleTaskManager contract.
type ContractOpenOracleTaskManagerAddressRemovedFromFeedlistIterator struct {
	Event *ContractOpenOracleTaskManagerAddressRemovedFromFeedlist // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOpenOracleTaskManagerAddressRemovedFromFeedlistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenOracleTaskManagerAddressRemovedFromFeedlist)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOpenOracleTaskManagerAddressRemovedFromFeedlist)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOpenOracleTaskManagerAddressRemovedFromFeedlistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenOracleTaskManagerAddressRemovedFromFeedlistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenOracleTaskManagerAddressRemovedFromFeedlist represents a AddressRemovedFromFeedlist event raised by the ContractOpenOracleTaskManager contract.
type ContractOpenOracleTaskManagerAddressRemovedFromFeedlist struct {
	Address common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddressRemovedFromFeedlist is a free log retrieval operation binding the contract event 0x28d57cd2f66bc88bf9eeda7d803691bff6e77bfdbf4212b41841a0d1b7189c5f.
//
// Solidity: event AddressRemovedFromFeedlist(address indexed _address)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) FilterAddressRemovedFromFeedlist(opts *bind.FilterOpts, _address []common.Address) (*ContractOpenOracleTaskManagerAddressRemovedFromFeedlistIterator, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _ContractOpenOracleTaskManager.contract.FilterLogs(opts, "AddressRemovedFromFeedlist", _addressRule)
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleTaskManagerAddressRemovedFromFeedlistIterator{contract: _ContractOpenOracleTaskManager.contract, event: "AddressRemovedFromFeedlist", logs: logs, sub: sub}, nil
}

// WatchAddressRemovedFromFeedlist is a free log subscription operation binding the contract event 0x28d57cd2f66bc88bf9eeda7d803691bff6e77bfdbf4212b41841a0d1b7189c5f.
//
// Solidity: event AddressRemovedFromFeedlist(address indexed _address)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) WatchAddressRemovedFromFeedlist(opts *bind.WatchOpts, sink chan<- *ContractOpenOracleTaskManagerAddressRemovedFromFeedlist, _address []common.Address) (event.Subscription, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _ContractOpenOracleTaskManager.contract.WatchLogs(opts, "AddressRemovedFromFeedlist", _addressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenOracleTaskManagerAddressRemovedFromFeedlist)
				if err := _ContractOpenOracleTaskManager.contract.UnpackLog(event, "AddressRemovedFromFeedlist", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddressRemovedFromFeedlist is a log parse operation binding the contract event 0x28d57cd2f66bc88bf9eeda7d803691bff6e77bfdbf4212b41841a0d1b7189c5f.
//
// Solidity: event AddressRemovedFromFeedlist(address indexed _address)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) ParseAddressRemovedFromFeedlist(log types.Log) (*ContractOpenOracleTaskManagerAddressRemovedFromFeedlist, error) {
	event := new(ContractOpenOracleTaskManagerAddressRemovedFromFeedlist)
	if err := _ContractOpenOracleTaskManager.contract.UnpackLog(event, "AddressRemovedFromFeedlist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOpenOracleTaskManagerAggregatorUpdatedIterator is returned from FilterAggregatorUpdated and is used to iterate over the raw logs and unpacked data for AggregatorUpdated events raised by the ContractOpenOracleTaskManager contract.
type ContractOpenOracleTaskManagerAggregatorUpdatedIterator struct {
	Event *ContractOpenOracleTaskManagerAggregatorUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOpenOracleTaskManagerAggregatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenOracleTaskManagerAggregatorUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOpenOracleTaskManagerAggregatorUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOpenOracleTaskManagerAggregatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenOracleTaskManagerAggregatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenOracleTaskManagerAggregatorUpdated represents a AggregatorUpdated event raised by the ContractOpenOracleTaskManager contract.
type ContractOpenOracleTaskManagerAggregatorUpdated struct {
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAggregatorUpdated is a free log retrieval operation binding the contract event 0x602cec4b1583b07d071161da5eb9589444d2459201e2fab7753dc941e9351c21.
//
// Solidity: event AggregatorUpdated(address aggregator)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) FilterAggregatorUpdated(opts *bind.FilterOpts) (*ContractOpenOracleTaskManagerAggregatorUpdatedIterator, error) {

	logs, sub, err := _ContractOpenOracleTaskManager.contract.FilterLogs(opts, "AggregatorUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleTaskManagerAggregatorUpdatedIterator{contract: _ContractOpenOracleTaskManager.contract, event: "AggregatorUpdated", logs: logs, sub: sub}, nil
}

// WatchAggregatorUpdated is a free log subscription operation binding the contract event 0x602cec4b1583b07d071161da5eb9589444d2459201e2fab7753dc941e9351c21.
//
// Solidity: event AggregatorUpdated(address aggregator)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) WatchAggregatorUpdated(opts *bind.WatchOpts, sink chan<- *ContractOpenOracleTaskManagerAggregatorUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractOpenOracleTaskManager.contract.WatchLogs(opts, "AggregatorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenOracleTaskManagerAggregatorUpdated)
				if err := _ContractOpenOracleTaskManager.contract.UnpackLog(event, "AggregatorUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAggregatorUpdated is a log parse operation binding the contract event 0x602cec4b1583b07d071161da5eb9589444d2459201e2fab7753dc941e9351c21.
//
// Solidity: event AggregatorUpdated(address aggregator)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) ParseAggregatorUpdated(log types.Log) (*ContractOpenOracleTaskManagerAggregatorUpdated, error) {
	event := new(ContractOpenOracleTaskManagerAggregatorUpdated)
	if err := _ContractOpenOracleTaskManager.contract.UnpackLog(event, "AggregatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOpenOracleTaskManagerFundsWithdrawnIterator is returned from FilterFundsWithdrawn and is used to iterate over the raw logs and unpacked data for FundsWithdrawn events raised by the ContractOpenOracleTaskManager contract.
type ContractOpenOracleTaskManagerFundsWithdrawnIterator struct {
	Event *ContractOpenOracleTaskManagerFundsWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOpenOracleTaskManagerFundsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenOracleTaskManagerFundsWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOpenOracleTaskManagerFundsWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOpenOracleTaskManagerFundsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenOracleTaskManagerFundsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenOracleTaskManagerFundsWithdrawn represents a FundsWithdrawn event raised by the ContractOpenOracleTaskManager contract.
type ContractOpenOracleTaskManagerFundsWithdrawn struct {
	TaskNum *big.Int
	Creator common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFundsWithdrawn is a free log retrieval operation binding the contract event 0xf440aec6b52895984d061d622e6edeba6210f7c3e059be920663140c084560d7.
//
// Solidity: event FundsWithdrawn(uint256 taskNum, address creator, uint256 amount)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) FilterFundsWithdrawn(opts *bind.FilterOpts) (*ContractOpenOracleTaskManagerFundsWithdrawnIterator, error) {

	logs, sub, err := _ContractOpenOracleTaskManager.contract.FilterLogs(opts, "FundsWithdrawn")
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleTaskManagerFundsWithdrawnIterator{contract: _ContractOpenOracleTaskManager.contract, event: "FundsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchFundsWithdrawn is a free log subscription operation binding the contract event 0xf440aec6b52895984d061d622e6edeba6210f7c3e059be920663140c084560d7.
//
// Solidity: event FundsWithdrawn(uint256 taskNum, address creator, uint256 amount)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) WatchFundsWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractOpenOracleTaskManagerFundsWithdrawn) (event.Subscription, error) {

	logs, sub, err := _ContractOpenOracleTaskManager.contract.WatchLogs(opts, "FundsWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenOracleTaskManagerFundsWithdrawn)
				if err := _ContractOpenOracleTaskManager.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFundsWithdrawn is a log parse operation binding the contract event 0xf440aec6b52895984d061d622e6edeba6210f7c3e059be920663140c084560d7.
//
// Solidity: event FundsWithdrawn(uint256 taskNum, address creator, uint256 amount)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) ParseFundsWithdrawn(log types.Log) (*ContractOpenOracleTaskManagerFundsWithdrawn, error) {
	event := new(ContractOpenOracleTaskManagerFundsWithdrawn)
	if err := _ContractOpenOracleTaskManager.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOpenOracleTaskManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractOpenOracleTaskManager contract.
type ContractOpenOracleTaskManagerInitializedIterator struct {
	Event *ContractOpenOracleTaskManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOpenOracleTaskManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenOracleTaskManagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOpenOracleTaskManagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOpenOracleTaskManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenOracleTaskManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenOracleTaskManagerInitialized represents a Initialized event raised by the ContractOpenOracleTaskManager contract.
type ContractOpenOracleTaskManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractOpenOracleTaskManagerInitializedIterator, error) {

	logs, sub, err := _ContractOpenOracleTaskManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleTaskManagerInitializedIterator{contract: _ContractOpenOracleTaskManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractOpenOracleTaskManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractOpenOracleTaskManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenOracleTaskManagerInitialized)
				if err := _ContractOpenOracleTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) ParseInitialized(log types.Log) (*ContractOpenOracleTaskManagerInitialized, error) {
	event := new(ContractOpenOracleTaskManagerInitialized)
	if err := _ContractOpenOracleTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOpenOracleTaskManagerNewTaskCreatedIterator is returned from FilterNewTaskCreated and is used to iterate over the raw logs and unpacked data for NewTaskCreated events raised by the ContractOpenOracleTaskManager contract.
type ContractOpenOracleTaskManagerNewTaskCreatedIterator struct {
	Event *ContractOpenOracleTaskManagerNewTaskCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOpenOracleTaskManagerNewTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenOracleTaskManagerNewTaskCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOpenOracleTaskManagerNewTaskCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOpenOracleTaskManagerNewTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenOracleTaskManagerNewTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenOracleTaskManagerNewTaskCreated represents a NewTaskCreated event raised by the ContractOpenOracleTaskManager contract.
type ContractOpenOracleTaskManagerNewTaskCreated struct {
	TaskIndex uint32
	Task      IOpenOracleTaskManagerTask
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewTaskCreated is a free log retrieval operation binding the contract event 0xdeb7c64f4d939c7b58f45fb9cd5b5118e0b5920d43390c9ba0d47032754c3121.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint8,uint32,uint8,uint96,address,uint256) task)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) FilterNewTaskCreated(opts *bind.FilterOpts, taskIndex []uint32) (*ContractOpenOracleTaskManagerNewTaskCreatedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractOpenOracleTaskManager.contract.FilterLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleTaskManagerNewTaskCreatedIterator{contract: _ContractOpenOracleTaskManager.contract, event: "NewTaskCreated", logs: logs, sub: sub}, nil
}

// WatchNewTaskCreated is a free log subscription operation binding the contract event 0xdeb7c64f4d939c7b58f45fb9cd5b5118e0b5920d43390c9ba0d47032754c3121.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint8,uint32,uint8,uint96,address,uint256) task)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) WatchNewTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractOpenOracleTaskManagerNewTaskCreated, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractOpenOracleTaskManager.contract.WatchLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenOracleTaskManagerNewTaskCreated)
				if err := _ContractOpenOracleTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewTaskCreated is a log parse operation binding the contract event 0xdeb7c64f4d939c7b58f45fb9cd5b5118e0b5920d43390c9ba0d47032754c3121.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint8,uint32,uint8,uint96,address,uint256) task)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) ParseNewTaskCreated(log types.Log) (*ContractOpenOracleTaskManagerNewTaskCreated, error) {
	event := new(ContractOpenOracleTaskManagerNewTaskCreated)
	if err := _ContractOpenOracleTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOpenOracleTaskManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractOpenOracleTaskManager contract.
type ContractOpenOracleTaskManagerOwnershipTransferredIterator struct {
	Event *ContractOpenOracleTaskManagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOpenOracleTaskManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenOracleTaskManagerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOpenOracleTaskManagerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOpenOracleTaskManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenOracleTaskManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenOracleTaskManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractOpenOracleTaskManager contract.
type ContractOpenOracleTaskManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOpenOracleTaskManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractOpenOracleTaskManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleTaskManagerOwnershipTransferredIterator{contract: _ContractOpenOracleTaskManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOpenOracleTaskManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractOpenOracleTaskManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenOracleTaskManagerOwnershipTransferred)
				if err := _ContractOpenOracleTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOpenOracleTaskManagerOwnershipTransferred, error) {
	event := new(ContractOpenOracleTaskManagerOwnershipTransferred)
	if err := _ContractOpenOracleTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOpenOracleTaskManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractOpenOracleTaskManager contract.
type ContractOpenOracleTaskManagerPausedIterator struct {
	Event *ContractOpenOracleTaskManagerPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOpenOracleTaskManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenOracleTaskManagerPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOpenOracleTaskManagerPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOpenOracleTaskManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenOracleTaskManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenOracleTaskManagerPaused represents a Paused event raised by the ContractOpenOracleTaskManager contract.
type ContractOpenOracleTaskManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractOpenOracleTaskManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractOpenOracleTaskManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleTaskManagerPausedIterator{contract: _ContractOpenOracleTaskManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractOpenOracleTaskManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractOpenOracleTaskManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenOracleTaskManagerPaused)
				if err := _ContractOpenOracleTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) ParsePaused(log types.Log) (*ContractOpenOracleTaskManagerPaused, error) {
	event := new(ContractOpenOracleTaskManagerPaused)
	if err := _ContractOpenOracleTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOpenOracleTaskManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractOpenOracleTaskManager contract.
type ContractOpenOracleTaskManagerPauserRegistrySetIterator struct {
	Event *ContractOpenOracleTaskManagerPauserRegistrySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOpenOracleTaskManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenOracleTaskManagerPauserRegistrySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOpenOracleTaskManagerPauserRegistrySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOpenOracleTaskManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenOracleTaskManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenOracleTaskManagerPauserRegistrySet represents a PauserRegistrySet event raised by the ContractOpenOracleTaskManager contract.
type ContractOpenOracleTaskManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractOpenOracleTaskManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractOpenOracleTaskManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleTaskManagerPauserRegistrySetIterator{contract: _ContractOpenOracleTaskManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractOpenOracleTaskManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractOpenOracleTaskManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenOracleTaskManagerPauserRegistrySet)
				if err := _ContractOpenOracleTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) ParsePauserRegistrySet(log types.Log) (*ContractOpenOracleTaskManagerPauserRegistrySet, error) {
	event := new(ContractOpenOracleTaskManagerPauserRegistrySet)
	if err := _ContractOpenOracleTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOpenOracleTaskManagerStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the ContractOpenOracleTaskManager contract.
type ContractOpenOracleTaskManagerStaleStakesForbiddenUpdateIterator struct {
	Event *ContractOpenOracleTaskManagerStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOpenOracleTaskManagerStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenOracleTaskManagerStaleStakesForbiddenUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOpenOracleTaskManagerStaleStakesForbiddenUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOpenOracleTaskManagerStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenOracleTaskManagerStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenOracleTaskManagerStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the ContractOpenOracleTaskManager contract.
type ContractOpenOracleTaskManagerStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractOpenOracleTaskManagerStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _ContractOpenOracleTaskManager.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleTaskManagerStaleStakesForbiddenUpdateIterator{contract: _ContractOpenOracleTaskManager.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractOpenOracleTaskManagerStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _ContractOpenOracleTaskManager.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenOracleTaskManagerStaleStakesForbiddenUpdate)
				if err := _ContractOpenOracleTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStaleStakesForbiddenUpdate is a log parse operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractOpenOracleTaskManagerStaleStakesForbiddenUpdate, error) {
	event := new(ContractOpenOracleTaskManagerStaleStakesForbiddenUpdate)
	if err := _ContractOpenOracleTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOpenOracleTaskManagerTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the ContractOpenOracleTaskManager contract.
type ContractOpenOracleTaskManagerTaskCompletedIterator struct {
	Event *ContractOpenOracleTaskManagerTaskCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOpenOracleTaskManagerTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenOracleTaskManagerTaskCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOpenOracleTaskManagerTaskCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOpenOracleTaskManagerTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenOracleTaskManagerTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenOracleTaskManagerTaskCompleted represents a TaskCompleted event raised by the ContractOpenOracleTaskManager contract.
type ContractOpenOracleTaskManagerTaskCompleted struct {
	TaskIndex uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskIndex []uint32) (*ContractOpenOracleTaskManagerTaskCompletedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractOpenOracleTaskManager.contract.FilterLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleTaskManagerTaskCompletedIterator{contract: _ContractOpenOracleTaskManager.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *ContractOpenOracleTaskManagerTaskCompleted, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractOpenOracleTaskManager.contract.WatchLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenOracleTaskManagerTaskCompleted)
				if err := _ContractOpenOracleTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskCompleted is a log parse operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) ParseTaskCompleted(log types.Log) (*ContractOpenOracleTaskManagerTaskCompleted, error) {
	event := new(ContractOpenOracleTaskManagerTaskCompleted)
	if err := _ContractOpenOracleTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOpenOracleTaskManagerTaskRespondedIterator is returned from FilterTaskResponded and is used to iterate over the raw logs and unpacked data for TaskResponded events raised by the ContractOpenOracleTaskManager contract.
type ContractOpenOracleTaskManagerTaskRespondedIterator struct {
	Event *ContractOpenOracleTaskManagerTaskResponded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOpenOracleTaskManagerTaskRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenOracleTaskManagerTaskResponded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOpenOracleTaskManagerTaskResponded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOpenOracleTaskManagerTaskRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenOracleTaskManagerTaskRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenOracleTaskManagerTaskResponded represents a TaskResponded event raised by the ContractOpenOracleTaskManager contract.
type ContractOpenOracleTaskManagerTaskResponded struct {
	Task                 IOpenOracleTaskManagerTask
	TaskResponse         IOpenOracleTaskManagerWeightedTaskResponse
	TaskResponseMetadata IOpenOracleTaskManagerTaskResponseMetadata
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTaskResponded is a free log retrieval operation binding the contract event 0x5d0b79c46a19832eb8bcd7612cf0231afd4412288c9010e3fae97a0d3375fa6f.
//
// Solidity: event TaskResponded((uint8,uint32,uint8,uint96,address,uint256) task, (uint32,uint256,uint256,uint256) taskResponse, (uint32) taskResponseMetadata)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) FilterTaskResponded(opts *bind.FilterOpts) (*ContractOpenOracleTaskManagerTaskRespondedIterator, error) {

	logs, sub, err := _ContractOpenOracleTaskManager.contract.FilterLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleTaskManagerTaskRespondedIterator{contract: _ContractOpenOracleTaskManager.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0x5d0b79c46a19832eb8bcd7612cf0231afd4412288c9010e3fae97a0d3375fa6f.
//
// Solidity: event TaskResponded((uint8,uint32,uint8,uint96,address,uint256) task, (uint32,uint256,uint256,uint256) taskResponse, (uint32) taskResponseMetadata)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) WatchTaskResponded(opts *bind.WatchOpts, sink chan<- *ContractOpenOracleTaskManagerTaskResponded) (event.Subscription, error) {

	logs, sub, err := _ContractOpenOracleTaskManager.contract.WatchLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenOracleTaskManagerTaskResponded)
				if err := _ContractOpenOracleTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskResponded is a log parse operation binding the contract event 0x5d0b79c46a19832eb8bcd7612cf0231afd4412288c9010e3fae97a0d3375fa6f.
//
// Solidity: event TaskResponded((uint8,uint32,uint8,uint96,address,uint256) task, (uint32,uint256,uint256,uint256) taskResponse, (uint32) taskResponseMetadata)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) ParseTaskResponded(log types.Log) (*ContractOpenOracleTaskManagerTaskResponded, error) {
	event := new(ContractOpenOracleTaskManagerTaskResponded)
	if err := _ContractOpenOracleTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOpenOracleTaskManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractOpenOracleTaskManager contract.
type ContractOpenOracleTaskManagerUnpausedIterator struct {
	Event *ContractOpenOracleTaskManagerUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOpenOracleTaskManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenOracleTaskManagerUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOpenOracleTaskManagerUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOpenOracleTaskManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenOracleTaskManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenOracleTaskManagerUnpaused represents a Unpaused event raised by the ContractOpenOracleTaskManager contract.
type ContractOpenOracleTaskManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractOpenOracleTaskManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractOpenOracleTaskManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleTaskManagerUnpausedIterator{contract: _ContractOpenOracleTaskManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractOpenOracleTaskManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractOpenOracleTaskManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenOracleTaskManagerUnpaused)
				if err := _ContractOpenOracleTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) ParseUnpaused(log types.Log) (*ContractOpenOracleTaskManagerUnpaused, error) {
	event := new(ContractOpenOracleTaskManagerUnpaused)
	if err := _ContractOpenOracleTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

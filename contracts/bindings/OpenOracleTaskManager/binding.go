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
	TaskType         uint8
	TaskCreatedBlock uint32
	ResponderNumber  uint8
	StakeThreshold   *big.Int
	Creator          common.Address
	CreationFee      *big.Int
}

// IOpenOracleTaskManagerTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type IOpenOracleTaskManagerTaskResponse struct {
	ReferenceTaskIndex uint32
	Result             *big.Int
	TimeStamp          *big.Int
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
	TimeStamp          *big.Int
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"taskType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"responderNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"stakeThreshold\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestTaskNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIOpenOracleTaskManager.Task\",\"components\":[{\"name\":\"taskType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"responderNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"stakeThreshold\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"creationFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"responses\",\"type\":\"tuple[]\",\"internalType\":\"structIOpenOracleTaskManager.OperatorResponse[]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"response\",\"type\":\"tuple\",\"internalType\":\"structIOpenOracleTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeStamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"weightedTaskResponse\",\"type\":\"tuple\",\"internalType\":\"structIOpenOracleTaskManager.WeightedTaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sd\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeStamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskCreationFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTaskCreationFee\",\"inputs\":[{\"name\":\"_newFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawTaskFunds\",\"inputs\":[{\"name\":\"taskNum\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIOpenOracleTaskManager.Task\",\"components\":[{\"name\":\"taskType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"responderNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"stakeThreshold\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"creationFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FundsWithdrawn\",\"inputs\":[{\"name\":\"taskNum\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"creator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOpenOracleTaskManager.Task\",\"components\":[{\"name\":\"taskType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"responderNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"stakeThreshold\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"creationFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOpenOracleTaskManager.WeightedTaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sd\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeStamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOpenOracleTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x610120604052655af3107a400060c9553480156200001c57600080fd5b5060405162006243380380620062438339810160408190526200003f9162000201565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000099573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000bf919062000248565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000117573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013d919062000248565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000197573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001bd919062000248565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff1661010052506200026f565b6001600160a01b0381168114620001fe57600080fd5b50565b600080604083850312156200021557600080fd5b82516200022281620001e8565b602084015190925063ffffffff811681146200023d57600080fd5b809150509250929050565b6000602082840312156200025b57600080fd5b81516200026881620001e8565b9392505050565b60805160a05160c05160e05161010051615f49620002fa600039600081816102c40152818161074a01526108060152600081816106f90152612c5101526000818161053c0152612e330152600081816105700152818161300901526131cb0152600081816105a4015281816117210152818161291c01528181612ab40152612cee0152615f496000f3fe60806040526004361061020f5760003560e01c80635c155662116101185780638b00ce7c116100a0578063cefdc1d41161006f578063cefdc1d4146106b9578063df5cf723146106e7578063f2fde38b1461071b578063f5c9899d1461073b578063fabc1cbc1461076e57600080fd5b80638b00ce7c146106445780638da5cb5b14610661578063b98d09081461067f578063c0c53b8b1461069957600080fd5b80636d14a987116100e75780636d14a987146105925780636efb4636146105c6578063715018a6146105f457806372d18e8d14610609578063886f11951461062457600080fd5b80635c155662146104e85780635c975abb146105155780635df459461461052a578063683048351461055e57600080fd5b80633563b0d11161019b5780634d075ce71161016a5780634d075ce7146104305780634f739f741461044657806352e9408b14610473578063595c6a67146104935780635ac86ab7146104a857600080fd5b80633563b0d1146103ae5780633ccfd60b146103db578063416c7e5e146103f0578063459b58b11461041057600080fd5b80631ad43189116101e25780631ad43189146102b25780632191297d146102fb578063245a7bfc1461030e5780632cb223d5146103465780632d89f6fc1461038157600080fd5b8063103642111461021457806310d67a2f14610236578063136439dd14610256578063171f1d5b14610276575b600080fd5b34801561022057600080fd5b5061023461022f366004614ab2565b61078e565b005b34801561024257600080fd5b50610234610251366004614b3b565b610d40565b34801561026257600080fd5b50610234610271366004614b58565b610df3565b34801561028257600080fd5b50610296610291366004614cf8565b610f32565b6040805192151583529015156020830152015b60405180910390f35b3480156102be57600080fd5b506102e67f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016102a9565b610234610309366004614d6d565b6110bc565b34801561031a57600080fd5b5060cd5461032e906001600160a01b031681565b6040516001600160a01b0390911681526020016102a9565b34801561035257600080fd5b50610373610361366004614dd5565b60cc6020526000908152604090205481565b6040519081526020016102a9565b34801561038d57600080fd5b5061037361039c366004614dd5565b60cb6020526000908152604090205481565b3480156103ba57600080fd5b506103ce6103c9366004614df2565b611255565b6040516102a99190614f4d565b3480156103e757600080fd5b506102346116eb565b3480156103fc57600080fd5b5061023461040b366004614f6e565b61171f565b34801561041c57600080fd5b5061023461042b366004614f8b565b611894565b34801561043c57600080fd5b5061037360c95481565b34801561045257600080fd5b50610466610461366004615077565b611ba7565b6040516102a9919061514c565b34801561047f57600080fd5b5061023461048e366004614b58565b6122cd565b34801561049f57600080fd5b506102346122da565b3480156104b457600080fd5b506104d86104c3366004615207565b606654600160ff9092169190911b9081161490565b60405190151581526020016102a9565b3480156104f457600080fd5b50610508610503366004615247565b6123a1565b6040516102a991906152f8565b34801561052157600080fd5b50606654610373565b34801561053657600080fd5b5061032e7f000000000000000000000000000000000000000000000000000000000000000081565b34801561056a57600080fd5b5061032e7f000000000000000000000000000000000000000000000000000000000000000081565b34801561059e57600080fd5b5061032e7f000000000000000000000000000000000000000000000000000000000000000081565b3480156105d257600080fd5b506105e66105e136600461548e565b612569565b6040516102a992919061565a565b34801561060057600080fd5b50610234613480565b34801561061557600080fd5b5060ca5463ffffffff166102e6565b34801561063057600080fd5b5060655461032e906001600160a01b031681565b34801561065057600080fd5b5060ca546102e69063ffffffff1681565b34801561066d57600080fd5b506033546001600160a01b031661032e565b34801561068b57600080fd5b506097546104d89060ff1681565b3480156106a557600080fd5b506102346106b43660046156a3565b613494565b3480156106c557600080fd5b506106d96106d43660046156e3565b6135cf565b6040516102a992919061571a565b3480156106f357600080fd5b5061032e7f000000000000000000000000000000000000000000000000000000000000000081565b34801561072757600080fd5b50610234610736366004614b3b565b613761565b34801561074757600080fd5b507f00000000000000000000000000000000000000000000000000000000000000006102e6565b34801561077a57600080fd5b50610234610789366004614b58565b6137d7565b60cd546001600160a01b031633146107ed5760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c657200000060448201526064015b60405180910390fd5b60006107ff6040860160208701614dd5565b905061082b7f000000000000000000000000000000000000000000000000000000000000000082615751565b63ffffffff164363ffffffff16111561089c5760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320726573706f6e64656420746f207468652060448201526c7461736b20746f6f206c61746560981b60648201526084016107e4565b826108f85760405162461bcd60e51b815260206004820152602660248201527f4f70657261746f7220726573706f6e7365732073686f756c64206e6f7420626560448201526520656d70747960d01b60648201526084016107e4565b60cb60006109096020850185614dd5565b63ffffffff1663ffffffff16815260200190815260200160002054856040516020016109359190615779565b60405160208183030381529060405280519060200120146109685760405162461bcd60e51b81526004016107e490615802565b600060cc8161097a6020860186614dd5565b63ffffffff1663ffffffff16815260200190815260200160002054146109f75760405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b60648201526084016107e4565b60005b83811015610c8d576000858583818110610a1657610a1661585f565b9050602002810190610a289190615875565b602001604051602001610a3b9190615895565b6040516020818303038152906040528051906020012090506000610aac826040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b9050868684818110610ac057610ac061585f565b9050602002810190610ad29190615875565b610ae0906020810190614b3b565b6001600160a01b0316610b5b888886818110610afe57610afe61585f565b9050602002810190610b109190615875565b610b1e9060808101906158c8565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525086939250506139339050565b6001600160a01b031614610bbf5760405162461bcd60e51b815260206004820152602560248201527f496e76616c6964207369676e6174757265206f72206f70657261746f72206164604482015264647265737360d81b60648201526084016107e4565b610bcc6020860186614dd5565b63ffffffff16878785818110610be457610be461585f565b9050602002810190610bf69190615875565b610c07906040810190602001614dd5565b63ffffffff1614610c785760405162461bcd60e51b815260206004820152603560248201527f41676772656761746f7220726573706f6e7365207461736b20696e6469636573604482015274081cda1bdd5b190818994818dbdb9cda5cdd195b9d605a1b60648201526084016107e4565b50508080610c859061590e565b9150506109fa565b50604080516020808201835263ffffffff4316825291519091610cb4918591849101615929565b6040516020818303038152906040528051906020012060cc6000856000016020810190610ce19190614dd5565b63ffffffff1663ffffffff168152602001908152602001600020819055507f5c817ab65498f09a081a8650849b89843e243485c961b53bd0eb8a80a7049af98382604051610d30929190615929565b60405180910390a1505050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d93573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610db79190615972565b6001600160a01b0316336001600160a01b031614610de75760405162461bcd60e51b81526004016107e49061598f565b610df081613959565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015610e3b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e5f91906159d9565b610e7b5760405162461bcd60e51b81526004016107e4906159f6565b60665481811614610ef45760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c697479000000000000000060648201526084016107e4565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000187876000015188602001518860000151600060028110610f7a57610f7a61585f565b60200201518951600160200201518a60200151600060028110610f9f57610f9f61585f565b60200201518b60200151600160028110610fbb57610fbb61585f565b602090810291909101518c518d8301516040516110189a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c61103b9190615a3e565b90506110ae61105461104d8884613a50565b8690613ae7565b61105c613b7b565b6110a46110958561108f604080518082018252600080825260209182015281518083019092526001825260029082015290565b90613a50565b61109e8c613c3b565b90613ae7565b886201d4c0613ccb565b909890975095505050505050565b60018060c9546110cc9190615a60565b34101561111b5760405162461bcd60e51b815260206004820152601f60248201527f4372656174696e672061207461736b2072657175697265732061206665652e0060448201526064016107e4565b6040805160c08101825260ff86811682524363ffffffff166020808401919091526001600160601b0386166060840152908616828401523360808301523460a08301529151909161116e91839101615a7f565b60408051601f19818403018152828252805160209182012060ca805463ffffffff908116600090815260cb90945293909220555416907fdeb7c64f4d939c7b58f45fb9cd5b5118e0b5920d43390c9ba0d47032754c3121906111d1908490615a7f565b60405180910390a260ca546111ed9063ffffffff166001615751565b60ca805463ffffffff191663ffffffff9290921691909117905560c95434111561124e5760c95433906108fc906112249034615ae0565b6040518115909202916000818181858888f1935050505015801561124c573d6000803e3d6000fd5b505b5050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015611297573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112bb9190615972565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa1580156112fd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113219190615972565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611363573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113879190615972565b9050600086516001600160401b038111156113a4576113a4614b71565b6040519080825280602002602001820160405280156113d757816020015b60608152602001906001900390816113c25790505b50905060005b87518110156116df5760008882815181106113fa576113fa61585f565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa15801561145b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526114839190810190615af7565b905080516001600160401b0381111561149e5761149e614b71565b6040519080825280602002602001820160405280156114e957816020015b60408051606081018252600080825260208083018290529282015282526000199092019101816114bc5790505b508484815181106114fc576114fc61585f565b602002602001018190525060005b81518110156116c9576040518060600160405280876001600160a01b03166347b314e885858151811061153f5761153f61585f565b60200260200101516040518263ffffffff1660e01b815260040161156591815260200190565b602060405180830381865afa158015611582573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115a69190615972565b6001600160a01b031681526020018383815181106115c6576115c661585f565b60200260200101518152602001896001600160a01b031663fa28c6278585815181106115f4576115f461585f565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015611650573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116749190615b87565b6001600160601b03168152508585815181106116925761169261585f565b602002602001015182815181106116ab576116ab61585f565b602002602001018190525080806116c19061590e565b91505061150a565b50505080806116d79061590e565b9150506113dd565b50979650505050505050565b6116f3613eef565b60405133904780156108fc02916000818181858888f19350505050158015610df0573d6000803e3d6000fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561177d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117a19190615972565b6001600160a01b0316336001600160a01b03161461184d5760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a4016107e4565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b60ca548290829063ffffffff90811690831611156118ed5760405162461bcd60e51b81526020600482015260166024820152757461736b206e756d626572206e6f742065786973747360501b60448201526064016107e4565b63ffffffff8216600090815260cb602090815260409182902054915161191591849101615a7f565b60405160208183030381529060405280519060200120146119485760405162461bcd60e51b81526004016107e490615802565b60808101516001600160a01b031633146119be5760405162461bcd60e51b815260206004820152603160248201527f4f6e6c7920746865207461736b2063726561746f722063616e20706572666f7260448201527036903a3434b99037b832b930ba34b7b71760791b60648201526084016107e4565b63ffffffff8416600090815260cc6020526040902054611a345760405162461bcd60e51b815260206004820152602b60248201527f43616e6e6f742077697468647261772066756e647320666f72206120636f6d7060448201526a3632ba32b2103a30b9b59760a91b60648201526084016107e4565b60008360a0015111611a805760405162461bcd60e51b8152602060048201526015602482015274273790333ab73239903a37903bb4ba34323930bb9760591b60448201526064016107e4565b60a0830180516000909152604051611a9c908590602001615a7f565b60408051601f19818403018152828252805160209182012063ffffffff8916600090815260cb90925291812091909155608086015190916001600160a01b039091169083908381818185875af1925050503d8060008114611b19576040519150601f19603f3d011682016040523d82523d6000602084013e611b1e565b606091505b5050905080611b645760405162461bcd60e51b81526020600482015260126024820152712bb4ba34323930bbb0b6103330b4b632b21760711b60448201526064016107e4565b6040805163ffffffff881681523360208201529081018390527ff440aec6b52895984d061d622e6edeba6210f7c3e059be920663140c084560d790606001610d30565b611bd26040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015611c12573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c369190615972565b9050611c636040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e90611c93908b9089908990600401615ba4565b600060405180830381865afa158015611cb0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611cd89190810190615bee565b81526040516340e03a8160e11b81526001600160a01b038316906381c0750290611d0a908b908b908b90600401615ca5565b600060405180830381865afa158015611d27573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611d4f9190810190615bee565b6040820152856001600160401b03811115611d6c57611d6c614b71565b604051908082528060200260200182016040528015611d9f57816020015b6060815260200190600190039081611d8a5790505b50606082015260005b60ff81168711156121de576000856001600160401b03811115611dcd57611dcd614b71565b604051908082528060200260200182016040528015611df6578160200160208202803683370190505b5083606001518360ff1681518110611e1057611e1061585f565b602002602001018190525060005b868110156120de5760008c6001600160a01b03166304ec63518a8a85818110611e4957611e4961585f565b905060200201358e88600001518681518110611e6757611e6761585f565b60200260200101516040518463ffffffff1660e01b8152600401611ea49392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611ec1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ee59190615cce565b90506001600160c01b038116611f895760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a4016107e4565b8a8a8560ff16818110611f9e57611f9e61585f565b6001600160c01b03841692013560f81c9190911c6001908116141590506120cb57856001600160a01b031663dd9846b98a8a85818110611fe057611fe061585f565b905060200201358d8d8860ff16818110611ffc57611ffc61585f565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015612052573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120769190615cf7565b85606001518560ff168151811061208f5761208f61585f565b602002602001015184815181106120a8576120a861585f565b63ffffffff90921660209283029190910190910152826120c78161590e565b9350505b50806120d68161590e565b915050611e1e565b506000816001600160401b038111156120f9576120f9614b71565b604051908082528060200260200182016040528015612122578160200160208202803683370190505b50905060005b828110156121a35784606001518460ff16815181106121495761214961585f565b602002602001015181815181106121625761216261585f565b602002602001015182828151811061217c5761217c61585f565b63ffffffff909216602092830291909101909101528061219b8161590e565b915050612128565b508084606001518460ff16815181106121be576121be61585f565b6020026020010181905250505080806121d690615d14565b915050611da8565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa15801561221f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122439190615972565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90612276908b908b908e90600401615d34565b600060405180830381865afa158015612293573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526122bb9190810190615bee565b60208301525098975050505050505050565b6122d5613eef565b60c955565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015612322573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061234691906159d9565b6123625760405162461bcd60e51b81526004016107e4906159f6565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b81526004016123d3929190615d5e565b600060405180830381865afa1580156123f0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526124189190810190615bee565b9050600084516001600160401b0381111561243557612435614b71565b60405190808252806020026020018201604052801561245e578160200160208202803683370190505b50905060005b855181101561255f57866001600160a01b03166304ec635187838151811061248e5761248e61585f565b6020026020010151878685815181106124a9576124a961585f565b60200260200101516040518463ffffffff1660e01b81526004016124e69392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015612503573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125279190615cce565b6001600160c01b03168282815181106125425761254261585f565b6020908102919091010152806125578161590e565b915050612464565b5095945050505050565b60408051808201909152606080825260208201526000846125e05760405162461bcd60e51b81526020600482015260376024820152600080516020615ef483398151915260448201527f7265733a20656d7074792071756f72756d20696e70757400000000000000000060648201526084016107e4565b604083015151851480156125f8575060a08301515185145b8015612608575060c08301515185145b8015612618575060e08301515185145b6126825760405162461bcd60e51b81526020600482015260416024820152600080516020615ef483398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a4016107e4565b825151602084015151146126fa5760405162461bcd60e51b815260206004820152604460248201819052600080516020615ef4833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a4016107e4565b4363ffffffff168463ffffffff16106127695760405162461bcd60e51b815260206004820152603c6024820152600080516020615ef483398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b0000000060648201526084016107e4565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b038111156127aa576127aa614b71565b6040519080825280602002602001820160405280156127d3578160200160208202803683370190505b506020820152866001600160401b038111156127f1576127f1614b71565b60405190808252806020026020018201604052801561281a578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b0381111561284e5761284e614b71565b604051908082528060200260200182016040528015612877578160200160208202803683370190505b5081526020860151516001600160401b0381111561289757612897614b71565b6040519080825280602002602001820160405280156128c0578160200160208202803683370190505b50816020018190525060006129928a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa158015612969573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061298d9190615db2565b613f49565b905060005b876020015151811015612c2d576129dc886020015182815181106129bd576129bd61585f565b6020026020010151805160009081526020918201519091526040902090565b836020015182815181106129f2576129f261585f565b60209081029190910101528015612ab2576020830151612a13600183615ae0565b81518110612a2357612a2361585f565b602002602001015160001c83602001518281518110612a4457612a4461585f565b602002602001015160001c11612ab2576040805162461bcd60e51b8152602060048201526024810191909152600080516020615ef483398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f7274656460648201526084016107e4565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec635184602001518381518110612af757612af761585f565b60200260200101518b8b600001518581518110612b1657612b1661585f565b60200260200101516040518463ffffffff1660e01b8152600401612b539392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015612b70573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b949190615cce565b6001600160c01b031683600001518281518110612bb357612bb361585f565b602002602001018181525050612c1961104d612bed8486600001518581518110612bdf57612bdf61585f565b602002602001015116613fda565b8a602001518481518110612c0357612c0361585f565b602002602001015161400590919063ffffffff16565b945080612c258161590e565b915050612997565b5050612c38836140e9565b60975490935060ff16600081612c4f576000612cd1565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa158015612cad573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612cd19190615dcf565b905060005b8a81101561334f578215612e31578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f86818110612d2d57612d2d61585f565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015612d6d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d919190615dcf565b612d9b9190615de8565b11612e315760405162461bcd60e51b81526020600482015260666024820152600080516020615ef483398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c4016107e4565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d84818110612e7257612e7261585f565b9050013560f81c60f81b60f81c8c8c60a001518581518110612e9657612e9661585f565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612ef2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f169190615e00565b6001600160401b031916612f398a6040015183815181106129bd576129bd61585f565b67ffffffffffffffff191614612fd55760405162461bcd60e51b81526020600482015260616024820152600080516020615ef483398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c4016107e4565b61300589604001518281518110612fee57612fee61585f565b602002602001015187613ae790919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d848181106130485761304861585f565b9050013560f81c60f81b60f81c8c8c60c00151858151811061306c5761306c61585f565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa1580156130c8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130ec9190615b87565b856020015182815181106131025761310261585f565b6001600160601b0390921660209283029190910182015285015180518290811061312e5761312e61585f565b60200260200101518560000151828151811061314c5761314c61585f565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a602001515181101561333a576131c4866000015182815181106131965761319661585f565b60200260200101518f8f868181106131b0576131b061585f565b600192013560f81c9290921c811614919050565b15613328577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f8681811061320a5761320a61585f565b9050013560f81c60f81b60f81c8e8960200151858151811061322e5761322e61585f565b60200260200101518f60e00151888151811061324c5761324c61585f565b602002602001015187815181106132655761326561585f565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa1580156132c9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132ed9190615b87565b87518051859081106133015761330161585f565b602002602001018181516133159190615e2b565b6001600160601b03169052506001909101905b806133328161590e565b915050613170565b505080806133479061590e565b915050612cd6565b5050506000806133698c868a606001518b60800151610f32565b91509150816133da5760405162461bcd60e51b81526020600482015260436024820152600080516020615ef483398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a4016107e4565b8061343b5760405162461bcd60e51b81526020600482015260396024820152600080516020615ef483398151915260448201527f7265733a207369676e617475726520697320696e76616c69640000000000000060648201526084016107e4565b50506000878260200151604051602001613456929190615e53565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b613488613eef565b6134926000614184565b565b600054610100900460ff16158080156134b45750600054600160ff909116105b806134ce5750303b1580156134ce575060005460ff166001145b6135315760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016107e4565b6000805460ff191660011790558015613554576000805461ff0019166101001790555b61355f8460006141d6565b61356883614184565b60cd80546001600160a01b0319166001600160a01b03841617905580156135c9576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b604080516001808252818301909252600091606091839160208083019080368337019050509050848160008151811061360a5761360a61585f565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e906136469088908690600401615d5e565b600060405180830381865afa158015613663573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261368b9190810190615bee565b60008151811061369d5761369d61585f565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa158015613709573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061372d9190615cce565b6001600160c01b031690506000613743826142c0565b9050816137518a838a611255565b9550955050505050935093915050565b613769613eef565b6001600160a01b0381166137ce5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016107e4565b610df081614184565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561382a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061384e9190615972565b6001600160a01b0316336001600160a01b03161461387e5760405162461bcd60e51b81526004016107e49061598f565b6066541981196066541916146138fc5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c697479000000000000000060648201526084016107e4565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610f27565b6000806000613942858561438c565b9150915061394f816143fc565b5090505b92915050565b6001600160a01b0381166139e75760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a4016107e4565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b6040805180820190915260008082526020820152613a6c614994565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa9050808015613a9f57613aa1565bfe5b5080613adf5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b60448201526064016107e4565b505092915050565b6040805180820190915260008082526020820152613b036149b2565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa9050808015613a9f575080613adf5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b60448201526064016107e4565b613b836149d0565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b604080518082019091526000808252602082015260008080613c6b600080516020615ed483398151915286615a3e565b90505b613c77816145b7565b9093509150600080516020615ed4833981519152828309831415613cb1576040805180820190915290815260208101919091529392505050565b600080516020615ed4833981519152600182089050613c6e565b604080518082018252868152602080820186905282518084019093528683528201849052600091829190613cfd6149f5565b60005b6002811015613ec2576000613d16826006615a60565b9050848260028110613d2a57613d2a61585f565b60200201515183613d3c836000615de8565b600c8110613d4c57613d4c61585f565b6020020152848260028110613d6357613d6361585f565b60200201516020015183826001613d7a9190615de8565b600c8110613d8a57613d8a61585f565b6020020152838260028110613da157613da161585f565b6020020151515183613db4836002615de8565b600c8110613dc457613dc461585f565b6020020152838260028110613ddb57613ddb61585f565b6020020151516001602002015183613df4836003615de8565b600c8110613e0457613e0461585f565b6020020152838260028110613e1b57613e1b61585f565b602002015160200151600060028110613e3657613e3661585f565b602002015183613e47836004615de8565b600c8110613e5757613e5761585f565b6020020152838260028110613e6e57613e6e61585f565b602002015160200151600160028110613e8957613e8961585f565b602002015183613e9a836005615de8565b600c8110613eaa57613eaa61585f565b60200201525080613eba8161590e565b915050613d00565b50613ecb614a14565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b6033546001600160a01b031633146134925760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016107e4565b600080613f5584614639565b9050808360ff166001901b11613fd35760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c75650060648201526084016107e4565b9392505050565b6000805b821561395357613fef600184615ae0565b9092169180613ffd81615e9b565b915050613fde565b60408051808201909152600080825260208201526102008261ffff16106140615760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b60448201526064016107e4565b8161ffff1660011415614075575081613953565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff16106140de57600161ffff871660ff83161c811614156140c1576140be8484613ae7565b93505b6140cb8384613ae7565b92506201fffe600192831b169101614091565b509195945050505050565b6040805180820190915260008082526020820152815115801561410e57506020820151155b1561412c575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020615ed4833981519152846020015161415f9190615a3e565b61417790600080516020615ed4833981519152615ae0565b905292915050565b919050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6065546001600160a01b03161580156141f757506001600160a01b03821615155b6142795760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a4016107e4565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a26142bc82613959565b5050565b60606000806142ce84613fda565b61ffff166001600160401b038111156142e9576142e9614b71565b6040519080825280601f01601f191660200182016040528015614313576020820181803683370190505b5090506000805b82518210801561432b575061010081105b15614382576001811b935085841615614372578060f81b8383815181106143545761435461585f565b60200101906001600160f81b031916908160001a9053508160010191505b61437b8161590e565b905061431a565b5090949350505050565b6000808251604114156143c35760208301516040840151606085015160001a6143b7878285856147c6565b945094505050506143f5565b8251604014156143ed57602083015160408401516143e28683836148b3565b9350935050506143f5565b506000905060025b9250929050565b600081600481111561441057614410615ebd565b14156144195750565b600181600481111561442d5761442d615ebd565b141561447b5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016107e4565b600281600481111561448f5761448f615ebd565b14156144dd5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016107e4565b60038160048111156144f1576144f1615ebd565b141561454a5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016107e4565b600481600481111561455e5761455e615ebd565b1415610df05760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b60648201526084016107e4565b60008080600080516020615ed48339815191526003600080516020615ed483398151915286600080516020615ed483398151915288890909089050600061462d827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020615ed48339815191526148ec565b91959194509092505050565b6000610100825111156146c25760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a4016107e4565b81516146d057506000919050565b600080836000815181106146e6576146e661585f565b0160200151600160f89190911c81901b92505b84518110156147bd578481815181106147145761471461585f565b0160200151600160f89190911c1b91508282116147a95760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a4016107e4565b918117916147b68161590e565b90506146f9565b50909392505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156147fd57506000905060036148aa565b8460ff16601b1415801561481557508460ff16601c14155b1561482657506000905060046148aa565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561487a573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166148a3576000600192509250506148aa565b9150600090505b94509492505050565b6000806001600160ff1b038316816148d060ff86901c601b615de8565b90506148de878288856147c6565b935093505050935093915050565b6000806148f7614a14565b6148ff614a32565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa9250828015613a9f5750826149895760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c75726500000000000060448201526064016107e4565b505195945050505050565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b60405180604001604052806149e3614a50565b81526020016149f0614a50565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b60008083601f840112614a8057600080fd5b5081356001600160401b03811115614a9757600080fd5b6020830191508360208260051b85010111156143f557600080fd5b600080600080848603610160811215614aca57600080fd5b60c0811215614ad857600080fd5b85945060c08601356001600160401b03811115614af457600080fd5b614b0088828901614a6e565b909550935050608060df1982011215614b1857600080fd5b5092959194509260e0019150565b6001600160a01b0381168114610df057600080fd5b600060208284031215614b4d57600080fd5b8135613fd381614b26565b600060208284031215614b6a57600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715614ba957614ba9614b71565b60405290565b60405160c081016001600160401b0381118282101715614ba957614ba9614b71565b60405161010081016001600160401b0381118282101715614ba957614ba9614b71565b604051601f8201601f191681016001600160401b0381118282101715614c1c57614c1c614b71565b604052919050565b600060408284031215614c3657600080fd5b614c3e614b87565b9050813581526020820135602082015292915050565b600082601f830112614c6557600080fd5b604051604081018181106001600160401b0382111715614c8757614c87614b71565b8060405250806040840185811115614c9e57600080fd5b845b818110156140de578035835260209283019201614ca0565b600060808284031215614cca57600080fd5b614cd2614b87565b9050614cde8383614c54565b8152614ced8360408401614c54565b602082015292915050565b6000806000806101208587031215614d0f57600080fd5b84359350614d208660208701614c24565b9250614d2f8660608701614cb8565b9150614d3e8660e08701614c24565b905092959194509250565b60ff81168114610df057600080fd5b6001600160601b0381168114610df057600080fd5b600080600060608486031215614d8257600080fd5b8335614d8d81614d49565b92506020840135614d9d81614d49565b91506040840135614dad81614d58565b809150509250925092565b63ffffffff81168114610df057600080fd5b803561417f81614db8565b600060208284031215614de757600080fd5b8135613fd381614db8565b600080600060608486031215614e0757600080fd5b8335614e1281614b26565b92506020848101356001600160401b0380821115614e2f57600080fd5b818701915087601f830112614e4357600080fd5b813581811115614e5557614e55614b71565b614e67601f8201601f19168501614bf4565b91508082528884828501011115614e7d57600080fd5b8084840185840137600084828401015250809450505050614ea060408501614dca565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015614f3f578385038a52825180518087529087019087870190845b81811015614f2a57835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101614ee6565b50509a87019a95505091850191600101614ec8565b509298975050505050505050565b602081526000613fd36020830184614ea9565b8015158114610df057600080fd5b600060208284031215614f8057600080fd5b8135613fd381614f60565b60008082840360e0811215614f9f57600080fd5b8335614faa81614db8565b925060c0601f1982011215614fbe57600080fd5b50614fc7614baf565b6020840135614fd581614d49565b81526040840135614fe581614db8565b60208201526060840135614ff881614d49565b6040820152608084013561500b81614d58565b606082015260a084013561501e81614b26565b608082015260c0939093013560a08401525092909150565b60008083601f84011261504857600080fd5b5081356001600160401b0381111561505f57600080fd5b6020830191508360208285010111156143f557600080fd5b6000806000806000806080878903121561509057600080fd5b863561509b81614b26565b955060208701356150ab81614db8565b945060408701356001600160401b03808211156150c757600080fd5b6150d38a838b01615036565b909650945060608901359150808211156150ec57600080fd5b506150f989828a01614a6e565b979a9699509497509295939492505050565b600081518084526020808501945080840160005b8381101561514157815163ffffffff168752958201959082019060010161511f565b509495945050505050565b60006020808352835160808285015261516860a085018261510b565b905081850151601f1980868403016040870152615185838361510b565b925060408701519150808684030160608701526151a2838361510b565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b828110156151f957848783030184526151e782875161510b565b958801959388019391506001016151cd565b509998505050505050505050565b60006020828403121561521957600080fd5b8135613fd381614d49565b60006001600160401b0382111561523d5761523d614b71565b5060051b60200190565b60008060006060848603121561525c57600080fd5b833561526781614b26565b92506020848101356001600160401b0381111561528357600080fd5b8501601f8101871361529457600080fd5b80356152a76152a282615224565b614bf4565b81815260059190911b820183019083810190898311156152c657600080fd5b928401925b828410156152e4578335825292840192908401906152cb565b8096505050505050614ea060408501614dca565b6020808252825182820181905260009190848201906040850190845b8181101561533057835183529284019291840191600101615314565b50909695505050505050565b600082601f83011261534d57600080fd5b8135602061535d6152a283615224565b82815260059290921b8401810191818101908684111561537c57600080fd5b8286015b848110156153a057803561539381614db8565b8352918301918301615380565b509695505050505050565b600082601f8301126153bc57600080fd5b813560206153cc6152a283615224565b82815260069290921b840181019181810190868411156153eb57600080fd5b8286015b848110156153a0576154018882614c24565b8352918301916040016153ef565b600082601f83011261542057600080fd5b813560206154306152a283615224565b82815260059290921b8401810191818101908684111561544f57600080fd5b8286015b848110156153a05780356001600160401b038111156154725760008081fd5b6154808986838b010161533c565b845250918301918301615453565b6000806000806000608086880312156154a657600080fd5b8535945060208601356001600160401b03808211156154c457600080fd5b6154d089838a01615036565b9096509450604088013591506154e582614db8565b909250606087013590808211156154fb57600080fd5b90870190610180828a03121561551057600080fd5b615518614bd1565b82358281111561552757600080fd5b6155338b82860161533c565b82525060208301358281111561554857600080fd5b6155548b8286016153ab565b60208301525060408301358281111561556c57600080fd5b6155788b8286016153ab565b60408301525061558b8a60608501614cb8565b606082015261559d8a60e08501614c24565b6080820152610120830135828111156155b557600080fd5b6155c18b82860161533c565b60a083015250610140830135828111156155da57600080fd5b6155e68b82860161533c565b60c083015250610160830135828111156155ff57600080fd5b61560b8b82860161540f565b60e0830152508093505050509295509295909350565b600081518084526020808501945080840160005b838110156151415781516001600160601b031687529582019590820190600101615635565b60408152600083516040808401526156756080840182615621565b90506020850151603f198483030160608501526156928282615621565b925050508260208301529392505050565b6000806000606084860312156156b857600080fd5b83356156c381614b26565b925060208401356156d381614b26565b91506040840135614dad81614b26565b6000806000606084860312156156f857600080fd5b833561570381614b26565b9250602084013591506040840135614dad81614db8565b8281526040602082015260006157336040830184614ea9565b949350505050565b634e487b7160e01b600052601160045260246000fd5b600063ffffffff8083168185168083038211156157705761577061573b565b01949350505050565b60c08101823561578881614d49565b60ff168252602083013561579b81614db8565b63ffffffff16602083015260408301356157b481614d49565b60ff16604083015260608301356157ca81614d58565b6001600160601b0316606083015260808301356157e681614b26565b6001600160a01b0316608083015260a092830135919092015290565b6020808252603d908201527f737570706c696564207461736b20646f6573206e6f74206d617463682074686560408201527f206f6e65207265636f7264656420696e2074686520636f6e7472616374000000606082015260800190565b634e487b7160e01b600052603260045260246000fd5b60008235609e1983360301811261588b57600080fd5b9190910192915050565b6060810182356158a481614db8565b63ffffffff8116835250602083013560208301526040830135604083015292915050565b6000808335601e198436030181126158df57600080fd5b8301803591506001600160401b038211156158f957600080fd5b6020019150368190038213156143f557600080fd5b60006000198214156159225761592261573b565b5060010190565b60a08101833561593881614db8565b63ffffffff808216845260208601356020850152604086013560408501526060860135606085015280855116608085015250509392505050565b60006020828403121561598457600080fd5b8151613fd381614b26565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b6000602082840312156159eb57600080fd5b8151613fd381614f60565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b600082615a5b57634e487b7160e01b600052601260045260246000fd5b500690565b6000816000190483118215151615615a7a57615a7a61573b565b500290565b600060c08201905060ff835116825263ffffffff602084015116602083015260ff60408401511660408301526001600160601b03606084015116606083015260018060a01b03608084015116608083015260a083015160a083015292915050565b600082821015615af257615af261573b565b500390565b60006020808385031215615b0a57600080fd5b82516001600160401b03811115615b2057600080fd5b8301601f81018513615b3157600080fd5b8051615b3f6152a282615224565b81815260059190911b82018301908381019087831115615b5e57600080fd5b928401925b82841015615b7c57835182529284019290840190615b63565b979650505050505050565b600060208284031215615b9957600080fd5b8151613fd381614d58565b63ffffffff84168152604060208201819052810182905260006001600160fb1b03831115615bd157600080fd5b8260051b8085606085013760009201606001918252509392505050565b60006020808385031215615c0157600080fd5b82516001600160401b03811115615c1757600080fd5b8301601f81018513615c2857600080fd5b8051615c366152a282615224565b81815260059190911b82018301908381019087831115615c5557600080fd5b928401925b82841015615b7c578351615c6d81614db8565b82529284019290840190615c5a565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff84168152604060208201526000615cc5604083018486615c7c565b95945050505050565b600060208284031215615ce057600080fd5b81516001600160c01b0381168114613fd357600080fd5b600060208284031215615d0957600080fd5b8151613fd381614db8565b600060ff821660ff811415615d2b57615d2b61573b565b60010192915050565b604081526000615d48604083018587615c7c565b905063ffffffff83166020830152949350505050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b81811015615da557845183529383019391830191600101615d89565b5090979650505050505050565b600060208284031215615dc457600080fd5b8151613fd381614d49565b600060208284031215615de157600080fd5b5051919050565b60008219821115615dfb57615dfb61573b565b500190565b600060208284031215615e1257600080fd5b815167ffffffffffffffff1981168114613fd357600080fd5b60006001600160601b0383811690831681811015615e4b57615e4b61573b565b039392505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b83811015615e8e57815185529382019390820190600101615e72565b5092979650505050505050565b600061ffff80831681811415615eb357615eb361573b565b6001019392505050565b634e487b7160e01b600052602160045260246000fdfe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a264697066735822122004f87bc98282f676aa054a66abe46f33b879a9889b33ecc4b5aecfa7481acbd664736f6c634300080c0033",
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

// CreateNewTask is a paid mutator transaction binding the contract method 0x2191297d.
//
// Solidity: function createNewTask(uint8 taskType, uint8 responderNumber, uint96 stakeThreshold) payable returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactor) CreateNewTask(opts *bind.TransactOpts, taskType uint8, responderNumber uint8, stakeThreshold *big.Int) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.contract.Transact(opts, "createNewTask", taskType, responderNumber, stakeThreshold)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x2191297d.
//
// Solidity: function createNewTask(uint8 taskType, uint8 responderNumber, uint96 stakeThreshold) payable returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) CreateNewTask(taskType uint8, responderNumber uint8, stakeThreshold *big.Int) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.CreateNewTask(&_ContractOpenOracleTaskManager.TransactOpts, taskType, responderNumber, stakeThreshold)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x2191297d.
//
// Solidity: function createNewTask(uint8 taskType, uint8 responderNumber, uint96 stakeThreshold) payable returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactorSession) CreateNewTask(taskType uint8, responderNumber uint8, stakeThreshold *big.Int) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.CreateNewTask(&_ContractOpenOracleTaskManager.TransactOpts, taskType, responderNumber, stakeThreshold)
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
	TaskResponse         IOpenOracleTaskManagerWeightedTaskResponse
	TaskResponseMetadata IOpenOracleTaskManagerTaskResponseMetadata
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTaskResponded is a free log retrieval operation binding the contract event 0x5c817ab65498f09a081a8650849b89843e243485c961b53bd0eb8a80a7049af9.
//
// Solidity: event TaskResponded((uint32,uint256,uint256,uint256) taskResponse, (uint32) taskResponseMetadata)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) FilterTaskResponded(opts *bind.FilterOpts) (*ContractOpenOracleTaskManagerTaskRespondedIterator, error) {

	logs, sub, err := _ContractOpenOracleTaskManager.contract.FilterLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleTaskManagerTaskRespondedIterator{contract: _ContractOpenOracleTaskManager.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0x5c817ab65498f09a081a8650849b89843e243485c961b53bd0eb8a80a7049af9.
//
// Solidity: event TaskResponded((uint32,uint256,uint256,uint256) taskResponse, (uint32) taskResponseMetadata)
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

// ParseTaskResponded is a log parse operation binding the contract event 0x5c817ab65498f09a081a8650849b89843e243485c961b53bd0eb8a80a7049af9.
//
// Solidity: event TaskResponded((uint32,uint256,uint256,uint256) taskResponse, (uint32) taskResponseMetadata)
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

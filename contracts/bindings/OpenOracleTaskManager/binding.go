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

// IOpenOracleTaskManagerTask is an auto generated low-level Go binding around an user-defined struct.
type IOpenOracleTaskManagerTask struct {
	TaskType                  uint8
	TaskCreatedBlock          uint32
	QuorumNumbers             []byte
	QuorumThresholdPercentage uint32
	Creator                   common.Address
	CreationFee               *big.Int
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
	HashOfNonSigners   [32]byte
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"taskType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestTaskNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIOpenOracleTaskManager.Task\",\"components\":[{\"name\":\"taskType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"creationFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIOpenOracleTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeStamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskCreationFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTaskCreationFee\",\"inputs\":[{\"name\":\"_newFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawTaskFunds\",\"inputs\":[{\"name\":\"taskNum\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIOpenOracleTaskManager.Task\",\"components\":[{\"name\":\"taskType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"creationFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FundsWithdrawn\",\"inputs\":[{\"name\":\"taskNum\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"creator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOpenOracleTaskManager.Task\",\"components\":[{\"name\":\"taskType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"creationFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOpenOracleTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeStamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOpenOracleTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x610120604052655af3107a400060c9553480156200001c57600080fd5b5060405162005eb438038062005eb48339810160408190526200003f9162000201565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000099573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000bf919062000248565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000117573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013d919062000248565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000197573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001bd919062000248565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff1661010052506200026f565b6001600160a01b0381168114620001fe57600080fd5b50565b600080604083850312156200021557600080fd5b82516200022281620001e8565b602084015190925063ffffffff811681146200023d57600080fd5b809150509250929050565b6000602082840312156200025b57600080fd5b81516200026881620001e8565b9392505050565b60805160a05160c05160e05161010051615bba620002fa600039600081816102a40152818161074a0152612ef90152600081816106f9015261251a01526000818161050901526126fd01526000818161053d015281816128d30152612a95015260008181610571015281816112fd015281816121e50152818161237d01526125b70152615bba6000f3fe60806040526004361061020f5760003560e01c80635df45946116101185780638da5cb5b116100a0578063cefdc1d41161006f578063cefdc1d4146106b9578063df5cf723146106e7578063f2fde38b1461071b578063f5c9899d1461073b578063fabc1cbc1461076e57600080fd5b80638da5cb5b1461064e578063a7a69b051461066c578063b98d09081461067f578063c0c53b8b1461069957600080fd5b8063715018a6116100e7578063715018a6146105c157806372d18e8d146105d6578063886f1195146105f157806388fee238146106115780638b00ce7c1461063157600080fd5b80635df45946146104f7578063683048351461052b5780636d14a9871461055f5780636efb46361461059357600080fd5b80633ccfd60b1161019b57806352e9408b1161016a57806352e9408b14610440578063595c6a67146104605780635ac86ab7146104755780635c155662146104b55780635c975abb146104e257600080fd5b80633ccfd60b146103c8578063416c7e5e146103dd5780634d075ce7146103fd5780634f739f741461041357600080fd5b8063245a7bfc116101e2578063245a7bfc146102db5780632cb223d5146103135780632d89f6fc1461034e578063352f52cb1461037b5780633563b0d11461039b57600080fd5b806310d67a2f14610214578063136439dd14610236578063171f1d5b146102565780631ad4318914610292575b600080fd5b34801561022057600080fd5b5061023461022f3660046145df565b61078e565b005b34801561024257600080fd5b506102346102513660046145fc565b61084a565b34801561026257600080fd5b5061027661027136600461479c565b610989565b6040805192151583529015156020830152015b60405180910390f35b34801561029e57600080fd5b506102c67f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff9091168152602001610289565b3480156102e757600080fd5b5060cd546102fb906001600160a01b031681565b6040516001600160a01b039091168152602001610289565b34801561031f57600080fd5b5061034061032e36600461480a565b60cc6020526000908152604090205481565b604051908152602001610289565b34801561035a57600080fd5b5061034061036936600461480a565b60cb6020526000908152604090205481565b34801561038757600080fd5b506102346103963660046148b0565b610b13565b3480156103a757600080fd5b506103bb6103b6366004614988565b610e31565b6040516102899190614a8f565b3480156103d457600080fd5b506102346112c7565b3480156103e957600080fd5b506102346103f8366004614ab7565b6112fb565b34801561040957600080fd5b5061034060c95481565b34801561041f57600080fd5b5061043361042e366004614b1c565b611470565b6040516102899190614c20565b34801561044c57600080fd5b5061023461045b3660046145fc565b611b96565b34801561046c57600080fd5b50610234611ba3565b34801561048157600080fd5b506104a5610490366004614cdb565b606654600160ff9092169190911b9081161490565b6040519015158152602001610289565b3480156104c157600080fd5b506104d56104d0366004614d1b565b611c6a565b6040516102899190614dd5565b3480156104ee57600080fd5b50606654610340565b34801561050357600080fd5b506102fb7f000000000000000000000000000000000000000000000000000000000000000081565b34801561053757600080fd5b506102fb7f000000000000000000000000000000000000000000000000000000000000000081565b34801561056b57600080fd5b506102fb7f000000000000000000000000000000000000000000000000000000000000000081565b34801561059f57600080fd5b506105b36105ae366004615095565b611e32565b604051610289929190615155565b3480156105cd57600080fd5b50610234612d4a565b3480156105e257600080fd5b5060ca5463ffffffff166102c6565b3480156105fd57600080fd5b506065546102fb906001600160a01b031681565b34801561061d57600080fd5b5061023461062c36600461519e565b612d5e565b34801561063d57600080fd5b5060ca546102c69063ffffffff1681565b34801561065a57600080fd5b506033546001600160a01b03166102fb565b61023461067a366004615226565b613187565b34801561068b57600080fd5b506097546104a59060ff1681565b3480156106a557600080fd5b506102346106b436600461528a565b613365565b3480156106c557600080fd5b506106d96106d43660046152ca565b6134a0565b604051610289929190615301565b3480156106f357600080fd5b506102fb7f000000000000000000000000000000000000000000000000000000000000000081565b34801561072757600080fd5b506102346107363660046145df565b613632565b34801561074757600080fd5b507f00000000000000000000000000000000000000000000000000000000000000006102c6565b34801561077a57600080fd5b506102346107893660046145fc565b6136a8565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156107e1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108059190615322565b6001600160a01b0316336001600160a01b03161461083e5760405162461bcd60e51b81526004016108359061533f565b60405180910390fd5b61084781613804565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015610892573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108b69190615389565b6108d25760405162461bcd60e51b8152600401610835906153a6565b6066548181161461094b5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610835565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001878760000151886020015188600001516000600281106109d1576109d16153ee565b60200201518951600160200201518a602001516000600281106109f6576109f66153ee565b60200201518b60200151600160028110610a1257610a126153ee565b602090810291909101518c518d830151604051610a6f9a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c610a929190615404565b9050610b05610aab610aa488846138fb565b8690613992565b610ab3613a26565b610afb610aec85610ae6604080518082018252600080825260209182015281518083019092526001825260029082015290565b906138fb565b610af58c613ae6565b90613992565b886201d4c0613b76565b909890975095505050505050565b60ca548290829063ffffffff9081169083161115610b6c5760405162461bcd60e51b81526020600482015260166024820152757461736b206e756d626572206e6f742065786973747360501b6044820152606401610835565b63ffffffff8216600090815260cb6020908152604091829020549151610b9491849101615426565b6040516020818303038152906040528051906020012014610bc75760405162461bcd60e51b8152600401610835906154d7565b60808101516001600160a01b03163314610c3d5760405162461bcd60e51b815260206004820152603160248201527f4f6e6c7920746865207461736b2063726561746f722063616e20706572666f7260448201527036903a3434b99037b832b930ba34b7b71760791b6064820152608401610835565b63ffffffff8416600090815260cc6020526040902054610cb35760405162461bcd60e51b815260206004820152602b60248201527f43616e6e6f742077697468647261772066756e647320666f72206120636f6d7060448201526a3632ba32b2103a30b9b59760a91b6064820152608401610835565b60008360a0015111610cff5760405162461bcd60e51b8152602060048201526015602482015274273790333ab73239903a37903bb4ba34323930bb9760591b6044820152606401610835565b60a0830180516000909152604051610d1b908590602001615426565b60408051601f19818403018152828252805160209182012063ffffffff8916600090815260cb90925291812091909155608086015190916001600160a01b039091169083908381818185875af1925050503d8060008114610d98576040519150601f19603f3d011682016040523d82523d6000602084013e610d9d565b606091505b5050905080610de35760405162461bcd60e51b81526020600482015260126024820152712bb4ba34323930bbb0b6103330b4b632b21760711b6044820152606401610835565b6040805163ffffffff881681523360208201529081018390527ff440aec6b52895984d061d622e6edeba6210f7c3e059be920663140c084560d79060600160405180910390a1505050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e73573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e979190615322565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ed9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610efd9190615322565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f3f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f639190615322565b9050600086516001600160401b03811115610f8057610f80614615565b604051908082528060200260200182016040528015610fb357816020015b6060815260200190600190039081610f9e5790505b50905060005b87518110156112bb576000888281518110610fd657610fd66153ee565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015611037573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261105f9190810190615534565b905080516001600160401b0381111561107a5761107a614615565b6040519080825280602002602001820160405280156110c557816020015b60408051606081018252600080825260208083018290529282015282526000199092019101816110985790505b508484815181106110d8576110d86153ee565b602002602001018190525060005b81518110156112a5576040518060600160405280876001600160a01b03166347b314e885858151811061111b5761111b6153ee565b60200260200101516040518263ffffffff1660e01b815260040161114191815260200190565b602060405180830381865afa15801561115e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111829190615322565b6001600160a01b031681526020018383815181106111a2576111a26153ee565b60200260200101518152602001896001600160a01b031663fa28c6278585815181106111d0576111d06153ee565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa15801561122c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061125091906155c4565b6001600160601b031681525085858151811061126e5761126e6153ee565b60200260200101518281518110611287576112876153ee565b6020026020010181905250808061129d90615603565b9150506110e6565b50505080806112b390615603565b915050610fb9565b50979650505050505050565b6112cf613d9a565b60405133904780156108fc02916000818181858888f19350505050158015610847573d6000803e3d6000fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611359573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061137d9190615322565b6001600160a01b0316336001600160a01b0316146114295760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a401610835565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b61149b6040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156114db573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114ff9190615322565b905061152c6040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e9061155c908b908990899060040161561e565b600060405180830381865afa158015611579573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526115a19190810190615668565b81526040516340e03a8160e11b81526001600160a01b038316906381c07502906115d3908b908b908b9060040161571f565b600060405180830381865afa1580156115f0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526116189190810190615668565b6040820152856001600160401b0381111561163557611635614615565b60405190808252806020026020018201604052801561166857816020015b60608152602001906001900390816116535790505b50606082015260005b60ff8116871115611aa7576000856001600160401b0381111561169657611696614615565b6040519080825280602002602001820160405280156116bf578160200160208202803683370190505b5083606001518360ff16815181106116d9576116d96153ee565b602002602001018190525060005b868110156119a75760008c6001600160a01b03166304ec63518a8a85818110611712576117126153ee565b905060200201358e88600001518681518110611730576117306153ee565b60200260200101516040518463ffffffff1660e01b815260040161176d9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa15801561178a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117ae9190615748565b90506001600160c01b0381166118525760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a401610835565b8a8a8560ff16818110611867576118676153ee565b6001600160c01b03841692013560f81c9190911c60019081161415905061199457856001600160a01b031663dd9846b98a8a858181106118a9576118a96153ee565b905060200201358d8d8860ff168181106118c5576118c56153ee565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa15801561191b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061193f9190615771565b85606001518560ff1681518110611958576119586153ee565b60200260200101518481518110611971576119716153ee565b63ffffffff909216602092830291909101909101528261199081615603565b9350505b508061199f81615603565b9150506116e7565b506000816001600160401b038111156119c2576119c2614615565b6040519080825280602002602001820160405280156119eb578160200160208202803683370190505b50905060005b82811015611a6c5784606001518460ff1681518110611a1257611a126153ee565b60200260200101518181518110611a2b57611a2b6153ee565b6020026020010151828281518110611a4557611a456153ee565b63ffffffff9092166020928302919091019091015280611a6481615603565b9150506119f1565b508084606001518460ff1681518110611a8757611a876153ee565b602002602001018190525050508080611a9f9061578e565b915050611671565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611ae8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b0c9190615322565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90611b3f908b908b908e906004016157ae565b600060405180830381865afa158015611b5c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611b849190810190615668565b60208301525098975050505050505050565b611b9e613d9a565b60c955565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015611beb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c0f9190615389565b611c2b5760405162461bcd60e51b8152600401610835906153a6565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b8152600401611c9c9291906157d8565b600060405180830381865afa158015611cb9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611ce19190810190615668565b9050600084516001600160401b03811115611cfe57611cfe614615565b604051908082528060200260200182016040528015611d27578160200160208202803683370190505b50905060005b8551811015611e2857866001600160a01b03166304ec6351878381518110611d5757611d576153ee565b602002602001015187868581518110611d7257611d726153ee565b60200260200101516040518463ffffffff1660e01b8152600401611daf9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611dcc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611df09190615748565b6001600160c01b0316828281518110611e0b57611e0b6153ee565b602090810291909101015280611e2081615603565b915050611d2d565b5095945050505050565b6040805180820190915260608082526020820152600084611ea95760405162461bcd60e51b81526020600482015260376024820152600080516020615b6583398151915260448201527f7265733a20656d7074792071756f72756d20696e7075740000000000000000006064820152608401610835565b60408301515185148015611ec1575060a08301515185145b8015611ed1575060c08301515185145b8015611ee1575060e08301515185145b611f4b5760405162461bcd60e51b81526020600482015260416024820152600080516020615b6583398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a401610835565b82515160208401515114611fc35760405162461bcd60e51b815260206004820152604460248201819052600080516020615b65833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610835565b4363ffffffff168463ffffffff16106120325760405162461bcd60e51b815260206004820152603c6024820152600080516020615b6583398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b000000006064820152608401610835565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b0381111561207357612073614615565b60405190808252806020026020018201604052801561209c578160200160208202803683370190505b506020820152866001600160401b038111156120ba576120ba614615565b6040519080825280602002602001820160405280156120e3578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b0381111561211757612117614615565b604051908082528060200260200182016040528015612140578160200160208202803683370190505b5081526020860151516001600160401b0381111561216057612160614615565b604051908082528060200260200182016040528015612189578160200160208202803683370190505b508160200181905250600061225b8a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa158015612232573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612256919061582c565b613df4565b905060005b8760200151518110156124f6576122a588602001518281518110612286576122866153ee565b6020026020010151805160009081526020918201519091526040902090565b836020015182815181106122bb576122bb6153ee565b6020908102919091010152801561237b5760208301516122dc600183615849565b815181106122ec576122ec6153ee565b602002602001015160001c8360200151828151811061230d5761230d6153ee565b602002602001015160001c1161237b576040805162461bcd60e51b8152602060048201526024810191909152600080516020615b6583398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610835565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec6351846020015183815181106123c0576123c06153ee565b60200260200101518b8b6000015185815181106123df576123df6153ee565b60200260200101516040518463ffffffff1660e01b815260040161241c9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015612439573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061245d9190615748565b6001600160c01b03168360000151828151811061247c5761247c6153ee565b6020026020010181815250506124e2610aa46124b684866000015185815181106124a8576124a86153ee565b602002602001015116613e87565b8a6020015184815181106124cc576124cc6153ee565b6020026020010151613eb290919063ffffffff16565b9450806124ee81615603565b915050612260565b505061250183613f96565b60975490935060ff1660008161251857600061259a565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa158015612576573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061259a9190615860565b905060005b8a811015612c195782156126fb578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f868181106125f6576125f66153ee565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015612636573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061265a9190615860565b6126649190615879565b10156126fb5760405162461bcd60e51b81526020600482015260666024820152600080516020615b6583398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c401610835565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d8481811061273c5761273c6153ee565b9050013560f81c60f81b60f81c8c8c60a001518581518110612760576127606153ee565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa1580156127bc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127e09190615891565b6001600160401b0319166128038a604001518381518110612286576122866153ee565b67ffffffffffffffff19161461289f5760405162461bcd60e51b81526020600482015260616024820152600080516020615b6583398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610835565b6128cf896040015182815181106128b8576128b86153ee565b60200260200101518761399290919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d84818110612912576129126153ee565b9050013560f81c60f81b60f81c8c8c60c001518581518110612936576129366153ee565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612992573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129b691906155c4565b856020015182815181106129cc576129cc6153ee565b6001600160601b039092166020928302919091018201528501518051829081106129f8576129f86153ee565b602002602001015185600001518281518110612a1657612a166153ee565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a6020015151811015612c0457612a8e86600001518281518110612a6057612a606153ee565b60200260200101518f8f86818110612a7a57612a7a6153ee565b600192013560f81c9290921c811614919050565b15612bf2577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f86818110612ad457612ad46153ee565b9050013560f81c60f81b60f81c8e89602001518581518110612af857612af86153ee565b60200260200101518f60e001518881518110612b1657612b166153ee565b60200260200101518781518110612b2f57612b2f6153ee565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa158015612b93573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612bb791906155c4565b8751805185908110612bcb57612bcb6153ee565b60200260200101818151612bdf91906158bc565b6001600160601b03169052506001909101905b80612bfc81615603565b915050612a3a565b50508080612c1190615603565b91505061259f565b505050600080612c338c868a606001518b60800151610989565b9150915081612ca45760405162461bcd60e51b81526020600482015260436024820152600080516020615b6583398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610835565b80612d055760405162461bcd60e51b81526020600482015260396024820152600080516020615b6583398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610835565b50506000878260200151604051602001612d209291906158e4565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b612d52613d9a565b612d5c6000614031565b565b60cd546001600160a01b03163314612db85760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c65720000006044820152606401610835565b6000612dca604085016020860161480a565b9050366000612ddc604087018761592c565b90925090506000612df3608088016060890161480a565b905060cb6000612e06602089018961480a565b63ffffffff1663ffffffff1681526020019081526020016000205487604051602001612e329190615972565b6040516020818303038152906040528051906020012014612e655760405162461bcd60e51b8152600401610835906154d7565b600060cc81612e7760208a018a61480a565b63ffffffff1663ffffffff1681526020019081526020016000205414612ef45760405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b6064820152608401610835565b612f1e7f000000000000000000000000000000000000000000000000000000000000000085615a4a565b63ffffffff164363ffffffff161115612f8f5760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320726573706f6e64656420746f207468652060448201526c7461736b20746f6f206c61746560981b6064820152608401610835565b600086604051602001612fa29190615a9a565b604051602081830303815290604052805190602001209050600080612fca8387878a8c611e32565b9150915060005b858110156130c9578460ff1683602001518281518110612ff357612ff36153ee565b60200260200101516130059190615aa8565b6001600160601b0316606484600001518381518110613026576130266153ee565b60200260200101516001600160601b03166130419190615ad7565b10156130b7576040805162461bcd60e51b81526020600482015260248101919091527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d6064820152608401610835565b806130c181615603565b915050612fd1565b5060408051808201825263ffffffff431681526020808201849052915190916130f6918c91849101615af6565b6040516020818303038152906040528051906020012060cc60008c6000016020810190613123919061480a565b63ffffffff1663ffffffff168152602001908152602001600020819055507fc7b8513936e453ed22d56dc5c765bceee58227d50ca6ae9ac750430339cd900b8a82604051613172929190615af6565b60405180910390a15050505050505050505050565b60018060c9546131979190615ad7565b3410156131e65760405162461bcd60e51b815260206004820152601f60248201527f4372656174696e672061207461736b2072657175697265732061206665652e006044820152606401610835565b6040805160c081018252606081830181905260006080830181905260a083015260ff8816825263ffffffff438116602080850191909152908816918301919091528251601f860182900482028101820190935284835290919085908590819084018382808284376000920191909152505050506040808301919091523360808301523460a08301525161327d908290602001615426565b60408051601f19818403018152828252805160209182012060ca805463ffffffff908116600090815260cb90945293909220555416907f767aa0f97f690f2aa73dc95d49893c4881872e9f1bd64e358a57a6e349f1dcd3906132e0908490615426565b60405180910390a260ca546132fc9063ffffffff166001615a4a565b60ca805463ffffffff191663ffffffff9290921691909117905560c95434111561335d5760c95433906108fc906133339034615849565b6040518115909202916000818181858888f1935050505015801561335b573d6000803e3d6000fd5b505b505050505050565b600054610100900460ff16158080156133855750600054600160ff909116105b8061339f5750303b15801561339f575060005460ff166001145b6134025760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610835565b6000805460ff191660011790558015613425576000805461ff0019166101001790555b613430846000614083565b61343983614031565b60cd80546001600160a01b0319166001600160a01b038416179055801561349a576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b60408051600180825281830190925260009160609183916020808301908036833701905050905084816000815181106134db576134db6153ee565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e9061351790889086906004016157d8565b600060405180830381865afa158015613534573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261355c9190810190615668565b60008151811061356e5761356e6153ee565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa1580156135da573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906135fe9190615748565b6001600160c01b0316905060006136148261416d565b9050816136228a838a610e31565b9550955050505050935093915050565b61363a613d9a565b6001600160a01b03811661369f5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610835565b61084781614031565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156136fb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061371f9190615322565b6001600160a01b0316336001600160a01b03161461374f5760405162461bcd60e51b81526004016108359061533f565b6066541981196066541916146137cd5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610835565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200161097e565b6001600160a01b0381166138925760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610835565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b60408051808201909152600080825260208201526139176144f0565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa905080801561394a5761394c565bfe5b508061398a5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610835565b505092915050565b60408051808201909152600080825260208201526139ae61450e565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa905080801561394a57508061398a5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610835565b613a2e61452c565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b604080518082019091526000808252602082015260008080613b16600080516020615b4583398151915286615404565b90505b613b2281614239565b9093509150600080516020615b45833981519152828309831415613b5c576040805180820190915290815260208101919091529392505050565b600080516020615b45833981519152600182089050613b19565b604080518082018252868152602080820186905282518084019093528683528201849052600091829190613ba8614551565b60005b6002811015613d6d576000613bc1826006615ad7565b9050848260028110613bd557613bd56153ee565b60200201515183613be7836000615879565b600c8110613bf757613bf76153ee565b6020020152848260028110613c0e57613c0e6153ee565b60200201516020015183826001613c259190615879565b600c8110613c3557613c356153ee565b6020020152838260028110613c4c57613c4c6153ee565b6020020151515183613c5f836002615879565b600c8110613c6f57613c6f6153ee565b6020020152838260028110613c8657613c866153ee565b6020020151516001602002015183613c9f836003615879565b600c8110613caf57613caf6153ee565b6020020152838260028110613cc657613cc66153ee565b602002015160200151600060028110613ce157613ce16153ee565b602002015183613cf2836004615879565b600c8110613d0257613d026153ee565b6020020152838260028110613d1957613d196153ee565b602002015160200151600160028110613d3457613d346153ee565b602002015183613d45836005615879565b600c8110613d5557613d556153ee565b60200201525080613d6581615603565b915050613bab565b50613d76614570565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b6033546001600160a01b03163314612d5c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610835565b600080613e00846142bb565b9050808360ff166001901b11613e7e5760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610835565b90505b92915050565b6000805b8215613e8157613e9c600184615849565b9092169180613eaa81615b22565b915050613e8b565b60408051808201909152600080825260208201526102008261ffff1610613f0e5760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610835565b8161ffff1660011415613f22575081613e81565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1610613f8b57600161ffff871660ff83161c81161415613f6e57613f6b8484613992565b93505b613f788384613992565b92506201fffe600192831b169101613f3e565b509195945050505050565b60408051808201909152600080825260208201528151158015613fbb57506020820151155b15613fd9575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020615b45833981519152846020015161400c9190615404565b61402490600080516020615b45833981519152615849565b905292915050565b919050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6065546001600160a01b03161580156140a457506001600160a01b03821615155b6141265760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610835565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a261416982613804565b5050565b606060008061417b84613e87565b61ffff166001600160401b0381111561419657614196614615565b6040519080825280601f01601f1916602001820160405280156141c0576020820181803683370190505b5090506000805b8251821080156141d8575061010081105b1561422f576001811b93508584161561421f578060f81b838381518110614201576142016153ee565b60200101906001600160f81b031916908160001a9053508160010191505b61422881615603565b90506141c7565b5090949350505050565b60008080600080516020615b458339815191526003600080516020615b4583398151915286600080516020615b458339815191528889090908905060006142af827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020615b45833981519152614448565b91959194509092505050565b6000610100825111156143445760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610835565b815161435257506000919050565b60008083600081518110614368576143686153ee565b0160200151600160f89190911c81901b92505b845181101561443f57848181518110614396576143966153ee565b0160200151600160f89190911c1b915082821161442b5760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610835565b9181179161443881615603565b905061437b565b50909392505050565b600080614453614570565b61445b61458e565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa925082801561394a5750826144e55760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610835565b505195945050505050565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b604051806040016040528061453f6145ac565b815260200161454c6145ac565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b038116811461084757600080fd5b6000602082840312156145f157600080fd5b8135613e7e816145ca565b60006020828403121561460e57600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171561464d5761464d614615565b60405290565b60405160c081016001600160401b038111828210171561464d5761464d614615565b60405161010081016001600160401b038111828210171561464d5761464d614615565b604051601f8201601f191681016001600160401b03811182821017156146c0576146c0614615565b604052919050565b6000604082840312156146da57600080fd5b6146e261462b565b9050813581526020820135602082015292915050565b600082601f83011261470957600080fd5b604051604081018181106001600160401b038211171561472b5761472b614615565b806040525080604084018581111561474257600080fd5b845b81811015613f8b578035835260209283019201614744565b60006080828403121561476e57600080fd5b61477661462b565b905061478283836146f8565b815261479183604084016146f8565b602082015292915050565b60008060008061012085870312156147b357600080fd5b843593506147c486602087016146c8565b92506147d3866060870161475c565b91506147e28660e087016146c8565b905092959194509250565b63ffffffff8116811461084757600080fd5b803561402c816147ed565b60006020828403121561481c57600080fd5b8135613e7e816147ed565b60ff8116811461084757600080fd5b600082601f83011261484757600080fd5b81356001600160401b0381111561486057614860614615565b614873601f8201601f1916602001614698565b81815284602083860101111561488857600080fd5b816020850160208301376000918101602001919091529392505050565b803561402c816145ca565b600080604083850312156148c357600080fd5b82356148ce816147ed565b915060208301356001600160401b03808211156148ea57600080fd5b9084019060c082870312156148fe57600080fd5b614906614653565b823561491181614827565b81526020830135614921816147ed565b602082015260408301358281111561493857600080fd5b61494488828601614836565b6040830152506060830135915061495a826147ed565b81606082015261496c608084016148a5565b608082015260a083013560a08201528093505050509250929050565b60008060006060848603121561499d57600080fd5b83356149a8816145ca565b925060208401356001600160401b038111156149c357600080fd5b6149cf86828701614836565b92505060408401356149e0816147ed565b809150509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015614a81578385038a52825180518087529087019087870190845b81811015614a6c57835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101614a28565b50509a87019a95505091850191600101614a0a565b509298975050505050505050565b602081526000614aa260208301846149eb565b9392505050565b801515811461084757600080fd5b600060208284031215614ac957600080fd5b8135613e7e81614aa9565b60008083601f840112614ae657600080fd5b5081356001600160401b03811115614afd57600080fd5b602083019150836020828501011115614b1557600080fd5b9250929050565b60008060008060008060808789031215614b3557600080fd5b8635614b40816145ca565b95506020870135614b50816147ed565b945060408701356001600160401b0380821115614b6c57600080fd5b614b788a838b01614ad4565b90965094506060890135915080821115614b9157600080fd5b818901915089601f830112614ba557600080fd5b813581811115614bb457600080fd5b8a60208260051b8501011115614bc957600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b83811015614c1557815163ffffffff1687529582019590820190600101614bf3565b509495945050505050565b600060208083528351608082850152614c3c60a0850182614bdf565b905081850151601f1980868403016040870152614c598383614bdf565b92506040870151915080868403016060870152614c768383614bdf565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b82811015614ccd5784878303018452614cbb828751614bdf565b95880195938801939150600101614ca1565b509998505050505050505050565b600060208284031215614ced57600080fd5b8135613e7e81614827565b60006001600160401b03821115614d1157614d11614615565b5060051b60200190565b600080600060608486031215614d3057600080fd5b8335614d3b816145ca565b92506020848101356001600160401b03811115614d5757600080fd5b8501601f81018713614d6857600080fd5b8035614d7b614d7682614cf8565b614698565b81815260059190911b82018301908381019089831115614d9a57600080fd5b928401925b82841015614db857833582529284019290840190614d9f565b8096505050505050614dcc604085016147ff565b90509250925092565b6020808252825182820181905260009190848201906040850190845b81811015614e0d57835183529284019291840191600101614df1565b50909695505050505050565b600082601f830112614e2a57600080fd5b81356020614e3a614d7683614cf8565b82815260059290921b84018101918181019086841115614e5957600080fd5b8286015b84811015614e7d578035614e70816147ed565b8352918301918301614e5d565b509695505050505050565b600082601f830112614e9957600080fd5b81356020614ea9614d7683614cf8565b82815260069290921b84018101918181019086841115614ec857600080fd5b8286015b84811015614e7d57614ede88826146c8565b835291830191604001614ecc565b600082601f830112614efd57600080fd5b81356020614f0d614d7683614cf8565b82815260059290921b84018101918181019086841115614f2c57600080fd5b8286015b84811015614e7d5780356001600160401b03811115614f4f5760008081fd5b614f5d8986838b0101614e19565b845250918301918301614f30565b60006101808284031215614f7e57600080fd5b614f86614675565b905081356001600160401b0380821115614f9f57600080fd5b614fab85838601614e19565b83526020840135915080821115614fc157600080fd5b614fcd85838601614e88565b60208401526040840135915080821115614fe657600080fd5b614ff285838601614e88565b6040840152615004856060860161475c565b60608401526150168560e086016146c8565b608084015261012084013591508082111561503057600080fd5b61503c85838601614e19565b60a084015261014084013591508082111561505657600080fd5b61506285838601614e19565b60c084015261016084013591508082111561507c57600080fd5b5061508984828501614eec565b60e08301525092915050565b6000806000806000608086880312156150ad57600080fd5b8535945060208601356001600160401b03808211156150cb57600080fd5b6150d789838a01614ad4565b9096509450604088013591506150ec826147ed565b9092506060870135908082111561510257600080fd5b5061510f88828901614f6b565b9150509295509295909350565b600081518084526020808501945080840160005b83811015614c155781516001600160601b031687529582019590820190600101615130565b6040815260008351604080840152615170608084018261511c565b90506020850151603f1984830301606085015261518d828261511c565b925050508260208301529392505050565b600080600083850360a08112156151b457600080fd5b84356001600160401b03808211156151cb57600080fd5b9086019060c082890312156151df57600080fd5b8195506060601f19840112156151f457600080fd5b602087019450608087013592508083111561520e57600080fd5b505061521c86828701614f6b565b9150509250925092565b6000806000806060858703121561523c57600080fd5b843561524781614827565b93506020850135615257816147ed565b925060408501356001600160401b0381111561527257600080fd5b61527e87828801614ad4565b95989497509550505050565b60008060006060848603121561529f57600080fd5b83356152aa816145ca565b925060208401356152ba816145ca565b915060408401356149e0816145ca565b6000806000606084860312156152df57600080fd5b83356152ea816145ca565b92506020840135915060408401356149e0816147ed565b82815260406020820152600061531a60408301846149eb565b949350505050565b60006020828403121561533457600080fd5b8151613e7e816145ca565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b60006020828403121561539b57600080fd5b8151613e7e81614aa9565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b60008261542157634e487b7160e01b600052601260045260246000fd5b500690565b6000602080835260ff8451168184015263ffffffff81850151166040840152604084015160c0606085015280518060e086015260005b81811015615479578281018401518682016101000152830161545c565b8181111561548c57600061010083880101525b50606086015163ffffffff81166080870152925060808601516001600160a01b03811660a0870152925060a0959095015160c0850152505050601f91909101601f1916016101000190565b6020808252603d908201527f737570706c696564207461736b20646f6573206e6f74206d617463682074686560408201527f206f6e65207265636f7264656420696e2074686520636f6e7472616374000000606082015260800190565b6000602080838503121561554757600080fd5b82516001600160401b0381111561555d57600080fd5b8301601f8101851361556e57600080fd5b805161557c614d7682614cf8565b81815260059190911b8201830190838101908783111561559b57600080fd5b928401925b828410156155b9578351825292840192908401906155a0565b979650505050505050565b6000602082840312156155d657600080fd5b81516001600160601b0381168114613e7e57600080fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415615617576156176155ed565b5060010190565b63ffffffff84168152604060208201819052810182905260006001600160fb1b0383111561564b57600080fd5b8260051b8085606085013760009201606001918252509392505050565b6000602080838503121561567b57600080fd5b82516001600160401b0381111561569157600080fd5b8301601f810185136156a257600080fd5b80516156b0614d7682614cf8565b81815260059190911b820183019083810190878311156156cf57600080fd5b928401925b828410156155b95783516156e7816147ed565b825292840192908401906156d4565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff8416815260406020820152600061573f6040830184866156f6565b95945050505050565b60006020828403121561575a57600080fd5b81516001600160c01b0381168114613e7e57600080fd5b60006020828403121561578357600080fd5b8151613e7e816147ed565b600060ff821660ff8114156157a5576157a56155ed565b60010192915050565b6040815260006157c26040830185876156f6565b905063ffffffff83166020830152949350505050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b8181101561581f57845183529383019391830191600101615803565b5090979650505050505050565b60006020828403121561583e57600080fd5b8151613e7e81614827565b60008282101561585b5761585b6155ed565b500390565b60006020828403121561587257600080fd5b5051919050565b6000821982111561588c5761588c6155ed565b500190565b6000602082840312156158a357600080fd5b815167ffffffffffffffff1981168114613e7e57600080fd5b60006001600160601b03838116908316818110156158dc576158dc6155ed565b039392505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b8381101561591f57815185529382019390820190600101615903565b5092979650505050505050565b6000808335601e1984360301811261594357600080fd5b8301803591506001600160401b0382111561595d57600080fd5b602001915036819003821315614b1557600080fd5b602081526000823561598381614827565b60ff8116602084015250602083013561599b816147ed565b63ffffffff81166040840152506040830135601e198436030181126159bf57600080fd5b830180356001600160401b038111156159d757600080fd5b8036038513156159e657600080fd5b60c060608501526159fe60e0850182602085016156f6565b915050615a0d606085016147ff565b63ffffffff8116608085015250615a26608085016148a5565b6001600160a01b03811660a08501525060a084013560c08401528091505092915050565b600063ffffffff808316818516808303821115615a6957615a696155ed565b01949350505050565b8035615a7d816147ed565b63ffffffff16825260208181013590830152604090810135910152565b60608101613e818284615a72565b60006001600160601b0380831681851681830481118215151615615ace57615ace6155ed565b02949350505050565b6000816000190483118215151615615af157615af16155ed565b500290565b60a08101615b048285615a72565b63ffffffff8351166060830152602083015160808301529392505050565b600061ffff80831681811415615b3a57615b3a6155ed565b600101939250505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a26469706673582212205499d9a4fbc6754d5c69b347e6cf320cb615426a7b55b1bfcf974307bca29ccc64736f6c634300080c0033",
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

// CreateNewTask is a paid mutator transaction binding the contract method 0xa7a69b05.
//
// Solidity: function createNewTask(uint8 taskType, uint32 quorumThresholdPercentage, bytes quorumNumbers) payable returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactor) CreateNewTask(opts *bind.TransactOpts, taskType uint8, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.contract.Transact(opts, "createNewTask", taskType, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xa7a69b05.
//
// Solidity: function createNewTask(uint8 taskType, uint32 quorumThresholdPercentage, bytes quorumNumbers) payable returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) CreateNewTask(taskType uint8, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.CreateNewTask(&_ContractOpenOracleTaskManager.TransactOpts, taskType, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xa7a69b05.
//
// Solidity: function createNewTask(uint8 taskType, uint32 quorumThresholdPercentage, bytes quorumNumbers) payable returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactorSession) CreateNewTask(taskType uint8, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.CreateNewTask(&_ContractOpenOracleTaskManager.TransactOpts, taskType, quorumThresholdPercentage, quorumNumbers)
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

// RespondToTask is a paid mutator transaction binding the contract method 0x88fee238.
//
// Solidity: function respondToTask((uint8,uint32,bytes,uint32,address,uint256) task, (uint32,uint256,uint256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactor) RespondToTask(opts *bind.TransactOpts, task IOpenOracleTaskManagerTask, taskResponse IOpenOracleTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.contract.Transact(opts, "respondToTask", task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x88fee238.
//
// Solidity: function respondToTask((uint8,uint32,bytes,uint32,address,uint256) task, (uint32,uint256,uint256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) RespondToTask(task IOpenOracleTaskManagerTask, taskResponse IOpenOracleTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.RespondToTask(&_ContractOpenOracleTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x88fee238.
//
// Solidity: function respondToTask((uint8,uint32,bytes,uint32,address,uint256) task, (uint32,uint256,uint256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactorSession) RespondToTask(task IOpenOracleTaskManagerTask, taskResponse IOpenOracleTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.RespondToTask(&_ContractOpenOracleTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
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

// WithdrawTaskFunds is a paid mutator transaction binding the contract method 0x352f52cb.
//
// Solidity: function withdrawTaskFunds(uint32 taskNum, (uint8,uint32,bytes,uint32,address,uint256) task) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerTransactor) WithdrawTaskFunds(opts *bind.TransactOpts, taskNum uint32, task IOpenOracleTaskManagerTask) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.contract.Transact(opts, "withdrawTaskFunds", taskNum, task)
}

// WithdrawTaskFunds is a paid mutator transaction binding the contract method 0x352f52cb.
//
// Solidity: function withdrawTaskFunds(uint32 taskNum, (uint8,uint32,bytes,uint32,address,uint256) task) returns()
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerSession) WithdrawTaskFunds(taskNum uint32, task IOpenOracleTaskManagerTask) (*types.Transaction, error) {
	return _ContractOpenOracleTaskManager.Contract.WithdrawTaskFunds(&_ContractOpenOracleTaskManager.TransactOpts, taskNum, task)
}

// WithdrawTaskFunds is a paid mutator transaction binding the contract method 0x352f52cb.
//
// Solidity: function withdrawTaskFunds(uint32 taskNum, (uint8,uint32,bytes,uint32,address,uint256) task) returns()
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

// FilterNewTaskCreated is a free log retrieval operation binding the contract event 0x767aa0f97f690f2aa73dc95d49893c4881872e9f1bd64e358a57a6e349f1dcd3.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint8,uint32,bytes,uint32,address,uint256) task)
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

// WatchNewTaskCreated is a free log subscription operation binding the contract event 0x767aa0f97f690f2aa73dc95d49893c4881872e9f1bd64e358a57a6e349f1dcd3.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint8,uint32,bytes,uint32,address,uint256) task)
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

// ParseNewTaskCreated is a log parse operation binding the contract event 0x767aa0f97f690f2aa73dc95d49893c4881872e9f1bd64e358a57a6e349f1dcd3.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint8,uint32,bytes,uint32,address,uint256) task)
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
	TaskResponse         IOpenOracleTaskManagerTaskResponse
	TaskResponseMetadata IOpenOracleTaskManagerTaskResponseMetadata
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTaskResponded is a free log retrieval operation binding the contract event 0xc7b8513936e453ed22d56dc5c765bceee58227d50ca6ae9ac750430339cd900b.
//
// Solidity: event TaskResponded((uint32,uint256,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractOpenOracleTaskManager *ContractOpenOracleTaskManagerFilterer) FilterTaskResponded(opts *bind.FilterOpts) (*ContractOpenOracleTaskManagerTaskRespondedIterator, error) {

	logs, sub, err := _ContractOpenOracleTaskManager.contract.FilterLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleTaskManagerTaskRespondedIterator{contract: _ContractOpenOracleTaskManager.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0xc7b8513936e453ed22d56dc5c765bceee58227d50ca6ae9ac750430339cd900b.
//
// Solidity: event TaskResponded((uint32,uint256,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata)
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

// ParseTaskResponded is a log parse operation binding the contract event 0xc7b8513936e453ed22d56dc5c765bceee58227d50ca6ae9ac750430339cd900b.
//
// Solidity: event TaskResponded((uint32,uint256,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata)
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

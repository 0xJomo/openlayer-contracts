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
	Price              *big.Int
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"taskType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestTaskNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIOpenOracleTaskManager.Task\",\"components\":[{\"name\":\"taskType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"creationFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIOpenOracleTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeStamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskCreationFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTaskCreationFee\",\"inputs\":[{\"name\":\"_newFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawTaskFunds\",\"inputs\":[{\"name\":\"taskNum\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIOpenOracleTaskManager.Task\",\"components\":[{\"name\":\"taskType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"creationFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FundsWithdrawn\",\"inputs\":[{\"name\":\"taskNum\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"creator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOpenOracleTaskManager.Task\",\"components\":[{\"name\":\"taskType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"creationFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOpenOracleTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeStamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOpenOracleTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x610120604052655af3107a40006098553480156200001c57600080fd5b5060405162005afd38038062005afd8339810160408190526200003f9162000201565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000099573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000bf919062000248565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000117573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013d919062000248565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000197573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001bd919062000248565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff1661010052506200026f565b6001600160a01b0381168114620001fe57600080fd5b50565b600080604083850312156200021557600080fd5b82516200022281620001e8565b602084015190925063ffffffff811681146200023d57600080fd5b809150509250929050565b6000602082840312156200025b57600080fd5b81516200026881620001e8565b9392505050565b60805160a05160c05160e05161010051615803620002fa60003960008181610299015281816107120152612c4c0152600081816106c101526122660152600081816104d101526124500152600081816105050152818161262601526127e8015260008181610539015281816112c501528181611f44015281816120dc015261230a01526158036000f3fe6080604052600436106102045760003560e01c80635df45946116101185780638da5cb5b116100a0578063cefdc1d41161006f578063cefdc1d414610681578063df5cf723146106af578063f2fde38b146106e3578063f5c9899d14610703578063fabc1cbc1461073657600080fd5b80638da5cb5b14610616578063a7a69b0514610634578063b98d090814610647578063c0c53b8b1461066157600080fd5b8063715018a6116100e7578063715018a61461058957806372d18e8d1461059e578063886f1195146105b957806388fee238146105d95780638b00ce7c146105f957600080fd5b80635df45946146104bf57806368304835146104f35780636d14a987146105275780636efb46361461055b57600080fd5b80633563b0d11161019b5780634f739f741161016a5780634f739f741461040857806352e9408b14610435578063595c6a67146104555780635ac86ab71461046a5780635c975abb146104aa57600080fd5b80633563b0d1146103905780633ccfd60b146103bd578063416c7e5e146103d25780634d075ce7146103f257600080fd5b8063245a7bfc116101d7578063245a7bfc146102d05780632cb223d5146103085780632d89f6fc14610343578063352f52cb1461037057600080fd5b806310d67a2f14610209578063136439dd1461022b578063171f1d5b1461024b5780631ad4318914610287575b600080fd5b34801561021557600080fd5b506102296102243660046142eb565b610756565b005b34801561023757600080fd5b50610229610246366004614308565b610812565b34801561025757600080fd5b5061026b6102663660046144a8565b610951565b6040805192151583529015156020830152015b60405180910390f35b34801561029357600080fd5b506102bb7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff909116815260200161027e565b3480156102dc57600080fd5b50609c546102f0906001600160a01b031681565b6040516001600160a01b03909116815260200161027e565b34801561031457600080fd5b50610335610323366004614516565b609b6020526000908152604090205481565b60405190815260200161027e565b34801561034f57600080fd5b5061033561035e366004614516565b609a6020526000908152604090205481565b34801561037c57600080fd5b5061022961038b3660046145bc565b610adb565b34801561039c57600080fd5b506103b06103ab366004614694565b610df9565b60405161027e919061479b565b3480156103c957600080fd5b5061022961128f565b3480156103de57600080fd5b506102296103ed3660046147c3565b6112c3565b3480156103fe57600080fd5b5061033560985481565b34801561041457600080fd5b50610428610423366004614828565b611438565b60405161027e919061492c565b34801561044157600080fd5b50610229610450366004614308565b611abc565b34801561046157600080fd5b50610229611ac9565b34801561047657600080fd5b5061049a6104853660046149e7565b606654600160ff9092169190911b9081161490565b604051901515815260200161027e565b3480156104b657600080fd5b50606654610335565b3480156104cb57600080fd5b506102f07f000000000000000000000000000000000000000000000000000000000000000081565b3480156104ff57600080fd5b506102f07f000000000000000000000000000000000000000000000000000000000000000081565b34801561053357600080fd5b506102f07f000000000000000000000000000000000000000000000000000000000000000081565b34801561056757600080fd5b5061057b610576366004614ca8565b611b90565b60405161027e929190614d68565b34801561059557600080fd5b50610229612a9d565b3480156105aa57600080fd5b5060995463ffffffff166102bb565b3480156105c557600080fd5b506065546102f0906001600160a01b031681565b3480156105e557600080fd5b506102296105f4366004614db1565b612ab1565b34801561060557600080fd5b506099546102bb9063ffffffff1681565b34801561062257600080fd5b506033546001600160a01b03166102f0565b610229610642366004614e39565b612eda565b34801561065357600080fd5b5060975461049a9060ff1681565b34801561066d57600080fd5b5061022961067c366004614e9d565b6130b8565b34801561068d57600080fd5b506106a161069c366004614edd565b6131f3565b60405161027e929190614f14565b3480156106bb57600080fd5b506102f07f000000000000000000000000000000000000000000000000000000000000000081565b3480156106ef57600080fd5b506102296106fe3660046142eb565b613385565b34801561070f57600080fd5b507f00000000000000000000000000000000000000000000000000000000000000006102bb565b34801561074257600080fd5b50610229610751366004614308565b6133fb565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156107a9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107cd9190614f35565b6001600160a01b0316336001600160a01b0316146108065760405162461bcd60e51b81526004016107fd90614f52565b60405180910390fd5b61080f81613557565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa15801561085a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061087e9190614f9c565b61089a5760405162461bcd60e51b81526004016107fd90614fb9565b606654818116146109135760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c697479000000000000000060648201526084016107fd565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018787600001518860200151886000015160006002811061099957610999615001565b60200201518951600160200201518a602001516000600281106109be576109be615001565b60200201518b602001516001600281106109da576109da615001565b602090810291909101518c518d830151604051610a379a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c610a5a9190615017565b9050610acd610a73610a6c888461364e565b86906136e5565b610a7b613779565b610ac3610ab485610aae604080518082018252600080825260209182015281518083019092526001825260029082015290565b9061364e565b610abd8c613839565b906136e5565b886201d4c06138c9565b909890975095505050505050565b6099548290829063ffffffff9081169083161115610b345760405162461bcd60e51b81526020600482015260166024820152757461736b206e756d626572206e6f742065786973747360501b60448201526064016107fd565b63ffffffff82166000908152609a6020908152604091829020549151610b5c91849101615065565b6040516020818303038152906040528051906020012014610b8f5760405162461bcd60e51b81526004016107fd906150f1565b60808101516001600160a01b03163314610c055760405162461bcd60e51b815260206004820152603160248201527f4f6e6c7920746865207461736b2063726561746f722063616e20706572666f7260448201527036903a3434b99037b832b930ba34b7b71760791b60648201526084016107fd565b63ffffffff84166000908152609b6020526040902054610c7b5760405162461bcd60e51b815260206004820152602b60248201527f43616e6e6f742077697468647261772066756e647320666f72206120636f6d7060448201526a3632ba32b2103a30b9b59760a91b60648201526084016107fd565b60008360a0015111610cc75760405162461bcd60e51b8152602060048201526015602482015274273790333ab73239903a37903bb4ba34323930bb9760591b60448201526064016107fd565b60a0830180516000909152604051610ce3908590602001615065565b60408051601f19818403018152828252805160209182012063ffffffff89166000908152609a90925291812091909155608086015190916001600160a01b039091169083908381818185875af1925050503d8060008114610d60576040519150601f19603f3d011682016040523d82523d6000602084013e610d65565b606091505b5050905080610dab5760405162461bcd60e51b81526020600482015260126024820152712bb4ba34323930bbb0b6103330b4b632b21760711b60448201526064016107fd565b6040805163ffffffff881681523360208201529081018390527ff440aec6b52895984d061d622e6edeba6210f7c3e059be920663140c084560d79060600160405180910390a1505050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e3b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e5f9190614f35565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ea1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ec59190614f35565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f07573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f2b9190614f35565b9050600086516001600160401b03811115610f4857610f48614321565b604051908082528060200260200182016040528015610f7b57816020015b6060815260200190600190039081610f665790505b50905060005b8751811015611283576000888281518110610f9e57610f9e615001565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610fff573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611027919081019061514e565b905080516001600160401b0381111561104257611042614321565b60405190808252806020026020018201604052801561108d57816020015b60408051606081018252600080825260208083018290529282015282526000199092019101816110605790505b508484815181106110a0576110a0615001565b602002602001018190525060005b815181101561126d576040518060600160405280876001600160a01b03166347b314e88585815181106110e3576110e3615001565b60200260200101516040518263ffffffff1660e01b815260040161110991815260200190565b602060405180830381865afa158015611126573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061114a9190614f35565b6001600160a01b0316815260200183838151811061116a5761116a615001565b60200260200101518152602001896001600160a01b031663fa28c62785858151811061119857611198615001565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa1580156111f4573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061121891906151de565b6001600160601b031681525085858151811061123657611236615001565b6020026020010151828151811061124f5761124f615001565b602002602001018190525080806112659061521d565b9150506110ae565b505050808061127b9061521d565b915050610f81565b50979650505050505050565b611297613aed565b60405133904780156108fc02916000818181858888f1935050505015801561080f573d6000803e3d6000fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611321573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113459190614f35565b6001600160a01b0316336001600160a01b0316146113f15760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a4016107fd565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b6114636040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156114a3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114c79190614f35565b90506114f46040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e90611524908b9089908990600401615238565b600060405180830381865afa158015611541573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526115699190810190615282565b81526040516340e03a8160e11b81526001600160a01b038316906381c075029061159b908b908b908b90600401615339565b600060405180830381865afa1580156115b8573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526115e09190810190615282565b6040820152856001600160401b038111156115fd576115fd614321565b60405190808252806020026020018201604052801561163057816020015b606081526020019060019003908161161b5790505b50606082015260005b60ff81168711156119cd576000856001600160401b0381111561165e5761165e614321565b604051908082528060200260200182016040528015611687578160200160208202803683370190505b5083606001518360ff16815181106116a1576116a1615001565b602002602001018190525060005b868110156118cd5760008c6001600160a01b03166304ec63518a8a858181106116da576116da615001565b905060200201358e886000015186815181106116f8576116f8615001565b60200260200101516040518463ffffffff1660e01b81526004016117359392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611752573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117769190615362565b90508a8a8560ff1681811061178d5761178d615001565b6001600160c01b03841692013560f81c9190911c6001908116141590506118ba57856001600160a01b031663dd9846b98a8a858181106117cf576117cf615001565b905060200201358d8d8860ff168181106117eb576117eb615001565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015611841573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611865919061538b565b85606001518560ff168151811061187e5761187e615001565b6020026020010151848151811061189757611897615001565b63ffffffff90921660209283029190910190910152826118b68161521d565b9350505b50806118c58161521d565b9150506116af565b506000816001600160401b038111156118e8576118e8614321565b604051908082528060200260200182016040528015611911578160200160208202803683370190505b50905060005b828110156119925784606001518460ff168151811061193857611938615001565b6020026020010151818151811061195157611951615001565b602002602001015182828151811061196b5761196b615001565b63ffffffff909216602092830291909101909101528061198a8161521d565b915050611917565b508084606001518460ff16815181106119ad576119ad615001565b6020026020010181905250505080806119c5906153a8565b915050611639565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611a0e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a329190614f35565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90611a65908b908b908e906004016153c8565b600060405180830381865afa158015611a82573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611aaa9190810190615282565b60208301525098975050505050505050565b611ac4613aed565b609855565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015611b11573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b359190614f9c565b611b515760405162461bcd60e51b81526004016107fd90614fb9565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b6040805180820190915260608082526020820152600084611c075760405162461bcd60e51b815260206004820152603760248201526000805160206157ae83398151915260448201527f7265733a20656d7074792071756f72756d20696e70757400000000000000000060648201526084016107fd565b60408301515185148015611c1f575060a08301515185145b8015611c2f575060c08301515185145b8015611c3f575060e08301515185145b611ca95760405162461bcd60e51b815260206004820152604160248201526000805160206157ae83398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a4016107fd565b82515160208401515114611d215760405162461bcd60e51b8152602060048201526044602482018190526000805160206157ae833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a4016107fd565b4363ffffffff168463ffffffff161115611d915760405162461bcd60e51b815260206004820152603c60248201526000805160206157ae83398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b0000000060648201526084016107fd565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b03811115611dd257611dd2614321565b604051908082528060200260200182016040528015611dfb578160200160208202803683370190505b506020820152866001600160401b03811115611e1957611e19614321565b604051908082528060200260200182016040528015611e42578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b03811115611e7657611e76614321565b604051908082528060200260200182016040528015611e9f578160200160208202803683370190505b5081526020860151516001600160401b03811115611ebf57611ebf614321565b604051908082528060200260200182016040528015611ee8578160200160208202803683370190505b5081602001819052506000611fba8a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa158015611f91573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fb591906153f2565b613b47565b905060005b8760200151518110156122555761200488602001518281518110611fe557611fe5615001565b6020026020010151805160009081526020918201519091526040902090565b8360200151828151811061201a5761201a615001565b602090810291909101015280156120da57602083015161203b60018361540f565b8151811061204b5761204b615001565b602002602001015160001c8360200151828151811061206c5761206c615001565b602002602001015160001c116120da576040805162461bcd60e51b81526020600482015260248101919091526000805160206157ae83398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f7274656460648201526084016107fd565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec63518460200151838151811061211f5761211f615001565b60200260200101518b8b60000151858151811061213e5761213e615001565b60200260200101516040518463ffffffff1660e01b815260040161217b9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015612198573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121bc9190615362565b6001600160c01b0316836000015182815181106121db576121db615001565b602002602001018181525050612241610a6c612215848660000151858151811061220757612207615001565b602002602001015116613c02565b8a60200151848151811061222b5761222b615001565b6020026020010151613c2d90919063ffffffff16565b94508061224d8161521d565b915050611fbf565b505061226083613d11565b925060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166350f73e7c6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156122c2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122e69190615426565b60975490915060ff1660005b8a81101561296c57811561244e578963ffffffff16837f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f8681811061234957612349615001565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015612389573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123ad9190615426565b6123b7919061543f565b101561244e5760405162461bcd60e51b815260206004820152606660248201526000805160206157ae83398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c4016107fd565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d8481811061248f5761248f615001565b9050013560f81c60f81b60f81c8c8c60a0015185815181106124b3576124b3615001565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa15801561250f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125339190615457565b6001600160401b0319166125568a604001518381518110611fe557611fe5615001565b67ffffffffffffffff1916146125f25760405162461bcd60e51b815260206004820152606160248201526000805160206157ae83398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c4016107fd565b6126228960400151828151811061260b5761260b615001565b6020026020010151876136e590919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d8481811061266557612665615001565b9050013560f81c60f81b60f81c8c8c60c00151858151811061268957612689615001565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa1580156126e5573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061270991906151de565b8560200151828151811061271f5761271f615001565b6001600160601b0390921660209283029190910182015285015180518290811061274b5761274b615001565b60200260200101518560000151828151811061276957612769615001565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a6020015151811015612957576127e1866000015182815181106127b3576127b3615001565b60200260200101518f8f868181106127cd576127cd615001565b600192013560f81c9290921c811614919050565b15612945577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f8681811061282757612827615001565b9050013560f81c60f81b60f81c8e8960200151858151811061284b5761284b615001565b60200260200101518f60e00151888151811061286957612869615001565b6020026020010151878151811061288257612882615001565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa1580156128e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061290a91906151de565b875180518590811061291e5761291e615001565b602002602001018181516129329190615482565b6001600160601b03169052506001909101905b8061294f8161521d565b91505061278d565b505080806129649061521d565b9150506122f2565b5050506000806129868c868a606001518b60800151610951565b91509150816129f75760405162461bcd60e51b815260206004820152604360248201526000805160206157ae83398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a4016107fd565b80612a585760405162461bcd60e51b815260206004820152603960248201526000805160206157ae83398151915260448201527f7265733a207369676e617475726520697320696e76616c69640000000000000060648201526084016107fd565b50506000878260200151604051602001612a739291906154aa565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b612aa5613aed565b612aaf6000613dac565b565b609c546001600160a01b03163314612b0b5760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c657200000060448201526064016107fd565b6000612b1d6040850160208601614516565b9050366000612b2f60408701876154f2565b90925090506000612b466080880160608901614516565b9050609a6000612b596020890189614516565b63ffffffff1663ffffffff1681526020019081526020016000205487604051602001612b859190615538565b6040516020818303038152906040528051906020012014612bb85760405162461bcd60e51b81526004016107fd906150f1565b6000609b81612bca60208a018a614516565b63ffffffff1663ffffffff1681526020019081526020016000205414612c475760405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b60648201526084016107fd565b612c717f000000000000000000000000000000000000000000000000000000000000000085615610565b63ffffffff164363ffffffff161115612ce25760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320726573706f6e64656420746f207468652060448201526c7461736b20746f6f206c61746560981b60648201526084016107fd565b600086604051602001612cf59190615660565b604051602081830303815290604052805190602001209050600080612d1d8387878a8c611b90565b9150915060005b85811015612e1c578460ff1683602001518281518110612d4657612d46615001565b6020026020010151612d58919061566e565b6001600160601b0316606484600001518381518110612d7957612d79615001565b60200260200101516001600160601b0316612d94919061569d565b1015612e0a576040805162461bcd60e51b81526020600482015260248101919091527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d60648201526084016107fd565b80612e148161521d565b915050612d24565b5060408051808201825263ffffffff43168152602080820184905291519091612e49918c918491016156bc565b60405160208183030381529060405280519060200120609b60008c6000016020810190612e769190614516565b63ffffffff1663ffffffff168152602001908152602001600020819055507fc7b8513936e453ed22d56dc5c765bceee58227d50ca6ae9ac750430339cd900b8a82604051612ec59291906156bc565b60405180910390a15050505050505050505050565b600180609854612eea919061569d565b341015612f395760405162461bcd60e51b815260206004820152601f60248201527f4372656174696e672061207461736b2072657175697265732061206665652e0060448201526064016107fd565b6040805160c081018252606081830181905260006080830181905260a083015260ff8816825263ffffffff438116602080850191909152908816918301919091528251601f860182900482028101820190935284835290919085908590819084018382808284376000920191909152505050506040808301919091523360808301523460a083015251612fd0908290602001615065565b60408051601f1981840301815282825280516020918201206099805463ffffffff9081166000908152609a90945293909220555416907f767aa0f97f690f2aa73dc95d49893c4881872e9f1bd64e358a57a6e349f1dcd390613033908490615065565b60405180910390a260995461304f9063ffffffff166001615610565b6099805463ffffffff191663ffffffff929092169190911790556098543411156130b05760985433906108fc90613086903461540f565b6040518115909202916000818181858888f193505050501580156130ae573d6000803e3d6000fd5b505b505050505050565b600054610100900460ff16158080156130d85750600054600160ff909116105b806130f25750303b1580156130f2575060005460ff166001145b6131555760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016107fd565b6000805460ff191660011790558015613178576000805461ff0019166101001790555b613183846000613dfe565b61318c83613dac565b609c80546001600160a01b0319166001600160a01b03841617905580156131ed576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b604080516001808252818301909252600091606091839160208083019080368337019050509050848160008151811061322e5761322e615001565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e9061326a90889086906004016156e8565b600060405180830381865afa158015613287573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526132af9190810190615282565b6000815181106132c1576132c1615001565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa15801561332d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133519190615362565b6001600160c01b03169050600061336782613ee8565b9050816133758a838a610df9565b9550955050505050935093915050565b61338d613aed565b6001600160a01b0381166133f25760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016107fd565b61080f81613dac565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561344e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134729190614f35565b6001600160a01b0316336001600160a01b0316146134a25760405162461bcd60e51b81526004016107fd90614f52565b6066541981196066541916146135205760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c697479000000000000000060648201526084016107fd565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610946565b6001600160a01b0381166135e55760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a4016107fd565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b604080518082019091526000808252602082015261366a6141fc565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa905080801561369d5761369f565bfe5b50806136dd5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b60448201526064016107fd565b505092915050565b604080518082019091526000808252602082015261370161421a565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa905080801561369d5750806136dd5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b60448201526064016107fd565b613781614238565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b60408051808201909152600080825260208201526000808061386960008051602061578e83398151915286615017565b90505b61387581613f45565b909350915060008051602061578e8339815191528283098314156138af576040805180820190915290815260208101919091529392505050565b60008051602061578e83398151915260018208905061386c565b6040805180820182528681526020808201869052825180840190935286835282018490526000918291906138fb61425d565b60005b6002811015613ac057600061391482600661569d565b905084826002811061392857613928615001565b6020020151518361393a83600061543f565b600c811061394a5761394a615001565b602002015284826002811061396157613961615001565b60200201516020015183826001613978919061543f565b600c811061398857613988615001565b602002015283826002811061399f5761399f615001565b60200201515151836139b283600261543f565b600c81106139c2576139c2615001565b60200201528382600281106139d9576139d9615001565b60200201515160016020020151836139f283600361543f565b600c8110613a0257613a02615001565b6020020152838260028110613a1957613a19615001565b602002015160200151600060028110613a3457613a34615001565b602002015183613a4583600461543f565b600c8110613a5557613a55615001565b6020020152838260028110613a6c57613a6c615001565b602002015160200151600160028110613a8757613a87615001565b602002015183613a9883600561543f565b600c8110613aa857613aa8615001565b60200201525080613ab88161521d565b9150506138fe565b50613ac961427c565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b6033546001600160a01b03163314612aaf5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016107fd565b600080613b5384613fc7565b90508015613bf9578260ff168460018651613b6e919061540f565b81518110613b7e57613b7e615001565b016020015160f81c10613bf95760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c75650060648201526084016107fd565b90505b92915050565b6000805b8215613bfc57613c1760018461540f565b9092169180613c258161573c565b915050613c06565b60408051808201909152600080825260208201526102008261ffff1610613c895760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b60448201526064016107fd565b8161ffff1660011415613c9d575081613bfc565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1610613d0657600161ffff871660ff83161c81161415613ce957613ce684846136e5565b93505b613cf383846136e5565b92506201fffe600192831b169101613cb9565b509195945050505050565b60408051808201909152600080825260208201528151158015613d3657506020820151155b15613d54575050604080518082019091526000808252602082015290565b60405180604001604052808360000151815260200160008051602061578e8339815191528460200151613d879190615017565b613d9f9060008051602061578e83398151915261540f565b905292915050565b919050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6065546001600160a01b0316158015613e1f57506001600160a01b03821615155b613ea15760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a4016107fd565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2613ee482613557565b5050565b60606000805b610100811015613f3e576001811b915083821615613f2e57828160f81b604051602001613f1c92919061575e565b60405160208183030381529060405292505b613f378161521d565b9050613eee565b5050919050565b6000808060008051602061578e833981519152600360008051602061578e8339815191528660008051602061578e833981519152888909090890506000613fbb827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260008051602061578e833981519152614154565b91959194509092505050565b6000610100825111156140505760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a4016107fd565b815161405e57506000919050565b6000808360008151811061407457614074615001565b0160200151600160f89190911c81901b92505b845181101561414b578481815181106140a2576140a2615001565b0160200151600160f89190911c1b91508282116141375760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a4016107fd565b918117916141448161521d565b9050614087565b50909392505050565b60008061415f61427c565b61416761429a565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa925082801561369d5750826141f15760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c75726500000000000060448201526064016107fd565b505195945050505050565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b604051806040016040528061424b6142b8565b81526020016142586142b8565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b038116811461080f57600080fd5b6000602082840312156142fd57600080fd5b8135613bf9816142d6565b60006020828403121561431a57600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171561435957614359614321565b60405290565b60405160c081016001600160401b038111828210171561435957614359614321565b60405161010081016001600160401b038111828210171561435957614359614321565b604051601f8201601f191681016001600160401b03811182821017156143cc576143cc614321565b604052919050565b6000604082840312156143e657600080fd5b6143ee614337565b9050813581526020820135602082015292915050565b600082601f83011261441557600080fd5b604051604081018181106001600160401b038211171561443757614437614321565b806040525080604084018581111561444e57600080fd5b845b81811015613d06578035835260209283019201614450565b60006080828403121561447a57600080fd5b614482614337565b905061448e8383614404565b815261449d8360408401614404565b602082015292915050565b60008060008061012085870312156144bf57600080fd5b843593506144d086602087016143d4565b92506144df8660608701614468565b91506144ee8660e087016143d4565b905092959194509250565b63ffffffff8116811461080f57600080fd5b8035613da7816144f9565b60006020828403121561452857600080fd5b8135613bf9816144f9565b60ff8116811461080f57600080fd5b600082601f83011261455357600080fd5b81356001600160401b0381111561456c5761456c614321565b61457f601f8201601f19166020016143a4565b81815284602083860101111561459457600080fd5b816020850160208301376000918101602001919091529392505050565b8035613da7816142d6565b600080604083850312156145cf57600080fd5b82356145da816144f9565b915060208301356001600160401b03808211156145f657600080fd5b9084019060c0828703121561460a57600080fd5b61461261435f565b823561461d81614533565b8152602083013561462d816144f9565b602082015260408301358281111561464457600080fd5b61465088828601614542565b60408301525060608301359150614666826144f9565b816060820152614678608084016145b1565b608082015260a083013560a08201528093505050509250929050565b6000806000606084860312156146a957600080fd5b83356146b4816142d6565b925060208401356001600160401b038111156146cf57600080fd5b6146db86828701614542565b92505060408401356146ec816144f9565b809150509250925092565b600081518084526020808501808196508360051b810191508286016000805b8681101561478d578385038a52825180518087529087019087870190845b8181101561477857835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101614734565b50509a87019a95505091850191600101614716565b509298975050505050505050565b6020815260006147ae60208301846146f7565b9392505050565b801515811461080f57600080fd5b6000602082840312156147d557600080fd5b8135613bf9816147b5565b60008083601f8401126147f257600080fd5b5081356001600160401b0381111561480957600080fd5b60208301915083602082850101111561482157600080fd5b9250929050565b6000806000806000806080878903121561484157600080fd5b863561484c816142d6565b9550602087013561485c816144f9565b945060408701356001600160401b038082111561487857600080fd5b6148848a838b016147e0565b9096509450606089013591508082111561489d57600080fd5b818901915089601f8301126148b157600080fd5b8135818111156148c057600080fd5b8a60208260051b85010111156148d557600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b8381101561492157815163ffffffff16875295820195908201906001016148ff565b509495945050505050565b60006020808352835160808285015261494860a08501826148eb565b905081850151601f198086840301604087015261496583836148eb565b9250604087015191508086840301606087015261498283836148eb565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b828110156149d957848783030184526149c78287516148eb565b958801959388019391506001016149ad565b509998505050505050505050565b6000602082840312156149f957600080fd5b8135613bf981614533565b60006001600160401b03821115614a1d57614a1d614321565b5060051b60200190565b600082601f830112614a3857600080fd5b81356020614a4d614a4883614a04565b6143a4565b82815260059290921b84018101918181019086841115614a6c57600080fd5b8286015b84811015614a90578035614a83816144f9565b8352918301918301614a70565b509695505050505050565b600082601f830112614aac57600080fd5b81356020614abc614a4883614a04565b82815260069290921b84018101918181019086841115614adb57600080fd5b8286015b84811015614a9057614af188826143d4565b835291830191604001614adf565b600082601f830112614b1057600080fd5b81356020614b20614a4883614a04565b82815260059290921b84018101918181019086841115614b3f57600080fd5b8286015b84811015614a905780356001600160401b03811115614b625760008081fd5b614b708986838b0101614a27565b845250918301918301614b43565b60006101808284031215614b9157600080fd5b614b99614381565b905081356001600160401b0380821115614bb257600080fd5b614bbe85838601614a27565b83526020840135915080821115614bd457600080fd5b614be085838601614a9b565b60208401526040840135915080821115614bf957600080fd5b614c0585838601614a9b565b6040840152614c178560608601614468565b6060840152614c298560e086016143d4565b6080840152610120840135915080821115614c4357600080fd5b614c4f85838601614a27565b60a0840152610140840135915080821115614c6957600080fd5b614c7585838601614a27565b60c0840152610160840135915080821115614c8f57600080fd5b50614c9c84828501614aff565b60e08301525092915050565b600080600080600060808688031215614cc057600080fd5b8535945060208601356001600160401b0380821115614cde57600080fd5b614cea89838a016147e0565b909650945060408801359150614cff826144f9565b90925060608701359080821115614d1557600080fd5b50614d2288828901614b7e565b9150509295509295909350565b600081518084526020808501945080840160005b838110156149215781516001600160601b031687529582019590820190600101614d43565b6040815260008351604080840152614d836080840182614d2f565b90506020850151603f19848303016060850152614da08282614d2f565b925050508260208301529392505050565b600080600083850360a0811215614dc757600080fd5b84356001600160401b0380821115614dde57600080fd5b9086019060c08289031215614df257600080fd5b8195506060601f1984011215614e0757600080fd5b6020870194506080870135925080831115614e2157600080fd5b5050614e2f86828701614b7e565b9150509250925092565b60008060008060608587031215614e4f57600080fd5b8435614e5a81614533565b93506020850135614e6a816144f9565b925060408501356001600160401b03811115614e8557600080fd5b614e91878288016147e0565b95989497509550505050565b600080600060608486031215614eb257600080fd5b8335614ebd816142d6565b92506020840135614ecd816142d6565b915060408401356146ec816142d6565b600080600060608486031215614ef257600080fd5b8335614efd816142d6565b92506020840135915060408401356146ec816144f9565b828152604060208201526000614f2d60408301846146f7565b949350505050565b600060208284031215614f4757600080fd5b8151613bf9816142d6565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b600060208284031215614fae57600080fd5b8151613bf9816147b5565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b60008261503457634e487b7160e01b600052601260045260246000fd5b500690565b60005b8381101561505457818101518382015260200161503c565b838111156131ed5750506000910152565b6020815260ff825116602082015263ffffffff60208301511660408201526000604083015160c0606084015280518060e08501526101006150ac8282870160208601615039565b606086015163ffffffff166080868101919091528601516001600160a01b031660a0808701919091529095015160c0850152601f01601f191690920190920192915050565b6020808252603d908201527f737570706c696564207461736b20646f6573206e6f74206d617463682074686560408201527f206f6e65207265636f7264656420696e2074686520636f6e7472616374000000606082015260800190565b6000602080838503121561516157600080fd5b82516001600160401b0381111561517757600080fd5b8301601f8101851361518857600080fd5b8051615196614a4882614a04565b81815260059190911b820183019083810190878311156151b557600080fd5b928401925b828410156151d3578351825292840192908401906151ba565b979650505050505050565b6000602082840312156151f057600080fd5b81516001600160601b0381168114613bf957600080fd5b634e487b7160e01b600052601160045260246000fd5b600060001982141561523157615231615207565b5060010190565b63ffffffff84168152604060208201819052810182905260006001600160fb1b0383111561526557600080fd5b8260051b8085606085013760009201606001918252509392505050565b6000602080838503121561529557600080fd5b82516001600160401b038111156152ab57600080fd5b8301601f810185136152bc57600080fd5b80516152ca614a4882614a04565b81815260059190911b820183019083810190878311156152e957600080fd5b928401925b828410156151d3578351615301816144f9565b825292840192908401906152ee565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff84168152604060208201526000615359604083018486615310565b95945050505050565b60006020828403121561537457600080fd5b81516001600160c01b0381168114613bf957600080fd5b60006020828403121561539d57600080fd5b8151613bf9816144f9565b600060ff821660ff8114156153bf576153bf615207565b60010192915050565b6040815260006153dc604083018587615310565b905063ffffffff83166020830152949350505050565b60006020828403121561540457600080fd5b8151613bf981614533565b60008282101561542157615421615207565b500390565b60006020828403121561543857600080fd5b5051919050565b6000821982111561545257615452615207565b500190565b60006020828403121561546957600080fd5b815167ffffffffffffffff1981168114613bf957600080fd5b60006001600160601b03838116908316818110156154a2576154a2615207565b039392505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b838110156154e5578151855293820193908201906001016154c9565b5092979650505050505050565b6000808335601e1984360301811261550957600080fd5b8301803591506001600160401b0382111561552357600080fd5b60200191503681900382131561482157600080fd5b602081526000823561554981614533565b60ff81166020840152506020830135615561816144f9565b63ffffffff81166040840152506040830135601e1984360301811261558557600080fd5b830180356001600160401b0381111561559d57600080fd5b8036038513156155ac57600080fd5b60c060608501526155c460e085018260208501615310565b9150506155d36060850161450b565b63ffffffff81166080850152506155ec608085016145b1565b6001600160a01b03811660a08501525060a084013560c08401528091505092915050565b600063ffffffff80831681851680830382111561562f5761562f615207565b01949350505050565b8035615643816144f9565b63ffffffff16825260208181013590830152604090810135910152565b60608101613bfc8284615638565b60006001600160601b038083168185168183048111821515161561569457615694615207565b02949350505050565b60008160001904831182151516156156b7576156b7615207565b500290565b60a081016156ca8285615638565b63ffffffff8351166060830152602083015160808301529392505050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b8181101561572f57845183529383019391830191600101615713565b5090979650505050505050565b600061ffff8083168181141561575457615754615207565b6001019392505050565b60008351615770818460208801615039565b6001600160f81b031993909316919092019081526001019291505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a2646970667358221220ce1cae364f3ce5f58fa339670e6211b3c0a999ddd884b00c020bec9b22e2276164736f6c634300080c0033",
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

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractOpenOracleIdenticalAnswerTaskManager

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

// IOpenOracleIdenticalAnswerTaskManagerAggregatedMsg is an auto generated low-level Go binding around an user-defined struct.
type IOpenOracleIdenticalAnswerTaskManagerAggregatedMsg struct {
	ReferenceTaskIndex uint32
	Result             []byte
}

// IOpenOracleIdenticalAnswerTaskManagerAggregatedTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type IOpenOracleIdenticalAnswerTaskManagerAggregatedTaskResponse struct {
	Msg                 IOpenOracleIdenticalAnswerTaskManagerAggregatedMsg
	Timestamp           *big.Int
	AggregatedSignature BN254G1Point
	ApkG2               BN254G2Point
}

// IOpenOracleIdenticalAnswerTaskManagerTask is an auto generated low-level Go binding around an user-defined struct.
type IOpenOracleIdenticalAnswerTaskManagerTask struct {
	TaskType           uint8
	TaskData           []byte
	TaskCreatedBlock   uint32
	ResponderThreshold uint8
	StakeThreshold     *big.Int
	Creator            common.Address
	CreationFee        *big.Int
}

// IOpenOracleIdenticalAnswerTaskManagerTaskResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type IOpenOracleIdenticalAnswerTaskManagerTaskResponseMetadata struct {
	TaskResponsedBlock uint32
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

// ContractOpenOracleIdenticalAnswerTaskManagerMetaData contains all meta data concerning the ContractOpenOracleIdenticalAnswerTaskManager contract.
var ContractOpenOracleIdenticalAnswerTaskManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"},{\"name\":\"_BLSApkRegistry\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_blsSignatureChecker\",\"type\":\"address\",\"internalType\":\"contractBLSSignatureChecker\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addToFeedlist\",\"inputs\":[{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsSignatureChecker\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractBLSSignatureChecker\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"taskType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"taskData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"responderThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"stakeThreshold\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isFeed\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestTaskNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeFromFeed\",\"inputs\":[{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIOpenOracleIdenticalAnswerTaskManager.Task\",\"components\":[{\"name\":\"taskType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"taskData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"responderThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"stakeThreshold\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"creationFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"response\",\"type\":\"tuple\",\"internalType\":\"structIOpenOracleIdenticalAnswerTaskManager.AggregatedTaskResponse\",\"components\":[{\"name\":\"msg\",\"type\":\"tuple\",\"internalType\":\"structIOpenOracleIdenticalAnswerTaskManager.AggregatedMsg\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatedSignature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskCreationFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAggregator\",\"inputs\":[{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTaskCreationFee\",\"inputs\":[{\"name\":\"_newFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawTaskFunds\",\"inputs\":[{\"name\":\"taskNum\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIOpenOracleIdenticalAnswerTaskManager.Task\",\"components\":[{\"name\":\"taskType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"taskData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"responderThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"stakeThreshold\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"creationFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AddressAddedToFeedlist\",\"inputs\":[{\"name\":\"_address\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AddressRemovedFromFeedlist\",\"inputs\":[{\"name\":\"_address\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AggregatorUpdated\",\"inputs\":[{\"name\":\"aggregator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FundsWithdrawn\",\"inputs\":[{\"name\":\"taskNum\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"creator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewIdenticalAnswerTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOpenOracleIdenticalAnswerTaskManager.Task\",\"components\":[{\"name\":\"taskType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"taskData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"responderThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"stakeThreshold\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"creationFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOpenOracleIdenticalAnswerTaskManager.Task\",\"components\":[{\"name\":\"taskType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"taskData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"responderThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"stakeThreshold\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"creationFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOpenOracleIdenticalAnswerTaskManager.AggregatedTaskResponse\",\"components\":[{\"name\":\"msg\",\"type\":\"tuple\",\"internalType\":\"structIOpenOracleIdenticalAnswerTaskManager.AggregatedMsg\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatedSignature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOpenOracleIdenticalAnswerTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x610100604052655af3107a40006097553480156200001c57600080fd5b5060405162003ecc38038062003ecc8339810160408190526200003f916200007f565b6001600160a01b0393841660805291831660a05263ffffffff1660e0521660c052620000ef565b6001600160a01b03811681146200007c57600080fd5b50565b600080600080608085870312156200009657600080fd5b8451620000a38162000066565b6020860151909450620000b68162000066565b604086015190935063ffffffff81168114620000d157600080fd5b6060860151909250620000e48162000066565b939692955090935050565b60805160a05160c05160e051613d7962000153600039600081816101f9015281816104d8015261192801526000818161023a0152611dd20152600081816103a901528181611c1f0152611ce30152600081816103d00152611b130152613d796000f3fe608060405234801561001057600080fd5b50600436106101a15760003560e01c8063033ff004146101a65780630c0f5792146101bb57806310d67a2f146101ce578063136439dd146101e15780631ad43189146101f45780631c178e9c14610235578063245a7bfc146102695780632cb223d51461027c5780632d89f6fc146102aa5780633563b0d1146102ca5780633ccfd60b146102ea5780634d075ce7146102f25780634f739f74146102fb57806352dd59da1461031b57806352e9408b1461032e578063595c6a67146103415780635ac86ab7146103495780635c1556621461037c5780635c975abb1461039c5780635df45946146103a457806368304835146103cb578063715018a6146103f257806372d18e8d146103fa578063839cb9b814610408578063886f11951461041b578063898bef3a1461042e5780638b00ce7c146104415780638da5cb5b146104515780639fe4ee4714610459578063c0c53b8b1461046c578063cefdc1d41461047f578063e58fdd04146104a0578063f2fde38b146104c3578063f5c9899d146104d6578063fabc1cbc146104fc575b600080fd5b6101b96101b4366004612b97565b61050f565b005b6101b96101c9366004612b97565b610563565b6101b96101dc366004612b97565b6105b4565b6101b96101ef366004612bbb565b610670565b61021b7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020015b60405180910390f35b61025c7f000000000000000000000000000000000000000000000000000000000000000081565b60405161022c9190612bd4565b609c5461025c906001600160a01b031681565b61029c61028a366004612c0a565b609a6020526000908152604090205481565b60405190815260200161022c565b61029c6102b8366004612c0a565b60996020526000908152604090205481565b6102dd6102d8366004612d04565b61079d565b60405161022c9190612e0b565b6101b9610c1c565b61029c60975481565b61030e610309366004612eaa565b610c50565b60405161022c9190612f7f565b6101b9610329366004613076565b61130b565b6101b961033c366004612bbb565b611629565b6101b9611636565b61036c610357366004613156565b606654600160ff9092169190911b9081161490565b604051901515815260200161022c565b61038f61038a366004613194565b6116f0565b60405161022c919061324e565b60665461029c565b61025c7f000000000000000000000000000000000000000000000000000000000000000081565b61025c7f000000000000000000000000000000000000000000000000000000000000000081565b6101b96118a1565b60985463ffffffff1661021b565b6101b9610416366004613292565b6118b5565b60655461025c906001600160a01b031681565b6101b961043c36600461332d565b612062565b60985461021b9063ffffffff1681565b61025c612214565b6101b9610467366004612b97565b612223565b6101b961047a3660046133a6565b612282565b61049261048d3660046133e6565b6123bd565b60405161022c92919061341d565b61036c6104ae366004612b97565b609b6020526000908152604090205460ff1681565b6101b96104d1366004612b97565b61254f565b7f000000000000000000000000000000000000000000000000000000000000000061021b565b6101b961050a366004612bbb565b6125c5565b61051761271c565b6001600160a01b0381166000818152609b6020526040808220805460ff19166001179055517fe0fb1d5933f4eca7f5928d3c3c41d075398c1d1a4c4b09a5771a1c5c67eb1d869190a250565b61056b61271c565b6001600160a01b0381166000818152609b6020526040808220805460ff19169055517f28d57cd2f66bc88bf9eeda7d803691bff6e77bfdbf4212b41841a0d1b7189c5f9190a250565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610607573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061062b919061343e565b6001600160a01b0316336001600160a01b0316146106645760405162461bcd60e51b815260040161065b9061345b565b60405180910390fd5b61066d8161277b565b50565b60655460405163237dfb4760e11b81526001600160a01b03909116906346fbf68e906106a0903390600401612bd4565b602060405180830381865afa1580156106bd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106e191906134a3565b6106fd5760405162461bcd60e51b815260040161065b906134be565b606654818116146107715760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d707420604482015277746f20756e70617573652066756e6374696f6e616c69747960401b606482015260840161065b565b60668190556040518181523390600080516020613d04833981519152906020015b60405180910390a250565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156107df573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610803919061343e565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610845573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610869919061343e565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156108ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108cf919061343e565b9050600086516001600160401b038111156108ec576108ec612c27565b60405190808252806020026020018201604052801561091f57816020015b606081526020019060019003908161090a5790505b50905060005b8751811015610c10576000888281518110610942576109426134f4565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa1580156109a3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526109cb919081019061350a565b905080516001600160401b038111156109e6576109e6612c27565b604051908082528060200260200182016040528015610a3157816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610a045790505b50848481518110610a4457610a446134f4565b602002602001018190525060005b8151811015610bfa576040518060600160405280876001600160a01b03166347b314e8858581518110610a8757610a876134f4565b60200260200101516040518263ffffffff1660e01b8152600401610aad91815260200190565b602060405180830381865afa158015610aca573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aee919061343e565b6001600160a01b03168152602001838381518110610b0e57610b0e6134f4565b60200260200101518152602001896001600160a01b031663fa28c627858581518110610b3c57610b3c6134f4565b6020026020010151878f6040518463ffffffff1660e01b8152600401610b649392919061359a565b602060405180830381865afa158015610b81573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ba591906135b9565b6001600160601b0316815250858581518110610bc357610bc36134f4565b60200260200101518281518110610bdc57610bdc6134f4565b60200260200101819052508080610bf2906135ec565b915050610a52565b5050508080610c08906135ec565b915050610925565b50979650505050505050565b610c2461271c565b60405133904780156108fc02916000818181858888f1935050505015801561066d573d6000803e3d6000fd5b610c58612b3c565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c98573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cbc919061343e565b9050610cc6612b3c565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e90610cf6908b9089908990600401613607565b600060405180830381865afa158015610d13573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610d3b9190810190613651565b81526040516340e03a8160e11b81526001600160a01b038316906381c0750290610d6d908b908b908b90600401613708565b600060405180830381865afa158015610d8a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610db29190810190613651565b6040820152856001600160401b03811115610dcf57610dcf612c27565b604051908082528060200260200182016040528015610e0257816020015b6060815260200190600190039081610ded5790505b50606082015260005b60ff811687111561121c576000856001600160401b03811115610e3057610e30612c27565b604051908082528060200260200182016040528015610e59578160200160208202803683370190505b5083606001518360ff1681518110610e7357610e736134f4565b602002602001018190525060005b8681101561111c5760008c6001600160a01b03166304ec63518a8a85818110610eac57610eac6134f4565b905060200201358e88600001518681518110610eca57610eca6134f4565b60200260200101516040518463ffffffff1660e01b8152600401610ef093929190613731565b602060405180830381865afa158015610f0d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f31919061374d565b90506001600160c01b038116610fd45760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527b3132903932b3b4b9ba32b932b21030ba10313637b1b5b73ab6b132b960211b608482015260a40161065b565b8a8a8560ff16818110610fe957610fe96134f4565b6001600160c01b03841692013560f81c9190911c60019081161415905061110957856001600160a01b031663dd9846b98a8a8581811061102b5761102b6134f4565b905060200201358d8d8860ff16818110611047576110476134f4565b9050013560f81c60f81b60f81c8f6040518463ffffffff1660e01b81526004016110739392919061359a565b602060405180830381865afa158015611090573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110b49190613776565b85606001518560ff16815181106110cd576110cd6134f4565b602002602001015184815181106110e6576110e66134f4565b63ffffffff9092166020928302919091019091015282611105816135ec565b9350505b5080611114816135ec565b915050610e81565b506000816001600160401b0381111561113757611137612c27565b604051908082528060200260200182016040528015611160578160200160208202803683370190505b50905060005b828110156111e15784606001518460ff1681518110611187576111876134f4565b602002602001015181815181106111a0576111a06134f4565b60200260200101518282815181106111ba576111ba6134f4565b63ffffffff90921660209283029190910190910152806111d9816135ec565b915050611166565b508084606001518460ff16815181106111fc576111fc6134f4565b60200260200101819052505050808061121490613793565b915050610e0b565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa15801561125d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611281919061343e565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c906112b4908b908b908e906004016137b3565b600060405180830381865afa1580156112d1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526112f99190810190613651565b60208301525098975050505050505050565b6098548290829063ffffffff90811690831611156113645760405162461bcd60e51b81526020600482015260166024820152757461736b206e756d626572206e6f742065786973747360501b604482015260640161065b565b63ffffffff821660009081526099602090815260409182902054915161138c918491016137dd565b60405160208183030381529060405280519060200120146113bf5760405162461bcd60e51b815260040161065b906138a5565b60a08101516001600160a01b031633146114355760405162461bcd60e51b815260206004820152603160248201527f4f6e6c7920746865207461736b2063726561746f722063616e20706572666f7260448201527036903a3434b99037b832b930ba34b7b71760791b606482015260840161065b565b63ffffffff84166000908152609a60205260409020546114ab5760405162461bcd60e51b815260206004820152602b60248201527f43616e6e6f742077697468647261772066756e647320666f72206120636f6d7060448201526a3632ba32b2103a30b9b59760a91b606482015260840161065b565b60008360c00151116114f75760405162461bcd60e51b8152602060048201526015602482015274273790333ab73239903a37903bb4ba34323930bb9760591b604482015260640161065b565b60c08301805160009091526040516115139085906020016137dd565b60408051601f19818403018152828252805160209182012063ffffffff8916600090815260999092529181209190915560a086015190916001600160a01b039091169083908381818185875af1925050503d8060008114611590576040519150601f19603f3d011682016040523d82523d6000602084013e611595565b606091505b50509050806115db5760405162461bcd60e51b81526020600482015260126024820152712bb4ba34323930bbb0b6103330b4b632b21760711b604482015260640161065b565b6040805163ffffffff881681523360208201529081018390527ff440aec6b52895984d061d622e6edeba6210f7c3e059be920663140c084560d79060600160405180910390a1505050505050565b61163161271c565b609755565b60655460405163237dfb4760e11b81526001600160a01b03909116906346fbf68e90611666903390600401612bd4565b602060405180830381865afa158015611683573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116a791906134a3565b6116c35760405162461bcd60e51b815260040161065b906134be565b60001960668190556040519081523390600080516020613d048339815191529060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b8152600401611722929190613902565b600060405180830381865afa15801561173f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526117679190810190613651565b9050600084516001600160401b0381111561178457611784612c27565b6040519080825280602002602001820160405280156117ad578160200160208202803683370190505b50905060005b855181101561189757866001600160a01b03166304ec63518783815181106117dd576117dd6134f4565b6020026020010151878685815181106117f8576117f86134f4565b60200260200101516040518463ffffffff1660e01b815260040161181e93929190613731565b602060405180830381865afa15801561183b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061185f919061374d565b6001600160c01b031682828151811061187a5761187a6134f4565b60209081029190910101528061188f816135ec565b9150506117b3565b5095945050505050565b6118a961271c565b6118b36000612872565b565b609c546001600160a01b0316331461190f5760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c6572000000604482015260640161065b565b60006119216060860160408701612c0a565b905061194d7f000000000000000000000000000000000000000000000000000000000000000082613956565b63ffffffff164363ffffffff1611156119be5760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320726573706f6e64656420746f207468652060448201526c7461736b20746f6f206c61746560981b606482015260840161065b565b609960006119cc848061397e565b6119da906020810190612c0a565b63ffffffff1663ffffffff1681526020019081526020016000205485604051602001611a069190613a8f565b6040516020818303038152906040528051906020012014611a395760405162461bcd60e51b815260040161065b906138a5565b6000609a81611a48858061397e565b611a56906020810190612c0a565b63ffffffff1663ffffffff1681526020019081526020016000205414611ad35760405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b606482015260840161065b565b611ae360a0860160808701613aa2565b60405162fcdba760e51b8152600060048201523360248201526001600160601b0391909116906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690631f9b74e090604401602060405180830381865afa158015611b5a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b7e91906135b9565b6001600160601b03161015611be75760405162461bcd60e51b815260206004820152602960248201527f4f70657261746f72207374616b65206e6f74206d656574206d696e696d756d206044820152681d1a1c995cda1bdb1960ba1b606482015260840161065b565b6000611bf3838061397e565b604051602001611c039190613af6565b60405160208183030381529060405280519060200120905060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316637ff81a8787876000818110611c5f57611c5f6134f4565b9050602002016020810190611c749190612b97565b6040518263ffffffff1660e01b8152600401611c909190612bd4565b606060405180830381865afa158015611cad573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611cd19190613b09565b50905060015b85811015611db65760007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316637ff81a87898985818110611d2257611d226134f4565b9050602002016020810190611d379190612b97565b6040518263ffffffff1660e01b8152600401611d539190612bd4565b606060405180830381865afa158015611d70573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d949190613b09565b509050611da183826128c4565b50508080611dae906135ec565b915050611cd7565b506040805163171f1d5b60e01b81526000916001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169163171f1d5b91611e11918791879160808c01918c0190600401613b8f565b6040805180830381865afa158015611e2d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e519190613bd0565b50905080611ea15760405162461bcd60e51b815260206004820181905260248201527f416767726567617465207369676e6174757265206973206e6f742076616c6964604482015260640161065b565b604080516020808201835263ffffffff4316825291519091611ec7918891849101613c64565b60408051601f198184030181529190528051602090910120609a6000611eed898061397e565b611efb906020810190612c0a565b63ffffffff1681526020810191909152604001600090812091909155609b90611f2a60c08c0160a08d01612b97565b6001600160a01b0316815260208101919091526040016000205460ff16611fa45760405162461bcd60e51b815260206004820152602860248201527f46656564206e6565647320746f2072656769737465722077697468207461736b6044820152671036b0b730b3b2b960c11b606482015260840161065b565b6000611fb660c08b0160a08c01612b97565b60405163a3a9a98360e01b81529091506001600160a01b0382169063a3a9a98390611fe9908d908b908790600401613c8d565b600060405180830381600087803b15801561200357600080fd5b505af1158015612017573d6000803e3d6000fd5b505050507f5859b9b29fa1ffdbe40fa8bb3cdf904a8c0b31eef57809c082fd5549f3df40058a888460405161204e93929190613c8d565b60405180910390a150505050505050505050565b336000908152609b602052604090205460ff166120b65760405162461bcd60e51b815260206004820152601260248201527110d85b1b195c881a5cc81b9bdd081999595960721b604482015260640161065b565b6040805160e0810182526060602080830182905260008385018190529183018290526080830182905260a0830182905260c083019190915260ff881682528251601f87018290048202810182019093528583529091908690869081908401838280828437600092018290525060208087019590955263ffffffff43166040808801919091526001600160601b038816608088015260ff891660608801523360a088015260c087019190915251612173948694500191506137dd9050565b60408051601f1981840301815282825280516020918201206098805463ffffffff9081166000908152609990945293909220555416907fdc1d72dbd45105676b54bc1aa51f223c99997d23756fc461872ef079635bad40906121d69084906137dd565b60405180910390a26098546121f29063ffffffff166001613956565b6098805463ffffffff191663ffffffff92909216919091179055505050505050565b6033546001600160a01b031690565b61222b61271c565b609c80546001600160a01b0319166001600160a01b0383169081179091556040517f602cec4b1583b07d071161da5eb9589444d2459201e2fab7753dc941e9351c219161227791612bd4565b60405180910390a150565b600054610100900460ff16158080156122a25750600054600160ff909116105b806122bc5750303b1580156122bc575060005460ff166001145b61231f5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161065b565b6000805460ff191660011790558015612342576000805461ff0019166101001790555b61234d846000612967565b61235683612872565b609c80546001600160a01b0319166001600160a01b03841617905580156123b7576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b60408051600180825281830190925260009160609183916020808301908036833701905050905084816000815181106123f8576123f86134f4565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e906124349088908690600401613902565b600060405180830381865afa158015612451573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526124799190810190613651565b60008151811061248b5761248b6134f4565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa1580156124f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061251b919061374d565b6001600160c01b03169050600061253182612a3f565b90508161253f8a838a61079d565b9550955050505050935093915050565b61255761271c565b6001600160a01b0381166125bc5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161065b565b61066d81612872565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612618573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061263c919061343e565b6001600160a01b0316336001600160a01b03161461266c5760405162461bcd60e51b815260040161065b9061345b565b6066541981196066541916146126e55760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d706044820152777420746f2070617573652066756e6374696f6e616c69747960401b606482015260840161065b565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610792565b33612725612214565b6001600160a01b0316146118b35760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161065b565b6001600160a01b0381166128095760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a40161065b565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60408051808201909152600080825260208201526128e0612b64565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa905080801561291f57612921565bfe5b508061295f5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b604482015260640161065b565b505092915050565b6065546001600160a01b031615801561298857506001600160a01b03821615155b612a0a5760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a40161065b565b60668190556040518181523390600080516020613d048339815191529060200160405180910390a2612a3b8261277b565b5050565b6060600080612a4d84612b0b565b61ffff166001600160401b03811115612a6857612a68612c27565b6040519080825280601f01601f191660200182016040528015612a92576020820181803683370190505b5090506000805b825182108015612aaa575061010081105b15612b01576001811b935085841615612af1578060f81b838381518110612ad357612ad36134f4565b60200101906001600160f81b031916908160001a9053508160010191505b612afa816135ec565b9050612a99565b5090949350505050565b6000805b8215612b3657612b20600184613cca565b9092169180612b2e81613ce1565b915050612b0f565b92915050565b6040518060800160405280606081526020016060815260200160608152602001606081525090565b60405180608001604052806004906020820280368337509192915050565b6001600160a01b038116811461066d57600080fd5b600060208284031215612ba957600080fd5b8135612bb481612b82565b9392505050565b600060208284031215612bcd57600080fd5b5035919050565b6001600160a01b0391909116815260200190565b63ffffffff8116811461066d57600080fd5b8035612c0581612be8565b919050565b600060208284031215612c1c57600080fd5b8135612bb481612be8565b634e487b7160e01b600052604160045260246000fd5b60405160e081016001600160401b0381118282101715612c5f57612c5f612c27565b60405290565b604051601f8201601f191681016001600160401b0381118282101715612c8d57612c8d612c27565b604052919050565b600082601f830112612ca657600080fd5b81356001600160401b03811115612cbf57612cbf612c27565b612cd2601f8201601f1916602001612c65565b818152846020838601011115612ce757600080fd5b816020850160208301376000918101602001919091529392505050565b600080600060608486031215612d1957600080fd5b8335612d2481612b82565b925060208401356001600160401b03811115612d3f57600080fd5b612d4b86828701612c95565b9250506040840135612d5c81612be8565b809150509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015612dfd578385038a52825180518087529087019087870190845b81811015612de857835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101612da4565b50509a87019a95505091850191600101612d86565b509298975050505050505050565b602081526000612bb46020830184612d67565b60008083601f840112612e3057600080fd5b5081356001600160401b03811115612e4757600080fd5b602083019150836020828501011115612e5f57600080fd5b9250929050565b60008083601f840112612e7857600080fd5b5081356001600160401b03811115612e8f57600080fd5b6020830191508360208260051b8501011115612e5f57600080fd5b60008060008060008060808789031215612ec357600080fd5b8635612ece81612b82565b95506020870135612ede81612be8565b945060408701356001600160401b0380821115612efa57600080fd5b612f068a838b01612e1e565b90965094506060890135915080821115612f1f57600080fd5b50612f2c89828a01612e66565b979a9699509497509295939492505050565b600081518084526020808501945080840160005b83811015612f7457815163ffffffff1687529582019590820190600101612f52565b509495945050505050565b600060208083528351608082850152612f9b60a0850182612f3e565b905081850151601f1980868403016040870152612fb88383612f3e565b92506040870151915080868403016060870152612fd58383612f3e565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b8281101561302c578487830301845261301a828751612f3e565b95880195938801939150600101613000565b509998505050505050505050565b803560ff81168114612c0557600080fd5b6001600160601b038116811461066d57600080fd5b8035612c058161304b565b8035612c0581612b82565b6000806040838503121561308957600080fd5b823561309481612be8565b915060208301356001600160401b03808211156130b057600080fd5b9084019060e082870312156130c457600080fd5b6130cc612c3d565b6130d58361303a565b81526020830135828111156130e957600080fd5b6130f588828601612c95565b60208301525061310760408401612bfa565b60408201526131186060840161303a565b606082015261312960808401613060565b608082015261313a60a0840161306b565b60a082015260c083013560c08201528093505050509250929050565b60006020828403121561316857600080fd5b612bb48261303a565b60006001600160401b0382111561318a5761318a612c27565b5060051b60200190565b6000806000606084860312156131a957600080fd5b83356131b481612b82565b92506020848101356001600160401b038111156131d057600080fd5b8501601f810187136131e157600080fd5b80356131f46131ef82613171565b612c65565b81815260059190911b8201830190838101908983111561321357600080fd5b928401925b8284101561323157833582529284019290840190613218565b809650505050505061324560408501612bfa565b90509250925092565b6020808252825182820181905260009190848201906040850190845b818110156132865783518352928401929184019160010161326a565b50909695505050505050565b600080600080606085870312156132a857600080fd5b84356001600160401b03808211156132bf57600080fd5b9086019060e082890312156132d357600080fd5b909450602086013590808211156132e957600080fd5b6132f588838901612e66565b9095509350604087013591508082111561330e57600080fd5b508501610100818803121561332257600080fd5b939692955090935050565b60008060008060006080868803121561334557600080fd5b61334e8661303a565b945060208601356001600160401b0381111561336957600080fd5b61337588828901612e1e565b909550935061338890506040870161303a565b915060608601356133988161304b565b809150509295509295909350565b6000806000606084860312156133bb57600080fd5b83356133c681612b82565b925060208401356133d681612b82565b91506040840135612d5c81612b82565b6000806000606084860312156133fb57600080fd5b833561340681612b82565b9250602084013591506040840135612d5c81612be8565b8281526040602082015260006134366040830184612d67565b949350505050565b60006020828403121561345057600080fd5b8151612bb481612b82565b6020808252602a90820152600080516020613d2483398151915260408201526939903ab73830bab9b2b960b11b606082015260800190565b80518015158114612c0557600080fd5b6000602082840312156134b557600080fd5b612bb482613493565b6020808252602890820152600080516020613d2483398151915260408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b6000602080838503121561351d57600080fd5b82516001600160401b0381111561353357600080fd5b8301601f8101851361354457600080fd5b80516135526131ef82613171565b81815260059190911b8201830190838101908783111561357157600080fd5b928401925b8284101561358f57835182529284019290840190613576565b979650505050505050565b92835260ff91909116602083015263ffffffff16604082015260600190565b6000602082840312156135cb57600080fd5b8151612bb48161304b565b634e487b7160e01b600052601160045260246000fd5b6000600019821415613600576136006135d6565b5060010190565b63ffffffff84168152604060208201819052810182905260006001600160fb1b0383111561363457600080fd5b8260051b8085606085013760009201606001918252509392505050565b6000602080838503121561366457600080fd5b82516001600160401b0381111561367a57600080fd5b8301601f8101851361368b57600080fd5b80516136996131ef82613171565b81815260059190911b820183019083810190878311156136b857600080fd5b928401925b8284101561358f5783516136d081612be8565b825292840192908401906136bd565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff841681526040602082015260006137286040830184866136df565b95945050505050565b92835263ffffffff918216602084015216604082015260600190565b60006020828403121561375f57600080fd5b81516001600160c01b0381168114612bb457600080fd5b60006020828403121561378857600080fd5b8151612bb481612be8565b600060ff821660ff8114156137aa576137aa6135d6565b60010192915050565b6040815260006137c76040830185876136df565b905063ffffffff83166020830152949350505050565b6000602080835260ff845116818401528084015160e0604085015280518061010086015260005b818110156138215782810184015186820161012001528301613804565b8181111561383457600061012083880101525b50604086015163ffffffff811660608701529250606086015160ff81166080870152925060808601516001600160601b03811660a0870152925060a08601516001600160a01b03811660c0870152925060c0959095015160e0850152505050601f91909101601f1916016101200190565b6020808252603d908201527f737570706c696564207461736b20646f6573206e6f74206d617463682074686560408201527f206f6e65207265636f7264656420696e2074686520636f6e7472616374000000606082015260800190565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b818110156139495784518352938301939183019160010161392d565b5090979650505050505050565b600063ffffffff808316818516808303821115613975576139756135d6565b01949350505050565b60008235603e1983360301811261399457600080fd5b9190910192915050565b6000808335601e198436030181126139b557600080fd5b83016020810192503590506001600160401b038111156139d457600080fd5b803603831315612e5f57600080fd5b60ff6139ee8261303a565b1682526000613a00602083018361399e565b60e06020860152613a1560e0860182846136df565b9150506040830135613a2681612be8565b63ffffffff16604085015260ff613a3f6060850161303a565b1660608501526080830135613a538161304b565b6001600160601b0316608085015260a0830135613a6f81612b82565b6001600160a01b031660a085015260c09283013592909301919091525090565b602081526000612bb460208301846139e3565b600060208284031215613ab457600080fd5b8135612bb48161304b565b60008135613acc81612be8565b63ffffffff168352613ae1602083018361399e565b604060208601526137286040860182846136df565b602081526000612bb46020830184613abf565b6000808284036060811215613b1d57600080fd5b6040811215613b2b57600080fd5b50604080519081016001600160401b0381118282101715613b4e57613b4e612c27565b60409081528451825260208086015190830152939093015192949293505050565b604081833760408201600081526040808301823750600060808301525050565b6000610120820190508582528451602083015260208501516040830152613bb96060830185613b6f565b823560e08301526020830135610100830152613728565b60008060408385031215613be357600080fd5b613bec83613493565b9150613bfa60208401613493565b90509250929050565b60006101008235603e19843603018112613c1c57600080fd5b818552613c2d828601858301613abf565b91505060208301356020850152613c54604085016040850180358252602090810135910152565b612bb46080850160808501613b6f565b604081526000613c776040830185613c03565b905063ffffffff83511660208301529392505050565b606081526000613ca060608301866139e3565b8281036020840152613cb28186613c03565b91505063ffffffff8351166040830152949350505050565b600082821015613cdc57613cdc6135d6565b500390565b600061ffff80831681811415613cf957613cf96135d6565b600101939250505056feab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d6d73672e73656e646572206973206e6f74207065726d697373696f6e65642061a2646970667358221220fa80b4d432c2dedc5c014776df51682d05ff9e0e1b6f0df45a45e03ff169feea64736f6c634300080c0033",
}

// ContractOpenOracleIdenticalAnswerTaskManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractOpenOracleIdenticalAnswerTaskManagerMetaData.ABI instead.
var ContractOpenOracleIdenticalAnswerTaskManagerABI = ContractOpenOracleIdenticalAnswerTaskManagerMetaData.ABI

// ContractOpenOracleIdenticalAnswerTaskManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractOpenOracleIdenticalAnswerTaskManagerMetaData.Bin instead.
var ContractOpenOracleIdenticalAnswerTaskManagerBin = ContractOpenOracleIdenticalAnswerTaskManagerMetaData.Bin

// DeployContractOpenOracleIdenticalAnswerTaskManager deploys a new Ethereum contract, binding an instance of ContractOpenOracleIdenticalAnswerTaskManager to it.
func DeployContractOpenOracleIdenticalAnswerTaskManager(auth *bind.TransactOpts, backend bind.ContractBackend, _stakeRegistry common.Address, _BLSApkRegistry common.Address, _taskResponseWindowBlock uint32, _blsSignatureChecker common.Address) (common.Address, *types.Transaction, *ContractOpenOracleIdenticalAnswerTaskManager, error) {
	parsed, err := ContractOpenOracleIdenticalAnswerTaskManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractOpenOracleIdenticalAnswerTaskManagerBin), backend, _stakeRegistry, _BLSApkRegistry, _taskResponseWindowBlock, _blsSignatureChecker)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractOpenOracleIdenticalAnswerTaskManager{ContractOpenOracleIdenticalAnswerTaskManagerCaller: ContractOpenOracleIdenticalAnswerTaskManagerCaller{contract: contract}, ContractOpenOracleIdenticalAnswerTaskManagerTransactor: ContractOpenOracleIdenticalAnswerTaskManagerTransactor{contract: contract}, ContractOpenOracleIdenticalAnswerTaskManagerFilterer: ContractOpenOracleIdenticalAnswerTaskManagerFilterer{contract: contract}}, nil
}

// ContractOpenOracleIdenticalAnswerTaskManager is an auto generated Go binding around an Ethereum contract.
type ContractOpenOracleIdenticalAnswerTaskManager struct {
	ContractOpenOracleIdenticalAnswerTaskManagerCaller     // Read-only binding to the contract
	ContractOpenOracleIdenticalAnswerTaskManagerTransactor // Write-only binding to the contract
	ContractOpenOracleIdenticalAnswerTaskManagerFilterer   // Log filterer for contract events
}

// ContractOpenOracleIdenticalAnswerTaskManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractOpenOracleIdenticalAnswerTaskManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOpenOracleIdenticalAnswerTaskManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractOpenOracleIdenticalAnswerTaskManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOpenOracleIdenticalAnswerTaskManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractOpenOracleIdenticalAnswerTaskManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOpenOracleIdenticalAnswerTaskManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractOpenOracleIdenticalAnswerTaskManagerSession struct {
	Contract     *ContractOpenOracleIdenticalAnswerTaskManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts                             // Transaction auth options to use throughout this session
}

// ContractOpenOracleIdenticalAnswerTaskManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractOpenOracleIdenticalAnswerTaskManagerCallerSession struct {
	Contract *ContractOpenOracleIdenticalAnswerTaskManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                                       // Call options to use throughout this session
}

// ContractOpenOracleIdenticalAnswerTaskManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractOpenOracleIdenticalAnswerTaskManagerTransactorSession struct {
	Contract     *ContractOpenOracleIdenticalAnswerTaskManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                                       // Transaction auth options to use throughout this session
}

// ContractOpenOracleIdenticalAnswerTaskManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractOpenOracleIdenticalAnswerTaskManagerRaw struct {
	Contract *ContractOpenOracleIdenticalAnswerTaskManager // Generic contract binding to access the raw methods on
}

// ContractOpenOracleIdenticalAnswerTaskManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractOpenOracleIdenticalAnswerTaskManagerCallerRaw struct {
	Contract *ContractOpenOracleIdenticalAnswerTaskManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractOpenOracleIdenticalAnswerTaskManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractOpenOracleIdenticalAnswerTaskManagerTransactorRaw struct {
	Contract *ContractOpenOracleIdenticalAnswerTaskManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractOpenOracleIdenticalAnswerTaskManager creates a new instance of ContractOpenOracleIdenticalAnswerTaskManager, bound to a specific deployed contract.
func NewContractOpenOracleIdenticalAnswerTaskManager(address common.Address, backend bind.ContractBackend) (*ContractOpenOracleIdenticalAnswerTaskManager, error) {
	contract, err := bindContractOpenOracleIdenticalAnswerTaskManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleIdenticalAnswerTaskManager{ContractOpenOracleIdenticalAnswerTaskManagerCaller: ContractOpenOracleIdenticalAnswerTaskManagerCaller{contract: contract}, ContractOpenOracleIdenticalAnswerTaskManagerTransactor: ContractOpenOracleIdenticalAnswerTaskManagerTransactor{contract: contract}, ContractOpenOracleIdenticalAnswerTaskManagerFilterer: ContractOpenOracleIdenticalAnswerTaskManagerFilterer{contract: contract}}, nil
}

// NewContractOpenOracleIdenticalAnswerTaskManagerCaller creates a new read-only instance of ContractOpenOracleIdenticalAnswerTaskManager, bound to a specific deployed contract.
func NewContractOpenOracleIdenticalAnswerTaskManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractOpenOracleIdenticalAnswerTaskManagerCaller, error) {
	contract, err := bindContractOpenOracleIdenticalAnswerTaskManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleIdenticalAnswerTaskManagerCaller{contract: contract}, nil
}

// NewContractOpenOracleIdenticalAnswerTaskManagerTransactor creates a new write-only instance of ContractOpenOracleIdenticalAnswerTaskManager, bound to a specific deployed contract.
func NewContractOpenOracleIdenticalAnswerTaskManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractOpenOracleIdenticalAnswerTaskManagerTransactor, error) {
	contract, err := bindContractOpenOracleIdenticalAnswerTaskManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleIdenticalAnswerTaskManagerTransactor{contract: contract}, nil
}

// NewContractOpenOracleIdenticalAnswerTaskManagerFilterer creates a new log filterer instance of ContractOpenOracleIdenticalAnswerTaskManager, bound to a specific deployed contract.
func NewContractOpenOracleIdenticalAnswerTaskManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractOpenOracleIdenticalAnswerTaskManagerFilterer, error) {
	contract, err := bindContractOpenOracleIdenticalAnswerTaskManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleIdenticalAnswerTaskManagerFilterer{contract: contract}, nil
}

// bindContractOpenOracleIdenticalAnswerTaskManager binds a generic wrapper to an already deployed contract.
func bindContractOpenOracleIdenticalAnswerTaskManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractOpenOracleIdenticalAnswerTaskManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.ContractOpenOracleIdenticalAnswerTaskManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.ContractOpenOracleIdenticalAnswerTaskManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.ContractOpenOracleIdenticalAnswerTaskManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.contract.Transact(opts, method, params...)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCaller) TASKRESPONSEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.Call(opts, &out, "TASK_RESPONSE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCallerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) Aggregator() (common.Address, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.Aggregator(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCallerSession) Aggregator() (common.Address, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.Aggregator(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCaller) AllTaskHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.Call(opts, &out, "allTaskHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.AllTaskHashes(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts, arg0)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCallerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.AllTaskHashes(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCaller) AllTaskResponses(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.Call(opts, &out, "allTaskResponses", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.AllTaskResponses(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCallerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.AllTaskResponses(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts, arg0)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.BlsApkRegistry(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.BlsApkRegistry(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts)
}

// BlsSignatureChecker is a free data retrieval call binding the contract method 0x1c178e9c.
//
// Solidity: function blsSignatureChecker() view returns(address)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCaller) BlsSignatureChecker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.Call(opts, &out, "blsSignatureChecker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsSignatureChecker is a free data retrieval call binding the contract method 0x1c178e9c.
//
// Solidity: function blsSignatureChecker() view returns(address)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) BlsSignatureChecker() (common.Address, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.BlsSignatureChecker(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts)
}

// BlsSignatureChecker is a free data retrieval call binding the contract method 0x1c178e9c.
//
// Solidity: function blsSignatureChecker() view returns(address)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCallerSession) BlsSignatureChecker() (common.Address, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.BlsSignatureChecker(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.GetCheckSignaturesIndices(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.GetCheckSignaturesIndices(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.GetOperatorState(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.GetOperatorState(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

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
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.GetOperatorState0(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.GetOperatorState0(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCaller) GetQuorumBitmapsAtBlockNumber(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.Call(opts, &out, "getQuorumBitmapsAtBlockNumber", registryCoordinator, operatorIds, blockNumber)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCallerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCaller) GetTaskResponseWindowBlock(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.Call(opts, &out, "getTaskResponseWindowBlock")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCallerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts)
}

// IsFeed is a free data retrieval call binding the contract method 0xe58fdd04.
//
// Solidity: function isFeed(address ) view returns(bool)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCaller) IsFeed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.Call(opts, &out, "isFeed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFeed is a free data retrieval call binding the contract method 0xe58fdd04.
//
// Solidity: function isFeed(address ) view returns(bool)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) IsFeed(arg0 common.Address) (bool, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.IsFeed(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts, arg0)
}

// IsFeed is a free data retrieval call binding the contract method 0xe58fdd04.
//
// Solidity: function isFeed(address ) view returns(bool)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCallerSession) IsFeed(arg0 common.Address) (bool, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.IsFeed(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts, arg0)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCaller) LatestTaskNum(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.Call(opts, &out, "latestTaskNum")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) LatestTaskNum() (uint32, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.LatestTaskNum(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCallerSession) LatestTaskNum() (uint32, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.LatestTaskNum(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) Owner() (common.Address, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.Owner(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCallerSession) Owner() (common.Address, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.Owner(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) Paused(index uint8) (bool, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.Paused(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.Paused(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) Paused0() (*big.Int, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.Paused0(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.Paused0(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.PauserRegistry(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.PauserRegistry(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) StakeRegistry() (common.Address, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.StakeRegistry(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.StakeRegistry(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts)
}

// TaskCreationFee is a free data retrieval call binding the contract method 0x4d075ce7.
//
// Solidity: function taskCreationFee() view returns(uint256)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCaller) TaskCreationFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.Call(opts, &out, "taskCreationFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaskCreationFee is a free data retrieval call binding the contract method 0x4d075ce7.
//
// Solidity: function taskCreationFee() view returns(uint256)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) TaskCreationFee() (*big.Int, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.TaskCreationFee(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts)
}

// TaskCreationFee is a free data retrieval call binding the contract method 0x4d075ce7.
//
// Solidity: function taskCreationFee() view returns(uint256)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCallerSession) TaskCreationFee() (*big.Int, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.TaskCreationFee(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCaller) TaskNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.Call(opts, &out, "taskNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) TaskNumber() (uint32, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.TaskNumber(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerCallerSession) TaskNumber() (uint32, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.TaskNumber(&_ContractOpenOracleIdenticalAnswerTaskManager.CallOpts)
}

// AddToFeedlist is a paid mutator transaction binding the contract method 0x033ff004.
//
// Solidity: function addToFeedlist(address _address) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerTransactor) AddToFeedlist(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.contract.Transact(opts, "addToFeedlist", _address)
}

// AddToFeedlist is a paid mutator transaction binding the contract method 0x033ff004.
//
// Solidity: function addToFeedlist(address _address) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) AddToFeedlist(_address common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.AddToFeedlist(&_ContractOpenOracleIdenticalAnswerTaskManager.TransactOpts, _address)
}

// AddToFeedlist is a paid mutator transaction binding the contract method 0x033ff004.
//
// Solidity: function addToFeedlist(address _address) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerTransactorSession) AddToFeedlist(_address common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.AddToFeedlist(&_ContractOpenOracleIdenticalAnswerTaskManager.TransactOpts, _address)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x898bef3a.
//
// Solidity: function createNewTask(uint8 taskType, bytes taskData, uint8 responderThreshold, uint96 stakeThreshold) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerTransactor) CreateNewTask(opts *bind.TransactOpts, taskType uint8, taskData []byte, responderThreshold uint8, stakeThreshold *big.Int) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.contract.Transact(opts, "createNewTask", taskType, taskData, responderThreshold, stakeThreshold)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x898bef3a.
//
// Solidity: function createNewTask(uint8 taskType, bytes taskData, uint8 responderThreshold, uint96 stakeThreshold) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) CreateNewTask(taskType uint8, taskData []byte, responderThreshold uint8, stakeThreshold *big.Int) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.CreateNewTask(&_ContractOpenOracleIdenticalAnswerTaskManager.TransactOpts, taskType, taskData, responderThreshold, stakeThreshold)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x898bef3a.
//
// Solidity: function createNewTask(uint8 taskType, bytes taskData, uint8 responderThreshold, uint96 stakeThreshold) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerTransactorSession) CreateNewTask(taskType uint8, taskData []byte, responderThreshold uint8, stakeThreshold *big.Int) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.CreateNewTask(&_ContractOpenOracleIdenticalAnswerTaskManager.TransactOpts, taskType, taskData, responderThreshold, stakeThreshold)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.contract.Transact(opts, "initialize", _pauserRegistry, initialOwner, _aggregator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.Initialize(&_ContractOpenOracleIdenticalAnswerTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerTransactorSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.Initialize(&_ContractOpenOracleIdenticalAnswerTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.Pause(&_ContractOpenOracleIdenticalAnswerTaskManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.Pause(&_ContractOpenOracleIdenticalAnswerTaskManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.PauseAll(&_ContractOpenOracleIdenticalAnswerTaskManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.PauseAll(&_ContractOpenOracleIdenticalAnswerTaskManager.TransactOpts)
}

// RemoveFromFeed is a paid mutator transaction binding the contract method 0x0c0f5792.
//
// Solidity: function removeFromFeed(address _address) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerTransactor) RemoveFromFeed(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.contract.Transact(opts, "removeFromFeed", _address)
}

// RemoveFromFeed is a paid mutator transaction binding the contract method 0x0c0f5792.
//
// Solidity: function removeFromFeed(address _address) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) RemoveFromFeed(_address common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.RemoveFromFeed(&_ContractOpenOracleIdenticalAnswerTaskManager.TransactOpts, _address)
}

// RemoveFromFeed is a paid mutator transaction binding the contract method 0x0c0f5792.
//
// Solidity: function removeFromFeed(address _address) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerTransactorSession) RemoveFromFeed(_address common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.RemoveFromFeed(&_ContractOpenOracleIdenticalAnswerTaskManager.TransactOpts, _address)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.RenounceOwnership(&_ContractOpenOracleIdenticalAnswerTaskManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.RenounceOwnership(&_ContractOpenOracleIdenticalAnswerTaskManager.TransactOpts)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x839cb9b8.
//
// Solidity: function respondToTask((uint8,bytes,uint32,uint8,uint96,address,uint256) task, address[] operators, ((uint32,bytes),uint256,(uint256,uint256),(uint256[2],uint256[2])) response) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerTransactor) RespondToTask(opts *bind.TransactOpts, task IOpenOracleIdenticalAnswerTaskManagerTask, operators []common.Address, response IOpenOracleIdenticalAnswerTaskManagerAggregatedTaskResponse) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.contract.Transact(opts, "respondToTask", task, operators, response)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x839cb9b8.
//
// Solidity: function respondToTask((uint8,bytes,uint32,uint8,uint96,address,uint256) task, address[] operators, ((uint32,bytes),uint256,(uint256,uint256),(uint256[2],uint256[2])) response) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) RespondToTask(task IOpenOracleIdenticalAnswerTaskManagerTask, operators []common.Address, response IOpenOracleIdenticalAnswerTaskManagerAggregatedTaskResponse) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.RespondToTask(&_ContractOpenOracleIdenticalAnswerTaskManager.TransactOpts, task, operators, response)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x839cb9b8.
//
// Solidity: function respondToTask((uint8,bytes,uint32,uint8,uint96,address,uint256) task, address[] operators, ((uint32,bytes),uint256,(uint256,uint256),(uint256[2],uint256[2])) response) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerTransactorSession) RespondToTask(task IOpenOracleIdenticalAnswerTaskManagerTask, operators []common.Address, response IOpenOracleIdenticalAnswerTaskManagerAggregatedTaskResponse) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.RespondToTask(&_ContractOpenOracleIdenticalAnswerTaskManager.TransactOpts, task, operators, response)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.SetPauserRegistry(&_ContractOpenOracleIdenticalAnswerTaskManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.SetPauserRegistry(&_ContractOpenOracleIdenticalAnswerTaskManager.TransactOpts, newPauserRegistry)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.TransferOwnership(&_ContractOpenOracleIdenticalAnswerTaskManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.TransferOwnership(&_ContractOpenOracleIdenticalAnswerTaskManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.Unpause(&_ContractOpenOracleIdenticalAnswerTaskManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.Unpause(&_ContractOpenOracleIdenticalAnswerTaskManager.TransactOpts, newPausedStatus)
}

// UpdateAggregator is a paid mutator transaction binding the contract method 0x9fe4ee47.
//
// Solidity: function updateAggregator(address _aggregator) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerTransactor) UpdateAggregator(opts *bind.TransactOpts, _aggregator common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.contract.Transact(opts, "updateAggregator", _aggregator)
}

// UpdateAggregator is a paid mutator transaction binding the contract method 0x9fe4ee47.
//
// Solidity: function updateAggregator(address _aggregator) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) UpdateAggregator(_aggregator common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.UpdateAggregator(&_ContractOpenOracleIdenticalAnswerTaskManager.TransactOpts, _aggregator)
}

// UpdateAggregator is a paid mutator transaction binding the contract method 0x9fe4ee47.
//
// Solidity: function updateAggregator(address _aggregator) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerTransactorSession) UpdateAggregator(_aggregator common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.UpdateAggregator(&_ContractOpenOracleIdenticalAnswerTaskManager.TransactOpts, _aggregator)
}

// UpdateTaskCreationFee is a paid mutator transaction binding the contract method 0x52e9408b.
//
// Solidity: function updateTaskCreationFee(uint256 _newFee) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerTransactor) UpdateTaskCreationFee(opts *bind.TransactOpts, _newFee *big.Int) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.contract.Transact(opts, "updateTaskCreationFee", _newFee)
}

// UpdateTaskCreationFee is a paid mutator transaction binding the contract method 0x52e9408b.
//
// Solidity: function updateTaskCreationFee(uint256 _newFee) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) UpdateTaskCreationFee(_newFee *big.Int) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.UpdateTaskCreationFee(&_ContractOpenOracleIdenticalAnswerTaskManager.TransactOpts, _newFee)
}

// UpdateTaskCreationFee is a paid mutator transaction binding the contract method 0x52e9408b.
//
// Solidity: function updateTaskCreationFee(uint256 _newFee) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerTransactorSession) UpdateTaskCreationFee(_newFee *big.Int) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.UpdateTaskCreationFee(&_ContractOpenOracleIdenticalAnswerTaskManager.TransactOpts, _newFee)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) Withdraw() (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.Withdraw(&_ContractOpenOracleIdenticalAnswerTaskManager.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerTransactorSession) Withdraw() (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.Withdraw(&_ContractOpenOracleIdenticalAnswerTaskManager.TransactOpts)
}

// WithdrawTaskFunds is a paid mutator transaction binding the contract method 0x52dd59da.
//
// Solidity: function withdrawTaskFunds(uint32 taskNum, (uint8,bytes,uint32,uint8,uint96,address,uint256) task) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerTransactor) WithdrawTaskFunds(opts *bind.TransactOpts, taskNum uint32, task IOpenOracleIdenticalAnswerTaskManagerTask) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.contract.Transact(opts, "withdrawTaskFunds", taskNum, task)
}

// WithdrawTaskFunds is a paid mutator transaction binding the contract method 0x52dd59da.
//
// Solidity: function withdrawTaskFunds(uint32 taskNum, (uint8,bytes,uint32,uint8,uint96,address,uint256) task) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerSession) WithdrawTaskFunds(taskNum uint32, task IOpenOracleIdenticalAnswerTaskManagerTask) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.WithdrawTaskFunds(&_ContractOpenOracleIdenticalAnswerTaskManager.TransactOpts, taskNum, task)
}

// WithdrawTaskFunds is a paid mutator transaction binding the contract method 0x52dd59da.
//
// Solidity: function withdrawTaskFunds(uint32 taskNum, (uint8,bytes,uint32,uint8,uint96,address,uint256) task) returns()
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerTransactorSession) WithdrawTaskFunds(taskNum uint32, task IOpenOracleIdenticalAnswerTaskManagerTask) (*types.Transaction, error) {
	return _ContractOpenOracleIdenticalAnswerTaskManager.Contract.WithdrawTaskFunds(&_ContractOpenOracleIdenticalAnswerTaskManager.TransactOpts, taskNum, task)
}

// ContractOpenOracleIdenticalAnswerTaskManagerAddressAddedToFeedlistIterator is returned from FilterAddressAddedToFeedlist and is used to iterate over the raw logs and unpacked data for AddressAddedToFeedlist events raised by the ContractOpenOracleIdenticalAnswerTaskManager contract.
type ContractOpenOracleIdenticalAnswerTaskManagerAddressAddedToFeedlistIterator struct {
	Event *ContractOpenOracleIdenticalAnswerTaskManagerAddressAddedToFeedlist // Event containing the contract specifics and raw log

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
func (it *ContractOpenOracleIdenticalAnswerTaskManagerAddressAddedToFeedlistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenOracleIdenticalAnswerTaskManagerAddressAddedToFeedlist)
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
		it.Event = new(ContractOpenOracleIdenticalAnswerTaskManagerAddressAddedToFeedlist)
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
func (it *ContractOpenOracleIdenticalAnswerTaskManagerAddressAddedToFeedlistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenOracleIdenticalAnswerTaskManagerAddressAddedToFeedlistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenOracleIdenticalAnswerTaskManagerAddressAddedToFeedlist represents a AddressAddedToFeedlist event raised by the ContractOpenOracleIdenticalAnswerTaskManager contract.
type ContractOpenOracleIdenticalAnswerTaskManagerAddressAddedToFeedlist struct {
	Address common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddressAddedToFeedlist is a free log retrieval operation binding the contract event 0xe0fb1d5933f4eca7f5928d3c3c41d075398c1d1a4c4b09a5771a1c5c67eb1d86.
//
// Solidity: event AddressAddedToFeedlist(address indexed _address)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) FilterAddressAddedToFeedlist(opts *bind.FilterOpts, _address []common.Address) (*ContractOpenOracleIdenticalAnswerTaskManagerAddressAddedToFeedlistIterator, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.FilterLogs(opts, "AddressAddedToFeedlist", _addressRule)
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleIdenticalAnswerTaskManagerAddressAddedToFeedlistIterator{contract: _ContractOpenOracleIdenticalAnswerTaskManager.contract, event: "AddressAddedToFeedlist", logs: logs, sub: sub}, nil
}

// WatchAddressAddedToFeedlist is a free log subscription operation binding the contract event 0xe0fb1d5933f4eca7f5928d3c3c41d075398c1d1a4c4b09a5771a1c5c67eb1d86.
//
// Solidity: event AddressAddedToFeedlist(address indexed _address)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) WatchAddressAddedToFeedlist(opts *bind.WatchOpts, sink chan<- *ContractOpenOracleIdenticalAnswerTaskManagerAddressAddedToFeedlist, _address []common.Address) (event.Subscription, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.WatchLogs(opts, "AddressAddedToFeedlist", _addressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenOracleIdenticalAnswerTaskManagerAddressAddedToFeedlist)
				if err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.UnpackLog(event, "AddressAddedToFeedlist", log); err != nil {
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
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) ParseAddressAddedToFeedlist(log types.Log) (*ContractOpenOracleIdenticalAnswerTaskManagerAddressAddedToFeedlist, error) {
	event := new(ContractOpenOracleIdenticalAnswerTaskManagerAddressAddedToFeedlist)
	if err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.UnpackLog(event, "AddressAddedToFeedlist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOpenOracleIdenticalAnswerTaskManagerAddressRemovedFromFeedlistIterator is returned from FilterAddressRemovedFromFeedlist and is used to iterate over the raw logs and unpacked data for AddressRemovedFromFeedlist events raised by the ContractOpenOracleIdenticalAnswerTaskManager contract.
type ContractOpenOracleIdenticalAnswerTaskManagerAddressRemovedFromFeedlistIterator struct {
	Event *ContractOpenOracleIdenticalAnswerTaskManagerAddressRemovedFromFeedlist // Event containing the contract specifics and raw log

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
func (it *ContractOpenOracleIdenticalAnswerTaskManagerAddressRemovedFromFeedlistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenOracleIdenticalAnswerTaskManagerAddressRemovedFromFeedlist)
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
		it.Event = new(ContractOpenOracleIdenticalAnswerTaskManagerAddressRemovedFromFeedlist)
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
func (it *ContractOpenOracleIdenticalAnswerTaskManagerAddressRemovedFromFeedlistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenOracleIdenticalAnswerTaskManagerAddressRemovedFromFeedlistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenOracleIdenticalAnswerTaskManagerAddressRemovedFromFeedlist represents a AddressRemovedFromFeedlist event raised by the ContractOpenOracleIdenticalAnswerTaskManager contract.
type ContractOpenOracleIdenticalAnswerTaskManagerAddressRemovedFromFeedlist struct {
	Address common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddressRemovedFromFeedlist is a free log retrieval operation binding the contract event 0x28d57cd2f66bc88bf9eeda7d803691bff6e77bfdbf4212b41841a0d1b7189c5f.
//
// Solidity: event AddressRemovedFromFeedlist(address indexed _address)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) FilterAddressRemovedFromFeedlist(opts *bind.FilterOpts, _address []common.Address) (*ContractOpenOracleIdenticalAnswerTaskManagerAddressRemovedFromFeedlistIterator, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.FilterLogs(opts, "AddressRemovedFromFeedlist", _addressRule)
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleIdenticalAnswerTaskManagerAddressRemovedFromFeedlistIterator{contract: _ContractOpenOracleIdenticalAnswerTaskManager.contract, event: "AddressRemovedFromFeedlist", logs: logs, sub: sub}, nil
}

// WatchAddressRemovedFromFeedlist is a free log subscription operation binding the contract event 0x28d57cd2f66bc88bf9eeda7d803691bff6e77bfdbf4212b41841a0d1b7189c5f.
//
// Solidity: event AddressRemovedFromFeedlist(address indexed _address)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) WatchAddressRemovedFromFeedlist(opts *bind.WatchOpts, sink chan<- *ContractOpenOracleIdenticalAnswerTaskManagerAddressRemovedFromFeedlist, _address []common.Address) (event.Subscription, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.WatchLogs(opts, "AddressRemovedFromFeedlist", _addressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenOracleIdenticalAnswerTaskManagerAddressRemovedFromFeedlist)
				if err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.UnpackLog(event, "AddressRemovedFromFeedlist", log); err != nil {
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
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) ParseAddressRemovedFromFeedlist(log types.Log) (*ContractOpenOracleIdenticalAnswerTaskManagerAddressRemovedFromFeedlist, error) {
	event := new(ContractOpenOracleIdenticalAnswerTaskManagerAddressRemovedFromFeedlist)
	if err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.UnpackLog(event, "AddressRemovedFromFeedlist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOpenOracleIdenticalAnswerTaskManagerAggregatorUpdatedIterator is returned from FilterAggregatorUpdated and is used to iterate over the raw logs and unpacked data for AggregatorUpdated events raised by the ContractOpenOracleIdenticalAnswerTaskManager contract.
type ContractOpenOracleIdenticalAnswerTaskManagerAggregatorUpdatedIterator struct {
	Event *ContractOpenOracleIdenticalAnswerTaskManagerAggregatorUpdated // Event containing the contract specifics and raw log

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
func (it *ContractOpenOracleIdenticalAnswerTaskManagerAggregatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenOracleIdenticalAnswerTaskManagerAggregatorUpdated)
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
		it.Event = new(ContractOpenOracleIdenticalAnswerTaskManagerAggregatorUpdated)
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
func (it *ContractOpenOracleIdenticalAnswerTaskManagerAggregatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenOracleIdenticalAnswerTaskManagerAggregatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenOracleIdenticalAnswerTaskManagerAggregatorUpdated represents a AggregatorUpdated event raised by the ContractOpenOracleIdenticalAnswerTaskManager contract.
type ContractOpenOracleIdenticalAnswerTaskManagerAggregatorUpdated struct {
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAggregatorUpdated is a free log retrieval operation binding the contract event 0x602cec4b1583b07d071161da5eb9589444d2459201e2fab7753dc941e9351c21.
//
// Solidity: event AggregatorUpdated(address aggregator)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) FilterAggregatorUpdated(opts *bind.FilterOpts) (*ContractOpenOracleIdenticalAnswerTaskManagerAggregatorUpdatedIterator, error) {

	logs, sub, err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.FilterLogs(opts, "AggregatorUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleIdenticalAnswerTaskManagerAggregatorUpdatedIterator{contract: _ContractOpenOracleIdenticalAnswerTaskManager.contract, event: "AggregatorUpdated", logs: logs, sub: sub}, nil
}

// WatchAggregatorUpdated is a free log subscription operation binding the contract event 0x602cec4b1583b07d071161da5eb9589444d2459201e2fab7753dc941e9351c21.
//
// Solidity: event AggregatorUpdated(address aggregator)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) WatchAggregatorUpdated(opts *bind.WatchOpts, sink chan<- *ContractOpenOracleIdenticalAnswerTaskManagerAggregatorUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.WatchLogs(opts, "AggregatorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenOracleIdenticalAnswerTaskManagerAggregatorUpdated)
				if err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.UnpackLog(event, "AggregatorUpdated", log); err != nil {
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
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) ParseAggregatorUpdated(log types.Log) (*ContractOpenOracleIdenticalAnswerTaskManagerAggregatorUpdated, error) {
	event := new(ContractOpenOracleIdenticalAnswerTaskManagerAggregatorUpdated)
	if err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.UnpackLog(event, "AggregatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOpenOracleIdenticalAnswerTaskManagerFundsWithdrawnIterator is returned from FilterFundsWithdrawn and is used to iterate over the raw logs and unpacked data for FundsWithdrawn events raised by the ContractOpenOracleIdenticalAnswerTaskManager contract.
type ContractOpenOracleIdenticalAnswerTaskManagerFundsWithdrawnIterator struct {
	Event *ContractOpenOracleIdenticalAnswerTaskManagerFundsWithdrawn // Event containing the contract specifics and raw log

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
func (it *ContractOpenOracleIdenticalAnswerTaskManagerFundsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenOracleIdenticalAnswerTaskManagerFundsWithdrawn)
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
		it.Event = new(ContractOpenOracleIdenticalAnswerTaskManagerFundsWithdrawn)
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
func (it *ContractOpenOracleIdenticalAnswerTaskManagerFundsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenOracleIdenticalAnswerTaskManagerFundsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenOracleIdenticalAnswerTaskManagerFundsWithdrawn represents a FundsWithdrawn event raised by the ContractOpenOracleIdenticalAnswerTaskManager contract.
type ContractOpenOracleIdenticalAnswerTaskManagerFundsWithdrawn struct {
	TaskNum *big.Int
	Creator common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFundsWithdrawn is a free log retrieval operation binding the contract event 0xf440aec6b52895984d061d622e6edeba6210f7c3e059be920663140c084560d7.
//
// Solidity: event FundsWithdrawn(uint256 taskNum, address creator, uint256 amount)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) FilterFundsWithdrawn(opts *bind.FilterOpts) (*ContractOpenOracleIdenticalAnswerTaskManagerFundsWithdrawnIterator, error) {

	logs, sub, err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.FilterLogs(opts, "FundsWithdrawn")
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleIdenticalAnswerTaskManagerFundsWithdrawnIterator{contract: _ContractOpenOracleIdenticalAnswerTaskManager.contract, event: "FundsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchFundsWithdrawn is a free log subscription operation binding the contract event 0xf440aec6b52895984d061d622e6edeba6210f7c3e059be920663140c084560d7.
//
// Solidity: event FundsWithdrawn(uint256 taskNum, address creator, uint256 amount)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) WatchFundsWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractOpenOracleIdenticalAnswerTaskManagerFundsWithdrawn) (event.Subscription, error) {

	logs, sub, err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.WatchLogs(opts, "FundsWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenOracleIdenticalAnswerTaskManagerFundsWithdrawn)
				if err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
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
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) ParseFundsWithdrawn(log types.Log) (*ContractOpenOracleIdenticalAnswerTaskManagerFundsWithdrawn, error) {
	event := new(ContractOpenOracleIdenticalAnswerTaskManagerFundsWithdrawn)
	if err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOpenOracleIdenticalAnswerTaskManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractOpenOracleIdenticalAnswerTaskManager contract.
type ContractOpenOracleIdenticalAnswerTaskManagerInitializedIterator struct {
	Event *ContractOpenOracleIdenticalAnswerTaskManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractOpenOracleIdenticalAnswerTaskManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenOracleIdenticalAnswerTaskManagerInitialized)
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
		it.Event = new(ContractOpenOracleIdenticalAnswerTaskManagerInitialized)
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
func (it *ContractOpenOracleIdenticalAnswerTaskManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenOracleIdenticalAnswerTaskManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenOracleIdenticalAnswerTaskManagerInitialized represents a Initialized event raised by the ContractOpenOracleIdenticalAnswerTaskManager contract.
type ContractOpenOracleIdenticalAnswerTaskManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractOpenOracleIdenticalAnswerTaskManagerInitializedIterator, error) {

	logs, sub, err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleIdenticalAnswerTaskManagerInitializedIterator{contract: _ContractOpenOracleIdenticalAnswerTaskManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractOpenOracleIdenticalAnswerTaskManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenOracleIdenticalAnswerTaskManagerInitialized)
				if err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) ParseInitialized(log types.Log) (*ContractOpenOracleIdenticalAnswerTaskManagerInitialized, error) {
	event := new(ContractOpenOracleIdenticalAnswerTaskManagerInitialized)
	if err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOpenOracleIdenticalAnswerTaskManagerNewIdenticalAnswerTaskCreatedIterator is returned from FilterNewIdenticalAnswerTaskCreated and is used to iterate over the raw logs and unpacked data for NewIdenticalAnswerTaskCreated events raised by the ContractOpenOracleIdenticalAnswerTaskManager contract.
type ContractOpenOracleIdenticalAnswerTaskManagerNewIdenticalAnswerTaskCreatedIterator struct {
	Event *ContractOpenOracleIdenticalAnswerTaskManagerNewIdenticalAnswerTaskCreated // Event containing the contract specifics and raw log

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
func (it *ContractOpenOracleIdenticalAnswerTaskManagerNewIdenticalAnswerTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenOracleIdenticalAnswerTaskManagerNewIdenticalAnswerTaskCreated)
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
		it.Event = new(ContractOpenOracleIdenticalAnswerTaskManagerNewIdenticalAnswerTaskCreated)
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
func (it *ContractOpenOracleIdenticalAnswerTaskManagerNewIdenticalAnswerTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenOracleIdenticalAnswerTaskManagerNewIdenticalAnswerTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenOracleIdenticalAnswerTaskManagerNewIdenticalAnswerTaskCreated represents a NewIdenticalAnswerTaskCreated event raised by the ContractOpenOracleIdenticalAnswerTaskManager contract.
type ContractOpenOracleIdenticalAnswerTaskManagerNewIdenticalAnswerTaskCreated struct {
	TaskIndex uint32
	Task      IOpenOracleIdenticalAnswerTaskManagerTask
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewIdenticalAnswerTaskCreated is a free log retrieval operation binding the contract event 0xdc1d72dbd45105676b54bc1aa51f223c99997d23756fc461872ef079635bad40.
//
// Solidity: event NewIdenticalAnswerTaskCreated(uint32 indexed taskIndex, (uint8,bytes,uint32,uint8,uint96,address,uint256) task)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) FilterNewIdenticalAnswerTaskCreated(opts *bind.FilterOpts, taskIndex []uint32) (*ContractOpenOracleIdenticalAnswerTaskManagerNewIdenticalAnswerTaskCreatedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.FilterLogs(opts, "NewIdenticalAnswerTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleIdenticalAnswerTaskManagerNewIdenticalAnswerTaskCreatedIterator{contract: _ContractOpenOracleIdenticalAnswerTaskManager.contract, event: "NewIdenticalAnswerTaskCreated", logs: logs, sub: sub}, nil
}

// WatchNewIdenticalAnswerTaskCreated is a free log subscription operation binding the contract event 0xdc1d72dbd45105676b54bc1aa51f223c99997d23756fc461872ef079635bad40.
//
// Solidity: event NewIdenticalAnswerTaskCreated(uint32 indexed taskIndex, (uint8,bytes,uint32,uint8,uint96,address,uint256) task)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) WatchNewIdenticalAnswerTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractOpenOracleIdenticalAnswerTaskManagerNewIdenticalAnswerTaskCreated, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.WatchLogs(opts, "NewIdenticalAnswerTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenOracleIdenticalAnswerTaskManagerNewIdenticalAnswerTaskCreated)
				if err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.UnpackLog(event, "NewIdenticalAnswerTaskCreated", log); err != nil {
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

// ParseNewIdenticalAnswerTaskCreated is a log parse operation binding the contract event 0xdc1d72dbd45105676b54bc1aa51f223c99997d23756fc461872ef079635bad40.
//
// Solidity: event NewIdenticalAnswerTaskCreated(uint32 indexed taskIndex, (uint8,bytes,uint32,uint8,uint96,address,uint256) task)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) ParseNewIdenticalAnswerTaskCreated(log types.Log) (*ContractOpenOracleIdenticalAnswerTaskManagerNewIdenticalAnswerTaskCreated, error) {
	event := new(ContractOpenOracleIdenticalAnswerTaskManagerNewIdenticalAnswerTaskCreated)
	if err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.UnpackLog(event, "NewIdenticalAnswerTaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOpenOracleIdenticalAnswerTaskManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractOpenOracleIdenticalAnswerTaskManager contract.
type ContractOpenOracleIdenticalAnswerTaskManagerOwnershipTransferredIterator struct {
	Event *ContractOpenOracleIdenticalAnswerTaskManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractOpenOracleIdenticalAnswerTaskManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenOracleIdenticalAnswerTaskManagerOwnershipTransferred)
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
		it.Event = new(ContractOpenOracleIdenticalAnswerTaskManagerOwnershipTransferred)
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
func (it *ContractOpenOracleIdenticalAnswerTaskManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenOracleIdenticalAnswerTaskManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenOracleIdenticalAnswerTaskManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractOpenOracleIdenticalAnswerTaskManager contract.
type ContractOpenOracleIdenticalAnswerTaskManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOpenOracleIdenticalAnswerTaskManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleIdenticalAnswerTaskManagerOwnershipTransferredIterator{contract: _ContractOpenOracleIdenticalAnswerTaskManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOpenOracleIdenticalAnswerTaskManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenOracleIdenticalAnswerTaskManagerOwnershipTransferred)
				if err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOpenOracleIdenticalAnswerTaskManagerOwnershipTransferred, error) {
	event := new(ContractOpenOracleIdenticalAnswerTaskManagerOwnershipTransferred)
	if err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOpenOracleIdenticalAnswerTaskManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractOpenOracleIdenticalAnswerTaskManager contract.
type ContractOpenOracleIdenticalAnswerTaskManagerPausedIterator struct {
	Event *ContractOpenOracleIdenticalAnswerTaskManagerPaused // Event containing the contract specifics and raw log

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
func (it *ContractOpenOracleIdenticalAnswerTaskManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenOracleIdenticalAnswerTaskManagerPaused)
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
		it.Event = new(ContractOpenOracleIdenticalAnswerTaskManagerPaused)
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
func (it *ContractOpenOracleIdenticalAnswerTaskManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenOracleIdenticalAnswerTaskManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenOracleIdenticalAnswerTaskManagerPaused represents a Paused event raised by the ContractOpenOracleIdenticalAnswerTaskManager contract.
type ContractOpenOracleIdenticalAnswerTaskManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractOpenOracleIdenticalAnswerTaskManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleIdenticalAnswerTaskManagerPausedIterator{contract: _ContractOpenOracleIdenticalAnswerTaskManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractOpenOracleIdenticalAnswerTaskManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenOracleIdenticalAnswerTaskManagerPaused)
				if err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) ParsePaused(log types.Log) (*ContractOpenOracleIdenticalAnswerTaskManagerPaused, error) {
	event := new(ContractOpenOracleIdenticalAnswerTaskManagerPaused)
	if err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOpenOracleIdenticalAnswerTaskManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractOpenOracleIdenticalAnswerTaskManager contract.
type ContractOpenOracleIdenticalAnswerTaskManagerPauserRegistrySetIterator struct {
	Event *ContractOpenOracleIdenticalAnswerTaskManagerPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *ContractOpenOracleIdenticalAnswerTaskManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenOracleIdenticalAnswerTaskManagerPauserRegistrySet)
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
		it.Event = new(ContractOpenOracleIdenticalAnswerTaskManagerPauserRegistrySet)
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
func (it *ContractOpenOracleIdenticalAnswerTaskManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenOracleIdenticalAnswerTaskManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenOracleIdenticalAnswerTaskManagerPauserRegistrySet represents a PauserRegistrySet event raised by the ContractOpenOracleIdenticalAnswerTaskManager contract.
type ContractOpenOracleIdenticalAnswerTaskManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractOpenOracleIdenticalAnswerTaskManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleIdenticalAnswerTaskManagerPauserRegistrySetIterator{contract: _ContractOpenOracleIdenticalAnswerTaskManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractOpenOracleIdenticalAnswerTaskManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenOracleIdenticalAnswerTaskManagerPauserRegistrySet)
				if err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) ParsePauserRegistrySet(log types.Log) (*ContractOpenOracleIdenticalAnswerTaskManagerPauserRegistrySet, error) {
	event := new(ContractOpenOracleIdenticalAnswerTaskManagerPauserRegistrySet)
	if err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOpenOracleIdenticalAnswerTaskManagerTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the ContractOpenOracleIdenticalAnswerTaskManager contract.
type ContractOpenOracleIdenticalAnswerTaskManagerTaskCompletedIterator struct {
	Event *ContractOpenOracleIdenticalAnswerTaskManagerTaskCompleted // Event containing the contract specifics and raw log

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
func (it *ContractOpenOracleIdenticalAnswerTaskManagerTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenOracleIdenticalAnswerTaskManagerTaskCompleted)
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
		it.Event = new(ContractOpenOracleIdenticalAnswerTaskManagerTaskCompleted)
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
func (it *ContractOpenOracleIdenticalAnswerTaskManagerTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenOracleIdenticalAnswerTaskManagerTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenOracleIdenticalAnswerTaskManagerTaskCompleted represents a TaskCompleted event raised by the ContractOpenOracleIdenticalAnswerTaskManager contract.
type ContractOpenOracleIdenticalAnswerTaskManagerTaskCompleted struct {
	TaskIndex uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskIndex []uint32) (*ContractOpenOracleIdenticalAnswerTaskManagerTaskCompletedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.FilterLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleIdenticalAnswerTaskManagerTaskCompletedIterator{contract: _ContractOpenOracleIdenticalAnswerTaskManager.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *ContractOpenOracleIdenticalAnswerTaskManagerTaskCompleted, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.WatchLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenOracleIdenticalAnswerTaskManagerTaskCompleted)
				if err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
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
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) ParseTaskCompleted(log types.Log) (*ContractOpenOracleIdenticalAnswerTaskManagerTaskCompleted, error) {
	event := new(ContractOpenOracleIdenticalAnswerTaskManagerTaskCompleted)
	if err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOpenOracleIdenticalAnswerTaskManagerTaskRespondedIterator is returned from FilterTaskResponded and is used to iterate over the raw logs and unpacked data for TaskResponded events raised by the ContractOpenOracleIdenticalAnswerTaskManager contract.
type ContractOpenOracleIdenticalAnswerTaskManagerTaskRespondedIterator struct {
	Event *ContractOpenOracleIdenticalAnswerTaskManagerTaskResponded // Event containing the contract specifics and raw log

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
func (it *ContractOpenOracleIdenticalAnswerTaskManagerTaskRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenOracleIdenticalAnswerTaskManagerTaskResponded)
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
		it.Event = new(ContractOpenOracleIdenticalAnswerTaskManagerTaskResponded)
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
func (it *ContractOpenOracleIdenticalAnswerTaskManagerTaskRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenOracleIdenticalAnswerTaskManagerTaskRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenOracleIdenticalAnswerTaskManagerTaskResponded represents a TaskResponded event raised by the ContractOpenOracleIdenticalAnswerTaskManager contract.
type ContractOpenOracleIdenticalAnswerTaskManagerTaskResponded struct {
	Task                 IOpenOracleIdenticalAnswerTaskManagerTask
	TaskResponse         IOpenOracleIdenticalAnswerTaskManagerAggregatedTaskResponse
	TaskResponseMetadata IOpenOracleIdenticalAnswerTaskManagerTaskResponseMetadata
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTaskResponded is a free log retrieval operation binding the contract event 0x5859b9b29fa1ffdbe40fa8bb3cdf904a8c0b31eef57809c082fd5549f3df4005.
//
// Solidity: event TaskResponded((uint8,bytes,uint32,uint8,uint96,address,uint256) task, ((uint32,bytes),uint256,(uint256,uint256),(uint256[2],uint256[2])) taskResponse, (uint32) taskResponseMetadata)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) FilterTaskResponded(opts *bind.FilterOpts) (*ContractOpenOracleIdenticalAnswerTaskManagerTaskRespondedIterator, error) {

	logs, sub, err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.FilterLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleIdenticalAnswerTaskManagerTaskRespondedIterator{contract: _ContractOpenOracleIdenticalAnswerTaskManager.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0x5859b9b29fa1ffdbe40fa8bb3cdf904a8c0b31eef57809c082fd5549f3df4005.
//
// Solidity: event TaskResponded((uint8,bytes,uint32,uint8,uint96,address,uint256) task, ((uint32,bytes),uint256,(uint256,uint256),(uint256[2],uint256[2])) taskResponse, (uint32) taskResponseMetadata)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) WatchTaskResponded(opts *bind.WatchOpts, sink chan<- *ContractOpenOracleIdenticalAnswerTaskManagerTaskResponded) (event.Subscription, error) {

	logs, sub, err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.WatchLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenOracleIdenticalAnswerTaskManagerTaskResponded)
				if err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
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

// ParseTaskResponded is a log parse operation binding the contract event 0x5859b9b29fa1ffdbe40fa8bb3cdf904a8c0b31eef57809c082fd5549f3df4005.
//
// Solidity: event TaskResponded((uint8,bytes,uint32,uint8,uint96,address,uint256) task, ((uint32,bytes),uint256,(uint256,uint256),(uint256[2],uint256[2])) taskResponse, (uint32) taskResponseMetadata)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) ParseTaskResponded(log types.Log) (*ContractOpenOracleIdenticalAnswerTaskManagerTaskResponded, error) {
	event := new(ContractOpenOracleIdenticalAnswerTaskManagerTaskResponded)
	if err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOpenOracleIdenticalAnswerTaskManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractOpenOracleIdenticalAnswerTaskManager contract.
type ContractOpenOracleIdenticalAnswerTaskManagerUnpausedIterator struct {
	Event *ContractOpenOracleIdenticalAnswerTaskManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractOpenOracleIdenticalAnswerTaskManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenOracleIdenticalAnswerTaskManagerUnpaused)
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
		it.Event = new(ContractOpenOracleIdenticalAnswerTaskManagerUnpaused)
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
func (it *ContractOpenOracleIdenticalAnswerTaskManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenOracleIdenticalAnswerTaskManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenOracleIdenticalAnswerTaskManagerUnpaused represents a Unpaused event raised by the ContractOpenOracleIdenticalAnswerTaskManager contract.
type ContractOpenOracleIdenticalAnswerTaskManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractOpenOracleIdenticalAnswerTaskManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleIdenticalAnswerTaskManagerUnpausedIterator{contract: _ContractOpenOracleIdenticalAnswerTaskManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractOpenOracleIdenticalAnswerTaskManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenOracleIdenticalAnswerTaskManagerUnpaused)
				if err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ContractOpenOracleIdenticalAnswerTaskManager *ContractOpenOracleIdenticalAnswerTaskManagerFilterer) ParseUnpaused(log types.Log) (*ContractOpenOracleIdenticalAnswerTaskManagerUnpaused, error) {
	event := new(ContractOpenOracleIdenticalAnswerTaskManagerUnpaused)
	if err := _ContractOpenOracleIdenticalAnswerTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractOpenOracleServiceManager

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

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// ContractOpenOracleServiceManagerMetaData contains all meta data concerning the ContractOpenOracleServiceManager contract.
var ContractOpenOracleServiceManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_delegationManager\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"},{\"name\":\"_openOracleTaskManager\",\"type\":\"address\",\"internalType\":\"contractIOpenOracleTaskManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"freezeOperator\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"openOracleTaskManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOpenOracleTaskManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x6101006040523480156200001257600080fd5b50604051620016d0380380620016d083398101604081905262000035916200014f565b6001600160a01b0380851660a052808416608052821660c0528383836200005b62000074565b5050506001600160a01b031660e05250620001b7915050565b600054610100900460ff1615620000e15760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116101562000134576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200014c57600080fd5b50565b600080600080608085870312156200016657600080fd5b8451620001738162000136565b6020860151909450620001868162000136565b6040860151909350620001998162000136565b6060860151909250620001ac8162000136565b939692955090935050565b60805160a05160c05160e05161147e620002526000396000818160f101526106860152600081816103a5015281816105010152818161059801528181610afb01528181610c7f0152610d1e0152600081816107540152818161081d01526108f10152600081816101d00152818161025f015281816102df015281816107c90152818161089501528181610a390152610bda015261147e6000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80638da5cb5b116100715780638da5cb5b146101465780639926ee7d14610157578063a364f4da1461016a578063c4d66de81461017d578063e481af9d14610190578063f2fde38b1461019857600080fd5b806333cfb7b7146100ae57806338c8ee64146100d75780633bb8f679146100ec578063715018a61461012b578063750521f514610133575b600080fd5b6100c16100bc366004610f91565b6101ab565b6040516100ce9190610fb5565b60405180910390f35b6100ea6100e5366004610f91565b61067b565b005b6101137f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100ce565b6100ea610721565b6100ea6101413660046110b7565b610735565b6033546001600160a01b0316610113565b6100ea610165366004611108565b6107be565b6100ea610178366004610f91565b61088a565b6100ea61018b366004610f91565b610920565b6100c1610a33565b6100ea6101a6366004610f91565b610dfd565b6040516309aa152760e11b81526001600160a01b0382811660048301526060916000917f000000000000000000000000000000000000000000000000000000000000000016906313542a4e90602401602060405180830381865afa158015610217573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061023b91906111b3565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa1580156102a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102ca91906111cc565b90506001600160c01b038116158061036457507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561033b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061035f91906111f5565b60ff16155b1561038057505060408051600081526020810190915292915050565b6000610394826001600160c01b0316610e73565b90506000805b825181101561046a577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f58483815181106103e4576103e4611218565b01602001516040516001600160e01b031960e084901b16815260f89190911c6004820152602401602060405180830381865afa158015610428573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061044c91906111b3565b6104569083611244565b9150806104628161125c565b91505061039a565b5060008167ffffffffffffffff81111561048657610486611002565b6040519080825280602002602001820160405280156104af578160200160208202803683370190505b5090506000805b845181101561066e5760008582815181106104d3576104d3611218565b0160200151604051633ca5a5f560e01b815260f89190911c6004820181905291506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590602401602060405180830381865afa158015610548573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061056c91906111b3565b905060005b81811015610658576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa1580156105e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061060a9190611277565b6000015186868151811061062057610620611218565b6001600160a01b0390921660209283029190910190910152846106428161125c565b95505080806106509061125c565b915050610571565b50505080806106669061125c565b9150506104b6565b5090979650505050505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461071e5760405162461bcd60e51b815260206004820152603c60248201527f6f6e6c794f70656e4f7261636c655461736b4d616e616765723a206e6f74206660448201527f726f6d206f70656e206f7261636c65207461736b206d616e616765720000000060648201526084015b60405180910390fd5b50565b610729610ed0565b6107336000610f2a565b565b61073d610ed0565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb35590610789908490600401611343565b600060405180830381600087803b1580156107a357600080fd5b505af11580156107b7573d6000803e3d6000fd5b5050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146108065760405162461bcd60e51b815260040161071590611356565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d9061085490859085906004016113ce565b600060405180830381600087803b15801561086e57600080fd5b505af1158015610882573d6000803e3d6000fd5b505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146108d25760405162461bcd60e51b815260040161071590611356565b6040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da90602401610789565b600054610100900460ff16158080156109405750600054600160ff909116105b8061095a5750303b15801561095a575060005460ff166001145b6109bd5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610715565b6000805460ff1916600117905580156109e0576000805461ff0019166101001790555b6109e982610f2a565b8015610a2f576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a95573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ab991906111f5565b60ff16905080610ad757505060408051600081526020810190915290565b6000805b82811015610b8c57604051633ca5a5f560e01b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610b4a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b6e91906111b3565b610b789083611244565b915080610b848161125c565b915050610adb565b5060008167ffffffffffffffff811115610ba857610ba8611002565b604051908082528060200260200182016040528015610bd1578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c36573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c5a91906111f5565b60ff16811015610df357604051633ca5a5f560e01b815260ff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610cce573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cf291906111b3565b905060005b81811015610dde576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015610d6c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d909190611277565b60000151858581518110610da657610da6611218565b6001600160a01b039092166020928302919091019091015283610dc88161125c565b9450508080610dd69061125c565b915050610cf7565b50508080610deb9061125c565b915050610bd8565b5090949350505050565b610e05610ed0565b6001600160a01b038116610e6a5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610715565b61071e81610f2a565b60606000805b610100811015610ec9576001811b915083821615610eb957828160f81b604051602001610ea7929190611419565b60405160208183030381529060405292505b610ec28161125c565b9050610e79565b5050919050565b6033546001600160a01b031633146107335760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610715565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600160a01b038116811461071e57600080fd5b600060208284031215610fa357600080fd5b8135610fae81610f7c565b9392505050565b6020808252825182820181905260009190848201906040850190845b81811015610ff65783516001600160a01b031683529284019291840191600101610fd1565b50909695505050505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff8111828210171561103b5761103b611002565b60405290565b600067ffffffffffffffff8084111561105c5761105c611002565b604051601f8501601f19908116603f0116810190828211818310171561108457611084611002565b8160405280935085815286868601111561109d57600080fd5b858560208301376000602087830101525050509392505050565b6000602082840312156110c957600080fd5b813567ffffffffffffffff8111156110e057600080fd5b8201601f810184136110f157600080fd5b61110084823560208401611041565b949350505050565b6000806040838503121561111b57600080fd5b823561112681610f7c565b9150602083013567ffffffffffffffff8082111561114357600080fd5b908401906060828703121561115757600080fd5b61115f611018565b82358281111561116e57600080fd5b83019150601f8201871361118157600080fd5b61119087833560208501611041565b815260208301356020820152604083013560408201528093505050509250929050565b6000602082840312156111c557600080fd5b5051919050565b6000602082840312156111de57600080fd5b81516001600160c01b0381168114610fae57600080fd5b60006020828403121561120757600080fd5b815160ff81168114610fae57600080fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600082198211156112575761125761122e565b500190565b60006000198214156112705761127061122e565b5060010190565b60006040828403121561128957600080fd5b6040516040810181811067ffffffffffffffff821117156112ac576112ac611002565b60405282516112ba81610f7c565b815260208301516bffffffffffffffffffffffff811681146112db57600080fd5b60208201529392505050565b60005b838110156113025781810151838201526020016112ea565b83811115611311576000848401525b50505050565b6000815180845261132f8160208601602086016112e7565b601f01601f19169290920160200192915050565b602081526000610fae6020830184611317565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b60018060a01b03831681526040602082015260008251606060408401526113f860a0840182611317565b90506020840151606084015260408401516080840152809150509392505050565b6000835161142b8184602088016112e7565b6001600160f81b031993909316919092019081526001019291505056fea2646970667358221220ae3c592d5ccb356ab9542c3f76e7236e261c34484dff102d348a6e8bfb803bf464736f6c634300080c0033",
}

// ContractOpenOracleServiceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractOpenOracleServiceManagerMetaData.ABI instead.
var ContractOpenOracleServiceManagerABI = ContractOpenOracleServiceManagerMetaData.ABI

// ContractOpenOracleServiceManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractOpenOracleServiceManagerMetaData.Bin instead.
var ContractOpenOracleServiceManagerBin = ContractOpenOracleServiceManagerMetaData.Bin

// DeployContractOpenOracleServiceManager deploys a new Ethereum contract, binding an instance of ContractOpenOracleServiceManager to it.
func DeployContractOpenOracleServiceManager(auth *bind.TransactOpts, backend bind.ContractBackend, _delegationManager common.Address, _registryCoordinator common.Address, _stakeRegistry common.Address, _openOracleTaskManager common.Address) (common.Address, *types.Transaction, *ContractOpenOracleServiceManager, error) {
	parsed, err := ContractOpenOracleServiceManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractOpenOracleServiceManagerBin), backend, _delegationManager, _registryCoordinator, _stakeRegistry, _openOracleTaskManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractOpenOracleServiceManager{ContractOpenOracleServiceManagerCaller: ContractOpenOracleServiceManagerCaller{contract: contract}, ContractOpenOracleServiceManagerTransactor: ContractOpenOracleServiceManagerTransactor{contract: contract}, ContractOpenOracleServiceManagerFilterer: ContractOpenOracleServiceManagerFilterer{contract: contract}}, nil
}

// ContractOpenOracleServiceManager is an auto generated Go binding around an Ethereum contract.
type ContractOpenOracleServiceManager struct {
	ContractOpenOracleServiceManagerCaller     // Read-only binding to the contract
	ContractOpenOracleServiceManagerTransactor // Write-only binding to the contract
	ContractOpenOracleServiceManagerFilterer   // Log filterer for contract events
}

// ContractOpenOracleServiceManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractOpenOracleServiceManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOpenOracleServiceManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractOpenOracleServiceManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOpenOracleServiceManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractOpenOracleServiceManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOpenOracleServiceManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractOpenOracleServiceManagerSession struct {
	Contract     *ContractOpenOracleServiceManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                     // Call options to use throughout this session
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ContractOpenOracleServiceManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractOpenOracleServiceManagerCallerSession struct {
	Contract *ContractOpenOracleServiceManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                           // Call options to use throughout this session
}

// ContractOpenOracleServiceManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractOpenOracleServiceManagerTransactorSession struct {
	Contract     *ContractOpenOracleServiceManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                           // Transaction auth options to use throughout this session
}

// ContractOpenOracleServiceManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractOpenOracleServiceManagerRaw struct {
	Contract *ContractOpenOracleServiceManager // Generic contract binding to access the raw methods on
}

// ContractOpenOracleServiceManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractOpenOracleServiceManagerCallerRaw struct {
	Contract *ContractOpenOracleServiceManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractOpenOracleServiceManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractOpenOracleServiceManagerTransactorRaw struct {
	Contract *ContractOpenOracleServiceManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractOpenOracleServiceManager creates a new instance of ContractOpenOracleServiceManager, bound to a specific deployed contract.
func NewContractOpenOracleServiceManager(address common.Address, backend bind.ContractBackend) (*ContractOpenOracleServiceManager, error) {
	contract, err := bindContractOpenOracleServiceManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleServiceManager{ContractOpenOracleServiceManagerCaller: ContractOpenOracleServiceManagerCaller{contract: contract}, ContractOpenOracleServiceManagerTransactor: ContractOpenOracleServiceManagerTransactor{contract: contract}, ContractOpenOracleServiceManagerFilterer: ContractOpenOracleServiceManagerFilterer{contract: contract}}, nil
}

// NewContractOpenOracleServiceManagerCaller creates a new read-only instance of ContractOpenOracleServiceManager, bound to a specific deployed contract.
func NewContractOpenOracleServiceManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractOpenOracleServiceManagerCaller, error) {
	contract, err := bindContractOpenOracleServiceManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleServiceManagerCaller{contract: contract}, nil
}

// NewContractOpenOracleServiceManagerTransactor creates a new write-only instance of ContractOpenOracleServiceManager, bound to a specific deployed contract.
func NewContractOpenOracleServiceManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractOpenOracleServiceManagerTransactor, error) {
	contract, err := bindContractOpenOracleServiceManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleServiceManagerTransactor{contract: contract}, nil
}

// NewContractOpenOracleServiceManagerFilterer creates a new log filterer instance of ContractOpenOracleServiceManager, bound to a specific deployed contract.
func NewContractOpenOracleServiceManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractOpenOracleServiceManagerFilterer, error) {
	contract, err := bindContractOpenOracleServiceManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleServiceManagerFilterer{contract: contract}, nil
}

// bindContractOpenOracleServiceManager binds a generic wrapper to an already deployed contract.
func bindContractOpenOracleServiceManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractOpenOracleServiceManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractOpenOracleServiceManager.Contract.ContractOpenOracleServiceManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.Contract.ContractOpenOracleServiceManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.Contract.ContractOpenOracleServiceManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractOpenOracleServiceManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.Contract.contract.Transact(opts, method, params...)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ContractOpenOracleServiceManager.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractOpenOracleServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractOpenOracleServiceManager.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractOpenOracleServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractOpenOracleServiceManager.CallOpts, operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractOpenOracleServiceManager.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractOpenOracleServiceManager.Contract.GetRestakeableStrategies(&_ContractOpenOracleServiceManager.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractOpenOracleServiceManager.Contract.GetRestakeableStrategies(&_ContractOpenOracleServiceManager.CallOpts)
}

// OpenOracleTaskManager is a free data retrieval call binding the contract method 0x3bb8f679.
//
// Solidity: function openOracleTaskManager() view returns(address)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerCaller) OpenOracleTaskManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOpenOracleServiceManager.contract.Call(opts, &out, "openOracleTaskManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OpenOracleTaskManager is a free data retrieval call binding the contract method 0x3bb8f679.
//
// Solidity: function openOracleTaskManager() view returns(address)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerSession) OpenOracleTaskManager() (common.Address, error) {
	return _ContractOpenOracleServiceManager.Contract.OpenOracleTaskManager(&_ContractOpenOracleServiceManager.CallOpts)
}

// OpenOracleTaskManager is a free data retrieval call binding the contract method 0x3bb8f679.
//
// Solidity: function openOracleTaskManager() view returns(address)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerCallerSession) OpenOracleTaskManager() (common.Address, error) {
	return _ContractOpenOracleServiceManager.Contract.OpenOracleTaskManager(&_ContractOpenOracleServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOpenOracleServiceManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerSession) Owner() (common.Address, error) {
	return _ContractOpenOracleServiceManager.Contract.Owner(&_ContractOpenOracleServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerCallerSession) Owner() (common.Address, error) {
	return _ContractOpenOracleServiceManager.Contract.Owner(&_ContractOpenOracleServiceManager.CallOpts)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractOpenOracleServiceManager.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractOpenOracleServiceManager.TransactOpts, operator)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerTransactor) FreezeOperator(opts *bind.TransactOpts, operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.contract.Transact(opts, "freezeOperator", operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.Contract.FreezeOperator(&_ContractOpenOracleServiceManager.TransactOpts, operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerTransactorSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.Contract.FreezeOperator(&_ContractOpenOracleServiceManager.TransactOpts, operatorAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.Contract.Initialize(&_ContractOpenOracleServiceManager.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.Contract.Initialize(&_ContractOpenOracleServiceManager.TransactOpts, initialOwner)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.Contract.RegisterOperatorToAVS(&_ContractOpenOracleServiceManager.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.Contract.RegisterOperatorToAVS(&_ContractOpenOracleServiceManager.TransactOpts, operator, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.Contract.RenounceOwnership(&_ContractOpenOracleServiceManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.Contract.RenounceOwnership(&_ContractOpenOracleServiceManager.TransactOpts)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string _metadataURI) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerTransactor) SetMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.contract.Transact(opts, "setMetadataURI", _metadataURI)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string _metadataURI) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerSession) SetMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.Contract.SetMetadataURI(&_ContractOpenOracleServiceManager.TransactOpts, _metadataURI)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string _metadataURI) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerTransactorSession) SetMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.Contract.SetMetadataURI(&_ContractOpenOracleServiceManager.TransactOpts, _metadataURI)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.Contract.TransferOwnership(&_ContractOpenOracleServiceManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.Contract.TransferOwnership(&_ContractOpenOracleServiceManager.TransactOpts, newOwner)
}

// ContractOpenOracleServiceManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractOpenOracleServiceManager contract.
type ContractOpenOracleServiceManagerInitializedIterator struct {
	Event *ContractOpenOracleServiceManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractOpenOracleServiceManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenOracleServiceManagerInitialized)
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
		it.Event = new(ContractOpenOracleServiceManagerInitialized)
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
func (it *ContractOpenOracleServiceManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenOracleServiceManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenOracleServiceManagerInitialized represents a Initialized event raised by the ContractOpenOracleServiceManager contract.
type ContractOpenOracleServiceManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractOpenOracleServiceManagerInitializedIterator, error) {

	logs, sub, err := _ContractOpenOracleServiceManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleServiceManagerInitializedIterator{contract: _ContractOpenOracleServiceManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractOpenOracleServiceManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractOpenOracleServiceManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenOracleServiceManagerInitialized)
				if err := _ContractOpenOracleServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerFilterer) ParseInitialized(log types.Log) (*ContractOpenOracleServiceManagerInitialized, error) {
	event := new(ContractOpenOracleServiceManagerInitialized)
	if err := _ContractOpenOracleServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOpenOracleServiceManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractOpenOracleServiceManager contract.
type ContractOpenOracleServiceManagerOwnershipTransferredIterator struct {
	Event *ContractOpenOracleServiceManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractOpenOracleServiceManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenOracleServiceManagerOwnershipTransferred)
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
		it.Event = new(ContractOpenOracleServiceManagerOwnershipTransferred)
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
func (it *ContractOpenOracleServiceManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenOracleServiceManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenOracleServiceManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractOpenOracleServiceManager contract.
type ContractOpenOracleServiceManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOpenOracleServiceManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractOpenOracleServiceManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleServiceManagerOwnershipTransferredIterator{contract: _ContractOpenOracleServiceManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOpenOracleServiceManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractOpenOracleServiceManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenOracleServiceManagerOwnershipTransferred)
				if err := _ContractOpenOracleServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOpenOracleServiceManagerOwnershipTransferred, error) {
	event := new(ContractOpenOracleServiceManagerOwnershipTransferred)
	if err := _ContractOpenOracleServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

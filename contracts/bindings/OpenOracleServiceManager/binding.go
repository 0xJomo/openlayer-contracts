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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"},{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"},{\"name\":\"_openOracleTaskManager\",\"type\":\"address\",\"internalType\":\"contractIOpenOracleTaskManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addTaskManager\",\"inputs\":[{\"name\":\"chainName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"taskManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"freezeOperator\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"openOracleTaskManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOpenOracleTaskManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeTaskManager\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"taskManagerCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskManagers\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"chainName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"taskManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskManagerAdded\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"chainName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"taskManagerAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskManagerRemoved\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101006040523480156200001257600080fd5b5060405162001ccf38038062001ccf83398101604081905262000035916200014f565b6001600160a01b0380851660c052808416608052821660a0528383836200005b62000074565b5050506001600160a01b031660e05250620001b7915050565b600054610100900460ff1615620000e15760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116101562000134576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200014c57600080fd5b50565b600080600080608085870312156200016657600080fd5b8451620001738162000136565b6020860151909450620001868162000136565b6040860151909350620001998162000136565b6060860151909250620001ac8162000136565b939692955090935050565b60805160a05160c05160e051611a756200025a6000396000818161012b015261080801526000818161015c01528181610924015281816109f00152610a79015260008181610522015281816106780152818161071a01528181610d6d01528181610ef20152610f9401526000818161032d015281816103dc0152818161045c015281816108d00152818161099c01528181610cac0152610e4e0152611a756000f3fe608060405234801561001057600080fd5b50600436106100d05760003560e01c806332f21894146100d557806333cfb7b7146100ea57806338c8ee64146101135780633bb8f679146101265780636b3aa72e1461015a578063715018a6146101805780637b5b3049146101885780638da5cb5b1461019f5780639926ee7d146101a7578063a364f4da146101ba578063a98fb355146101cd578063c4d66de8146101e0578063ca94cfe0146101f3578063e481af9d14610206578063f2fde38b1461020e578063fd69a60f14610221575b600080fd5b6100e86100e336600461153a565b610243565b005b6100fd6100f836600461158b565b610327565b60405161010a91906115a8565b60405180910390f35b6100e861012136600461158b565b6107fd565b61014d7f000000000000000000000000000000000000000000000000000000000000000081565b60405161010a91906115f5565b7f000000000000000000000000000000000000000000000000000000000000000061014d565b6100e86108a2565b61019160985481565b60405190815260200161010a565b61014d6108b6565b6100e86101b5366004611609565b6108c5565b6100e86101c836600461158b565b610991565b6100e86101db3660046116b3565b610a5a565b6100e86101ee36600461158b565b610aae565b6100e86102013660046116ef565b610bc6565b6100fd610ca6565b6100e861021c36600461158b565b611073565b61023461022f3660046116ef565b6110e9565b60405161010a93929190611755565b61024b6111a3565b609880546000918261025c8361179e565b90915550604080516060810182528581526001600160a01b0385166020808301919091526001828401526000848152609782529290922081518051949550919390926102ac9284929101906113b2565b506020820151600190910180546040938401511515600160a01b026001600160a81b03199091166001600160a01b0390931692909217919091179055517f661e3d4c4a13612fad1504b96725fcedfb885ff3ac490c18004f57ad9f4e87789061031a908390869086906117b9565b60405180910390a1505050565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166313542a4e846040518263ffffffff1660e01b815260040161037791906115f5565b602060405180830381865afa158015610394573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103b891906117ea565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa158015610423573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104479190611803565b90506001600160c01b03811615806104e157507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104b8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104dc919061182c565b60ff16155b156104fd57505060408051600081526020810190915292915050565b6000610511826001600160c01b0316611202565b90506000805b82518110156105e9577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f58483815181106105615761056161184f565b01602001516040516001600160e01b031960e084901b16815261058a9160f81c90600401611865565b602060405180830381865afa1580156105a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105cb91906117ea565b6105d59083611873565b9150806105e18161179e565b915050610517565b506000816001600160401b038111156106045761060461144b565b60405190808252806020026020018201604052801561062d578160200160208202803683370190505b5090506000805b84518110156107f05760008582815181106106515761065161184f565b0160200151604051633ca5a5f560e01b815260f89190911c91506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f5906106ad908590600401611865565b602060405180830381865afa1580156106ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106ee91906117ea565b905060005b818110156107da576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015610768573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061078c919061188b565b600001518686815181106107a2576107a261184f565b6001600160a01b0390921660209283029190910190910152846107c48161179e565b95505080806107d29061179e565b9150506106f3565b50505080806107e89061179e565b915050610634565b5090979650505050505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461089f5760405162461bcd60e51b815260206004820152603c60248201527f6f6e6c794f70656e4f7261636c655461736b4d616e616765723a206e6f74206660448201527b3937b69037b832b71037b930b1b632903a30b9b59036b0b730b3b2b960211b60648201526084015b60405180910390fd5b50565b6108aa6111a3565b6108b460006112c4565b565b6033546001600160a01b031690565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461090d5760405162461bcd60e51b8152600401610896906118f5565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d9061095b908590859060040161196d565b600060405180830381600087803b15801561097557600080fd5b505af1158015610989573d6000803e3d6000fd5b505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146109d95760405162461bcd60e51b8152600401610896906118f5565b6040516351b27a6d60e11b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da90610a259084906004016115f5565b600060405180830381600087803b158015610a3f57600080fd5b505af1158015610a53573d6000803e3d6000fd5b5050505050565b610a626111a3565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb35590610a259084906004016119b8565b600054610100900460ff1615808015610ace5750600054600160ff909116105b80610ae85750303b158015610ae8575060005460ff166001145b610b4b5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610896565b6000805460ff191660011790558015610b6e576000805461ff0019166101001790555b610b7782611316565b8015610bc2576000805461ff00191690556040517f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890610bb990600190611865565b60405180910390a15b5050565b610bce6111a3565b600081815260976020526040902060010154600160a01b900460ff16610c4d5760405162461bcd60e51b815260206004820152602e60248201527f5461736b204d616e6167657220646f6573206e6f74206578697374206f72206160448201526d1b1c9958591e481c995b5bdd995960921b6064820152608401610896565b60008181526097602052604090819020600101805460ff60a01b19169055517f191603ba9e3b1261ce14bffd91013f2cc8137f0292821bae13360bfc0e21d44c90610c9b9083815260200190565b60405180910390a150565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d08573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d2c919061182c565b60ff16905080610d4a57505060408051600081526020810190915290565b6000805b82811015610e0157604051633ca5a5f560e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590610da2908490600401611865565b602060405180830381865afa158015610dbf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610de391906117ea565b610ded9083611873565b915080610df98161179e565b915050610d4e565b506000816001600160401b03811115610e1c57610e1c61144b565b604051908082528060200260200182016040528015610e45578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610eaa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ece919061182c565b60ff1681101561106957604051633ca5a5f560e01b81526000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590610f27908590600401611865565b602060405180830381865afa158015610f44573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f6891906117ea565b905060005b81811015611054576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015610fe2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611006919061188b565b6000015185858151811061101c5761101c61184f565b6001600160a01b03909216602092830291909101909101528361103e8161179e565b945050808061104c9061179e565b915050610f6d565b505080806110619061179e565b915050610e4c565b5090949350505050565b61107b6111a3565b6001600160a01b0381166110e05760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610896565b61089f816112c4565b609760205260009081526040902080548190611104906119cb565b80601f0160208091040260200160405190810160405280929190818152602001828054611130906119cb565b801561117d5780601f106111525761010080835404028352916020019161117d565b820191906000526020600020905b81548152906001019060200180831161116057829003601f168201915b505050600190930154919250506001600160a01b0381169060ff600160a01b9091041683565b336111ac6108b6565b6001600160a01b0316146108b45760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610896565b606060008061121084611381565b61ffff166001600160401b0381111561122b5761122b61144b565b6040519080825280601f01601f191660200182016040528015611255576020820181803683370190505b5090506000805b82518210801561126d575061010081105b15611069576001811b9350858416156112b4578060f81b8383815181106112965761129661184f565b60200101906001600160f81b031916908160001a9053508160010191505b6112bd8161179e565b905061125c565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff166110e05760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608401610896565b6000805b82156113ac57611396600184611a06565b90921691806113a481611a1d565b915050611385565b92915050565b8280546113be906119cb565b90600052602060002090601f0160209004810192826113e05760008555611426565b82601f106113f957805160ff1916838001178555611426565b82800160010185558215611426579182015b8281111561142657825182559160200191906001019061140b565b50611432929150611436565b5090565b5b808211156114325760008155600101611437565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b03811182821017156114835761148361144b565b60405290565b60006001600160401b03808411156114a3576114a361144b565b604051601f8501601f19908116603f011681019082821181831017156114cb576114cb61144b565b816040528093508581528686860111156114e457600080fd5b858560208301376000602087830101525050509392505050565b600082601f83011261150f57600080fd5b61151e83833560208501611489565b9392505050565b6001600160a01b038116811461089f57600080fd5b6000806040838503121561154d57600080fd5b82356001600160401b0381111561156357600080fd5b61156f858286016114fe565b925050602083013561158081611525565b809150509250929050565b60006020828403121561159d57600080fd5b813561151e81611525565b6020808252825182820181905260009190848201906040850190845b818110156115e95783516001600160a01b0316835292840192918401916001016115c4565b50909695505050505050565b6001600160a01b0391909116815260200190565b6000806040838503121561161c57600080fd5b823561162781611525565b915060208301356001600160401b038082111561164357600080fd5b908401906060828703121561165757600080fd5b61165f611461565b82358281111561166e57600080fd5b83019150601f8201871361168157600080fd5b61169087833560208501611489565b815260208301356020820152604083013560408201528093505050509250929050565b6000602082840312156116c557600080fd5b81356001600160401b038111156116db57600080fd5b6116e7848285016114fe565b949350505050565b60006020828403121561170157600080fd5b5035919050565b6000815180845260005b8181101561172e57602081850181015186830182015201611712565b81811115611740576000602083870101525b50601f01601f19169290920160200192915050565b6060815260006117686060830186611708565b6001600160a01b0394909416602083015250901515604090910152919050565b634e487b7160e01b600052601160045260246000fd5b60006000198214156117b2576117b2611788565b5060010190565b8381526060602082015260006117d26060830185611708565b905060018060a01b0383166040830152949350505050565b6000602082840312156117fc57600080fd5b5051919050565b60006020828403121561181557600080fd5b81516001600160c01b038116811461151e57600080fd5b60006020828403121561183e57600080fd5b815160ff8116811461151e57600080fd5b634e487b7160e01b600052603260045260246000fd5b60ff91909116815260200190565b6000821982111561188657611886611788565b500190565b60006040828403121561189d57600080fd5b604080519081016001600160401b03811182821017156118bf576118bf61144b565b60405282516118cd81611525565b815260208301516001600160601b03811681146118e957600080fd5b60208201529392505050565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b60018060a01b038316815260406020820152600082516060604084015261199760a0840182611708565b90506020840151606084015260408401516080840152809150509392505050565b60208152600061151e6020830184611708565b600181811c908216806119df57607f821691505b60208210811415611a0057634e487b7160e01b600052602260045260246000fd5b50919050565b600082821015611a1857611a18611788565b500390565b600061ffff80831681811415611a3557611a35611788565b600101939250505056fea2646970667358221220c33da8dbfb3d898978d4602d6be26e70a3063075a8f5b8be8d8dfdd7da97c77264736f6c634300080c0033",
}

// ContractOpenOracleServiceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractOpenOracleServiceManagerMetaData.ABI instead.
var ContractOpenOracleServiceManagerABI = ContractOpenOracleServiceManagerMetaData.ABI

// ContractOpenOracleServiceManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractOpenOracleServiceManagerMetaData.Bin instead.
var ContractOpenOracleServiceManagerBin = ContractOpenOracleServiceManagerMetaData.Bin

// DeployContractOpenOracleServiceManager deploys a new Ethereum contract, binding an instance of ContractOpenOracleServiceManager to it.
func DeployContractOpenOracleServiceManager(auth *bind.TransactOpts, backend bind.ContractBackend, _avsDirectory common.Address, _registryCoordinator common.Address, _stakeRegistry common.Address, _openOracleTaskManager common.Address) (common.Address, *types.Transaction, *ContractOpenOracleServiceManager, error) {
	parsed, err := ContractOpenOracleServiceManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractOpenOracleServiceManagerBin), backend, _avsDirectory, _registryCoordinator, _stakeRegistry, _openOracleTaskManager)
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

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOpenOracleServiceManager.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerSession) AvsDirectory() (common.Address, error) {
	return _ContractOpenOracleServiceManager.Contract.AvsDirectory(&_ContractOpenOracleServiceManager.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerCallerSession) AvsDirectory() (common.Address, error) {
	return _ContractOpenOracleServiceManager.Contract.AvsDirectory(&_ContractOpenOracleServiceManager.CallOpts)
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

// TaskManagerCount is a free data retrieval call binding the contract method 0x7b5b3049.
//
// Solidity: function taskManagerCount() view returns(uint256)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerCaller) TaskManagerCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractOpenOracleServiceManager.contract.Call(opts, &out, "taskManagerCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaskManagerCount is a free data retrieval call binding the contract method 0x7b5b3049.
//
// Solidity: function taskManagerCount() view returns(uint256)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerSession) TaskManagerCount() (*big.Int, error) {
	return _ContractOpenOracleServiceManager.Contract.TaskManagerCount(&_ContractOpenOracleServiceManager.CallOpts)
}

// TaskManagerCount is a free data retrieval call binding the contract method 0x7b5b3049.
//
// Solidity: function taskManagerCount() view returns(uint256)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerCallerSession) TaskManagerCount() (*big.Int, error) {
	return _ContractOpenOracleServiceManager.Contract.TaskManagerCount(&_ContractOpenOracleServiceManager.CallOpts)
}

// TaskManagers is a free data retrieval call binding the contract method 0xfd69a60f.
//
// Solidity: function taskManagers(uint256 ) view returns(string chainName, address taskManagerAddress, bool isActive)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerCaller) TaskManagers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ChainName          string
	TaskManagerAddress common.Address
	IsActive           bool
}, error) {
	var out []interface{}
	err := _ContractOpenOracleServiceManager.contract.Call(opts, &out, "taskManagers", arg0)

	outstruct := new(struct {
		ChainName          string
		TaskManagerAddress common.Address
		IsActive           bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChainName = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.TaskManagerAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.IsActive = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// TaskManagers is a free data retrieval call binding the contract method 0xfd69a60f.
//
// Solidity: function taskManagers(uint256 ) view returns(string chainName, address taskManagerAddress, bool isActive)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerSession) TaskManagers(arg0 *big.Int) (struct {
	ChainName          string
	TaskManagerAddress common.Address
	IsActive           bool
}, error) {
	return _ContractOpenOracleServiceManager.Contract.TaskManagers(&_ContractOpenOracleServiceManager.CallOpts, arg0)
}

// TaskManagers is a free data retrieval call binding the contract method 0xfd69a60f.
//
// Solidity: function taskManagers(uint256 ) view returns(string chainName, address taskManagerAddress, bool isActive)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerCallerSession) TaskManagers(arg0 *big.Int) (struct {
	ChainName          string
	TaskManagerAddress common.Address
	IsActive           bool
}, error) {
	return _ContractOpenOracleServiceManager.Contract.TaskManagers(&_ContractOpenOracleServiceManager.CallOpts, arg0)
}

// AddTaskManager is a paid mutator transaction binding the contract method 0x32f21894.
//
// Solidity: function addTaskManager(string chainName, address taskManagerAddress) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerTransactor) AddTaskManager(opts *bind.TransactOpts, chainName string, taskManagerAddress common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.contract.Transact(opts, "addTaskManager", chainName, taskManagerAddress)
}

// AddTaskManager is a paid mutator transaction binding the contract method 0x32f21894.
//
// Solidity: function addTaskManager(string chainName, address taskManagerAddress) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerSession) AddTaskManager(chainName string, taskManagerAddress common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.Contract.AddTaskManager(&_ContractOpenOracleServiceManager.TransactOpts, chainName, taskManagerAddress)
}

// AddTaskManager is a paid mutator transaction binding the contract method 0x32f21894.
//
// Solidity: function addTaskManager(string chainName, address taskManagerAddress) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerTransactorSession) AddTaskManager(chainName string, taskManagerAddress common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.Contract.AddTaskManager(&_ContractOpenOracleServiceManager.TransactOpts, chainName, taskManagerAddress)
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

// RemoveTaskManager is a paid mutator transaction binding the contract method 0xca94cfe0.
//
// Solidity: function removeTaskManager(uint256 id) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerTransactor) RemoveTaskManager(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.contract.Transact(opts, "removeTaskManager", id)
}

// RemoveTaskManager is a paid mutator transaction binding the contract method 0xca94cfe0.
//
// Solidity: function removeTaskManager(uint256 id) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerSession) RemoveTaskManager(id *big.Int) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.Contract.RemoveTaskManager(&_ContractOpenOracleServiceManager.TransactOpts, id)
}

// RemoveTaskManager is a paid mutator transaction binding the contract method 0xca94cfe0.
//
// Solidity: function removeTaskManager(uint256 id) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerTransactorSession) RemoveTaskManager(id *big.Int) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.Contract.RemoveTaskManager(&_ContractOpenOracleServiceManager.TransactOpts, id)
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

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.contract.Transact(opts, "updateAVSMetadataURI", _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.Contract.UpdateAVSMetadataURI(&_ContractOpenOracleServiceManager.TransactOpts, _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerTransactorSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.Contract.UpdateAVSMetadataURI(&_ContractOpenOracleServiceManager.TransactOpts, _metadataURI)
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

// ContractOpenOracleServiceManagerTaskManagerAddedIterator is returned from FilterTaskManagerAdded and is used to iterate over the raw logs and unpacked data for TaskManagerAdded events raised by the ContractOpenOracleServiceManager contract.
type ContractOpenOracleServiceManagerTaskManagerAddedIterator struct {
	Event *ContractOpenOracleServiceManagerTaskManagerAdded // Event containing the contract specifics and raw log

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
func (it *ContractOpenOracleServiceManagerTaskManagerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenOracleServiceManagerTaskManagerAdded)
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
		it.Event = new(ContractOpenOracleServiceManagerTaskManagerAdded)
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
func (it *ContractOpenOracleServiceManagerTaskManagerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenOracleServiceManagerTaskManagerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenOracleServiceManagerTaskManagerAdded represents a TaskManagerAdded event raised by the ContractOpenOracleServiceManager contract.
type ContractOpenOracleServiceManagerTaskManagerAdded struct {
	Id                 *big.Int
	ChainName          string
	TaskManagerAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTaskManagerAdded is a free log retrieval operation binding the contract event 0x661e3d4c4a13612fad1504b96725fcedfb885ff3ac490c18004f57ad9f4e8778.
//
// Solidity: event TaskManagerAdded(uint256 id, string chainName, address taskManagerAddress)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerFilterer) FilterTaskManagerAdded(opts *bind.FilterOpts) (*ContractOpenOracleServiceManagerTaskManagerAddedIterator, error) {

	logs, sub, err := _ContractOpenOracleServiceManager.contract.FilterLogs(opts, "TaskManagerAdded")
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleServiceManagerTaskManagerAddedIterator{contract: _ContractOpenOracleServiceManager.contract, event: "TaskManagerAdded", logs: logs, sub: sub}, nil
}

// WatchTaskManagerAdded is a free log subscription operation binding the contract event 0x661e3d4c4a13612fad1504b96725fcedfb885ff3ac490c18004f57ad9f4e8778.
//
// Solidity: event TaskManagerAdded(uint256 id, string chainName, address taskManagerAddress)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerFilterer) WatchTaskManagerAdded(opts *bind.WatchOpts, sink chan<- *ContractOpenOracleServiceManagerTaskManagerAdded) (event.Subscription, error) {

	logs, sub, err := _ContractOpenOracleServiceManager.contract.WatchLogs(opts, "TaskManagerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenOracleServiceManagerTaskManagerAdded)
				if err := _ContractOpenOracleServiceManager.contract.UnpackLog(event, "TaskManagerAdded", log); err != nil {
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

// ParseTaskManagerAdded is a log parse operation binding the contract event 0x661e3d4c4a13612fad1504b96725fcedfb885ff3ac490c18004f57ad9f4e8778.
//
// Solidity: event TaskManagerAdded(uint256 id, string chainName, address taskManagerAddress)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerFilterer) ParseTaskManagerAdded(log types.Log) (*ContractOpenOracleServiceManagerTaskManagerAdded, error) {
	event := new(ContractOpenOracleServiceManagerTaskManagerAdded)
	if err := _ContractOpenOracleServiceManager.contract.UnpackLog(event, "TaskManagerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOpenOracleServiceManagerTaskManagerRemovedIterator is returned from FilterTaskManagerRemoved and is used to iterate over the raw logs and unpacked data for TaskManagerRemoved events raised by the ContractOpenOracleServiceManager contract.
type ContractOpenOracleServiceManagerTaskManagerRemovedIterator struct {
	Event *ContractOpenOracleServiceManagerTaskManagerRemoved // Event containing the contract specifics and raw log

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
func (it *ContractOpenOracleServiceManagerTaskManagerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenOracleServiceManagerTaskManagerRemoved)
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
		it.Event = new(ContractOpenOracleServiceManagerTaskManagerRemoved)
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
func (it *ContractOpenOracleServiceManagerTaskManagerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenOracleServiceManagerTaskManagerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenOracleServiceManagerTaskManagerRemoved represents a TaskManagerRemoved event raised by the ContractOpenOracleServiceManager contract.
type ContractOpenOracleServiceManagerTaskManagerRemoved struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTaskManagerRemoved is a free log retrieval operation binding the contract event 0x191603ba9e3b1261ce14bffd91013f2cc8137f0292821bae13360bfc0e21d44c.
//
// Solidity: event TaskManagerRemoved(uint256 id)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerFilterer) FilterTaskManagerRemoved(opts *bind.FilterOpts) (*ContractOpenOracleServiceManagerTaskManagerRemovedIterator, error) {

	logs, sub, err := _ContractOpenOracleServiceManager.contract.FilterLogs(opts, "TaskManagerRemoved")
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleServiceManagerTaskManagerRemovedIterator{contract: _ContractOpenOracleServiceManager.contract, event: "TaskManagerRemoved", logs: logs, sub: sub}, nil
}

// WatchTaskManagerRemoved is a free log subscription operation binding the contract event 0x191603ba9e3b1261ce14bffd91013f2cc8137f0292821bae13360bfc0e21d44c.
//
// Solidity: event TaskManagerRemoved(uint256 id)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerFilterer) WatchTaskManagerRemoved(opts *bind.WatchOpts, sink chan<- *ContractOpenOracleServiceManagerTaskManagerRemoved) (event.Subscription, error) {

	logs, sub, err := _ContractOpenOracleServiceManager.contract.WatchLogs(opts, "TaskManagerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenOracleServiceManagerTaskManagerRemoved)
				if err := _ContractOpenOracleServiceManager.contract.UnpackLog(event, "TaskManagerRemoved", log); err != nil {
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

// ParseTaskManagerRemoved is a log parse operation binding the contract event 0x191603ba9e3b1261ce14bffd91013f2cc8137f0292821bae13360bfc0e21d44c.
//
// Solidity: event TaskManagerRemoved(uint256 id)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerFilterer) ParseTaskManagerRemoved(log types.Log) (*ContractOpenOracleServiceManagerTaskManagerRemoved, error) {
	event := new(ContractOpenOracleServiceManagerTaskManagerRemoved)
	if err := _ContractOpenOracleServiceManager.contract.UnpackLog(event, "TaskManagerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

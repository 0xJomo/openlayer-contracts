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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"},{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addOperatorToRegistryWhitelist\",\"inputs\":[{\"name\":\"operatorsToWhitelist\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addTaskManager\",\"inputs\":[{\"name\":\"chainName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"taskManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorIsWhitelistedForRegister\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeOperatorFromRegistryWhitelist\",\"inputs\":[{\"name\":\"operatorsToRemoveFromWhitelist\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeTaskManager\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"taskManagerCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskManagers\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"chainName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"taskManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToRegistryWhitelist\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromRegistryWhitelist\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskManagerAdded\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"chainName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"taskManagerAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskManagerRemoved\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162001f6038038062001f60833981016040819052620000349162000141565b6001600160a01b0380841660c052808316608052811660a0528282826200005a62000066565b50505050505062000195565b600054610100900460ff1615620000d35760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116101562000126576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200013e57600080fd5b50565b6000806000606084860312156200015757600080fd5b8351620001648162000128565b6020850151909350620001778162000128565b60408501519092506200018a8162000128565b809150509250925092565b60805160a05160c051611d39620002276000396000818161014601528181610b7001528181610c3d0152610cc601526000818161066a015281816107c00152818161086201528181610fba0152818161113f01526111e101526000818161047501528181610524015281816105a401528181610a8e01528181610be901528181610ef9015261109b0152611d396000f3fe608060405234801561001057600080fd5b50600436106100db5760003560e01c8063183b657d146100e057806332f21894146100f557806333cfb7b7146101085780634ce5075b146101315780636b3aa72e14610144578063715018a6146101735780637b5b30491461017b5780638da5cb5b146101925780639926ee7d1461019a578063a364f4da146101ad578063a98fb355146101c0578063c4d66de8146101d3578063c8689338146101e6578063ca94cfe014610219578063e481af9d1461022c578063f2fde38b14610234578063fd69a60f14610247575b600080fd5b6100f36100ee36600461169b565b610269565b005b6100f36101033660046117fe565b61038b565b61011b61011636600461184f565b61046f565b604051610128919061186c565b60405180910390f35b6100f361013f36600461169b565b610945565b7f00000000000000000000000000000000000000000000000000000000000000005b60405161012891906118b9565b6100f3610a60565b61018460995481565b604051908152602001610128565b610166610a74565b6100f36101a83660046118cd565b610a83565b6100f36101bb36600461184f565b610bde565b6100f36101ce366004611977565b610ca7565b6100f36101e136600461184f565b610cfb565b6102096101f436600461184f565b60976020526000908152604090205460ff1681565b6040519015158152602001610128565b6100f36102273660046119b3565b610e13565b61011b610ef3565b6100f361024236600461184f565b6112c0565b61025a6102553660046119b3565b611339565b60405161012893929190611a19565b6102716113f3565b8060005b81811015610385576097600085858481811061029357610293611a4c565b90506020020160208101906102a8919061184f565b6001600160a01b0316815260208101919091526040016000205460ff161561037d576000609760008686858181106102e2576102e2611a4c565b90506020020160208101906102f7919061184f565b6001600160a01b031681526020810191909152604001600020805460ff19169115159190911790557f3b14e126d552f04f750ed8b17f3402b755ef087779a91250f50a0cb8a95fd87384848381811061035257610352611a4c565b9050602002016020810190610367919061184f565b60405161037491906118b9565b60405180910390a15b600101610275565b50505050565b6103936113f3565b60998054600091826103a483611a78565b90915550604080516060810182528581526001600160a01b0385166020808301919091526001828401526000848152609882529290922081518051949550919390926103f4928492910190611602565b506020820151600190910180546040938401511515600160a01b026001600160a81b03199091166001600160a01b0390931692909217919091179055517f661e3d4c4a13612fad1504b96725fcedfb885ff3ac490c18004f57ad9f4e87789061046290839086908690611a93565b60405180910390a1505050565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166313542a4e846040518263ffffffff1660e01b81526004016104bf91906118b9565b602060405180830381865afa1580156104dc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105009190611ac4565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa15801561056b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061058f9190611add565b90506001600160c01b038116158061062957507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610600573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106249190611b06565b60ff16155b1561064557505060408051600081526020810190915292915050565b6000610659826001600160c01b0316611452565b90506000805b8251811015610731577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f58483815181106106a9576106a9611a4c565b01602001516040516001600160e01b031960e084901b1681526106d29160f81c90600401611b29565b602060405180830381865afa1580156106ef573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107139190611ac4565b61071d9083611b37565b91508061072981611a78565b91505061065f565b506000816001600160401b0381111561074c5761074c61170f565b604051908082528060200260200182016040528015610775578160200160208202803683370190505b5090506000805b845181101561093857600085828151811061079957610799611a4c565b0160200151604051633ca5a5f560e01b815260f89190911c91506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f5906107f5908590600401611b29565b602060405180830381865afa158015610812573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108369190611ac4565b905060005b81811015610922576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa1580156108b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108d49190611b4f565b600001518686815181106108ea576108ea611a4c565b6001600160a01b03909216602092830291909101909101528461090c81611a78565b955050808061091a90611a78565b91505061083b565b505050808061093090611a78565b91505061077c565b5090979650505050505050565b61094d6113f3565b8060005b81811015610385576097600085858481811061096f5761096f611a4c565b9050602002016020810190610984919061184f565b6001600160a01b0316815260208101919091526040016000205460ff16610a58576001609760008686858181106109bd576109bd611a4c565b90506020020160208101906109d2919061184f565b6001600160a01b031681526020810191909152604001600020805460ff19169115159190911790557f9408e0bce32ef2257130df793b1f3f21e664eea865cefce7475ab6e0982f98a2848483818110610a2d57610a2d611a4c565b9050602002016020810190610a42919061184f565b604051610a4f91906118b9565b60405180910390a15b600101610951565b610a686113f3565b610a726000611514565b565b6033546001600160a01b031690565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610ad45760405162461bcd60e51b8152600401610acb90611bb9565b60405180910390fd5b6001600160a01b038216600090815260976020526040902054829060ff16610b595760405162461bcd60e51b815260206004820152603260248201527f6f6e6c7957686974654c69737465644f70657261746f723a206f70657261746f6044820152711c881b9bdd081a5b881dda1a5d195b1a5cdd60721b6064820152608401610acb565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d90610ba79086908690600401611c31565b600060405180830381600087803b158015610bc157600080fd5b505af1158015610bd5573d6000803e3d6000fd5b50505050505050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610c265760405162461bcd60e51b8152600401610acb90611bb9565b6040516351b27a6d60e11b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da90610c729084906004016118b9565b600060405180830381600087803b158015610c8c57600080fd5b505af1158015610ca0573d6000803e3d6000fd5b5050505050565b610caf6113f3565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb35590610c72908490600401611c7c565b600054610100900460ff1615808015610d1b5750600054600160ff909116105b80610d355750303b158015610d35575060005460ff166001145b610d985760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610acb565b6000805460ff191660011790558015610dbb576000805461ff0019166101001790555b610dc482611566565b8015610e0f576000805461ff00191690556040517f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890610e0690600190611b29565b60405180910390a15b5050565b610e1b6113f3565b600081815260986020526040902060010154600160a01b900460ff16610e9a5760405162461bcd60e51b815260206004820152602e60248201527f5461736b204d616e6167657220646f6573206e6f74206578697374206f72206160448201526d1b1c9958591e481c995b5bdd995960921b6064820152608401610acb565b60008181526098602052604090819020600101805460ff60a01b19169055517f191603ba9e3b1261ce14bffd91013f2cc8137f0292821bae13360bfc0e21d44c90610ee89083815260200190565b60405180910390a150565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f55573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f799190611b06565b60ff16905080610f9757505060408051600081526020810190915290565b6000805b8281101561104e57604051633ca5a5f560e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590610fef908490600401611b29565b602060405180830381865afa15801561100c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110309190611ac4565b61103a9083611b37565b91508061104681611a78565b915050610f9b565b506000816001600160401b038111156110695761106961170f565b604051908082528060200260200182016040528015611092578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156110f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061111b9190611b06565b60ff168110156112b657604051633ca5a5f560e01b81526000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590611174908590600401611b29565b602060405180830381865afa158015611191573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111b59190611ac4565b905060005b818110156112a1576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa15801561122f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112539190611b4f565b6000015185858151811061126957611269611a4c565b6001600160a01b03909216602092830291909101909101528361128b81611a78565b945050808061129990611a78565b9150506111ba565b505080806112ae90611a78565b915050611099565b5090949350505050565b6112c86113f3565b6001600160a01b03811661132d5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610acb565b61133681611514565b50565b60986020526000908152604090208054819061135490611c8f565b80601f016020809104026020016040519081016040528092919081815260200182805461138090611c8f565b80156113cd5780601f106113a2576101008083540402835291602001916113cd565b820191906000526020600020905b8154815290600101906020018083116113b057829003601f168201915b505050600190930154919250506001600160a01b0381169060ff600160a01b9091041683565b336113fc610a74565b6001600160a01b031614610a725760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610acb565b6060600080611460846115d1565b61ffff166001600160401b0381111561147b5761147b61170f565b6040519080825280601f01601f1916602001820160405280156114a5576020820181803683370190505b5090506000805b8251821080156114bd575061010081105b156112b6576001811b935085841615611504578060f81b8383815181106114e6576114e6611a4c565b60200101906001600160f81b031916908160001a9053508160010191505b61150d81611a78565b90506114ac565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff1661132d5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608401610acb565b6000805b82156115fc576115e6600184611cca565b90921691806115f481611ce1565b9150506115d5565b92915050565b82805461160e90611c8f565b90600052602060002090601f0160209004810192826116305760008555611676565b82601f1061164957805160ff1916838001178555611676565b82800160010185558215611676579182015b8281111561167657825182559160200191906001019061165b565b50611682929150611686565b5090565b5b808211156116825760008155600101611687565b600080602083850312156116ae57600080fd5b82356001600160401b03808211156116c557600080fd5b818501915085601f8301126116d957600080fd5b8135818111156116e857600080fd5b8660208260051b85010111156116fd57600080fd5b60209290920196919550909350505050565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b03811182821017156117475761174761170f565b60405290565b60006001600160401b03808411156117675761176761170f565b604051601f8501601f19908116603f0116810190828211818310171561178f5761178f61170f565b816040528093508581528686860111156117a857600080fd5b858560208301376000602087830101525050509392505050565b600082601f8301126117d357600080fd5b6117e28383356020850161174d565b9392505050565b6001600160a01b038116811461133657600080fd5b6000806040838503121561181157600080fd5b82356001600160401b0381111561182757600080fd5b611833858286016117c2565b9250506020830135611844816117e9565b809150509250929050565b60006020828403121561186157600080fd5b81356117e2816117e9565b6020808252825182820181905260009190848201906040850190845b818110156118ad5783516001600160a01b031683529284019291840191600101611888565b50909695505050505050565b6001600160a01b0391909116815260200190565b600080604083850312156118e057600080fd5b82356118eb816117e9565b915060208301356001600160401b038082111561190757600080fd5b908401906060828703121561191b57600080fd5b611923611725565b82358281111561193257600080fd5b83019150601f8201871361194557600080fd5b6119548783356020850161174d565b815260208301356020820152604083013560408201528093505050509250929050565b60006020828403121561198957600080fd5b81356001600160401b0381111561199f57600080fd5b6119ab848285016117c2565b949350505050565b6000602082840312156119c557600080fd5b5035919050565b6000815180845260005b818110156119f2576020818501810151868301820152016119d6565b81811115611a04576000602083870101525b50601f01601f19169290920160200192915050565b606081526000611a2c60608301866119cc565b6001600160a01b0394909416602083015250901515604090910152919050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415611a8c57611a8c611a62565b5060010190565b838152606060208201526000611aac60608301856119cc565b905060018060a01b0383166040830152949350505050565b600060208284031215611ad657600080fd5b5051919050565b600060208284031215611aef57600080fd5b81516001600160c01b03811681146117e257600080fd5b600060208284031215611b1857600080fd5b815160ff811681146117e257600080fd5b60ff91909116815260200190565b60008219821115611b4a57611b4a611a62565b500190565b600060408284031215611b6157600080fd5b604080519081016001600160401b0381118282101715611b8357611b8361170f565b6040528251611b91816117e9565b815260208301516001600160601b0381168114611bad57600080fd5b60208201529392505050565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b60018060a01b0383168152604060208201526000825160606040840152611c5b60a08401826119cc565b90506020840151606084015260408401516080840152809150509392505050565b6020815260006117e260208301846119cc565b600181811c90821680611ca357607f821691505b60208210811415611cc457634e487b7160e01b600052602260045260246000fd5b50919050565b600082821015611cdc57611cdc611a62565b500390565b600061ffff80831681811415611cf957611cf9611a62565b600101939250505056fea2646970667358221220b7a8a887e3b823dc94341b4e8e83c46dda5ef3a9f2489e792b0700f1a1c8d67d64736f6c634300080c0033",
}

// ContractOpenOracleServiceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractOpenOracleServiceManagerMetaData.ABI instead.
var ContractOpenOracleServiceManagerABI = ContractOpenOracleServiceManagerMetaData.ABI

// ContractOpenOracleServiceManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractOpenOracleServiceManagerMetaData.Bin instead.
var ContractOpenOracleServiceManagerBin = ContractOpenOracleServiceManagerMetaData.Bin

// DeployContractOpenOracleServiceManager deploys a new Ethereum contract, binding an instance of ContractOpenOracleServiceManager to it.
func DeployContractOpenOracleServiceManager(auth *bind.TransactOpts, backend bind.ContractBackend, _avsDirectory common.Address, _registryCoordinator common.Address, _stakeRegistry common.Address) (common.Address, *types.Transaction, *ContractOpenOracleServiceManager, error) {
	parsed, err := ContractOpenOracleServiceManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractOpenOracleServiceManagerBin), backend, _avsDirectory, _registryCoordinator, _stakeRegistry)
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

// OperatorIsWhitelistedForRegister is a free data retrieval call binding the contract method 0xc8689338.
//
// Solidity: function operatorIsWhitelistedForRegister(address ) view returns(bool)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerCaller) OperatorIsWhitelistedForRegister(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ContractOpenOracleServiceManager.contract.Call(opts, &out, "operatorIsWhitelistedForRegister", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorIsWhitelistedForRegister is a free data retrieval call binding the contract method 0xc8689338.
//
// Solidity: function operatorIsWhitelistedForRegister(address ) view returns(bool)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerSession) OperatorIsWhitelistedForRegister(arg0 common.Address) (bool, error) {
	return _ContractOpenOracleServiceManager.Contract.OperatorIsWhitelistedForRegister(&_ContractOpenOracleServiceManager.CallOpts, arg0)
}

// OperatorIsWhitelistedForRegister is a free data retrieval call binding the contract method 0xc8689338.
//
// Solidity: function operatorIsWhitelistedForRegister(address ) view returns(bool)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerCallerSession) OperatorIsWhitelistedForRegister(arg0 common.Address) (bool, error) {
	return _ContractOpenOracleServiceManager.Contract.OperatorIsWhitelistedForRegister(&_ContractOpenOracleServiceManager.CallOpts, arg0)
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

// AddOperatorToRegistryWhitelist is a paid mutator transaction binding the contract method 0x4ce5075b.
//
// Solidity: function addOperatorToRegistryWhitelist(address[] operatorsToWhitelist) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerTransactor) AddOperatorToRegistryWhitelist(opts *bind.TransactOpts, operatorsToWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.contract.Transact(opts, "addOperatorToRegistryWhitelist", operatorsToWhitelist)
}

// AddOperatorToRegistryWhitelist is a paid mutator transaction binding the contract method 0x4ce5075b.
//
// Solidity: function addOperatorToRegistryWhitelist(address[] operatorsToWhitelist) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerSession) AddOperatorToRegistryWhitelist(operatorsToWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.Contract.AddOperatorToRegistryWhitelist(&_ContractOpenOracleServiceManager.TransactOpts, operatorsToWhitelist)
}

// AddOperatorToRegistryWhitelist is a paid mutator transaction binding the contract method 0x4ce5075b.
//
// Solidity: function addOperatorToRegistryWhitelist(address[] operatorsToWhitelist) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerTransactorSession) AddOperatorToRegistryWhitelist(operatorsToWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.Contract.AddOperatorToRegistryWhitelist(&_ContractOpenOracleServiceManager.TransactOpts, operatorsToWhitelist)
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

// RemoveOperatorFromRegistryWhitelist is a paid mutator transaction binding the contract method 0x183b657d.
//
// Solidity: function removeOperatorFromRegistryWhitelist(address[] operatorsToRemoveFromWhitelist) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerTransactor) RemoveOperatorFromRegistryWhitelist(opts *bind.TransactOpts, operatorsToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.contract.Transact(opts, "removeOperatorFromRegistryWhitelist", operatorsToRemoveFromWhitelist)
}

// RemoveOperatorFromRegistryWhitelist is a paid mutator transaction binding the contract method 0x183b657d.
//
// Solidity: function removeOperatorFromRegistryWhitelist(address[] operatorsToRemoveFromWhitelist) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerSession) RemoveOperatorFromRegistryWhitelist(operatorsToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.Contract.RemoveOperatorFromRegistryWhitelist(&_ContractOpenOracleServiceManager.TransactOpts, operatorsToRemoveFromWhitelist)
}

// RemoveOperatorFromRegistryWhitelist is a paid mutator transaction binding the contract method 0x183b657d.
//
// Solidity: function removeOperatorFromRegistryWhitelist(address[] operatorsToRemoveFromWhitelist) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerTransactorSession) RemoveOperatorFromRegistryWhitelist(operatorsToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.Contract.RemoveOperatorFromRegistryWhitelist(&_ContractOpenOracleServiceManager.TransactOpts, operatorsToRemoveFromWhitelist)
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

// ContractOpenOracleServiceManagerOperatorAddedToRegistryWhitelistIterator is returned from FilterOperatorAddedToRegistryWhitelist and is used to iterate over the raw logs and unpacked data for OperatorAddedToRegistryWhitelist events raised by the ContractOpenOracleServiceManager contract.
type ContractOpenOracleServiceManagerOperatorAddedToRegistryWhitelistIterator struct {
	Event *ContractOpenOracleServiceManagerOperatorAddedToRegistryWhitelist // Event containing the contract specifics and raw log

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
func (it *ContractOpenOracleServiceManagerOperatorAddedToRegistryWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenOracleServiceManagerOperatorAddedToRegistryWhitelist)
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
		it.Event = new(ContractOpenOracleServiceManagerOperatorAddedToRegistryWhitelist)
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
func (it *ContractOpenOracleServiceManagerOperatorAddedToRegistryWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenOracleServiceManagerOperatorAddedToRegistryWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenOracleServiceManagerOperatorAddedToRegistryWhitelist represents a OperatorAddedToRegistryWhitelist event raised by the ContractOpenOracleServiceManager contract.
type ContractOpenOracleServiceManagerOperatorAddedToRegistryWhitelist struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorAddedToRegistryWhitelist is a free log retrieval operation binding the contract event 0x9408e0bce32ef2257130df793b1f3f21e664eea865cefce7475ab6e0982f98a2.
//
// Solidity: event OperatorAddedToRegistryWhitelist(address operator)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerFilterer) FilterOperatorAddedToRegistryWhitelist(opts *bind.FilterOpts) (*ContractOpenOracleServiceManagerOperatorAddedToRegistryWhitelistIterator, error) {

	logs, sub, err := _ContractOpenOracleServiceManager.contract.FilterLogs(opts, "OperatorAddedToRegistryWhitelist")
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleServiceManagerOperatorAddedToRegistryWhitelistIterator{contract: _ContractOpenOracleServiceManager.contract, event: "OperatorAddedToRegistryWhitelist", logs: logs, sub: sub}, nil
}

// WatchOperatorAddedToRegistryWhitelist is a free log subscription operation binding the contract event 0x9408e0bce32ef2257130df793b1f3f21e664eea865cefce7475ab6e0982f98a2.
//
// Solidity: event OperatorAddedToRegistryWhitelist(address operator)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerFilterer) WatchOperatorAddedToRegistryWhitelist(opts *bind.WatchOpts, sink chan<- *ContractOpenOracleServiceManagerOperatorAddedToRegistryWhitelist) (event.Subscription, error) {

	logs, sub, err := _ContractOpenOracleServiceManager.contract.WatchLogs(opts, "OperatorAddedToRegistryWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenOracleServiceManagerOperatorAddedToRegistryWhitelist)
				if err := _ContractOpenOracleServiceManager.contract.UnpackLog(event, "OperatorAddedToRegistryWhitelist", log); err != nil {
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

// ParseOperatorAddedToRegistryWhitelist is a log parse operation binding the contract event 0x9408e0bce32ef2257130df793b1f3f21e664eea865cefce7475ab6e0982f98a2.
//
// Solidity: event OperatorAddedToRegistryWhitelist(address operator)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerFilterer) ParseOperatorAddedToRegistryWhitelist(log types.Log) (*ContractOpenOracleServiceManagerOperatorAddedToRegistryWhitelist, error) {
	event := new(ContractOpenOracleServiceManagerOperatorAddedToRegistryWhitelist)
	if err := _ContractOpenOracleServiceManager.contract.UnpackLog(event, "OperatorAddedToRegistryWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOpenOracleServiceManagerOperatorRemovedFromRegistryWhitelistIterator is returned from FilterOperatorRemovedFromRegistryWhitelist and is used to iterate over the raw logs and unpacked data for OperatorRemovedFromRegistryWhitelist events raised by the ContractOpenOracleServiceManager contract.
type ContractOpenOracleServiceManagerOperatorRemovedFromRegistryWhitelistIterator struct {
	Event *ContractOpenOracleServiceManagerOperatorRemovedFromRegistryWhitelist // Event containing the contract specifics and raw log

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
func (it *ContractOpenOracleServiceManagerOperatorRemovedFromRegistryWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenOracleServiceManagerOperatorRemovedFromRegistryWhitelist)
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
		it.Event = new(ContractOpenOracleServiceManagerOperatorRemovedFromRegistryWhitelist)
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
func (it *ContractOpenOracleServiceManagerOperatorRemovedFromRegistryWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenOracleServiceManagerOperatorRemovedFromRegistryWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenOracleServiceManagerOperatorRemovedFromRegistryWhitelist represents a OperatorRemovedFromRegistryWhitelist event raised by the ContractOpenOracleServiceManager contract.
type ContractOpenOracleServiceManagerOperatorRemovedFromRegistryWhitelist struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemovedFromRegistryWhitelist is a free log retrieval operation binding the contract event 0x3b14e126d552f04f750ed8b17f3402b755ef087779a91250f50a0cb8a95fd873.
//
// Solidity: event OperatorRemovedFromRegistryWhitelist(address operator)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerFilterer) FilterOperatorRemovedFromRegistryWhitelist(opts *bind.FilterOpts) (*ContractOpenOracleServiceManagerOperatorRemovedFromRegistryWhitelistIterator, error) {

	logs, sub, err := _ContractOpenOracleServiceManager.contract.FilterLogs(opts, "OperatorRemovedFromRegistryWhitelist")
	if err != nil {
		return nil, err
	}
	return &ContractOpenOracleServiceManagerOperatorRemovedFromRegistryWhitelistIterator{contract: _ContractOpenOracleServiceManager.contract, event: "OperatorRemovedFromRegistryWhitelist", logs: logs, sub: sub}, nil
}

// WatchOperatorRemovedFromRegistryWhitelist is a free log subscription operation binding the contract event 0x3b14e126d552f04f750ed8b17f3402b755ef087779a91250f50a0cb8a95fd873.
//
// Solidity: event OperatorRemovedFromRegistryWhitelist(address operator)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerFilterer) WatchOperatorRemovedFromRegistryWhitelist(opts *bind.WatchOpts, sink chan<- *ContractOpenOracleServiceManagerOperatorRemovedFromRegistryWhitelist) (event.Subscription, error) {

	logs, sub, err := _ContractOpenOracleServiceManager.contract.WatchLogs(opts, "OperatorRemovedFromRegistryWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenOracleServiceManagerOperatorRemovedFromRegistryWhitelist)
				if err := _ContractOpenOracleServiceManager.contract.UnpackLog(event, "OperatorRemovedFromRegistryWhitelist", log); err != nil {
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

// ParseOperatorRemovedFromRegistryWhitelist is a log parse operation binding the contract event 0x3b14e126d552f04f750ed8b17f3402b755ef087779a91250f50a0cb8a95fd873.
//
// Solidity: event OperatorRemovedFromRegistryWhitelist(address operator)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerFilterer) ParseOperatorRemovedFromRegistryWhitelist(log types.Log) (*ContractOpenOracleServiceManagerOperatorRemovedFromRegistryWhitelist, error) {
	event := new(ContractOpenOracleServiceManagerOperatorRemovedFromRegistryWhitelist)
	if err := _ContractOpenOracleServiceManager.contract.UnpackLog(event, "OperatorRemovedFromRegistryWhitelist", log); err != nil {
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

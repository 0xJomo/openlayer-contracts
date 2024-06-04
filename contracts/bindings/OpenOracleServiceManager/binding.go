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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"},{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addOperatorToRegistryWhitelist\",\"inputs\":[{\"name\":\"operatorsToWhitelist\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addTaskManager\",\"inputs\":[{\"name\":\"chainName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"taskManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addTaskManager\",\"inputs\":[{\"name\":\"chainName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"taskManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"managerType\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorIsWhitelistedForRegister\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeOperatorFromRegistryWhitelist\",\"inputs\":[{\"name\":\"operatorsToRemoveFromWhitelist\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeTaskManager\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"taskManagerCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskManagers\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"chainName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"taskManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"managerType\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToRegistryWhitelist\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromRegistryWhitelist\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskManagerAdded\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"chainName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"taskManagerAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskManagerRemoved\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60e06040523480156200001157600080fd5b506040516200212938038062002129833981016040819052620000349162000141565b6001600160a01b0380841660c052808316608052811660a0528282826200005a62000066565b50505050505062000195565b600054610100900460ff1615620000d35760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116101562000126576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200013e57600080fd5b50565b6000806000606084860312156200015757600080fd5b8351620001648162000128565b6020850151909350620001778162000128565b60408501519092506200018a8162000128565b809150509250925092565b60805160a05160c051611f02620002276000396000818161015101528181610c9c01528181610d690152610df201526000818161069c015281816107f201528181610894015281816110e60152818161126b015261130d0152600081816104a701528181610556015281816105d601528181610bba01528181610d150152818161102501526111c70152611f026000f3fe608060405234801561001057600080fd5b50600436106100e65760003560e01c8063183b657d146100eb57806332f218941461010057806333cfb7b7146101135780634ce5075b1461013c5780636b3aa72e1461014f578063715018a61461017e5780637b5b304914610186578063868529c31461019d5780638da5cb5b146101b05780639926ee7d146101b8578063a364f4da146101cb578063a98fb355146101de578063c4d66de8146101f1578063c868933814610204578063ca94cfe014610237578063e481af9d1461024a578063f2fde38b14610252578063fd69a60f14610265575b600080fd5b6100fe6100f93660046117d0565b610288565b005b6100fe61010e366004611933565b6103aa565b610126610121366004611984565b6104a1565b60405161013391906119a1565b60405180910390f35b6100fe61014a3660046117d0565b610977565b7f00000000000000000000000000000000000000000000000000000000000000005b60405161013391906119ee565b6100fe610a92565b61018f60995481565b604051908152602001610133565b6100fe6101ab366004611a11565b610aa6565b610171610ba0565b6100fe6101c6366004611a74565b610baf565b6100fe6101d9366004611984565b610d0a565b6100fe6101ec366004611b1e565b610dd3565b6100fe6101ff366004611984565b610e27565b610227610212366004611984565b60976020526000908152604090205460ff1681565b6040519015158152602001610133565b6100fe610245366004611b5a565b610f3f565b61012661101f565b6100fe610260366004611984565b6113ec565b610278610273366004611b5a565b611465565b6040516101339493929190611bc0565b610290611528565b8060005b818110156103a457609760008585848181106102b2576102b2611bfb565b90506020020160208101906102c79190611984565b6001600160a01b0316815260208101919091526040016000205460ff161561039c5760006097600086868581811061030157610301611bfb565b90506020020160208101906103169190611984565b6001600160a01b031681526020810191909152604001600020805460ff19169115159190911790557f3b14e126d552f04f750ed8b17f3402b755ef087779a91250f50a0cb8a95fd87384848381811061037157610371611bfb565b90506020020160208101906103869190611984565b60405161039391906119ee565b60405180910390a15b600101610294565b50505050565b6103b2611528565b60998054600091826103c383611c27565b90915550604080516080810182528581526001600160a01b0385166020808301919091526001828401526000606083018190528481526098825292909220815180519495509193909261041a928492910190611737565b5060208201516001909101805460408085015160609095015160ff16600160a81b0260ff60a81b19951515600160a01b026001600160a81b03199093166001600160a01b0390951694909417919091179390931691909117905551600080516020611ead8339815191529061049490839086908690611c42565b60405180910390a1505050565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166313542a4e846040518263ffffffff1660e01b81526004016104f191906119ee565b602060405180830381865afa15801561050e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105329190611c73565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa15801561059d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105c19190611c8c565b90506001600160c01b038116158061065b57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610632573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106569190611cb5565b60ff16155b1561067757505060408051600081526020810190915292915050565b600061068b826001600160c01b0316611587565b90506000805b8251811015610763577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f58483815181106106db576106db611bfb565b01602001516040516001600160e01b031960e084901b1681526107049160f81c90600401611cd2565b602060405180830381865afa158015610721573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107459190611c73565b61074f9083611ce0565b91508061075b81611c27565b915050610691565b506000816001600160401b0381111561077e5761077e611844565b6040519080825280602002602001820160405280156107a7578160200160208202803683370190505b5090506000805b845181101561096a5760008582815181106107cb576107cb611bfb565b0160200151604051633ca5a5f560e01b815260f89190911c91506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590610827908590600401611cd2565b602060405180830381865afa158015610844573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108689190611c73565b905060005b81811015610954576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa1580156108e2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109069190611cf8565b6000015186868151811061091c5761091c611bfb565b6001600160a01b03909216602092830291909101909101528461093e81611c27565b955050808061094c90611c27565b91505061086d565b505050808061096290611c27565b9150506107ae565b5090979650505050505050565b61097f611528565b8060005b818110156103a457609760008585848181106109a1576109a1611bfb565b90506020020160208101906109b69190611984565b6001600160a01b0316815260208101919091526040016000205460ff16610a8a576001609760008686858181106109ef576109ef611bfb565b9050602002016020810190610a049190611984565b6001600160a01b031681526020810191909152604001600020805460ff19169115159190911790557f9408e0bce32ef2257130df793b1f3f21e664eea865cefce7475ab6e0982f98a2848483818110610a5f57610a5f611bfb565b9050602002016020810190610a749190611984565b604051610a8191906119ee565b60405180910390a15b600101610983565b610a9a611528565b610aa46000611649565b565b610aae611528565b6099805460009182610abf83611c27565b90915550604080516080810182528681526001600160a01b03861660208083019190915260018284015260ff86166060830152600084815260988252929092208151805194955091939092610b18928492910190611737565b5060208201516001909101805460408085015160609095015160ff16600160a81b0260ff60a81b19951515600160a01b026001600160a81b03199093166001600160a01b0390951694909417919091179390931691909117905551600080516020611ead83398151915290610b9290839087908790611c42565b60405180910390a150505050565b6033546001600160a01b031690565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610c005760405162461bcd60e51b8152600401610bf790611d62565b60405180910390fd5b6001600160a01b038216600090815260976020526040902054829060ff16610c855760405162461bcd60e51b815260206004820152603260248201527f6f6e6c7957686974654c69737465644f70657261746f723a206f70657261746f6044820152711c881b9bdd081a5b881dda1a5d195b1a5cdd60721b6064820152608401610bf7565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d90610cd39086908690600401611dda565b600060405180830381600087803b158015610ced57600080fd5b505af1158015610d01573d6000803e3d6000fd5b50505050505050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610d525760405162461bcd60e51b8152600401610bf790611d62565b6040516351b27a6d60e11b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da90610d9e9084906004016119ee565b600060405180830381600087803b158015610db857600080fd5b505af1158015610dcc573d6000803e3d6000fd5b5050505050565b610ddb611528565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb35590610d9e908490600401611e25565b600054610100900460ff1615808015610e475750600054600160ff909116105b80610e615750303b158015610e61575060005460ff166001145b610ec45760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610bf7565b6000805460ff191660011790558015610ee7576000805461ff0019166101001790555b610ef08261169b565b8015610f3b576000805461ff00191690556040517f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890610f3290600190611cd2565b60405180910390a15b5050565b610f47611528565b600081815260986020526040902060010154600160a01b900460ff16610fc65760405162461bcd60e51b815260206004820152602e60248201527f5461736b204d616e6167657220646f6573206e6f74206578697374206f72206160448201526d1b1c9958591e481c995b5bdd995960921b6064820152608401610bf7565b60008181526098602052604090819020600101805460ff60a01b19169055517f191603ba9e3b1261ce14bffd91013f2cc8137f0292821bae13360bfc0e21d44c906110149083815260200190565b60405180910390a150565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611081573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110a59190611cb5565b60ff169050806110c357505060408051600081526020810190915290565b6000805b8281101561117a57604051633ca5a5f560e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f59061111b908490600401611cd2565b602060405180830381865afa158015611138573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061115c9190611c73565b6111669083611ce0565b91508061117281611c27565b9150506110c7565b506000816001600160401b0381111561119557611195611844565b6040519080825280602002602001820160405280156111be578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611223573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112479190611cb5565b60ff168110156113e257604051633ca5a5f560e01b81526000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f5906112a0908590600401611cd2565b602060405180830381865afa1580156112bd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112e19190611c73565b905060005b818110156113cd576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa15801561135b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061137f9190611cf8565b6000015185858151811061139557611395611bfb565b6001600160a01b0390921660209283029190910190910152836113b781611c27565b94505080806113c590611c27565b9150506112e6565b505080806113da90611c27565b9150506111c5565b5090949350505050565b6113f4611528565b6001600160a01b0381166114595760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610bf7565b61146281611649565b50565b60986020526000908152604090208054819061148090611e38565b80601f01602080910402602001604051908101604052809291908181526020018280546114ac90611e38565b80156114f95780601f106114ce576101008083540402835291602001916114f9565b820191906000526020600020905b8154815290600101906020018083116114dc57829003601f168201915b505050600190930154919250506001600160a01b0381169060ff600160a01b8204811691600160a81b90041684565b33611531610ba0565b6001600160a01b031614610aa45760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610bf7565b606060008061159584611706565b61ffff166001600160401b038111156115b0576115b0611844565b6040519080825280601f01601f1916602001820160405280156115da576020820181803683370190505b5090506000805b8251821080156115f2575061010081105b156113e2576001811b935085841615611639578060f81b83838151811061161b5761161b611bfb565b60200101906001600160f81b031916908160001a9053508160010191505b61164281611c27565b90506115e1565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff166114595760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608401610bf7565b6000805b82156117315761171b600184611e73565b909216918061172981611e8a565b91505061170a565b92915050565b82805461174390611e38565b90600052602060002090601f01602090048101928261176557600085556117ab565b82601f1061177e57805160ff19168380011785556117ab565b828001600101855582156117ab579182015b828111156117ab578251825591602001919060010190611790565b506117b79291506117bb565b5090565b5b808211156117b757600081556001016117bc565b600080602083850312156117e357600080fd5b82356001600160401b03808211156117fa57600080fd5b818501915085601f83011261180e57600080fd5b81358181111561181d57600080fd5b8660208260051b850101111561183257600080fd5b60209290920196919550909350505050565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b038111828210171561187c5761187c611844565b60405290565b60006001600160401b038084111561189c5761189c611844565b604051601f8501601f19908116603f011681019082821181831017156118c4576118c4611844565b816040528093508581528686860111156118dd57600080fd5b858560208301376000602087830101525050509392505050565b600082601f83011261190857600080fd5b61191783833560208501611882565b9392505050565b6001600160a01b038116811461146257600080fd5b6000806040838503121561194657600080fd5b82356001600160401b0381111561195c57600080fd5b611968858286016118f7565b92505060208301356119798161191e565b809150509250929050565b60006020828403121561199657600080fd5b81356119178161191e565b6020808252825182820181905260009190848201906040850190845b818110156119e25783516001600160a01b0316835292840192918401916001016119bd565b50909695505050505050565b6001600160a01b0391909116815260200190565b60ff8116811461146257600080fd5b600080600060608486031215611a2657600080fd5b83356001600160401b03811115611a3c57600080fd5b611a48868287016118f7565b9350506020840135611a598161191e565b91506040840135611a6981611a02565b809150509250925092565b60008060408385031215611a8757600080fd5b8235611a928161191e565b915060208301356001600160401b0380821115611aae57600080fd5b9084019060608287031215611ac257600080fd5b611aca61185a565b823582811115611ad957600080fd5b83019150601f82018713611aec57600080fd5b611afb87833560208501611882565b815260208301356020820152604083013560408201528093505050509250929050565b600060208284031215611b3057600080fd5b81356001600160401b03811115611b4657600080fd5b611b52848285016118f7565b949350505050565b600060208284031215611b6c57600080fd5b5035919050565b6000815180845260005b81811015611b9957602081850181015186830182015201611b7d565b81811115611bab576000602083870101525b50601f01601f19169290920160200192915050565b608081526000611bd36080830187611b73565b6001600160a01b0395909516602083015250911515604083015260ff16606090910152919050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415611c3b57611c3b611c11565b5060010190565b838152606060208201526000611c5b6060830185611b73565b905060018060a01b0383166040830152949350505050565b600060208284031215611c8557600080fd5b5051919050565b600060208284031215611c9e57600080fd5b81516001600160c01b038116811461191757600080fd5b600060208284031215611cc757600080fd5b815161191781611a02565b60ff91909116815260200190565b60008219821115611cf357611cf3611c11565b500190565b600060408284031215611d0a57600080fd5b604080519081016001600160401b0381118282101715611d2c57611d2c611844565b6040528251611d3a8161191e565b815260208301516001600160601b0381168114611d5657600080fd5b60208201529392505050565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b60018060a01b0383168152604060208201526000825160606040840152611e0460a0840182611b73565b90506020840151606084015260408401516080840152809150509392505050565b6020815260006119176020830184611b73565b600181811c90821680611e4c57607f821691505b60208210811415611e6d57634e487b7160e01b600052602260045260246000fd5b50919050565b600082821015611e8557611e85611c11565b500390565b600061ffff80831681811415611ea257611ea2611c11565b600101939250505056fe661e3d4c4a13612fad1504b96725fcedfb885ff3ac490c18004f57ad9f4e8778a2646970667358221220ef907363ee3de9869000c678e77c4f050cefc7e3893e9deb2bdf588cb54f04ad64736f6c634300080c0033",
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
// Solidity: function taskManagers(uint256 ) view returns(string chainName, address taskManagerAddress, bool isActive, uint8 managerType)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerCaller) TaskManagers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ChainName          string
	TaskManagerAddress common.Address
	IsActive           bool
	ManagerType        uint8
}, error) {
	var out []interface{}
	err := _ContractOpenOracleServiceManager.contract.Call(opts, &out, "taskManagers", arg0)

	outstruct := new(struct {
		ChainName          string
		TaskManagerAddress common.Address
		IsActive           bool
		ManagerType        uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChainName = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.TaskManagerAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.IsActive = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.ManagerType = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}

// TaskManagers is a free data retrieval call binding the contract method 0xfd69a60f.
//
// Solidity: function taskManagers(uint256 ) view returns(string chainName, address taskManagerAddress, bool isActive, uint8 managerType)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerSession) TaskManagers(arg0 *big.Int) (struct {
	ChainName          string
	TaskManagerAddress common.Address
	IsActive           bool
	ManagerType        uint8
}, error) {
	return _ContractOpenOracleServiceManager.Contract.TaskManagers(&_ContractOpenOracleServiceManager.CallOpts, arg0)
}

// TaskManagers is a free data retrieval call binding the contract method 0xfd69a60f.
//
// Solidity: function taskManagers(uint256 ) view returns(string chainName, address taskManagerAddress, bool isActive, uint8 managerType)
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerCallerSession) TaskManagers(arg0 *big.Int) (struct {
	ChainName          string
	TaskManagerAddress common.Address
	IsActive           bool
	ManagerType        uint8
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

// AddTaskManager0 is a paid mutator transaction binding the contract method 0x868529c3.
//
// Solidity: function addTaskManager(string chainName, address taskManagerAddress, uint8 managerType) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerTransactor) AddTaskManager0(opts *bind.TransactOpts, chainName string, taskManagerAddress common.Address, managerType uint8) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.contract.Transact(opts, "addTaskManager0", chainName, taskManagerAddress, managerType)
}

// AddTaskManager0 is a paid mutator transaction binding the contract method 0x868529c3.
//
// Solidity: function addTaskManager(string chainName, address taskManagerAddress, uint8 managerType) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerSession) AddTaskManager0(chainName string, taskManagerAddress common.Address, managerType uint8) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.Contract.AddTaskManager0(&_ContractOpenOracleServiceManager.TransactOpts, chainName, taskManagerAddress, managerType)
}

// AddTaskManager0 is a paid mutator transaction binding the contract method 0x868529c3.
//
// Solidity: function addTaskManager(string chainName, address taskManagerAddress, uint8 managerType) returns()
func (_ContractOpenOracleServiceManager *ContractOpenOracleServiceManagerTransactorSession) AddTaskManager0(chainName string, taskManagerAddress common.Address, managerType uint8) (*types.Transaction, error) {
	return _ContractOpenOracleServiceManager.Contract.AddTaskManager0(&_ContractOpenOracleServiceManager.TransactOpts, chainName, taskManagerAddress, managerType)
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

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package generated

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IOdinDatasetReferenceData is an auto generated low-level Go binding around an user-defined struct.
type IOdinDatasetReferenceData struct {
	Count       *big.Int
	LastUpdated *big.Int
}

// OdinDatasetABI is the input ABI used to generate the binding from.
const OdinDatasetABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_models\",\"type\":\"string[]\"}],\"name\":\"getReferenceData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structIOdinDataset.ReferenceData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICacheBridge\",\"name\":\"_bridge\",\"type\":\"address\"}],\"name\":\"setBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OdinDatasetBin is the compiled bytecode used for deploying new contracts.
var OdinDatasetBin = "0x6080604052600167ffffffffffffffff811180156200001d57600080fd5b506040519080825280602002602001820160405280156200005357816020015b60608152602001906001900390816200003d5790505b50600390805190602001906200006b929190620004e8565b503480156200007957600080fd5b5060405162001e8638038062001e8683398181016040528101906200009f919062000707565b6000620000b1620004e060201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060405180606001604052806036815260200162001e50603691396003600081548110620001b957fe5b906000526020600020019080519060200190620001d89291906200054f565b506040518060600160405280600267ffffffffffffffff168152602001600060ff168152602001600060ff168152506002604051620002179062000880565b908152602001604051809103902060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548160ff021916908360ff16021790555060408201518160000160096101000a81548160ff021916908360ff1602179055509050506040518060600160405280600267ffffffffffffffff168152602001600060ff168152602001600160ff168152506002604051620002d79062000869565b908152602001604051809103902060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548160ff021916908360ff16021790555060408201518160000160096101000a81548160ff021916908360ff1602179055509050506040518060600160405280600267ffffffffffffffff168152602001600060ff168152602001600260ff16815250600260405162000397906200083b565b908152602001604051809103902060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548160ff021916908360ff16021790555060408201518160000160096101000a81548160ff021916908360ff1602179055509050506040518060600160405280600267ffffffffffffffff168152602001600060ff168152602001600360ff168152506002604051620004579062000852565b908152602001604051809103902060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548160ff021916908360ff16021790555060408201518160000160096101000a81548160ff021916908360ff16021790555090505050620008f0565b600033905090565b8280548282559060005260206000209081019282156200053c579160200282015b828111156200053b5782518290805190602001906200052a929190620005d6565b509160200191906001019062000509565b5b5090506200054b91906200065d565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200059257805160ff1916838001178555620005c3565b82800160010185558215620005c3579182015b82811115620005c2578251825591602001919060010190620005a5565b5b509050620005d2919062000685565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200061957805160ff19168380011785556200064a565b828001600101855582156200064a579182015b82811115620006495782518255916020019190600101906200062c565b5b50905062000659919062000685565b5090565b5b80821115620006815760008181620006779190620006a4565b506001016200065e565b5090565b5b80821115620006a057600081600090555060010162000686565b5090565b50805460018160011615610100020316600290046000825580601f10620006cc5750620006ed565b601f016020900490600052602060002090810190620006ec919062000685565b5b50565b6000815190506200070181620008d6565b92915050565b6000602082840312156200071a57600080fd5b60006200072a84828501620006f0565b91505092915050565b60006200074260098362000897565b91507f6950686f6e6520313100000000000000000000000000000000000000000000006000830152600982019050919050565b60006200078460098362000897565b91507f6950686f6e6520313200000000000000000000000000000000000000000000006000830152600982019050919050565b6000620007c660088362000897565b91507f6950686f6e6520580000000000000000000000000000000000000000000000006000830152600882019050919050565b60006200080860088362000897565b91507f6950686f6e6520380000000000000000000000000000000000000000000000006000830152600882019050919050565b6000620008488262000733565b9150819050919050565b60006200085f8262000775565b9150819050919050565b60006200087682620007b7565b9150819050919050565b60006200088d82620007f9565b9150819050919050565b600081905092915050565b6000620008af82620008b6565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b620008e181620008a2565b8114620008ed57600080fd5b50565b61155080620009006000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063715018a61461005c5780638da5cb5b146100665780638dd1480214610084578063d486151b146100a0578063f2fde38b146100d0575b600080fd5b6100646100ec565b005b61006e610226565b60405161007b9190611203565b60405180910390f35b61009e60048036038101906100999190610eaa565b61024f565b005b6100ba60048036038101906100b59190610e69565b61030f565b6040516100c7919061121e565b60405180910390f35b6100ea60048036038101906100e59190610e40565b6103fd565b005b6100f46105a6565b73ffffffffffffffffffffffffffffffffffffffff16610112610226565b73ffffffffffffffffffffffffffffffffffffffff1614610168576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161015f90611280565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6102576105a6565b73ffffffffffffffffffffffffffffffffffffffff16610275610226565b73ffffffffffffffffffffffffffffffffffffffff16146102cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102c290611280565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b606080825167ffffffffffffffff8111801561032a57600080fd5b5060405190808252806020026020018201604052801561036457816020015b610351610aa0565b8152602001906001900390816103495790505b50905060005b83518110156103f35760008061039286848151811061038557fe5b60200260200101516105ae565b915091508167ffffffffffffffff168484815181106103ad57fe5b602002602001015160000181815250508067ffffffffffffffff168484815181106103d457fe5b602002602001015160200181815250505050808060010191505061036a565b5080915050919050565b6104056105a6565b73ffffffffffffffffffffffffffffffffffffffff16610423610226565b73ffffffffffffffffffffffffffffffffffffffff1614610479576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161047090611280565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156104e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104e090611240565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600033905090565b6000806105b9610aba565b6040518060400160405280600181526020017f310000000000000000000000000000000000000000000000000000000000000081525081600001819052506001816060019067ffffffffffffffff16908167ffffffffffffffff16815250506001816080019067ffffffffffffffff16908167ffffffffffffffff1681525050600060028560405161064b91906111ec565b908152602001604051809103902090508060000160009054906101000a900467ffffffffffffffff16826020019067ffffffffffffffff16908167ffffffffffffffff168152505060038160000160089054906101000a900460ff1660ff16815481106106b457fe5b906000526020600020018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107525780601f1061072757610100808354040283529160200191610752565b820191906000526020600020905b81548152906001019060200180831161073557829003601f168201915b50505050508260400181905250610767610b07565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631102965c846040518263ffffffff1660e01b81526004016107c291906112a0565b60006040518083038186803b1580156107da57600080fd5b505afa1580156107ee573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906108179190610ed3565b9050610821610b6f565b61082e8260c0015161086e565b905080600001518360000160099054906101000a900460ff1660ff168151811061085457fe5b602002602001015182608001519550955050505050915091565b610876610b6f565b61087e610b82565b61088783610956565b905060006108948261097b565b905060608163ffffffff1667ffffffffffffffff811180156108b557600080fd5b506040519080825280602002602001820160405280156108e45781602001602082028036833780820191505090505b50905060005b8263ffffffff1681101561093c57610901846109ac565b82828151811061090d57fe5b602002602001019067ffffffffffffffff16908167ffffffffffffffff168152505080806001019150506108ea565b506040518060200160405280828152509350505050919050565b61095e610b82565b604051806040016040528060008152602001838152509050919050565b60006010610988836109e5565b61ffff1663ffffffff16901b905061099f826109e5565b61ffff1681179050919050565b600060206109b98361097b565b63ffffffff1667ffffffffffffffff16901b90506109d68261097b565b63ffffffff1681179050919050565b600060086109f283610a12565b60ff1661ffff16901b9050610a0682610a12565b60ff1681179050919050565b6000816001808260000151018260200151511015610a65576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a5c90611260565b60405180910390fd5b8360200151846000015181518110610a7957fe5b602001015160f81c60f81b60f81c9250808260000181815101915081815250505050919050565b604051806040016040528060008152602001600081525090565b6040518060a0016040528060608152602001600067ffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b6040518060e0016040528060608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600060ff168152602001606081525090565b6040518060200160405280606081525090565b604051806040016040528060008152602001606081525090565b600081359050610bab816114be565b92915050565b600082601f830112610bc257600080fd5b8135610bd5610bd0826112ef565b6112c2565b9150818183526020840193506020810190508360005b83811015610c1b5781358601610c018882610c8e565b845260208401935060208301925050600181019050610beb565b5050505092915050565b600082601f830112610c3657600080fd5b8151610c49610c4482611317565b6112c2565b91508082526020830160208301858383011115610c6557600080fd5b610c7083828461147a565b50505092915050565b600081359050610c88816114d5565b92915050565b600082601f830112610c9f57600080fd5b8135610cb2610cad82611343565b6112c2565b91508082526020830160208301858383011115610cce57600080fd5b610cd983828461146b565b50505092915050565b600082601f830112610cf357600080fd5b8151610d06610d0182611343565b6112c2565b91508082526020830160208301858383011115610d2257600080fd5b610d2d83828461147a565b50505092915050565b600060e08284031215610d4857600080fd5b610d5260e06112c2565b9050600082015167ffffffffffffffff811115610d6e57600080fd5b610d7a84828501610ce2565b6000830152506020610d8e84828501610e16565b6020830152506040610da284828501610e16565b6040830152506060610db684828501610e16565b6060830152506080610dca84828501610e16565b60808301525060a0610dde84828501610e2b565b60a08301525060c082015167ffffffffffffffff811115610dfe57600080fd5b610e0a84828501610c25565b60c08301525092915050565b600081519050610e25816114ec565b92915050565b600081519050610e3a81611503565b92915050565b600060208284031215610e5257600080fd5b6000610e6084828501610b9c565b91505092915050565b600060208284031215610e7b57600080fd5b600082013567ffffffffffffffff811115610e9557600080fd5b610ea184828501610bb1565b91505092915050565b600060208284031215610ebc57600080fd5b6000610eca84828501610c79565b91505092915050565b600060208284031215610ee557600080fd5b600082015167ffffffffffffffff811115610eff57600080fd5b610f0b84828501610d36565b91505092915050565b6000610f208383611122565b60408301905092915050565b610f35816113fc565b82525050565b6000610f468261137f565b610f5081856113ad565b9350610f5b8361136f565b8060005b83811015610f8c578151610f738882610f14565b9750610f7e836113a0565b925050600181019050610f5f565b5085935050505092915050565b6000610fa48261138a565b610fae81856113be565b9350610fbe81856020860161147a565b610fc7816114ad565b840191505092915050565b6000610fdd82611395565b610fe781856113cf565b9350610ff781856020860161147a565b611000816114ad565b840191505092915050565b600061101682611395565b61102081856113f1565b935061103081856020860161147a565b80840191505092915050565b60006110496026836113e0565b91507f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008301527f64647265737300000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006110af6011836113e0565b91507f4f62693a204f7574206f662072616e67650000000000000000000000000000006000830152602082019050919050565b60006110ef6020836113e0565b91507f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726000830152602082019050919050565b60408201600082015161113860008501826111ce565b50602082015161114b60208501826111ce565b50505050565b600060a083016000830151848203600086015261116e8282610fd2565b915050602083015161118360208601826111dd565b506040830151848203604086015261119b8282610f99565b91505060608301516111b060608601826111dd565b5060808301516111c360808601826111dd565b508091505092915050565b6111d781611440565b82525050565b6111e68161144a565b82525050565b60006111f8828461100b565b915081905092915050565b60006020820190506112186000830184610f2c565b92915050565b600060208201905081810360008301526112388184610f3b565b905092915050565b600060208201905081810360008301526112598161103c565b9050919050565b60006020820190508181036000830152611279816110a2565b9050919050565b60006020820190508181036000830152611299816110e2565b9050919050565b600060208201905081810360008301526112ba8184611151565b905092915050565b6000604051905081810181811067ffffffffffffffff821117156112e557600080fd5b8060405250919050565b600067ffffffffffffffff82111561130657600080fd5b602082029050602081019050919050565b600067ffffffffffffffff82111561132e57600080fd5b601f19601f8301169050602081019050919050565b600067ffffffffffffffff82111561135a57600080fd5b601f19601f8301169050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600061140782611420565b9050919050565b6000611419826113fc565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b8381101561149857808201518184015260208101905061147d565b838111156114a7576000848401525b50505050565b6000601f19601f8301169050919050565b6114c7816113fc565b81146114d257600080fd5b50565b6114de8161140e565b81146114e957600080fd5b50565b6114f58161144a565b811461150057600080fd5b50565b61150c8161145e565b811461151757600080fd5b5056fea26469706673582212207de5d1d021212b28daa9110cddd191b63615904d7b362aecfb51417e8942036464736f6c6343000700003300000004000000086950686f6e652038000000086950686f6e652058000000096950686f6e65203131000000096950686f6e65203132"

// DeployOdinDataset deploys a new Ethereum contract, binding an instance of OdinDataset to it.
func DeployOdinDataset(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge common.Address) (common.Address, *types.Transaction, *OdinDataset, error) {
	parsed, err := abi.JSON(strings.NewReader(OdinDatasetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OdinDatasetBin), backend, _bridge)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OdinDataset{OdinDatasetCaller: OdinDatasetCaller{contract: contract}, OdinDatasetTransactor: OdinDatasetTransactor{contract: contract}, OdinDatasetFilterer: OdinDatasetFilterer{contract: contract}}, nil
}

// OdinDataset is an auto generated Go binding around an Ethereum contract.
type OdinDataset struct {
	OdinDatasetCaller     // Read-only binding to the contract
	OdinDatasetTransactor // Write-only binding to the contract
	OdinDatasetFilterer   // Log filterer for contract events
}

// OdinDatasetCaller is an auto generated read-only Go binding around an Ethereum contract.
type OdinDatasetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OdinDatasetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OdinDatasetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OdinDatasetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OdinDatasetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OdinDatasetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OdinDatasetSession struct {
	Contract     *OdinDataset      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OdinDatasetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OdinDatasetCallerSession struct {
	Contract *OdinDatasetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// OdinDatasetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OdinDatasetTransactorSession struct {
	Contract     *OdinDatasetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// OdinDatasetRaw is an auto generated low-level Go binding around an Ethereum contract.
type OdinDatasetRaw struct {
	Contract *OdinDataset // Generic contract binding to access the raw methods on
}

// OdinDatasetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OdinDatasetCallerRaw struct {
	Contract *OdinDatasetCaller // Generic read-only contract binding to access the raw methods on
}

// OdinDatasetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OdinDatasetTransactorRaw struct {
	Contract *OdinDatasetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOdinDataset creates a new instance of OdinDataset, bound to a specific deployed contract.
func NewOdinDataset(address common.Address, backend bind.ContractBackend) (*OdinDataset, error) {
	contract, err := bindOdinDataset(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OdinDataset{OdinDatasetCaller: OdinDatasetCaller{contract: contract}, OdinDatasetTransactor: OdinDatasetTransactor{contract: contract}, OdinDatasetFilterer: OdinDatasetFilterer{contract: contract}}, nil
}

// NewOdinDatasetCaller creates a new read-only instance of OdinDataset, bound to a specific deployed contract.
func NewOdinDatasetCaller(address common.Address, caller bind.ContractCaller) (*OdinDatasetCaller, error) {
	contract, err := bindOdinDataset(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OdinDatasetCaller{contract: contract}, nil
}

// NewOdinDatasetTransactor creates a new write-only instance of OdinDataset, bound to a specific deployed contract.
func NewOdinDatasetTransactor(address common.Address, transactor bind.ContractTransactor) (*OdinDatasetTransactor, error) {
	contract, err := bindOdinDataset(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OdinDatasetTransactor{contract: contract}, nil
}

// NewOdinDatasetFilterer creates a new log filterer instance of OdinDataset, bound to a specific deployed contract.
func NewOdinDatasetFilterer(address common.Address, filterer bind.ContractFilterer) (*OdinDatasetFilterer, error) {
	contract, err := bindOdinDataset(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OdinDatasetFilterer{contract: contract}, nil
}

// bindOdinDataset binds a generic wrapper to an already deployed contract.
func bindOdinDataset(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OdinDatasetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OdinDataset *OdinDatasetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OdinDataset.Contract.OdinDatasetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OdinDataset *OdinDatasetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OdinDataset.Contract.OdinDatasetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OdinDataset *OdinDatasetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OdinDataset.Contract.OdinDatasetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OdinDataset *OdinDatasetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OdinDataset.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OdinDataset *OdinDatasetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OdinDataset.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OdinDataset *OdinDatasetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OdinDataset.Contract.contract.Transact(opts, method, params...)
}

// GetReferenceData is a free data retrieval call binding the contract method 0xd486151b.
//
// Solidity: function getReferenceData(string[] _models) view returns((uint256,uint256)[])
func (_OdinDataset *OdinDatasetCaller) GetReferenceData(opts *bind.CallOpts, _models []string) ([]IOdinDatasetReferenceData, error) {
	var out []interface{}
	err := _OdinDataset.contract.Call(opts, &out, "getReferenceData", _models)

	if err != nil {
		return *new([]IOdinDatasetReferenceData), err
	}

	out0 := *abi.ConvertType(out[0], new([]IOdinDatasetReferenceData)).(*[]IOdinDatasetReferenceData)

	return out0, err

}

// GetReferenceData is a free data retrieval call binding the contract method 0xd486151b.
//
// Solidity: function getReferenceData(string[] _models) view returns((uint256,uint256)[])
func (_OdinDataset *OdinDatasetSession) GetReferenceData(_models []string) ([]IOdinDatasetReferenceData, error) {
	return _OdinDataset.Contract.GetReferenceData(&_OdinDataset.CallOpts, _models)
}

// GetReferenceData is a free data retrieval call binding the contract method 0xd486151b.
//
// Solidity: function getReferenceData(string[] _models) view returns((uint256,uint256)[])
func (_OdinDataset *OdinDatasetCallerSession) GetReferenceData(_models []string) ([]IOdinDatasetReferenceData, error) {
	return _OdinDataset.Contract.GetReferenceData(&_OdinDataset.CallOpts, _models)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OdinDataset *OdinDatasetCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OdinDataset.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OdinDataset *OdinDatasetSession) Owner() (common.Address, error) {
	return _OdinDataset.Contract.Owner(&_OdinDataset.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OdinDataset *OdinDatasetCallerSession) Owner() (common.Address, error) {
	return _OdinDataset.Contract.Owner(&_OdinDataset.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OdinDataset *OdinDatasetTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OdinDataset.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OdinDataset *OdinDatasetSession) RenounceOwnership() (*types.Transaction, error) {
	return _OdinDataset.Contract.RenounceOwnership(&_OdinDataset.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OdinDataset *OdinDatasetTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OdinDataset.Contract.RenounceOwnership(&_OdinDataset.TransactOpts)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address _bridge) returns()
func (_OdinDataset *OdinDatasetTransactor) SetBridge(opts *bind.TransactOpts, _bridge common.Address) (*types.Transaction, error) {
	return _OdinDataset.contract.Transact(opts, "setBridge", _bridge)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address _bridge) returns()
func (_OdinDataset *OdinDatasetSession) SetBridge(_bridge common.Address) (*types.Transaction, error) {
	return _OdinDataset.Contract.SetBridge(&_OdinDataset.TransactOpts, _bridge)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address _bridge) returns()
func (_OdinDataset *OdinDatasetTransactorSession) SetBridge(_bridge common.Address) (*types.Transaction, error) {
	return _OdinDataset.Contract.SetBridge(&_OdinDataset.TransactOpts, _bridge)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OdinDataset *OdinDatasetTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OdinDataset.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OdinDataset *OdinDatasetSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OdinDataset.Contract.TransferOwnership(&_OdinDataset.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OdinDataset *OdinDatasetTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OdinDataset.Contract.TransferOwnership(&_OdinDataset.TransactOpts, newOwner)
}

// OdinDatasetOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OdinDataset contract.
type OdinDatasetOwnershipTransferredIterator struct {
	Event *OdinDatasetOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OdinDatasetOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OdinDatasetOwnershipTransferred)
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
		it.Event = new(OdinDatasetOwnershipTransferred)
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
func (it *OdinDatasetOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OdinDatasetOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OdinDatasetOwnershipTransferred represents a OwnershipTransferred event raised by the OdinDataset contract.
type OdinDatasetOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OdinDataset *OdinDatasetFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OdinDatasetOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OdinDataset.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OdinDatasetOwnershipTransferredIterator{contract: _OdinDataset.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OdinDataset *OdinDatasetFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OdinDatasetOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OdinDataset.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OdinDatasetOwnershipTransferred)
				if err := _OdinDataset.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OdinDataset *OdinDatasetFilterer) ParseOwnershipTransferred(log types.Log) (*OdinDatasetOwnershipTransferred, error) {
	event := new(OdinDatasetOwnershipTransferred)
	if err := _OdinDataset.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

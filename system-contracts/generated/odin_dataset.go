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

// OdinDatasetABI is the input ABI used to generate the binding from.
const OdinDatasetABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_oracleScriptId\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"getDeviceCount\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleScriptId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICacheBridge\",\"name\":\"_bridge\",\"type\":\"address\"}],\"name\":\"setBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OdinDatasetBin is the compiled bytecode used for deploying new contracts.
var OdinDatasetBin = "0x60806040523480156200001157600080fd5b50604051620014873803806200148783398181016040528101906200003791906200018f565b6000620000496200015960201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600060146101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050506200024c565b600033905090565b600081519050620001728162000218565b92915050565b600081519050620001898162000232565b92915050565b60008060408385031215620001a357600080fd5b6000620001b38582860162000161565b9250506020620001c68582860162000178565b9150509250929050565b6000620001dd82620001e4565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600067ffffffffffffffff82169050919050565b6200022381620001d0565b81146200022f57600080fd5b50565b6200023d8162000204565b81146200024957600080fd5b50565b61122b806200025c6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806347613f4514610067578063715018a6146100975780638da5cb5b146100a15780638dd14802146100bf578063dda683dc146100db578063f2fde38b146100f9575b600080fd5b610081600480360381019061007c9190610c78565b610115565b60405161008e9190610f67565b60405180910390f35b61009f610273565b005b6100a96103ad565b6040516100b69190610f2a565b60405180910390f35b6100d960048036038101906100d49190610cb9565b6103d6565b005b6100e3610496565b6040516100f09190611009565b60405180910390f35b610113600480360381019061010e9190610c4f565b6104b0565b005b606061011f610885565b6101276108d2565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ad373732856040518263ffffffff1660e01b81526004016101829190610f45565b600060405180830381600087803b15801561019c57600080fd5b505af11580156101b0573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906101d99190610ce2565b91509150600060149054906101000a900467ffffffffffffffff1667ffffffffffffffff16826020015167ffffffffffffffff161461024d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161024490610fc9565b60405180910390fd5b61025561093a565b6102628260c00151610659565b905080600001519350505050919050565b61027b610697565b73ffffffffffffffffffffffffffffffffffffffff166102996103ad565b73ffffffffffffffffffffffffffffffffffffffff16146102ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102e690610fe9565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6103de610697565b73ffffffffffffffffffffffffffffffffffffffff166103fc6103ad565b73ffffffffffffffffffffffffffffffffffffffff1614610452576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161044990610fe9565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600060149054906101000a900467ffffffffffffffff1681565b6104b8610697565b73ffffffffffffffffffffffffffffffffffffffff166104d66103ad565b73ffffffffffffffffffffffffffffffffffffffff161461052c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161052390610fe9565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561059c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161059390610f89565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b61066161093a565b61066961094d565b6106728361069f565b905061067c61093a565b610685826106c4565b81600001819052508092505050919050565b600033905090565b6106a761094d565b604051806040016040528060008152602001838152509050919050565b60606106cf826106d6565b9050919050565b60606106e182610799565b63ffffffff1667ffffffffffffffff811180156106fd57600080fd5b506040519080825280601f01601f1916602001820160405280156107305781602001600182028036833780820191505090505b50905060005b815181101561079357610748836107ca565b60f81b82828151811061075757fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080600101915050610736565b50919050565b600060106107a683610858565b61ffff1663ffffffff16901b90506107bd82610858565b61ffff1681179050919050565b600081600180826000015101826020015151101561081d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161081490610fa9565b60405180910390fd5b836020015184600001518151811061083157fe5b602001015160f81c60f81b60f81c9250808260000181815101915081815250505050919050565b60006008610865836107ca565b60ff1661ffff16901b9050610879826107ca565b60ff1681179050919050565b6040518060a0016040528060608152602001600067ffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b6040518060e0016040528060608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600060ff168152602001606081525090565b6040518060200160405280606081525090565b604051806040016040528060008152602001606081525090565b60008135905061097681611199565b92915050565b600082601f83011261098d57600080fd5b81356109a061099b82611051565b611024565b915080825260208301602083018583830111156109bc57600080fd5b6109c7838284611146565b50505092915050565b600082601f8301126109e157600080fd5b81516109f46109ef82611051565b611024565b91508082526020830160208301858383011115610a1057600080fd5b610a1b838284611155565b50505092915050565b600081359050610a33816111b0565b92915050565b600082601f830112610a4a57600080fd5b8151610a5d610a588261107d565b611024565b91508082526020830160208301858383011115610a7957600080fd5b610a84838284611155565b50505092915050565b600060a08284031215610a9f57600080fd5b610aa960a0611024565b9050600082015167ffffffffffffffff811115610ac557600080fd5b610ad184828501610a39565b6000830152506020610ae584828501610c25565b602083015250604082015167ffffffffffffffff811115610b0557600080fd5b610b11848285016109d0565b6040830152506060610b2584828501610c25565b6060830152506080610b3984828501610c25565b60808301525092915050565b600060e08284031215610b5757600080fd5b610b6160e0611024565b9050600082015167ffffffffffffffff811115610b7d57600080fd5b610b8984828501610a39565b6000830152506020610b9d84828501610c25565b6020830152506040610bb184828501610c25565b6040830152506060610bc584828501610c25565b6060830152506080610bd984828501610c25565b60808301525060a0610bed84828501610c3a565b60a08301525060c082015167ffffffffffffffff811115610c0d57600080fd5b610c19848285016109d0565b60c08301525092915050565b600081519050610c34816111c7565b92915050565b600081519050610c49816111de565b92915050565b600060208284031215610c6157600080fd5b6000610c6f84828501610967565b91505092915050565b600060208284031215610c8a57600080fd5b600082013567ffffffffffffffff811115610ca457600080fd5b610cb08482850161097c565b91505092915050565b600060208284031215610ccb57600080fd5b6000610cd984828501610a24565b91505092915050565b60008060408385031215610cf557600080fd5b600083015167ffffffffffffffff811115610d0f57600080fd5b610d1b85828601610a8d565b925050602083015167ffffffffffffffff811115610d3857600080fd5b610d4485828601610b45565b9150509250929050565b610d57816110e1565b82525050565b6000610d68826110a9565b610d7281856110bf565b9350610d82818560208601611155565b610d8b81611188565b840191505092915050565b6000610da1826110b4565b610dab81856110d0565b9350610dbb818560208601611155565b610dc481611188565b840191505092915050565b6000610ddc6026836110d0565b91507f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008301527f64647265737300000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000610e426011836110d0565b91507f4f62693a204f7574206f662072616e67650000000000000000000000000000006000830152602082019050919050565b6000610e826035836110d0565b91507f4552524f525f4f5241434c455f5343524950545f49445f444f45535f4e4f545f60008301527f4d415443485f574954485f5448455f434f4e46494700000000000000000000006020830152604082019050919050565b6000610ee86020836110d0565b91507f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726000830152602082019050919050565b610f2481611125565b82525050565b6000602082019050610f3f6000830184610d4e565b92915050565b60006020820190508181036000830152610f5f8184610d5d565b905092915050565b60006020820190508181036000830152610f818184610d96565b905092915050565b60006020820190508181036000830152610fa281610dcf565b9050919050565b60006020820190508181036000830152610fc281610e35565b9050919050565b60006020820190508181036000830152610fe281610e75565b9050919050565b6000602082019050818103600083015261100281610edb565b9050919050565b600060208201905061101e6000830184610f1b565b92915050565b6000604051905081810181811067ffffffffffffffff8211171561104757600080fd5b8060405250919050565b600067ffffffffffffffff82111561106857600080fd5b601f19601f8301169050602081019050919050565b600067ffffffffffffffff82111561109457600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b60006110ec82611105565b9050919050565b60006110fe826110e1565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600067ffffffffffffffff82169050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b83811015611173578082015181840152602081019050611158565b83811115611182576000848401525b50505050565b6000601f19601f8301169050919050565b6111a2816110e1565b81146111ad57600080fd5b50565b6111b9816110f3565b81146111c457600080fd5b50565b6111d081611125565b81146111db57600080fd5b50565b6111e781611139565b81146111f257600080fd5b5056fea264697066735822122037ba2632cfd29469bed9b639100af43ee92965e02657a25e3790edb6eddda0da64736f6c63430007000033"

// DeployOdinDataset deploys a new Ethereum contract, binding an instance of OdinDataset to it.
func DeployOdinDataset(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge common.Address, _oracleScriptId uint64) (common.Address, *types.Transaction, *OdinDataset, error) {
	parsed, err := abi.JSON(strings.NewReader(OdinDatasetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OdinDatasetBin), backend, _bridge, _oracleScriptId)
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

// OracleScriptId is a free data retrieval call binding the contract method 0xdda683dc.
//
// Solidity: function oracleScriptId() view returns(uint64)
func (_OdinDataset *OdinDatasetCaller) OracleScriptId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OdinDataset.contract.Call(opts, &out, "oracleScriptId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// OracleScriptId is a free data retrieval call binding the contract method 0xdda683dc.
//
// Solidity: function oracleScriptId() view returns(uint64)
func (_OdinDataset *OdinDatasetSession) OracleScriptId() (uint64, error) {
	return _OdinDataset.Contract.OracleScriptId(&_OdinDataset.CallOpts)
}

// OracleScriptId is a free data retrieval call binding the contract method 0xdda683dc.
//
// Solidity: function oracleScriptId() view returns(uint64)
func (_OdinDataset *OdinDatasetCallerSession) OracleScriptId() (uint64, error) {
	return _OdinDataset.Contract.OracleScriptId(&_OdinDataset.CallOpts)
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

// GetDeviceCount is a paid mutator transaction binding the contract method 0x47613f45.
//
// Solidity: function getDeviceCount(bytes _data) returns(string)
func (_OdinDataset *OdinDatasetTransactor) GetDeviceCount(opts *bind.TransactOpts, _data []byte) (*types.Transaction, error) {
	return _OdinDataset.contract.Transact(opts, "getDeviceCount", _data)
}

// GetDeviceCount is a paid mutator transaction binding the contract method 0x47613f45.
//
// Solidity: function getDeviceCount(bytes _data) returns(string)
func (_OdinDataset *OdinDatasetSession) GetDeviceCount(_data []byte) (*types.Transaction, error) {
	return _OdinDataset.Contract.GetDeviceCount(&_OdinDataset.TransactOpts, _data)
}

// GetDeviceCount is a paid mutator transaction binding the contract method 0x47613f45.
//
// Solidity: function getDeviceCount(bytes _data) returns(string)
func (_OdinDataset *OdinDatasetTransactorSession) GetDeviceCount(_data []byte) (*types.Transaction, error) {
	return _OdinDataset.Contract.GetDeviceCount(&_OdinDataset.TransactOpts, _data)
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

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// PromptVaultMetaData contains all meta data concerning the PromptVault contract.
var PromptVaultMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_ticketPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_signerAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"MAX_TICKET_PRICE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_TICKET_PRICE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimPrize\",\"inputs\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getNonce\",\"inputs\":[{\"name\":\"player\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVaultState\",\"inputs\":[],\"outputs\":[{\"name\":\"_prizePool\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_protocolRevenue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_ticketPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_signerAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"play\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"prizePool\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"protocolRevenue\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"seedPrizePool\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setSigner\",\"inputs\":[{\"name\":\"_newSigner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTicketPrice\",\"inputs\":[{\"name\":\"_newPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"signerAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ticketPrice\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawRevenue\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FundsReceived\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PrizeClaimed\",\"inputs\":[{\"name\":\"winner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RevenueWithdrawn\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignerUpdated\",\"inputs\":[{\"name\":\"oldSigner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newSigner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TicketPriceUpdated\",\"inputs\":[{\"name\":\"oldPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TicketPurchased\",\"inputs\":[{\"name\":\"player\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"prizePoolNew\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
}

// PromptVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use PromptVaultMetaData.ABI instead.
var PromptVaultABI = PromptVaultMetaData.ABI

// PromptVault is an auto generated Go binding around an Ethereum contract.
type PromptVault struct {
	PromptVaultCaller     // Read-only binding to the contract
	PromptVaultTransactor // Write-only binding to the contract
	PromptVaultFilterer   // Log filterer for contract events
}

// PromptVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type PromptVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PromptVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PromptVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PromptVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PromptVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PromptVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PromptVaultSession struct {
	Contract     *PromptVault      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PromptVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PromptVaultCallerSession struct {
	Contract *PromptVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PromptVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PromptVaultTransactorSession struct {
	Contract     *PromptVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PromptVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type PromptVaultRaw struct {
	Contract *PromptVault // Generic contract binding to access the raw methods on
}

// PromptVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PromptVaultCallerRaw struct {
	Contract *PromptVaultCaller // Generic read-only contract binding to access the raw methods on
}

// PromptVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PromptVaultTransactorRaw struct {
	Contract *PromptVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPromptVault creates a new instance of PromptVault, bound to a specific deployed contract.
func NewPromptVault(address common.Address, backend bind.ContractBackend) (*PromptVault, error) {
	contract, err := bindPromptVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PromptVault{PromptVaultCaller: PromptVaultCaller{contract: contract}, PromptVaultTransactor: PromptVaultTransactor{contract: contract}, PromptVaultFilterer: PromptVaultFilterer{contract: contract}}, nil
}

// NewPromptVaultCaller creates a new read-only instance of PromptVault, bound to a specific deployed contract.
func NewPromptVaultCaller(address common.Address, caller bind.ContractCaller) (*PromptVaultCaller, error) {
	contract, err := bindPromptVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PromptVaultCaller{contract: contract}, nil
}

// NewPromptVaultTransactor creates a new write-only instance of PromptVault, bound to a specific deployed contract.
func NewPromptVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*PromptVaultTransactor, error) {
	contract, err := bindPromptVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PromptVaultTransactor{contract: contract}, nil
}

// NewPromptVaultFilterer creates a new log filterer instance of PromptVault, bound to a specific deployed contract.
func NewPromptVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*PromptVaultFilterer, error) {
	contract, err := bindPromptVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PromptVaultFilterer{contract: contract}, nil
}

// bindPromptVault binds a generic wrapper to an already deployed contract.
func bindPromptVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PromptVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PromptVault *PromptVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PromptVault.Contract.PromptVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PromptVault *PromptVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PromptVault.Contract.PromptVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PromptVault *PromptVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PromptVault.Contract.PromptVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PromptVault *PromptVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PromptVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PromptVault *PromptVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PromptVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PromptVault *PromptVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PromptVault.Contract.contract.Transact(opts, method, params...)
}

// MAXTICKETPRICE is a free data retrieval call binding the contract method 0xb5e41287.
//
// Solidity: function MAX_TICKET_PRICE() view returns(uint256)
func (_PromptVault *PromptVaultCaller) MAXTICKETPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PromptVault.contract.Call(opts, &out, "MAX_TICKET_PRICE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTICKETPRICE is a free data retrieval call binding the contract method 0xb5e41287.
//
// Solidity: function MAX_TICKET_PRICE() view returns(uint256)
func (_PromptVault *PromptVaultSession) MAXTICKETPRICE() (*big.Int, error) {
	return _PromptVault.Contract.MAXTICKETPRICE(&_PromptVault.CallOpts)
}

// MAXTICKETPRICE is a free data retrieval call binding the contract method 0xb5e41287.
//
// Solidity: function MAX_TICKET_PRICE() view returns(uint256)
func (_PromptVault *PromptVaultCallerSession) MAXTICKETPRICE() (*big.Int, error) {
	return _PromptVault.Contract.MAXTICKETPRICE(&_PromptVault.CallOpts)
}

// MINTICKETPRICE is a free data retrieval call binding the contract method 0x1af2a065.
//
// Solidity: function MIN_TICKET_PRICE() view returns(uint256)
func (_PromptVault *PromptVaultCaller) MINTICKETPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PromptVault.contract.Call(opts, &out, "MIN_TICKET_PRICE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTICKETPRICE is a free data retrieval call binding the contract method 0x1af2a065.
//
// Solidity: function MIN_TICKET_PRICE() view returns(uint256)
func (_PromptVault *PromptVaultSession) MINTICKETPRICE() (*big.Int, error) {
	return _PromptVault.Contract.MINTICKETPRICE(&_PromptVault.CallOpts)
}

// MINTICKETPRICE is a free data retrieval call binding the contract method 0x1af2a065.
//
// Solidity: function MIN_TICKET_PRICE() view returns(uint256)
func (_PromptVault *PromptVaultCallerSession) MINTICKETPRICE() (*big.Int, error) {
	return _PromptVault.Contract.MINTICKETPRICE(&_PromptVault.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address player) view returns(uint256)
func (_PromptVault *PromptVaultCaller) GetNonce(opts *bind.CallOpts, player common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PromptVault.contract.Call(opts, &out, "getNonce", player)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address player) view returns(uint256)
func (_PromptVault *PromptVaultSession) GetNonce(player common.Address) (*big.Int, error) {
	return _PromptVault.Contract.GetNonce(&_PromptVault.CallOpts, player)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address player) view returns(uint256)
func (_PromptVault *PromptVaultCallerSession) GetNonce(player common.Address) (*big.Int, error) {
	return _PromptVault.Contract.GetNonce(&_PromptVault.CallOpts, player)
}

// GetVaultState is a free data retrieval call binding the contract method 0x4a8c110a.
//
// Solidity: function getVaultState() view returns(uint256 _prizePool, uint256 _protocolRevenue, uint256 _ticketPrice, address _signerAddress)
func (_PromptVault *PromptVaultCaller) GetVaultState(opts *bind.CallOpts) (struct {
	PrizePool       *big.Int
	ProtocolRevenue *big.Int
	TicketPrice     *big.Int
	SignerAddress   common.Address
}, error) {
	var out []interface{}
	err := _PromptVault.contract.Call(opts, &out, "getVaultState")

	outstruct := new(struct {
		PrizePool       *big.Int
		ProtocolRevenue *big.Int
		TicketPrice     *big.Int
		SignerAddress   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PrizePool = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ProtocolRevenue = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TicketPrice = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SignerAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetVaultState is a free data retrieval call binding the contract method 0x4a8c110a.
//
// Solidity: function getVaultState() view returns(uint256 _prizePool, uint256 _protocolRevenue, uint256 _ticketPrice, address _signerAddress)
func (_PromptVault *PromptVaultSession) GetVaultState() (struct {
	PrizePool       *big.Int
	ProtocolRevenue *big.Int
	TicketPrice     *big.Int
	SignerAddress   common.Address
}, error) {
	return _PromptVault.Contract.GetVaultState(&_PromptVault.CallOpts)
}

// GetVaultState is a free data retrieval call binding the contract method 0x4a8c110a.
//
// Solidity: function getVaultState() view returns(uint256 _prizePool, uint256 _protocolRevenue, uint256 _ticketPrice, address _signerAddress)
func (_PromptVault *PromptVaultCallerSession) GetVaultState() (struct {
	PrizePool       *big.Int
	ProtocolRevenue *big.Int
	TicketPrice     *big.Int
	SignerAddress   common.Address
}, error) {
	return _PromptVault.Contract.GetVaultState(&_PromptVault.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_PromptVault *PromptVaultCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PromptVault.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_PromptVault *PromptVaultSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _PromptVault.Contract.Nonces(&_PromptVault.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_PromptVault *PromptVaultCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _PromptVault.Contract.Nonces(&_PromptVault.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PromptVault *PromptVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PromptVault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PromptVault *PromptVaultSession) Owner() (common.Address, error) {
	return _PromptVault.Contract.Owner(&_PromptVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PromptVault *PromptVaultCallerSession) Owner() (common.Address, error) {
	return _PromptVault.Contract.Owner(&_PromptVault.CallOpts)
}

// PrizePool is a free data retrieval call binding the contract method 0x719ce73e.
//
// Solidity: function prizePool() view returns(uint256)
func (_PromptVault *PromptVaultCaller) PrizePool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PromptVault.contract.Call(opts, &out, "prizePool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrizePool is a free data retrieval call binding the contract method 0x719ce73e.
//
// Solidity: function prizePool() view returns(uint256)
func (_PromptVault *PromptVaultSession) PrizePool() (*big.Int, error) {
	return _PromptVault.Contract.PrizePool(&_PromptVault.CallOpts)
}

// PrizePool is a free data retrieval call binding the contract method 0x719ce73e.
//
// Solidity: function prizePool() view returns(uint256)
func (_PromptVault *PromptVaultCallerSession) PrizePool() (*big.Int, error) {
	return _PromptVault.Contract.PrizePool(&_PromptVault.CallOpts)
}

// ProtocolRevenue is a free data retrieval call binding the contract method 0x7af3816c.
//
// Solidity: function protocolRevenue() view returns(uint256)
func (_PromptVault *PromptVaultCaller) ProtocolRevenue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PromptVault.contract.Call(opts, &out, "protocolRevenue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolRevenue is a free data retrieval call binding the contract method 0x7af3816c.
//
// Solidity: function protocolRevenue() view returns(uint256)
func (_PromptVault *PromptVaultSession) ProtocolRevenue() (*big.Int, error) {
	return _PromptVault.Contract.ProtocolRevenue(&_PromptVault.CallOpts)
}

// ProtocolRevenue is a free data retrieval call binding the contract method 0x7af3816c.
//
// Solidity: function protocolRevenue() view returns(uint256)
func (_PromptVault *PromptVaultCallerSession) ProtocolRevenue() (*big.Int, error) {
	return _PromptVault.Contract.ProtocolRevenue(&_PromptVault.CallOpts)
}

// SignerAddress is a free data retrieval call binding the contract method 0x5b7633d0.
//
// Solidity: function signerAddress() view returns(address)
func (_PromptVault *PromptVaultCaller) SignerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PromptVault.contract.Call(opts, &out, "signerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerAddress is a free data retrieval call binding the contract method 0x5b7633d0.
//
// Solidity: function signerAddress() view returns(address)
func (_PromptVault *PromptVaultSession) SignerAddress() (common.Address, error) {
	return _PromptVault.Contract.SignerAddress(&_PromptVault.CallOpts)
}

// SignerAddress is a free data retrieval call binding the contract method 0x5b7633d0.
//
// Solidity: function signerAddress() view returns(address)
func (_PromptVault *PromptVaultCallerSession) SignerAddress() (common.Address, error) {
	return _PromptVault.Contract.SignerAddress(&_PromptVault.CallOpts)
}

// TicketPrice is a free data retrieval call binding the contract method 0x1209b1f6.
//
// Solidity: function ticketPrice() view returns(uint256)
func (_PromptVault *PromptVaultCaller) TicketPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PromptVault.contract.Call(opts, &out, "ticketPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TicketPrice is a free data retrieval call binding the contract method 0x1209b1f6.
//
// Solidity: function ticketPrice() view returns(uint256)
func (_PromptVault *PromptVaultSession) TicketPrice() (*big.Int, error) {
	return _PromptVault.Contract.TicketPrice(&_PromptVault.CallOpts)
}

// TicketPrice is a free data retrieval call binding the contract method 0x1209b1f6.
//
// Solidity: function ticketPrice() view returns(uint256)
func (_PromptVault *PromptVaultCallerSession) TicketPrice() (*big.Int, error) {
	return _PromptVault.Contract.TicketPrice(&_PromptVault.CallOpts)
}

// ClaimPrize is a paid mutator transaction binding the contract method 0x0dcba9c5.
//
// Solidity: function claimPrize(bytes signature, uint256 signedAmount) returns()
func (_PromptVault *PromptVaultTransactor) ClaimPrize(opts *bind.TransactOpts, signature []byte, signedAmount *big.Int) (*types.Transaction, error) {
	return _PromptVault.contract.Transact(opts, "claimPrize", signature, signedAmount)
}

// ClaimPrize is a paid mutator transaction binding the contract method 0x0dcba9c5.
//
// Solidity: function claimPrize(bytes signature, uint256 signedAmount) returns()
func (_PromptVault *PromptVaultSession) ClaimPrize(signature []byte, signedAmount *big.Int) (*types.Transaction, error) {
	return _PromptVault.Contract.ClaimPrize(&_PromptVault.TransactOpts, signature, signedAmount)
}

// ClaimPrize is a paid mutator transaction binding the contract method 0x0dcba9c5.
//
// Solidity: function claimPrize(bytes signature, uint256 signedAmount) returns()
func (_PromptVault *PromptVaultTransactorSession) ClaimPrize(signature []byte, signedAmount *big.Int) (*types.Transaction, error) {
	return _PromptVault.Contract.ClaimPrize(&_PromptVault.TransactOpts, signature, signedAmount)
}

// Play is a paid mutator transaction binding the contract method 0x93e84cd9.
//
// Solidity: function play() payable returns()
func (_PromptVault *PromptVaultTransactor) Play(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PromptVault.contract.Transact(opts, "play")
}

// Play is a paid mutator transaction binding the contract method 0x93e84cd9.
//
// Solidity: function play() payable returns()
func (_PromptVault *PromptVaultSession) Play() (*types.Transaction, error) {
	return _PromptVault.Contract.Play(&_PromptVault.TransactOpts)
}

// Play is a paid mutator transaction binding the contract method 0x93e84cd9.
//
// Solidity: function play() payable returns()
func (_PromptVault *PromptVaultTransactorSession) Play() (*types.Transaction, error) {
	return _PromptVault.Contract.Play(&_PromptVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PromptVault *PromptVaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PromptVault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PromptVault *PromptVaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _PromptVault.Contract.RenounceOwnership(&_PromptVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PromptVault *PromptVaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PromptVault.Contract.RenounceOwnership(&_PromptVault.TransactOpts)
}

// SeedPrizePool is a paid mutator transaction binding the contract method 0x1f9d7e2e.
//
// Solidity: function seedPrizePool() payable returns()
func (_PromptVault *PromptVaultTransactor) SeedPrizePool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PromptVault.contract.Transact(opts, "seedPrizePool")
}

// SeedPrizePool is a paid mutator transaction binding the contract method 0x1f9d7e2e.
//
// Solidity: function seedPrizePool() payable returns()
func (_PromptVault *PromptVaultSession) SeedPrizePool() (*types.Transaction, error) {
	return _PromptVault.Contract.SeedPrizePool(&_PromptVault.TransactOpts)
}

// SeedPrizePool is a paid mutator transaction binding the contract method 0x1f9d7e2e.
//
// Solidity: function seedPrizePool() payable returns()
func (_PromptVault *PromptVaultTransactorSession) SeedPrizePool() (*types.Transaction, error) {
	return _PromptVault.Contract.SeedPrizePool(&_PromptVault.TransactOpts)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address _newSigner) returns()
func (_PromptVault *PromptVaultTransactor) SetSigner(opts *bind.TransactOpts, _newSigner common.Address) (*types.Transaction, error) {
	return _PromptVault.contract.Transact(opts, "setSigner", _newSigner)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address _newSigner) returns()
func (_PromptVault *PromptVaultSession) SetSigner(_newSigner common.Address) (*types.Transaction, error) {
	return _PromptVault.Contract.SetSigner(&_PromptVault.TransactOpts, _newSigner)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address _newSigner) returns()
func (_PromptVault *PromptVaultTransactorSession) SetSigner(_newSigner common.Address) (*types.Transaction, error) {
	return _PromptVault.Contract.SetSigner(&_PromptVault.TransactOpts, _newSigner)
}

// SetTicketPrice is a paid mutator transaction binding the contract method 0x15981650.
//
// Solidity: function setTicketPrice(uint256 _newPrice) returns()
func (_PromptVault *PromptVaultTransactor) SetTicketPrice(opts *bind.TransactOpts, _newPrice *big.Int) (*types.Transaction, error) {
	return _PromptVault.contract.Transact(opts, "setTicketPrice", _newPrice)
}

// SetTicketPrice is a paid mutator transaction binding the contract method 0x15981650.
//
// Solidity: function setTicketPrice(uint256 _newPrice) returns()
func (_PromptVault *PromptVaultSession) SetTicketPrice(_newPrice *big.Int) (*types.Transaction, error) {
	return _PromptVault.Contract.SetTicketPrice(&_PromptVault.TransactOpts, _newPrice)
}

// SetTicketPrice is a paid mutator transaction binding the contract method 0x15981650.
//
// Solidity: function setTicketPrice(uint256 _newPrice) returns()
func (_PromptVault *PromptVaultTransactorSession) SetTicketPrice(_newPrice *big.Int) (*types.Transaction, error) {
	return _PromptVault.Contract.SetTicketPrice(&_PromptVault.TransactOpts, _newPrice)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PromptVault *PromptVaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PromptVault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PromptVault *PromptVaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PromptVault.Contract.TransferOwnership(&_PromptVault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PromptVault *PromptVaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PromptVault.Contract.TransferOwnership(&_PromptVault.TransactOpts, newOwner)
}

// WithdrawRevenue is a paid mutator transaction binding the contract method 0x4f573cb2.
//
// Solidity: function withdrawRevenue() returns()
func (_PromptVault *PromptVaultTransactor) WithdrawRevenue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PromptVault.contract.Transact(opts, "withdrawRevenue")
}

// WithdrawRevenue is a paid mutator transaction binding the contract method 0x4f573cb2.
//
// Solidity: function withdrawRevenue() returns()
func (_PromptVault *PromptVaultSession) WithdrawRevenue() (*types.Transaction, error) {
	return _PromptVault.Contract.WithdrawRevenue(&_PromptVault.TransactOpts)
}

// WithdrawRevenue is a paid mutator transaction binding the contract method 0x4f573cb2.
//
// Solidity: function withdrawRevenue() returns()
func (_PromptVault *PromptVaultTransactorSession) WithdrawRevenue() (*types.Transaction, error) {
	return _PromptVault.Contract.WithdrawRevenue(&_PromptVault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PromptVault *PromptVaultTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PromptVault.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PromptVault *PromptVaultSession) Receive() (*types.Transaction, error) {
	return _PromptVault.Contract.Receive(&_PromptVault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PromptVault *PromptVaultTransactorSession) Receive() (*types.Transaction, error) {
	return _PromptVault.Contract.Receive(&_PromptVault.TransactOpts)
}

// PromptVaultFundsReceivedIterator is returned from FilterFundsReceived and is used to iterate over the raw logs and unpacked data for FundsReceived events raised by the PromptVault contract.
type PromptVaultFundsReceivedIterator struct {
	Event *PromptVaultFundsReceived // Event containing the contract specifics and raw log

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
func (it *PromptVaultFundsReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PromptVaultFundsReceived)
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
		it.Event = new(PromptVaultFundsReceived)
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
func (it *PromptVaultFundsReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PromptVaultFundsReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PromptVaultFundsReceived represents a FundsReceived event raised by the PromptVault contract.
type PromptVaultFundsReceived struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFundsReceived is a free log retrieval operation binding the contract event 0x8e47b87b0ef542cdfa1659c551d88bad38aa7f452d2bbb349ab7530dfec8be8f.
//
// Solidity: event FundsReceived(address indexed sender, uint256 amount)
func (_PromptVault *PromptVaultFilterer) FilterFundsReceived(opts *bind.FilterOpts, sender []common.Address) (*PromptVaultFundsReceivedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PromptVault.contract.FilterLogs(opts, "FundsReceived", senderRule)
	if err != nil {
		return nil, err
	}
	return &PromptVaultFundsReceivedIterator{contract: _PromptVault.contract, event: "FundsReceived", logs: logs, sub: sub}, nil
}

// WatchFundsReceived is a free log subscription operation binding the contract event 0x8e47b87b0ef542cdfa1659c551d88bad38aa7f452d2bbb349ab7530dfec8be8f.
//
// Solidity: event FundsReceived(address indexed sender, uint256 amount)
func (_PromptVault *PromptVaultFilterer) WatchFundsReceived(opts *bind.WatchOpts, sink chan<- *PromptVaultFundsReceived, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PromptVault.contract.WatchLogs(opts, "FundsReceived", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PromptVaultFundsReceived)
				if err := _PromptVault.contract.UnpackLog(event, "FundsReceived", log); err != nil {
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

// ParseFundsReceived is a log parse operation binding the contract event 0x8e47b87b0ef542cdfa1659c551d88bad38aa7f452d2bbb349ab7530dfec8be8f.
//
// Solidity: event FundsReceived(address indexed sender, uint256 amount)
func (_PromptVault *PromptVaultFilterer) ParseFundsReceived(log types.Log) (*PromptVaultFundsReceived, error) {
	event := new(PromptVaultFundsReceived)
	if err := _PromptVault.contract.UnpackLog(event, "FundsReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PromptVaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PromptVault contract.
type PromptVaultOwnershipTransferredIterator struct {
	Event *PromptVaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PromptVaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PromptVaultOwnershipTransferred)
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
		it.Event = new(PromptVaultOwnershipTransferred)
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
func (it *PromptVaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PromptVaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PromptVaultOwnershipTransferred represents a OwnershipTransferred event raised by the PromptVault contract.
type PromptVaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PromptVault *PromptVaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PromptVaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PromptVault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PromptVaultOwnershipTransferredIterator{contract: _PromptVault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PromptVault *PromptVaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PromptVaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PromptVault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PromptVaultOwnershipTransferred)
				if err := _PromptVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PromptVault *PromptVaultFilterer) ParseOwnershipTransferred(log types.Log) (*PromptVaultOwnershipTransferred, error) {
	event := new(PromptVaultOwnershipTransferred)
	if err := _PromptVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PromptVaultPrizeClaimedIterator is returned from FilterPrizeClaimed and is used to iterate over the raw logs and unpacked data for PrizeClaimed events raised by the PromptVault contract.
type PromptVaultPrizeClaimedIterator struct {
	Event *PromptVaultPrizeClaimed // Event containing the contract specifics and raw log

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
func (it *PromptVaultPrizeClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PromptVaultPrizeClaimed)
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
		it.Event = new(PromptVaultPrizeClaimed)
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
func (it *PromptVaultPrizeClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PromptVaultPrizeClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PromptVaultPrizeClaimed represents a PrizeClaimed event raised by the PromptVault contract.
type PromptVaultPrizeClaimed struct {
	Winner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPrizeClaimed is a free log retrieval operation binding the contract event 0x95681e512bc0fe659e195e06c283eada494316f3d801213e48e7101af92bf770.
//
// Solidity: event PrizeClaimed(address indexed winner, uint256 amount)
func (_PromptVault *PromptVaultFilterer) FilterPrizeClaimed(opts *bind.FilterOpts, winner []common.Address) (*PromptVaultPrizeClaimedIterator, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _PromptVault.contract.FilterLogs(opts, "PrizeClaimed", winnerRule)
	if err != nil {
		return nil, err
	}
	return &PromptVaultPrizeClaimedIterator{contract: _PromptVault.contract, event: "PrizeClaimed", logs: logs, sub: sub}, nil
}

// WatchPrizeClaimed is a free log subscription operation binding the contract event 0x95681e512bc0fe659e195e06c283eada494316f3d801213e48e7101af92bf770.
//
// Solidity: event PrizeClaimed(address indexed winner, uint256 amount)
func (_PromptVault *PromptVaultFilterer) WatchPrizeClaimed(opts *bind.WatchOpts, sink chan<- *PromptVaultPrizeClaimed, winner []common.Address) (event.Subscription, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _PromptVault.contract.WatchLogs(opts, "PrizeClaimed", winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PromptVaultPrizeClaimed)
				if err := _PromptVault.contract.UnpackLog(event, "PrizeClaimed", log); err != nil {
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

// ParsePrizeClaimed is a log parse operation binding the contract event 0x95681e512bc0fe659e195e06c283eada494316f3d801213e48e7101af92bf770.
//
// Solidity: event PrizeClaimed(address indexed winner, uint256 amount)
func (_PromptVault *PromptVaultFilterer) ParsePrizeClaimed(log types.Log) (*PromptVaultPrizeClaimed, error) {
	event := new(PromptVaultPrizeClaimed)
	if err := _PromptVault.contract.UnpackLog(event, "PrizeClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PromptVaultRevenueWithdrawnIterator is returned from FilterRevenueWithdrawn and is used to iterate over the raw logs and unpacked data for RevenueWithdrawn events raised by the PromptVault contract.
type PromptVaultRevenueWithdrawnIterator struct {
	Event *PromptVaultRevenueWithdrawn // Event containing the contract specifics and raw log

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
func (it *PromptVaultRevenueWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PromptVaultRevenueWithdrawn)
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
		it.Event = new(PromptVaultRevenueWithdrawn)
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
func (it *PromptVaultRevenueWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PromptVaultRevenueWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PromptVaultRevenueWithdrawn represents a RevenueWithdrawn event raised by the PromptVault contract.
type PromptVaultRevenueWithdrawn struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRevenueWithdrawn is a free log retrieval operation binding the contract event 0x86a5cc8fe9af9ae68fb50f62885307a7755a30cbd290131644377f0bd94a7181.
//
// Solidity: event RevenueWithdrawn(address indexed to, uint256 amount)
func (_PromptVault *PromptVaultFilterer) FilterRevenueWithdrawn(opts *bind.FilterOpts, to []common.Address) (*PromptVaultRevenueWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PromptVault.contract.FilterLogs(opts, "RevenueWithdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return &PromptVaultRevenueWithdrawnIterator{contract: _PromptVault.contract, event: "RevenueWithdrawn", logs: logs, sub: sub}, nil
}

// WatchRevenueWithdrawn is a free log subscription operation binding the contract event 0x86a5cc8fe9af9ae68fb50f62885307a7755a30cbd290131644377f0bd94a7181.
//
// Solidity: event RevenueWithdrawn(address indexed to, uint256 amount)
func (_PromptVault *PromptVaultFilterer) WatchRevenueWithdrawn(opts *bind.WatchOpts, sink chan<- *PromptVaultRevenueWithdrawn, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PromptVault.contract.WatchLogs(opts, "RevenueWithdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PromptVaultRevenueWithdrawn)
				if err := _PromptVault.contract.UnpackLog(event, "RevenueWithdrawn", log); err != nil {
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

// ParseRevenueWithdrawn is a log parse operation binding the contract event 0x86a5cc8fe9af9ae68fb50f62885307a7755a30cbd290131644377f0bd94a7181.
//
// Solidity: event RevenueWithdrawn(address indexed to, uint256 amount)
func (_PromptVault *PromptVaultFilterer) ParseRevenueWithdrawn(log types.Log) (*PromptVaultRevenueWithdrawn, error) {
	event := new(PromptVaultRevenueWithdrawn)
	if err := _PromptVault.contract.UnpackLog(event, "RevenueWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PromptVaultSignerUpdatedIterator is returned from FilterSignerUpdated and is used to iterate over the raw logs and unpacked data for SignerUpdated events raised by the PromptVault contract.
type PromptVaultSignerUpdatedIterator struct {
	Event *PromptVaultSignerUpdated // Event containing the contract specifics and raw log

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
func (it *PromptVaultSignerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PromptVaultSignerUpdated)
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
		it.Event = new(PromptVaultSignerUpdated)
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
func (it *PromptVaultSignerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PromptVaultSignerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PromptVaultSignerUpdated represents a SignerUpdated event raised by the PromptVault contract.
type PromptVaultSignerUpdated struct {
	OldSigner common.Address
	NewSigner common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSignerUpdated is a free log retrieval operation binding the contract event 0x2d025324f0a785e8c12d0a0d91a9caa49df4ef20ff87e0df7213a1d4f3157beb.
//
// Solidity: event SignerUpdated(address indexed oldSigner, address indexed newSigner)
func (_PromptVault *PromptVaultFilterer) FilterSignerUpdated(opts *bind.FilterOpts, oldSigner []common.Address, newSigner []common.Address) (*PromptVaultSignerUpdatedIterator, error) {

	var oldSignerRule []interface{}
	for _, oldSignerItem := range oldSigner {
		oldSignerRule = append(oldSignerRule, oldSignerItem)
	}
	var newSignerRule []interface{}
	for _, newSignerItem := range newSigner {
		newSignerRule = append(newSignerRule, newSignerItem)
	}

	logs, sub, err := _PromptVault.contract.FilterLogs(opts, "SignerUpdated", oldSignerRule, newSignerRule)
	if err != nil {
		return nil, err
	}
	return &PromptVaultSignerUpdatedIterator{contract: _PromptVault.contract, event: "SignerUpdated", logs: logs, sub: sub}, nil
}

// WatchSignerUpdated is a free log subscription operation binding the contract event 0x2d025324f0a785e8c12d0a0d91a9caa49df4ef20ff87e0df7213a1d4f3157beb.
//
// Solidity: event SignerUpdated(address indexed oldSigner, address indexed newSigner)
func (_PromptVault *PromptVaultFilterer) WatchSignerUpdated(opts *bind.WatchOpts, sink chan<- *PromptVaultSignerUpdated, oldSigner []common.Address, newSigner []common.Address) (event.Subscription, error) {

	var oldSignerRule []interface{}
	for _, oldSignerItem := range oldSigner {
		oldSignerRule = append(oldSignerRule, oldSignerItem)
	}
	var newSignerRule []interface{}
	for _, newSignerItem := range newSigner {
		newSignerRule = append(newSignerRule, newSignerItem)
	}

	logs, sub, err := _PromptVault.contract.WatchLogs(opts, "SignerUpdated", oldSignerRule, newSignerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PromptVaultSignerUpdated)
				if err := _PromptVault.contract.UnpackLog(event, "SignerUpdated", log); err != nil {
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

// ParseSignerUpdated is a log parse operation binding the contract event 0x2d025324f0a785e8c12d0a0d91a9caa49df4ef20ff87e0df7213a1d4f3157beb.
//
// Solidity: event SignerUpdated(address indexed oldSigner, address indexed newSigner)
func (_PromptVault *PromptVaultFilterer) ParseSignerUpdated(log types.Log) (*PromptVaultSignerUpdated, error) {
	event := new(PromptVaultSignerUpdated)
	if err := _PromptVault.contract.UnpackLog(event, "SignerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PromptVaultTicketPriceUpdatedIterator is returned from FilterTicketPriceUpdated and is used to iterate over the raw logs and unpacked data for TicketPriceUpdated events raised by the PromptVault contract.
type PromptVaultTicketPriceUpdatedIterator struct {
	Event *PromptVaultTicketPriceUpdated // Event containing the contract specifics and raw log

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
func (it *PromptVaultTicketPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PromptVaultTicketPriceUpdated)
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
		it.Event = new(PromptVaultTicketPriceUpdated)
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
func (it *PromptVaultTicketPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PromptVaultTicketPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PromptVaultTicketPriceUpdated represents a TicketPriceUpdated event raised by the PromptVault contract.
type PromptVaultTicketPriceUpdated struct {
	OldPrice *big.Int
	NewPrice *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTicketPriceUpdated is a free log retrieval operation binding the contract event 0xd4c5e06b1ae097ba02372652a7adaa6e4a8e00be527497a3ad0ebc3f761ef3fb.
//
// Solidity: event TicketPriceUpdated(uint256 oldPrice, uint256 newPrice)
func (_PromptVault *PromptVaultFilterer) FilterTicketPriceUpdated(opts *bind.FilterOpts) (*PromptVaultTicketPriceUpdatedIterator, error) {

	logs, sub, err := _PromptVault.contract.FilterLogs(opts, "TicketPriceUpdated")
	if err != nil {
		return nil, err
	}
	return &PromptVaultTicketPriceUpdatedIterator{contract: _PromptVault.contract, event: "TicketPriceUpdated", logs: logs, sub: sub}, nil
}

// WatchTicketPriceUpdated is a free log subscription operation binding the contract event 0xd4c5e06b1ae097ba02372652a7adaa6e4a8e00be527497a3ad0ebc3f761ef3fb.
//
// Solidity: event TicketPriceUpdated(uint256 oldPrice, uint256 newPrice)
func (_PromptVault *PromptVaultFilterer) WatchTicketPriceUpdated(opts *bind.WatchOpts, sink chan<- *PromptVaultTicketPriceUpdated) (event.Subscription, error) {

	logs, sub, err := _PromptVault.contract.WatchLogs(opts, "TicketPriceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PromptVaultTicketPriceUpdated)
				if err := _PromptVault.contract.UnpackLog(event, "TicketPriceUpdated", log); err != nil {
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

// ParseTicketPriceUpdated is a log parse operation binding the contract event 0xd4c5e06b1ae097ba02372652a7adaa6e4a8e00be527497a3ad0ebc3f761ef3fb.
//
// Solidity: event TicketPriceUpdated(uint256 oldPrice, uint256 newPrice)
func (_PromptVault *PromptVaultFilterer) ParseTicketPriceUpdated(log types.Log) (*PromptVaultTicketPriceUpdated, error) {
	event := new(PromptVaultTicketPriceUpdated)
	if err := _PromptVault.contract.UnpackLog(event, "TicketPriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PromptVaultTicketPurchasedIterator is returned from FilterTicketPurchased and is used to iterate over the raw logs and unpacked data for TicketPurchased events raised by the PromptVault contract.
type PromptVaultTicketPurchasedIterator struct {
	Event *PromptVaultTicketPurchased // Event containing the contract specifics and raw log

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
func (it *PromptVaultTicketPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PromptVaultTicketPurchased)
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
		it.Event = new(PromptVaultTicketPurchased)
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
func (it *PromptVaultTicketPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PromptVaultTicketPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PromptVaultTicketPurchased represents a TicketPurchased event raised by the PromptVault contract.
type PromptVaultTicketPurchased struct {
	Player       common.Address
	Amount       *big.Int
	PrizePoolNew *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTicketPurchased is a free log retrieval operation binding the contract event 0x2a91574e12ad96234e84923e146b0946ecfb871cd8d5534dc1fdcbe87a7c01b3.
//
// Solidity: event TicketPurchased(address indexed player, uint256 amount, uint256 prizePoolNew)
func (_PromptVault *PromptVaultFilterer) FilterTicketPurchased(opts *bind.FilterOpts, player []common.Address) (*PromptVaultTicketPurchasedIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _PromptVault.contract.FilterLogs(opts, "TicketPurchased", playerRule)
	if err != nil {
		return nil, err
	}
	return &PromptVaultTicketPurchasedIterator{contract: _PromptVault.contract, event: "TicketPurchased", logs: logs, sub: sub}, nil
}

// WatchTicketPurchased is a free log subscription operation binding the contract event 0x2a91574e12ad96234e84923e146b0946ecfb871cd8d5534dc1fdcbe87a7c01b3.
//
// Solidity: event TicketPurchased(address indexed player, uint256 amount, uint256 prizePoolNew)
func (_PromptVault *PromptVaultFilterer) WatchTicketPurchased(opts *bind.WatchOpts, sink chan<- *PromptVaultTicketPurchased, player []common.Address) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _PromptVault.contract.WatchLogs(opts, "TicketPurchased", playerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PromptVaultTicketPurchased)
				if err := _PromptVault.contract.UnpackLog(event, "TicketPurchased", log); err != nil {
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

// ParseTicketPurchased is a log parse operation binding the contract event 0x2a91574e12ad96234e84923e146b0946ecfb871cd8d5534dc1fdcbe87a7c01b3.
//
// Solidity: event TicketPurchased(address indexed player, uint256 amount, uint256 prizePoolNew)
func (_PromptVault *PromptVaultFilterer) ParseTicketPurchased(log types.Log) (*PromptVaultTicketPurchased, error) {
	event := new(PromptVaultTicketPurchased)
	if err := _PromptVault.contract.UnpackLog(event, "TicketPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

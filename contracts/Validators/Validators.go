// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Validators

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

// ValidatorsABI is the input ABI used to generate the binding from.
const ValidatorsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogAddToTopValidators\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogCreateValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogDistributeBlockReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogEditValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogReactive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogRemoveFromTopValidators\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hb\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogRemoveValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hb\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogRemoveValidatorIncoming\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"staking\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogUnstake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"newSet\",\"type\":\"address[]\"}],\"name\":\"LogUpdateValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hb\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogWithdrawProfits\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogWithdrawStaking\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MaxValidators\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MinimalStakingCoin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ProposalAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PunishContractAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"StakingLockPeriod\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ValidatorContractAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WithdrawProfitPeriod\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"currentValidatorSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestValidatorsSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalJailedHB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"vals\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"stake\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"feeAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"moniker\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identity\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"name\":\"createOrEditValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"tryReactive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"unstake\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"withdrawStaking\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"withdrawProfits\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributeBlockReward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newSet\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"updateActiveValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"removeValidatorIncoming\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"getValidatorDescription\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"getValidatorInfo\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"enumValidators.Status\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"getStakingInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalStakeOfActiveValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"len\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"isActiveValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"isTopValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTopValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"moniker\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identity\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"name\":\"validateDescription\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// Validators is an auto generated Go binding around an Ethereum contract.
type Validators struct {
	ValidatorsCaller     // Read-only binding to the contract
	ValidatorsTransactor // Write-only binding to the contract
	ValidatorsFilterer   // Log filterer for contract events
}

// ValidatorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorsSession struct {
	Contract     *Validators       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorsCallerSession struct {
	Contract *ValidatorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ValidatorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorsTransactorSession struct {
	Contract     *ValidatorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ValidatorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorsRaw struct {
	Contract *Validators // Generic contract binding to access the raw methods on
}

// ValidatorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorsCallerRaw struct {
	Contract *ValidatorsCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorsTransactorRaw struct {
	Contract *ValidatorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidators creates a new instance of Validators, bound to a specific deployed contract.
func NewValidators(address common.Address, backend bind.ContractBackend) (*Validators, error) {
	contract, err := bindValidators(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validators{ValidatorsCaller: ValidatorsCaller{contract: contract}, ValidatorsTransactor: ValidatorsTransactor{contract: contract}, ValidatorsFilterer: ValidatorsFilterer{contract: contract}}, nil
}

// NewValidatorsCaller creates a new read-only instance of Validators, bound to a specific deployed contract.
func NewValidatorsCaller(address common.Address, caller bind.ContractCaller) (*ValidatorsCaller, error) {
	contract, err := bindValidators(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorsCaller{contract: contract}, nil
}

// NewValidatorsTransactor creates a new write-only instance of Validators, bound to a specific deployed contract.
func NewValidatorsTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorsTransactor, error) {
	contract, err := bindValidators(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorsTransactor{contract: contract}, nil
}

// NewValidatorsFilterer creates a new log filterer instance of Validators, bound to a specific deployed contract.
func NewValidatorsFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorsFilterer, error) {
	contract, err := bindValidators(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorsFilterer{contract: contract}, nil
}

// bindValidators binds a generic wrapper to an already deployed contract.
func bindValidators(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validators *ValidatorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validators.Contract.ValidatorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validators *ValidatorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validators.Contract.ValidatorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validators *ValidatorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validators.Contract.ValidatorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validators *ValidatorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validators.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validators *ValidatorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validators.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validators *ValidatorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validators.Contract.contract.Transact(opts, method, params...)
}

// MaxValidators is a free data retrieval call binding the contract method 0xc967f90f.
//
// Solidity: function MaxValidators() view returns(uint16)
func (_Validators *ValidatorsCaller) MaxValidators(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "MaxValidators")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MaxValidators is a free data retrieval call binding the contract method 0xc967f90f.
//
// Solidity: function MaxValidators() view returns(uint16)
func (_Validators *ValidatorsSession) MaxValidators() (uint16, error) {
	return _Validators.Contract.MaxValidators(&_Validators.CallOpts)
}

// MaxValidators is a free data retrieval call binding the contract method 0xc967f90f.
//
// Solidity: function MaxValidators() view returns(uint16)
func (_Validators *ValidatorsCallerSession) MaxValidators() (uint16, error) {
	return _Validators.Contract.MaxValidators(&_Validators.CallOpts)
}

// MinimalStakingCoin is a free data retrieval call binding the contract method 0xbe645692.
//
// Solidity: function MinimalStakingCoin() view returns(uint256)
func (_Validators *ValidatorsCaller) MinimalStakingCoin(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "MinimalStakingCoin")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimalStakingCoin is a free data retrieval call binding the contract method 0xbe645692.
//
// Solidity: function MinimalStakingCoin() view returns(uint256)
func (_Validators *ValidatorsSession) MinimalStakingCoin() (*big.Int, error) {
	return _Validators.Contract.MinimalStakingCoin(&_Validators.CallOpts)
}

// MinimalStakingCoin is a free data retrieval call binding the contract method 0xbe645692.
//
// Solidity: function MinimalStakingCoin() view returns(uint256)
func (_Validators *ValidatorsCallerSession) MinimalStakingCoin() (*big.Int, error) {
	return _Validators.Contract.MinimalStakingCoin(&_Validators.CallOpts)
}

// ProposalAddr is a free data retrieval call binding the contract method 0x6233be5d.
//
// Solidity: function ProposalAddr() view returns(address)
func (_Validators *ValidatorsCaller) ProposalAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "ProposalAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposalAddr is a free data retrieval call binding the contract method 0x6233be5d.
//
// Solidity: function ProposalAddr() view returns(address)
func (_Validators *ValidatorsSession) ProposalAddr() (common.Address, error) {
	return _Validators.Contract.ProposalAddr(&_Validators.CallOpts)
}

// ProposalAddr is a free data retrieval call binding the contract method 0x6233be5d.
//
// Solidity: function ProposalAddr() view returns(address)
func (_Validators *ValidatorsCallerSession) ProposalAddr() (common.Address, error) {
	return _Validators.Contract.ProposalAddr(&_Validators.CallOpts)
}

// PunishContractAddr is a free data retrieval call binding the contract method 0x1b5e358c.
//
// Solidity: function PunishContractAddr() view returns(address)
func (_Validators *ValidatorsCaller) PunishContractAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "PunishContractAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PunishContractAddr is a free data retrieval call binding the contract method 0x1b5e358c.
//
// Solidity: function PunishContractAddr() view returns(address)
func (_Validators *ValidatorsSession) PunishContractAddr() (common.Address, error) {
	return _Validators.Contract.PunishContractAddr(&_Validators.CallOpts)
}

// PunishContractAddr is a free data retrieval call binding the contract method 0x1b5e358c.
//
// Solidity: function PunishContractAddr() view returns(address)
func (_Validators *ValidatorsCallerSession) PunishContractAddr() (common.Address, error) {
	return _Validators.Contract.PunishContractAddr(&_Validators.CallOpts)
}

// StakingLockPeriod is a free data retrieval call binding the contract method 0xdb78dd28.
//
// Solidity: function StakingLockPeriod() view returns(uint64)
func (_Validators *ValidatorsCaller) StakingLockPeriod(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "StakingLockPeriod")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// StakingLockPeriod is a free data retrieval call binding the contract method 0xdb78dd28.
//
// Solidity: function StakingLockPeriod() view returns(uint64)
func (_Validators *ValidatorsSession) StakingLockPeriod() (uint64, error) {
	return _Validators.Contract.StakingLockPeriod(&_Validators.CallOpts)
}

// StakingLockPeriod is a free data retrieval call binding the contract method 0xdb78dd28.
//
// Solidity: function StakingLockPeriod() view returns(uint64)
func (_Validators *ValidatorsCallerSession) StakingLockPeriod() (uint64, error) {
	return _Validators.Contract.StakingLockPeriod(&_Validators.CallOpts)
}

// ValidatorContractAddr is a free data retrieval call binding the contract method 0x3a061bd3.
//
// Solidity: function ValidatorContractAddr() view returns(address)
func (_Validators *ValidatorsCaller) ValidatorContractAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "ValidatorContractAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorContractAddr is a free data retrieval call binding the contract method 0x3a061bd3.
//
// Solidity: function ValidatorContractAddr() view returns(address)
func (_Validators *ValidatorsSession) ValidatorContractAddr() (common.Address, error) {
	return _Validators.Contract.ValidatorContractAddr(&_Validators.CallOpts)
}

// ValidatorContractAddr is a free data retrieval call binding the contract method 0x3a061bd3.
//
// Solidity: function ValidatorContractAddr() view returns(address)
func (_Validators *ValidatorsCallerSession) ValidatorContractAddr() (common.Address, error) {
	return _Validators.Contract.ValidatorContractAddr(&_Validators.CallOpts)
}

// WithdrawProfitPeriod is a free data retrieval call binding the contract method 0xefd8d8e2.
//
// Solidity: function WithdrawProfitPeriod() view returns(uint64)
func (_Validators *ValidatorsCaller) WithdrawProfitPeriod(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "WithdrawProfitPeriod")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// WithdrawProfitPeriod is a free data retrieval call binding the contract method 0xefd8d8e2.
//
// Solidity: function WithdrawProfitPeriod() view returns(uint64)
func (_Validators *ValidatorsSession) WithdrawProfitPeriod() (uint64, error) {
	return _Validators.Contract.WithdrawProfitPeriod(&_Validators.CallOpts)
}

// WithdrawProfitPeriod is a free data retrieval call binding the contract method 0xefd8d8e2.
//
// Solidity: function WithdrawProfitPeriod() view returns(uint64)
func (_Validators *ValidatorsCallerSession) WithdrawProfitPeriod() (uint64, error) {
	return _Validators.Contract.WithdrawProfitPeriod(&_Validators.CallOpts)
}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address)
func (_Validators *ValidatorsCaller) CurrentValidatorSet(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "currentValidatorSet", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address)
func (_Validators *ValidatorsSession) CurrentValidatorSet(arg0 *big.Int) (common.Address, error) {
	return _Validators.Contract.CurrentValidatorSet(&_Validators.CallOpts, arg0)
}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address)
func (_Validators *ValidatorsCallerSession) CurrentValidatorSet(arg0 *big.Int) (common.Address, error) {
	return _Validators.Contract.CurrentValidatorSet(&_Validators.CallOpts, arg0)
}

// GetActiveValidators is a free data retrieval call binding the contract method 0x9de70258.
//
// Solidity: function getActiveValidators() view returns(address[])
func (_Validators *ValidatorsCaller) GetActiveValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "getActiveValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetActiveValidators is a free data retrieval call binding the contract method 0x9de70258.
//
// Solidity: function getActiveValidators() view returns(address[])
func (_Validators *ValidatorsSession) GetActiveValidators() ([]common.Address, error) {
	return _Validators.Contract.GetActiveValidators(&_Validators.CallOpts)
}

// GetActiveValidators is a free data retrieval call binding the contract method 0x9de70258.
//
// Solidity: function getActiveValidators() view returns(address[])
func (_Validators *ValidatorsCallerSession) GetActiveValidators() ([]common.Address, error) {
	return _Validators.Contract.GetActiveValidators(&_Validators.CallOpts)
}

// GetStakingInfo is a free data retrieval call binding the contract method 0x7f4f95fa.
//
// Solidity: function getStakingInfo(address staker, address val) view returns(uint256, uint256, uint256)
func (_Validators *ValidatorsCaller) GetStakingInfo(opts *bind.CallOpts, staker common.Address, val common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "getStakingInfo", staker, val)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetStakingInfo is a free data retrieval call binding the contract method 0x7f4f95fa.
//
// Solidity: function getStakingInfo(address staker, address val) view returns(uint256, uint256, uint256)
func (_Validators *ValidatorsSession) GetStakingInfo(staker common.Address, val common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Validators.Contract.GetStakingInfo(&_Validators.CallOpts, staker, val)
}

// GetStakingInfo is a free data retrieval call binding the contract method 0x7f4f95fa.
//
// Solidity: function getStakingInfo(address staker, address val) view returns(uint256, uint256, uint256)
func (_Validators *ValidatorsCallerSession) GetStakingInfo(staker common.Address, val common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Validators.Contract.GetStakingInfo(&_Validators.CallOpts, staker, val)
}

// GetTopValidators is a free data retrieval call binding the contract method 0xafeea115.
//
// Solidity: function getTopValidators() view returns(address[])
func (_Validators *ValidatorsCaller) GetTopValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "getTopValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTopValidators is a free data retrieval call binding the contract method 0xafeea115.
//
// Solidity: function getTopValidators() view returns(address[])
func (_Validators *ValidatorsSession) GetTopValidators() ([]common.Address, error) {
	return _Validators.Contract.GetTopValidators(&_Validators.CallOpts)
}

// GetTopValidators is a free data retrieval call binding the contract method 0xafeea115.
//
// Solidity: function getTopValidators() view returns(address[])
func (_Validators *ValidatorsCallerSession) GetTopValidators() ([]common.Address, error) {
	return _Validators.Contract.GetTopValidators(&_Validators.CallOpts)
}

// GetTotalStakeOfActiveValidators is a free data retrieval call binding the contract method 0xc253c384.
//
// Solidity: function getTotalStakeOfActiveValidators() view returns(uint256 total, uint256 len)
func (_Validators *ValidatorsCaller) GetTotalStakeOfActiveValidators(opts *bind.CallOpts) (struct {
	Total *big.Int
	Len   *big.Int
}, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "getTotalStakeOfActiveValidators")

	outstruct := new(struct {
		Total *big.Int
		Len   *big.Int
	})

	outstruct.Total = out[0].(*big.Int)
	outstruct.Len = out[1].(*big.Int)

	return *outstruct, err

}

// GetTotalStakeOfActiveValidators is a free data retrieval call binding the contract method 0xc253c384.
//
// Solidity: function getTotalStakeOfActiveValidators() view returns(uint256 total, uint256 len)
func (_Validators *ValidatorsSession) GetTotalStakeOfActiveValidators() (struct {
	Total *big.Int
	Len   *big.Int
}, error) {
	return _Validators.Contract.GetTotalStakeOfActiveValidators(&_Validators.CallOpts)
}

// GetTotalStakeOfActiveValidators is a free data retrieval call binding the contract method 0xc253c384.
//
// Solidity: function getTotalStakeOfActiveValidators() view returns(uint256 total, uint256 len)
func (_Validators *ValidatorsCallerSession) GetTotalStakeOfActiveValidators() (struct {
	Total *big.Int
	Len   *big.Int
}, error) {
	return _Validators.Contract.GetTotalStakeOfActiveValidators(&_Validators.CallOpts)
}

// GetValidatorDescription is a free data retrieval call binding the contract method 0xa43569b3.
//
// Solidity: function getValidatorDescription(address val) view returns(string, string, string, string, string)
func (_Validators *ValidatorsCaller) GetValidatorDescription(opts *bind.CallOpts, val common.Address) (string, string, string, string, string, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "getValidatorDescription", val)

	if err != nil {
		return *new(string), *new(string), *new(string), *new(string), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(string)).(*string)
	out4 := *abi.ConvertType(out[4], new(string)).(*string)

	return out0, out1, out2, out3, out4, err

}

// GetValidatorDescription is a free data retrieval call binding the contract method 0xa43569b3.
//
// Solidity: function getValidatorDescription(address val) view returns(string, string, string, string, string)
func (_Validators *ValidatorsSession) GetValidatorDescription(val common.Address) (string, string, string, string, string, error) {
	return _Validators.Contract.GetValidatorDescription(&_Validators.CallOpts, val)
}

// GetValidatorDescription is a free data retrieval call binding the contract method 0xa43569b3.
//
// Solidity: function getValidatorDescription(address val) view returns(string, string, string, string, string)
func (_Validators *ValidatorsCallerSession) GetValidatorDescription(val common.Address) (string, string, string, string, string, error) {
	return _Validators.Contract.GetValidatorDescription(&_Validators.CallOpts, val)
}

// GetValidatorInfo is a free data retrieval call binding the contract method 0x8a11d7c9.
//
// Solidity: function getValidatorInfo(address val) view returns(address, uint8, uint256, uint256, uint256, uint256, address[])
func (_Validators *ValidatorsCaller) GetValidatorInfo(opts *bind.CallOpts, val common.Address) (common.Address, uint8, *big.Int, *big.Int, *big.Int, *big.Int, []common.Address, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "getValidatorInfo", val)

	if err != nil {
		return *new(common.Address), *new(uint8), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new([]common.Address)).(*[]common.Address)

	return out0, out1, out2, out3, out4, out5, out6, err

}

// GetValidatorInfo is a free data retrieval call binding the contract method 0x8a11d7c9.
//
// Solidity: function getValidatorInfo(address val) view returns(address, uint8, uint256, uint256, uint256, uint256, address[])
func (_Validators *ValidatorsSession) GetValidatorInfo(val common.Address) (common.Address, uint8, *big.Int, *big.Int, *big.Int, *big.Int, []common.Address, error) {
	return _Validators.Contract.GetValidatorInfo(&_Validators.CallOpts, val)
}

// GetValidatorInfo is a free data retrieval call binding the contract method 0x8a11d7c9.
//
// Solidity: function getValidatorInfo(address val) view returns(address, uint8, uint256, uint256, uint256, uint256, address[])
func (_Validators *ValidatorsCallerSession) GetValidatorInfo(val common.Address) (common.Address, uint8, *big.Int, *big.Int, *big.Int, *big.Int, []common.Address, error) {
	return _Validators.Contract.GetValidatorInfo(&_Validators.CallOpts, val)
}

// HighestValidatorsSet is a free data retrieval call binding the contract method 0x4b3d500b.
//
// Solidity: function highestValidatorsSet(uint256 ) view returns(address)
func (_Validators *ValidatorsCaller) HighestValidatorsSet(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "highestValidatorsSet", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HighestValidatorsSet is a free data retrieval call binding the contract method 0x4b3d500b.
//
// Solidity: function highestValidatorsSet(uint256 ) view returns(address)
func (_Validators *ValidatorsSession) HighestValidatorsSet(arg0 *big.Int) (common.Address, error) {
	return _Validators.Contract.HighestValidatorsSet(&_Validators.CallOpts, arg0)
}

// HighestValidatorsSet is a free data retrieval call binding the contract method 0x4b3d500b.
//
// Solidity: function highestValidatorsSet(uint256 ) view returns(address)
func (_Validators *ValidatorsCallerSession) HighestValidatorsSet(arg0 *big.Int) (common.Address, error) {
	return _Validators.Contract.HighestValidatorsSet(&_Validators.CallOpts, arg0)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Validators *ValidatorsCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Validators *ValidatorsSession) Initialized() (bool, error) {
	return _Validators.Contract.Initialized(&_Validators.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Validators *ValidatorsCallerSession) Initialized() (bool, error) {
	return _Validators.Contract.Initialized(&_Validators.CallOpts)
}

// IsActiveValidator is a free data retrieval call binding the contract method 0x40550a1c.
//
// Solidity: function isActiveValidator(address who) view returns(bool)
func (_Validators *ValidatorsCaller) IsActiveValidator(opts *bind.CallOpts, who common.Address) (bool, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "isActiveValidator", who)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveValidator is a free data retrieval call binding the contract method 0x40550a1c.
//
// Solidity: function isActiveValidator(address who) view returns(bool)
func (_Validators *ValidatorsSession) IsActiveValidator(who common.Address) (bool, error) {
	return _Validators.Contract.IsActiveValidator(&_Validators.CallOpts, who)
}

// IsActiveValidator is a free data retrieval call binding the contract method 0x40550a1c.
//
// Solidity: function isActiveValidator(address who) view returns(bool)
func (_Validators *ValidatorsCallerSession) IsActiveValidator(who common.Address) (bool, error) {
	return _Validators.Contract.IsActiveValidator(&_Validators.CallOpts, who)
}

// IsTopValidator is a free data retrieval call binding the contract method 0x98e3b626.
//
// Solidity: function isTopValidator(address who) view returns(bool)
func (_Validators *ValidatorsCaller) IsTopValidator(opts *bind.CallOpts, who common.Address) (bool, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "isTopValidator", who)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTopValidator is a free data retrieval call binding the contract method 0x98e3b626.
//
// Solidity: function isTopValidator(address who) view returns(bool)
func (_Validators *ValidatorsSession) IsTopValidator(who common.Address) (bool, error) {
	return _Validators.Contract.IsTopValidator(&_Validators.CallOpts, who)
}

// IsTopValidator is a free data retrieval call binding the contract method 0x98e3b626.
//
// Solidity: function isTopValidator(address who) view returns(bool)
func (_Validators *ValidatorsCallerSession) IsTopValidator(who common.Address) (bool, error) {
	return _Validators.Contract.IsTopValidator(&_Validators.CallOpts, who)
}

// TotalJailedHB is a free data retrieval call binding the contract method 0x1303f7cf.
//
// Solidity: function totalJailedHB() view returns(uint256)
func (_Validators *ValidatorsCaller) TotalJailedHB(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "totalJailedHB")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalJailedHB is a free data retrieval call binding the contract method 0x1303f7cf.
//
// Solidity: function totalJailedHB() view returns(uint256)
func (_Validators *ValidatorsSession) TotalJailedHB() (*big.Int, error) {
	return _Validators.Contract.TotalJailedHB(&_Validators.CallOpts)
}

// TotalJailedHB is a free data retrieval call binding the contract method 0x1303f7cf.
//
// Solidity: function totalJailedHB() view returns(uint256)
func (_Validators *ValidatorsCallerSession) TotalJailedHB() (*big.Int, error) {
	return _Validators.Contract.TotalJailedHB(&_Validators.CallOpts)
}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_Validators *ValidatorsCaller) TotalStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "totalStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_Validators *ValidatorsSession) TotalStake() (*big.Int, error) {
	return _Validators.Contract.TotalStake(&_Validators.CallOpts)
}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_Validators *ValidatorsCallerSession) TotalStake() (*big.Int, error) {
	return _Validators.Contract.TotalStake(&_Validators.CallOpts)
}

// ValidateDescription is a free data retrieval call binding the contract method 0xb6c88519.
//
// Solidity: function validateDescription(string moniker, string identity, string website, string email, string details) pure returns(bool)
func (_Validators *ValidatorsCaller) ValidateDescription(opts *bind.CallOpts, moniker string, identity string, website string, email string, details string) (bool, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "validateDescription", moniker, identity, website, email, details)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateDescription is a free data retrieval call binding the contract method 0xb6c88519.
//
// Solidity: function validateDescription(string moniker, string identity, string website, string email, string details) pure returns(bool)
func (_Validators *ValidatorsSession) ValidateDescription(moniker string, identity string, website string, email string, details string) (bool, error) {
	return _Validators.Contract.ValidateDescription(&_Validators.CallOpts, moniker, identity, website, email, details)
}

// ValidateDescription is a free data retrieval call binding the contract method 0xb6c88519.
//
// Solidity: function validateDescription(string moniker, string identity, string website, string email, string details) pure returns(bool)
func (_Validators *ValidatorsCallerSession) ValidateDescription(moniker string, identity string, website string, email string, details string) (bool, error) {
	return _Validators.Contract.ValidateDescription(&_Validators.CallOpts, moniker, identity, website, email, details)
}

// CreateOrEditValidator is a paid mutator transaction binding the contract method 0xa406fcb7.
//
// Solidity: function createOrEditValidator(address feeAddr, string moniker, string identity, string website, string email, string details) returns(bool)
func (_Validators *ValidatorsTransactor) CreateOrEditValidator(opts *bind.TransactOpts, feeAddr common.Address, moniker string, identity string, website string, email string, details string) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "createOrEditValidator", feeAddr, moniker, identity, website, email, details)
}

// CreateOrEditValidator is a paid mutator transaction binding the contract method 0xa406fcb7.
//
// Solidity: function createOrEditValidator(address feeAddr, string moniker, string identity, string website, string email, string details) returns(bool)
func (_Validators *ValidatorsSession) CreateOrEditValidator(feeAddr common.Address, moniker string, identity string, website string, email string, details string) (*types.Transaction, error) {
	return _Validators.Contract.CreateOrEditValidator(&_Validators.TransactOpts, feeAddr, moniker, identity, website, email, details)
}

// CreateOrEditValidator is a paid mutator transaction binding the contract method 0xa406fcb7.
//
// Solidity: function createOrEditValidator(address feeAddr, string moniker, string identity, string website, string email, string details) returns(bool)
func (_Validators *ValidatorsTransactorSession) CreateOrEditValidator(feeAddr common.Address, moniker string, identity string, website string, email string, details string) (*types.Transaction, error) {
	return _Validators.Contract.CreateOrEditValidator(&_Validators.TransactOpts, feeAddr, moniker, identity, website, email, details)
}

// DistributeBlockReward is a paid mutator transaction binding the contract method 0xd6c0edad.
//
// Solidity: function distributeBlockReward() payable returns()
func (_Validators *ValidatorsTransactor) DistributeBlockReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "distributeBlockReward")
}

// DistributeBlockReward is a paid mutator transaction binding the contract method 0xd6c0edad.
//
// Solidity: function distributeBlockReward() payable returns()
func (_Validators *ValidatorsSession) DistributeBlockReward() (*types.Transaction, error) {
	return _Validators.Contract.DistributeBlockReward(&_Validators.TransactOpts)
}

// DistributeBlockReward is a paid mutator transaction binding the contract method 0xd6c0edad.
//
// Solidity: function distributeBlockReward() payable returns()
func (_Validators *ValidatorsTransactorSession) DistributeBlockReward() (*types.Transaction, error) {
	return _Validators.Contract.DistributeBlockReward(&_Validators.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] vals) returns()
func (_Validators *ValidatorsTransactor) Initialize(opts *bind.TransactOpts, vals []common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "initialize", vals)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] vals) returns()
func (_Validators *ValidatorsSession) Initialize(vals []common.Address) (*types.Transaction, error) {
	return _Validators.Contract.Initialize(&_Validators.TransactOpts, vals)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] vals) returns()
func (_Validators *ValidatorsTransactorSession) Initialize(vals []common.Address) (*types.Transaction, error) {
	return _Validators.Contract.Initialize(&_Validators.TransactOpts, vals)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address val) returns()
func (_Validators *ValidatorsTransactor) RemoveValidator(opts *bind.TransactOpts, val common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "removeValidator", val)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address val) returns()
func (_Validators *ValidatorsSession) RemoveValidator(val common.Address) (*types.Transaction, error) {
	return _Validators.Contract.RemoveValidator(&_Validators.TransactOpts, val)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address val) returns()
func (_Validators *ValidatorsTransactorSession) RemoveValidator(val common.Address) (*types.Transaction, error) {
	return _Validators.Contract.RemoveValidator(&_Validators.TransactOpts, val)
}

// RemoveValidatorIncoming is a paid mutator transaction binding the contract method 0x5dd09590.
//
// Solidity: function removeValidatorIncoming(address val) returns()
func (_Validators *ValidatorsTransactor) RemoveValidatorIncoming(opts *bind.TransactOpts, val common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "removeValidatorIncoming", val)
}

// RemoveValidatorIncoming is a paid mutator transaction binding the contract method 0x5dd09590.
//
// Solidity: function removeValidatorIncoming(address val) returns()
func (_Validators *ValidatorsSession) RemoveValidatorIncoming(val common.Address) (*types.Transaction, error) {
	return _Validators.Contract.RemoveValidatorIncoming(&_Validators.TransactOpts, val)
}

// RemoveValidatorIncoming is a paid mutator transaction binding the contract method 0x5dd09590.
//
// Solidity: function removeValidatorIncoming(address val) returns()
func (_Validators *ValidatorsTransactorSession) RemoveValidatorIncoming(val common.Address) (*types.Transaction, error) {
	return _Validators.Contract.RemoveValidatorIncoming(&_Validators.TransactOpts, val)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address validator) payable returns(bool)
func (_Validators *ValidatorsTransactor) Stake(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "stake", validator)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address validator) payable returns(bool)
func (_Validators *ValidatorsSession) Stake(validator common.Address) (*types.Transaction, error) {
	return _Validators.Contract.Stake(&_Validators.TransactOpts, validator)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address validator) payable returns(bool)
func (_Validators *ValidatorsTransactorSession) Stake(validator common.Address) (*types.Transaction, error) {
	return _Validators.Contract.Stake(&_Validators.TransactOpts, validator)
}

// TryReactive is a paid mutator transaction binding the contract method 0x82bd3d92.
//
// Solidity: function tryReactive(address validator) returns(bool)
func (_Validators *ValidatorsTransactor) TryReactive(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "tryReactive", validator)
}

// TryReactive is a paid mutator transaction binding the contract method 0x82bd3d92.
//
// Solidity: function tryReactive(address validator) returns(bool)
func (_Validators *ValidatorsSession) TryReactive(validator common.Address) (*types.Transaction, error) {
	return _Validators.Contract.TryReactive(&_Validators.TransactOpts, validator)
}

// TryReactive is a paid mutator transaction binding the contract method 0x82bd3d92.
//
// Solidity: function tryReactive(address validator) returns(bool)
func (_Validators *ValidatorsTransactorSession) TryReactive(validator common.Address) (*types.Transaction, error) {
	return _Validators.Contract.TryReactive(&_Validators.TransactOpts, validator)
}

// Unstake is a paid mutator transaction binding the contract method 0xf2888dbb.
//
// Solidity: function unstake(address validator) returns(bool)
func (_Validators *ValidatorsTransactor) Unstake(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "unstake", validator)
}

// Unstake is a paid mutator transaction binding the contract method 0xf2888dbb.
//
// Solidity: function unstake(address validator) returns(bool)
func (_Validators *ValidatorsSession) Unstake(validator common.Address) (*types.Transaction, error) {
	return _Validators.Contract.Unstake(&_Validators.TransactOpts, validator)
}

// Unstake is a paid mutator transaction binding the contract method 0xf2888dbb.
//
// Solidity: function unstake(address validator) returns(bool)
func (_Validators *ValidatorsTransactorSession) Unstake(validator common.Address) (*types.Transaction, error) {
	return _Validators.Contract.Unstake(&_Validators.TransactOpts, validator)
}

// UpdateActiveValidatorSet is a paid mutator transaction binding the contract method 0x6846992a.
//
// Solidity: function updateActiveValidatorSet(address[] newSet, uint256 epoch) returns()
func (_Validators *ValidatorsTransactor) UpdateActiveValidatorSet(opts *bind.TransactOpts, newSet []common.Address, epoch *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "updateActiveValidatorSet", newSet, epoch)
}

// UpdateActiveValidatorSet is a paid mutator transaction binding the contract method 0x6846992a.
//
// Solidity: function updateActiveValidatorSet(address[] newSet, uint256 epoch) returns()
func (_Validators *ValidatorsSession) UpdateActiveValidatorSet(newSet []common.Address, epoch *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.UpdateActiveValidatorSet(&_Validators.TransactOpts, newSet, epoch)
}

// UpdateActiveValidatorSet is a paid mutator transaction binding the contract method 0x6846992a.
//
// Solidity: function updateActiveValidatorSet(address[] newSet, uint256 epoch) returns()
func (_Validators *ValidatorsTransactorSession) UpdateActiveValidatorSet(newSet []common.Address, epoch *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.UpdateActiveValidatorSet(&_Validators.TransactOpts, newSet, epoch)
}

// WithdrawProfits is a paid mutator transaction binding the contract method 0x00362a77.
//
// Solidity: function withdrawProfits(address validator) returns(bool)
func (_Validators *ValidatorsTransactor) WithdrawProfits(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "withdrawProfits", validator)
}

// WithdrawProfits is a paid mutator transaction binding the contract method 0x00362a77.
//
// Solidity: function withdrawProfits(address validator) returns(bool)
func (_Validators *ValidatorsSession) WithdrawProfits(validator common.Address) (*types.Transaction, error) {
	return _Validators.Contract.WithdrawProfits(&_Validators.TransactOpts, validator)
}

// WithdrawProfits is a paid mutator transaction binding the contract method 0x00362a77.
//
// Solidity: function withdrawProfits(address validator) returns(bool)
func (_Validators *ValidatorsTransactorSession) WithdrawProfits(validator common.Address) (*types.Transaction, error) {
	return _Validators.Contract.WithdrawProfits(&_Validators.TransactOpts, validator)
}

// WithdrawStaking is a paid mutator transaction binding the contract method 0x222d3b05.
//
// Solidity: function withdrawStaking(address validator) returns(bool)
func (_Validators *ValidatorsTransactor) WithdrawStaking(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "withdrawStaking", validator)
}

// WithdrawStaking is a paid mutator transaction binding the contract method 0x222d3b05.
//
// Solidity: function withdrawStaking(address validator) returns(bool)
func (_Validators *ValidatorsSession) WithdrawStaking(validator common.Address) (*types.Transaction, error) {
	return _Validators.Contract.WithdrawStaking(&_Validators.TransactOpts, validator)
}

// WithdrawStaking is a paid mutator transaction binding the contract method 0x222d3b05.
//
// Solidity: function withdrawStaking(address validator) returns(bool)
func (_Validators *ValidatorsTransactorSession) WithdrawStaking(validator common.Address) (*types.Transaction, error) {
	return _Validators.Contract.WithdrawStaking(&_Validators.TransactOpts, validator)
}

// ValidatorsLogAddToTopValidatorsIterator is returned from FilterLogAddToTopValidators and is used to iterate over the raw logs and unpacked data for LogAddToTopValidators events raised by the Validators contract.
type ValidatorsLogAddToTopValidatorsIterator struct {
	Event *ValidatorsLogAddToTopValidators // Event containing the contract specifics and raw log

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
func (it *ValidatorsLogAddToTopValidatorsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsLogAddToTopValidators)
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
		it.Event = new(ValidatorsLogAddToTopValidators)
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
func (it *ValidatorsLogAddToTopValidatorsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsLogAddToTopValidatorsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsLogAddToTopValidators represents a LogAddToTopValidators event raised by the Validators contract.
type ValidatorsLogAddToTopValidators struct {
	Val  common.Address
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddToTopValidators is a free log retrieval operation binding the contract event 0x1e3310ad6891b30e03874ec3d1422a6386c5da63d9faf595f5d99eeaf443b99a.
//
// Solidity: event LogAddToTopValidators(address indexed val, uint256 time)
func (_Validators *ValidatorsFilterer) FilterLogAddToTopValidators(opts *bind.FilterOpts, val []common.Address) (*ValidatorsLogAddToTopValidatorsIterator, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "LogAddToTopValidators", valRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsLogAddToTopValidatorsIterator{contract: _Validators.contract, event: "LogAddToTopValidators", logs: logs, sub: sub}, nil
}

// WatchLogAddToTopValidators is a free log subscription operation binding the contract event 0x1e3310ad6891b30e03874ec3d1422a6386c5da63d9faf595f5d99eeaf443b99a.
//
// Solidity: event LogAddToTopValidators(address indexed val, uint256 time)
func (_Validators *ValidatorsFilterer) WatchLogAddToTopValidators(opts *bind.WatchOpts, sink chan<- *ValidatorsLogAddToTopValidators, val []common.Address) (event.Subscription, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "LogAddToTopValidators", valRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsLogAddToTopValidators)
				if err := _Validators.contract.UnpackLog(event, "LogAddToTopValidators", log); err != nil {
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

// ParseLogAddToTopValidators is a log parse operation binding the contract event 0x1e3310ad6891b30e03874ec3d1422a6386c5da63d9faf595f5d99eeaf443b99a.
//
// Solidity: event LogAddToTopValidators(address indexed val, uint256 time)
func (_Validators *ValidatorsFilterer) ParseLogAddToTopValidators(log types.Log) (*ValidatorsLogAddToTopValidators, error) {
	event := new(ValidatorsLogAddToTopValidators)
	if err := _Validators.contract.UnpackLog(event, "LogAddToTopValidators", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsLogCreateValidatorIterator is returned from FilterLogCreateValidator and is used to iterate over the raw logs and unpacked data for LogCreateValidator events raised by the Validators contract.
type ValidatorsLogCreateValidatorIterator struct {
	Event *ValidatorsLogCreateValidator // Event containing the contract specifics and raw log

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
func (it *ValidatorsLogCreateValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsLogCreateValidator)
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
		it.Event = new(ValidatorsLogCreateValidator)
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
func (it *ValidatorsLogCreateValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsLogCreateValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsLogCreateValidator represents a LogCreateValidator event raised by the Validators contract.
type ValidatorsLogCreateValidator struct {
	Val  common.Address
	Fee  common.Address
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogCreateValidator is a free log retrieval operation binding the contract event 0x887eec9d757b7247dd8e51198f9d1b8f27979bceb34bdcc1bffd4ec5ec736c22.
//
// Solidity: event LogCreateValidator(address indexed val, address indexed fee, uint256 time)
func (_Validators *ValidatorsFilterer) FilterLogCreateValidator(opts *bind.FilterOpts, val []common.Address, fee []common.Address) (*ValidatorsLogCreateValidatorIterator, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "LogCreateValidator", valRule, feeRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsLogCreateValidatorIterator{contract: _Validators.contract, event: "LogCreateValidator", logs: logs, sub: sub}, nil
}

// WatchLogCreateValidator is a free log subscription operation binding the contract event 0x887eec9d757b7247dd8e51198f9d1b8f27979bceb34bdcc1bffd4ec5ec736c22.
//
// Solidity: event LogCreateValidator(address indexed val, address indexed fee, uint256 time)
func (_Validators *ValidatorsFilterer) WatchLogCreateValidator(opts *bind.WatchOpts, sink chan<- *ValidatorsLogCreateValidator, val []common.Address, fee []common.Address) (event.Subscription, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "LogCreateValidator", valRule, feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsLogCreateValidator)
				if err := _Validators.contract.UnpackLog(event, "LogCreateValidator", log); err != nil {
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

// ParseLogCreateValidator is a log parse operation binding the contract event 0x887eec9d757b7247dd8e51198f9d1b8f27979bceb34bdcc1bffd4ec5ec736c22.
//
// Solidity: event LogCreateValidator(address indexed val, address indexed fee, uint256 time)
func (_Validators *ValidatorsFilterer) ParseLogCreateValidator(log types.Log) (*ValidatorsLogCreateValidator, error) {
	event := new(ValidatorsLogCreateValidator)
	if err := _Validators.contract.UnpackLog(event, "LogCreateValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsLogDistributeBlockRewardIterator is returned from FilterLogDistributeBlockReward and is used to iterate over the raw logs and unpacked data for LogDistributeBlockReward events raised by the Validators contract.
type ValidatorsLogDistributeBlockRewardIterator struct {
	Event *ValidatorsLogDistributeBlockReward // Event containing the contract specifics and raw log

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
func (it *ValidatorsLogDistributeBlockRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsLogDistributeBlockReward)
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
		it.Event = new(ValidatorsLogDistributeBlockReward)
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
func (it *ValidatorsLogDistributeBlockRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsLogDistributeBlockRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsLogDistributeBlockReward represents a LogDistributeBlockReward event raised by the Validators contract.
type ValidatorsLogDistributeBlockReward struct {
	Coinbase    common.Address
	BlockReward *big.Int
	Time        *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogDistributeBlockReward is a free log retrieval operation binding the contract event 0x7dc4e5df59513708dca355b8706273a5df7b810a4cec8019f2a4b9bb166a1a04.
//
// Solidity: event LogDistributeBlockReward(address indexed coinbase, uint256 blockReward, uint256 time)
func (_Validators *ValidatorsFilterer) FilterLogDistributeBlockReward(opts *bind.FilterOpts, coinbase []common.Address) (*ValidatorsLogDistributeBlockRewardIterator, error) {

	var coinbaseRule []interface{}
	for _, coinbaseItem := range coinbase {
		coinbaseRule = append(coinbaseRule, coinbaseItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "LogDistributeBlockReward", coinbaseRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsLogDistributeBlockRewardIterator{contract: _Validators.contract, event: "LogDistributeBlockReward", logs: logs, sub: sub}, nil
}

// WatchLogDistributeBlockReward is a free log subscription operation binding the contract event 0x7dc4e5df59513708dca355b8706273a5df7b810a4cec8019f2a4b9bb166a1a04.
//
// Solidity: event LogDistributeBlockReward(address indexed coinbase, uint256 blockReward, uint256 time)
func (_Validators *ValidatorsFilterer) WatchLogDistributeBlockReward(opts *bind.WatchOpts, sink chan<- *ValidatorsLogDistributeBlockReward, coinbase []common.Address) (event.Subscription, error) {

	var coinbaseRule []interface{}
	for _, coinbaseItem := range coinbase {
		coinbaseRule = append(coinbaseRule, coinbaseItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "LogDistributeBlockReward", coinbaseRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsLogDistributeBlockReward)
				if err := _Validators.contract.UnpackLog(event, "LogDistributeBlockReward", log); err != nil {
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

// ParseLogDistributeBlockReward is a log parse operation binding the contract event 0x7dc4e5df59513708dca355b8706273a5df7b810a4cec8019f2a4b9bb166a1a04.
//
// Solidity: event LogDistributeBlockReward(address indexed coinbase, uint256 blockReward, uint256 time)
func (_Validators *ValidatorsFilterer) ParseLogDistributeBlockReward(log types.Log) (*ValidatorsLogDistributeBlockReward, error) {
	event := new(ValidatorsLogDistributeBlockReward)
	if err := _Validators.contract.UnpackLog(event, "LogDistributeBlockReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsLogEditValidatorIterator is returned from FilterLogEditValidator and is used to iterate over the raw logs and unpacked data for LogEditValidator events raised by the Validators contract.
type ValidatorsLogEditValidatorIterator struct {
	Event *ValidatorsLogEditValidator // Event containing the contract specifics and raw log

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
func (it *ValidatorsLogEditValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsLogEditValidator)
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
		it.Event = new(ValidatorsLogEditValidator)
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
func (it *ValidatorsLogEditValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsLogEditValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsLogEditValidator represents a LogEditValidator event raised by the Validators contract.
type ValidatorsLogEditValidator struct {
	Val  common.Address
	Fee  common.Address
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogEditValidator is a free log retrieval operation binding the contract event 0xb8421f65501371f54d58de1937ff1e1ccdb76423ef6f84acea1814a0f6362ca0.
//
// Solidity: event LogEditValidator(address indexed val, address indexed fee, uint256 time)
func (_Validators *ValidatorsFilterer) FilterLogEditValidator(opts *bind.FilterOpts, val []common.Address, fee []common.Address) (*ValidatorsLogEditValidatorIterator, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "LogEditValidator", valRule, feeRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsLogEditValidatorIterator{contract: _Validators.contract, event: "LogEditValidator", logs: logs, sub: sub}, nil
}

// WatchLogEditValidator is a free log subscription operation binding the contract event 0xb8421f65501371f54d58de1937ff1e1ccdb76423ef6f84acea1814a0f6362ca0.
//
// Solidity: event LogEditValidator(address indexed val, address indexed fee, uint256 time)
func (_Validators *ValidatorsFilterer) WatchLogEditValidator(opts *bind.WatchOpts, sink chan<- *ValidatorsLogEditValidator, val []common.Address, fee []common.Address) (event.Subscription, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "LogEditValidator", valRule, feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsLogEditValidator)
				if err := _Validators.contract.UnpackLog(event, "LogEditValidator", log); err != nil {
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

// ParseLogEditValidator is a log parse operation binding the contract event 0xb8421f65501371f54d58de1937ff1e1ccdb76423ef6f84acea1814a0f6362ca0.
//
// Solidity: event LogEditValidator(address indexed val, address indexed fee, uint256 time)
func (_Validators *ValidatorsFilterer) ParseLogEditValidator(log types.Log) (*ValidatorsLogEditValidator, error) {
	event := new(ValidatorsLogEditValidator)
	if err := _Validators.contract.UnpackLog(event, "LogEditValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsLogReactiveIterator is returned from FilterLogReactive and is used to iterate over the raw logs and unpacked data for LogReactive events raised by the Validators contract.
type ValidatorsLogReactiveIterator struct {
	Event *ValidatorsLogReactive // Event containing the contract specifics and raw log

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
func (it *ValidatorsLogReactiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsLogReactive)
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
		it.Event = new(ValidatorsLogReactive)
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
func (it *ValidatorsLogReactiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsLogReactiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsLogReactive represents a LogReactive event raised by the Validators contract.
type ValidatorsLogReactive struct {
	Val  common.Address
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogReactive is a free log retrieval operation binding the contract event 0xd8b2c426ec1be69ca7583d26b1e893946e3227430d3ebc3bd64d9e1c469cb400.
//
// Solidity: event LogReactive(address indexed val, uint256 time)
func (_Validators *ValidatorsFilterer) FilterLogReactive(opts *bind.FilterOpts, val []common.Address) (*ValidatorsLogReactiveIterator, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "LogReactive", valRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsLogReactiveIterator{contract: _Validators.contract, event: "LogReactive", logs: logs, sub: sub}, nil
}

// WatchLogReactive is a free log subscription operation binding the contract event 0xd8b2c426ec1be69ca7583d26b1e893946e3227430d3ebc3bd64d9e1c469cb400.
//
// Solidity: event LogReactive(address indexed val, uint256 time)
func (_Validators *ValidatorsFilterer) WatchLogReactive(opts *bind.WatchOpts, sink chan<- *ValidatorsLogReactive, val []common.Address) (event.Subscription, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "LogReactive", valRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsLogReactive)
				if err := _Validators.contract.UnpackLog(event, "LogReactive", log); err != nil {
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

// ParseLogReactive is a log parse operation binding the contract event 0xd8b2c426ec1be69ca7583d26b1e893946e3227430d3ebc3bd64d9e1c469cb400.
//
// Solidity: event LogReactive(address indexed val, uint256 time)
func (_Validators *ValidatorsFilterer) ParseLogReactive(log types.Log) (*ValidatorsLogReactive, error) {
	event := new(ValidatorsLogReactive)
	if err := _Validators.contract.UnpackLog(event, "LogReactive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsLogRemoveFromTopValidatorsIterator is returned from FilterLogRemoveFromTopValidators and is used to iterate over the raw logs and unpacked data for LogRemoveFromTopValidators events raised by the Validators contract.
type ValidatorsLogRemoveFromTopValidatorsIterator struct {
	Event *ValidatorsLogRemoveFromTopValidators // Event containing the contract specifics and raw log

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
func (it *ValidatorsLogRemoveFromTopValidatorsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsLogRemoveFromTopValidators)
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
		it.Event = new(ValidatorsLogRemoveFromTopValidators)
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
func (it *ValidatorsLogRemoveFromTopValidatorsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsLogRemoveFromTopValidatorsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsLogRemoveFromTopValidators represents a LogRemoveFromTopValidators event raised by the Validators contract.
type ValidatorsLogRemoveFromTopValidators struct {
	Val  common.Address
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogRemoveFromTopValidators is a free log retrieval operation binding the contract event 0x7521e44559c870c316e84e60bc4785d9c034a8ab1d6acdce8134ac03f946c6ed.
//
// Solidity: event LogRemoveFromTopValidators(address indexed val, uint256 time)
func (_Validators *ValidatorsFilterer) FilterLogRemoveFromTopValidators(opts *bind.FilterOpts, val []common.Address) (*ValidatorsLogRemoveFromTopValidatorsIterator, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "LogRemoveFromTopValidators", valRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsLogRemoveFromTopValidatorsIterator{contract: _Validators.contract, event: "LogRemoveFromTopValidators", logs: logs, sub: sub}, nil
}

// WatchLogRemoveFromTopValidators is a free log subscription operation binding the contract event 0x7521e44559c870c316e84e60bc4785d9c034a8ab1d6acdce8134ac03f946c6ed.
//
// Solidity: event LogRemoveFromTopValidators(address indexed val, uint256 time)
func (_Validators *ValidatorsFilterer) WatchLogRemoveFromTopValidators(opts *bind.WatchOpts, sink chan<- *ValidatorsLogRemoveFromTopValidators, val []common.Address) (event.Subscription, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "LogRemoveFromTopValidators", valRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsLogRemoveFromTopValidators)
				if err := _Validators.contract.UnpackLog(event, "LogRemoveFromTopValidators", log); err != nil {
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

// ParseLogRemoveFromTopValidators is a log parse operation binding the contract event 0x7521e44559c870c316e84e60bc4785d9c034a8ab1d6acdce8134ac03f946c6ed.
//
// Solidity: event LogRemoveFromTopValidators(address indexed val, uint256 time)
func (_Validators *ValidatorsFilterer) ParseLogRemoveFromTopValidators(log types.Log) (*ValidatorsLogRemoveFromTopValidators, error) {
	event := new(ValidatorsLogRemoveFromTopValidators)
	if err := _Validators.contract.UnpackLog(event, "LogRemoveFromTopValidators", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsLogRemoveValidatorIterator is returned from FilterLogRemoveValidator and is used to iterate over the raw logs and unpacked data for LogRemoveValidator events raised by the Validators contract.
type ValidatorsLogRemoveValidatorIterator struct {
	Event *ValidatorsLogRemoveValidator // Event containing the contract specifics and raw log

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
func (it *ValidatorsLogRemoveValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsLogRemoveValidator)
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
		it.Event = new(ValidatorsLogRemoveValidator)
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
func (it *ValidatorsLogRemoveValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsLogRemoveValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsLogRemoveValidator represents a LogRemoveValidator event raised by the Validators contract.
type ValidatorsLogRemoveValidator struct {
	Val  common.Address
	Hb   *big.Int
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogRemoveValidator is a free log retrieval operation binding the contract event 0xa26de7ab324eac08c596549f421e5c8741213d237d2e9a2c9c0ebde0a7a849fe.
//
// Solidity: event LogRemoveValidator(address indexed val, uint256 hb, uint256 time)
func (_Validators *ValidatorsFilterer) FilterLogRemoveValidator(opts *bind.FilterOpts, val []common.Address) (*ValidatorsLogRemoveValidatorIterator, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "LogRemoveValidator", valRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsLogRemoveValidatorIterator{contract: _Validators.contract, event: "LogRemoveValidator", logs: logs, sub: sub}, nil
}

// WatchLogRemoveValidator is a free log subscription operation binding the contract event 0xa26de7ab324eac08c596549f421e5c8741213d237d2e9a2c9c0ebde0a7a849fe.
//
// Solidity: event LogRemoveValidator(address indexed val, uint256 hb, uint256 time)
func (_Validators *ValidatorsFilterer) WatchLogRemoveValidator(opts *bind.WatchOpts, sink chan<- *ValidatorsLogRemoveValidator, val []common.Address) (event.Subscription, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "LogRemoveValidator", valRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsLogRemoveValidator)
				if err := _Validators.contract.UnpackLog(event, "LogRemoveValidator", log); err != nil {
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

// ParseLogRemoveValidator is a log parse operation binding the contract event 0xa26de7ab324eac08c596549f421e5c8741213d237d2e9a2c9c0ebde0a7a849fe.
//
// Solidity: event LogRemoveValidator(address indexed val, uint256 hb, uint256 time)
func (_Validators *ValidatorsFilterer) ParseLogRemoveValidator(log types.Log) (*ValidatorsLogRemoveValidator, error) {
	event := new(ValidatorsLogRemoveValidator)
	if err := _Validators.contract.UnpackLog(event, "LogRemoveValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsLogRemoveValidatorIncomingIterator is returned from FilterLogRemoveValidatorIncoming and is used to iterate over the raw logs and unpacked data for LogRemoveValidatorIncoming events raised by the Validators contract.
type ValidatorsLogRemoveValidatorIncomingIterator struct {
	Event *ValidatorsLogRemoveValidatorIncoming // Event containing the contract specifics and raw log

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
func (it *ValidatorsLogRemoveValidatorIncomingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsLogRemoveValidatorIncoming)
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
		it.Event = new(ValidatorsLogRemoveValidatorIncoming)
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
func (it *ValidatorsLogRemoveValidatorIncomingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsLogRemoveValidatorIncomingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsLogRemoveValidatorIncoming represents a LogRemoveValidatorIncoming event raised by the Validators contract.
type ValidatorsLogRemoveValidatorIncoming struct {
	Val  common.Address
	Hb   *big.Int
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogRemoveValidatorIncoming is a free log retrieval operation binding the contract event 0xe294e9d73f8eee23e21b2e1567960625a6b5d339cb127b55d0d09473a9951235.
//
// Solidity: event LogRemoveValidatorIncoming(address indexed val, uint256 hb, uint256 time)
func (_Validators *ValidatorsFilterer) FilterLogRemoveValidatorIncoming(opts *bind.FilterOpts, val []common.Address) (*ValidatorsLogRemoveValidatorIncomingIterator, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "LogRemoveValidatorIncoming", valRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsLogRemoveValidatorIncomingIterator{contract: _Validators.contract, event: "LogRemoveValidatorIncoming", logs: logs, sub: sub}, nil
}

// WatchLogRemoveValidatorIncoming is a free log subscription operation binding the contract event 0xe294e9d73f8eee23e21b2e1567960625a6b5d339cb127b55d0d09473a9951235.
//
// Solidity: event LogRemoveValidatorIncoming(address indexed val, uint256 hb, uint256 time)
func (_Validators *ValidatorsFilterer) WatchLogRemoveValidatorIncoming(opts *bind.WatchOpts, sink chan<- *ValidatorsLogRemoveValidatorIncoming, val []common.Address) (event.Subscription, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "LogRemoveValidatorIncoming", valRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsLogRemoveValidatorIncoming)
				if err := _Validators.contract.UnpackLog(event, "LogRemoveValidatorIncoming", log); err != nil {
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

// ParseLogRemoveValidatorIncoming is a log parse operation binding the contract event 0xe294e9d73f8eee23e21b2e1567960625a6b5d339cb127b55d0d09473a9951235.
//
// Solidity: event LogRemoveValidatorIncoming(address indexed val, uint256 hb, uint256 time)
func (_Validators *ValidatorsFilterer) ParseLogRemoveValidatorIncoming(log types.Log) (*ValidatorsLogRemoveValidatorIncoming, error) {
	event := new(ValidatorsLogRemoveValidatorIncoming)
	if err := _Validators.contract.UnpackLog(event, "LogRemoveValidatorIncoming", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsLogStakeIterator is returned from FilterLogStake and is used to iterate over the raw logs and unpacked data for LogStake events raised by the Validators contract.
type ValidatorsLogStakeIterator struct {
	Event *ValidatorsLogStake // Event containing the contract specifics and raw log

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
func (it *ValidatorsLogStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsLogStake)
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
		it.Event = new(ValidatorsLogStake)
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
func (it *ValidatorsLogStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsLogStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsLogStake represents a LogStake event raised by the Validators contract.
type ValidatorsLogStake struct {
	Staker  common.Address
	Val     common.Address
	Staking *big.Int
	Time    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogStake is a free log retrieval operation binding the contract event 0xb9ba725934532316cffe10975da6eb25ad49c2d1c294d982c46c9f8d684ee075.
//
// Solidity: event LogStake(address indexed staker, address indexed val, uint256 staking, uint256 time)
func (_Validators *ValidatorsFilterer) FilterLogStake(opts *bind.FilterOpts, staker []common.Address, val []common.Address) (*ValidatorsLogStakeIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "LogStake", stakerRule, valRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsLogStakeIterator{contract: _Validators.contract, event: "LogStake", logs: logs, sub: sub}, nil
}

// WatchLogStake is a free log subscription operation binding the contract event 0xb9ba725934532316cffe10975da6eb25ad49c2d1c294d982c46c9f8d684ee075.
//
// Solidity: event LogStake(address indexed staker, address indexed val, uint256 staking, uint256 time)
func (_Validators *ValidatorsFilterer) WatchLogStake(opts *bind.WatchOpts, sink chan<- *ValidatorsLogStake, staker []common.Address, val []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "LogStake", stakerRule, valRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsLogStake)
				if err := _Validators.contract.UnpackLog(event, "LogStake", log); err != nil {
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

// ParseLogStake is a log parse operation binding the contract event 0xb9ba725934532316cffe10975da6eb25ad49c2d1c294d982c46c9f8d684ee075.
//
// Solidity: event LogStake(address indexed staker, address indexed val, uint256 staking, uint256 time)
func (_Validators *ValidatorsFilterer) ParseLogStake(log types.Log) (*ValidatorsLogStake, error) {
	event := new(ValidatorsLogStake)
	if err := _Validators.contract.UnpackLog(event, "LogStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsLogUnstakeIterator is returned from FilterLogUnstake and is used to iterate over the raw logs and unpacked data for LogUnstake events raised by the Validators contract.
type ValidatorsLogUnstakeIterator struct {
	Event *ValidatorsLogUnstake // Event containing the contract specifics and raw log

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
func (it *ValidatorsLogUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsLogUnstake)
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
		it.Event = new(ValidatorsLogUnstake)
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
func (it *ValidatorsLogUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsLogUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsLogUnstake represents a LogUnstake event raised by the Validators contract.
type ValidatorsLogUnstake struct {
	Staker common.Address
	Val    common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogUnstake is a free log retrieval operation binding the contract event 0x449002ae18e748d69a55f38514400d64f966492e593e32d6e9b8b24db98a0bc1.
//
// Solidity: event LogUnstake(address indexed staker, address indexed val, uint256 amount, uint256 time)
func (_Validators *ValidatorsFilterer) FilterLogUnstake(opts *bind.FilterOpts, staker []common.Address, val []common.Address) (*ValidatorsLogUnstakeIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "LogUnstake", stakerRule, valRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsLogUnstakeIterator{contract: _Validators.contract, event: "LogUnstake", logs: logs, sub: sub}, nil
}

// WatchLogUnstake is a free log subscription operation binding the contract event 0x449002ae18e748d69a55f38514400d64f966492e593e32d6e9b8b24db98a0bc1.
//
// Solidity: event LogUnstake(address indexed staker, address indexed val, uint256 amount, uint256 time)
func (_Validators *ValidatorsFilterer) WatchLogUnstake(opts *bind.WatchOpts, sink chan<- *ValidatorsLogUnstake, staker []common.Address, val []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "LogUnstake", stakerRule, valRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsLogUnstake)
				if err := _Validators.contract.UnpackLog(event, "LogUnstake", log); err != nil {
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

// ParseLogUnstake is a log parse operation binding the contract event 0x449002ae18e748d69a55f38514400d64f966492e593e32d6e9b8b24db98a0bc1.
//
// Solidity: event LogUnstake(address indexed staker, address indexed val, uint256 amount, uint256 time)
func (_Validators *ValidatorsFilterer) ParseLogUnstake(log types.Log) (*ValidatorsLogUnstake, error) {
	event := new(ValidatorsLogUnstake)
	if err := _Validators.contract.UnpackLog(event, "LogUnstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsLogUpdateValidatorIterator is returned from FilterLogUpdateValidator and is used to iterate over the raw logs and unpacked data for LogUpdateValidator events raised by the Validators contract.
type ValidatorsLogUpdateValidatorIterator struct {
	Event *ValidatorsLogUpdateValidator // Event containing the contract specifics and raw log

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
func (it *ValidatorsLogUpdateValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsLogUpdateValidator)
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
		it.Event = new(ValidatorsLogUpdateValidator)
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
func (it *ValidatorsLogUpdateValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsLogUpdateValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsLogUpdateValidator represents a LogUpdateValidator event raised by the Validators contract.
type ValidatorsLogUpdateValidator struct {
	NewSet []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogUpdateValidator is a free log retrieval operation binding the contract event 0xeacea8f3c22f06c0b18306bdb04d0a967255129e8ce0094debb0a0ff89d006b5.
//
// Solidity: event LogUpdateValidator(address[] newSet)
func (_Validators *ValidatorsFilterer) FilterLogUpdateValidator(opts *bind.FilterOpts) (*ValidatorsLogUpdateValidatorIterator, error) {

	logs, sub, err := _Validators.contract.FilterLogs(opts, "LogUpdateValidator")
	if err != nil {
		return nil, err
	}
	return &ValidatorsLogUpdateValidatorIterator{contract: _Validators.contract, event: "LogUpdateValidator", logs: logs, sub: sub}, nil
}

// WatchLogUpdateValidator is a free log subscription operation binding the contract event 0xeacea8f3c22f06c0b18306bdb04d0a967255129e8ce0094debb0a0ff89d006b5.
//
// Solidity: event LogUpdateValidator(address[] newSet)
func (_Validators *ValidatorsFilterer) WatchLogUpdateValidator(opts *bind.WatchOpts, sink chan<- *ValidatorsLogUpdateValidator) (event.Subscription, error) {

	logs, sub, err := _Validators.contract.WatchLogs(opts, "LogUpdateValidator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsLogUpdateValidator)
				if err := _Validators.contract.UnpackLog(event, "LogUpdateValidator", log); err != nil {
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

// ParseLogUpdateValidator is a log parse operation binding the contract event 0xeacea8f3c22f06c0b18306bdb04d0a967255129e8ce0094debb0a0ff89d006b5.
//
// Solidity: event LogUpdateValidator(address[] newSet)
func (_Validators *ValidatorsFilterer) ParseLogUpdateValidator(log types.Log) (*ValidatorsLogUpdateValidator, error) {
	event := new(ValidatorsLogUpdateValidator)
	if err := _Validators.contract.UnpackLog(event, "LogUpdateValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsLogWithdrawProfitsIterator is returned from FilterLogWithdrawProfits and is used to iterate over the raw logs and unpacked data for LogWithdrawProfits events raised by the Validators contract.
type ValidatorsLogWithdrawProfitsIterator struct {
	Event *ValidatorsLogWithdrawProfits // Event containing the contract specifics and raw log

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
func (it *ValidatorsLogWithdrawProfitsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsLogWithdrawProfits)
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
		it.Event = new(ValidatorsLogWithdrawProfits)
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
func (it *ValidatorsLogWithdrawProfitsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsLogWithdrawProfitsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsLogWithdrawProfits represents a LogWithdrawProfits event raised by the Validators contract.
type ValidatorsLogWithdrawProfits struct {
	Val  common.Address
	Fee  common.Address
	Hb   *big.Int
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogWithdrawProfits is a free log retrieval operation binding the contract event 0x51a69b4502f660774c9339825c7b5adbf0b8622289134647e29728ec5d9b3bb9.
//
// Solidity: event LogWithdrawProfits(address indexed val, address indexed fee, uint256 hb, uint256 time)
func (_Validators *ValidatorsFilterer) FilterLogWithdrawProfits(opts *bind.FilterOpts, val []common.Address, fee []common.Address) (*ValidatorsLogWithdrawProfitsIterator, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "LogWithdrawProfits", valRule, feeRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsLogWithdrawProfitsIterator{contract: _Validators.contract, event: "LogWithdrawProfits", logs: logs, sub: sub}, nil
}

// WatchLogWithdrawProfits is a free log subscription operation binding the contract event 0x51a69b4502f660774c9339825c7b5adbf0b8622289134647e29728ec5d9b3bb9.
//
// Solidity: event LogWithdrawProfits(address indexed val, address indexed fee, uint256 hb, uint256 time)
func (_Validators *ValidatorsFilterer) WatchLogWithdrawProfits(opts *bind.WatchOpts, sink chan<- *ValidatorsLogWithdrawProfits, val []common.Address, fee []common.Address) (event.Subscription, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "LogWithdrawProfits", valRule, feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsLogWithdrawProfits)
				if err := _Validators.contract.UnpackLog(event, "LogWithdrawProfits", log); err != nil {
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

// ParseLogWithdrawProfits is a log parse operation binding the contract event 0x51a69b4502f660774c9339825c7b5adbf0b8622289134647e29728ec5d9b3bb9.
//
// Solidity: event LogWithdrawProfits(address indexed val, address indexed fee, uint256 hb, uint256 time)
func (_Validators *ValidatorsFilterer) ParseLogWithdrawProfits(log types.Log) (*ValidatorsLogWithdrawProfits, error) {
	event := new(ValidatorsLogWithdrawProfits)
	if err := _Validators.contract.UnpackLog(event, "LogWithdrawProfits", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsLogWithdrawStakingIterator is returned from FilterLogWithdrawStaking and is used to iterate over the raw logs and unpacked data for LogWithdrawStaking events raised by the Validators contract.
type ValidatorsLogWithdrawStakingIterator struct {
	Event *ValidatorsLogWithdrawStaking // Event containing the contract specifics and raw log

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
func (it *ValidatorsLogWithdrawStakingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsLogWithdrawStaking)
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
		it.Event = new(ValidatorsLogWithdrawStaking)
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
func (it *ValidatorsLogWithdrawStakingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsLogWithdrawStakingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsLogWithdrawStaking represents a LogWithdrawStaking event raised by the Validators contract.
type ValidatorsLogWithdrawStaking struct {
	Staker common.Address
	Val    common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogWithdrawStaking is a free log retrieval operation binding the contract event 0xa70cd94070cd852339a76b32cf2d95a3c8f2a322269163d276071c1c14955619.
//
// Solidity: event LogWithdrawStaking(address indexed staker, address indexed val, uint256 amount, uint256 time)
func (_Validators *ValidatorsFilterer) FilterLogWithdrawStaking(opts *bind.FilterOpts, staker []common.Address, val []common.Address) (*ValidatorsLogWithdrawStakingIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "LogWithdrawStaking", stakerRule, valRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsLogWithdrawStakingIterator{contract: _Validators.contract, event: "LogWithdrawStaking", logs: logs, sub: sub}, nil
}

// WatchLogWithdrawStaking is a free log subscription operation binding the contract event 0xa70cd94070cd852339a76b32cf2d95a3c8f2a322269163d276071c1c14955619.
//
// Solidity: event LogWithdrawStaking(address indexed staker, address indexed val, uint256 amount, uint256 time)
func (_Validators *ValidatorsFilterer) WatchLogWithdrawStaking(opts *bind.WatchOpts, sink chan<- *ValidatorsLogWithdrawStaking, staker []common.Address, val []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "LogWithdrawStaking", stakerRule, valRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsLogWithdrawStaking)
				if err := _Validators.contract.UnpackLog(event, "LogWithdrawStaking", log); err != nil {
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

// ParseLogWithdrawStaking is a log parse operation binding the contract event 0xa70cd94070cd852339a76b32cf2d95a3c8f2a322269163d276071c1c14955619.
//
// Solidity: event LogWithdrawStaking(address indexed staker, address indexed val, uint256 amount, uint256 time)
func (_Validators *ValidatorsFilterer) ParseLogWithdrawStaking(log types.Log) (*ValidatorsLogWithdrawStaking, error) {
	event := new(ValidatorsLogWithdrawStaking)
	if err := _Validators.contract.UnpackLog(event, "LogWithdrawStaking", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

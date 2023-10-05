package near_rpc

import (
	"fmt"

	types "github.com/nexeranet/gonear/near_rpc/types"
)

// Returns the contract code (Wasm binary) deployed to the account. Please note
// that the returned code will be encoded in base64.
func (a *Request) ViewContractCode(accountId string) (*types.ContractCodeView, error) {
	type Params struct {
		RequestType string `json:"request_type"`
		Finality    string `json:"finality"`
		AccountID   string `json:"account_id"`
	}
	params := Params{"view_code", "final", accountId}
	response, err := a.Call("query", &params)
	if err != nil {
		return nil, err
	}
	var raw types.ContractCodeView
	return &raw, response.GetObject(&raw)
}

func (a *NearRpc) ViewContractCode(accountId string) (*types.ContractCodeView, error) {
	return a.Request().ViewContractCode(accountId)
}

// Returns the contract code (Wasm binary) deployed to the account of a specific block. Please note
// that the returned code will be encoded in base64.
func (a *Request) ViewContractCodeByBlockId(accountId string, blockId uint64) (*types.ContractCodeView, error) {
	type Params struct {
		RequestType string `json:"request_type"`
		BlockId     uint64 `json:"block_id"`
		AccountID   string `json:"account_id"`
	}
	params := Params{"view_code", blockId, accountId}
	response, err := a.Call("query", &params)
	if err != nil {
		return nil, err
	}
	var raw types.ContractCodeView
	return &raw, response.GetObject(&raw)
}

func (a *NearRpc) ViewContractCodeByBlockId(accountId string, blockId uint64) (*types.ContractCodeView, error) {
	return a.Request().ViewContractCodeByBlockId(accountId, blockId)
}

// Returns the state (key value pairs) of a contract based on the key prefix
// (base64 encoded). Pass an empty string for prefix_base64 if you would like to
// return the entire state. Please note that the returned state will be base64
// encoded as well.
func (a *Request) ViewContractState(accountId, prefixBase64 string) (*types.ContractStateView, error) {
	type Params struct {
		RequestType  string `json:"request_type"`
		Finality     string `json:"finality"`
		AccountID    string `json:"account_id"`
		PrefixBase64 string `json:"prefix_base64"`
	}
	params := Params{"view_code", "final", accountId, prefixBase64}
	response, err := a.Call("query", &params)
	if err != nil {
		return nil, err
	}
	var raw types.ContractStateView
	return &raw, response.GetObject(&raw)
}

func (a *NearRpc) ViewContractState(accountId, prefixBase64 string) (*types.ContractStateView, error) {
	return a.Request().ViewContractState(accountId, prefixBase64)
}

// Returns the state (key value pairs) of a contract based on the key prefix
// (base64 encoded) of a specific block. Pass an empty string for prefix_base64 if you would like to
// return the entire state. Please note that the returned state will be base64
// encoded as well.
func (a *Request) ViewContractStateByBlockId(accountId, prefixBase64 string, blockId uint64) (*types.ContractStateView, error) {
	type Params struct {
		RequestType  string `json:"request_type"`
		BlockId      uint64 `json:"block_id"`
		AccountID    string `json:"account_id"`
		PrefixBase64 string `json:"prefix_base64"`
	}
	params := Params{"view_code", blockId, accountId, prefixBase64}
	response, err := a.Call("query", &params)
	if err != nil {
		return nil, err
	}
	var raw types.ContractStateView
	return &raw, response.GetObject(&raw)
}

func (a *NearRpc) ViewContractStateByBlockId(accountId, prefixBase64 string, blockId uint64) (*types.ContractStateView, error) {
	return a.Request().ViewContractStateByBlockId(accountId, prefixBase64, blockId)
}

// Call a contract method as a view function.
func (a *Request) CallContractFunc(accountId, method_name, args_base64 string) (*types.ContractFuncResult, error) {
	type Params struct {
		RequestType string `json:"request_type"`
		Finality    string `json:"finality"`
		AccountID   string `json:"account_id"`
		MethodName  string `json:"method_name"`
		ArgsBase64  string `json:"args_base64"`
	}
	params := Params{"call_function", "final", accountId, method_name, args_base64}
	response, err := a.Call("query", &params)
	if err != nil {
		return nil, err
	}
	var raw types.ContractFuncResult
	err = response.GetObject(&raw)
	if err != nil {
		return nil, err
	}
	if raw.Error != "" {
		return nil, fmt.Errorf(raw.Error)
	}
	return &raw, nil
}

func (a *NearRpc) CallContractFunc(accountId, method_name, args_base64 string) (*types.ContractFuncResult, error) {
	return a.Request().CallContractFunc(accountId, method_name, args_base64)
}

// Call a contract method as a view function by block id.
func (a *Request) CallContractFuncByBlockId(accountId, method_name, args_base64 string, block_id uint64) (*types.ContractFuncResult, error) {
	type Params struct {
		RequestType string `json:"request_type"`
		BlockId     uint64 `json:"block_id"`
		AccountID   string `json:"account_id"`
		MethodName  string `json:"method_name"`
		ArgsBase64  string `json:"args_base64"`
	}
	params := Params{"call_function", block_id, accountId, method_name, args_base64}
	response, err := a.Call("query", &params)
	if err != nil {
		return nil, err
	}
	var raw types.ContractFuncResult
	err = response.GetObject(&raw)
	if err != nil {
		return nil, err
	}
	if raw.Error != "" {
		return nil, fmt.Errorf(raw.Error)
	}
	return &raw, nil
}

func (a *NearRpc) CallContractFuncByBlockId(accountId, method_name, args_base64 string, block_id uint64) (*types.ContractFuncResult, error) {
	return a.Request().CallContractFuncByBlockId(accountId, method_name, args_base64, block_id)
}

// Returns the state change details of a contract based on the key prefix (encoded
// to base64). Pass an empty string for this param if you would like to return all
// state changes.
func (a *Request) ViewContractStateChanges(accountIds []string, keyPrefixBase64 string) (*types.ContractStateChangesView, error) {
	type Params struct {
		ChangesType     string   `json:"changes_type"`
		Finality        string   `json:"finality"`
		AccountIds      []string `json:"account_ids"`
		KeyPrefixBase64 string   `json:"key_prefix_base64"`
	}
	params := Params{
		ChangesType:     "data_changes",
		Finality:        "final",
		AccountIds:      accountIds,
		KeyPrefixBase64: keyPrefixBase64,
	}
	response, err := a.Call("EXPERIMENTAL_changes", &params)
	if err != nil {
		return nil, err
	}
	var raw types.ContractStateChangesView
	return &raw, response.GetObject(&raw)
}

func (a *NearRpc) ViewContractStateChanges(accountIds []string, keyPrefixBase64 string) (*types.ContractStateChangesView, error) {
	return a.Request().ViewContractStateChanges(accountIds, keyPrefixBase64)
}

// Returns the state change details of a contract based on the key prefix (encoded
// to base64) of a specific block. Pass an empty string for this param if you would like to return all
// state changes.
func (a *Request) ViewContractStateChangesByBlockId(accountIds []string, keyPrefixBase64 string, blockId uint64) (*types.ContractStateChangesView, error) {
	type Params struct {
		ChangesType     string   `json:"changes_type"`
		AccountIds      []string `json:"account_ids"`
		KeyPrefixBase64 string   `json:"key_prefix_base64"`
		BlockId         uint64   `json:"block_id"`
	}
	params := Params{
		ChangesType:     "data_changes",
		AccountIds:      accountIds,
		KeyPrefixBase64: keyPrefixBase64,
		BlockId:         blockId,
	}
	response, err := a.Call("EXPERIMENTAL_changes", &params)
	if err != nil {
		return nil, err
	}
	var raw types.ContractStateChangesView
	return &raw, response.GetObject(&raw)
}

func (a *NearRpc) ViewContractStateChangesByBlockId(accountIds []string, keyPrefixBase64 string, blockId uint64) (*types.ContractStateChangesView, error) {
	return a.Request().ViewContractStateChangesByBlockId(accountIds, keyPrefixBase64, blockId)
}

// Returns code changes made when deploying a contract. Change is returned is a
// base64 encoded WASM file.
func (a *Request) ViewContractCodeChanges(accountIds []string) (*types.ContractCodeChangesView, error) {
	type Params struct {
		ChangesType string   `json:"changes_type"`
		AccountIds  []string `json:"account_ids"`
		Finality    string   `json:"finality"`
	}
	params := Params{
		ChangesType: "contract_code_changes",
		AccountIds:  accountIds,
		Finality:    "final",
	}
	response, err := a.Call("EXPERIMENTAL_changes", &params)
	if err != nil {
		return nil, err
	}
	var raw types.ContractCodeChangesView
	return &raw, response.GetObject(&raw)
}

func (a *NearRpc) ViewContractCodeChanges(accountIds []string) (*types.ContractCodeChangesView, error) {
	return a.Request().ViewContractCodeChanges(accountIds)
}

// Returns code changes made when deploying a contract of a specific block.
// Change is returned is a base64 encoded WASM file.
func (a *Request) ViewContractCodeChangesByBlockId(accountIds []string, blockId uint64) (*types.ContractCodeChangesView, error) {
	type Params struct {
		ChangesType string   `json:"changes_type"`
		AccountIds  []string `json:"account_ids"`
		BlockId     uint64   `json:"block_id"`
	}
	params := Params{
		ChangesType: "contract_code_changes",
		AccountIds:  accountIds,
		BlockId:     blockId,
	}
	response, err := a.Call("EXPERIMENTAL_changes", &params)
	if err != nil {
		return nil, err
	}
	var raw types.ContractCodeChangesView
	return &raw, response.GetObject(&raw)
}

func (a *NearRpc) ViewContractCodeChangesByBlockId(accountIds []string, blockId uint64) (*types.ContractCodeChangesView, error) {
	return a.Request().ViewContractCodeChangesByBlockId(accountIds, blockId)
}

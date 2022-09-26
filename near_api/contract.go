package near_api

import (
	"fmt"

	types "github.com/nexeranet/gonear/near_api/types"
)

func (a *NearApi) ViewContractCode(accountId string) (*types.ContractCodeView, error) {
	type Params struct {
		RequestType string `json:"request_type"`
		Finality    string `json:"finality"`
		AccountID   string `json:"account_id"`
	}
	params := Params{"view_code", "final", accountId}
	response, err := a.c.Call("query", &params)
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	var raw types.ContractCodeView
	return &raw, response.GetObject(&raw)
}

func (a *NearApi) ViewContractState(accountId, prefixBase64 string) (*types.ContractStateView, error) {
	type Params struct {
		RequestType  string `json:"request_type"`
		Finality     string `json:"finality"`
		AccountID    string `json:"account_id"`
		PrefixBase64 string `json:"prefix_base64"`
	}
	params := Params{"view_code", "final", accountId, prefixBase64}
	response, err := a.c.Call("query", &params)
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	var raw types.ContractStateView
	return &raw, response.GetObject(&raw)
}

func (a *NearApi) CallContractFunc(accountId, method_name, args_base64 string) (*types.ContractFuncResult, error) {
	type Params struct {
		RequestType string `json:"request_type"`
		Finality    string `json:"finality"`
		AccountID   string `json:"account_id"`
		MethodName  string `json:"method_name"`
		ArgsBase64  string `json:"args_base64"`
	}
	params := Params{"call_function", "final", accountId, method_name, args_base64}
	response, err := a.c.Call("query", &params)
	if err := a.checkError(err, response); err != nil {
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

func (a *NearApi) ViewContractStateChanges(accountIds []string, keyPrefixBase64 string) (raw *types.ContractStateChangesView, err error) {
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
	response, err := a.c.Call("EXPERIMENTAL_changes", &params)
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	return raw, response.GetObject(raw)
}

func (a *NearApi) ViewContractStateChangesByBlockId(accountIds []string, keyPrefixBase64 string, blockId uint64) (raw *types.ContractStateChangesView, err error) {
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
	response, err := a.c.Call("EXPERIMENTAL_changes", &params)
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	return raw, response.GetObject(raw)
}

func (a *NearApi) ViewContractCodeChanges(accountIds []string) (raw *types.ContractCodeChangesView, err error) {
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
	response, err := a.c.Call("EXPERIMENTAL_changes", &params)
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	return raw, response.GetObject(raw)
}

func (a *NearApi) ViewContractCodeChangesByBlockId(accountIds []string, blockId uint64) (raw *types.ContractCodeChangesView, err error) {
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
	response, err := a.c.Call("EXPERIMENTAL_changes", &params)
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	return raw, response.GetObject(raw)
}

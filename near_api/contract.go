package near_api

import (
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
		RequestType string `json:"request_type"`
		Finality    string `json:"finality"`
		AccountID   string `json:"account_id"`
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

func (a *NearApi) CallContractFunc(account, method_name, args_base64 string) (*types.ContractFuncResult, error) {
	type Params struct {
		RequestType string `json:"request_type"`
		Finality    string `json:"finality"`
		AccountID   string `json:"account_id"`
		MethodName  string `json:"method_name"`
		ArgsBase64  string `json:"args_base64"`
	}
	params := Params{"call_function", "final", account, method_name, args_base64}
	response, err := a.c.Call("query", &params)
	if err != nil {
		return nil, err
	}
	var raw types.ContractFuncResult
	err = response.GetObject(&raw)
	if err != nil {
		return nil, err
	}
	return &raw, nil
}

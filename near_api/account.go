package near_api

import (
	types "github.com/nexeranet/gonear/near_api/types"
)

func (a *NearApi) ViewAccount(accountId string) (*types.AccountView, error) {
	type Params struct {
		RequestType string `json:"request_type"`
		Finality    string `json:"finality"`
		AccountID   string `json:"account_id"`
	}
	params := Params{"view_account", "final", accountId}
	response, err := a.c.Call("query", &params)
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	var raw types.AccountView
	return &raw, response.GetObject(&raw)
}

func (a *NearApi) ViewAccountByBlockId(accountId string, blockId uint64) (*types.AccountView, error) {
	type Params struct {
		RequestType string `json:"request_type"`
		BlockId     uint64 `json:"block_id"`
		AccountID   string `json:"account_id"`
	}
	params := Params{"view_account", blockId, accountId}
	response, err := a.c.Call("query", &params)
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	var raw types.AccountView
	return &raw, response.GetObject(&raw)
}

func (a *NearApi) ViewAccountChanges(accountIds []string) (*types.AccountChangesView, error) {
	type Params struct {
		ChangesType string   `json:"changes_type"`
		Finality    string   `json:"finality"`
		AccountIds  []string `json:"account_ids"`
	}
	params := Params{"account_changes", "final", accountIds}
	response, err := a.c.Call("EXPERIMENTAL_changes", &params)
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	var raw types.AccountChangesView
	return &raw, response.GetObject(&raw)
}

func (a *NearApi) ViewAccountChangesByBlockId(accountIds []string, blockId uint64) (*types.AccountChangesView, error) {
	type Params struct {
		ChangesType string   `json:"changes_type"`
		AccountIds  []string `json:"account_ids"`
		BlockId     uint64   `json:"block_id"`
	}
	params := Params{"account_changes", accountIds, blockId}
	response, err := a.c.Call("EXPERIMENTAL_changes", &params)
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	var raw types.AccountChangesView
	return &raw, response.GetObject(&raw)
}

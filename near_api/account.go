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

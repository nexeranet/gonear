package near_api

import (
	"fmt"

	types "github.com/nexeranet/gonear/near_api/types"
)

func (a *NearApi) ViewAccessKey(account, publicKey string) (*types.AccessKeysView, error) {
	type Params struct {
		RequestType string `json:"request_type"`
		Finality    string `json:"finality"`
		AccountID   string `json:"account_id"`
		PublicKey   string `json:"public_key"`
	}
	params := Params{"view_access_key", "final", account, publicKey}
	response, err := a.c.Call("query", &params)
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	var raw types.AccessKeysView
	err = response.GetObject(&raw)
	if err != nil {
		return nil, err
	}
	if raw.Error != "" {
		return nil, fmt.Errorf(raw.Error)
	}
	return &raw, nil
}

func (a *NearApi) ViewAccessKeyList(account string) (*types.AccessKeysListViev, error){
	type Params struct {
		RequestType string `json:"request_type"`
		Finality    string `json:"finality"`
		AccountID   string `json:"account_id"`
	}
	params := Params{"view_access_key_list", "final", account}
	response, err := a.c.Call("query", &params)
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	var raw types.AccessKeysListViev
	return &raw, response.GetObject(&raw)
}

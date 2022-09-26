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

func (a *NearApi) ViewAccessKeyByBlockId(account, publicKey string, blockId uint64) (*types.AccessKeysView, error) {
	type Params struct {
		RequestType string `json:"request_type"`
		AccountID   string `json:"account_id"`
		PublicKey   string `json:"public_key"`
		BlockId     uint64 `json:"block_id"`
	}
	params := Params{"view_access_key", account, publicKey, blockId}
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

func (a *NearApi) ViewAccessKeyList(account string) (*types.AccessKeysListViev, error) {
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

func (a *NearApi) ViewAccessKeyChanges(accountId, publicKey string) (*types.AccessKeyChangesView, error) {
	type Keys struct {
		AccountID string `json:"account_id"`
		PublicKey string `json:"public_key"`
	}
	type Params struct {
		ChangesType string `json:"changes_type"`
		Finality    string `json:"finality"`
		Keys        Keys   `json:"keys"`
	}
	params := Params{
		ChangesType: "single_access_key_changes",
		Finality:    "final",
		Keys: Keys{
			AccountID: accountId,
			PublicKey: publicKey,
		},
	}
	response, err := a.c.Call("EXPERIMENTAL_changes", &params)
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	var raw types.AccessKeyChangesView
	return &raw, response.GetObject(&raw)
}

func (a *NearApi) ViewAccessKeyChangesByBlockId(accountId, publicKey string, blockId uint64) (*types.AccessKeyChangesView, error) {
	type Keys struct {
		AccountID string `json:"account_id"`
		PublicKey string `json:"public_key"`
	}
	type Params struct {
		ChangesType string `json:"changes_type"`
		BlockId     uint64 `json:"block_id"`
		Keys        Keys   `json:"keys"`
	}
	params := Params{
		ChangesType: "single_access_key_changes",
		BlockId:     blockId,
		Keys: Keys{
			AccountID: accountId,
			PublicKey: publicKey,
		},
	}
	response, err := a.c.Call("EXPERIMENTAL_changes", &params)
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	var raw types.AccessKeyChangesView
	return &raw, response.GetObject(&raw)
}

func (a *NearApi) ViewAllAccessKeyChanges(accountIds []string) (*types.AccessKeyChangesView, error) {
	type Params struct {
		ChangesType string   `json:"changes_type"`
		AccountIds  []string `json:"account_ids"`
		Finality    string   `json:"finality"`
	}
	params := Params{
		ChangesType: "all_access_key_changes",
		Finality:    "final",
		AccountIds:  accountIds,
	}
	response, err := a.c.Call("EXPERIMENTAL_changes", &params)
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	var raw types.AccessKeyChangesView
	return &raw, response.GetObject(&raw)
}

func (a *NearApi) ViewAllAccessKeyChangesByBlockId(accountIds []string, blockId uint64) (*types.AccessKeyChangesView, error) {
	type Params struct {
		ChangesType string   `json:"changes_type"`
		AccountIds  []string `json:"account_ids"`
		BlockId     uint64   `json:"block_id"`
	}
	params := Params{
		ChangesType: "all_access_key_changes",
		BlockId:     blockId,
		AccountIds:  accountIds,
	}
	response, err := a.c.Call("EXPERIMENTAL_changes", &params)
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	var raw types.AccessKeyChangesView
	return &raw, response.GetObject(&raw)
}

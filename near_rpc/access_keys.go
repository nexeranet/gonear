package near_rpc

import (
	"fmt"

	types "github.com/nexeranet/gonear/near_rpc/types"
)

// Returns information about a single access key for given account.
func (a *Request) ViewAccessKey(account, publicKey string) (*types.AccessKeysView, error) {
	type Params struct {
		RequestType string `json:"request_type"`
		Finality    string `json:"finality"`
		AccountID   string `json:"account_id"`
		PublicKey   string `json:"public_key"`
	}
	params := Params{"view_access_key", "final", account, publicKey}
	response, err := a.Call("query", &params)
	if err != nil {
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

func (a *NearRpc) ViewAccessKey(account, publicKey string) (*types.AccessKeysView, error) {
    return a.Request().ViewAccessKey(account, publicKey)
}

// Returns information about a single access key for given account and block id
func (a *Request) ViewAccessKeyByBlockId(account, publicKey string, blockId uint64) (*types.AccessKeysView, error) {
	type Params struct {
		RequestType string `json:"request_type"`
		AccountID   string `json:"account_id"`
		PublicKey   string `json:"public_key"`
		BlockId     uint64 `json:"block_id"`
	}
	params := Params{"view_access_key", account, publicKey, blockId}
	response, err := a.Call("query", &params)
	if err != nil {
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

func (a *NearRpc) ViewAccessKeyByBlockId(account, publicKey string, blockId uint64) (*types.AccessKeysView, error) {
    return a.Request().ViewAccessKeyByBlockId(account, publicKey, blockId)
}

// Access keys for a given account.
func (a *Request) ViewAccessKeyList(account string) (*types.AccessKeysListViev, error) {
	type Params struct {
		RequestType string `json:"request_type"`
		Finality    string `json:"finality"`
		AccountID   string `json:"account_id"`
	}
	params := Params{"view_access_key_list", "final", account}
	response, err := a.Call("query", &params)
	if err != nil {
		return nil, err
	}
	var raw types.AccessKeysListViev
	return &raw, response.GetObject(&raw)
}

func (a *NearRpc) ViewAccessKeyList(account string) (*types.AccessKeysListViev, error) {
    return a.Request().ViewAccessKeyList(account)
}

//Returns individual access key changes in a last block.
func (a *Request) ViewAccessKeyChanges(accountId, publicKey string) (*types.AccessKeyChangesView, error) {
	type Keys struct {
		AccountID string `json:"account_id"`
		PublicKey string `json:"public_key"`
	}
	type Params struct {
		ChangesType string `json:"changes_type"`
		Finality    string `json:"finality"`
		Keys        []Keys `json:"keys"`
	}
	params := Params{
		ChangesType: "single_access_key_changes",
		Finality:    "final",
		Keys: []Keys{
			{
				AccountID: accountId,
				PublicKey: publicKey,
			},
		},
	}
	response, err := a.Call("EXPERIMENTAL_changes", &params)
	if err != nil {
		return nil, err
	}
	var raw types.AccessKeyChangesView
	return &raw, response.GetObject(&raw)
}

func (a *NearRpc) ViewAccessKeyChanges(accountId, publicKey string) (*types.AccessKeyChangesView, error) {
    return a.Request().ViewAccessKeyChanges(accountId, publicKey)
}

//Returns individual access key changes in a specific block.
func (a *Request) ViewAccessKeyChangesByBlockId(accountId, publicKey string, blockId uint64) (*types.AccessKeyChangesView, error) {
	type Keys struct {
		AccountID string `json:"account_id"`
		PublicKey string `json:"public_key"`
	}
	type Params struct {
		ChangesType string `json:"changes_type"`
		BlockId     uint64 `json:"block_id"`
		Keys        []Keys `json:"keys"`
	}
	params := Params{
		ChangesType: "single_access_key_changes",
		BlockId:     blockId,
		Keys: []Keys{
			{
				AccountID: accountId,
				PublicKey: publicKey,
			},
		},
	}
	response, err := a.Call("EXPERIMENTAL_changes", &params)
	if err != nil {
		return nil, err
	}
	var raw types.AccessKeyChangesView
	return &raw, response.GetObject(&raw)
}

func (a *NearRpc) ViewAccessKeyChangesByBlockId(accountId, publicKey string, blockId uint64) (*types.AccessKeyChangesView, error) {
    return a.Request().ViewAccessKeyChangesByBlockId(accountId, publicKey, blockId)
}


// Returns changes to all access keys of a last block.
func (a *Request) ViewAllAccessKeyChanges(accountIds []string) (*types.AccessKeyChangesView, error) {
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
	response, err := a.Call("EXPERIMENTAL_changes", &params)
	if err != nil {
		return nil, err
	}
	var raw types.AccessKeyChangesView
	return &raw, response.GetObject(&raw)
}

func (a *NearRpc) ViewAllAccessKeyChanges(accountIds []string) (*types.AccessKeyChangesView, error) {
    return a.Request().ViewAllAccessKeyChanges(accountIds)
}
// Returns changes to all access keys of a specific block.
func (a *Request) ViewAllAccessKeyChangesByBlockId(accountIds []string, blockId uint64) (*types.AccessKeyChangesView, error) {
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
	response, err := a.Call("EXPERIMENTAL_changes", &params)
	if err != nil {
		return nil, err
	}
	var raw types.AccessKeyChangesView
	return &raw, response.GetObject(&raw)
}

func (a *NearRpc) ViewAllAccessKeyChangesByBlockId(accountIds []string, blockId uint64) (*types.AccessKeyChangesView, error) {
    return a.Request().ViewAllAccessKeyChangesByBlockId(accountIds, blockId)
}

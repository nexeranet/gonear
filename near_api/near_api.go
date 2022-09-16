package near_api

import (
	"math/big"

	"github.com/nexeranet/gonear/jsonrpc"
	"github.com/nexeranet/gonear/types"
)

type NearApi struct {
	c   jsonrpc.RPCClient
	url string
}

// type NearApiI interface {
// 	ViewAccessKey(account, publicKey string) (*types.Permission, string, uint64, error)
// 	BalanceAt(account string) (*big.Int, error)
// 	CheckTx(hash, sender string) (*types.Transaction, error)
// 	SendAsyncTx(signedTx string) (string, error)
// 	SendAwaitTx(signedTx string) (*types.Transaction, error)
// }

func New(url string) *NearApi {
    rpc := &NearApi{
        url: url,
    }
	rpc.c = jsonrpc.NewClient(rpc.url)
    return rpc
}

func (a *NearApi) checkError(err error, response *jsonrpc.RPCResponse) error {
	if err != nil {
		return err
	}
	if response.Error != nil {
		return response.Error
	}
	return nil
}

func (a *NearApi) GetLastBlock() (string, error) {
	type Params struct {
		Finality string `json:"finality"`
	}
	params := &Params{"final"}
	response, err := a.c.Call("status", params)
	if err := a.checkError(err, response); err != nil {
		return "", err
	}
	var raw types.Block
	err = response.GetObject(&raw)
	if err != nil {
		return "", err
	}
	return raw.Header.Hash, nil
}

func (a *NearApi) ChainID() (string, error) {
	type Status struct {
		ChainID string `json:"chain_id"`
	}
	response, err := a.c.Call("status", []string{})
	if err := a.checkError(err, response); err != nil {
		return "", err
	}
	var raw Status
	err = response.GetObject(&raw)
	if err != nil {
		return "", err
	}
	return raw.ChainID, nil
}

func (a *NearApi) BalanceAt(account string) (*big.Int, error) {
	type Params struct {
		RequestType string `json:"request_type"`
		Finality    string `json:"finality"`
		AccountID   string `json:"account_id"`
	}
	type Account struct {
		Amount string `json:"amount"`
		Locked string `json:"locked"`
	}
	params := Params{"view_account", "final", account}
	response, err := a.c.Call("query", &params)

	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	var raw Account
	err = response.GetObject(&raw)
	if err != nil {
		return nil, err
	}
	i := new(big.Int)
	i.SetString(raw.Amount, 10)
	return i, nil
}


func (a *NearApi) CallContractFunc(account, method_name, args_base64 string) ([]rune, error) {
	type Params struct {
		RequestType string `json:"request_type"`
		Finality    string `json:"finality"`
		AccountID   string `json:"account_id"`
		MethodName  string `json:"method_name"`
		ArgsBase64  string `json:"args_base64"`
	}
	type Result struct {
		Result      []rune `json:"result"`
		BlockHeight uint64 `json:"block_height"`
		BlockHash   string `json:"block_hash"`
	}
	params := Params{"call_function", "final", account, method_name, args_base64}
	response, err := a.c.Call("query", &params)
	if err != nil {
		return nil, err
	}
	var raw Result
	err = response.GetObject(&raw)
	if err != nil {
		return nil, err
	}
	return raw.Result, nil
}

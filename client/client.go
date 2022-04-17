package client

import (
	"fmt"
	"math/big"

	"github.com/nexeranet/gonear/jsonrpc"
	"github.com/nexeranet/gonear/types"
)

type IClient interface {
	BalanceAt(account string) (*big.Int, error)
	CheckTx(hash, sender string) (*types.Transaction, error)
	GetAccessKeys(account, publicKey string) (string, string, uint64, error)
	SendAsyncTx(signedTx string) (string, error)
	SendAwaitTx(signedTx string) (bool, string, error)
	SendTransferTx(amount *big.Int, key, publicKey, addrFrom, addrTo string) (string, error)
	SendCallFunctionTx(methodName string, args []byte, amount, gas *big.Int, key, publicKey, addrFrom, addrTo string) (string, error)
}

type Client struct {
	c   jsonrpc.RPCClient
	url string
}

func NewClient(url string) *Client {
	return &Client{url: url}
}

func (a *Client) Init() (*Client, error) {
	a.c = jsonrpc.NewClient(a.url)
	return a, nil
}

func (a *Client) checkError(err error, response *jsonrpc.RPCResponse) error {
	if err != nil {
		return err
	}
	if response.Error != nil {
		return response.Error
	}
	return nil
}

func (a *Client) CheckTx(hash, sender string) (*types.Transaction, error) {
	response, err := a.c.Call("tx", [2]string{hash, sender})
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	var raw types.Transaction
	err = response.GetObject(&raw)
	if err != nil {
		return nil, err
	}
	if raw.Status.IsError() {
		return nil, raw.Status.Failure.Error()
	}
	if raw.Status.IsSuccess() {
		return &raw, nil
	}
	return nil, types.ErrUnknown
}

func (a *Client) GetLastBlock() (string, error) {
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

func (a *Client) ChainID() (string, error) {
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

func (a *Client) BalanceAt(account string) (*big.Int, error) {
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

// signed tx in base64
func (a *Client) SendAsyncTx(signedTx string) (string, error) {
	response, err := a.c.Call("broadcast_tx_async", [1]string{signedTx})
	if err := a.checkError(err, response); err != nil {
		return "", err
	}
	return response.GetString()
}

func (a *Client) SendAwaitTx(signedTx string) (bool, string, error) {
	response, err := a.c.Call("broadcast_tx_commit", [1]string{signedTx})
	if err := a.checkError(err, response); err != nil {
		return false, "", err
	}
	var raw types.Transaction
	err = response.GetObject(&raw)
	if err != nil {
		return false, "", err
	}
	if raw.Status.IsError() {
		return false, "", raw.Status.Failure.Error()
	}
	if raw.Status.IsSuccess() {
		return true, raw.Transaction.Hash, nil
	}
	return false, "", types.ErrUnknown
}

func (a *Client) GetAccessKeys(account, publicKey string) (*types.Permission, string, uint64, error) {
	type Params struct {
		RequestType string `json:"request_type"`
		Finality    string `json:"finality"`
		AccountID   string `json:"account_id"`
		PublicKey   string `json:"public_key"`
	}
	params := Params{"view_access_key", "final", account, publicKey}
	response, err := a.c.Call("query", &params)
	if err := a.checkError(err, response); err != nil {
		return nil, "", 0, err
	}
	var raw types.AccessKeys
	err = response.GetObject(&raw)
	if err != nil {
		return nil, "", 0, err
	}
	if raw.Error != "" {
		return &raw.Permission, raw.BlockHash, raw.Nonce, fmt.Errorf(raw.Error)
	}
	return &raw.Permission, raw.BlockHash, raw.Nonce, nil
}

func (a *Client) CallContractFunc(account, method_name, args_base64 string) ([]rune, error) {
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

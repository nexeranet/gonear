package client

import (
	"fmt"
	"math/big"

	"github.com/mr-tron/base58"
	"github.com/nexeranet/gonear/client/types"
	"github.com/nexeranet/gonear/near_rpc"
	near_api_types "github.com/nexeranet/gonear/near_rpc/types"
)

type IClient interface {
	API() near_rpc.NearApiI
	BalanceAt(string) (*big.Int, error)
	CheckTx(hash, sender string) (*near_api_types.TxView, error)
	SendTransferTx(amount *big.Int, key, publicKey, addrFrom, addrTo string) (*near_api_types.TxView, error)
	SendFunctionCallTx(methodName string, args []byte, amount, gas *big.Int, key, publicKey, addrFrom, addrTo string) (*near_api_types.TxView, error)
	SendActionsTx(key, publicKey, addrFrom, addrTo string, actions []types.Action) (*near_api_types.TxView, error)
	AsyncSendActionsTx(key, publicKey, addrFrom, addrTo string, actions []types.Action) (string, error)
}

type Client struct {
	C near_rpc.NearApiI
}

func NewClient(url string) *Client {
	return &Client{near_rpc.New(url)}
}

func (a *Client) API() near_rpc.NearApiI {
	return a.C
}

func (a *Client) CheckTx(hash, sender string) (*near_api_types.TxView, error) {
	return a.C.CheckTx(hash, sender)
}

func (a *Client) BalanceAt(accountId string) (*big.Int, error) {
	acc, err := a.C.ViewAccount(accountId)
	if err != nil {
		return nil, err
	}
	i := new(big.Int)
	i.SetString(acc.Amount, 10)
	return i, nil
}

func (a *Client) validateAccess(account, publicKey string) (nonce uint64, blockHash [32]byte, err error) {
	access_key, err := a.C.ViewAccessKey(account, publicKey)
	if err != nil {
		return nonce, blockHash, err
	}
	// TODO: if this condition is needed or not?
	if !access_key.Permission.IsFullAccess() {
		return nonce, blockHash, fmt.Errorf("`Account %s does not have permission to send tokens using key: %s", account, publicKey)
	}
	nonce = access_key.Nonce + 1
	hash, err := base58.Decode(access_key.BlockHash)
	if err != nil {
		return nonce, blockHash, err
	}
	return nonce, *(*[32]byte)(hash), nil
}

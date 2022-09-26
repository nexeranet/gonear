package client

import (
	"math/big"

	"github.com/nexeranet/gonear/near_api"
	types "github.com/nexeranet/gonear/near_api/types"
)

type IClient interface {
	BalanceAt(string) (*big.Int, error)
	CheckTx(hash, sender string) (*types.TxView, error)
	SendTransferTx(amount *big.Int, key, publicKey, addrFrom, addrTo string) (*types.TxView, error)
	SendCallFunctionTx(methodName string, args []byte, amount, gas *big.Int, key, publicKey, addrFrom, addrTo string) (*types.TxView, error)
}

type Client struct {
	C near_api.NearApiI
}

func NewClient(url string) *Client {
	return &Client{near_api.New(url)}
}

func (a *Client) CheckTx(hash, sender string) (*types.TxView, error) {
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

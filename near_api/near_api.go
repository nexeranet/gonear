package near_api

import (
	"github.com/nexeranet/gonear/jsonrpc"
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

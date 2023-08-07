// Near client
package client

import (
	"math/big"

	"github.com/mr-tron/base58"
	"github.com/nexeranet/gonear/client/types"
	"github.com/nexeranet/gonear/near_rpc"
	near_api_types "github.com/nexeranet/gonear/near_rpc/types"
)

const MaxGas uint64 = 300000000000000

type IClient interface {
	Rpc() near_rpc.INearRpc
	BalanceAt(string) (*big.Int, error)
	BalanceAtByBlockId(string, uint64) (*big.Int, error)
	CheckTx(hash, sender string) (*near_api_types.TxView, error)
	TransferTx(amount *big.Int, key, publicKey, addrFrom, addrTo string) (*near_api_types.TxView, error)
	FunctionCallTx(methodName string, args []byte, amount *big.Int, gas uint64, key, publicKey, addrFrom, addrTo string) (*near_api_types.TxView, error)
	ActionsTx(key, publicKey, addrFrom, addrTo string, actions []types.Action) (*near_api_types.TxView, error)
	AsyncActionsTx(key, publicKey, addrFrom, addrTo string, actions []types.Action) (string, error)
	CallContractFunc(accountId, method_name, args_base64 string) (*near_api_types.ContractFuncResult, error)
	CallContractFuncByBlockId(accountId, method_name, args_base64 string, block_id uint64) (*near_api_types.ContractFuncResult, error)
}

type Client struct {
	C near_rpc.INearRpc
}

func NewClient(url string) *Client {
	return &Client{near_rpc.NewInterface(url)}
}

func NewIClient(url string) IClient {
	return NewClient(url)
}

func (a *Client) Rpc() near_rpc.INearRpc {
	return a.C
}

func (a *Client) CallContractFunc(accountId, method_name, args_base64 string) (*near_api_types.ContractFuncResult, error) {
	return a.C.CallContractFunc(accountId, method_name, args_base64)
}

func (a *Client) CallContractFuncByBlockId(accountId, method_name, args_base64 string, block_id uint64) (*near_api_types.ContractFuncResult, error) {
	return a.C.CallContractFuncByBlockId(accountId, method_name, args_base64, block_id)
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

func (a *Client) BalanceAtByBlockId(accountId string, block_id uint64) (*big.Int, error) {
	acc, err := a.C.ViewAccountByBlockId(accountId, block_id)
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
	nonce = access_key.Nonce + 1
	hash, err := base58.Decode(access_key.BlockHash)
	if err != nil {
		return nonce, blockHash, err
	}
	return nonce, *(*[32]byte)(hash), nil
}

package near_api

import (
	"github.com/nexeranet/gonear/jsonrpc"
	types "github.com/nexeranet/gonear/near_api/types"
)

type NearApi struct {
	c   jsonrpc.RPCClient
	url string
}

type NearApiI interface {
	ViewAccessKey(account, publicKey string) (*types.AccessKeysView, error)
	ViewAccessKeyList(account string) (*types.AccessKeysListViev, error)
	ViewAccount(accountId string) (*types.AccountView, error)
	Block() (*types.BlockView, error)
	BlockByNumber(number uint64) (*types.BlockView, error)
	BlockByHash(hash string) (*types.BlockView, error)
	ChunkDetailsByHash(hash string) (*types.ChunkDetailsView, error)
	ChunkDetailsByIds(blockId, shardId uint64) (*types.ChunkDetailsView, error)
	ViewContractCode(accountId string) (*types.ContractCodeView, error)
	ViewContractState(accountId, prefixBase64 string) (*types.ContractStateView, error)
	CallContractFunc(accountId, method_name, args_base64 string) (*types.ContractFuncResult, error)
	GasPriceByHeight(height uint64) (*types.GasPriceView, error)
	GasPriceByHash(hash string) (*types.GasPriceView, error)
	GasPrice() (*types.GasPriceView, error)
	NodeStatus() (*types.NodeStatusView, error)
	NetworkInfo() (*types.NetworkInfoView, error)
	ValidationStatus() (*types.ValidationStatusView, error)
	ValidationStatusById(blockNumber uint64) (*types.ValidationStatusView, error)
	ValidationStatusByHash(hash string) (*types.ValidationStatusView, error)
	CheckTx(hash, sender string) (*types.TxView, error)
	SendAsyncTx(signedTx string) (string, error)
	SendAwaitTx(signedTx string) (*types.TxView, error)
}

func New(url string) NearApiI {
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

func (a *NearApi) GetUrl() string {
    return a.url
}

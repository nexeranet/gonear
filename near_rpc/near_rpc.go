// Near rpc api client (only rpc methods)
//
// The RPC API allows you to communicate directly with the NEAR network.
package near_rpc

//go:generate mockgen -source near_rpc.go -destination mocks/near_rpc.go
import (
	"github.com/nexeranet/gonear/jsonrpc"
	types "github.com/nexeranet/gonear/near_rpc/types"
)

type NearApi struct {
	c   jsonrpc.RPCClient
	url string
}

type NearApiI interface {
	Call(method string, params ...interface{}) (*jsonrpc.RPCResponse, error)
	ViewAccessKey(account, publicKey string) (*types.AccessKeysView, error)
	ViewAccessKeyByBlockId(account, publicKey string, blockId uint64) (*types.AccessKeysView, error)
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
	GenesisConfig() (raw *types.GenesisConfigView, err error)
	ProtocolConfig() (raw *types.ProtocolConfigView, err error)
	ProtocolConfigByBlockId(blockId uint64) (raw *types.ProtocolConfigView, err error)
	ViewAccessKeyChanges(accountId, publicKey string) (*types.AccessKeyChangesView, error)
	ViewAccessKeyChangesByBlockId(accountId, publicKey string, blockId uint64) (*types.AccessKeyChangesView, error)
	ViewAllAccessKeyChanges(accountIds []string) (*types.AccessKeyChangesView, error)
	ViewAllAccessKeyChangesByBlockId(accountIds []string, blockId uint64) (*types.AccessKeyChangesView, error)
	ViewAccountByBlockId(accountId string, blockId uint64) (*types.AccountView, error)
	ViewAccountChanges(accountIds []string) (*types.AccountChangesView, error)
	ViewAccountChangesByBlockId(accountIds []string, blockId uint64) (*types.AccountChangesView, error)
	ChangesInBlock() (*types.BlockChangesView, error)
	ChangesInBlockByHash(hash string) (*types.BlockChangesView, error)
	ChangesInBlockById(id uint64) (*types.BlockChangesView, error)
	ViewContractCodeChanges(accountIds []string) (raw *types.ContractCodeChangesView, err error)
	ViewContractCodeChangesByBlockId(accountIds []string, blockId uint64) (raw *types.ContractCodeChangesView, err error)
	ViewContractStateChanges(accountIds []string, keyPrefixBase64 string) (*types.ContractStateChangesView, error)
	ViewContractStateChangesByBlockId(accountIds []string, keyPrefixBase64 string, blockId uint64) (*types.ContractStateChangesView, error)
	TxStatusWithReceipts(txHash, sender string) (*types.TxView, error)
	ReceiptbyId(receiptId string) (*types.ViewReceipt, error)
}

func New(url string) NearApiI {
	rpc := &NearApi{
		url: url,
	}
	rpc.c = jsonrpc.NewClient(rpc.url)
	return rpc
}

func (a *NearApi) Call(method string, params ...interface{}) (*jsonrpc.RPCResponse, error) {
	response, err := a.c.Call(method, params...)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, types.ConvertError(response.Error)
	}
	return response, nil
}

func (a *NearApi) GetUrl() string {
	return a.url
}

package client

import (
	"github.com/nexeranet/gonear/client/types"
	near_rpc_types "github.com/nexeranet/gonear/near_rpc/types"
)

func (a *Client) ActionsTx(key, publicKey, addrFrom, addrTo string, actions []types.Action) (*near_rpc_types.TxView, error) {
	nonce, blockHash, err := a.validateAccess(addrFrom, publicKey)
	if err != nil {
		return nil, err
	}
    data, err := GenerateActionsTransactionHash(addrFrom, addrTo, key, nonce, blockHash, actions)
    if err != nil {
        return nil, err
    }
	return a.C.SendAwaitTx(data)
}

func (a *Client) AsyncActionsTx(key, publicKey, addrFrom, addrTo string, actions []types.Action) (string, error) {
	nonce, blockHash, err := a.validateAccess(addrFrom, publicKey)
	if err != nil {
		return "", err
	}
    data, err := GenerateActionsTransactionHash(addrFrom, addrTo, key, nonce, blockHash, actions)
    if err != nil {
        return "", err
    }
	return a.C.SendAsyncTx(data)
}

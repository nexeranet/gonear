package client

import (
	"math/big"

	"github.com/nexeranet/gonear/client/types"
	near_api_types "github.com/nexeranet/gonear/near_api/types"
)

func (a *Client) SendFunctionCallTx(methodName string, args []byte, deposit *big.Int, gas uint64, key, publicKey, addrFrom, addrTo string) (*near_api_types.TxView, error) {
	nonce, blockHash, err := a.validateAccess(addrFrom, publicKey)
	if err != nil {
		return nil, err
	}
    data, err := GenerateActionsTransactionHash(addrFrom, addrTo, key, nonce, blockHash, []types.Action{
        types.FunctionCallAction(methodName, args, gas, *deposit),
    })
    if err != nil {
        return nil, err
    }
	return a.C.SendAwaitTx(data)
}

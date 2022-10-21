package near_rpc

import (
	types "github.com/nexeranet/gonear/near_rpc/types"
)

func (a *NearApi) CheckTx(hash, sender string) (*types.TxView, error) {
	response, err := a.c.Call("tx", [2]string{hash, sender})
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	var raw types.TxView
	err = response.GetObject(&raw)
	if err != nil {
		return nil, err
	}
	return &raw, raw.Status.CheckError()
}

func (a *NearApi) SendAsyncTx(signedTx string) (string, error) {
	response, err := a.c.Call("broadcast_tx_async", [1]string{signedTx})
	if err := a.checkError(err, response); err != nil {
		return "", err
	}
	return response.GetString()
}

func (a *NearApi) SendAwaitTx(signedTx string) (*types.TxView, error) {
	response, err := a.c.Call("broadcast_tx_commit", [1]string{signedTx})
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	var raw types.TxView
	err = response.GetObject(&raw)
	if err != nil {
		return nil, err
	}
	return &raw, raw.Status.CheckError()
}

func (a *NearApi) TxStatusWithReceipts(txHash, sender string) (*types.TxView, error) {
	response, err := a.c.Call("EXPERIMENTAL_tx_status", [2]string{txHash, sender})
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	var raw types.TxView
	err = response.GetObject(&raw)
	if err != nil {
		return nil, err
	}
	return &raw, raw.Status.CheckError()
}

func (a *NearApi) ReceiptbyId(receiptId string) (*types.ViewReceipt, error) {
	type Params struct {
		ReceiptId string `json:"receipt_id"`
	}
	params := &Params{receiptId}
	response, err := a.c.Call("EXPERIMENTAL_receipt", params)
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	var raw types.ViewReceipt
	return &raw, response.GetObject(&raw)
}

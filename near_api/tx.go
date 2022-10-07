package near_api

import (

	types "github.com/nexeranet/gonear/near_api/types"
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
	if raw.Status.IsError() {
		return nil, raw.Status.Failure.Error()
	}
	if raw.Status.IsSuccess() {
		return &raw, nil
	}
	return nil, types.ErrUnknown
}

// signed tx in base64
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
	if raw.Status.IsError() {
		return &raw, raw.Status.Failure.Error()
	}
	return &raw, nil
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
	if raw.Status.IsError() {
		return nil, raw.Status.Failure.Error()
	}
	if raw.Status.IsSuccess() {
		return &raw, nil
	}
	return nil, types.ErrUnknown
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

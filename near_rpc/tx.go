package near_rpc

import (
	types "github.com/nexeranet/gonear/near_rpc/types"
)

// Queries status of a transaction by hash with sender account and returns the
// final transaction result.
func (a *Request) CheckTx(hash, sender string) (*types.TxView, error) {
	response, err := a.Call("tx", [2]string{hash, sender})
	if err != nil {
		return nil, err
	}
	var raw types.TxView
	return &raw, response.GetObject(&raw)
}

func (a *NearRpc) CheckTx(hash, sender string) (*types.TxView, error) {
	return a.Request().CheckTx(hash, sender)
}

// Sends a transaction and immediately returns transaction hash.
func (a *Request) SendAsyncTx(signedTx string) (string, error) {
	response, err := a.Call("broadcast_tx_async", [1]string{signedTx})
	if err != nil {
		return "", err
	}
	return response.GetString()
}

func (a *NearRpc) SendAsyncTx(signedTx string) (string, error) {
	return a.Request().SendAsyncTx(signedTx)
}

// Sends a transaction and waits until transaction is fully complete.
func (a *Request) SendAwaitTx(signedTx string) (*types.TxView, error) {
	response, err := a.Call("broadcast_tx_commit", [1]string{signedTx})
	if err != nil {
		return nil, err
	}
	var raw types.TxView
	return &raw, response.GetObject(&raw)
}

func (a *NearRpc) SendAwaitTx(signedTx string) (*types.TxView, error) {
	return a.Request().SendAwaitTx(signedTx)
}

// Queries status of a transaction by hash, returning the final transaction result
// and details of all receipts.
func (a *Request) TxStatusWithReceipts(txHash, sender string) (*types.TxView, error) {
	response, err := a.Call("EXPERIMENTAL_tx_status", [2]string{txHash, sender})
	if err != nil {
		return nil, err
	}
	var raw types.TxView
	return &raw, response.GetObject(&raw)
}

func (a *NearRpc) TxStatusWithReceipts(txHash, sender string) (*types.TxView, error) {
	return a.Request().TxStatusWithReceipts(txHash, sender)
}

// Fetches a receipt by it's ID (as is, without a status or execution outcome)
func (a *Request) ReceiptbyId(receiptId string) (*types.ViewReceipt, error) {
	type Params struct {
		ReceiptId string `json:"receipt_id"`
	}
	params := &Params{receiptId}
	response, err := a.Call("EXPERIMENTAL_receipt", params)
	if err != nil {
		return nil, err
	}
	var raw types.ViewReceipt
	return &raw, response.GetObject(&raw)
}

func (a *NearRpc) ReceiptbyId(receiptId string) (*types.ViewReceipt, error) {
	return a.Request().ReceiptbyId(receiptId)
}

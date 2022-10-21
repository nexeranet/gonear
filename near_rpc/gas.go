package near_rpc

import (
    types "github.com/nexeranet/gonear/near_rpc/types"
)

func (a *NearApi) GasPriceByHeight(height uint64) (*types.GasPriceView, error) {
	response, err := a.c.Call("gas_price", [1]uint64{height})
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	var raw types.GasPriceView
	return &raw, response.GetObject(&raw)
}

func (a *NearApi) GasPriceByHash(hash string) (*types.GasPriceView, error) {
	response, err := a.c.Call("gas_price", [1]string{hash})
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	var raw types.GasPriceView
	return &raw, response.GetObject(&raw)
}

func (a *NearApi) GasPrice() (*types.GasPriceView, error) {
	response, err := a.c.Call("gas_price", []interface{}{nil})
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	var raw types.GasPriceView
	return &raw, response.GetObject(&raw)
}

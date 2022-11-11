package near_rpc

import (
    types "github.com/nexeranet/gonear/near_rpc/types"
)

//Returns gas price for a specific block height.
func (a *NearRpc) GasPriceByHeight(height uint64) (*types.GasPriceView, error) {
	response, err := a.Call("gas_price", [1]uint64{height})
	if err != nil {
		return nil, err
	}
	var raw types.GasPriceView
	return &raw, response.GetObject(&raw)
}

//Returns gas price for a specific block hash
func (a *NearRpc) GasPriceByHash(hash string) (*types.GasPriceView, error) {
	response, err := a.Call("gas_price", [1]string{hash})
	if err != nil {
		return nil, err
	}
	var raw types.GasPriceView
	return &raw, response.GetObject(&raw)
}

// Returns gas price of a most recent block
func (a *NearRpc) GasPrice() (*types.GasPriceView, error) {
	response, err := a.Call("gas_price", []interface{}{nil})
	if err != nil {
		return nil, err
	}
	var raw types.GasPriceView
	return &raw, response.GetObject(&raw)
}

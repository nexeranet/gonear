package near_rpc

import (
	types "github.com/nexeranet/gonear/near_rpc/types"
)

func (a *NearApi) NodeStatus() (*types.NodeStatusView, error) {
	response, err := a.Call("status", []interface{}{})
	if err != nil {
		return nil, err
	}
	var raw types.NodeStatusView
	return &raw, response.GetObject(&raw)
}

func (a *NearApi) NetworkInfo() (*types.NetworkInfoView, error) {
	response, err := a.Call("network_info", []interface{}{})
	if err != nil {
		return nil, err
	}
	var raw types.NetworkInfoView
	return &raw, response.GetObject(&raw)
}

func (a *NearApi) ValidationStatusById(blockNumber uint64) (*types.ValidationStatusView, error) {
	response, err := a.Call("validators", []uint64{blockNumber})
	if err != nil {
		return nil, err
	}
	var raw types.ValidationStatusView
	return &raw, response.GetObject(&raw)
}

func (a *NearApi) ValidationStatusByHash(hash string) (*types.ValidationStatusView, error) {
	response, err := a.Call("validators", []string{hash})
	if err != nil {
		return nil, err
	}
	var raw types.ValidationStatusView
	return &raw, response.GetObject(&raw)
}

func (a *NearApi) ValidationStatus() (*types.ValidationStatusView, error) {
	response, err := a.Call("validators", []interface{}{nil})
	if err != nil {
		return nil, err
	}
	var raw types.ValidationStatusView
	return &raw, response.GetObject(&raw)
}

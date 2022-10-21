package near_rpc

import (
	types "github.com/nexeranet/gonear/near_rpc/types"
)

func (a *NearApi) NodeStatus() (*types.NodeStatusView, error) {
	response, err := a.c.Call("status", []interface{}{})
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	var raw types.NodeStatusView
	return &raw, response.GetObject(&raw)
}

func (a *NearApi) NetworkInfo() (*types.NetworkInfoView, error) {
	response, err := a.c.Call("network_info", []interface{}{})
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	var raw types.NetworkInfoView
	return &raw, response.GetObject(&raw)
}

func (a *NearApi) ValidationStatusById (blockNumber uint64) (*types.ValidationStatusView, error) {
	response, err := a.c.Call("validators", []uint64{blockNumber})
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	var raw types.ValidationStatusView
	return &raw, response.GetObject(&raw)
}

func (a *NearApi) ValidationStatusByHash (hash string) (*types.ValidationStatusView, error) {
	response, err := a.c.Call("validators", []string{hash})
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	var raw types.ValidationStatusView
	return &raw, response.GetObject(&raw)
}

func (a *NearApi) ValidationStatus () (*types.ValidationStatusView, error) {
	response, err := a.c.Call("validators", []interface{}{nil})
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	var raw types.ValidationStatusView
	return &raw, response.GetObject(&raw)
}

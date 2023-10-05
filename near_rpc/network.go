package near_rpc

import (
	types "github.com/nexeranet/gonear/near_rpc/types"
)

// Returns general status of a given node (sync status, nearcore node version,
// protocol version, etc), and the current set of validators.
func (a *Request) NodeStatus() (*types.NodeStatusView, error) {
	response, err := a.Call("status", []interface{}{})
	if err != nil {
		return nil, err
	}
	var raw types.NodeStatusView
	return &raw, response.GetObject(&raw)
}

func (a *NearRpc) NodeStatus() (*types.NodeStatusView, error) {
	return a.Request().NodeStatus()
}

// Returns the current state of node network connections (active peers,
// transmitted data, etc.)
func (a *Request) NetworkInfo() (*types.NetworkInfoView, error) {
	response, err := a.Call("network_info", []interface{}{})
	if err != nil {
		return nil, err
	}
	var raw types.NetworkInfoView
	return &raw, response.GetObject(&raw)
}

func (a *NearRpc) NetworkInfo() (*types.NetworkInfoView, error) {
	return a.Request().NetworkInfo()
}

// Queries active validators on the network returning details and the state of
// validation on the blockchain.
func (a *Request) ValidationStatusById(blockNumber uint64) (*types.ValidationStatusView, error) {
	response, err := a.Call("validators", []uint64{blockNumber})
	if err != nil {
		return nil, err
	}
	var raw types.ValidationStatusView
	return &raw, response.GetObject(&raw)
}

func (a *NearRpc) ValidationStatusById(blockNumber uint64) (*types.ValidationStatusView, error) {
	return a.Request().ValidationStatusById(blockNumber)
}

// Queries active validators on the network returning details and the state of
// validation on the blockchain.
func (a *Request) ValidationStatusByHash(hash string) (*types.ValidationStatusView, error) {
	response, err := a.Call("validators", []string{hash})
	if err != nil {
		return nil, err
	}
	var raw types.ValidationStatusView
	return &raw, response.GetObject(&raw)
}

func (a *NearRpc) ValidationStatusByHash(hash string) (*types.ValidationStatusView, error) {
	return a.Request().ValidationStatusByHash(hash)
}

// Queries active validators on the network returning details and the state of
// validation on the blockchain.
func (a *Request) ValidationStatus() (*types.ValidationStatusView, error) {
	response, err := a.Call("validators", []interface{}{nil})
	if err != nil {
		return nil, err
	}
	var raw types.ValidationStatusView
	return &raw, response.GetObject(&raw)
}

func (a *NearRpc) ValidationStatus() (*types.ValidationStatusView, error) {
	return a.Request().ValidationStatus()
}

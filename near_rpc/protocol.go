package near_rpc

import (
	types "github.com/nexeranet/gonear/near_rpc/types"
)

//Returns current genesis configuration.
func (a *Request) GenesisConfig() (*types.GenesisConfigView, error) {
	response, err := a.Call("EXPERIMENTAL_genesis_config")
	if err != nil {
		return nil, err
	}
	var raw types.GenesisConfigView
	return &raw, response.GetObject(&raw)
}

func (a *NearRpc) GenesisConfig() (*types.GenesisConfigView, error) {
    return a.Request().GenesisConfig()
}

//Returns most recent protocol configuration last block.
//Useful for finding current storage and transaction costs.
func (a *Request) ProtocolConfig() (*types.ProtocolConfigView, error) {
	type Params struct {
		Finality string `json:"finality"`
	}
	params := &Params{"final"}
	response, err := a.Call("EXPERIMENTAL_protocol_config", params)
	if err != nil {
		return nil, err
	}
	var raw types.ProtocolConfigView
	return &raw, response.GetObject(&raw)
}

func (a *NearRpc) ProtocolConfig() (*types.ProtocolConfigView, error) {
    return a.Request().ProtocolConfig()
}

//Returns most recent protocol configuration of a specific queried block.
//Useful for finding current storage and transaction costs.
func (a *Request) ProtocolConfigByBlockId(blockId uint64) (*types.ProtocolConfigView, error) {
	type Params struct {
		BlockId uint64 `json:"block_id"`
	}
	params := &Params{blockId}
	response, err := a.Call("EXPERIMENTAL_protocol_config", params)
	if err != nil {
		return nil, err
	}
	var raw types.ProtocolConfigView
	return &raw, response.GetObject(&raw)
}

func (a *NearRpc) ProtocolConfigByBlockId(blockId uint64) (*types.ProtocolConfigView, error) {
    return a.Request().ProtocolConfigByBlockId(blockId)
}

package near_rpc

import (
	types "github.com/nexeranet/gonear/near_rpc/types"
)

func (a *NearApi) GenesisConfig() (*types.GenesisConfigView, error) {
	response, err := a.Call("EXPERIMENTAL_genesis_config")
	if err != nil {
		return nil, err
	}
	var raw types.GenesisConfigView
	return &raw, response.GetObject(&raw)
}

func (a *NearApi) ProtocolConfig() (*types.ProtocolConfigView, error) {
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

func (a *NearApi) ProtocolConfigByBlockId(blockId uint64) (*types.ProtocolConfigView, error) {
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

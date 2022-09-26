package near_api

import (
    types "github.com/nexeranet/gonear/near_api/types"
)

func (a *NearApi) GenesisConfig() (raw *types.GenesisConfigView, err error) {
	response, err := a.c.Call("EXPERIMENTAL_genesis_config")
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	return raw, response.GetObject(raw)
}

func (a *NearApi) ProtocolConfig() (raw *types.ProtocolConfigView, err error) {
	type Params struct {
		Finality string `json:"finality"`
	}
	params := &Params{"final"}
	response, err := a.c.Call("EXPERIMENTAL_protocol_config", params)
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	return raw, response.GetObject(raw)
}

func (a *NearApi) ProtocolConfigByBlockId(blockId uint64) (raw *types.ProtocolConfigView, err error) {
	type Params struct {
		BlockId uint64 `json:"block_id"`
	}
	params := &Params{blockId}
	response, err := a.c.Call("EXPERIMENTAL_protocol_config", params)
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	return raw, response.GetObject(raw)
}

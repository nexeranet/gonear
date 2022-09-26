package near_api

import (
    types "github.com/nexeranet/gonear/near_api/types"
)


func (a *NearApi) Block() (*types.BlockView, error) {
	type Params struct {
		Finality string `json:"finality"`
	}
	params := &Params{"final"}
	response, err := a.c.Call("block", params)
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	var raw types.BlockView
	return &raw,  response.GetObject(&raw)
}

func (a *NearApi) BlockByNumber(number uint64) (*types.BlockView, error) {
	type Params struct {
		BlockId uint64 `json:"block_id"`
	}
	params := &Params{number}
	response, err := a.c.Call("block", params)
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	var raw types.BlockView
	return &raw,  response.GetObject(&raw)
}

func (a *NearApi) BlockByHash(hash string) (*types.BlockView, error) {
	type Params struct {
		BlockId string `json:"block_id"`
	}
	params := &Params{hash}
	response, err := a.c.Call("block", params)
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	var raw types.BlockView
	return &raw,  response.GetObject(&raw)
}

func (a *NearApi) ChangesInBlock(hash string) (*types.BlockChangesView, error) {
	type Params struct {
		Finality string `json:"finality"`
	}
	params := &Params{"final"}
	response, err := a.c.Call("EXPERIMENTAL_changes_in_block", params)
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	var raw types.BlockChangesView
	return &raw,  response.GetObject(&raw)
}

func (a *NearApi) ChangesInBlockByHash(hash string) (*types.BlockChangesView, error) {
	type Params struct {
		BlockId string `json:"block_id"`
	}
	params := &Params{hash}
	response, err := a.c.Call("EXPERIMENTAL_changes_in_block", params)
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	var raw types.BlockChangesView
	return &raw,  response.GetObject(&raw)
}

func (a *NearApi) ChangesInBlockById(id uint64) (*types.BlockChangesView, error) {
	type Params struct {
		BlockId uint64 `json:"block_id"`
	}
	params := &Params{id}
	response, err := a.c.Call("EXPERIMENTAL_changes_in_block", params)
	if err := a.checkError(err, response); err != nil {
		return nil, err
	}
	var raw types.BlockChangesView
	return &raw,  response.GetObject(&raw)
}

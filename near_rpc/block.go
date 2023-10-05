package near_rpc

import (
	types "github.com/nexeranet/gonear/near_rpc/types"
)

// Queries network and returns latest block details
func (a *Request) Block() (*types.BlockView, error) {
	type Params struct {
		Finality string `json:"finality"`
	}
	params := &Params{"final"}
	response, err := a.Call("block", params)
	if err != nil {
		return nil, err
	}
	var raw types.BlockView
	return &raw, response.GetObject(&raw)
}

func (a *NearRpc) Block() (*types.BlockView, error) {
	return a.Request().Block()
}

// Queries network and returns block for given height.
func (a *Request) BlockByNumber(number uint64) (*types.BlockView, error) {
	type Params struct {
		BlockId uint64 `json:"block_id"`
	}
	params := &Params{number}
	response, err := a.Call("block", params)
	if err != nil {
		return nil, err
	}
	var raw types.BlockView
	return &raw, response.GetObject(&raw)
}

func (a *NearRpc) BlockByNumber(number uint64) (*types.BlockView, error) {
	return a.Request().BlockByNumber(number)
}

// Queries network and returns block for given hash.
func (a *Request) BlockByHash(hash string) (*types.BlockView, error) {
	type Params struct {
		BlockId string `json:"block_id"`
	}
	params := &Params{hash}
	response, err := a.Call("block", params)
	if err != nil {
		return nil, err
	}
	var raw types.BlockView
	return &raw, response.GetObject(&raw)
}

func (a *NearRpc) BlockByHash(hash string) (*types.BlockView, error) {
	return a.Request().BlockByHash(hash)
}

// Returns changes in block for given latest block details.
func (a *Request) ChangesInBlock() (*types.BlockChangesView, error) {
	type Params struct {
		Finality string `json:"finality"`
	}
	params := &Params{"final"}
	response, err := a.Call("EXPERIMENTAL_changes_in_block", params)
	if err != nil {
		return nil, err
	}
	var raw types.BlockChangesView
	return &raw, response.GetObject(&raw)
}

func (a *NearRpc) ChangesInBlock() (*types.BlockChangesView, error) {
	return a.Request().ChangesInBlock()
}

// Returns changes in block for given block hash.
func (a *Request) ChangesInBlockByHash(hash string) (*types.BlockChangesView, error) {
	type Params struct {
		BlockId string `json:"block_id"`
	}
	params := &Params{hash}
	response, err := a.Call("EXPERIMENTAL_changes_in_block", params)
	if err != nil {
		return nil, err
	}
	var raw types.BlockChangesView
	return &raw, response.GetObject(&raw)
}

func (a *NearRpc) ChangesInBlockByHash(hash string) (*types.BlockChangesView, error) {
	return a.Request().ChangesInBlockByHash(hash)
}

// Returns changes in block for given block height.
func (a *Request) ChangesInBlockById(id uint64) (*types.BlockChangesView, error) {
	type Params struct {
		BlockId uint64 `json:"block_id"`
	}
	params := &Params{id}
	response, err := a.Call("EXPERIMENTAL_changes_in_block", params)
	if err != nil {
		return nil, err
	}
	var raw types.BlockChangesView
	return &raw, response.GetObject(&raw)
}

func (a *NearRpc) ChangesInBlockById(id uint64) (*types.BlockChangesView, error) {
	return a.Request().ChangesInBlockById(id)
}

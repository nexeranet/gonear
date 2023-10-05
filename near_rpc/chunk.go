package near_rpc

import (
	types "github.com/nexeranet/gonear/near_rpc/types"
)

// Returns details of a specific chunk
func (a *Request) ChunkDetailsByHash(hash string)  (*types.ChunkDetailsView, error) {
	type Params struct {
		ChunkId string `json:"chunk_id"`
	}
	params := &Params{hash}
	response, err := a.Call("chunk", params)
	if err != nil {
		return nil, err
	}
	var raw types.ChunkDetailsView
	return &raw,  response.GetObject(&raw)
}

func (a *NearRpc) ChunkDetailsByHash(hash string)  (*types.ChunkDetailsView, error) {
    return a.Request().ChunkDetailsByHash(hash)
}

//Returns details of a specific chunk. You can run a block details query to get
//a valid chunk hash.
func (a *Request) ChunkDetailsByIds(blockId, shardId uint64)  (*types.ChunkDetailsView, error) {
	type Params struct {
        BlockID uint64 `json:"block_id"`
        ShardID uint64 `json:"shard_id"`
	}
	params := &Params{blockId, shardId}
	response, err := a.Call("chunk", params)
	if err != nil {
		return nil, err
	}
	var raw types.ChunkDetailsView
	return &raw,  response.GetObject(&raw)
}

func (a *NearRpc) ChunkDetailsByIds(blockId, shardId uint64)  (*types.ChunkDetailsView, error) {
    return a.Request().ChunkDetailsByIds(blockId, shardId)
}

package near_rpc

import (
	types "github.com/nexeranet/gonear/near_rpc/types"
)

func (a *NearApi) ChunkDetailsByHash(hash string)  (*types.ChunkDetailsView, error) {
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

func (a *NearApi) ChunkDetailsByIds(blockId, shardId uint64)  (*types.ChunkDetailsView, error) {
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

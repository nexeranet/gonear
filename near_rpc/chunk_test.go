package near_rpc

import (
	"reflect"
	"testing"

	types "github.com/nexeranet/gonear/near_rpc/types"
)

func TestChunkDetailsByHash(t *testing.T) {
	type Test struct {
		name    string
		hash    string
		errType error
	}

	api := initTesnetApi()
	tests := []Test{
		{
			name:    "Valid chunk hash",
			hash:    "7dXNhXe9KRREZFTpD94jFKmgfyLTAD3TBWEoFAX5JtYh",
			errType: &types.ErrorUnknownChunk{},
		},
		{
			name:    "Invalid chunk hash",
			hash:    "SSSSSS",
			errType: &types.ErrorParseError{},
		},
		{
			name:    "Chunk not found",
			hash:    "EBM2qg5cGr47EjMPtH88uvmXHDHqmWPzKaQadbWhdw22",
			errType: &types.ErrorUnknownChunk{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := api.ChunkDetailsByHash(tt.hash)
			if err != nil {
				expect := reflect.TypeOf(tt.errType)
				have := reflect.TypeOf(err)
				if have != expect {
					t.Fatalf("Unexpected error %s, have type: %s, expect type %#v",
						err,
						have.String(),
						expect,
					)
				}
			} else {
				if result == nil {
					t.Fatalf("Expect struct, not nil")
				}
			}
		})
	}
}

func TestChunkDetailsByIds(t *testing.T) {
	type Test struct {
		name    string
		blockId uint64
		shardId uint64
        errType error
	}

	api := initTesnetApi()
	tests := []Test{
		{
			name:    "Valid block id and shard id",
			shardId: 1,
			blockId: 100655760,
			errType: &types.ErrorUnknownBlock{},
		},
		{
			name:    "Invalid block id",
			blockId: 0,
			errType: &types.ErrorUnknownBlock{},
		},
		{
			name:    "Chunk not found",
			blockId: 58934027,
			shardId: 0,
			errType: &types.ErrorUnknownBlock{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := api.ChunkDetailsByIds(tt.blockId, tt.shardId)
			if err != nil {
				expect := reflect.TypeOf(tt.errType)
				have := reflect.TypeOf(err)
				if have != expect {
					t.Fatalf("Unexpected error %s, have type: %s, expect type %#v",
						err,
						have.String(),
						expect,
					)
				}
			} else {
				if result == nil {
					t.Fatalf("Expect struct, not nil")
				}
			}
		})
	}
}

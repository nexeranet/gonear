package near_api

import (
	"testing"
)

func TestChunkDetailsByHash(t *testing.T) {
	type Test struct {
		name    string
		hash    string
		isError bool
	}

	api := initTesnetApi()
	tests := []Test{
		{
			name:    "Valid chunk hash",
			hash:    "7dXNhXe9KRREZFTpD94jFKmgfyLTAD3TBWEoFAX5JtYh",
			isError: false,
		},
		{
			name:    "Invalid chunk hash",
			hash:    "SSSSSS",
			isError: true,
		},
		{
			name:    "Chunk not found",
			hash:    "EBM2qg5cGr47EjMPtH88uvmXHDHqmWPzKaQadbWhdw22",
			isError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := api.ChunkDetailsByHash(tt.hash)
			if err != nil && !tt.isError {
				t.Fatalf("expected not error, actual %s", err)
			}
			if err == nil && tt.isError {
				t.Fatalf("Expect error, have nil")
			}
			if !tt.isError && result == nil {
				t.Fatalf("Expect struct, not nil")
			}
			if result != nil && result.Header.ChunkHash != tt.hash {
				t.Fatalf("Chank hash is not equal, result: %s, argument %s", result.Header.ChunkHash, tt.hash)
			}
		})
	}
}

func TestChunkDetailsByIds(t *testing.T) {
	type Test struct {
		name    string
		blockId uint64
		shardId uint64
		isError bool
	}

	api := initTesnetApi()
	tests := []Test{
		{
			name:    "Valid block id and shard id",
			isError: false,
			shardId: 1,
			blockId: 100655760,
		},
		{
			name:    "Invalid block id",
			isError: true,
			blockId: 0,
		},
		{
			name:    "Chunk not found",
			isError: true,
			blockId: 58934027,
			shardId: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := api.ChunkDetailsByIds(tt.blockId, tt.shardId)
			if err != nil && !tt.isError {
				t.Fatalf("expected not error, actual %s", err)
			}
			if err == nil && tt.isError {
				t.Fatalf("Expect error, have nil")
			}
			if !tt.isError && result == nil {
				t.Fatalf("Expect struct, not nil")
			}
		})
	}
}

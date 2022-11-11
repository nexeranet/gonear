package near_rpc

import (
	"testing"
)

func TestGenesisConfig(t *testing.T) {
	api := initTesnetApi()
	result, err := api.GenesisConfig()
	if err != nil {
		t.Fatalf("Error %s", err)
	}
	if result == nil {
		t.Fatalf("Result view is nil")
	}
}

func TestProtocolConfig(t *testing.T) {
	api := initTesnetApi()
	result, err := api.ProtocolConfig()
	if err != nil {
		t.Fatalf("Error %s", err)
	}
	if result == nil {
		t.Fatalf("Result view is nil")
	}
}

func TestProtocolConfigByBlockId(t *testing.T) {
    type Test struct{
        name string
        blockId uint64
        isError bool
    }

	api := initTesnetApi()
    tests := []Test{
        {
            name:"Valid block number",
            blockId: 100655760,
            isError: false,
        },
        {
            name:"Invalid block number",
            blockId: 100,
            isError: true,
        },
        {
            name:"Block number not found",
            blockId: 17821130,
            isError: true,
        },
    }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
            result, err := api.ProtocolConfigByBlockId(tt.blockId)
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

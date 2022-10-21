package near_rpc

import (
	"testing"
)

func TestBlock(t *testing.T) {
	api := initTesnetApi()
	block, err := api.Block()
	if err != nil {
		t.Fatalf("Error %s", err)
	}
	if block == nil {
		t.Fatalf("Block view is nil")
	}
}

func TestBlockByNumber(t *testing.T) {
	type Test struct {
		name    string
		number  uint64
		isError bool
	}
	api := initTesnetApi()
    tests := []Test{
        {
            name:"Valid block number",
            number: 100655760,
            isError: false,
        },
        {
            name:"Invalid block number",
            number: 100,
            isError: true,
        },
        {
            name:"Block number not found",
            number: 17821130,
            isError: true,
        },
    }
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
            block, err := api.BlockByNumber(tt.number)
			if err != nil && !tt.isError {
				t.Fatalf("expected not error, actual %s", err)
			}
			if err == nil && tt.isError {
				t.Fatalf("Expect error, have nil")
			}
            if !tt.isError && block == nil {
				t.Fatalf("Expect struct, not nil")
            }
        })
    }
}

func TestBlockByHash(t *testing.T) {
	type Test struct {
		name    string
		hash  string
		isError bool
	}
	api := initTesnetApi()
    tests := []Test{
        {
            name:"Valid block hash",
            hash: "AVXswVKwfAsUAqfFY3feMgf7GanN6GwPYtgnPiifWQGS",
            isError: false,
        },
        {
            name:"Invalid block hash",
            hash: "SSSSSS",
            isError: true,
        },
        {
            name:"Block hash not found",
            hash: "7nsuuitwS7xcdGnD9JgrE22cRB2vf2VS4yh1N9S71F4d",
            isError: true,
        },
    }
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
            block, err := api.BlockByHash(tt.hash)
			if err != nil && !tt.isError {
				t.Fatalf("expected not error, actual %s", err)
			}
			if err == nil && tt.isError {
				t.Fatalf("Expect error, have nil")
			}
            if !tt.isError && block == nil {
				t.Fatalf("Expect struct, not nil")
            }
        })
    }
}

func TestChangesInBlock(t *testing.T) {
	api := initTesnetApi()
	block, err := api.ChangesInBlock()
	if err != nil {
		t.Fatalf("Error %s", err)
	}
	if block == nil {
		t.Fatalf("Block view is nil")
	}
}

func TestChangesInBlockByHash(t *testing.T) {
	type Test struct {
		name    string
		hash  string
		isError bool
	}
	api := initTesnetApi()
    tests := []Test{
        {
            name:"Valid hash block",
            hash: "8uC449X4YtJXaCVo2NfS2LMvEhv769joDXEo8uzLuJE4",
            isError: false,
        },
        {
            name:"Invalid hash",
            hash: "asdfasdfasdf$$$$$$",
            isError: true,
        },
        {
            name:"Block not found",
            hash: "8uC449X4YtJXaCVo2NfS2LMvEhv769joDXEo8uzLuW33",
            isError: true,
        },
    }
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
            block, err := api.ChangesInBlockByHash(tt.hash)
			if err != nil && !tt.isError {
				t.Fatalf("expected not error, actual %s", err)
			}
			if err == nil && tt.isError {
				t.Fatalf("Expect error, have nil")
			}
            if !tt.isError && block == nil {
				t.Fatalf("Expect struct, not nil")
            }
        })
    }
}


func TestChangesInBlockById(t *testing.T) {
	type Test struct {
		name    string
        id uint64
		isError bool
	}
	api := initTesnetApi()
    tests := []Test{
        {
            name:"Valid block",
            id: 102109027,
            isError: false,
        },
        {
            name:"Invalid id",
            id: 0,
            isError: true,
        },
        {
            name:"Block not found",
            id: 1,
            isError: true,
        },
    }
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
            block, err := api.ChangesInBlockById(tt.id)
			if err != nil && !tt.isError {
				t.Fatalf("expected not error, actual %s", err)
			}
			if err == nil && tt.isError {
				t.Fatalf("Expect error, have nil")
			}
            if !tt.isError && block == nil {
				t.Fatalf("Expect struct, not nil")
            }
        })
    }
}

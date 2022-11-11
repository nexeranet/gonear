package near_rpc

import (
	"reflect"
	"testing"

	types "github.com/nexeranet/gonear/near_rpc/types"
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
        errType error
	}
	api := initTesnetApi()
    tests := []Test{
        {
            name:"Valid block number",
            number: 100655760,
			errType: &types.ErrorUnknownBlock{},
        },
        {
            name:"Invalid block number",
            number: 100,
			errType: &types.ErrorUnknownBlock{},
        },
        {
            name:"Block number not found",
            number: 17821130,
			errType: &types.ErrorUnknownBlock{},
        },
    }
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
            result, err := api.BlockByNumber(tt.number)
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

func TestBlockByHash(t *testing.T) {
	type Test struct {
		name    string
		hash  string
        errType error
	}
	api := initTesnetApi()
    tests := []Test{
        {
            name:"Valid block hash",
            hash: "AVXswVKwfAsUAqfFY3feMgf7GanN6GwPYtgnPiifWQGS",
			errType: &types.ErrorUnknownBlock{},
        },
        {
            name:"Invalid block hash",
            hash: "SSSSSS",
			errType: &types.ErrorParseError{},
        },
        {
            name:"Block hash not found",
            hash: "7nsuuitwS7xcdGnD9JgrE22cRB2vf2VS4yh1N9S71F4d",
			errType: &types.ErrorUnknownBlock{},
        },
    }
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
            result, err := api.BlockByHash(tt.hash)
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
        errType error
	}
	api := initTesnetApi()
    tests := []Test{
        {
            name:"Valid hash block",
            hash: "8uC449X4YtJXaCVo2NfS2LMvEhv769joDXEo8uzLuJE4",
			errType: &types.ErrorUnknownBlock{},
        },
        {
            name:"Invalid hash",
            hash: "asdfasdfasdf$$$$$$",
			errType: &types.ErrorParseError{},
        },
        {
            name:"Block not found",
            hash: "8uC449X4YtJXaCVo2NfS2LMvEhv769joDXEo8uzLuW33",
			errType: &types.ErrorUnknownBlock{},
        },
    }
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
            result, err := api.ChangesInBlockByHash(tt.hash)
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


func TestChangesInBlockById(t *testing.T) {
	type Test struct {
		name    string
        id uint64
        errType error
	}
	api := initTesnetApi()
    tests := []Test{
        {
            name:"Valid block",
            id: 102109027,
			errType: &types.ErrorUnknownBlock{},
        },
        {
            name:"Invalid id",
            id: 0,
			errType: &types.ErrorUnknownBlock{},
        },
        {
            name:"Block not found",
            id: 1,
			errType: &types.ErrorUnknownBlock{},
        },
    }
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
            result, err := api.ChangesInBlockById(tt.id)
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

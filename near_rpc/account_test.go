package near_rpc

import (
	"reflect"
	"testing"

	types "github.com/nexeranet/gonear/near_rpc/types"
)

func TestViewAccount(t *testing.T) {
	type Test struct {
		name      string
		accountId string
		isError   bool
	}

	client := initTesnetApi()
	tests := []Test{
		{
			name:      "Valid accound id",
			accountId: "nexeranet.testnet",
		},
		{
			name:      "Contract accound id",
			accountId: "vfinal.token.sweat.testnet",
		},
		{
			name:      "invalid account id",
			accountId: "___",
			isError:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list, err := client.ViewAccount(tt.accountId)
			if err != nil && !tt.isError {
				t.Fatalf("expected not error, actual %s", err)
			}
			if err == nil && tt.isError {
				t.Fatalf("Expect error, have nil")
			}
			if !tt.isError && list == nil {
				t.Fatalf("Expect struct, not nil")
			}
		})
	}
}

func TestViewAccountByBlockId(t *testing.T) {
	type Test struct {
		name      string
		accountId string
		blockId   uint64
		errType   error
	}

	client := initTesnetApi()
	tests := []Test{
		{
			name:      "Valid accound id",
			accountId: "nexeranet.testnet",
			blockId:   102109027,
			errType:   &types.ErrorGarbageCollectedBlock{},
		},
		{
			name:      "Contract accound id",
			accountId: "vfinal.token.sweat.testnet",
			blockId:   102109027,
			errType:   &types.ErrorGarbageCollectedBlock{},
		},
		{
			name:      "invalid account id",
			accountId: "___",
			blockId:   102109027,
			errType:   &types.ErrorParseError{},
		},
		{
			name:      "invalid block id",
			accountId: "nexeranet.testnet",
			blockId:   0,
			errType:   &types.ErrorUnknownBlock{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := client.ViewAccountByBlockId(tt.accountId, tt.blockId)
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

func TestViewAccountChanges(t *testing.T) {
	type Test struct {
		name     string
		accounts []string
		isError  bool
	}

	client := initTesnetApi()
	tests := []Test{
		{
			name:     "Valid accound id",
			accounts: []string{"nexeranet.testnet", "vfinal.token.sweat.testnet"},
		},
		{
			name:     "Contract accound id",
			accounts: []string{"vfinal.token.sweat.testnet"},
		},
		{
			name:     "invalid account id",
			accounts: []string{"___"},
			isError:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list, err := client.ViewAccountChanges(tt.accounts)
			if err != nil && !tt.isError {
				t.Fatalf("expected not error, actual %s", err)
			}
			if err == nil && tt.isError {
				t.Fatalf("Expect error, have nil")
			}
			if !tt.isError && list == nil {
				t.Fatalf("Expect struct, not nil")
			}
		})
	}
}

func TestViewAccountChangesByBlockId(t *testing.T) {
	type Test struct {
		name     string
		accounts []string
		blockId  uint64
		errType  error
	}

	client := initTesnetApi()
	tests := []Test{
		{
			name:     "Valid accound id",
			accounts: []string{"nexeranet.testnet"},
			blockId:  102109027,
			errType:  &types.ErrorUnknownBlock{},
		},
		{
			name:     "Contract accound id",
			accounts: []string{"vfinal.token.sweat.testnet"},
			blockId:  102109027,
			errType:  &types.ErrorUnknownBlock{},
		},
		{
			name:     "invalid block id",
			accounts: []string{"nexeranet.testnet"},
			blockId:  0,
			errType:  &types.ErrorUnknownBlock{},
		},
		{
			name:     "invalid account id",
			accounts: []string{"___"},
			errType:  &types.ErrorParseError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := client.ViewAccountChangesByBlockId(tt.accounts, tt.blockId)
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

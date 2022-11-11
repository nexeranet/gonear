package near_rpc

import (
	"reflect"
	"testing"

	types "github.com/nexeranet/gonear/near_rpc/types"
)

func TestViewAccessKey(t *testing.T) {
	type Test struct {
		name    string
		account string
		pubKey  string
		isError bool
	}
	tests := []Test{
		{
			name:    "simple addr",
			account: "nexeranet.testnet",
			pubKey:  "ed25519:7phkB1HWhWETQ1WkErTUS58s1EjMr4F8JFYg9VTQDk3X",
		},
		{
			name:    "get contract access keys",
			account: "client.chainlink.testnet",
			pubKey:  "ed25519:H9k5eiU4xXS3M4z8HzKJSLaZdqGdGwBG49o7orNC4eZW",
		},
		{
			name:    "get contract access keys 2",
			account: "token.arhius.testnet",
			pubKey:  "ed25519:9f42REGgZBENqEFSoQkfMwyv2VChsR7Lpy1tvWmYS6mL",
		},
		{
			name:    "invalid account id",
			account: "asdfasdf",
			pubKey:  "ed25519:9f42REGgZBENqEFSoQkfMwyv2VChsR7Lpy1tvWmYS6mL",
			isError: true,
		},
		{
			name:    "invalid public key",
			account: "asdfasdf",
			pubKey:  "asdfasfdsadf",
			isError: true,
		},
	}
	client := initTesnetApi()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			view_key, err := client.ViewAccessKey(tt.account, tt.pubKey)
			if err != nil && !tt.isError {
				t.Fatalf("expected not error, actual %s", err)
			}
			if err == nil && tt.isError {
				t.Fatalf("Expect error, have nil")
			}
			if !tt.isError && view_key == nil {
				t.Fatalf("Expect struct, not nil")
			}
		})
	}
}

func TestViewAccessKeyByBlockId(t *testing.T) {
	type Test struct {
		name    string
		account string
		blockId uint64
		pubKey  string
		errType error
	}
	tests := []Test{
		{
			name:    "simple addr",
			account: "nexeranet.testnet",
			pubKey:  "ed25519:7phkB1HWhWETQ1WkErTUS58s1EjMr4F8JFYg9VTQDk3X",
			blockId: 101076582,
			errType: &types.ErrorGarbageCollectedBlock{},
		},
		{
			name:    "invalid account id",
			account: "asdfasdf",
			pubKey:  "ed25519:9f42REGgZBENqEFSoQkfMwyv2VChsR7Lpy1tvWmYS6mL",
			blockId: 0,
			errType: &types.ErrorUnknownBlock{},
		},
		{
			name:    "invalid public key",
			account: "asdfasdf",
			pubKey:  "asdfasfdsadf",
			blockId: 0,
			errType: &types.ErrorParseError{},
		},
	}
	client := initTesnetApi()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			view_key, err := client.ViewAccessKeyByBlockId(tt.account, tt.pubKey, tt.blockId)
			if err != nil {
				expect := reflect.TypeOf(tt.errType)
				have := reflect.TypeOf(err)
				if expect != have {
					t.Fatalf("Unexpected error %#v, have type: %s, expect type %#v",
						err,
						have.String(),
						expect,
					)
				}
			} else {
				if view_key == nil {
					t.Fatalf("Expect struct, not nil")
				}
			}
		})
	}
}

func TestViewAccessKeyList(t *testing.T) {
	type Test struct {
		name    string
		account string
		isError bool
	}

	client := initTesnetApi()
	tests := []Test{
		{
			name:    "Base case",
			account: "nexeranet.testnet",
			isError: false,
		},
		{
			name: "SORRY valid account id",
			//INFO: valid accound id wow
			account: "asdfasdf",
			isError: false,
		},
		{
			name:    "invalid account id",
			account: "adfsa___adsfasdf",
			isError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list, err := client.ViewAccessKeyList(tt.account)
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

func TestViewAccessKeyChanges(t *testing.T) {
	type Test struct {
		name    string
		account string
		pubKey  string
		isError bool
	}
	client := initTesnetApi()
	tests := []Test{
		{
			name:    "Simple test",
			account: "nexeranet.testnet",
			pubKey:  "ed25519:7phkB1HWhWETQ1WkErTUS58s1EjMr4F8JFYg9VTQDk3X",
			isError: false,
		},
		{
			name:    "Invalid account id",
			account: "asdfasdf$$$$$",
			pubKey:  "ed25519:7phkB1HWhWETQ1WkErTUS58s1EjMr4F8JFYg9VTQDk3X",
			isError: true,
		},
		{
			name:    "Invalid public key",
			account: "nexeranet.testnet",
			pubKey:  "WkErT11$$$US58s1EjMr4F8JFYg9VTQDk3X",
			isError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := client.ViewAccessKeyChanges(tt.account, tt.pubKey)
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

func TestViewAccessKeyChangesByBlockId(t *testing.T) {
	type Test struct {
		name    string
		account string
		blockId uint64
		pubKey  string
		errType error
	}
	client := initTesnetApi()
	tests := []Test{
		{
			name:    "Simple test",
			account: "nexeranet.testnet",
			pubKey:  "ed25519:7phkB1HWhWETQ1WkErTUS58s1EjMr4F8JFYg9VTQDk3X",
			blockId: 102109027,
			errType: &types.ErrorUnknownBlock{},
		},
		{
			name:    "Invalid account id",
			account: "asdfasdf$$$$$",
			pubKey:  "ed25519:7phkB1HWhWETQ1WkErTUS58s1EjMr4F8JFYg9VTQDk3X",
			blockId: 102109027,
			errType: &types.ErrorParseError{},
		},
		{
			name:    "Invalid public key",
			account: "nexeranet.testnet",
			pubKey:  "WkErT11$$$US58s1EjMr4F8JFYg9VTQDk3X",
			blockId: 102109027,
			errType: &types.ErrorParseError{},
		},
		{
			name:    "Invalid block id",
			account: "nexeranet.testnet",
			pubKey:  "WkErT11$$$US58s1EjMr4F8JFYg9VTQDk3X",
			blockId: 0,
			errType: &types.ErrorParseError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := client.ViewAccessKeyChangesByBlockId(tt.account, tt.pubKey, tt.blockId)
			if err != nil {
				expect := reflect.TypeOf(tt.errType)
				have := reflect.TypeOf(err)
				if have != expect {
					t.Fatalf("Unexpected error %#v, have type: %s, expect type %#v",
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

func TestViewAllAccessKeyChanges(t *testing.T) {
	type Test struct {
		name     string
		accounts []string
		isError  bool
	}
	client := initTesnetApi()
	tests := []Test{
		{
			name:     "Simple test",
			accounts: []string{"nexeranet.testnet"},
			isError:  false,
		},
		{
			name:     "Invalid account id",
			accounts: []string{"asdfasdf$$$$$"},
			isError:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := client.ViewAllAccessKeyChanges(tt.accounts)
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
func TestViewAllAccessKeyChangesByBlockId(t *testing.T) {
	type Test struct {
		name     string
		accounts []string
		blockId  uint64
		errType  error
	}
	client := initTesnetApi()
	tests := []Test{
		{
			name:     "Simple test",
			accounts: []string{"nexeranet.testnet"},
			blockId:  102109027,
			errType:  &types.ErrorUnknownBlock{},
		},
		{
			name:     "Invalid account id",
			accounts: []string{"asdfasdf$$$$$"},
			blockId:  102109027,
			errType:  &types.ErrorParseError{},
		},
		{
			name:     "Invalid block id",
			accounts: []string{"nexeranet.testnet"},
			blockId:  0,
			errType:  &types.ErrorUnknownBlock{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := client.ViewAllAccessKeyChangesByBlockId(tt.accounts, tt.blockId)
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

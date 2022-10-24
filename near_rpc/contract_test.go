package near_rpc

import (
	"fmt"
	"reflect"
	"testing"

	types "github.com/nexeranet/gonear/near_rpc/types"
)

func TestViewContractCode(t *testing.T) {
	type Test struct {
		name      string
		accountId string
		isError   bool
	}

	api := initTesnetApi()
	tests := []Test{
		{
			name:      "Contract address",
			accountId: "vfinal.token.sweat.testnet",
			isError:   false,
		},
		{
			name:      "Account id doesn't belong to the contract",
			accountId: "nexeranet.testnet",
			isError:   true,
		},
		{
			name:      "Invalid account id",
			accountId: "asdfas____",
			isError:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := api.ViewContractCode(tt.accountId)
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

func TestViewContractState(t *testing.T) {
	type Test struct {
		name      string
		accountId string
		isError   bool
	}

	api := initTesnetApi()
	tests := []Test{
		{
			name:      "Contract address",
			accountId: "vfinal.token.sweat.testnet",
			isError:   false,
		},
		{
			name:      "Account id doesn't belong to the contract",
			accountId: "nexeranet.testnet",
			isError:   true,
		},
		{
			name:      "Invalid account id",
			accountId: "asdfas____",
			isError:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := api.ViewContractState(tt.accountId, "")
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

func TestCallContractFunc(t *testing.T) {
	type Test struct {
		name       string
		accountId  string
		methodName string
		argsBase64 string
		isError    bool
	}

	api := initTesnetApi()
	tests := []Test{
		{
			name:       "Base case",
			accountId:  "dev-1588039999690",
			methodName: "get_num",
			argsBase64: "e30=",
			isError:    false,
		},
		{
			name:       "Undefined method",
			accountId:  "dev-1588039999690",
			methodName: "get_num_asdfas_fffff",
			argsBase64: "e30=",
			isError:    true,
		},
		{
			name:       "Account id doesn't belong to the contract",
			accountId:  "nexeranet.testnet",
			methodName: "get_num",
			argsBase64: "e30=",
			isError:    true,
		},
		{
			name:       "invalid args",
			accountId:  "nexeranet.testnet",
			methodName: "get_num",
			argsBase64: "e30=asdfasdfasdfsa",
			isError:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := api.CallContractFunc(tt.accountId, tt.methodName, tt.argsBase64)
			if err != nil && !tt.isError {
				t.Fatalf("expected not error, actual %s", err)
			}
			if err == nil && tt.isError {
				t.Fatalf("Expect error, have nil, %s", tt.name)
			}
			if !tt.isError && result == nil {
				t.Fatalf("Expect struct, not nil")
			}
		})
	}
}

func TestViewContractCodeChanges(t *testing.T) {
	type Test struct {
		name     string
		accounts []string
		isError  bool
	}

	api := initTesnetApi()
	tests := []Test{
		{
			name:     "Valid contract",
			accounts: []string{"dev-1588039999690"},
			isError:  false,
		},
		// INFO: without error, HMMMMMMMM
		{
			name:     "Invalid contract address, user address",
			accounts: []string{"nexeranet.testnet"},
			isError:  false,
		},
		{
			name:     "Invalid accounts id",
			accounts: []string{"___"},
			isError:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block, err := api.ViewContractCodeChanges(tt.accounts)
			fmt.Println(err)
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

func TestViewContractCodeChangesByBlockId(t *testing.T) {
	type Test struct {
		name     string
		accounts []string
		blockId  uint64
        errType error
	}
	api := initTesnetApi()
	tests := []Test{
		{
			name:     "Valid contract",
			accounts: []string{"dev-1588039999690"},
			blockId:  103143176,
			errType: &types.ErrorUnknownBlock{},
		},
		{
			name:     "Invalid contract address, user address",
			accounts: []string{"nexeranet.testnet"},
			blockId:  103143176,
			errType: &types.ErrorUnknownBlock{},
		},
		{
			name:     "Invalid accounts id",
			accounts: []string{"___"},
			errType: &types.ErrorParseError{},
		},
		{
			name:     "Invalid block id",
			accounts: []string{"dev-1588039999690"},
			blockId:  0,
			errType: &types.ErrorUnknownBlock{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := api.ViewContractCodeChangesByBlockId(tt.accounts, tt.blockId)
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

func TestViewContractStateChanges(t *testing.T) {
	type Test struct {
		name     string
		accounts []string
		args     string
		isError  bool
	}
	api := initTesnetApi()
	tests := []Test{
		{
			name:     "Valid contract",
			accounts: []string{"dev-1588039999690"},
			isError:  false,
		},
		// INFO: without error, HMMMMMMMM
		{
			name:     "Invalid contract address, user address",
			accounts: []string{"nexeranet.testnet"},
			isError:  false,
		},
		{
			name:     "Invalid accounts id",
			accounts: []string{"___"},
			isError:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := api.ViewContractStateChanges(tt.accounts, tt.args)
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

func TestViewContractStateChangesByBlockId(t *testing.T) {
	type Test struct {
		name     string
		accounts []string
		args     string
		blockId  uint64
        errType error
	}
	api := initTesnetApi()
	tests := []Test{
		{
			name:     "Valid contract",
			accounts: []string{"dev-1588039999690"},
			blockId:  102118895,
			errType: &types.ErrorUnknownBlock{},
		},
		{
			name:     "Invalid accounts id",
			accounts: []string{"___"},
			blockId:  102118895,
			errType: &types.ErrorParseError{},
		},
		{
			name:     "Invalid block id",
			accounts: []string{"dev-1588039999690"},
			blockId:  0,
			errType: &types.ErrorUnknownBlock{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := api.ViewContractStateChangesByBlockId(tt.accounts, tt.args, tt.blockId)
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

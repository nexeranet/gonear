package near_api

import (
	"testing"
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

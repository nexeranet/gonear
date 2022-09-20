package near_api

import (
	"testing"
)

func TestGasPriceByHeight(t *testing.T) {
	type Test struct {
		name    string
		height  uint64
		isError bool
	}
	api := initTesnetApi()
	tests := []Test{
		{
			name:    "Valid block height",
			height:  100655760,
			isError: false,
		},
		{
			name:    "Invalid block height",
			height:  0,
			isError: true,
		},
		{
			name:    "Block height not found",
			height:  17824600,
			isError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := api.GasPriceByHeight(tt.height)
			if err != nil && !tt.isError {
				t.Fatalf("Test: %s, expected not error, actual %s", tt.name, err)
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

func TestGasPriceByHash(t *testing.T) {
	type Test struct {
		name    string
		hash    string
		isError bool
	}
	api := initTesnetApi()
	tests := []Test{
		{
			name:    "Valid block hash",
			hash:    "AVXswVKwfAsUAqfFY3feMgf7GanN6GwPYtgnPiifWQGS",
			isError: false,
		},
		{
			name:    "Invalid block hash",
			hash:    "SSSSSS",
			isError: true,
		},
		{
			name:    "Block hash not found",
			hash:    "7nsuuitwS7xcdGnD9JgrE22cRB2vf2VS4yh1N9S71F4d",
			isError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := api.GasPriceByHash(tt.hash)
			if err != nil && !tt.isError {
				t.Fatalf("Test: %s, expected not error, actual %s", tt.name, err)
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

func TestGasPrice(t *testing.T) {
	api := initTesnetApi()
	block, err := api.GasPrice()
	if err != nil {
		t.Fatalf("Error %s", err)
	}
	if block == nil {
		t.Fatalf("Gas view is nil")
	}
}

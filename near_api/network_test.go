package near_api

import (
	"testing"
)

func TestNodeStatus(t *testing.T) {
	api := initTesnetApi()
	result, err := api.NodeStatus()
	if err != nil {
		t.Fatalf("Error %s", err)
	}
	if result == nil {
		t.Fatalf("Result view is nil")
	}
}

func TestNetworkInfo(t *testing.T) {
	api := initTesnetApi()
	result, err := api.NetworkInfo()
	if err != nil {
		t.Fatalf("Error %s", err)
	}
	if result == nil {
		t.Fatalf("Result view is nil")
	}
}


func TestValidationStatus(t *testing.T) {
	api := initTesnetApi()
	result, err := api.ValidationStatus()
	if err != nil {
		t.Fatalf("Error %s", err)
	}
	if result == nil {
		t.Fatalf("Result view is nil")
	}
}

func TestValidationStatusById(t *testing.T) {
	type Test struct {
		name    string
		number  uint64
		isError bool
	}
	api := initTesnetApi()
	tests := []Test{
		// {
		// 	name:    "Valid block number",
		// 	number:  100671417,
		// 	isError: false,
		// },
		{
			name:    "Invalid block number",
			number:  0,
			isError: true,
		},
		{
			name:    "Block number not found",
			number:  17824600,
			isError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := api.ValidationStatusById(tt.number)
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

func TestValidationStatusByHash(t *testing.T) {
	type Test struct {
		name    string
		hash    string
		isError bool
	}
	api := initTesnetApi()
	tests := []Test{
		// {
		// 	name:    "Valid block hash",
		// 	hash:    "AVXswVKwfAsUAqfFY3feMgf7GanN6GwPYtgnPiifWQGS",
		// 	isError: false,
		// },
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
			result, err := api.ValidationStatusByHash(tt.hash)
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

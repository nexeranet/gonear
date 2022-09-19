package near_api

import (
	"testing"
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
	client := initApi()
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

func TestViewAccessKeyList(t *testing.T) {
	type Test struct {
		name    string
		account string
		isError bool
	}

	client := initApi()
    tests := []Test{
		{
			name:     "Base case",
			account: "nexeranet.testnet",
			isError: false,
		},
		{
			name:    "SORRY valid account id",
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

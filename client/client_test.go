package client

import (
	"log"
	"testing"
)

func initTestClient(t *testing.T) *Client {
	Url := "https://rpc.testnet.near.org"
	return NewClient(Url)
}

func TestGetAccessKeys(t *testing.T) {
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
	}
	client := initTestClient(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			permission, blockHash, nonce, err := client.GetAccessKeys(tt.account, tt.pubKey)
			log.Println(permission, blockHash, nonce)
			if err != nil && !tt.isError {
				t.Fatalf("expected not error, actual %s", err)
			}
			if err == nil && tt.isError {
				t.Fatalf("Expect error, have nil")
			}
		})
	}
}

func TestClient__GetBalance(t *testing.T) {
	type Test struct {
		isError bool
		addr    string
		name    string
	}
	tests := []Test{
		{
			name:    "simple addr",
			addr:    "nearkat.testnet",
			isError: false,
		},
		{
			name:    "invalid addr",
			addr:    "c292",
			isError: true,
		},
	}
	client := initTestClient(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			balance, err := client.BalanceAt(tt.addr)
			if err != nil && !tt.isError {
				t.Fatalf("expected not error, actual %s", err)
			}
			if err == nil && tt.isError {
				t.Fatalf("Expect error, have nil")
			}
			if balance == nil && !tt.isError {
				t.Fatalf("Balance is nil")
			}
		})
	}
}

package client

import (
	"testing"
)

func initTestClient(t *testing.T) *Client {
	Url := "https://rpc.testnet.near.org"
	return NewClient(Url)
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

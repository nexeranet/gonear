package client

import (
	"math/big"
	"testing"

	"github.com/nexeranet/gonear/client/types"
)

func TestClient__SendTransferTx(t *testing.T) {
	type Test struct {
		name     string
		addrTo   string
		addrFrom string
		amount   *big.Int
		isError  bool
	}
	key := "ed25519:5XKLL4yQoBVyHCUyXrMt9898VG7My2iWomu1GC3wAW4V6eBwZGmreqpMiWfC1HiVpmAAWCe1pJ6RKNuEFgupbPjK"
	pubKey := "ed25519:7phkB1HWhWETQ1WkErTUS58s1EjMr4F8JFYg9VTQDk3X"
	tests := []Test{
		{
			name:     "Simple data",
			addrFrom: "nexeranet.testnet",
			addrTo:   "token.arhius.testnet",
			amount:   types.NewNear(1).BigInt(),
		},
	}
	client := initTestClient(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tx, err := client.SendTransferTx(tt.amount, key, pubKey, tt.addrFrom, tt.addrTo)

			if err != nil && !tt.isError {
				t.Fatalf("expected not error, actual %s", err)
			}

			if tx.Status.IsError() && !tt.isError {
				t.Fatalf("expected not error, actual %s", tx.Status.Failure.Error())
			}

			if err == nil && tt.isError {
				t.Fatalf("Expect error, have nil")
			}
		})
	}
}

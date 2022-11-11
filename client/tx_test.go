package client

import (
	"math/big"
	"os"
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
	key := os.Getenv("PRIVATE")
	pubKey := os.Getenv("PUBLIC")
	acc := os.Getenv("ACCOUNT")
	tests := []Test{
		{
			name:     "Simple data",
			addrFrom: acc,
			addrTo:   "token.arhius.testnet",
			amount:   types.NewNear(1).BigInt(),
		},
	}
	client := initTestClient(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := client.TransferTx(tt.amount, key, pubKey, tt.addrFrom, tt.addrTo)

			if err != nil && !tt.isError {
				t.Fatalf("expected not error, actual %s", err)
			}

			if err == nil && tt.isError {
				t.Fatalf("Expect error, have nil")
			}
		})
	}
}

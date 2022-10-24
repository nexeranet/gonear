package client

import (
	"encoding/json"
	"log"
	"math/big"
	"testing"

	"github.com/nexeranet/gonear/client/types"
)

func prettyPrint(i interface{}) string {
    s, _ := json.MarshalIndent(i, "", " ")
    return string(s)
}

func TestClient__SendCallFunctionTx(t *testing.T) {
	type Test struct {
		isError  bool
		addrTo   string
		addrFrom string
		gas      uint64
		name     string
	}
	type Args struct {
		ReceiverId string  `json:"receiver_id"`
		Amount     string  `json:"amount"`
		Memo       *string `json:"memo"`
	}
	key := "ed25519:5XKLL4yQoBVyHCUyXrMt9898VG7My2iWomu1GC3wAW4V6eBwZGmreqpMiWfC1HiVpmAAWCe1pJ6RKNuEFgupbPjK"
	pubKey := "ed25519:7phkB1HWhWETQ1WkErTUS58s1EjMr4F8JFYg9VTQDk3X"
	tests := []Test{
		{
			name:     "Simple data",
			addrFrom: "nexeranet.testnet",
			addrTo:   "token.arhius.testnet",
			gas:      300000000000000,
		},
	}
	args := Args{
		ReceiverId: "token.arhius.testnet",
		Amount:     types.NewNear(1).String(),
		Memo:       nil,
	}
	bytes, err := json.Marshal(&args)
	log.Println(args)
	if err != nil {
		t.Fatalf("JSON Marshal: %s", err.Error())
	}
	client := initTestClient(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tx, err := client.FunctionCallTx("ft_transfer", bytes, big.NewInt(1), tt.gas, key, pubKey, tt.addrFrom, tt.addrTo)
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

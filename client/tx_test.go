package client

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/nexeranet/gonear/types"
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
			fmt.Println(prettyPrint(tx))
			fmt.Println(err)

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

func TestClient__SendTransaction(t *testing.T) {
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
            // types.TransferAction | types.types.FunctionCallAction
			actions := types.Actions[types.TransferAction]{
				types.TransferAction{
					Enum: types.TransferEnum,
					Transfer: types.Transfer{
						Deposit: *tt.amount,
					},
				},
				// types.FunctionCallAction{
				// 	Enum: types.FunctionCallEnum,
				// 	FunctionCall: types.FunctionCall{
				// 		Deposit:    *big.NewInt(0),
				// 		MethodName: "ft_transfer",
				// 		Gas:        300000000000000,
				// 		Args:       []byte{},
				// 	},
				// },
			}
			tx, err := SendTransaction(client, key, pubKey, tt.addrFrom, tt.addrTo, actions)
			fmt.Println(prettyPrint(tx))
			fmt.Println(err)

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

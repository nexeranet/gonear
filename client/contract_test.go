package client

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"testing"

	"github.com/nexeranet/gonear/client/types"
)

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

	key := os.Getenv("PRIVATE")
	pubKey := os.Getenv("PUBLIC")
	acc := os.Getenv("ACCOUNT")
	tests := []Test{
		{
			name:     "Simple data",
			addrFrom: acc,
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

			_, err := client.FunctionCallTx("ft_transfer", bytes, big.NewInt(1), tt.gas, key, pubKey, tt.addrFrom, tt.addrTo)
			if err != nil && !tt.isError {
				t.Fatalf("expected not error, actual %s", err)
			}
			if err == nil && tt.isError {
				t.Fatalf("Expect error, have nil")
			}
		})
	}
}

func TestClient__CallContract(t *testing.T) {
	type Args struct {
		ProjectId string `json:"project_id"`
	}
	var gas uint64 = 300000000000000
	key := os.Getenv("PRIVATE")
	pubKey := os.Getenv("PUBLIC")
	addrFrom := os.Getenv("ACCOUNT")
	addrTo := "deploy.ofp_collateral.testnet"
	args := Args{
		ProjectId: "1666081062930",
	}
	bytes, err := json.Marshal(&args)
	if err != nil {
		t.Fatalf("JSON Marshal: %s", err.Error())
	}
	deposit := types.NewNear(2).BigInt()
	deposit = deposit.Div(deposit, big.NewInt(10))
	client := initTestClient(t)
	tx, err := client.FunctionCallTx("stake_funds", bytes, deposit, gas, key, pubKey, addrFrom, addrTo)
	if err != nil {
		t.Fatalf("RPC Error: %s", err.Error())
	}
	var result bool
	err = tx.Status.Result(&result)
	if err != nil {
		t.Fatalf("Result Error: %s", err.Error())
	}
	fmt.Println(result)
}

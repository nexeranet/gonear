package client

import (
	"fmt"
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

func TestFuncCall(t *testing.T) {
	client := initTestClient(t)

	type Params struct {
		ProjectID string `json:"project_id"`
		StageID   string `json:"stage_id"`
	}
	params := Params{fmt.Sprint(1670332900497), fmt.Sprint(4)}
	args, err := EncodeBase64Args(params)
	if err != nil {
        t.Fatal(err)
	}
	response, err := client.CallContractFunc(
		"test_v1_nft.openforest.testnet",
		"get_stage",
		args)
	// check return value in response result
	if err != nil {
        t.Fatal(err)
	}
    t.Log(string(response.Result))
}

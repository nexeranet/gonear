package client

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/nexeranet/gonear/client/types"
	"golang.org/x/sync/errgroup"
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
func TestGoroutines(t *testing.T) {
	client := initTestClient(t)
	key := os.Getenv("PRIVATE")
	pubKey := os.Getenv("PUBLIC")
	acc := os.Getenv("ACCOUNT")
	call := func(t *testing.T, projectId, stageIndex string) error {
		type Params struct {
			ProjectID string `json:"project_id"`
			StageID   string `json:"stage_id"`
		}
		params := Params{projectId, stageIndex}
		bytes, err := json.Marshal(&params)
		if err != nil {
			return err
		}
		tx, err := client.FunctionCallTx(
			"allocate_stage_reward",
			bytes,
			big.NewInt(1), // deposit
			300000000,  // gas
			key,
			pubKey,
			acc,
			"test_v1_nft.openforest.testnet")

		if err != nil {
			fmt.Println(params, err)
			return err
		}
		if !tx.Status.IsSuccess() {
			return tx.Status.GetError()
		}
		return nil
	}
	type Test struct {
		id    string
		index string
	}
	tests := []Test{
		{
			id:    "1671096069968",
			index: "0",
		},
		{
			id:    "1671186172279",
			index: "3",
		},
		{
			id:    "1671096000441",
			index: "0",
		},
		{
			id:    "1671089866466",
			index: "0",
		},
		{
			id:    "1671089997354",
			index: "0",
		},
	}
	ctx := context.Background()
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(5)
	callback := func(tt Test) {
		// g.Go(func() error {
			err := call(t, tt.id, tt.index)
			t.Log(tt.id, tt.index, err)
		// 	return err
		// })
	}
	for _, tt := range tests {
		callback(tt)
	}
	if err := g.Wait(); err != nil {
		t.Fatal(err)
	}
	t.Log("SUCCESS")
}

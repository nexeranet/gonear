package client

import (
	"encoding/json"
	"fmt"
	"testing"
)

type projectStage struct {
	ValidationEndsAt string `json:"validation_ends_at"`
}

func TestViewStage(t *testing.T) {
	client := initTestClient(t)
	contractAddress := "test_v1_nft.openforest.testnet"
	projectId := "1660809317950"
	stageIndex := "1"
	type Params struct {
		ProjectID  string `json:"project_id"`
		StageID string `json:"stage_id"`
	}
	params := Params{projectId, stageIndex}
    args, err := EncodeBase64Args(params)
    if err != nil {
        t.Fatal(err)
    }
	response, err := client.CallContractFunc(
		contractAddress,
		"get_stage",
		args)
	// check return value in response result
	if err != nil {
		t.Fatal(err)
	}
	var resp projectStage
	err = json.Unmarshal(response.Result, &resp)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}

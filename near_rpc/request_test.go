package near_rpc

import (
	"encoding/base64"
	"encoding/json"
	"testing"
)

func EncodeBase64Args(value interface{}) (string, error){
	bytes, err := json.Marshal(value)
	if err != nil {
        return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}

func TestCallContractFuncByBlockId(t *testing.T) {

    type ValidatorProfile struct {
        AccountId    string `json:"account_id"`
        UsedBalance  string `json:"used_balance"`
        LastRewardId string `json:"last_reward_id"`
    }
	api := initTesnetApi()
	type Params struct {
		FromIndex string `json:"from_index"`
		Limit     string `json:"limit"`
	}
	//
	params := Params{"0", "100"}
	args, err := EncodeBase64Args(params)
    if err != nil {
        t.Fatalf("Encode %v", err)
    }
	response, err := api.RequestEndpoint("https://archival-rpc.testnet.near.org").CallContractFuncByBlockId(
		"dev_v1_reward.openforest.testnet",
		"get_validation_profiles",
		args,
		135724564,
	)
    var validators []ValidatorProfile
	// check return value in response result
	if err != nil {
	    t.Fatalf("%v",err)
	}

	json.Unmarshal(response.Result, &validators)
    t.Log(validators)
}

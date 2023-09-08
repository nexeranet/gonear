package jsonrpc

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"testing"
	"time"
)
func EncodeBase64Args(value interface{}) (string, error){
	bytes, err := json.Marshal(value)
	if err != nil {
        return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}

func TestRequest(t *testing.T) {
    httpClient :=&http.Client{
        Timeout: time.Second * 10,
    }
    request := NewRequest("https://archival-rpc.testnet.near.org", httpClient)

	type Params struct {
		RequestType string `json:"request_type"`
		BlockId     uint64 `json:"block_id"`
		AccountID   string `json:"account_id"`
		MethodName  string `json:"method_name"`
		ArgsBase64  string `json:"args_base64"`
	}
	type Args struct {
		FromIndex string `json:"from_index"`
		Limit     string `json:"limit"`
	}
	//
	args := Args{"0", "100"}
	argsbase64, err := EncodeBase64Args(args)
    if err != nil {
        t.Fatalf("Encode %v", err)
    }

	params := Params{"call_function", 135724564, "dev_v1_reward.openforest.testnet","get_validation_profiles", argsbase64}

    result, err :=  request.Call("query", &params)
    if err != nil {
        t.Fatalf("Error %v",err)
    }
    t.Log(result)
}

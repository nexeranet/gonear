package near_api

import (
	// "log"
	"testing"
)

func initApi(t *testing.T) *NearApi {
	Url := "https://rpc.testnet.near.org"
	return New(Url)
}

func TestCheckTx(t *testing.T) {

}

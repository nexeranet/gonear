package near_api

import (
	"fmt"
	types "github.com/nexeranet/gonear/near_api/types"
	"testing"
)

func initTesnetApi() NearApiI {
	Url := "https://rpc.testnet.near.org"
	return New(Url)
}
func TestErrorBlockByNumber(t *testing.T) {
	client := initTesnetApi()
	res, err := client.BlockByNumber(100655760)
	if err != nil {
		switch err.(type) {
		case *types.ErrorUnknownBlock:
			unk := err.(*types.ErrorUnknownBlock)
			fmt.Println("ErrorUnknownBlock", unk.Cause())
		case *types.ErrorParseError:
			unk := err.(*types.ErrorParseError)
			fmt.Println("ErrorParseError", unk.Cause())
		}
	}
	fmt.Println(res)
}

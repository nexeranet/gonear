package near_api

func initTesnetApi() *NearApi {
	Url := "https://rpc.testnet.near.org"
	return New(Url)
}

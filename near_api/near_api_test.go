package near_api

func initTesnetApi() NearApiI {
	Url := "https://rpc.testnet.near.org"
	return New(Url)
}

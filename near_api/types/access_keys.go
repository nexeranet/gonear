package near_api_types

type AccessKey struct {
	Nonce      uint64     `json:"nonce"`
	Permission Permission `json:"permission"`
}

type KeyItem struct {
	PublicKey string    `json:"public_key"`
	AccessKey AccessKey `json:"access_key"`
}

type AccessKeysView struct {
	Permission  Permission `json:"permission"`
	BlockHash   string     `json:"block_hash"`
	BlockHeight string     `json:"block_height"`
	Nonce       uint64     `json:"nonce"`
	Error       string     `json:"error"`
}
type AccessKeysListViev struct {
	Keys []KeyItem
}

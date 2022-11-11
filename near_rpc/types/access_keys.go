package near_rpc_types

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
	BlockHeight uint64     `json:"block_height"`
	Nonce       uint64     `json:"nonce"`
	Error       string     `json:"error"`
}
type AccessKeysListViev struct {
	Keys []KeyItem
}

type AccessKeyChange struct {
	Cause struct {
		Type   string `json:"type"`
		TxHash string `json:"tx_hash"`
	} `json:"cause"`
	Type   string `json:"type"`
	Change struct {
		AccountID string `json:"account_id"`
		PublicKey string `json:"public_key"`
		AccessKey struct {
			Nonce      int    `json:"nonce"`
			Permission string `json:"permission"`
		} `json:"access_key"`
	} `json:"change"`
}

type AccessKeyChangesView struct {
	BlockHash string            `json:"block_hash"`
	Changes   []AccessKeyChange `json:"changes"`
}

package near_api_types

type AccountView struct {
	Amount        string `json:"amount"`
	Locked        string `json:"locked"`
	CodeHash      string `json:"code_hash"`
	StorageUsage  uint64 `json:"storage_usage"`
	StoragePaidAt uint64 `json:"storage_paid_at"`
	BlockHash     string `json:"block_hash"`
	BlockHeight   uint64 `json:"block_height"`
}
type AccountChangesView struct {
	BlockHash string `json:"block_hash"`
	Changes   []struct {
		Cause struct {
			Type   string `json:"type"`
			TxHash string `json:"tx_hash"`
		} `json:"cause"`
		Type   string `json:"type"`
		Change struct {
			AccountID     string `json:"account_id"`
			Amount        string `json:"amount"`
			Locked        string `json:"locked"`
			CodeHash      string `json:"code_hash"`
			StorageUsage  int    `json:"storage_usage"`
			StoragePaidAt int    `json:"storage_paid_at"`
		} `json:"change"`
	} `json:"changes"`
}

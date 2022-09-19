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

package near_rpc_types

type ContractCodeView struct {
	CodeBase64  string `json:"code_base64"`
	Hash        string `json:"hash"`
	BlockHash   string `json:"block_hash"`
	BlockHeight uint64 `json:"block_height"`
}

type ContractStateItem struct {
	Key   string   `json:"key"`
	Value string   `json:"value"`
	Proof []string `json:"proof"`
}

type ContractStateView struct {
	Values      []ContractStateItem `json:"values"`
	Proof       []string            `json:"proof"`
	BlockHash   string              `json:"block_hash"`
	BlockHeight uint64              `json:"block_height"`
}

type ContractFuncResult struct {
	Result      []byte `json:"result"`
	BlockHeight uint64 `json:"block_height"`
	BlockHash   string `json:"block_hash"`
	Error       string `json:"error,omitempty"`
}

type ContractStateChangesView struct {
	BlockHash string `json:"block_hash"`
	Changes   []struct {
		Cause struct {
			Type        string `json:"type"`
			ReceiptHash string `json:"receipt_hash"`
		} `json:"cause"`
		Type   string `json:"type"`
		Change struct {
			AccountID   string `json:"account_id"`
			KeyBase64   string `json:"key_base64"`
			ValueBase64 string `json:"value_base64"`
		} `json:"change"`
	} `json:"changes"`
}

type ContractCodeChangesView struct {
	BlockHash string `json:"block_hash"`
	Changes   []struct {
		Cause struct {
			Type        string `json:"type"`
			ReceiptHash string `json:"receipt_hash"`
		} `json:"cause"`
		Type   string `json:"type"`
		Change struct {
			AccountID  string `json:"account_id"`
			CodeBase64 string `json:"code_base64"`
		} `json:"change"`
	} `json:"changes"`
}

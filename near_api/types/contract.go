package near_api_types

type ContractCodeView struct {
	CodeBase64  string `json:"code_base64"`
	Hash        string `json:"hash"`
	BlockHash   string `json:"block_hash"`
	BlockHeight string `json:"block_height"`
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
	BlockHeight string              `json:"block_height"`
}

type ContractFuncResult struct {
		Result      []rune `json:"result"`
		BlockHeight uint64 `json:"block_height"`
		BlockHash   string `json:"block_hash"`
}

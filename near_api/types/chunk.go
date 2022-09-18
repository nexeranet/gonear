package near_api_types

type ChunkDetailsView struct {
	Author string `json:"author"`
	Header struct {
		ChunkHash            string        `json:"chunk_hash"`
		PrevBlockHash        string        `json:"prev_block_hash"`
		OutcomeRoot          string        `json:"outcome_root"`
		PrevStateRoot        string        `json:"prev_state_root"`
		EncodedMerkleRoot    string        `json:"encoded_merkle_root"`
		EncodedLength        int           `json:"encoded_length"`
		HeightCreated        int           `json:"height_created"`
		HeightIncluded       int           `json:"height_included"`
		ShardID              int           `json:"shard_id"`
		GasUsed              int           `json:"gas_used"`
		GasLimit             int64         `json:"gas_limit"`
		RentPaid             string        `json:"rent_paid"`
		ValidatorReward      string        `json:"validator_reward"`
		BalanceBurnt         string        `json:"balance_burnt"`
		OutgoingReceiptsRoot string        `json:"outgoing_receipts_root"`
		TxRoot               string        `json:"tx_root"`
		ValidatorProposals   []interface{} `json:"validator_proposals"`
		Signature            string        `json:"signature"`
	} `json:"header"`
	Transactions []interface{} `json:"transactions"`
	Receipts     []interface{} `json:"receipts"`
}

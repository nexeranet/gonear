package near_api_types

import "fmt"

type StatusTx struct {
	SuccessValue *string   `json:"SuccessValue,omitempty"`
	Failure      *FailerTx `json:"Failure,omitempty"`
}

func (s StatusTx) IsError() bool {
	return s.Failure != nil
}

func (s StatusTx) IsSuccess() bool {
	if s.SuccessValue != nil {
		if *s.SuccessValue == "" {
			return true
		}
	}
	return false
}

type FailerTx struct {
	ActionError ActionError `json:"ActionError"`
}

func (f FailerTx) Error() error {
	return f.ActionError.Kind.ReturnError()
}

type ActionError struct {
	Index int             `json:"index"`
	Kind  ActionErrorKind `json:"kind"`
}

type AccountDoesNotExist struct {
	AccountId string `json:"account_id"`
}

func (a AccountDoesNotExist) Error() error {
	return fmt.Errorf("Account does not exist: %s", a.AccountId)
}

type ActionErrorKind struct {
	AccountDoesNotExist *AccountDoesNotExist `json:"AccountDoesNotExist,omitempty"`
}

func (a ActionErrorKind) ReturnError() error {
	if a.AccountDoesNotExist != nil {
		return a.AccountDoesNotExist.Error()
	}
	return ErrUnknown
}

type Transaction struct {
	SignerID   string                   `json:"signer_id"`
	PublicKey  string                   `json:"public_key"`
	Nonce      int                      `json:"nonce"`
	ReceiverID string                   `json:"receiver_id"`
	Actions    []map[string]interface{} `json:"actions"`
	Signature  string                   `json:"signature"`
	Hash       string                   `json:"hash"`
}
type TransactionOutcome struct {
	Proof []struct {
		Hash      string `json:"hash"`
		Direction string `json:"direction"`
	} `json:"proof"`
	BlockHash string `json:"block_hash"`
	ID        string `json:"id"`
	Outcome   struct {
		Logs        []interface{} `json:"logs"`
		ReceiptIds  []string      `json:"receipt_ids"`
		GasBurnt    int64         `json:"gas_burnt"`
		TokensBurnt string        `json:"tokens_burnt"`
		ExecutorID  string        `json:"executor_id"`
		Status      struct {
			SuccessReceiptID string `json:"SuccessReceiptId"`
		} `json:"status"`
	} `json:"outcome"`
}
type ReceiptOutcome struct {
	BlockHash string `json:"block_hash"`
	ID        string `json:"id"`
	Proof     []struct {
		Hash      string `json:"hash"`
		Direction string `json:"direction"`
	} `json:"proof"`
	Outcome struct {
		Logs        []interface{} `json:"logs"`
		ReceiptIds  []string      `json:"receipt_ids"`
		GasBurnt    int64         `json:"gas_burnt"`
		TokensBurnt string        `json:"tokens_burnt"`
		ExecutorID  string        `json:"executor_id"`
		Status      struct {
			SuccessValue string `json:"SuccessValue"`
		} `json:"status"`
	} `json:"outcome"`
}

type TxView struct {
	Status             StatusTx           `json:"status"`
	Transaction        Transaction        `json:"transaction"`
	TransactionOutcome TransactionOutcome `json:"transaction_outcome"`
	ReceiptsOutcome    []ReceiptOutcome   `json:"receipts_outcome"`
}

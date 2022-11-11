package near_rpc_types

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

type StatusTx struct {
	SuccessValue *string   `json:"SuccessValue,omitempty"`
	Failure      *FailerTx `json:"Failure,omitempty"`
}

func (t StatusTx) Result(value interface{}) error {
    if !t.IsSuccess() {
        return fmt.Errorf("Success value is nil")
    }
	decoded64, err := base64.StdEncoding.DecodeString(*t.SuccessValue)
	if err != nil {
		return err
	}
	return json.Unmarshal(decoded64, value)
}

// func (s StatusTx) CheckError() error {
// 	if !s.IsSuccess() {
// 		return fmt.Errorf("Unknown error")
// 	}
// 	return nil
// }

func (s StatusTx) IsSuccess() bool {
	if s.SuccessValue != nil {
		return true
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

func (a AccountDoesNotExist) Error() string {
	return fmt.Sprintf("Account does not exist: %s", a.AccountId)
}

type ActionErrorKind struct {
	AccountDoesNotExist *AccountDoesNotExist `json:"AccountDoesNotExist,omitempty"`
}

func (a ActionErrorKind) ReturnError() error {
	if a.AccountDoesNotExist != nil {
		return a.AccountDoesNotExist
	}
	return fmt.Errorf("Unknown error")
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

type EventLog struct {
	Standard string      `json:"standard"`
	Version  string      `json:"version"`
	Event    string      `json:"event"`
	Data     interface{} `json:"data"`
}

func (r ReceiptOutcome) GetLogs() (list []EventLog, err error) {
	for _, log := range r.Outcome.Logs {
		strlog, ok := log.(string)
		if !ok {
			continue
		}
		var event_log EventLog
		str := strings.Replace(strlog, "EVENT_JSON:", "", -1)
		err = json.Unmarshal([]byte(str), &event_log)
		if err != nil {
			return list, err
		}
		list = append(list, event_log)
	}
	return list, err
}

type TxView struct {
	Status             StatusTx           `json:"status"`
	Transaction        Transaction        `json:"transaction"`
	TransactionOutcome TransactionOutcome `json:"transaction_outcome"`
	ReceiptsOutcome    []ReceiptOutcome   `json:"receipts_outcome"`
}

type ViewReceipt struct {
	PredecessorID string `json:"predecessor_id"`
	Receipt       struct {
		Action struct {
			Actions             []interface{} `json:"actions"`
			GasPrice            string        `json:"gas_price"`
			InputDataIds        []interface{} `json:"input_data_ids"`
			OutputDataReceivers []interface{} `json:"output_data_receivers"`
			SignerID            string        `json:"signer_id"`
			SignerPublicKey     string        `json:"signer_public_key"`
		} `json:"Action"`
	} `json:"receipt"`
	ReceiptID  string `json:"receipt_id"`
	ReceiverID string `json:"receiver_id"`
}

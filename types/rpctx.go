package types

import "fmt"

type TxInfo struct {
	Hash string `json:"hash"`
}

type TransactionOutcome struct {
	BlockHash string `json:"block_hash"`
}

type Transaction struct {
	Status             StatusTx           `json:"status"`
	Transaction        TxInfo             `json:"transaction"`
	TransactionOutcome TransactionOutcome `json:"transaction_outcome"`
}

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

package near_rpc_types

import (
	"fmt"
	"strings"
)

type FailerTx struct {
	ActionError ActionError `json:"ActionError"`
}

func (f FailerTx) Error() string {
	var list []string
	for key, value := range f.ActionError.Kind {
		list = append(list, fmt.Sprintf("%s: %v", key, value))
	}
    return fmt.Sprintf("Index: %d, Kind: %v", f.ActionError.Index, strings.Join(list, ", "))
}

type ActionError struct {
	Index int                    `json:"index"`
	Kind  map[string]interface{} `json:"kind"`
}

// type ActionErrorKind struct {
// 	AccountDoesNotExist *AccountDoesNotExist `json:"AccountDoesNotExist,omitempty"`
// 	FunctionCallError   *FunctionCallError   `json:"FunctionCallError,omitempty"`
// }
//
// func (a ActionErrorKind) GetError() error {
// 	if a.AccountDoesNotExist != nil {
// 		return a.AccountDoesNotExist
// 	}
// 	if a.FunctionCallError != nil {
// 		return a.FunctionCallError
// 	}
// 	return fmt.Errorf("Untyped unknown error")
// }
//
// type FunctionCallError struct {
// 	ExecutionError string `json:"ExecutionError"`
// }
//
// func (f *FunctionCallError) Error() string {
// 	return f.ExecutionError
// }
//
// type AccountDoesNotExist struct {
// 	AccountId string `json:"account_id"`
// }
//
// func (a *AccountDoesNotExist) Error() string {
// 	return fmt.Sprintf("Account does not exist: %s", a.AccountId)
// }

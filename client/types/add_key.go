package types

import (
	"math/big"

	"github.com/near/borsh-go"
)

type AccessKeyPermissionEnum borsh.Enum

const (
	FunctionCallPermissionEnum AccessKeyPermissionEnum = iota
	FullAccessPermissionEnum
)

type FunctionCallPermission struct {
	Allowance   *big.Int
	ReceiverID  string
	MethodNames []string
}

type FullAccessPermission struct{}

type AccessKeyPermission struct {
	Enum         AccessKeyPermissionEnum `borsh_enum:"true"`
	FunctionCall FunctionCallPermission
	FullAccess   FullAccessPermission
}

type AccessKey struct {
	Nonce      uint64
	Permission AccessKeyPermission
}

type AddKey struct {
	PublicKey PublicKey
	AccessKey AccessKey
}

func AddKeyAction(addkey AddKey) Action {
    return Action{
        Enum: AddKeyEnum,
        AddKey: addkey,
    }
}

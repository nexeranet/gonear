package types

import (
	"errors"
	"math/big"

	"github.com/near/borsh-go"
)

var ErrUnknown = errors.New("Unknown error")

type AccessKeyPermissionEnum borsh.Enum

const (
	FunctionCallPermissionEnum AccessKeyPermissionEnum = iota
	FullAccessPermissionEnum
)

type ActionEnum borsh.Enum

const (
	CreateAccountEnum ActionEnum = iota
	DeployContractEnum
	FunctionCallEnum
	TransferEnum
	StakeEnum
	AddKeyEnum
	DeleteKeyEnum
	DeleteAccountEnum
)

type Signature struct {
	KeyType uint8
	Data    [64]byte
}

type PublicKey struct {
	KeyType uint8
	Data    [32]byte
}

type AccessKey struct {
	Nonce      uint64
	Permission AccessKeyPermission
}

type AccessKeyPermission struct {
	Enum         AccessKeyPermissionEnum
	FunctionCall FunctionCallPermission
	FullAccess   FullAccessPermission
}

type FunctionCallPermission struct {
	Allowance   *big.Int
	ReceiverID  string
	MethodNames []string
}

type FullAccessPermission struct{}

// type Action struct {
// 	Enum           ActionEnum
// 	CreateAccount  *CreateAccount
// 	DeployContract *DeployContract
// 	FunctionCall   *FunctionCall
// 	Transfer       *Transfer
// 	Stake          *Stake
// 	AddKey         *AddKey
// 	DeleteKey      *DeleteKey
// 	DeleteAccount  *DeleteAccount
// }

type CreateAccount struct{}

type DeployContract struct {
	Code []byte
}

type Stake struct {
	Stake     big.Int
	PublicKey PublicKey
}

type AddKey struct {
	PublicKey PublicKey
	AccessKey AccessKey
}

type DeleteKey struct {
	PublicKey PublicKey
}

type DeleteAccount struct {
	BeneficiaryID string
}

package types

import (
	"errors"

	"github.com/near/borsh-go"
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

var ErrUnknown = errors.New("Unknown error")

type AccessKeys struct {
	Permission string `json:"permission"`
	BlockHash  string `json:"block_hash"`
	Nonce      uint64 `json:"nonce"`
}

type HeaderBlock struct {
	Hash string `json:"hash"`
}

type Block struct {
	Author string      `json:"author"`
	Header HeaderBlock `json:"header"`
}

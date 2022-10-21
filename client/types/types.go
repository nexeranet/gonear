package types

import (
	"crypto/sha256"
	"encoding/base64"
	"github.com/near/borsh-go"
	"golang.org/x/crypto/nacl/sign"
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

func NewSignature(data [64]byte) Signature {
	return Signature{
		KeyType: 0,
		Data:    data,
	}
}

type PublicKey struct {
	KeyType uint8
	Data    [32]byte
}

func NewPublicKey(data [32]byte) PublicKey {
	return PublicKey{
		KeyType: 0,
		Data:    data,
	}
}

type Action struct {
	Enum           ActionEnum `borsh_enum:"true"`
	CreateAccount  CreateAccount
	DeployContract DeployContract
	FunctionCall   FunctionCall
	Transfer       Transfer
	Stake          Stake
	AddKey         AddKey
	DeleteKey      DeleteKey
	DeleteAccount  DeleteAccount
}

type Transaction struct {
	SignerId   string
	PublicKey  PublicKey
	Nonce      uint64
	ReceiverId string
	BlockHash  [32]byte
	Actions    []Action
}

func (t Transaction) Sign(privKey *[64]byte) (signTx [64]byte, err error) {
	tx, err := borsh.Serialize(t)
	if err != nil {
		return signTx, err
	}
	hash := sha256.Sum256(tx)
	signature := sign.Sign(nil, hash[:], privKey)
	return *(*[64]byte)(signature), nil
}

type SignedTransaction struct {
	Transaction Transaction
	Signature   Signature
}

func (t SignedTransaction) Base64Encode() (string, error) {
	data, err := borsh.Serialize(t)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

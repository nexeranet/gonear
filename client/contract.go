package client

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"

	"github.com/mr-tron/base58"
	"github.com/near/borsh-go"
	"github.com/nexeranet/gonear/types"
	"golang.org/x/crypto/nacl/sign"
)

func (a *Client) SendCallFunctionTx(methodName string, args []byte, gas uint64, key, publicKey, addrFrom, addrTo string) (string, error) {
	permission, block_hash, nonce, err := a.GetAccessKeys(addrFrom, publicKey)
	if err != nil {
		return "", err
	}
	if permission.String != "FullAccess" {
		return "", fmt.Errorf("`Account %s does not have permission to send tokens using key: %s", addrFrom, string(publicKey[:]))
	}
	publicKeyBytes, privKeyBytes, err := getKeys(key)
	//_, privKeyBytes, err := getKeys(key)
	if err != nil {
		return "", err
	}
	nonce_tx := nonce + 1
	block_hash_dec, err := base58.Decode(block_hash)
	if err != nil {
		return "", err
	}
	deposit := big.NewInt(1)
	action := types.FunctionCallAction{
		Enum: types.FunctionCallEnum,
		FunctionCall: types.FunctionCall{
			Deposit:    *deposit,
			MethodName: methodName,
			Gas:        gas,
			Args:       args,
		},
	}
	actions := []types.FunctionCallAction{action}
	block_hash_dec_static := (*[32]byte)(block_hash_dec)
	tx := types.TxFunctionCall{
		SignerId: addrFrom,
		PublicKey: types.PublicKey{
			Data: *publicKeyBytes,
		},
		Nonce:      nonce_tx,
		ReceiverId: addrTo,
		Actions:    actions,
		BlockHash:  *block_hash_dec_static,
	}
	//log.Println(tx)
	serialized_tx, err := borsh.Serialize(tx)
	if err != nil {
		return "", err
	}
	serializedTxHash := sha256.Sum256(serialized_tx)
	signature := sign.Sign(nil, serializedTxHash[:], privKeyBytes)
	signature_fixed := (*[64]byte)(signature)
	signed_tx := types.SignedTxFunctionCall{
		Transaction: tx,
		Signature: types.Signature{
			KeyType: 0,
			Data:    *signature_fixed,
		},
	}
	data, err := borsh.Serialize(signed_tx)
	if err != nil {
		return "", err
	}
	encoded_bs64 := base64.StdEncoding.EncodeToString(data)
	hash, err := a.SendAsyncTx(encoded_bs64)
	fmt.Println(encoded_bs64, err)
	return hash, err
}

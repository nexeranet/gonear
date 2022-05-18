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

func (a *Client) SendTransferTx(amount *big.Int, key, publicKey, addrFrom, addrTo string) (*types.Transaction, error) {
	permission, block_hash, nonce, err := a.GetAccessKeys(addrFrom, publicKey)
	if err != nil {
		return nil, err
	}
	if permission.String != "FullAccess" {
		return nil, fmt.Errorf("`Account %s does not have permission to send tokens using key: %s", addrFrom, string(publicKey[:]))
	}
	publicKeyBytes, privKeyBytes, err := getKeys(key)
	if err != nil {
		return nil, err
	}
	nonce_tx := nonce + 1
	block_hash_dec, err := base58.Decode(block_hash)
	if err != nil {
		return nil, err
	}
	action := types.TransferAction{
		Enum: types.TransferEnum,
		Transfer: types.Transfer{
			Deposit: *formatAmount(amount),
		},
	}
	actions := []types.TransferAction{action}
	block_hash_dec_fix := (*[32]byte)(block_hash_dec)
	tx := types.TxTransfer{
		SignerId: addrFrom,
		PublicKey: types.PublicKey{
			Data: *publicKeyBytes,
		},
		Nonce:      nonce_tx,
		ReceiverId: addrTo,
		Actions:    actions,
		BlockHash:  *block_hash_dec_fix,
	}
	serialized_tx, err := borsh.Serialize(tx)
	if err != nil {
		return nil, err
	}
	serializedTxHash := sha256.Sum256(serialized_tx)
	signature := sign.Sign(nil, serializedTxHash[:], privKeyBytes)
	signature_fixed := (*[64]byte)(signature)
	signed_tx := types.SignedTxTransfer{
		Transaction: tx,
		Signature: types.Signature{
			KeyType: 0,
			Data:    *signature_fixed,
		},
	}
	data, err := borsh.Serialize(signed_tx)
	if err != nil {
		return nil, err
	}
	encoded_bs64 := base64.StdEncoding.EncodeToString(data)
	//fmt.Println(encoded_bs64)
	return a.SendAwaitTx(encoded_bs64)
}

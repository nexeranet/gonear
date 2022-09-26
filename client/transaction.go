package client

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/mr-tron/base58"
	"github.com/near/borsh-go"
	near_api_types "github.com/nexeranet/gonear/near_api/types"
	"github.com/nexeranet/gonear/types"
	"golang.org/x/crypto/nacl/sign"
)

func SendTransaction[T types.Action] (a *Client, key, publicKey, addrFrom, addrTo string, actions types.Actions[T]) (*near_api_types.TxView, error) {
	access_key, err := a.C.ViewAccessKey(addrFrom, publicKey)
	if err != nil {
		return nil, err
	}
	if access_key.Permission.String != "FullAccess" {
		return nil, fmt.Errorf("`Account %s does not have permission to send tokens using key: %s", addrFrom, string(publicKey[:]))
	}
	publicKeyBytes, privKeyBytes, err := getKeys(key)
	if err != nil {
		return nil, err
	}
	nonce_tx := access_key.Nonce + 1
	block_hash_dec, err := base58.Decode(access_key.BlockHash)
	if err != nil {
		return nil, err
	}
	block_hash_dec_fix := (*[32]byte)(block_hash_dec)
	tx := types.Transaction[T]{
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
	signed_tx := types.SignedTransaction[T]{
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
    fmt.Println(encoded_bs64)
	return a.C.SendAwaitTx(encoded_bs64)
}

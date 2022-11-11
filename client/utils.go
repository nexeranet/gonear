package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mr-tron/base58"
	"github.com/nexeranet/gonear/client/types"
	"golang.org/x/crypto/nacl/sign"
)

func validatePrivateKey(key string) ([]byte, error) {
	parts := strings.Split(key, ":")
	if len(parts) == 1 {
		return base58.Decode(parts[0])
	} else if len(parts) == 2 {
		switch v := strings.ToUpper(parts[0]); v {
		case "ED25519":
			return base58.Decode(parts[1])
		default:
			return nil, fmt.Errorf("Unknown curve: %s", parts[0])
		}
	} else {
		return nil, fmt.Errorf("Invalid encoded key format, must be <curve>:<encoded key>'")
	}
}

func getKeys(key string) (publicKey *[32]byte, privateKey *[64]byte, err error) {
	validKey, err := validatePrivateKey(key)
	if err != nil {
		return nil, nil, nil
	}
	public, private, err := sign.GenerateKey(bytes.NewReader(validKey))
	return public, private, err
}

// Generate base64 signed transaction hash
func GenerateActionsTransactionHash(addrFrom, addrTo, key string, nonce uint64, blockHash [32]byte, actions []types.Action) (string, error) {
	publicKey, privKey, err := getKeys(key)
	if err != nil {
		return "", err
	}
	tx := types.Transaction{
		SignerId:   addrFrom,
		PublicKey:  types.NewPublicKey(*publicKey),
		Nonce:      nonce,
		ReceiverId: addrTo,
		Actions:    actions,
		BlockHash:  blockHash,
	}
	signatureData, err := tx.Sign(privKey)
	if err != nil {
		return "", err
	}
	signed_tx := types.SignedTransaction{
		Transaction: tx,
		Signature:   types.NewSignature(signatureData),
	}
	return signed_tx.Base64Encode()
}

func EncodeToBase64(v interface{}) (string, error) {
    var buf bytes.Buffer
    encoder := base64.NewEncoder(base64.StdEncoding, &buf)
    err := json.NewEncoder(encoder).Encode(v)
    if err != nil {
        return "", err
    }
    encoder.Close()
    return buf.String(), nil
}

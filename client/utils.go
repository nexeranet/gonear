package client

import (
	"bytes"
	"fmt"
	"math/big"
	"strings"

	"github.com/mr-tron/base58"
	"golang.org/x/crypto/nacl/sign"
)

func formatAmount(amount *big.Int) *big.Int {
	yotto := big.NewInt(0)
	yotto.SetString("1000000000000000000000000", 10)
	yotto.Mul(yotto, amount)
	return yotto
}

func NewYottoNear(num *big.Int) *big.Int {
	return formatAmount(num)
}

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

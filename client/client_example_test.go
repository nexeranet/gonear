package client_test

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/nexeranet/gonear/client"
	"github.com/nexeranet/gonear/client/types"
	near_rpc_types "github.com/nexeranet/gonear/near_rpc/types"
)

func ExampleClient_BalanceAt() {
	// Connect the client.
	account := "nexeranet.testnet"
	url := "https://rpc.testnet.near.org"
	near := client.NewClient(url)
	balance, err := near.BalanceAt(account)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(balance)
}

func ExampleClient_TransferTx() {
	// Connect the client.
	url := "https://rpc.testnet.near.org"
	near := client.NewClient(url)

	// private key
	key := "ed25519:111111"
	// public key
	pubKey := "ed25519:22222"
	// address from
	addrFrom := "account.mainnet"
	// address to
	addrTo := "token.arhius.testnet"
	// amount of near (*big.Int yoctoNear)
	amount := types.NewNear(1).BigInt()
	tx, err := near.TransferTx(amount, key, pubKey, addrFrom, addrTo)
	if err != nil {
		switch err.(type) {
		// Check error type
		case *near_rpc_types.ErrorParseError:
			fmt.Println("Error parse error")
		}
		return
	}
	fmt.Println(tx)
}

func ExampleClient_CheckTx() {
	// Connect the client.
	url := "https://rpc.testnet.near.org"
	near := client.NewClient(url)
	// sender account address
	sender := "prod-users.kaiching.testnet"
	// transaction hash
	txHash := "DL9Jw3pRduc9tVRjCibGtzBPMLDUwNnpcJm2aDUWAQeu"
	tx, err := near.CheckTx(txHash, sender)
	if err != nil {
		switch err.(type) {
		// Check error type
		case *near_rpc_types.ErrorUnknownTransaction:
			fmt.Println("Error unknown transaction")
		default:
			fmt.Println(err)
		}
		return
	}
	fmt.Println(tx)
}

func ExampleClient_ActionsTx() {
	// Connect the client.
	url := "https://rpc.testnet.near.org"
	near := client.NewClient(url)

	// private key
	key := "ed25519:111111"
	// public key
	pubKey := "ed25519:222222"
	// address from
	addrFrom := "account.mainnet"
	// address to
	addrTo := "token.arhius.testnet"
	// amount of near (*big.Int yoctoNear)
	amount := types.NewNear(1).BigInt()
	tx, err := near.ActionsTx(key, pubKey, addrFrom, addrTo, []types.Action{
		types.TransferAction(*amount),
		types.CreateAccountAction(),
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tx)
}

func ExampleClient_AsyncActionsTx() {
	// Connect the client.
	url := "https://rpc.testnet.near.org"
	near := client.NewClient(url)

	// private key
	key := "ed25519:111111"
	// public key
	pubKey := "ed25519:222222"
	// address from
	addrFrom := "account.mainnet"
	// address to
	addrTo := "token.arhius.testnet"
	// amount of near (*big.Int yoctoNear)
	amount := types.NewNear(1).BigInt()
	txHash, err := near.AsyncActionsTx(key, pubKey, addrFrom, addrTo, []types.Action{
		types.TransferAction(*amount),
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(txHash)
}

func ExampleClient_FunctionCallTx() {
	// Connect the client.
	url := "https://rpc.testnet.near.org"
	near := client.NewClient(url)
    // method name
    method := "get_method"
	// private key
	key := "ed25519:111111"
	// public key
	pubKey := "ed25519:2222222"
	// address from
	addrFrom := "nexeranet.testnet"
	// address to
	addrTo := "token.arhius.testnet"
    // function argumets in json
    type args struct {
        I int `json:"i"`
    }
    g := args{1}
	bytes, err := json.Marshal(&g)
    if err != nil{
        fmt.Println(err)
        return
    }
    // deposit
    deposit := big.NewInt(1)
    // gas
    var gas uint64 = 1
	tx, err := near.FunctionCallTx(method, bytes, deposit, gas, key, pubKey, addrFrom, addrTo)
	if err != nil {
		switch err.(type) {
		// Check error type
		case *near_rpc_types.ErrorParseError:
			fmt.Println("Error parse")
		default:
			fmt.Println(err)
		}
		return
	}
	fmt.Println(tx)
}

func ExampleClient_CallContractFunc() {
	// Connect the client.
	url := "https://rpc.testnet.near.org"
	near := client.NewClient(url)
    // contract address
    contract := "contract_address"
    // method
    method := "get_method"
    // arguments in base 64
    type V struct {
        V string `json:"v"`
    }

    args, err := client.EncodeToBase64(V{"Test"})
    if err  != nil {
        fmt.Println(err)
        return
    }
    result, err := near.CallContractFunc(contract, method, args)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(result)
}


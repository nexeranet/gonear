package client_test

import (
	"fmt"

	"github.com/nexeranet/gonear/client"
	"github.com/nexeranet/gonear/client/types"
	near_rpc_types "github.com/nexeranet/gonear/near_rpc/types"
)

func ExampleClient_BalanceAt() {
	// Connect the client.
	account := "nexeranet.testnet"
	url := "https://rpc.testnet.near.org"
	client := client.NewClient(url)
	balance, err := client.BalanceAt(account)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(balance)
}

func ExampleClient_TransferTx() {
	// Connect the client.
	url := "https://rpc.testnet.near.org"
	client := client.NewClient(url)

	// private key
	key := "ed25519:5XKLL4yQoBVyHCUyXrMt9898VG7My2iWomu1GC3wAW4V6eBwZGmreqpMiWfC1HiVpmAAWCe1pJ6RKNuEFgupbPjK"
	// public key
	pubKey := "ed25519:7phkB1HWhWETQ1WkErTUS58s1EjMr4F8JFYg9VTQDk3X"
	// address from
	addrFrom := "nexeranet.testnet"
	// address to
	addrTo := "token.arhius.testnet"
	// amount of near (*big.Int yoctoNear)
	amount := types.NewNear(1).BigInt()
	tx, err := client.TransferTx(amount, key, pubKey, addrFrom, addrTo)
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
	client := client.NewClient(url)
    // sender account address
	sender := "prod-users.kaiching.testnet"
    // transaction hash
	txHash := "DL9Jw3pRduc9tVRjCibGtzBPMLDUwNnpcJm2aDUWAQeu"
	tx, err := client.CheckTx(txHash, sender)
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

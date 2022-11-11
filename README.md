 <pre>
                                          _
 _ __   _____  _____ _ __ __ _ _ __   ___| |_
| '_ \ / _ \ \/ / _ \ '__/ _` | '_ \ / _ \ __|
| | | |  __/>  <  __/ | | (_| | | | |  __/ |_
|_| |_|\___/_/\_\___|_|  \__,_|_| |_|\___|\__|
</pre>
# gonear

A NEAR client written in Go

```go
package main

import (
	"fmt"
	"time"
	"github.com/nexeranet/gonear/client"
)

func main() {

	// Connect the client.
	url := "https://rpc.mainnet.near.org"
	client := client.NewClient(url)

	// private key
	key := "ed25519:private key"
	// public key
	pubKey := "ed25519:public key"
	// address from
	addrFrom := "sender.near"
	// address to
	addrTo := "reciver.near"
	// amount of near (*big.Int yoctoNear)
    // 1 near = 1 * 10^23
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
```

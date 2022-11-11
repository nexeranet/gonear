package near_rpc

import (
	"reflect"
	"testing"

	types "github.com/nexeranet/gonear/near_rpc/types"
)

func TestCheckTx(t *testing.T) {
	type Test struct {
		name    string
		hash    string
		sender  string
		errType error
	}
	api := initTesnetApi()
	tests := []Test{
		{
			name:    "base case",
			hash:    "5xTRhNFtFsUEaBoZk9eEZjcLqTb8SAEAw4EdfjPFj4vZ",
			sender:  "perp.spin-fi.testnet",
			errType: &types.ErrorUnknownTransaction{},
		},
		{
			name:    "Tx hash invalid",
			hash:    "6zgh2u9DqHHiXz111111111",
			sender:  "perp.spin-fi.testnet",
			errType: &types.ErrorParseError{},
		},
		{
			name:    "Tx hash not found",
			hash:    "6zgh2u9DqHHiXzdy9ouTP7oGky2T4nugqzqt9wJZwNFm",
			sender:  "perp.spin-fi.testnet",
			errType: &types.ErrorUnknownTransaction{},
		},
		{
			name:    "Tx for contract call",
			hash:    "9pGS3NpV8dY87oRw1Yf4KAygwG2BCUT5imdQgVA8fd5T",
			sender:  "sbv2-authority.testnet",
			errType: &types.ErrorUnknownTransaction{},
		},
		{
			name:    "Tx for delete account",
			hash:    "8wwpDsv4LaBSeb2mCgVmC2igeMav3T7rqQZ38RnwT5Zy",
			sender:  "dev-1663602270988-50546215737119",
			errType: &types.ErrorUnknownTransaction{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := api.CheckTx(tt.hash, tt.sender)
			if err != nil {
				expect := reflect.TypeOf(tt.errType)
				have := reflect.TypeOf(err)
				if have != expect {
					t.Fatalf("Unexpected error %s, have type: %s, expect type %#v",
						err,
						have.String(),
						expect,
					)
				}
			} else {
				if result == nil {
					t.Fatalf("Expect struct, not nil")
				}
			}
		})
	}

}
func TestSendAsyncTx(t *testing.T) {
	type Test struct {
		name         string
		signedTxHash string
		isError      bool
	}
	api := initTesnetApi()

	tests := []Test{
		{
			name:         "base case",
			signedTxHash: "DgAAAHNlbmRlci50ZXN0bmV0AOrmAai64SZOv9e/naX4W15pJx0GAap35wTT1T/DwcbbDwAAAAAAAAAQAAAAcmVjZWl2ZXIudGVzdG5ldNMnL7URB1cxPOu3G8jTqlEwlcasagIbKlAJlF5ywVFLAQAAAAMAAACh7czOG8LTAAAAAAAAAGQcOG03xVSFQFjoagOb4NBBqWhERnnz45LY4+52JgZhm1iQKz7qAdPByrGFDQhQ2Mfga8RlbysuQ8D8LlA6bQE=",
			isError:      false,
		},
		{
			name:         "invalid signed tx",
			signedTxHash: "asdfasdfasfdas1241213",
			isError:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := api.SendAsyncTx(tt.signedTxHash)
			if err != nil && !tt.isError {
				t.Fatalf("expected not error, actual %s", err)
			}
			if err == nil && tt.isError {
				t.Fatalf("Expect error, have nil")
			}
		})
	}
}

func TestSendAwaitTx(t *testing.T) {
	type Test struct {
		name         string
		signedTxHash string
		isError      bool
	}
	api := initTesnetApi()

	tests := []Test{
		{
			name:         "expired signed tx",
			signedTxHash: "DgAAAHNlbmRlci50ZXN0bmV0AOrmAai64SZOv9e/naX4W15pJx0GAap35wTT1T/DwcbbDwAAAAAAAAAQAAAAcmVjZWl2ZXIudGVzdG5ldNMnL7URB1cxPOu3G8jTqlEwlcasagIbKlAJlF5ywVFLAQAAAAMAAACh7czOG8LTAAAAAAAAAGQcOG03xVSFQFjoagOb4NBBqWhERnnz45LY4+52JgZhm1iQKz7qAdPByrGFDQhQ2Mfga8RlbysuQ8D8LlA6bQE=",
			isError:      true,
		},
		{
			name:         "invalid signed tx",
			signedTxHash: "asdfasdfasfdas1241213",
			isError:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := api.SendAwaitTx(tt.signedTxHash)
			if err != nil && !tt.isError {
				t.Fatalf("Test %s, expected not error, actual %s", tt.name, err)
			}
			if err == nil && tt.isError {
				t.Fatalf("Test %s, Expect error, have nil", tt.name)
			}
		})
	}
}

func TestTxStatusWithReceipts(t *testing.T) {
	type Test struct {
		name    string
		txHash  string
		sender  string
		errType error
	}
	api := initTesnetApi()
	tests := []Test{
		{
			name:    "Base case",
			txHash:  "CBCFeceYUgSknaV7TBjofX4Zg6geGJyZqpxcxFnnogiA",
			sender:  "harr4.testnet",
			errType: &types.ErrorUnknownTransaction{},
		},
		{
			name:    "invalid tx hash",
			txHash:  "asdfasdfasfdas1241213",
			sender:  "harr4.testnet",
			errType: &types.ErrorParseError{},
		},
		{
			name:    "invalid sender",
			txHash:  "CBCFeceYUgSknaV7TBjofX4Zg6geGJyZqpxcxFnnogiA",
			sender:  "asdfsa$$$AAA",
			errType: &types.ErrorParseError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := api.TxStatusWithReceipts(tt.txHash, tt.sender)
			if err != nil {
				expect := reflect.TypeOf(tt.errType)
				have := reflect.TypeOf(err)
				if have != expect {
					t.Fatalf("Unexpected error %s, have type: %s, expect type %#v",
						err,
						have.String(),
						expect,
					)
				}
			} else {
				if result == nil {
					t.Fatalf("Expect struct, not nil")
				}
			}
		})
	}
}

func TestReceiptbyId(t *testing.T) {
	type Test struct {
		name      string
		receiptId string
		errType   error
	}
	api := initTesnetApi()
	tests := []Test{
		{
			name:      "Base case",
			receiptId: "Hfe4QVnXxJLMpmjKAss8SnMhHgV55ZDAfFcEavXXcqD4",
			errType:   &types.ErrorUnknownReceipt{},
		},
		{
			name:      "invalid receipt id",
			receiptId: "asdfsa$$$$$$$",
			errType:   &types.ErrorParseError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := api.ReceiptbyId(tt.receiptId)
			if err != nil {
				expect := reflect.TypeOf(tt.errType)
				have := reflect.TypeOf(err)
				if have != expect {
					t.Fatalf("Unexpected error %s, have type: %s, expect type %#v",
						err,
						have.String(),
						expect,
					)
				}
			} else {
				if result == nil {
					t.Fatalf("Expect struct, not nil")
				}
			}
		})
	}
}

// func TestGetLogs(t *testing.T) {
// 	api := initTesnetApi()
// 	result, err := api.CheckTx("DRMZCyj1F5kjCjrdnTQRGmSswac6RZetQ6r2LGs1c4dg", "sbv2-authority.testnet")
// 	if err != nil {
// 		t.Fatalf("%v", err)
// 	}
// 	log := result.ReceiptsOutcome[0]
//     event_log, err := log.GetLogs()
//     if err != nil {
//         t.Fatalf("%v", err)
//     }
// 	fmt.Println(event_log)
// }

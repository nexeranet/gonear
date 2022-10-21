package near_rpc

import (
	"testing"
)

func TestCheckTx(t *testing.T) {
	type Test struct {
		name    string
		hash    string
		sender  string
		isError bool
	}
	api := initTesnetApi()
	tests := []Test{
		{
			name:    "base case",
			hash:    "5xTRhNFtFsUEaBoZk9eEZjcLqTb8SAEAw4EdfjPFj4vZ",
			sender:  "perp.spin-fi.testnet",
			isError: false,
		},
		{
			name:    "Tx hash invalid",
			hash:    "6zgh2u9DqHHiXz111111111",
			sender:  "perp.spin-fi.testnet",
			isError: true,
		},
		{
			name:    "Tx hash not found",
			hash:    "6zgh2u9DqHHiXzdy9ouTP7oGky2T4nugqzqt9wJZwNFm",
			sender:  "perp.spin-fi.testnet",
			isError: true,
		},
		{
			name:    "Tx for contract call",
			hash:    "9pGS3NpV8dY87oRw1Yf4KAygwG2BCUT5imdQgVA8fd5T",
			sender:  "sbv2-authority.testnet",
			isError: false,
		},
		{
			name:    "Tx for delete account",
			hash:    "8wwpDsv4LaBSeb2mCgVmC2igeMav3T7rqQZ38RnwT5Zy",
			sender:  "dev-1663602270988-50546215737119",
			isError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, err := api.CheckTx(tt.hash, tt.sender)
			if err != nil && !tt.isError {
				t.Fatalf("expected not error, actual %s", err)
			}
			if err == nil && tt.isError {
				t.Fatalf("Expect error, have nil")
			}
			if tx == nil && !tt.isError {
				t.Fatalf("Tx is nil, expect tx view struct, have %s", err)
			}
			if tx != nil && tx.Transaction.Hash != tt.hash {
				t.Fatalf("Expect %s, have %s", tt.hash, tx.Transaction.Hash)
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
		isError bool
	}
	api := initTesnetApi()
	tests := []Test{
		{
			name:    "Base case",
			txHash:  "CBCFeceYUgSknaV7TBjofX4Zg6geGJyZqpxcxFnnogiA",
			sender:  "harr4.testnet",
			isError: false,
		},
		{
			name:    "invalid tx hash",
			txHash:  "asdfasdfasfdas1241213",
			sender:  "harr4.testnet",
			isError: true,
		},
		{
			name:    "invalid sender",
			txHash:  "CBCFeceYUgSknaV7TBjofX4Zg6geGJyZqpxcxFnnogiA",
			sender:  "asdfsa$$$AAA",
			isError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, err := api.TxStatusWithReceipts(tt.txHash, tt.sender)
			if err != nil && !tt.isError {
				t.Fatalf("Test %s, expected not error, actual %s", tt.name, err)
			}
			if err == nil && tt.isError {
				t.Fatalf("Test %s, Expect error, have nil", tt.name)
			}
			if tx != nil && tx.Transaction.SignerID != tt.sender {
				t.Fatalf("Test %s, Expect sender %s, have %s", tt.name, tt.sender, tx.Transaction.SignerID)
			}
		})
	}
}

func TestReceiptbyId(t *testing.T) {
	type Test struct {
		name      string
		receiptId string
		isError   bool
	}
	api := initTesnetApi()
	tests := []Test{
		{
			name:      "Base case",
			receiptId: "Hfe4QVnXxJLMpmjKAss8SnMhHgV55ZDAfFcEavXXcqD4",
			isError:   false,
		},
		{
			name:      "invalid receipt id",
			receiptId: "asdfsa$$$$$$$",
			isError:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := api.ReceiptbyId(tt.receiptId)
			if err != nil && !tt.isError {
				t.Fatalf("Test %s, expected not error, actual %s", tt.name, err)
			}
			if err == nil && tt.isError {
				t.Fatalf("Test %s, Expect error, have nil", tt.name)
			}
			if err == nil && result == nil {
				t.Fatalf("Test %s, Error and result is nil", tt.name)
			}
		})
	}
}

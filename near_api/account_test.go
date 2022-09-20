package near_api

import (
    "testing"
)

func TestViewAccount(t *testing.T){
    type Test struct {
        name string
        accountId string
        isError bool
    }

	client := initTesnetApi()
    tests := []Test{
        {
            name: "Valid accound id",
            accountId: "nexeranet.testnet",
        },
        {
            name: "Contract accound id",
            accountId: "vfinal.token.sweat.testnet",
        },
        {
            name: "invalid account id",
            accountId: "___",
            isError: true,
        },
    }
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list, err := client.ViewAccount(tt.accountId)
			if err != nil && !tt.isError {
				t.Fatalf("expected not error, actual %s", err)
			}
			if err == nil && tt.isError {
				t.Fatalf("Expect error, have nil")
			}
            if !tt.isError && list == nil {
				t.Fatalf("Expect struct, not nil")
            }
		})
	}
}

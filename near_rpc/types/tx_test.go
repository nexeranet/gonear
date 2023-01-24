package near_rpc_types

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeString(t *testing.T) {
	startValue := "aaaaa"
	uEnc := base64.URLEncoding.EncodeToString([]byte(startValue))
	status := StatusTx{
		SuccessValue: &uEnc,
	}
	valuePointer := ""
	err := status.Result(&valuePointer)
	if err != nil {
		t.Fatal(err)
	}
	if valuePointer != startValue {
		t.Fatalf("Expect %s, have %s", startValue, valuePointer)
	}
}

func TestDecodeBool(t *testing.T) {
	uEnc := base64.URLEncoding.EncodeToString([]byte("true"))
	status := StatusTx{
		SuccessValue: &uEnc,
	}
	valuePointer := false
	err := status.Result(&valuePointer)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDecodeInt(t *testing.T) {
	startValue := 123
	uEnc := base64.URLEncoding.EncodeToString([]byte(fmt.Sprint(startValue)))
	status := StatusTx{
		SuccessValue: &uEnc,
	}
	var valuePointer int = 0
	err := status.Result(&valuePointer)
	if err != nil {
		t.Fatal(err)
	}
    if valuePointer != startValue {
		t.Fatalf("Expect %d, have %d", startValue, valuePointer)
    }

}

func TestDecodeStruct(t *testing.T) {
	type Prm struct {
		V string `json:"v"`
		D int    `json:"d"`
	}
	jsonPrm, _ := json.Marshal(Prm{"hello", 12})
	uEnc := base64.URLEncoding.EncodeToString(jsonPrm)
	status := StatusTx{
		SuccessValue: &uEnc,
	}
	var valuePointer Prm
	err := status.Result(&valuePointer)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDecodeArray(t *testing.T) {
	jsonPrm, _ := json.Marshal([3]string{"a", "b", "c"})
	uEnc := base64.URLEncoding.EncodeToString(jsonPrm)
	status := StatusTx{
		SuccessValue: &uEnc,
	}
	var list [3]string
	err := status.Result(&list)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDecodeMap(t *testing.T) {
	jsonPrm, _ := json.Marshal(map[string]string{
		"h": "a",
	})
	uEnc := base64.URLEncoding.EncodeToString(jsonPrm)
	status := StatusTx{
		SuccessValue: &uEnc,
	}
	var list map[string]string
	err := status.Result(&list)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDecodeSliceStruct(t *testing.T) {
	type Prm struct {
		V string `json:"v"`
		D int    `json:"d"`
	}
	jsonPrm, _ := json.Marshal([]Prm{
		{
			V: "h",
			D: 1,
		},
	})
	uEnc := base64.URLEncoding.EncodeToString(jsonPrm)
	status := StatusTx{
		SuccessValue: &uEnc,
	}
	var value []Prm
	err := status.Result(&value)
	if err != nil {
		t.Fatal(err)
	}
}

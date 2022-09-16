package near_api_types

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Permission struct {
	String       string                  `json:"permission"`
	FunctionCall *FunctionCallPermission `json:"FunctionCall"`
}

type FunctionCallPermission struct {
	Allowance   string   `json:"allowance"`
	MethodNames []string `json:"method_names"`
	ReceiverId  string   `json:"receiver_id"`
}

func (p *Permission) UnmarshalJSON(data []byte) error {
	var base interface{}
	err := json.Unmarshal(data, &base)
	if err != nil {
		return err
	}
	item := reflect.ValueOf(base)
	switch item.Kind() {
	case reflect.String:
		p.String = item.String()
	case reflect.Map:
		fCallValue := item.MapIndex(reflect.ValueOf("FunctionCall"))
		if !fCallValue.IsNil() {
			p.FunctionCall = new(FunctionCallPermission)
			fMap, ok := fCallValue.Interface().(map[string]interface{})
			if !ok {
				return fmt.Errorf("Can't convert to map string interface()")
			}
			allowance, ok := fMap["allowance"]
			if ok {
				p.FunctionCall.Allowance = allowance.(string)
			}
			methodNames, ok := fMap["method_names"]
			if ok {
				array := []string{}
				methodNamesArray := methodNames.([]interface{})
				for _, value := range methodNamesArray {
					array = append(array, value.(string))
				}
				p.FunctionCall.MethodNames = array
			}
			receiver_id, ok := fMap["receiver_id"]
			if ok {
				p.FunctionCall.ReceiverId = receiver_id.(string)
			}
		}
	}
	return nil
}

package near_rpc

import (
	"context"
	"fmt"

	"github.com/nexeranet/gonear/jsonrpc"
	types "github.com/nexeranet/gonear/near_rpc/types"
)
type IRequestBase interface {
   Ctx(ctx context.Context) *Request
}

type IRequest interface {
    IRequestBase
    ICallMethod
    IClient
}

func NewRequest(request *jsonrpc.Request) *Request{
    return &Request{request}
}

type Request struct {
    request *jsonrpc.Request
}


func (a *Request) Ctx(ctx context.Context) *Request {
    a.request.Ctx(ctx)
    return a
}

func (a *Request) Call(method string, params ...interface{}) (*jsonrpc.RPCResponse, error) {
	response, err := a.request.Call(method, params...)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, types.ConvertError(response.Error)
	}
	return response, nil
}


// Call a contract method as a view function by block id.
func (a *Request) CallContractFuncByBlockId(accountId, method_name, args_base64 string, block_id uint64) (*types.ContractFuncResult, error) {
	type Params struct {
		RequestType string `json:"request_type"`
		BlockId     uint64 `json:"block_id"`
		AccountID   string `json:"account_id"`
		MethodName  string `json:"method_name"`
		ArgsBase64  string `json:"args_base64"`
	}
	params := Params{"call_function", block_id, accountId, method_name, args_base64}
	response, err := a.Call("query", &params)
	if err != nil {
		return nil, err
	}
	var raw types.ContractFuncResult
	err = response.GetObject(&raw)
	if err != nil {
		return nil, err
	}
	if raw.Error != "" {
		return nil, fmt.Errorf(raw.Error)
	}
	return &raw, nil
}

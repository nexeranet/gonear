package near_rpc

import (
	"context"
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


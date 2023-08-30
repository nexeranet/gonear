package jsonrpc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	Endpoint string
    Headers map[string]string
	Context  context.Context
    HttpClient *http.Client
}

func NewRequest(endpoint string, client *http.Client)  *Request{
   return &Request{
        Endpoint: endpoint,
        Headers: make(map[string]string),
        Context: context.Background(),
   }
}

func (r *Request) Call(method string, params ...interface{}) (*RPCResponse, error) {
	request := &RPCRequest{
		Method:  method,
		Params:  Params(params...),
		JSONRPC: jsonrpcVersion,
	}

	return r.doCall(r.Context, request)
}

func (r *Request) doCall(ctx context.Context, RPCRequest *RPCRequest) (*RPCResponse, error) {
	httpRequest, err := NewHttpRequest(ctx, RPCRequest, r.Endpoint, r.Headers)
	if err != nil {
		return nil, fmt.Errorf("rpc call %v() on %v: %v", RPCRequest.Method, r.Endpoint, err.Error())
	}
	httpResponse, err := r.HttpClient.Do(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("rpc call %v() on %v: %v", RPCRequest.Method, httpRequest.URL.String(), err.Error())
	}
	defer httpResponse.Body.Close()

	var rpcResponse *RPCResponse
	decoder := json.NewDecoder(httpResponse.Body)
	decoder.DisallowUnknownFields()
	decoder.UseNumber()
	err = decoder.Decode(&rpcResponse)

	// parsing error
	if err != nil {
		// if we have some http error, return it
		if httpResponse.StatusCode >= 400 {
			return nil, &HTTPError{
				Code: httpResponse.StatusCode,
				err:  fmt.Errorf("rpc call %v() on %v status code: %v. could not decode body to rpc response: %v", RPCRequest.Method, httpRequest.URL.String(), httpResponse.StatusCode, err.Error()),
			}
		}
		return nil, fmt.Errorf("rpc call %v() on %v status code: %v. could not decode body to rpc response: %v", RPCRequest.Method, httpRequest.URL.String(), httpResponse.StatusCode, err.Error())
	}

	// response body empty
	if rpcResponse == nil {
		// if we have some http error, return it
		if httpResponse.StatusCode >= 400 {
			return nil, &HTTPError{
				Code: httpResponse.StatusCode,
				err:  fmt.Errorf("rpc call %v() on %v status code: %v. rpc response missing", RPCRequest.Method, httpRequest.URL.String(), httpResponse.StatusCode),
			}
		}
		return nil, fmt.Errorf("rpc call %v() on %v status code: %v. rpc response missing", RPCRequest.Method, httpRequest.URL.String(), httpResponse.StatusCode)
	}

	return rpcResponse, nil
}

func (r *Request) Ctx(ctx context.Context) *Request {
	r.Context = ctx
	return r
}

// Json rpc client 2.0 (http)
//
//
package jsonrpc

//go:generate mockgen -source jsonrpc.go -destination mocks/jsonrpc.go.go

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
)

const (
	jsonrpcVersion = "2.0"
)

type RPCClient interface {
	Call(method string, params ...interface{}) (*RPCResponse, error)
	CallRaw(request *RPCRequest) (*RPCResponse, error)
	CallFor(out interface{}, method string, params ...interface{}) error
	CallBatch(requests RPCRequests) (RPCResponses, error)
	CallBatchRaw(requests RPCRequests) (RPCResponses, error)
}

type RPCRequest struct {
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
	ID      int         `json:"id"`
	JSONRPC string      `json:"jsonrpc"`
}

func NewRequest(method string, params ...interface{}) *RPCRequest {
	request := &RPCRequest{
		Method:  method,
		Params:  Params(params...),
		JSONRPC: jsonrpcVersion,
	}

	return request
}

type RPCResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result,omitempty"`
	Error   *RPCError   `json:"error,omitempty"`
	ID      int         `json:"id"`
}

type RPCErrorInfo struct {
	ErrorMessage string `json:"error_message,omitempty"`
}
type RPCErrorCause struct {
	Name string                 `json:"name"`
	Info map[string]interface{} `json:"info"`
}
type RPCError struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    interface{}   `json:"data,omitempty"`
	Name    string        `json:"name"`
	Cause   RPCErrorCause `json:"cause"`
}

func (e *RPCError) Error() string {
	return fmt.Sprintf("%d:%s[%v]", e.Code, e.Message, e.Data)
}

type HTTPError struct {
	Code int
	err  error
}

func (e *HTTPError) Error() string {
	return e.err.Error()
}

type rpcClient struct {
	endpoint      string
	httpClient    *http.Client
	customHeaders map[string]string
}

type RPCClientOpts struct {
	HTTPClient    *http.Client
	CustomHeaders map[string]string
}

type RPCResponses []*RPCResponse

func (res RPCResponses) AsMap() map[int]*RPCResponse {
	resMap := make(map[int]*RPCResponse)
	for _, r := range res {
		resMap[r.ID] = r
	}

	return resMap
}

func (res RPCResponses) GetByID(id int) *RPCResponse {
	for _, r := range res {
		if r.ID == id {
			return r
		}
	}

	return nil
}

func (res RPCResponses) HasError() bool {
	for _, res := range res {
		if res.Error != nil {
			return true
		}
	}
	return false
}

type RPCRequests []*RPCRequest

func NewClient(endpoint string) RPCClient {
	return NewClientWithOpts(endpoint, nil)
}
func NewClientWithOpts(endpoint string, opts *RPCClientOpts) RPCClient {
	rpcClient := &rpcClient{
		endpoint:      endpoint,
		httpClient:    &http.Client{},
		customHeaders: make(map[string]string),
	}

	if opts == nil {
		return rpcClient
	}

	if opts.HTTPClient != nil {
		rpcClient.httpClient = opts.HTTPClient
	}

	if opts.CustomHeaders != nil {
		for k, v := range opts.CustomHeaders {
			rpcClient.customHeaders[k] = v
		}
	}

	return rpcClient
}

func (client *rpcClient) Call(method string, params ...interface{}) (*RPCResponse, error) {

	request := &RPCRequest{
		Method:  method,
		Params:  Params(params...),
		JSONRPC: jsonrpcVersion,
	}

	return client.doCall(request)
}

func (client *rpcClient) CallRaw(request *RPCRequest) (*RPCResponse, error) {

	return client.doCall(request)
}

func (client *rpcClient) CallFor(out interface{}, method string, params ...interface{}) error {
	rpcResponse, err := client.Call(method, params...)
	if err != nil {
		return err
	}

	if rpcResponse.Error != nil {
		return rpcResponse.Error
	}

	return rpcResponse.GetObject(out)
}

func (client *rpcClient) CallBatch(requests RPCRequests) (RPCResponses, error) {
	if len(requests) == 0 {
		return nil, errors.New("empty request list")
	}

	for i, req := range requests {
		req.ID = i
		req.JSONRPC = jsonrpcVersion
	}

	return client.doBatchCall(requests)
}

func (client *rpcClient) CallBatchRaw(requests RPCRequests) (RPCResponses, error) {
	if len(requests) == 0 {
		return nil, errors.New("empty request list")
	}

	return client.doBatchCall(requests)
}

func (client *rpcClient) newRequest(req interface{}) (*http.Request, error) {

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", client.endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	// set default headers first, so that even content type and accept can be overwritten
	for k, v := range client.customHeaders {
		request.Header.Set(k, v)
	}

	return request, nil
}

func (client *rpcClient) doCall(RPCRequest *RPCRequest) (*RPCResponse, error) {

	httpRequest, err := client.newRequest(RPCRequest)
	if err != nil {
		return nil, fmt.Errorf("rpc call %v() on %v: %v", RPCRequest.Method, client.endpoint, err.Error())
	}
	httpResponse, err := client.httpClient.Do(httpRequest)
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

func (client *rpcClient) doBatchCall(rpcRequest []*RPCRequest) ([]*RPCResponse, error) {
	httpRequest, err := client.newRequest(rpcRequest)
	if err != nil {
		return nil, fmt.Errorf("rpc batch call on %v: %v", client.endpoint, err.Error())
	}
	httpResponse, err := client.httpClient.Do(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("rpc batch call on %v: %v", httpRequest.URL.String(), err.Error())
	}
	defer httpResponse.Body.Close()

	var rpcResponse RPCResponses
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
				err:  fmt.Errorf("rpc batch call on %v status code: %v. could not decode body to rpc response: %v", httpRequest.URL.String(), httpResponse.StatusCode, err.Error()),
			}
		}
		return nil, fmt.Errorf("rpc batch call on %v status code: %v. could not decode body to rpc response: %v", httpRequest.URL.String(), httpResponse.StatusCode, err.Error())
	}

	// response body empty
	if rpcResponse == nil || len(rpcResponse) == 0 {
		// if we have some http error, return it
		if httpResponse.StatusCode >= 400 {
			return nil, &HTTPError{
				Code: httpResponse.StatusCode,
				err:  fmt.Errorf("rpc batch call on %v status code: %v. rpc response missing", httpRequest.URL.String(), httpResponse.StatusCode),
			}
		}
		return nil, fmt.Errorf("rpc batch call on %v status code: %v. rpc response missing", httpRequest.URL.String(), httpResponse.StatusCode)
	}

	return rpcResponse, nil
}

func Params(params ...interface{}) interface{} {
	var finalParams interface{}

	// if params was nil skip this and p stays nil
	if params != nil {
		switch len(params) {
		case 0: // no parameters were provided, do nothing so finalParam is nil and will be omitted
		case 1: // one param was provided, use it directly as is, or wrap primitive types in array
			if params[0] != nil {
				var typeOf reflect.Type

				// traverse until nil or not a pointer type
				for typeOf = reflect.TypeOf(params[0]); typeOf != nil && typeOf.Kind() == reflect.Ptr; typeOf = typeOf.Elem() {
				}

				if typeOf != nil {
					// now check if we can directly marshal the type or if it must be wrapped in an array
					switch typeOf.Kind() {
					// for these types we just do nothing, since value of p is already unwrapped from the array params
					case reflect.Struct:
						finalParams = params[0]
					case reflect.Array:
						finalParams = params[0]
					case reflect.Slice:
						finalParams = params[0]
					case reflect.Interface:
						finalParams = params[0]
					case reflect.Map:
						finalParams = params[0]
					default: // everything else must stay in an array (int, string, etc)
						finalParams = params
					}
				}
			} else {
				finalParams = params
			}
		default: // if more than one parameter was provided it should be treated as an array
			finalParams = params
		}
	}

	return finalParams
}

func (RPCResponse *RPCResponse) GetInt() (int64, error) {
	val, ok := RPCResponse.Result.(json.Number)
	if !ok {
		return 0, fmt.Errorf("could not parse int64 from %s", RPCResponse.Result)
	}

	i, err := val.Int64()
	if err != nil {
		return 0, err
	}

	return i, nil
}

func (RPCResponse *RPCResponse) GetFloat() (float64, error) {
	val, ok := RPCResponse.Result.(json.Number)
	if !ok {
		return 0, fmt.Errorf("could not parse float64 from %s", RPCResponse.Result)
	}

	f, err := val.Float64()
	if err != nil {
		return 0, err
	}

	return f, nil
}

func (RPCResponse *RPCResponse) GetBool() (bool, error) {
	val, ok := RPCResponse.Result.(bool)
	if !ok {
		return false, fmt.Errorf("could not parse bool from %s", RPCResponse.Result)
	}

	return val, nil
}

func (RPCResponse *RPCResponse) GetString() (string, error) {
	val, ok := RPCResponse.Result.(string)
	if !ok {
		return "", fmt.Errorf("could not parse string from %s", RPCResponse.Result)
	}

	return val, nil
}

func (RPCResponse *RPCResponse) GetObject(toType interface{}) error {
	js, err := json.Marshal(RPCResponse.Result)
	if err != nil {
		return err
	}

	err = json.Unmarshal(js, toType)
	if err != nil {
		return err
	}

	return nil
}

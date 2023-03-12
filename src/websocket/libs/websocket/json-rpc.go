package websocket

import (
	"encoding/json"
)

// JSON-RPC request and response structs
type JsonRPCRequest struct {
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
	ID     string          `json:"id"`
}

type JsonRPCResponse struct {
	ID     string          `json:"id"`
	Result json.RawMessage `json:"result"`
	Error  interface{}     `json:"error"`
}

func (r *JsonRPCRequest) GetParams() map[string]interface{} {
	var params map[string]interface{}
	json.Unmarshal(r.Params, &params)
	return params
}

func GenerateRequest(method string, params map[string]interface{}, ID string) ([]byte, error) {
	paramsBytes, _ := json.Marshal(params)
	jsonRPC := JsonRPCRequest{
		Method: method,
		Params: paramsBytes,
		ID:     ID,
	}
	return json.Marshal(jsonRPC)
}

func ResponseBuilder(id string, params json.RawMessage) JsonRPCResponse {
	return JsonRPCResponse{
		ID:     id,
		Result: params,
	}
}

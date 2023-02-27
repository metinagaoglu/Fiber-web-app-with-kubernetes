package websocket

import (
	"encoding/json"
)

// JSON-RPC request and response structs
type JsonRPCRequest struct {
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
	ID     uint64          `json:"id"`
}

type JsonRPCResponse struct {
	ID     uint64          `json:"id"`
	Result json.RawMessage `json:"result"`
	Error  interface{}     `json:"error"`
}

func ResponseBuilder(id uint64, params json.RawMessage) JsonRPCResponse {
	return JsonRPCResponse{
		ID:     id,
		Result: params,
	}
}

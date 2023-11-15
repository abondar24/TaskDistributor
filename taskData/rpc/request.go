package rpc

// TaskRPCRequest request for json-rpc api
type TaskRPCRequest struct {
	// json-rpc version 2
	JSONRPC string `json:"jsonrpc"`

	// json-rpc method: TaskRPC.GetTask, TaskRPC.GetTaskHistory
	Method string `json:"method"`

	// params
	Params []string `json:"params"`

	//request id
	ID int `json:"id"`
}

// TaskRPCStatusRequest request for json-rpc api with status arguments
type TaskRPCStatusRequest struct {
	// json-rpc version 2
	JSONRPC string `json:"jsonrpc"`

	// json-rpc method: TaskRPC.GetTasksByStatus,
	Method string `json:"method"`

	// params
	Params []StatusArgs `json:"params"`

	//request id
	ID int `json:"id"`
}

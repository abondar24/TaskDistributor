package client

import (
	"bytes"
	"github.com/abondar24/TaskDistributor/taskData/config"
	"github.com/abondar24/TaskDistributor/taskData/data"
	args "github.com/abondar24/TaskDistributor/taskData/rpc"
	"github.com/gorilla/rpc/v2/json"
	"log"
	"net/http"

	"strconv"
)

type TaskRpcClient struct {
	uri    string
	client *http.Client
}

const (
	GetTask          string = "TaskRPC.GetTask"
	GetTasksByStatus string = "TaskRPC.GetTasksByStatus"

	GetTaskHistory string = "TaskRPC.GetTaskHistory"
)

func NewTaskRpcClient(conf *config.Config) *TaskRpcClient {
	uri := "http://" + conf.RPC.Host + ":" + strconv.Itoa(conf.RPC.Port) + "/rpc"
	client := &http.Client{}
	return &TaskRpcClient{
		uri,
		client,
	}
}

func (cl *TaskRpcClient) GetTask(id *string) (*data.Task, error) {
	resp, err := cl.sendRequest(GetTask, id)
	if err != nil {
		log.Println("Error sending HTTP request:", err)
		return nil, err
	}

	defer resp.Body.Close()

	var response data.Task
	err = json.DecodeClientResponse(resp.Body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (cl *TaskRpcClient) GetTasksByStatus(args *args.StatusArgs) (*[]*data.Task, error) {
	resp, err := cl.sendRequest(GetTasksByStatus, args)
	if err != nil {
		log.Println("Error sending HTTP request:", err)
		return nil, err
	}

	defer resp.Body.Close()

	var response []*data.Task
	err = json.DecodeClientResponse(resp.Body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (cl *TaskRpcClient) GetTaskHistory(id *string) (*data.TaskHistory, error) {

	resp, err := cl.sendRequest(GetTaskHistory, id)
	if err != nil {
		log.Println("Error sending HTTP request:", err)
		return nil, err
	}

	defer resp.Body.Close()

	var response data.TaskHistory
	err = json.DecodeClientResponse(resp.Body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (cl *TaskRpcClient) sendRequest(method string, params interface{}) (*http.Response, error) {

	requestJSON, err := json.EncodeClientRequest(method, params)
	if err != nil {
		return nil, err
	}

	log.Println("Sending request to store:", string(requestJSON))

	req, err := http.NewRequest("POST", cl.uri, bytes.NewBuffer(requestJSON))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return cl.client.Do(req)
}

package handler

import (
	"context"
	"encoding/json"
	"github.com/abondar24/TaskDisrtibutor/taskApi/model"
	"github.com/gorilla/mux"
	"net/http"
)

func decodeTaskRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if ok {
		req.ID = id
	}

	return req, nil
}

func decodeHealthRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

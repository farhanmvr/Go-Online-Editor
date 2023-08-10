package handlers

import (
	"encoding/json"
	"github.com/farhanmvr/go-editor/utils"
	"net/http"
	"time"
)

type ExecuteResponse struct {
	Status    *string    `json:"status"`
	Timestamp *time.Time `json:"timestamp"`
	Output    *string    `json:"output,omitempty"`
	Error     *string    `json:"error,omitempty"`
}

type ExecuteRequest struct {
	Code *string `json:"code"`
}

func ExecuteCodeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	time := time.Now()

	// Response contract
	var response ExecuteResponse
	var status string
	response.Status = &status
	response.Timestamp = &time

	// Parse the request and validate
	var request ExecuteRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil || request.Code == nil || *request.Code == "" {
		status = "error"
		errorMsg := "invalid request"
		response.Error = &errorMsg
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	result, err := utils.RunGoCode(*request.Code)
	if err != nil {
		status = "error"
		errorMsg := "Something went wrong, please try again later"
		response.Error = &errorMsg
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}
	if result.Error != nil {
		status = "fail"
		response.Error = result.Error
	} else {
		status = "success"
	}
	response.Output = result.Output
	json.NewEncoder(w).Encode(response)
}

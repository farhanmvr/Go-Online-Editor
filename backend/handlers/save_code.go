package handlers

import (
	"encoding/json"
	"github.com/farhanmvr/go-editor/db"
	"github.com/farhanmvr/go-editor/models"
	"github.com/farhanmvr/go-editor/utils"
	"net/http"
)

type SaveCodeRequest struct {
	Code *string `json:"code"`
	Name *string `json:"name"`
}

type SaveCodeResponse struct {
	Status      *string             `json:"status"`
	CodeSnippet *models.CodeSnippet `json:"code_snippet,omitempty"`
	Error       *string             `json:"error,omitempty"`
}

// This handler will save code into db [Only successfull compiled codes are stored in db]
func SaveCodeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Response contract
	var response SaveCodeResponse
	var status string
	response.Status = &status

	var request SaveCodeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil ||
		request.Code == nil || *request.Code == "" || request.Name == nil ||
		*request.Name == "" {
		status = "fail"
		errorMsg := "invalid request, please provide code and name for saving"
		response.Error = &errorMsg
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Run code and test
	result, err := utils.RunGoCode(*request.Code)
	if err != nil {
		status = "error"
		errorMsg := "Something went wrong, please try again later"
		response.Error = &errorMsg
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}
	// Return in case of any error (not storing in db)
	if result.Error != nil {
		status = "fail"
		response.Error = result.Output
		json.NewEncoder(w).Encode(response)
		return
	} else {
		status = "success"
	}

	codeSnippet, err := db.InsertCodeSnippet(*request.Code, *request.Name, status)
	if err != nil {
		status = "error"
		errorMsg := "Something went wrong, please try again later"
		response.Error = &errorMsg
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}
	response.CodeSnippet = codeSnippet

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

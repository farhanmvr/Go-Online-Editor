package handlers

import (
	"encoding/json"
	"github.com/farhanmvr/go-editor/db"
	"github.com/farhanmvr/go-editor/models"
	"net/http"
)

type GetAllCodeSnippetsResponse struct {
	Status       *string                `json:"status"`
	CodeSnippets *[]*models.CodeSnippet `json:"code_snippets,omitempty"`
	Error        *string                `json:"error,omitempty"`
}

// Get all code snippets from db
func GetAllCodeSnippetsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response GetAllCodeSnippetsResponse
	var status string
	response.Status = &status

	snippets, err := db.GetAllCodeSnippets()
	if err != nil {
		status = "error"
		errorMsg := "Something went wrong, please try again later"
		response.Error = &errorMsg
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}
	status = "success"
	response.CodeSnippets = &snippets

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

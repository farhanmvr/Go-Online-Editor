package handlers

import (
	"encoding/json"
	"github.com/farhanmvr/go-editor/db"
	"github.com/gorilla/mux"
	"net/http"
)

// This handler will delete a code snippet by id given
func DeleteCodeSnippetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	snippetID := vars["id"]

	rows, err := db.DeleteSnippetById(snippetID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"status": "error",
			"error":  "Something went wrong, please try again later",
		})
		return
	}

	if rows < 1 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"status": "fail",
			"error":  "No record found for the given id",
		})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

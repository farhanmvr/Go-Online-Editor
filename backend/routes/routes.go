package routes

import (
	"github.com/farhanmvr/go-editor/handlers"
	"github.com/gorilla/mux"
)

// All routes to the server
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/code/execute", handlers.ExecuteCodeHandler).Methods("POST")
	router.HandleFunc("/api/v1/code/save", handlers.SaveCodeHandler).Methods("POST")
	router.HandleFunc("/api/v1/code/snippets", handlers.GetAllCodeSnippetsHandler).Methods("GET")
	router.HandleFunc("/api/v1/code/snippets/{id}", handlers.DeleteCodeSnippetHandler).Methods("DELETE")

	return router
}

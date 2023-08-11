package models

import "time"

// CodeSnippet contract
type CodeSnippet struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	Status      string    `json:"status"`
	DateCreated time.Time `json:"date_created"`
}

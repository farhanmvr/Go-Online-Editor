package utils

import (
	"os"
	"os/exec"
	"strings"
)

type GoRunResponse struct {
	Output *string `json:"output,omitempty"`
	Error  *string `json:"error,omitempty"`
}

// Helper function to run go code
func RunGoCode(code string) (*GoRunResponse, error) {
	// Create a temporary file to save the code
	tempFile, err := os.CreateTemp("", "temp-*.go")
	if err != nil {
		return nil, err
	}
	// Clean up
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	_, err = tempFile.WriteString(code)
	if err != nil {
		return nil, err
	}

	var result GoRunResponse

	// Run the code
	cmd := exec.Command("go", "run", tempFile.Name())
	opt, err := cmd.CombinedOutput()
	if err != nil {
		errorMsg := err.Error()
		result.Error = &errorMsg
	}

	output := string(opt)
	// Replace file path with main.go
	output = strings.Replace(output, tempFile.Name(), "main.go", -1)
	result.Output = &output

	return &result, nil
}

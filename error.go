package fingerspot

import (
	"encoding/json"
	"fmt"
)

type Err struct {
	Success   bool   `json:"success,omitempty"`
	ErrorCode string `json:"error_code"`
	Message   string `json:"message"`
}

func (e Err) Error() string {
	err, _ := json.Marshal(e)
	return fmt.Sprintf("Error Response: %s\n", string(err))
}

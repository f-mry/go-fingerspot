package fingerspot

import (
	"encoding/json"
)

type Err struct {
	Success   bool   `json:"success"`
	ErrorCode string `json:"error_code"`
	Message   string `json:"message"`
}

func (e Err) Error() string {
	err, _ := json.Marshal(e)
	return string(err)
}

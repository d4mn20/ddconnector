package utils

import (
	"encoding/json"
	"fmt"
)

func GetErrorMessageFromBody(body []byte) string {
	var errorMsg map[string]interface{}
	if err := json.Unmarshal(body, &errorMsg); err == nil {
		if msg, exists := errorMsg["detail"]; exists {
			return fmt.Sprintf("%v", msg)
		}
		if msg, exists := errorMsg["error"]; exists {
			return fmt.Sprintf("%v", msg)
		}
	}
	return "unknown error"
}

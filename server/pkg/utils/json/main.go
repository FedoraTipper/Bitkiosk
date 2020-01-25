package json

import (
	"encoding/json"
	"github.com/fedoratipper/bitkiosk/server/internal/logger"
)

func ObjectToJsonString(obj interface{}) string{
	bytes, err := json.Marshal(obj)

	if err != nil {
		logger.Error("Unable to parse obj")
		return ""
	}

	return string(bytes)
}
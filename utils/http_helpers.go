package utils

import (
	"bytes"
	"encoding/json"
	"log"
)

func PrettifyJSON(data string) string {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, []byte(data), "", "\t")
	if err != nil {
		log.Panic("JSON parse error")
	}
	return prettyJSON.String()
}

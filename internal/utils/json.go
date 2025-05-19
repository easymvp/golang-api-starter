package utils

import (
	"encoding/json"
	"strings"
)

func ToJSON(data interface{}) []byte {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return bytes
}

func ToJSONText(data interface{}) string {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func FromJSON(bytes []byte, data interface{}) error {
	err := json.Unmarshal(bytes, data)
	return err
}

func FromJSONText(text string, data interface{}) error {
	err := json.Unmarshal([]byte(text), data)
	return err
}

func ConvertObject(from interface{}, data interface{}) error {
	bytes := ToJSON(from)
	err := FromJSON(bytes, data)
	return err
}

func IsJSON(data []byte) bool {
	var js json.RawMessage
	return json.Unmarshal(data, &js) == nil
}

func IsJSONText(data string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(data), &js) == nil
}

func RetrieveJSON(data string) string {
	// Find the first opening brace
	startIndex := strings.Index(data, "{")
	if startIndex < 0 {
		return "{}"
	}

	// Find the last closing brace
	endIndex := strings.LastIndex(data, "}")
	if endIndex < 0 {
		return "{}"
	}

	jsonStr := data[startIndex : endIndex+1]
	jsonStr = strings.ReplaceAll(jsonStr, "\n", "")
	jsonStr = strings.ReplaceAll(jsonStr, "}{", ",")
	return jsonStr
}

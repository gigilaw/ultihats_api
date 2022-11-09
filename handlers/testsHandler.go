package handlers

import (
	"bytes"
	"encoding/json"
)

func DecodeResponse(body *bytes.Buffer) map[string]interface{} {
	res := make(map[string]interface{})
	json.NewDecoder(body).Decode(&res)
	return res
}

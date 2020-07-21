package models

import (
	"encoding/json"
)

type Message struct {
	Code int    `json:"code"`
	Data string `json: "data"`
}

func SetValueMessage(code int, message string) ([]byte, error) {
	messageRaw := Message{
		Code: code,
		Data: message,
	}
	messageJSON, err := json.Marshal(messageRaw)
	if err != nil {
		return nil, err
	}
	return messageJSON, nil
}

package fingerspot

import (
	"encoding/json"
	"io"
	"net/http"
)

type Webhook struct {
	handler http.Handler
	Message chan interface{}
}

func (w *Webhook) Server() http.Handler {
	return w.handler
}

func NewWebhookWithServer() Webhook {
	var wh Webhook
	handler := http.NewServeMux()
	handler.HandleFunc("/", wh.MessageHandler)
	wh.handler = handler
	wh.Message = make(chan interface{})
	return wh
}

func (wh *Webhook) MessageHandler(_ http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		wh.Message <- Err{
			ErrorCode: "webhook message handler",
			Message:   "failed to parse raw message",
		}
	}
	wh.Message <- getMessage(body)
}

func getMessage(rawMessage []byte) interface{} {
	var data map[string]interface{}
	err := json.Unmarshal(rawMessage, &data)
	if err != nil {
		return Err{
			ErrorCode: "webhook parsing",
			Message:   "invalid payload",
		}
	}
	return decodeMessage(rawMessage, data)
}

func decodeMessage(rawMessage []byte, data map[string]interface{}) interface{} {
	messageType, ok := data["type"]
	if !ok {
		return Err{
			ErrorCode: "webhook decode message",
			Message:   "fail to get message type",
		}
	}

	messageKey, ok := messageType.(string)
	if !ok {
		return Err{
			ErrorCode: "webhook decode message",
			Message:   "invalid message type",
		}
	}

	switch messageKey {
	case "attlog":
		message := new(RealtimeAttLogMessage)
		err := json.Unmarshal(rawMessage, message)
		if err != nil {
			return Err{
				ErrorCode: "webhook decode casting",
				Message:   err.Error(),
			}
		}
		return message
	case "get_userinfo":
		message := new(GetUserInfoMessage)
		err := json.Unmarshal(rawMessage, message)
		if err != nil {
			return Err{
				ErrorCode: "webhook decode casting",
				Message:   err.Error(),
			}
		}
		return message
	case "set_userinfo":
		message := new(SetUserInfoMessage)
		err := json.Unmarshal(rawMessage, message)
		if err != nil {
			return Err{
				ErrorCode: "webhook decode casting",
				Message:   err.Error(),
			}
		}
		return message
	case "delete_userinfo":
		message := new(DeleteUserInfoMessage)
		err := json.Unmarshal(rawMessage, message)
		if err != nil {
			return Err{
				ErrorCode: "webhook decode casting",
				Message:   err.Error(),
			}
		}
		return message
	default:
		return data
	}
}

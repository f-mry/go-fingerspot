package fingerspot

import (
	"fmt"
	"io"
	"net/http"
)

type WebhookServer struct {
	handler http.Handler
}

func (w *WebhookServer) Server() http.Handler {
	return w.handler
}

func handleMessage(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}

func NewWebhookServer() WebhookServer {
	handler := http.NewServeMux()
	handler.HandleFunc("/", handleMessage)
	return WebhookServer{handler: handler}
}

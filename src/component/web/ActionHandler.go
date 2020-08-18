package web

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Action func(response http.ResponseWriter, request *http.Request) ActionResultInterface

func HandleAPIAction(action Action, response http.ResponseWriter, request *http.Request) {
	new(actionHandler).Handle(action, response, request)
}

type actionHandler struct {
}

func (h *actionHandler) Handle(action Action, response http.ResponseWriter, request *http.Request) {
	content := action(response, request)
	statusCode := h.getStatusCode(content.Error(), content.StatusCode())

	log.Print(content.Error())

	response.WriteHeader(statusCode)
	response.Header().Set("Content-EntityType", "application/json")

	serializerContent, _ := json.Marshal(content.Content())
	_, err := io.WriteString(response, string(serializerContent))
	log.Print(err)
}

func (h *actionHandler) getStatusCode(err error, forcedCode int) int {
	statusCode := http.StatusOK

	if forcedCode != 0 {
		statusCode = forcedCode
	} else if err != nil {
		statusCode = http.StatusBadRequest
	}

	return statusCode
}

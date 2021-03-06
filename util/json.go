package util

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// SendJSON marshals entity to a json struct and sends appropriate headers to writer
func SendJSON(writer http.ResponseWriter, request *http.Request, entity interface{}, code int) {
	writer.Header().Set("Content-Type", "application/json")

	encodedEntity, err := json.Marshal(entity)

	if err != nil {
		log.Print(fmt.Sprintf("Error while encoding JSON: %entity", entity))
		writer.WriteHeader(http.StatusInternalServerError)
		io.WriteString(writer, `{"error": "Internal server error"}`)
	} else {
		writer.WriteHeader(code)
		io.WriteString(writer, string(encodedEntity))
	}
}

package utils

import (
    "log"
    "net/http"
)

func ErrorResponse(w http.ResponseWriter, message string, statusCode int) {
    w.WriteHeader(statusCode)
    w.Write([]byte(message))
    log.Printf("Error %d: %s", statusCode, message)
}

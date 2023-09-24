package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type jsonWriter struct {
	log ILogger
	w   http.ResponseWriter
}

func (jw jsonWriter) Write(v any) {
	content, err := json.Marshal(v)
	statusCode := http.StatusOK
	if nil == err {
		jw.w.Header().Set("Content-Type", "application/json")
	} else {
		jw.log.Error(fmt.Sprintf("JSON marshal error: %v", err))
		statusCode = http.StatusInternalServerError
		content = []byte(err.Error())
	}
	jw.w.WriteHeader(statusCode)
	_, writeErr := jw.w.Write(content)
	if writeErr != nil {
		jw.log.Error(fmt.Sprintf("response write error: %v", err))
	}
}

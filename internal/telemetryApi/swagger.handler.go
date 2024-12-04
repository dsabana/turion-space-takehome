package telemetryApi

import (
	"bytes"
	"net/http"
)

func getDocumentation(w http.ResponseWriter, r *http.Request) {
	schema := bytes.Replace(Schema, []byte("3.1.0"), []byte("3.0.3"), 1)

	if _, err := w.Write(schema); err != nil {
		generateErrorResponse(w, http.StatusInternalServerError, ErrEncodingResponse)
		return
	}
}

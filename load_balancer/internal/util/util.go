package util

import (
	"net/http"
	"net/http/httptest"
)

func CopyHeadersAndBody(w http.ResponseWriter, recorder *httptest.ResponseRecorder) {
	headers := w.Header()
	for k, vs := range recorder.Header() {
		headers[k] = vs
	}
	w.WriteHeader(recorder.Code)
	recorder.Body.WriteTo(w)
}

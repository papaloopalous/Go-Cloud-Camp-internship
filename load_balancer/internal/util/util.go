package util

import (
	"crypto/sha256"
	"encoding/hex"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
)

func CopyHeadersAndBody(w http.ResponseWriter, recorder *httptest.ResponseRecorder) {
	headers := w.Header()
	for k, vs := range recorder.Header() {
		headers[k] = vs
	}
	w.WriteHeader(recorder.Code)
	recorder.Body.WriteTo(w)
}

func HashIP(ip, salt string) string {
	hasher := sha256.New()
	hasher.Write([]byte(salt + ip))
	return hex.EncodeToString(hasher.Sum(nil))
}

func GetClientIP(r *http.Request) string {
	if realIP := r.Header.Get("X-Real-IP"); realIP != "" {
		return realIP
	}
	if fwd := r.Header.Get("X-Forwarded-For"); fwd != "" {
		ips := strings.Split(fwd, ",")
		return strings.TrimSpace(ips[0])
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

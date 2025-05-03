package handler

import (
	ratelimiter "load_balancer/rate_limiter"
	"net/http"
	"strconv"
)

type LimiterHandler struct {
	Limiter ratelimiter.BucketIface
}

// обработчик установки rate
func (lh *LimiterHandler) SetRateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := r.URL.Query().Get("ip")
		valStr := r.URL.Query().Get("value")

		if ip == "" || valStr == "" {
			http.Error(w, "Missing 'ip' or 'value' parameter", http.StatusBadRequest)
			return
		}

		rate, err := strconv.Atoi(valStr)
		if err != nil {
			http.Error(w, "Invalid 'value' parameter", http.StatusBadRequest)
			return
		}

		if err := lh.Limiter.SetRate(ip, rate); err != nil {
			http.Error(w, "Failed to set rate: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Rate updated"))
	}
}

// обработчик установки max tokens
func (lh *LimiterHandler) SetMaxHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := r.URL.Query().Get("ip")
		valStr := r.URL.Query().Get("value")

		if ip == "" || valStr == "" {
			http.Error(w, "Missing 'ip' or 'value' parameter", http.StatusBadRequest)
			return
		}

		max, err := strconv.Atoi(valStr)
		if err != nil {
			http.Error(w, "Invalid 'value' parameter", http.StatusBadRequest)
			return
		}

		if err := lh.Limiter.SetMaxTokens(ip, max); err != nil {
			http.Error(w, "Failed to set max tokens: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Max tokens updated"))
	}
}

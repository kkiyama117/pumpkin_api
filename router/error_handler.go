package router

import (
	"net/http"

	"github.com/goware/httplog"
)

type BetterHandler func(http.ResponseWriter, *http.Request) error

// cache error and return HandlerFunc
func ErrorHandler(f BetterHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		logger := httplog.LogEntry(r.Context())
		if err != nil {

			http.Error(w, err.Error(), http.StatusInternalServerError)
			logger.Error().Msg(err.Error())
		} else {
			logger.Info().Msg("No error cached at ErrorHandler")
		}
	}
}

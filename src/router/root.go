package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/docgen"
	"hintan.jp/pumpkin_api/src/entities"
	"hintan.jp/pumpkin_api/src/pkg/version"
	"hintan.jp/pumpkin_api/src/utils"
)

func router() chi.Router {
	r := chi.NewRouter()
	r.Get("/", utils.ErrorHandler(rootFunc()))
	r.Get("/health_check", utils.ErrorHandler(statusFunc()))
	r.Get("/routes", utils.ErrorHandler(getRoutes(r)))
	return r
}

func rootFunc() func(w http.ResponseWriter, request *http.Request) error {
	return func(w http.ResponseWriter, request *http.Request) error {
		_, err := fmt.Fprint(w, "")
		if err != nil {
			return fmt.Errorf("error: %v", err)
		}
		return nil
	}
}

// docs
func getRoutes(router chi.Router) func(w http.ResponseWriter, request *http.Request) error {
	return func(w http.ResponseWriter, request *http.Request) error {
		if _, err := fmt.Fprint(w, docgen.JSONRoutesDoc(router)); err != nil {
			return fmt.Errorf("error: %v", err)
		}
		return nil
	}
}

func statusFunc() func(w http.ResponseWriter, request *http.Request) error {
	return func(w http.ResponseWriter, request *http.Request) error {
		currentStatus := entities.ServerStatus{Code: 0, Version: version.VERSION, Value: "OK"}
		// To JSON Format
		currentStatusJson, _ := json.Marshal(currentStatus)
		_, err := fmt.Fprint(w, string(currentStatusJson))
		if err != nil {
			return fmt.Errorf("error: %v", err)
		}
		return nil
	}
}

var Router = router

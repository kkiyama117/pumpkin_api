package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/docgen"
	"hintan.jp/pumpkin_api/entities"
)

func router() chi.Router {
	r := chi.NewRouter()
	r.Get("/", ErrorHandler(rootFunc()))
	r.Get("/status", ErrorHandler(statusFunc()))
	r.Get("/routes", ErrorHandler(getRoutes(r)))
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
		currentStatus := entities.ServerStatus{Code: 0, Value: "OK"}
		_, err := fmt.Fprint(w, currentStatus)
		if err != nil {
			return fmt.Errorf("error: %v", err)
		}
		return nil
	}
}

var Router = router

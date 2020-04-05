package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/docgen"
)

func router() chi.Router {
	r := chi.NewRouter()
	r.Get("/", ErrorHandler(rootFunc()))
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

var Router = router

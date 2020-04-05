package router

import (
	"github.com/go-chi/chi"
)

func debugRouter(router chi.Router) chi.Router {
	r := chi.NewRouter()
	r.Get("/", ErrorHandler(getRoutes(router)))
	return r
}

var DebugRouter = debugRouter

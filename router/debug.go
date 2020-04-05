package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func debugRouter(router chi.Router) chi.Router {
	r := chi.NewRouter()
	r.Get("/", ErrorHandler(getRoutes(router)))
	r.Mount("/profiler", middleware.Profiler())

	return r
}

var DebugRouter = debugRouter

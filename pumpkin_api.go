package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/goware/httplog"
	"hintan.jp/pumpkin_api/middlewares"
	"hintan.jp/pumpkin_api/router"
)

func main() {
	r := chi.NewRouter()
	// must be before router
	setMiddleware(r)

	// set port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Use(middleware.Heartbeat("/ping"))
	r.Mount("/", router.Router())
	r.Mount("/debug", router.DebugRouter(r))

	// Passing -router to the program will generate docs for the above
	// router definition. See the `router.json` file in this folder for
	// the output.
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func setMiddleware(router chi.Router) {
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RedirectSlashes)
	router.Use(middleware.RequestID)
	router.Use(middlewares.CorsMiddleware("localhost:3000"))
	// Logger
	logger := httplog.NewLogger("httplog-example", httplog.Options{
		JSON: true,
	})
	router.Use(httplog.RequestLogger(logger))
}

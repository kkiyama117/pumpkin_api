package factories

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/goware/httplog"
	"hintan.jp/pumpkin_api/src/middlewares"
	"hintan.jp/pumpkin_api/src/router"
)

type server struct {
	router chi.Router
}

type Server interface {
	// call http.ListenAndServe
	Run() error
}

func NewServer() Server {
	// create instance
	sv := &server{}
	// initialize
	sv = Inject(sv)
	// return injected web server
	return sv
}

// inject router and handler and usecase
func Inject(server *server) *server {
	server.router = chi.NewRouter()
	// Initialize web
	err := InjectMiddleware(server)
	if err != nil {
		panic("error with inject middleware")
	}
	// set router
	InjectRouter(server)
	return server
}

func InjectMiddleware(server *server) error {

	if server.router == nil {
		return errors.New("nil web, please set web instance")
	}

	server.router.Use(middleware.RedirectSlashes)
	server.router.Use(middleware.RequestID)
	server.router.Use(middleware.Logger)
	server.router.Use(middleware.Recoverer)
	server.router.Use(middlewares.CorsMiddleware("localhost:3000"))
	// for ping
	server.router.Use(middleware.Heartbeat("/ping"))
	server.router.Use(middleware.URLFormat)
	server.router.Use(render.SetContentType(render.ContentTypeJSON))
	// Logger
	logger := httplog.NewLogger("httplog-example", httplog.Options{
		JSON: true,
	})
	server.router.Use(httplog.RequestLogger(logger))
	return nil
}

func InjectRouter(server *server) {
	server.router.Mount("/", router.Router())
	server.router.Mount("/debug", router.DebugRouter(server.router))
}
func (server *server) Run() error {
	// return error
	config := GetConfigs()
	portString := ":" + config.GetValue("Port")
	return http.ListenAndServe(portString, server.router)
}

package server

import (
	"log"
	"os"

	"github.com/cleitonmarx/gowebapp/config"
	"github.com/cleitonmarx/gowebapp/server/controllers"
	"github.com/codegangsta/negroni"
	"github.com/dimfeld/httptreemux"
)

//Server encapsulates a HTTP server
type Server struct {
	configuration config.EnvironmentConfig
	router        *httptreemux.TreeMux
	httpServer    *negroni.Negroni
	logger        *log.Logger
}

func (s *Server) Init() {

	s.logger = log.New(os.Stdout, "[GoWebApp] ", 0)
	mainController := controllers.NewMainController()

	s.router = httptreemux.New()
	s.router.GET("/", mainController.GetHandler)
	s.router.GET("/_version", mainController.GetVersionHandler)

	s.httpServer = negroni.Classic()
	s.httpServer.UseHandler(s.router)
}

func (s *Server) Run() {
	s.logger.Printf("Environment: %s", s.configuration.Name)

	s.httpServer.Run(s.configuration.HTTPServer.GetFormatedAddress())
}

func New(environmentConfig config.EnvironmentConfig) *Server {
	return &Server{configuration: environmentConfig}
}

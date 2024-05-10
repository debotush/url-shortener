package http

import (
	"github.com/gin-gonic/gin"
	"url-shortener-service/config"
	"url-shortener-service/internal/handlers"
)

// Server represents the HTTP server.
type Server struct {
	cfg    *config.Config
	router *gin.Engine
}

// NewServer creates a new HTTP server instance.
func NewServer(cfg *config.Config) *Server {
	route := gin.Default()
	route.SetTrustedProxies(nil)
	return &Server{
		cfg:    cfg,
		router: route,
	}
}

// Start starts the HTTP server.
func (srv *Server) Start() error {
	return srv.router.Run(srv.cfg.ServerAddress)
}

// RegisterHandlers registers the HTTP request handlers.
func (srv *Server) RegisterHandlers() {
	srv.router.GET("/ping", handlers.PingHandler)
	srv.router.POST("/url-shortener", handlers.UrlShortener)
	srv.router.GET("/:customPath", handlers.RedirectUrl)
}

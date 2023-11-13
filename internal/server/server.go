package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *http.Server
}

func New(addr string, router *gin.Engine) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           addr,
			Handler:        router,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown() error {
	return s.httpServer.Close()
}

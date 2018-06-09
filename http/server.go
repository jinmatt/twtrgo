package http

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/jinmatt/twtrgo/config"

	"github.com/jinmatt/twtrgo/http/handler"
)

// Server type to hold http components
type Server struct {
	handler  *handler.Handler
	server   *http.Server
	listener net.Listener
}

// NewServer inits type Server with http handlers
func NewServer(handler *handler.Handler) *Server {
	return &Server{
		handler: handler,
	}
}

// Start sets up and starts the http server with `Config`
func (s *Server) Start(config *config.Config) error {
	addr := ":" + config.Port
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	s.listener = ln

	handler, err := s.handler.MakeHandler()
	if err != nil {
		return err
	}

	// Start http server
	log.Printf("Listening on '%s'...\n", addr)
	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	s.server = server

	go func() {
		err = server.Serve(ln)
		if err != http.ErrServerClosed {
			log.Fatalln("Error on server.Serve:", err)
			os.Exit(1)
		}
	}()

	return nil
}

// Stop shuts down http server in a given grace time period
func (s *Server) Stop(grace time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), grace)
	err := s.server.Shutdown(ctx)
	if err != nil {
		log.Fatalln("Graceful shutdown expired:", err.Error())
	}
	cancel()
	s.listener.Close()
}

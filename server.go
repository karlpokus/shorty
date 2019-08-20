package shorty

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	*http.Server
	Addr string
}

var Version = "vX.Y.Z" // injected at build time

var Stdout = log.New(os.Stdout, "", 0)
var Stderr = log.New(os.Stderr, "", 0)

func cleanupOnExit(s *Server) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := s.Server.Shutdown(ctx); err != nil {
		Stderr.Printf("Server shutdown err: %s", err)
	} else {
		Stdout.Println("Server shutdown complete")
	}
}

func (s *Server) Start() error {
	go cleanupOnExit(s)
	Stdout.Println(fmt.Sprintf("shorty %s listening on %s", Version, s.Addr))
	return s.ListenAndServe()
}

func New(host, port string) (*Server, error) {
	addr := fmt.Sprintf("%s:%s", host, port)
	return &Server{
		Server: &http.Server{
			Addr:              addr,
			Handler:           newRouter(addr),
			ReadTimeout:       10 * time.Second,
			WriteTimeout:      10 * time.Second,
			ReadHeaderTimeout: 10 * time.Second,
			MaxHeaderBytes:    1 << 20, // 1 MB
		},
		Addr: addr,
	}, nil
}

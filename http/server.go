package http

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

// Server represents an HTTP server.
type Server struct {
	*http.Server
}

// NewServer returns a new instance of Server.
func NewServer() *Server {
	return &Server{
		Server: &http.Server{},
	}
}

// Open opens the server.
func (s *Server) Open() error {
	go s.ListenAndServe()
	return nil
}

// Close shuts down the server.
func (s *Server) Close() error {
	return s.Server.Close()
}

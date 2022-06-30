package server

import (
	"github.com/XM-GO/panda-kit/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(addr string) *http.Server {
	srv := http.NewServer(addr)
	return srv
}

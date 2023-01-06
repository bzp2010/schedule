package server

import (
	"github.com/bzp2010/schedule/internal/config"
	"github.com/gin-gonic/gin"
)

// Server data structure contains references to HTTP and HTTPS services
type Server struct {
	gin     *gin.Engine
	options *Options
}

// Options stores some option values for the server.
type Options struct {
	Config config.Config
}

var (
	server *Server
)

// SetupServer creates server instances
func SetupServer(options *Options) error {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	go r.Run(options.Config.Server.HTTPListen)
	if options.Config.Server.TLS.CertFile != "" && options.Config.Server.TLS.KeyFile != "" {
		go r.RunTLS(
			options.Config.Server.HTTPSListen,
			options.Config.Server.TLS.CertFile,
			options.Config.Server.TLS.KeyFile,
		)
	}

	server = &Server{options: options, gin: r}

	return nil
}

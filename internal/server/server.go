package server

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/bzp2010/schedule/internal/config"
	"github.com/bzp2010/schedule/internal/handler/graphql"
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

	h := graphql.NewGraphQLHandler()
	r.POST("/graphql", func(ctx *gin.Context) {
		h.ServeHTTP(ctx.Writer, ctx.Request)
	})
	r.GET("/playground", func(ctx *gin.Context) {
		playground.
			Handler("GraphQL playground", "/graphql").
			ServeHTTP(ctx.Writer, ctx.Request)
	})

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

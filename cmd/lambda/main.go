package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/handlerfunc"

	"github.com/Je33/packager/internal/config"
	"github.com/Je33/packager/internal/repository/mem"
	"github.com/Je33/packager/internal/service/packer"
	"github.com/Je33/packager/internal/transport/graphql"
	"github.com/Je33/packager/pkg/logger"
)

var (
	// Global variables for execution environment reuse (AWS best practice)
	log     logger.Logger
	handler http.Handler
)

// init runs once during Lambda initialization phase
func init() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	log = logger.New(logger.Config{
		Level: cfg.Log.Level,
	})

	// Initialize repository and service once
	repo := mem.New(log)
	pack := packer.New(repo, log)

	// Create GraphQL handler
	h := graphql.NewHandler(pack)

	// Wrap with logging middleware
	handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("Request", "method", r.Method, "path", r.URL.Path)
		h.ServeHTTP(w, r)
	})
}

func main() {
	lambda.Start(handlerfunc.NewV2(handler.ServeHTTP).ProxyWithContext)
}

package main

import (
	"bytes"
	"io"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"

	"github.com/Je33/packager/internal/config"
	"github.com/Je33/packager/internal/repository/mem"
	"github.com/Je33/packager/internal/service/packer"
	"github.com/Je33/packager/internal/transport/graphql"
	"github.com/Je33/packager/pkg/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	log := logger.New(logger.Config{
		Level: cfg.Log.Level,
	})

	// Initialize repository and service
	repo := mem.New(log)
	pack := packer.New(repo, log)

	// Create GraphQL handler
	h := graphql.NewHandler(pack)
	
	// Wrap with logging middleware to debug request
	loggedHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Read body for logging
		bodyBytes, _ := io.ReadAll(r.Body)
		r.Body.Close()
		
		log.Info("Incoming request", 
			"method", r.Method,
			"path", r.URL.Path,
			"content-type", r.Header.Get("Content-Type"),
			"body", string(bodyBytes))
		
		// Restore body for handler
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		
		h.ServeHTTP(w, r)
	})

	// Create Lambda adapter
	httpLambda := httpadapter.New(loggedHandler)

	log.Info("Lambda handler initialized")

	lambda.Start(httpLambda.ProxyWithContext)
}

package main

import (
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

	// Create Lambda adapter
	httpLambda := httpadapter.New(h)

	log.Info("Lambda handler initialized")

	lambda.Start(httpLambda.ProxyWithContext)
}

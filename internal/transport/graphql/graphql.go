package graphql

import (
	"context"
	"errors"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"
	"github.com/vektah/gqlparser/v2/ast"

	"github.com/Je33/packager/internal/config"
	"github.com/Je33/packager/internal/transport/graphql/generated"
	"github.com/Je33/packager/internal/transport/graphql/resolver"
	"github.com/Je33/packager/pkg/logger"
)

type Transport struct {
	config *config.GraphQLConfig
	server *http.Server
	packer resolver.Packer
	log    logger.Logger
}

func New(config *config.GraphQLConfig, packer resolver.Packer, log logger.Logger) *Transport {

	return &Transport{
		config: config,
		packer: packer,
		server: &http.Server{
			Addr:    ":" + config.Port,
			Handler: NewHandler(packer),
			// TODO: add timeouts
		},
		log: log,
	}
}

func NewHandler(packer resolver.Packer) http.Handler {
	h := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: resolver.New(packer)}))

	h.AddTransport(transport.Options{})
	h.AddTransport(transport.GET{})
	h.AddTransport(transport.POST{})

	h.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	h.Use(extension.Introspection{})
	h.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	mux := http.NewServeMux()

	mux.Handle("/query", h)
	mux.Handle("/", h) // Lambda Function URLs strip the path, so handle root too
	mux.Handle("/playground", playground.Handler("GraphQL", "/query"))

	// TODO: allow only actual domains
	return cors.AllowAll().Handler(mux)
}

func (t *Transport) Start(ctx context.Context) error {
	err := t.server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (t *Transport) Stop(ctx context.Context) error {
	return t.server.Shutdown(ctx)
}

package config

type Config struct {
	Log     *LogConfig     `envconfig:"LOG"`
	GraphQL *GraphQLConfig `envconfig:"GRAPHQL"`
}

type LogConfig struct {
	Level string `envconfig:"LEVEL" default:"info"`
}

type GraphQLConfig struct {
	Port string `envconfig:"PORT" default:"8080"`
}

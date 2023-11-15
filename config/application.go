package config

type Application struct {
	// Application Name
	Name string `env:"APP_NAME"`
	// Application HTTP Port
	HTTPPort int `env:"APP_HTTP_PORT"`
	// Application GRPC Port
	GRPCPort int `env:"APP_GRPC_PORT"`
	// Application Environment
	Environment string `env:"APP_ENV"`
	// Application Debug
	Debug bool `env:"APP_DEBUG"`
}

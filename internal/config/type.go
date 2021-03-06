package config

// Environment holds the values for environment loaded configuration
type Environment struct {
	EnvType                      string `validate:"required,oneof=DEV QA PROD"`
	DBHost, DBUser, DBPass, DBDB string
	DBPort, APIHost, APIPort     string
}

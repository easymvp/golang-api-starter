package swagger

import "os"

const (
	SwaggerEnabled = "SWAGGER_ENABLED"
)

type Config struct {
	Enabled bool
}

func NewConfig() (*Config, error) {
	cfg := Config{
		Enabled: os.Getenv(SwaggerEnabled) == "true",
	}
	return &cfg, nil
}

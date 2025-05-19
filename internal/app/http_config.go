package app

import (
	"os"
	"strconv"
)

const (
	HTTPHost       = "HTTP_HOST"
	HTTPPort       = "HTTP_PORT"
	HTTPJwtSecret  = "HTTP_JWT_SECRET"
	HTTPJwtRealm   = "HTTP_JWT_REALM"
	HTTPJwtExpire  = "HTTP_JWT_EXPIRE"
	HTTPJwtRefresh = "HTTP_JWT_REFRESH"
)

type HTTPConfig struct {
	Host string
	Port string
	Jwt  JwtConfig
}

func NewHttpConfig() (*HTTPConfig, error) {
	httpConfig := HTTPConfig{
		Host: os.Getenv(HTTPHost),
		Port: os.Getenv(HTTPPort),
		Jwt: JwtConfig{
			Secret: os.Getenv(HTTPJwtSecret),
			Realm:  os.Getenv(HTTPJwtRealm),
			ExpirationTime: func() int {
				v, err := strconv.Atoi(os.Getenv(HTTPJwtExpire))
				if err != nil {
					return 3600
				}
				return v
			}(),
			RefreshTime: func() int {
				v, err := strconv.Atoi(os.Getenv(HTTPJwtRefresh))
				if err != nil {
					return 7200
				}
				return v
			}(),
		},
	}
	return &httpConfig, nil
}

func NewJwtConfig(config *HTTPConfig) *JwtConfig {
	return &config.Jwt
}

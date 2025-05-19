package app

import (
	"time"
)

func NewJwtEnvVars(config *JwtConfig) (JwtEnvVars, error) {
	return &jwtEnvVars{
		secret:         config.Secret,
		realm:          config.Realm,
		expirationTime: time.Duration(config.ExpirationTime) * time.Second,
		maxRefreshTime: time.Duration(config.RefreshTime) * time.Second,
	}, nil
}

type JwtEnvVars interface {
	Secret() string
	Realm() string
	Expiration() time.Duration
	RefreshTime() time.Duration
}

type jwtEnvVars struct {
	secret         string
	realm          string
	expirationTime time.Duration
	maxRefreshTime time.Duration
}

func (jwt *jwtEnvVars) Secret() string {
	return jwt.secret
}

func (jwt *jwtEnvVars) Realm() string {
	return jwt.secret
}

func (jwt *jwtEnvVars) Expiration() time.Duration {
	return jwt.expirationTime
}

func (jwt *jwtEnvVars) RefreshTime() time.Duration {
	return jwt.maxRefreshTime
}

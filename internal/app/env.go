package app

import (
	"os"
)

const (
	GoEvnKey string = "GO_ENV"
)

type Env string

const (
	GoEvnTest Env = "TEST"
	GoEvnDev  Env = "DEV"
	GoEvnProd Env = "PROD"
)

func GetGoEnv() Env {
	goEnv := os.Getenv(GoEvnKey)
	if goEnv == "" {
		return GoEvnDev
	}
	return Env(goEnv)
}

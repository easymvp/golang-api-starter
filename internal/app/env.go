package app

import (
	"fmt"
	"github.com/duke-git/lancet/v2/xerror"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
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

func LoadEnv() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	for {
		envPath := filepath.Join(dir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			break
		}
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			// Reached the root directory
			break
		}
		dir = parentDir
	}

	envFile := filepath.Join(dir, ".env")
	err = godotenv.Load(envFile)
	fmt.Println("using env", envFile)
	if err != nil {
		panic(xerror.Wrap(err, fmt.Sprintf("Error loading .env file from %s", dir)))
	}
}

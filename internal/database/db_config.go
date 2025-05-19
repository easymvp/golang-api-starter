package database

import (
	"os"
	"strconv"
)

const (
	DBUrl           = "DB_URL"
	DBDriver        = "DB_DRIVER"
	DBMaxOpenConns  = "DB_MAX_OPEN_CONNS"
	DBMaxIdleConns  = "DB_MAX_IDLE_CONNS"
	DBConnMaxLife   = "DB_CONN_MAX_LIFE"
	DBAutoMigration = "DB_AUTO_MIGRATION"
)

type DBConfig struct {
	Url           string
	Driver        string
	MaxOpenConns  int
	MaxIdleConns  int
	ConnMaxLife   int
	AutoMigration bool
}

func NewDBConfig() (*DBConfig, error) {
	dbConfig := DBConfig{
		Url:    os.Getenv(DBUrl),
		Driver: os.Getenv(DBDriver),
		MaxOpenConns: func() int {
			v, err := strconv.Atoi(os.Getenv(DBMaxOpenConns))
			if err != nil {
				return 10
			}
			return v
		}(),
		MaxIdleConns: func() int {
			v, err := strconv.Atoi(os.Getenv(DBMaxIdleConns))
			if err != nil {
				return 10
			}
			return v
		}(),
		ConnMaxLife: func() int {
			v, err := strconv.Atoi(os.Getenv(DBConnMaxLife))
			if err != nil {
				return 10
			}
			return v
		}(),
		AutoMigration: func() bool {
			v, err := strconv.ParseBool(os.Getenv(DBAutoMigration))
			if err != nil {
				return true
			}
			return v
		}(),
	}
	return &dbConfig, nil
}

package tests

import (
	"context"
	"easymvp_api/internal/app"
	"easymvp_api/internal/database"
	"fmt"
	"go.uber.org/fx"
	"gorm.io/driver/sqlite"
	"os"
	"strconv"
)

func NewDBConfig() (*database.DBConfig, error) {
	fmt.Println(sqlite.Config{})
	dbConfig := database.DBConfig{
		Url:    "file::memory:?cache=shared",
		Driver: "sqlite3",
		MaxOpenConns: func() int {
			v, err := strconv.Atoi(os.Getenv(database.DBMaxOpenConns))
			if err != nil {
				return 10
			}
			return v
		}(),
		MaxIdleConns: func() int {
			v, err := strconv.Atoi(os.Getenv(database.DBMaxIdleConns))
			if err != nil {
				return 10
			}
			return v
		}(),
		ConnMaxLife: func() int {
			v, err := strconv.Atoi(os.Getenv(database.DBConnMaxLife))
			if err != nil {
				return 10
			}
			return v
		}(),
		AutoMigration: func() bool {
			v, err := strconv.ParseBool(os.Getenv(database.DBAutoMigration))
			if err != nil {
				return true
			}
			return v
		}(),
	}
	return &dbConfig, nil
}

var Module = fx.Options(
	fx.Provide(
		database.NewDB, database.NewPGX, database.NewEntityRegistry, NewDBConfig, database.NewGormLogger,
	),
	fx.Invoke(func(lifecycle fx.Lifecycle, registry *database.ModelRegistry) {
		lifecycle.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				if app.GetGoEnv() == app.GoEvnTest {
					err := registry.Migrate()
					if err != nil {
						return err
					}
				}
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return nil
			},
		})
	}),
)

package database

import (
	"context"
	"database/sql"
	"easymvp_api/internal/app"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func NewDB(
	cfg *DBConfig,
	GormLogger GormLogger,
) *gorm.DB {
	sqlDB, err := sql.Open(cfg.Driver, cfg.Url)
	if err != nil {
		panic(err.Error())
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: GormLogger,
	})
	if err != nil {
		panic(err.Error())
	}

	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLife) * time.Second)

	return gormDB
}

func NewPGX(
	cfg *DBConfig,
) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(cfg.Url)
	if err != nil {
		return nil, err
	}

	config.MaxConns = int32(cfg.MaxOpenConns)
	pgxPool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	err = pgxPool.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return pgxPool, nil
}

var Module = fx.Options(
	fx.Provide(
		NewDB, NewPGX, NewEntityRegistry, NewDBConfig, NewGormLogger,
	),
	fx.Invoke(func(lifecycle fx.Lifecycle, registry *ModelRegistry) {
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

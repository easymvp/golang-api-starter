package database

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ModelRegistry struct {
	entities []interface{}
	DB       *gorm.DB
	logger   *zap.Logger
}

func NewEntityRegistry(DB *gorm.DB, logger *zap.Logger) *ModelRegistry {
	return &ModelRegistry{
		DB:     DB,
		logger: logger,
	}
}

func (e *ModelRegistry) Register(entity interface{}) {
	e.entities = append(e.entities, entity)
}

func (e *ModelRegistry) Migrate() error {
	e.logger.Info("migrating entity", zap.Any("entity", e.entities))
	return e.DB.AutoMigrate(e.entities...)
}

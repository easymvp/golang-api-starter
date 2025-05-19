package app

import (
	"easymvp_api/internal/app/handlers"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type userProviderImpl struct {
	db *gorm.DB
}

func NewUserProvider(db *gorm.DB) UserProvider {
	return &userProviderImpl{db: db}
}

func (u *userProviderImpl) Get(id string) (*UserInfo, error) {
	var user UserInfo
	err := u.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

type workspaceProviderImpl struct {
	db *gorm.DB
}

func NewWorkspaceProvider(db *gorm.DB) WorkspaceProvider {
	return &workspaceProviderImpl{db: db}
}

func (w *workspaceProviderImpl) Get(id string) (*WorkspaceInfo, error) {
	var workspace WorkspaceInfo
	err := w.db.Where("id = ?", id).First(&workspace).Error
	if err != nil {
		return nil, err
	}
	return &workspace, nil
}

var Module = fx.Module("web",
	fx.Provide(NewWebApp, NewHttpConfig, NewJwtConfig),
	fx.Provide(handlers.NewHealthCheckHandler),
	fx.Invoke(func(s *App, handler *handlers.HealthCheckHandler) {
		s.NoAuth.GET("/api/health-check", handler.HealthCheck())
	}),
	fx.Provide(NewUserProvider),
	fx.Provide(NewWorkspaceProvider),
)

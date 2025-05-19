package users

import (
	"easymvp_api/internal/database"
	"easymvp_api/internal/users/entities"
	"go.uber.org/fx"
)

var Module = fx.Module("users",
	fx.Provide(NewUserService),
	fx.Invoke(func(registry *database.ModelRegistry) {
		registry.Register(entities.User{})
	}),
)

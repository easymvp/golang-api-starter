package users

import (
	"context"
	"easymvp_api/internal/app"
	"easymvp_api/internal/log"
	"easymvp_api/internal/tests"
	"easymvp_api/internal/users/entities"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"testing"
)

var userService *UserService

var (
	mockApp = tests.NewTestApp(Module, app.Module, log.Module, tests.Module, fx.Populate(&userService))
)

func TestMain(m *testing.M) {
	mockApp.Run(m)
}

func NewTestContext(t *testing.T) context.Context {
	ctx := context.Background()
	user := &entities.User{
		Username: "testuser",
		Password: "password",
	}
	err := userService.Save(ctx, user)
	require.NoError(t, err)
	return ctx
}

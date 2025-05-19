package users

import (
	"context"
	"easymvp_api/internal/users/entities"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserService_Get(t *testing.T) {
	mockApp.Reset()

	// create user
	user := &entities.User{
		Username: "testuser",
		Password: "password",
	}
	err := userService.Save(context.Background(), user)
	require.NoError(t, err)

	// get user
	user2, err := userService.Get(context.Background(), user.ID)
	require.NoError(t, err)
	require.Equal(t, user.ID, user2.ID)
	require.Equal(t, user.Username, user2.Username)
	require.Equal(t, user.Password, user2.Password)
}

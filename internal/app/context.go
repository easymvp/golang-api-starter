package app

import (
	"context"
	"gorm.io/gorm"
)

const (
	DBKey        string = "db"
	UserKey      string = "user_info"
	WorkspaceKey string = "workspace_info"
)

func DatabaseOf(ctx context.Context) *gorm.DB {
	value := ctx.Value(DBKey)
	return value.(*gorm.DB)
}

func UserOf(ctx context.Context) *UserInfo {
	value := ctx.Value(UserKey)
	return value.(*UserInfo)
}

func WorkspaceOf(ctx context.Context) *WorkspaceInfo {
	value := ctx.Value(WorkspaceKey)
	return value.(*WorkspaceInfo)
}

func WithDatabase(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, DBKey, db.WithContext(ctx))
}

func WithUser(ctx context.Context, user *UserInfo) context.Context {
	return context.WithValue(ctx, UserKey, user)
}

func WithWorkspace(ctx context.Context, workspace *WorkspaceInfo) context.Context {
	return context.WithValue(ctx, WorkspaceKey, workspace)
}

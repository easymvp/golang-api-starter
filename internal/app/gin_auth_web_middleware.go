package app

import (
	"context"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GinAuthWebMiddleware(userInfoProvider UserProvider, workspaceInfoProvider WorkspaceProvider, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		userId := claims[IdentityKey].(string)
		if userId == "" {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		user, err := userInfoProvider.Get(userId)
		if err != nil {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		workspaceInfo, err := workspaceInfoProvider.Get(user.WorkspaceID)
		if err != nil {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, UserKey, user)
		ctx = context.WithValue(ctx, DBKey, db.WithContext(ctx))
		ctx = context.WithValue(ctx, WorkspaceKey, workspaceInfo)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

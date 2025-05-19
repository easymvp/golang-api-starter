package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GinNoAuthWebMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, DBKey, db.WithContext(ctx))
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

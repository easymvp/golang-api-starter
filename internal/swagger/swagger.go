package swagger

import (
	"easymvp_api/docs"
	"easymvp_api/internal/app"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
)

var Module = fx.Module("swagger",
	fx.Provide(NewConfig),
	fx.Invoke(func(server *app.App, config *Config) {
		if config.Enabled {
			server.Gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
		print(docs.SwaggerInfo.BasePath)
	}),
)

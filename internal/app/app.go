package app

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

type App struct {
	HttpConfig            *HTTPConfig
	Logger                *zap.Logger
	Gin                   *gin.Engine
	DB                    *gorm.DB
	UserInfoProvider      UserProvider
	WorkspaceInfoProvider WorkspaceProvider
	NoAuth                *gin.RouterGroup
	NeedsAuth             *gin.RouterGroup
}

func NewWebApp(
	Logger *zap.Logger,
	userInfoProvider UserProvider,
	workspaceInfoProvider WorkspaceProvider,
	DB *gorm.DB,
	httpConfig *HTTPConfig,
) *App {
	g := gin.Default()
	_ = g.SetTrustedProxies(nil)

	l := Logger
	if l == nil {
		l = zap.NewNop()
	}

	s := &App{
		HttpConfig:            httpConfig,
		Logger:                l,
		Gin:                   g,
		DB:                    DB,
		UserInfoProvider:      userInfoProvider,
		WorkspaceInfoProvider: workspaceInfoProvider,
	}
	s.Init()
	return s
}

func (s *App) Run(addr string) error {
	return s.Gin.Run(addr)
}

func (s *App) Init() {
	s.Gin.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	})).Use(func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.Status(http.StatusNoContent)
			return
		}
		c.Next()
	})
	s.NoAuth = s.Gin.Group("/")
	s.NoAuth.Use(GinNoAuthWebMiddleware(s.DB))

	jwtMiddleware := InitParams(&s.HttpConfig.Jwt)
	s.NeedsAuth = s.Gin.Group("/", jwtMiddleware.MiddlewareFunc())
	authMiddleware, err := jwt.New(jwtMiddleware)
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	s.NeedsAuth.Use(HandlerMiddleware(authMiddleware))
	s.NeedsAuth.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Next()
	})
	s.NeedsAuth.Use(GinAuthWebMiddleware(s.UserInfoProvider, s.WorkspaceInfoProvider, s.DB))
}

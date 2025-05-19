package tests

import (
	"context"
	"database/sql"
	"easymvp_api/internal/app"
	"fmt"
	"os"
	"testing"

	"easymvp_api/internal/database"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type TestApp struct {
	app     *fx.App
	options []fx.Option
	server  *app.App
}

func NewTestApp(options ...fx.Option) *TestApp {
	os.Setenv("GO_ENV", "TEST")
	return &TestApp{
		options: options,
	}
}

func (s *TestApp) Run(m *testing.M) {
	err := s.Start()
	if err != nil {
		println("failed to start app", err.Error())
	}
	m.Run()
	s.Stop()
}

func (s *TestApp) Router() *gin.Engine {
	return s.server.Gin
}

func (s *TestApp) Reset() error {
	db, err := sql.Open("sqlite", os.Getenv(database.DBUrl))
	if err != nil {
		return err
	}
	query := `TRUNCATE SCHEMA public AND COMMIT;`

	_, err = db.Exec(query)
	return err
}

func (s *TestApp) Start() error {
	fmt.Println("setup")
	var server *app.App
	options := append(s.options,
		fx.Populate(&server),
	)
	app := fx.New(
		options...,
	)
	s.app = app
	s.server = server

	err := s.app.Start(context.Background())
	return err
}

func (s *TestApp) Stop() error {
	fmt.Println("teardown")
	return nil
}

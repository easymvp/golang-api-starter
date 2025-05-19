package log

import (
	"easymvp_api/internal/app"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"strings"
)

// Custom caller encoder to log relative file paths.
func relativeCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	basePath, _ := filepath.Abs(".")
	absPath := caller.File
	gopath := os.Getenv("GOPATH")

	relPath, err := filepath.Rel(basePath, absPath)
	if err != nil || strings.HasPrefix(relPath, "../") || filepath.IsAbs(relPath) {
		if gopath != "" && strings.HasPrefix(absPath, gopath) {
			absPath = strings.Replace(absPath, gopath, "$GOPATH", 1)
		}
		enc.AppendString(fmt.Sprintf("%s:%d", absPath, caller.Line))
	} else {
		enc.AppendString(fmt.Sprintf("%s:%d", relPath, caller.Line))
	}
}

// Provide a custom zap logger with the relative caller encoder.
func NewCustomZapLogger() (*zap.Logger, error) {
	env := app.GetGoEnv()
	if env == app.GoEvnProd {
		config := zap.Config{
			Encoding:    "json",
			Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
			OutputPaths: []string{"stdout"},
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey:   "message",
				LevelKey:     "level",
				TimeKey:      "time",
				CallerKey:    "caller",
				EncodeLevel:  zapcore.LowercaseLevelEncoder,
				EncodeTime:   zapcore.ISO8601TimeEncoder,
				EncodeCaller: relativeCallerEncoder,
			},
		}
		return config.Build()
	} else {
		logger := zap.Must(zap.NewDevelopment())
		return logger, nil
	}
}

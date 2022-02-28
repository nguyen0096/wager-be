package log

import (
	"encoding/json"
	"fmt"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
)

const DefaultZapConfig = `
{
	"level": "debug",
	"encoding": "json",
	"outputPaths": ["stdout", "/tmp/logs"],
	"errorOutputPaths": ["stderr"],
	"encoderConfig": {
		"messageKey": "message",
		"levelKey": "level",
		"levelEncoder": "lowercase"
	}
}
`

func NewLogger() logr.Logger {
	cfg := &zap.Config{}

	if err := json.Unmarshal([]byte(DefaultZapConfig), &cfg); err != nil {
		panic(fmt.Errorf("failed to unmarshal zap log config. err: %w", err))
	}

	zapLog, err := cfg.Build()
	if err != nil {
		panic(fmt.Errorf("failed to build zap log. err: %w", err))
	}

	return zapr.NewLogger(zapLog)
}

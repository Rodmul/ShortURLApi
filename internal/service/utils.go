package service

import (
	"encoding/json"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
)

func NewLogger() *zap.SugaredLogger {
	loggerConfig := zap.NewProductionEncoderConfig()
	loggerConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	defaultLogLevel := zapcore.DebugLevel
	consoleEncoder := zapcore.NewConsoleEncoder(loggerConfig)
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	sugar := logger.Sugar()

	return sugar
}

func (srv *Service) error(w http.ResponseWriter, code int, err error) {
	srv.respond(w, code, map[string]string{"message": err.Error()})
	u, ok := err.(interface {
		ErrorEx() string
	})
	if !ok {
		srv.Logger.Error(err.Error())
	} else {
		srv.Logger.Error(u.ErrorEx())
	}
}

func (srv *Service) warning(w http.ResponseWriter, code int, err error) {
	srv.respond(w, code, map[string]string{"message": err.Error()})
	u, ok := err.(interface {
		ErrorEx() string
	})
	if !ok {
		srv.Logger.Warn(err.Error())
	} else {
		srv.Logger.Warn(u.ErrorEx())
	}
}

func (srv *Service) respond(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			srv.Logger.Errorf("failed to encode json %v", err)
		}
	}
}

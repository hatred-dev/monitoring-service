package logger

import "go.uber.org/zap"

var Log *zap.SugaredLogger

func InitLogger() {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()
	Log = sugar
}

package logger

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func InitLogger(appName string, env string) {
	config := zap.NewDevelopmentConfig()
	if env == "production" {
		config = zap.NewProductionConfig()
	}

	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	config.DisableStacktrace = true

	log, err := config.Build()
	if err != nil {
		fmt.Errorf("error init logger: %s", err.Error())
	}
	log = log.With(zap.String("app_name", appName))

	logger = log
	defer logger.Sync()
}

func initBaseLoggerFields(ctx context.Context) []zapcore.Field {
	traceID, ok := ctx.Value(TraceID).(string)
	if !ok {
		traceID = uuid.NewString()
	}
	fileName, methodName := traceFileNameAndMethodName()

	return []zapcore.Field{
		zap.String("trace_id", traceID),
		zap.String("method_name", methodName),
		zap.String("file_name", fileName),
	}
}

func traceFileNameAndMethodName() (fileName, methodName string) {
	pc, file, line, _ := runtime.Caller(3)
	details := runtime.FuncForPC(pc)

	return fmt.Sprintf("%s:%d", file, line), details.Name()
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, initBaseLoggerFields(ctx)...)
	logger.Info(msg, fields...)
}

func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, initBaseLoggerFields(ctx)...)
	logger.Debug(msg, fields...)
}

func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, initBaseLoggerFields(ctx)...)
	logger.Warn(msg, fields...)
}

func Error(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, initBaseLoggerFields(ctx)...)
	logger.Error(msg, fields...)
}

func Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, initBaseLoggerFields(ctx)...)
	logger.Fatal(msg, fields...)
}

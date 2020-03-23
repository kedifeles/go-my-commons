package logger

import (
	"context"

	cmnCtx "github.com/kedifeles/go-my-commons/context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

// DefaultEncoderConfig is default encoder config of zapcore logging
var DefaultEncoderConfig zapcore.EncoderConfig = zapcore.EncoderConfig{
	TimeKey:        "ts",
	LevelKey:       "level",
	NameKey:        "logger",
	CallerKey:      "caller",
	MessageKey:     "msg",
	StacktraceKey:  "trace",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    zapcore.LowercaseLevelEncoder,
	EncodeTime:     zapcore.RFC3339NanoTimeEncoder,
	EncodeDuration: zapcore.SecondsDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}

// DefaultConfig is default config of zapcore logging
var DefaultConfig zap.Config = zap.Config{
	Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
	Development:      false,
	Sampling:         nil,
	Encoding:         "json",
	EncoderConfig:    DefaultEncoderConfig,
	OutputPaths:      []string{"stderr"},
	ErrorOutputPaths: []string{"stderr"},
}

func init() {
	// config := defaultConfig
	// config.OutputPaths = []string{"stdout"}
	newLogger, err := DefaultConfig.Build()

	//   logger := zap.New(core)
	// a fallback/root logger for events without context

	// logger = zap.New(core)
	// logger.WithOptions()

	// newLogger, err := config.Build(zap.WrapCore(func(core zapcore.Core) zapcore.Core {
	// w := zapcore.AddSync(&lumberjack.Logger{
	// 	Filename: "foo.log",
	// 	MaxSize:  500, // megabytes
	// 	// MaxBackups: 3,
	// 	// MaxAge:     28, // days
	// })
	// return zapcore.NewCore(
	// 	zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
	// 	w,
	// 	zap.InfoLevel,
	// )

	// 	return core
	// }))

	if err != nil {
		logger.Fatalf("Failed to init logger")
	}

	logger = newLogger.Sugar()
}

// Logger return zap Sugared Logger with fields from context
func Logger(ctx context.Context) *zap.SugaredLogger {
	newLogger := logger
	if ctx != nil {
		if ctxHTTPReqID, ok := ctx.Value(cmnCtx.HTTPReqIDKey).(string); ok {
			newLogger = newLogger.With(zap.String(cmnCtx.HTTPReqIDKey.String(), ctxHTTPReqID))
		}
		if ctxHTTPSessID, ok := ctx.Value(cmnCtx.HTTPSessIDKey).(string); ok {
			newLogger = newLogger.With(zap.String(cmnCtx.HTTPSessIDKey.String(), ctxHTTPSessID))
		}
		if ctxBNISessID, ok := ctx.Value(cmnCtx.SessIDKey).(string); ok {
			newLogger = newLogger.With(zap.String(cmnCtx.SessIDKey.String(), ctxBNISessID))
		}
	}
	return newLogger
}

// SetOptions set zap options
func SetOptions(opts ...zap.Option) {
	logger = logger.Desugar().WithOptions(opts...).Sugar()
}

// With compose new zap Sugared Logger with provided fields
func With(args ...interface{}) *zap.SugaredLogger {
	return logger.With(args)
}

// Debug uses fmt.Sprint to construct and log a message.
func Debug(args ...interface{}) {
	logger.Debug(args)
}

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
	logger.Info(args)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	logger.Warn(args)
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	logger.Error(args)
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanic(args ...interface{}) {
	logger.DPanic(args)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func Panic(args ...interface{}) {
	logger.Panic(args)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatal(args ...interface{}) {
	logger.Fatal(args)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	logger.Infof(template, args)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanicf(template string, args ...interface{}) {
	logger.DPanicf(template, args)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func Panicf(template string, args ...interface{}) {
	logger.Panicf(template, args)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//  s.With(keysAndValues).Debug(msg)
func Debugw(msg string, keysAndValues ...interface{}) {
	logger.Debugw(msg, keysAndValues)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Infow(msg string, keysAndValues ...interface{}) {
	logger.Infow(msg, keysAndValues)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Warnw(msg string, keysAndValues ...interface{}) {
	logger.Warnw(msg, keysAndValues)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Errorw(msg string, keysAndValues ...interface{}) {
	logger.Errorw(msg, keysAndValues)
}

// DPanicw logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func DPanicw(msg string, keysAndValues ...interface{}) {
	logger.DPanicw(msg, keysAndValues)
}

// Panicw logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With.
func Panicw(msg string, keysAndValues ...interface{}) {
	logger.Panicw(msg, keysAndValues)
}

// Fatalw logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func Fatalw(msg string, keysAndValues ...interface{}) {
	logger.Fatalw(msg, keysAndValues)
}

// Sync flushes any buffered log entries.
func Sync() error {
	return logger.Sync()
}

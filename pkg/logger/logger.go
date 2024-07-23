package logger

import (
	"my_shop/pkg/setting"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerConfig) *LoggerZap {
	logLever := config.LogLevel
	//debug --> info --> warn --> error --> fatal --> panic
	var lever zapcore.Level
	switch logLever {
		case "debug":
            lever = zap.DebugLevel
        case "info":
            lever = zap.InfoLevel
        case "warn":
            lever = zap.WarnLevel
        case "error":
            lever = zap.ErrorLevel
        case "fatal":
            lever = zap.FatalLevel
        default:
            lever = zap.InfoLevel
	}

	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename: config.Filename,
		MaxSize: config.MaxSize, // megabytes
		MaxBackups: config.MaxBackups,
		MaxAge: config.MaxAge, // dealine days to delete file
		Compress: config.Compress, // disable compression
	}
	core := zapcore.NewCore(
		encoder, 
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), 
		lever,
	)

	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}

//format logs of a message
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	// 176842932.877321 --> 2024-07-23T9:42:02.343+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// ts --> time
	encodeConfig.TimeKey = "time"

	// from info INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// caller initalize/run.go
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}
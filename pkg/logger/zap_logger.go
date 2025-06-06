package logger

import (
	"os"

	"github.com/lyonnee/go-template/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZapLogger(
	logConfig config.LogConfig,
) (*zap.Logger, error) {
	encoder := getEncoder(logConfig.Format, logConfig.Caller)

	var cores = make([]zapcore.Core, 0)

	consoleCore, err := getConsoleWriterCore(encoder, logConfig.Level)
	if err != nil {
		return nil, err
	}
	cores = append(cores, consoleCore)

	if logConfig.ToFile {
		fileCore, err := getFileWriterCore(
			encoder,
			logConfig.LogFileConfig.Filename,
			logConfig.LogFileConfig.LogLevel,
			logConfig.LogFileConfig.MaxSize,
			logConfig.LogFileConfig.MaxBackups,
			logConfig.LogFileConfig.MaxAge,
			logConfig.LogFileConfig.IsCompression,
		)
		if err != nil {
			return nil, err
		}
		cores = append(cores, fileCore)
	}

	core := zapcore.NewTee(cores...)

	return zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)), nil
}

func getEncoder(format, encodeCaller string) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder

	if encodeCaller == "full" {
		encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	} else {
		encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	}

	var encoder func(cfg zapcore.EncoderConfig) zapcore.Encoder
	if format == "console" {
		encoder = zapcore.NewConsoleEncoder
	} else {
		encoder = zapcore.NewJSONEncoder
	}

	return encoder(encoderConfig)
}

func getConsoleWriterCore(encoder zapcore.Encoder, levelStr string) (zapcore.Core, error) {
	level, err := zapcore.ParseLevel(levelStr)
	if err != nil {
		return nil, err
	}

	return zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level), nil
}

func getFileWriterCore(encoder zapcore.Encoder, filename, levelStr string, maxSize, maxBackups, maxAge int, isCompression bool) (zapcore.Core, error) {
	level, err := zapcore.ParseLevel(levelStr)
	if err != nil {
		return nil, err
	}

	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   isCompression,
	}

	syncer := zapcore.AddSync(lumberJackLogger)
	return zapcore.NewCore(encoder, syncer, level), nil
}

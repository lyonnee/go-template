package log

import (
	"os"

	"github.com/lyonnee/go-template/internal/infrastructure/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newZapLogger(
	logConfig config.LogConfig,
) (*zap.Logger, error) {
	var cores = make([]zapcore.Core, 0)

	consoleCore, err := getConsoleWriterCore(logConfig.ConsoleWriterConfig)
	if err != nil {
		return nil, err
	}
	if consoleCore != nil {
		cores = append(cores, consoleCore)
	}

	fileCore, err := getFileWriterCore(logConfig.FileWriterConfig)
	if err != nil {
		return nil, err
	}
	if fileCore != nil {
		cores = append(cores, fileCore)
	}

	core := zapcore.NewTee(cores...)

	return zap.New(core, zap.AddCaller()), nil
}

func getConsoleWriterCore(conf config.LogConsoleWriterConfig) (zapcore.Core, error) {
	level, err := zapcore.ParseLevel(conf.Level)
	if err != nil {
		return nil, err
	}

	encode := getEncoder(conf.Format, conf.Caller)

	return zapcore.NewCore(encode, zapcore.AddSync(os.Stdout), level), nil
}

func getFileWriterCore(conf config.LogFileWriterConfig) (zapcore.Core, error) {
	level, err := zapcore.ParseLevel(conf.Level)
	if err != nil {
		return nil, err
	}

	encoder := getEncoder(conf.Format, conf.Caller)

	lumberJackLogger := &lumberjack.Logger{
		Filename:   conf.Filename,
		MaxSize:    conf.MaxSize,
		MaxBackups: conf.MaxBackups,
		MaxAge:     conf.MaxAge,
		Compress:   conf.IsCompression,
	}
	syncer := zapcore.AddSync(lumberJackLogger)

	return zapcore.NewCore(encoder, syncer, level), nil
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

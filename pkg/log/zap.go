package log

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newZap(
	enableToConsole bool, toConsoleLevel string,
	filename, toFileLevel string, maxSize, maxBackups, maxAge int,
) (*zap.Logger, error) {
	encoder := getEncoder()

	var cores = make([]zapcore.Core, 0)

	fileCore, err := getFileWriterCore(encoder, filename, toFileLevel, maxSize, maxBackups, maxAge)
	if err != nil {
		return nil, err
	}
	cores = append(cores, fileCore)

	if enableToConsole {
		consoleCore, err := getConsoleWriterCore(encoder, toConsoleLevel)
		if err != nil {
			return nil, err
		}
		cores = append(cores, consoleCore)
	}

	core := zapcore.NewTee(cores...)

	return zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)), nil
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}

func getConsoleWriterCore(encoder zapcore.Encoder, levelStr string) (zapcore.Core, error) {
	level, err := zapcore.ParseLevel(levelStr)
	if err != nil {
		return nil, err
	}

	return zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level), nil
}

func getFileWriterCore(encoder zapcore.Encoder, filename, levelStr string, maxSize, maxBackups, maxAge int) (zapcore.Core, error) {
	level, err := zapcore.ParseLevel(levelStr)
	if err != nil {
		return nil, err
	}

	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
	}

	syncer := zapcore.AddSync(lumberJackLogger)
	return zapcore.NewCore(encoder, syncer, level), nil
}

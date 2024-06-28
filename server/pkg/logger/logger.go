package logger

import (
	"fmt"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	B *zap.Logger
)

func init() {
	B = getLogger("base")
}

func getLogger(name string) *zap.Logger {
	var (
		debugLogFile    = os.Stdout
		infoLogFile     *lumberjack.Logger
		infoLogFileName string

		fileEncoder zapcore.Encoder
	)

	infoLogFileName = fmt.Sprintf("/tmp/log/%s_info.log", name)
	infoLogFile = &lumberjack.Logger{
		Filename:   infoLogFileName,
		MaxSize:    500, // megabytes
		MaxBackups: 5,
		MaxAge:     28,    //days
		Compress:   false, // disabled by default
	}

	fileEncoder = zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		CallerKey:      "caller",
		TimeKey:        "t",
		LevelKey:       "level",
		NameKey:        "logger",
		MessageKey:     "msg",
		StacktraceKey:  "trace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})

	teecore := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(debugLogFile), zap.DebugLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(infoLogFile), zap.InfoLevel),
	)

	return zap.New(teecore, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
}

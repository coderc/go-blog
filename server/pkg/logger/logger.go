package logger

import (
	"fmt"
	"os"
	"time"

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
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        name,
		CallerKey:      "caller",
		FunctionKey:    "func",
		StacktraceKey:  "stacktrace",
		SkipLineEnding: true,
		EncodeLevel: func(zapcore.Level, zapcore.PrimitiveArrayEncoder) {
		},
		EncodeTime: func(time.Time, zapcore.PrimitiveArrayEncoder) {
		},
		EncodeDuration: func(time.Duration, zapcore.PrimitiveArrayEncoder) {
		},
		EncodeCaller: func(zapcore.EntryCaller, zapcore.PrimitiveArrayEncoder) {
		},
		EncodeName: func(string, zapcore.PrimitiveArrayEncoder) {
		},
		// NewReflectedEncoder: func( io.Writer) zapcore.ReflectedEncoder {
		// },
		ConsoleSeparator: "",
	})

	teecore := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(debugLogFile), zap.DebugLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(infoLogFile), zap.InfoLevel),
	)

	return zap.New(teecore, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
}

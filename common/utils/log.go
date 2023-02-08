package utils

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/kitex-contrib/obs-opentelemetry/logging/zap"
	"go.uber.org/zap/zapcore"
)

// Initializes logger for Kitex with Zap.
func InitKlog() {
	klog.SetLogger(zap.NewLogger(zap.WithCoreEnc(GetZapEncoder())))
	klog.SetLevel(klog.LevelDebug)
}

func GetZapEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		LevelKey:      "level",
		TimeKey:       "ts",
		MessageKey:    "msg",
		CallerKey:     "caller",
		NameKey:       "logger",
		StacktraceKey: "stacktrace",
		EncodeLevel:   zapcore.CapitalColorLevelEncoder,
		EncodeTime:    zapcore.TimeEncoderOfLayout("15:04:05 Mon"),
	})
}

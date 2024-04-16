package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func init() {
	logFile := "df_zap.log"

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	file.WriteString("\n")
	
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		MessageKey:     "msg",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.AddSync(file),
		zap.DebugLevel,
	)

	logger := zap.New(core)

	zap.ReplaceGlobals(logger)
	zap.S().Info("Logger initialized")
}


func Info(s string){
	zap.S().Info(s)
}
func Error(s string, err error){
	zap.S().Error(s, err)
}
func Debug(s string){
	zap.S().Debug(s)
}
func Sync(){
	zap.L().Sync()
}

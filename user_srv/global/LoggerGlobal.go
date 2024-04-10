package global

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var (
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
)

func InitLogger() {
	//writeSyncer := getLogWriter()
	//encoder := getEncoder()
	//core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	//Logger = zap.New(core)
	logger, _ := zap.NewDevelopment()
	Logger = logger
	zap.ReplaceGlobals(Logger)
	defer Logger.Sync()
	Sugar = Logger.Sugar()
}

func getLogWriter() zapcore.WriteSyncer {
	//输出到文件
	file, _ := os.OpenFile("./zap2023.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666) // 打开文件（不存在创建）
	return zapcore.AddSync(file)
}

func getEncoder() zapcore.Encoder {
	//encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = customTimeEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

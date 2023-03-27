package main

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
)

var logger *zap.Logger

func main() {
	InitLogger()
	defer logger.Sync()

	//simpleHttpGet("www.sogo.com")
	//simpleHttpGet("http://www.sogo.com")
	for i := 0; i < 10000; i++ {
		logger.Info("testing...")
	}
}
func InitLogger() {
	writerSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)

	logger = zap.New(core, zap.AddCaller())
	//AddCaller :记录调用信息
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

//原版
//func getEncoder() zapcore.Encoder {
//
//	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
//}

//func getLogWriter() zapcore.WriteSyncer {
//	file, _ := os.Create("./test2.log")
//	return zapcore.AddSync(file)
//}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url...",
			zap.String("url", url),
			zap.Error(err))

	} else {
		logger.Info("Success...",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}

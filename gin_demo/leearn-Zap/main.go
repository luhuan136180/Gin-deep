package main

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"

	"net/http"
)

var (
	SugarLogger *zap.SugaredLogger
	logger      *zap.Logger
)

func main() {
	InitLogger()

	defer logger.Sync()
	for i := 0; i < 10000; i++ {
		logger.Info("testing info ......")
	}
	simpleHttpGet2("www.baidu.com")
	fmt.Println("////")
	simpleHttpGet2("https://www.baidu.com")

}

//给定的创造loger
//func InitLogger() {
//	logger, _ = zap.NewProduction()
//	SugarLogger = logger.Sugar()
//}
//以下，自定义创建
func InitLogger() {
	//指定日志将写到哪里去
	writeSyncer := GetLogWriter2()
	//规定写入日志的格式（编码器）
	Encoder := GetEncoder()
	//func NewCore(enc Encoder, ws WriteSyncer, enab LevelEnabler) Core
	//打造核心
	core := zapcore.NewCore(Encoder, writeSyncer, zapcore.DebugLevel)

	//func New(core zapcore.Core, options ...Option) *Logger

	//zap.AddCaller():添加将调用函数信息记录到日志中的功能
	logger = zap.New(core, zap.AddCaller())
	SugarLogger = logger.Sugar()

}

//Encoder:编码器(如何写入日志)。
func GetEncoder() zapcore.Encoder {
	//用json格式包装
	//日志内容：使用预先设置的ProductionEncoderConfig()格式
	//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	//将编码器从JSON Encoder更改为普通Encoder。为此，
	//我们需要将NewJSONEncoder()更改为NewConsoleEncoder()
	//return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())

	//需求：用正常时间输出
	//对日志内容格式进行修改
	//encoderConfig:=zap.NewProductionEncoderConfig(),同下
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder, //人类可读时间
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

//WriterSyncer ：指定日志将写到哪里去
//zap库中的原生模式
func GetLogWriter1() zapcore.WriteSyncer {
	//os包，创建文件test.log
	//file, _ := os.Create("leearn-Zap/test.log")

	//1可添加日志
	file, _ := os.OpenFile("leearn-Zap/test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	//指定日志传入的文件
	return zapcore.AddSync(file)
}

//引入github.com/natefinch/lumberjack 库，对日志进行切割存储
func GetLogWriter2() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    1,     // .M 文件大小
		MaxBackups: 5,     //  备份数量
		MaxAge:     30,    //最大备份天数
		Compress:   false, //是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}

func simpleHttpGet1(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		logger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}

func simpleHttpGet2(url string) {
	SugarLogger.Debug("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		SugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		SugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}

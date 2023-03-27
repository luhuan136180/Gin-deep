package main

import (
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

func main() {
	InitLogger()
	//r := gin.Default()  ==	engine := New()
	//							engine.Use(Logger(), Recovery())

	//手动自建引擎
	r := gin.New()
	r.Use(GinLogger(logger), GinRecovery(logger, true))

	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello liwenzhou.com!")
	})
	r.Run()
}

//主动实现Engine里的logger和Recovery中间件
// GinLogger 接收gin框架默认的日志
func GinLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		logger.Info(path, //强制要求日志内容里的内容
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic
func GinRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

var (
	logger      *zap.Logger
	SugarLogger *zap.SugaredLogger
)

//创建logger
func InitLogger() {
	writeSyncer := GetLogWriter2()
	Encoder := GetEncoder()
	//func NewCore(enc Encoder, ws WriteSyncer, enab LevelEnabler) Core
	core := zapcore.NewCore(Encoder, writeSyncer, zapcore.DebugLevel)

	//func New(core zapcore.Core, options ...Option) *Logger
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
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func GetLogWriter2() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    1,     // .M
		MaxBackups: 5,     //  备份数量
		MaxAge:     30,    //最大备份天数
		Compress:   false, //是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}

package logging

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"onlinemall/config"
	"os"
	"path"
	"time"
)

func LoggerHandler() gin.HandlerFunc {
	// get log file
	logFilePath := config.ServerConfig.LogDir
	logFileName := config.ServerConfig.LogFile

	//init log
	initLog(logFilePath, logFileName, localLogger)

	//format logger
	return formatLog
}

func initLog(logFilePath string, logFileName string, logger *logrus.Logger) {
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)

	// 写入文件
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("os open file error ", err)
	}

	// 设置输出
	logger.Out = src

	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+"-%Y%m%d"+".log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 新增 Hook
	logger.AddHook(lfHook)
}

func formatLog(c *gin.Context) {
	// 开始时间
	startTime := time.Now()

	// 处理请求
	c.Next()

	// 结束时间
	endTime := time.Now()

	// 执行时间
	latencyTime := endTime.Sub(startTime)

	// 请求方式
	reqMethod := c.Request.Method

	// 请求路由
	reqUri := c.Request.RequestURI

	// 状态码
	statusCode := c.Writer.Status()

	// 请求IP
	clientIP := c.ClientIP()

	// 日志格式
	localLogger.WithFields(logrus.Fields{
		"status_code":  statusCode,
		"latency_time": latencyTime,
		"client_ip":    clientIP,
		"req_method":   reqMethod,
		"req_uri":      reqUri,
	}).Info()
}

func DBLogger() *logrus.Logger {
	var dbLogger = logrus.New()
	// get log file
	logFilePath := config.DatasourceConfig.LogDir
	logFileName := config.DatasourceConfig.LogFile

	//init log
	initLog(logFilePath, logFileName, dbLogger)

	return dbLogger
}



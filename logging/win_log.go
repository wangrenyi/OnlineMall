package logging

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"onlinemall/config"
	"os"
	"path"
	"time"
)

//win10不能生成软连接,使用无分割日志
func WinLoggerHandler() gin.HandlerFunc {
	// get log file
	logFilePath := config.ServerConfig.LogDir
	logFileName := config.ServerConfig.LogFile

	return initWinLogger(logFilePath, logFileName, localLogger)
}

func initWinLogger(logFilePath string, logFileName string, logger *logrus.Logger) gin.HandlerFunc {
	//日志文件
	fileName := path.Join(logFilePath, logFileName)

	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	//设置输出
	logger.Out = src

	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{})

	return formatWinLog
}

func formatWinLog(c *gin.Context) {
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
	localLogger.Infof("| %3d | %13v | %15s | %s | %s |",
		statusCode,
		latencyTime,
		clientIP,
		reqMethod,
		reqUri,
	)
}

func WinDBLogger() *logrus.Logger {
	var dbLogger = logrus.New()
	// get log file
	logFilePath := config.DatasourceConfig.LogDir
	logFileName := config.DatasourceConfig.LogFile

	//init log
	initWinLogger(logFilePath, logFileName, dbLogger)

	return dbLogger
}

package logger

import (
	"com_sgrid/src/components/constant"
	"fmt"
	"os"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func CreateLogger(logName string) *logrus.Logger {
	logger := logrus.New()

	// 设置日志格式
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.DateTime,
	})
	SGRID_LOG_DIR := os.Getenv(constant.ENV_LOG_DIR)
	if SGRID_LOG_DIR == "" {
		return logger
	}
	logPath := filepath.Join(SGRID_LOG_DIR, fmt.Sprintf("%s.log", logName))
	fmt.Println(logPath)
	// 配置日志轮转
	writer, _ := rotatelogs.New(
		logPath+".%Y%m%d",
		rotatelogs.WithLinkName(logPath),
		rotatelogs.WithMaxAge(time.Duration(14*24)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
	)

	// 设置日志级别映射
	writerMap := lfshook.WriterMap{
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}

	// 设置 Hook
	logger.AddHook(lfshook.NewHook(
		writerMap,
		&logrus.JSONFormatter{
			TimestampFormat: time.DateTime,
		},
	))

	return logger
}

// 使用示例

// func init() {
// 	log_service := CreateLogger("service")
// 	log_service.Info("hello")
// }

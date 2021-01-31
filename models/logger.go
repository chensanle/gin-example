package models

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/chensanle/gin-example/helper"
	"github.com/gin-gonic/gin"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var log = logrus.StandardLogger()
var logger *logrus.Entry

const DebugLevel = logrus.DebugLevel

func InitLogger(path string, level logrus.Level) error {
	exist, err := pathExists(path)
	if err != nil {
		return err
	}
	if !exist {
		fmt.Printf("warning: invalid log path: %v\n", path)
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}

	log.SetOutput(ioutil.Discard) //使用hook输出日志，丢弃原有的write操作

	logrus.SetLevel(level)

	// fixme, need config
	logrus.SetOutput(os.Stdout)

	logMain, err := rotateLogs.New(path+"/main.log.%Y%m%d%H", rotateLogs.WithMaxAge(30*24*time.Hour),
		rotateLogs.WithRotationTime(time.Hour))
	if err != nil {
		return err
	}
	logError, err := rotateLogs.New(path+"/error.log.%Y%m%d%H", rotateLogs.WithMaxAge(30*24*time.Hour),
		rotateLogs.WithRotationTime(time.Hour))
	if err != nil {
		return err
	}

	lfHook := lfshook.NewHook(
		lfshook.WriterMap{
			logrus.DebugLevel: logMain,
			logrus.InfoLevel:  logMain,
			logrus.WarnLevel:  logMain,
			logrus.ErrorLevel: logError,
			logrus.FatalLevel: logError,
			logrus.PanicLevel: logError,
		},
		&logrus.JSONFormatter{},
	)

	log.AddHook(lfHook)

	ip, _ := helper.GetIp()
	logger = log.WithField("host", ip)

	return nil
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func WithContext(c *gin.Context, fields ...logrus.Fields) *logrus.Entry {
	if c == nil && logger == nil {
		return logrus.NewEntry(log)
	}

	if c == nil {
		return logger
	}

	if c.ClientIP() != "" {
		logger = logger.WithField("client-ip", c.ClientIP())
	}

	if c.GetInt("uid") != 0 {
		logger = logger.WithField("uid", c.GetInt("uid"))
	}

	if c.GetInt("target_uid") != 0 {
		logger = logger.WithField("uid", c.GetInt("target_uid"))
	}

	if c.Request.URL.Path != "" {
		logger = logger.WithField("path", c.Request.URL.Path)
	}

	if c.GetHeader("AppVersion") != "" {
		logger = logger.WithField("app_version", c.GetInt(c.GetHeader("AppVersion")))
	}

	if c.GetHeader("Platform") != "" {
		logger = logger.WithField("platform", c.GetInt(c.GetHeader("Platform")))
	}

	if len(fields) > 0 {
		for key, val := range fields[0] {
			logger = logger.WithField(key, val)
		}
	}

	return logger
}

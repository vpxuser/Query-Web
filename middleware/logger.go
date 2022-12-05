package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"log"
	"math"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	filePath := "log/log"
	//linkName := "latest_log.log"
	scr,err := os.OpenFile(filePath,os.O_RDWR|os.O_CREATE,0755)
	if err != nil {
		log.Println("err :",err)
	}
	logger := logrus.New()

	logger.Out = scr

	logger.SetLevel(logrus.DebugLevel)

	logWriter,_ := retalog.New(
		filePath+"%Y%m%d.log",
		retalog.WithMaxAge(7*24*time.Hour),
		retalog.WithRotationTime(24*time.Hour),
		//retalog.WithLinkName(linkName),
		)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel: logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel: logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	Hook := lfshook.NewHook(writeMap,&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.AddHook(Hook)

	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms",int(math.Ceil(float64(stopTime.Nanoseconds())/1000000.0)))
		hostName,err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		statusCode := ctx.Writer.Status()
		clientIp := ctx.ClientIP()
		userAgent := ctx.Request.UserAgent()
		dataSize := ctx.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := ctx.Request.Method
		path := ctx.Request.RequestURI

		entry := logger.WithFields(logrus.Fields{
			"HostName" : hostName,
			"status" : statusCode,
			"SpendTime" : spendTime,
			"Ip" : clientIp,
			"Method" : method,
			"Path" : path,
			"DataSize" : dataSize,
			"Agent" : userAgent,
		})
		if len(ctx.Errors) > 0 {
			entry.Error(ctx.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
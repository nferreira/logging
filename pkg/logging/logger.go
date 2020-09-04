package logging

import (
	"github.com/nferreira/app/pkg/service"
)

const (
	LogService = "LogService"
)

type Organization struct {
	Id   string
	Name string
	Unit string
}

type System struct {
	Environment string
	Id          string
	Hostname    string
	AppName     string
}

type Logger interface {
	service.Service
	Fatalf(correlationId string, format string, args ...interface{})
	Fatal(correlationId string, message string)
	Panicf(correlationId string, format string, args ...interface{})
	Panic(correlationId string, message string)
	Errorf(correlationId string, format string, args ...interface{})
	Error(correlationId string, message string)
	Infof(correlationId string, format string, args ...interface{})
	Info(correlationId string, message string)
	Warnf(correlationId string, format string, args ...interface{})
	Warn(correlationId string, message string)
	Debugf(correlationId string, format string, args ...interface{})
	Debug(correlationId string, message string)
}

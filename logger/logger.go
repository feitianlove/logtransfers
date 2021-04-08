package logger

import (
	goliblogger "github.com/feitianlove/golib/common/logger"
	"github.com/sirupsen/logrus"
)

var Ctrl *logrus.Logger

func init() {
	Ctrl = goliblogger.NewLoggerInstance()
}
func InitCtrlLog(conf *goliblogger.LogConf) error {
	logger, err := goliblogger.InitLogger(conf)
	if err != nil {
		return err
	}
	Ctrl = logger
	return nil
}

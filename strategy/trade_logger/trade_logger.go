package tradelogger

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var TradeLogger = func() *logrus.Logger {

	log := logrus.New()

	log.Formatter = new(logrus.TextFormatter)

	level, err := logrus.ParseLevel(logrus.TraceLevel.String())
	if err != nil {
		panic(err)
	}

	lumberjackLogger := &lumberjack.Logger{
		Filename:   "./logs/trading_log.log",
		MaxSize:    100,
		MaxBackups: 4,
		MaxAge:     30,
	}

	log.SetOutput(lumberjackLogger)

	log.Level = level
	return log

}

package logger

import "github.com/sirupsen/logrus"

var Logger *logrus.Logger

func SetupLogger() {	

	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})
	log.Info("logger initialized using logrus") 
	Logger = log
}
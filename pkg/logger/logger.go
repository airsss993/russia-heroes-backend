package logger

import "go.uber.org/zap"

var L *zap.Logger

func Init(env string) error {
	var err error
	if env == "prod" {
		L, err = zap.NewProduction()
	} else if env == "dev" {
		L, err = zap.NewDevelopment()
	}
	return err
}

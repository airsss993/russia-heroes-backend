package logger

import "go.uber.org/zap"

// L глобальный инстанс логгера
var L *zap.Logger

// Init инициализирует логгер в зависимости от окружения
func Init(env string) error {
	var err error
	if env == "prod" {
		L, err = zap.NewProduction()
	} else if env == "dev" {
		L, err = zap.NewDevelopment()
	}
	return err
}

package logs

import (
	"fmt"

	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

// InitLogger function to initialize zap sugar
func InitLogger() error {
	l, err := zap.NewDevelopment()

	if err != nil {
		_ = fmt.Errorf("Cant create zap logger %s", err.Error())
		return err
	}
	sugar = l.Sugar()
	return nil
}

// Log internal function prepared to logging
func Log() *zap.SugaredLogger {
	return sugar
}

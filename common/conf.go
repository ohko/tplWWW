// Package common 这个目录放一些配置和公用函数什么的
package common

import (
	"os"
	"time"

	"github.com/ohko/logger"
)

// var
var (
	LL           = logger.NewLogger(logger.NewDefaultWriter(&logger.DefaultWriterOption{Clone: os.Stdout, Path: "./log", Label: "tpler", Name: "log_"}))
	TimeLocation *time.Location
	SessionName  = "TPLER"
)

// Init 初始化系统配置
func Init() error {
	var err error

	TimeLocation, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return err
	}

	return nil
}

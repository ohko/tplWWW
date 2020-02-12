// Package common 这个目录放一些配置和公用函数什么的
package common

import (
	"os"
	"time"

	"github.com/ohko/logger"
)

// var
var (
	// 日志写到文件
	LLFile = logger.NewDefaultWriter(&logger.DefaultWriterOption{
		Clone: os.Stdout,
		Path:  "./log",
		Label: "tpler",
		Name:  "log_",
	})
	// 日志接口
	LL = logger.NewLogger(LLFile)
	// 时间地区
	TimeLocation *time.Location
	// Session名称
	SessionName = "TPLER"
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

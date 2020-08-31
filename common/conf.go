// Package common 这个目录放一些配置和公用函数什么的
package common

import (
	"os"
	"time"

	"github.com/ohko/logger"
)

// var
var (
	BuildInfo = ""
	// 日志写到文件
	LLFile = logger.NewDefaultWriter(&logger.DefaultWriterOption{
		CompressMode:  "day",     // 日志压缩模式 [month|day] month=按月压缩，day=按日压缩
		CompressCount: 3,         // 仅在按日压缩模式下有效，设置为压缩几天前的日志，支持大于等于1的数字
		CompressKeep:  30,        // 前多少次的压缩文件删除掉，支持month和day模式。默认为0，不删除。例如：1=保留最近1个压缩日志，2=保留最近2个压缩日志，依次类推。。。
		Clone:         os.Stdout, // 日志克隆输出接口
		Path:          "./log",   // 日志目录，默认目录：./log
		Label:         "tpler",   // 日志标签
		Name:          "log_",    // 日志文件名
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

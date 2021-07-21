package common

import (
	"github.com/bingoohuang/golog/pkg/logfmt"
	"github.com/bingoohuang/golog/pkg/spec"
	"github.com/bingoohuang/golog/pkg/timex"
)

func InitSetup() {
	var size spec.Size
	size.Parse("10M") // 最大单个日志文件10M
	layout := `%t{yyyy-MM-dd HH:mm:ss.SSS} [%-5l{length=5}] ☆ %msg ☆ %caller %fields %n`
	maxAge, _ := timex.ParseDuration("10d") // 最大保留3年
	gzipAge, _ := timex.ParseDuration("3d") // 归档压缩3天前的日志

	logs := logfmt.LogrusOption{
		Level:       "debug",
		LogPath:     "./logs/easyNmon.log",
		Rotate:      ".yyyy-MM-dd",
		MaxAge:      maxAge,
		GzipAge:     gzipAge,
		MaxSize:     int64(size),
		PrintColor:  true,
		PrintCaller: true,
		Stdout:      true,
		Layout:      layout,
	}
	logs.Setup(nil)
}

package common

import (
	"os"

	"github.com/bingoohuang/golog"
)

func InitLogs() {
	layout := `%t{yyyy-MM-dd HH:mm:ss.SSS} [%-5l{length=5}] ☆ %msg ☆ %caller %fields%n`
	spec := "printColor=true,file=./logs/easyNmon.log,maxSize=10M,maxAge=365d,gzipAge=3d,stdout=true"
	if *F.Debug {
		spec = "level=debug," + spec
	}
	os.Setenv("GOLOG_DEBUG", "off")
	golog.Setup(golog.Layout(layout), golog.Spec(spec))
}

package common

import (
	"easyNmon/pkg"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"time"

	"github.com/mzky/utils/cmd"
	"github.com/mzky/utils/memdb"
)

func Run(delay, count string) {
	cmdOptions := cmd.Options{Buffered: false, Streaming: true}
	envCmd := cmd.NewCmdOptions(cmdOptions, pkg.Njmon, "-n", "-s", delay, "-c", count)
	ticker := time.NewTicker(time.Second)
	envCmd.Start()
	//go kill(pkg.NjMon, time.Second*time.Duration(t*c+2))
	var m Mem
	m.DB = memdb.New()
	for range ticker.C {
		if envCmd.Stdout != nil {
			line, open := <-envCmd.Stdout
			if !open {
				envCmd.Stdout = nil
				continue
			}
			fmt.Println(line)
			var md MonData

			Handle(jsoniter.UnmarshalFromString(line, &md))
			jBytes, _ := jsoniter.Marshal(md)
			m.Parser(jBytes)
			Handle(m.DB.Save("data.json"))
			fmt.Println(m.GetKeys("SysInfo"))
		}
	}
}

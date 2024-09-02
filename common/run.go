package common

import (
	"easyNmon/pkg"
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/mzky/utils/cmd"
	"github.com/mzky/utils/memdb"
)

func Run(t, c int) {
	cmdOptions := cmd.Options{Buffered: false, Streaming: true}
	envCmd := cmd.NewCmdOptions(cmdOptions, pkg.NjMon, "-n", "-s", strconv.Itoa(t), "-c", strconv.Itoa(c))
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
			json.Unmarshal([]byte(line), &md)
			jBytes, _ := json.Marshal(md)
			m.Parser(jBytes)
			Handle(m.DB.Save("data.json"))
			fmt.Println(m.GetKeys("SysInfo"))
		}
	}
}

func kill(name string, d time.Duration) {
	<-time.After(d)
	ret := exec.Command("pidof", name)
	buf, _ := ret.Output()
	for _, v := range strings.Fields(string(buf)) {
		pid, _ := strconv.Atoi(v)
		fmt.Println("kill", v, name, syscall.Kill(pid, syscall.SIGKILL) == nil)
	}
}
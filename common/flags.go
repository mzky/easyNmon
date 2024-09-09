package common

import (
	"flag"
	"fmt"
	"github.com/gookit/color"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"syscall"
)

type Flag struct {
	IP        string
	Debug     bool
	Version   bool
	Port      string
	Dir       string
	Analysis  string
	Host      string
	NjmonPath string
	R         echo.Echo
	ReportDir string
	Address   string
}

func (f *Flag) InitFlag() {
	flag.BoolVar(&f.Debug, "debug", false, "Debug mode")
	flag.BoolVar(&f.Version, "v", false, "Show version")
	flag.StringVar(&f.Port, "p", "9999", "Service port")
	flag.StringVar(&f.Dir, "d", "report", "Default reporting directory")
	flag.StringVar(&f.NjmonPath, "n", "njmon", "Specify the njmon version for the platform")
	f.GetExternalIP()
	f.Address = fmt.Sprintf("http://%s:%s", f.IP, f.Port)
	color.Gray.Print(art)
	flag.Usage = f.usage
	flag.Parse()

	if f.Version {
		fmt.Println("Version: " + Version)
		fmt.Println("Compile: " + Compile)
		os.Exit(0)
	}

	if f.Analysis != "" {
		os.Exit(0)
	}

	fmt.Println("Management Page:", f.Address)

	f.ReportDir, _ = filepath.Abs(f.Dir) //绝对路径*dir
	syscall.Umask(0)
	if os.MkdirAll(f.ReportDir, os.ModePerm) != nil {
		logrus.Error("EasyNjmon does not have permission or the directory does not have permission to write!")
		os.Exit(0)
	}
}

func printf(format string, a ...interface{}) {
	fmt.Printf(format+"\n", a...)
}

func (f *Flag) usage() {
	printf("Version: %s", Version)
	printf("BuildTime: %s", Compile)
	printf("Usage: %s [OPTIONS] args", os.Args[0])
	printf("OPTIONS:")
	flag.PrintDefaults()
	printf("")
	printf("USAGES:")
	printf("   Examples of parameters")
	printf("      %s -d /mnt/reports", os.Args[0])
	printf("      %s -n ./njmon", os.Args[0])
	printf("   Management Page")
	printf("      %s", f.Address)
	printf("   API [GET]")
	printf("      Start Monitoring")
	printf("         %s/start?n=name&t=time&f=frequency", f.Address)
	printf("         [n] The [name] of the file to generate the report")
	printf("         [t] Monitoring [time] (Unit: minute)")
	printf("         [f] Monitoring [frequency] (Unit: seconds)")
	printf("      Stop Monitoring")
	printf("         %s/stop", f.Address)
	printf("      View Reports")
	printf("         %s/report", f.Address)
	printf("      Close Self")
	printf("         %s/close", f.Address)

}

const art = `
-888888888---------------------------------88b-----88--88-------------------------------------
-88----------------------------------------888b----88-----------------------------------------
-88--------,adPPYba,-,adPPYba,-8b-------d8-88-8b---88--88--,adba..adba,---,adPPba,---adPPba,--
-88aaaaaa--------'Y8-a8(-------'8b-----d8'-88--8b--88--88-88"--"88"--"8a-a8"----"8a-88"---"8a-
-88--------,adPPPP88--'"Y8ba,---'8b---d8'--88---8b-88--88-88----88----88-8b------d8-88-----88-
-88--------88,---,88-aa----)8a---'8b,d8'---88----8888--88-88----88----88-8a------8a-88-----88-
-888888888-'"8bbP8P--'"YbbdP"'-----&88'----88-----888--88-88----88----88--"adGGba"--88-----88-
-----------------------------------d8'-----------------,88-------------------------------------
---------------------------------ad8'----------------888P"-------------------------------------
`

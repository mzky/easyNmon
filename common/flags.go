package common

import (
	"flag"
	"fmt"
	"github.com/arsham/figurine/figurine"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"syscall"

	"github.com/labstack/echo/v4"
)

type Flag struct {
	IP        string
	Debug     bool
	V         bool
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
	flag.BoolVar(&f.V, "v", false, "\nShow version")
	flag.StringVar(&f.Port, "p", "9999", "Web service port")
	flag.StringVar(&f.Dir, "d", "report", "Default reporting directory")
	flag.StringVar(&f.Analysis, "a", "", "Specify the Nmon report file to generate HTML")
	flag.StringVar(&f.NjmonPath, "n", "njmon", "Specify the njmon version for the platform")

	figurine.Write(os.Stdout, "EasyNjmon", "Doom.flf")

	flag.Usage = f.usage
	flag.Parse()

	if f.V {
		fmt.Println("Version: " + Version)
		fmt.Println("Compile: " + Compile)
		os.Exit(0)
	}

	if f.Analysis != "" {
		os.Exit(0)
	}

	f.IP = GetExternalIP()
	f.Address = fmt.Sprintf("Management Page: http://%s:%s", f.IP, f.Port)
	fmt.Println(f.Address)

	f.ReportDir, _ = filepath.Abs(f.Dir) //绝对路径*dir
	syscall.Umask(0)
	if os.MkdirAll(f.ReportDir, os.ModePerm) != nil {
		logrus.Error("EasyNmon does not have permission or the directory does not have permission to write!")
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
	printf("      %s -a ./report/testName", os.Args[0])
	printf("      %s -d /mnt/reports", os.Args[0])
	printf("      %s -n ./nmon/nmon_centos7", os.Args[0])
	printf("   Web Management Page")
	printf("      %s", f.Address)
	printf("   Web Interface [GET]")
	printf("      Start monitoring")
	printf("         %s/start?n=name&t=30&f=30", f.Address)
	printf("         [n] The name of the file to generate the report")
	printf("         [t] The monitoring time (Unit: minute)")
	printf("         [f] This is the monitoring frequency (Unit: seconds)")
	printf("      Stop monitoring")
	printf("         %s/stop", f.Address)
	printf("      View Reports")
	printf("         %s/report", f.Address)
	printf("      Close EasyNmon")
	printf("         %s/close", f.Address)

}

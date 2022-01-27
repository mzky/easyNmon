package common

import (
	"easyNmon/utils"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"syscall"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

var F Flag

type Flag struct {
	IP        string
	Debug     *bool
	V         *bool
	Port      *string
	Dir       *string
	Analysis  *string
	Host      *string
	NmonPath  *string
	R         *gin.Engine
	ReportDir string
}

func InitFlag() {
	F.Debug = flag.Bool("debug", false, "Debug mode")
	F.V = flag.Bool("v", false, "\nShow version")
	F.Port = flag.String("p", "9999", "Web service port")
	F.Dir = flag.String("d", "report", "Default reporting directory")
	F.Analysis = flag.String("a", "", "Specify the Nmon report file to generate HTML")
	F.NmonPath = flag.String("n", "nmon/nmon", "Specify the nmon version for the platform")

	flag.Usage = usage
	flag.Parse()

	if *F.V {
		fmt.Println("Version: " + Version)
		fmt.Println("Compile: " + Compile)
		os.Exit(0)
	}

	if *F.Analysis != "" {
		path, name := filepath.Split(*F.Analysis)
		utils.GetNmonReport(path, name)
		os.Exit(0)
	}

	if !*F.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	F.ReportDir, _ = filepath.Abs(*F.Dir) //绝对路径*dir
	syscall.Umask(0)
	if os.MkdirAll(F.ReportDir, os.ModePerm) != nil {
		logrus.Error("EasyNmon does not have permission or the directory does not have permission to write!")
		os.Exit(0)
	}
}

func printf(format string, a ...interface{}) {
	fmt.Printf(format+"\n", a...)
}

func usage() {
	F.IP = utils.GetExternalIP()
	address := fmt.Sprintf("http://%s:%s", F.IP, *F.Port)
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
	printf("      %s", address)
	printf("   Web Interface [GET]")
	printf("      Start monitoring")
	printf("         %s/start?n=name&t=30&f=30", address)
	printf("         [n] The name of the file to generate the report")
	printf("         [t] The monitoring time (Unit: minute)")
	printf("         [f] This is the monitoring frequency (Unit: seconds)")
	printf("      Stop monitoring")
	printf("         %s/stop", address)
	printf("      View Reports")
	printf("         %s/report", address)
	printf("      Close EasyNmon")
	printf("         %s/close", address)

}

package main

import (
	"easyNmon/internal"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	// 版本号
	Version = "<Version>"
	// 构建时间
	BuildTime = "<BuildTime>"
	//全局
	ReportDir string
	NmonPath  string
)

func main() {
	ip := "127.0.0.1"
	netaddr, _ := net.InterfaceAddrs()
	networkIp, _ := netaddr[1].(*net.IPNet)
	if !networkIp.IP.IsLoopback() && networkIp.IP.To4() != nil {
		ip = networkIp.IP.String()
	}

	version := flag.Bool("v", false, "version:显示版本号")
	port := flag.String("p", "9999", "port:默认监听端口9999,自定义端口加 -p 端口号\n示例：./easyNmon -p 9999")
	dir := flag.String("d", "report", "directory:指定生成报告的路径\n示例：./easyNmon -d /mnt/rep")
	analysis := flag.String("a", "", "analysis:生成html图表，参数指定nmon报告文件，同目录生成html图表\n示例：./easyNmon -a ./report/nmonTestName")
	nmonpath := flag.String("np", "nmon/nmon", "nmonpath：指定对应系统版本的nmon文件\n示例：./easyNmon -np ./nmon/nmon_xxx")
	readme := "接口(Get)：\n\t/start\t启动监控,接口方式时,所有参数非必选\n\t\t参数n为生成报告的文件名,\n\t\t参数t为监控时长(单位分钟),\n\t\t参数f为监控频率，每隔多少秒收集一次;\n\t\thttp://" + ip + ":9999/start?n=name&t=30&f=30\n\t/stop\t停止所有监控任务：\n\t\thttp://" + ip + ":9999/stop\n\t/report\t查看报告：\n\t\thttp://" + ip + ":9999/report\n\t/close\t关闭自身：\n\t\thttp://" + ip + ":9999/close\n管理页面：\n\t通过浏览器访问web管理页面：\n\thttp://" + ip + ":9999"
	flag.Bool("操作说明", false, readme)
	flag.Parse()

	ReportDir = *dir
	NmonPath = *nmonpath
	err := os.MkdirAll(ReportDir, 755)
	if err!=nil{
		fmt.Println("easyNmon启动权限不足!")
		return
	}
	if *version {
		fmt.Println("Version: " + Version)
		fmt.Println("BuildTime: " + BuildTime)
		os.Exit(0)
	}

	analy := *analysis
	if analy != "" {
		paths, fileName := filepath.Split(analy)
		internal.GetNmonReport(paths, fileName)
		os.Exit(0)
	}

	sysinfo := internal.SysInfo()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	
	//重定向首页--解决静态文件与接口共存
	// r.GET("/", func(c *gin.Context) {
	// 	//c.Request.URL.Host = "http://127.0.0.1:8090"
	// 	c.Header("Content-Type", "text/html; charset=utf-8")
	// 	c.HTML(200, "./web/index.html", nil)
	// 	//c.Request.URL.Path = "/web"
	// 	defer r.HandleContext(c)
	// })

	//首页
	r.LoadHTMLFiles("web/index.html")
	r.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.Static("/js", "web/js")
	// 浏览报告
	r.StaticFS("/report", http.Dir(ReportDir))
	//生成报告,用于实时更新报告
	r.GET("/generate/:name/", func(c *gin.Context) {
		name := c.Param("name")
		internal.GetNmonReport(filepath.Join(ReportDir, name), name[:len(name)-14])
		c.JSON(http.StatusOK, gin.H{
			"message": "更新生成报告",
		})
	})
	r.GET("/sysinfo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": sysinfo,
		})
	})
	r.GET("/start", start)
	r.GET("/close", close)
	r.GET("/stop", stop)
	readme = strings.ReplaceAll(readme, "9999", *port)
	fmt.Println(readme)
	fmt.Println("执行的nmon文件：" + *nmonpath)
	fmt.Println("存放报告的目录：" + *dir)
	r.Run(":" + *port) // listen and serve on 0.0.0.0:8080
	fmt.Println("easyNmon启动失败，查看端口是否被占用!")
}

func start(c *gin.Context) { // 格式 ?n=name&t=time&f=60 参数均可为空 默认30分钟
	name := c.DefaultQuery("n", "name")    // 取name值
	timeStr := c.DefaultQuery("t", "30")   // 取执行时长,单位分钟
	frequency := c.DefaultQuery("f", "30") //频率，多少秒取一次
	fileName := strings.Join([]string{name,time.Now().Format("20060102150405")},"")

	t, _ := strconv.Atoi(timeStr)
	f, _ := strconv.Atoi(frequency)
	if t == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": string("执行错误，请检查参数是否正确！"),
		})
		return
	}
	go func() {
		fp := filepath.Join(ReportDir, fileName)
		os.MkdirAll(fp, 777)

		buf, err := ioutil.ReadFile("web/chart/index.html")
		if err != nil {
			fmt.Println(err)
		}
		content := string(buf)
		newContent := strings.ReplaceAll(content, "{{loopTime}}", strings.Join([]string{frequency,"000"},""))

		//重新写入
		ioutil.WriteFile(filepath.Join(fp, "index.html"), []byte(newContent), 0)

		exec.Command("cp", "-f", "web/js/echarts.min.js", fp).Run()
		//	exec.Command("cp", "-f", "web/chart/index.html", fp).Run()
		exec.Command("/bin/bash", "-c", strings.Join([]string{NmonPath,"-f -t -s",frequency,"-c",strconv.Itoa(t*60/f),"-m",fp,"-F",name}," ")).Run()
		time.Sleep(time.Second * 2)
		internal.GetNmonReport(fp, name)
		time.Sleep(time.Second * time.Duration(t*60+2))
		internal.GetNmonReport(fp, name)
	}()
	c.JSON(http.StatusOK, gin.H{
		"message": strings.Join([]string{"已执行", name,"场景，监控时长",timeStr,"分钟，频率为" ,frequency,"秒！"},""),
	})
}

func close(c *gin.Context) { //结束自身进程
	c.JSON(http.StatusOK, gin.H{
		"message": "已结束EasyNmon服务!",
	})
	go func() {
		getAllReport()
		killNmon()
		os.Exit(0)
	}()
}

func stop(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "已结束所有服务器监控任务!",
	})
	go func() {
		getAllReport()
		killNmon()
	}()
}

//重新生成所有报告
func getAllReport() {
	list := getDirList(ReportDir)
	for _, v := range list {
		internal.GetNmonReport(filepath.Join(ReportDir, v), v[:len(v)-14])
	}
}

//获取文件夹列表
func getDirList(dirpath string) []string {
	var dir_list []string
	filepath.Walk(dirpath,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				if path != dirpath {
					dir_list = append(dir_list, path[len(ReportDir)+1:])
					return nil
				}
			}
			return nil
		})
	return dir_list
}

//杀掉所有nmon进程
func killNmon() {
	ret := exec.Command("pidof", NmonPath)
	buf, err := ret.Output()
	if err == nil {
		pids := strings.Split(strings.ReplaceAll(string(buf), "\n", ""), " ")
		for _, value := range pids {
			pid, _ := strconv.Atoi(value)
			syscall.Kill(pid, syscall.SIGKILL)
		}
	}
}

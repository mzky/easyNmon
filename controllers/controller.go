package controllers

import (
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Start(c *gin.Context) { // 格式 ?n=name&t=time&f=60 参数均可为空 默认30分钟
	name := c.DefaultQuery("n", "name")    // 取name值
	timeStr := c.DefaultQuery("t", "30")   // 时长 单位分钟
	frequency := c.DefaultQuery("f", "30") //频率，多少秒取一次
	//fileName := strings.Join([]string{name, time.Now().Format("20060102150405")}, "")

	//go func() {
	//	fullPath := filepath.Join(F.ReportDir, fileName)
	//	os.MkdirAll(fullPath, os.ModePerm)
	//
	//	buf := common.Wfs.Files["/chart/index.html"].Data
	//	content := string(buf)
	//	newContent := strings.ReplaceAll(content, "{{loopTime}}", strings.Join([]string{frequency, "000"}, ""))
	//
	//	//重新写入
	//	ioutil.WriteFile(filepath.Join(fullPath, "index.html"), []byte(newContent), 0)
	//
	//	utils.InitFile(common.Wfs, "/js/echarts.min.js", filepath.Join(fullPath, "echarts.min.js"))
	//	//exec.Command("cp", "-f", "web/js/echarts.min.js", fullPath).Run()
	//	os.Chmod(filepath.Join(fullPath, "index.html"), os.ModePerm)
	//	os.Chmod(filepath.Join(fullPath, "echarts.min.js"), os.ModePerm)
	//	os.Chmod(filepath.Join(fullPath, name), os.ModePerm)
	//
	//	t, _ := strconv.Atoi(timeStr)
	//	f, _ := strconv.Atoi(frequency)
	//
	//	lib.Agent(fullPath, name, frequency, strconv.Itoa(t*60/f))
	//
	//	<-time.After(1 * time.Second)
	//	utils.GetNmonReport(fullPath, name)
	//}()
	logrus.Info("已执行%s场景，监控时长%s分钟，频率为%s秒！", name, timeStr, frequency)
	c.JSON(http.StatusOK, gin.H{
		"message": strings.Join([]string{"已执行", name, "场景，监控时长", timeStr, "分钟，频率为", frequency, "秒！"}, ""),
	})
}

func Close(c *gin.Context) { //结束自身进程
	logrus.Info("已结束EasyNmon服务!")
	c.JSON(http.StatusOK, gin.H{
		"message": "已结束EasyNmon服务!",
	})
	go func() {
		getAllReport()
		killNmon()
		os.Exit(0)
	}()
}

func Stop(c *gin.Context) {
	logrus.Info("已结束所有服务器监控任务!")
	c.JSON(http.StatusOK, gin.H{
		"message": "已结束所有服务器监控任务!",
	})
	go func() {
		getAllReport()
		killNmon()
	}()
}

// 重新生成所有报告
func getAllReport() {
	//list := getDirList(common.ReportDir)
	//for _, v := range list {
	//	utils.GetNmonReport(filepath.Join(common.ReportDir, v), v[:len(v)-14])
	//}
}

// 获取文件夹列表
func getDirList(dirpath string) []string {
	var dirList []string
	filepath.Walk(dirpath,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				if path != dirpath {
					//dirList = append(dirList, path[len(common.ReportDir)+1:])
					return nil
				}
			}
			return nil
		})
	return dirList
}

// 杀掉所有nmon进程
func killNmon() {
	//ret := exec.Command("pidof", common.NmonPath)
	//buf, err := ret.Output()
	//if err == nil {
	//	pids := strings.Split(strings.ReplaceAll(string(buf), "\n", ""), " ")
	//	for _, value := range pids {
	//		pid, _ := strconv.Atoi(value)
	//		syscall.Kill(pid, syscall.SIGKILL)
	//	}
	//}

}

func GetSystemInfo(c *gin.Context) {
	sysInfo := utils.SysInfo()
	logrus.Info(sysInfo)
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.JSON(http.StatusOK, gin.H{"message": sysInfo})
}

func ShowIndex(c *gin.Context) {
	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Path = "/web"
		req.URL.Host = c.Request.Host
		req.Host = c.Request.Host
	}
	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(c.Writer, c.Request)
}

func Generate(c *gin.Context) {
	name := c.Param("name")
	logrus.Info("更新%s报告", name)
	c.JSON(http.StatusOK, gin.H{
		"message": "更新生成报告",
	})
	go func() {
		//utils.GetNmonReport(filepath.Join(common.ReportDir, name), name[:len(name)-14])
	}()
}

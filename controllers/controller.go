package controllers

import (
	"easyNmon/common"
	"easyNmon/pkg"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func Start(c echo.Context) error { // 格式 ?n=name&t=time&f=60 参数均可为空 默认30分钟
	name := c.QueryParam("n")      // 取name值
	timeStr := c.QueryParam("t")   // 时长 单位分钟
	frequency := c.QueryParam("f") //频率，多少秒取一次
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

	ms := fmt.Sprintf("已执行%s场景，监控时长%s分钟，频率为%s秒！", name, timeStr, frequency)
	logrus.Info(ms)
	return c.JSON(http.StatusOK, common.RspOK(ms, ""))
}

func Close(c echo.Context) error { //结束自身进程
	logrus.Info("已结束EasyNmon服务!")
	go func() {
		getAllReport()
		killNjmon()
		os.Exit(0)
	}()
	return c.JSON(http.StatusOK, common.RspOK("已结束EasyNmon服务!", ""))
}

func Stop(c echo.Context) error {
	logrus.Info("已结束所有服务器监控任务!")
	go func() {
		getAllReport()
		killNjmon()
	}()
	return c.JSON(http.StatusOK, common.RspOK("已结束所有服务器监控任务!", ""))
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

// 杀掉所有njmon进程
func killNjmon() {
	<-time.After(time.Second)
	ret := exec.Command("pidof", pkg.Njmon)
	buf, _ := ret.Output()
	for _, v := range strings.Fields(string(buf)) {
		pid, _ := strconv.Atoi(v)
		logrus.Warnln("kill", v, pkg.Njmon, syscall.Kill(pid, syscall.SIGKILL) == nil)
	}
}

func GetSystemInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, common.RspOK(common.SysInfo(), ""))
}

func Generate(c echo.Context) error {
	name := c.Param("name")
	logrus.Info("更新%s报告", name)
	return c.JSON(http.StatusOK, common.RspOK("更新生成报告", ""))
}

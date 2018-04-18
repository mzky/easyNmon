# 为了方便多场景批量性能测试，用golang写了个监控程序，可以通过get url方式启动和停止nmon服务，
非常适合配合Loadrunner性能测试框架使用

1.下载解压到被测性能的服务器，并附权限
执行./monitor& 即可（后缀加&为后台执行）

2.默认端口8080，如需修改 加上-port 端口号 如下图

3.执行后通过loadrunner调用开始和结束

4.Loadrunner例子：

将脚本放到init下，每次测试仅执行一次

	web_custom_request("stop",//停止所有nmon监控
		       "URL=http://192.168.136.91:8080/stop",
                       "Method=GET",
		       LAST);
	web_custom_request("pdf",//pdf文字修改为实际场景名，不支持中文
		       "URL=http://192.168.136.91:8080/pdf",
                       "Method=GET",
		       LAST);

5.nmon运行结束后，自动生成html格式报告

有问题可以联系我：mzky@163.com

说明：
go的http框架采用gin：https://gin-gonic.github.io/gin/
图表插件采用百度的echarts：http://echarts.baidu.com/
新版nmon下载地址：http://nmon.sourceforge.net/
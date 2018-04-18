# 为了方便多场景批量性能测试，用golang写了个监控程序，可以通过get url方式启动和停止nmon服务，
非常适合配合Loadrunner性能测试框架使用
nmon服务包下载：monitor.zip

下载解压到被测性能的服务器，并附权限



执行./monitor& 即可（后缀加&为后台执行）


默认端口8080，如需修改 加上-port 端口号 如下图


执行时可通过loadrunner调用开始和结束

例如下图的pdf，那么文件名为pdf.nmon

如将pdf修改为stop，将停止所有nmon进程

Loadrunner例子：

将脚本放到init下，每次测试仅执行一次

	web_custom_request("stop",//停止所有nmon监控
		       "URL=http://192.168.136.91:8080/stop",
                       "Method=GET",
		       LAST);
	web_custom_request("pdf",//pdf文字修改为实际场景名，不支持中文
		       "URL=http://192.168.136.91:8080/pdf",
                       "Method=GET",
		       LAST);
执行效果：



如需go源码，联系我获取

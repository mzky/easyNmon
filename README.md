为了方便多场景批量性能测试，用golang写了个监控程序，可以通过get url方式启动和停止nmon服务，非常适合配合Loadrunner性能测试框架使用，可以批量执行场景并生成报告

使用说明：
https://www.jianshu.com/p/c7c36ba14d3e




更新日志：

2018.04.27

1.变更get参数格式，增加监控时间参数

2.使用日期+场景名方式保存报告，避免场景重复



后期规划：

与LR批量测试自动化框架结合 达到自动执行性能，自动监控服务器，自动生成html报告（包括TPS、RT和服务器性能报告）





联系：mzky@163.com


引用：

go的http框架采用gin：https://gin-gonic.github.io/gin/

图表插件采用百度的echarts：http://echarts.baidu.com/

新版nmon下载地址：http://nmon.sourceforge.net/


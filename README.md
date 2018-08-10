# 正式更名为：easyNmon


为了方便多场景批量性能测试，用golang写了个监控程序，可以通过get url方式启动和停止nmon服务，
适合配合Loadrunner和jmeter进行性能测试，可以z做到批量执行场景并生成监控报告！

## 使用说明：

https://www.jianshu.com/p/c7c36ba14d3e

执行文件下载：（以下执行文件不包含源码）

https://pan.baidu.com/s/1_A5TnW_tG1bTmX-gCNjOzw



#
## 更新日志：

### 2018.08.08

1.将nmonCTL.sh去掉，功能集成到go代码中

2.未到nmon设置的预定时间，结束监控服务和nmon进程时，自动生成index.html图表页面文件

3.将-port参数改为-p

4.增加build脚本，构建go代码和进行upx压缩


### 2018.06.05

1.监控接口返回值添加报告列表的url地址，可自己使用地址在浏览器中查看html报告

2.使用upx将执行文件压缩，打包添加到wiki，方便大家仅下载工具，不需要clone工程


### 2018.05.16

1.增加-h帮助和示例

2.增加/report在线显示图表报告

3.修改参数n为文件名，t为时长

4.增加/close关闭自身接口

5.修改线程执行方式


### 2018.04.27

1.变更get参数格式，增加监控时间参数

2.使用日期+场景名方式保存报告，避免场景重复


#
## 后期规划：

与LR和jmeter批量测试自动化框架结合 达到自动执行性能，自动监控服务器，自动生成html报告（包括TPS、RT和服务器性能报告）


jmeter4.0的html报告汉化：

https://github.com/mzky/jmeter4.0-cn-report-template

#
问题反馈：mzky@163.com



#
## 引用：

go的http框架采用gin：https://gin-gonic.github.io/gin/

图表插件采用百度的echarts：http://echarts.baidu.com/

新版nmon下载地址：http://nmon.sourceforge.net/





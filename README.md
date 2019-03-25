# EasyNmon

为了方便多场景批量性能测试，用golang写了个监控程序，可以通过get url方式启动和停止nmon服务，
适合配合Loadrunner和jmeter进行性能测试，可以做到批量执行场景并生成监控报告！

## 使用说明：

https://www.jianshu.com/p/c7c36ba14d3e


## 执行文件下载：（以下执行文件不包含源码）

https://pan.baidu.com/s/1XCeNQPMtymlI79kgNCg1ZA


### 为方便沟通，建了一个QQ群：
点击链接加入群聊【EasyNmon交流】：https://jq.qq.com/?_wv=1027&k=5sgrpm9



#
## 更新日志：
### 2019.03.21
1.修复用户提交的shell脚本缺陷

2.修复网络数据的写为负值

### 2019.01.31

1.之前因功能简单，所以写到一个main函数中，现计划新增功能，将main函数内功能分解成子函数

2.增加执行命令后的提示，如非双层网络，可直接复制提示的地址进行访问，更便捷

3.因8080经常与tomcat端口冲突，默认端口改为9999


### 2018.11.19

1.增加报告图表平均值线

2.增加web管理页面，可以通过管理页面提交监控、结束监控和查看报告


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


jmeter4.0/5.0的html报告汉化模版：

https://github.com/mzky/jmeter4.0-cn-report-template

https://github.com/mzky/jmeter5.0-cn-report-template



#
## 引用：

go的http框架采用gin：https://gin-gonic.github.io/gin/

图表插件采用百度的echarts：http://echarts.baidu.com/

新版nmon下载地址：http://nmon.sourceforge.net/

#注意： 默认nmon为CentOS版本（CentOS6.5~7.4正常），Ubuntu和SUSE需要下载对应版本的nmon替换（SUSE11.4测试正常）



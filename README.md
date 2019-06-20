### 因新版改动过大，并且不再使用Nmon作为采集工具，另开一个分支项目
新项目将支持mysql、pgsql、tomcat、redis、Nginx等监控

https://github.com/mzky/mesro

#
### EasyNmon将继续完善和优化现有功能
#
# EasyNmon

为了方便多场景批量性能测试，用golang写了个监控程序，可以通过get url方式启动和停止nmon服务，
适合配合Loadrunner和jmeter进行性能测试，可以做到批量执行场景并生成监控报告！

#### easyNmon的目标很明确：简单、轻量、绿色,在不需要安装任何语言环境和插件的情况下进行Linux系统资源监控
如在固定服务器上进行长期监控，建议使用open-falcon、Telegraf+Influxdb+grafana或NetData等优秀的监控工具

#
## 使用说明：

https://www.jianshu.com/p/c7c36ba14d3e

#
## 执行文件下载：（以下执行文件不包含源码）
https://github.com/mzky/easyNmon/releases

镜像：
https://pan.baidu.com/s/1XCeNQPMtymlI79kgNCg1ZA

#
### 为方便沟通，建了一个QQ群：
点击链接加入群聊【EasyNmon交流】：https://jq.qq.com/?_wv=1027&k=5sgrpm9


#
### ☆ 特别感谢[ElectricBubble](https://github.com/ElectricBubble)提交的代码，使EasyNmon实现完全go语言编码
#


## 更新日志：

https://github.com/mzky/easyNmon/wiki/ChangeLog



#
## 近期规划：（将在mesro项目中实现）

1.增加系统识别，计划支持ubuntu、suse、centos

2.去掉shell脚本，全部使用go编写

3.通过模版导出word报告（可能会通过多选生成汇总报告）

#
## 长期规划：（将在mesro项目中实现）

与LR和jmeter批量测试自动化框架结合 达到自动执行性能，自动监控服务器，自动生成html报告（包括TPS、RT和服务器性能报告）

jmeter4.0/5.0的html报告汉化模版：

https://github.com/mzky/jmeter4.0-cn-report-template

https://github.com/mzky/jmeter5.0-cn-report-template


#注意： 默认nmon为CentOS版本（CentOS6.5~7.4正常），Ubuntu和SUSE需要下载对应版本的nmon替换（SUSE11.4测试正常）

#
### 感谢：

nmon：http://nmon.sourceforge.net/

go的web框架gin：https://github.com/gin-gonic/gin

图表插件echarts：http://echarts.baidu.com/

前端amazeui：http://amazeui.org




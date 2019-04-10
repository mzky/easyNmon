# EasyNmon

为了方便多场景批量性能测试，用golang写了个监控程序，可以通过get url方式启动和停止nmon服务，
适合配合Loadrunner和jmeter进行性能测试，可以做到批量执行场景并生成监控报告！


## 使用说明：

https://www.jianshu.com/p/c7c36ba14d3e


## 执行文件下载：（以下执行文件不包含源码）

https://pan.baidu.com/s/1XCeNQPMtymlI79kgNCg1ZA


## 为方便沟通，建了一个QQ群：
点击链接加入群聊【EasyNmon交流】：https://jq.qq.com/?_wv=1027&k=5sgrpm9


## 更新日志：

https://github.com/mzky/easyNmon/wiki/更新日志


## 近期规划：

1.增加系统识别，计划支持ubuntu、suse、centos

2.去掉shell脚本，全部使用go编写

3.通过模版导出word报告（可能会通过多选生成汇总报告）


## 长期规划：

与LR和jmeter批量测试自动化框架结合 达到自动执行性能，自动监控服务器，自动生成html报告（包括TPS、RT和服务器性能报告）

jmeter4.0/5.0的html报告汉化模版：

https://github.com/mzky/jmeter4.0-cn-report-template

https://github.com/mzky/jmeter5.0-cn-report-template



## 引用与感谢：

go的web框架采用gin：https://github.com/gin-gonic/gin

图表插件采用百度的echarts：http://echarts.baidu.com/

前端amazeui：http://amazeui.org

新版nmon下载地址：http://nmon.sourceforge.net/


#注意： 默认nmon为CentOS版本（CentOS6.5~7.4正常），Ubuntu和SUSE需要下载对应版本的nmon替换（SUSE11.4测试正常）



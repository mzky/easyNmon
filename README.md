
#
# EasyNmon

为了方便多场景批量性能测试，用golang写了个监控程序，可以通过get url方式启动和停止nmon服务，
适合配合Loadrunner和jmeter进行性能测试，可以做到批量执行场景并生成监控报告！

#### 弱水三千只取一瓢，easyNmon的目标很明确：简单、轻量、绿色,在不需要安装任何语言环境和插件的情况下进行Linux系统资源监控
如在固定服务器上进行长期监控，建议使用open-falcon、Telegraf+Influxdb+grafana或NetData等优秀的监控工具

##
## 操作说明：

http://mzky.cc/post/9.html

##
## 执行文件下载：（以下执行文件不包含源码）
https://github.com/mzky/easyNmon/releases

镜像：
https://pan.baidu.com/s/1XCeNQPMtymlI79kgNCg1ZA

##
## 为方便沟通，建了一个QQ群：
点击链接加入群聊【EasyNmon交流】：https://jq.qq.com/?_wv=1027&k=5sgrpm9

##
## 扩展工具

goodhal的批量部署及监控服务：https://gitee.com/goodhal/ezNmon-Manager

jmeter优化版：https://github.com/mzky/Jmeter-Extension

## 注意

### 1.x版使用nmon，最新版本下载 https://github.com/mzky/easyNmon/releases/tag/v1.9

### 2.x版使用njmon，最新版本下载 https://github.com/mzky/easyNmon/releases/latest

## 巨人肩膀：

njmon：https://nmon.sourceforge.io/pmwiki.php?n=Site.Njmon

1.x版使用gin框架：https://github.com/gin-gonic/gin

2.x版使用echo框架（生成的二进制文件非常小）：https://github.com/labstack/echo

图表插件echarts：http://echarts.baidu.com/

前端amazeui：http://amazeui.org



## FAQ
1、无法创建报告文件（html）

查看是否权限正常，非root用户可以将程序放在当前用户的目录下，例如/home/user

2、无法创建data文件（不显示图表）


3、同一个架构不同系统可以使用同一个二进制文件，但需要安装缺少的依赖包

openAnolis、openEuler、CentOS等系统，有可能需安装依赖包： yum install ncurses*

Ubuntu、debian等系统，有可能需安装依赖包： apt install libncurses5




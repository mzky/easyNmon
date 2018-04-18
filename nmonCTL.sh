#!/bin/bash
if [ $1 == "stop" ];then
ps -ef|grep nmon16e|grep -v grep|awk {'print $2'}|xargs kill -9 
else
cp -rf templet report/$1
./nmon16e -f -t -s 30 -c 62 -m report/$1 -F $1.nmon 
cd report/$1 
sleep 1866
./toHtml.sh $1
fi

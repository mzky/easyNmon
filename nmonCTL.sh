#!/bin/bash
file=$1_`date "+%Y%m%d%H%M%S"`
cp -rf templet report/$file
./nmon -f -t -s $2 -c 60 -m report/$file -F $1
cd report/$file
sleep `expr 60 \* $2 + 2`
./toHtml.sh $1

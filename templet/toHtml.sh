#!/bin/bash
#将nmon报告筛选出关键信息
array=(ZZZZ CPU_ALL DISKREAD DISKWRITE MEM NET,)
for name in ${array[@]}  
do 
	cat $1 |grep $name|awk -F',' '{OFS=","}{NF=NF;$1="tmp";print}'|sed 's/tmp,//g' >$name
done  
#net总列数
n=`cat NET,|awk -F',' 'NR==1{print NF}'`
#net中间的列数
c=$(((n+1)/2+1))
#net取read列
cat NET,|awk -F',' 'BEGIN{n='$n'}{for(i=2;i<n/2;i++)printf $i",";print $i}' >read
#net取write列
cat NET,|awk -F',' 'BEGIN{c='$c';n='$n'}{for(i=c;i<n;i++)printf "-"$i",";print $i}' >write
#汇总总数
cat read |awk -F',' '{n=0;for(i=1;i<=NF;i++) n+=$i;print n }' >NETREAD
cat write|awk -F',' '{n=0;for(i=1;i<=NF;i++) n+=$i;print n }' >NETWRITE
cp -f templet index.html
#替换脚本名称
sed -i "s/scripts/$1/g" index.html
xAxisdatas=""
cpuUsers=""
cpuSyss=""
cpuWaits=""
memtotals=""
memfrees=""
actives=""
NetReads=""
NetWrites=""
DiskReads=""
DiskWrites=""
#取时间列表
while read LINE
do
	if [[ "$xAxisdatas" == "" ]]
	then	
		xAxisdatas="\"`echo $LINE |awk -F',' '{print $2}'`\""
	else
		xAxisdatas=$xAxisdatas",\"`echo $LINE |awk -F',' '{print $2}'`\""
	fi
done < ZZZZ
#xAxisdatas=($xAxisdatas)
sed -i "s/xAxisdatas/$xAxisdatas/g" index.html
#取CPU指标
while read LINE
do
	if [[ "$cpuUsers" == "" ]]
	then	
		cpuUsers=`echo $LINE |awk -F',' '{print $2}'`
		cpuSyss=`echo $LINE |awk -F',' '{print $3}'`
		cpuWaits=`echo $LINE |awk -F',' '{print $4}'`
	else
		cpuUsers=$cpuUsers","`echo $LINE |awk -F',' '{print $2}'`
		cpuSyss=$cpuSyss","`echo $LINE |awk -F',' '{print $3}'`
		cpuWaits=$cpuWaits","`echo $LINE |awk -F',' '{print $4}'`
	fi
done < CPU_ALL
cpuUsers=(`echo $cpuUsers|awk -F',' '{OFS=","}{NF=NF;$1="?";print}'|sed 's/?,//g'`)
cpuSyss=(`echo $cpuSyss|awk -F',' '{OFS=","}{NF=NF;$1="?";print}'|sed 's/?,//g'`)
cpuWaits=(`echo $cpuWaits|awk -F',' '{OFS=","}{NF=NF;$1="?";print}'|sed 's/?,//g'`)
sed -i "s/cpuUsers/$cpuUsers/g" index.html
sed -i "s/cpuSyss/$cpuSyss/g" index.html
sed -i "s/cpuWaits/$cpuWaits/g" index.html
#取mem列表
while read LINE
do
	if [[ "$memfrees" == "" ]]
	then	
		memtotals=`echo $LINE |awk -F',' '{print $2}'`
		memfrees=`echo $LINE |awk -F',' '{print $6}'`
		actives=`echo $LINE |awk -F',' '{print $12}'`
	else
		memtotals=$memtotals","`echo $LINE |awk -F',' '{print $2}'`
		memfrees=$memfrees","`echo $LINE |awk -F',' '{print $6}'`
		actives=$actives","`echo $LINE |awk -F',' '{print $12}'`
	fi
done < MEM
memtotals=(`echo $memtotals|awk -F',' '{OFS=","}{NF=NF;$1="?";print}'|sed 's/?,//g'`)
memfrees=(`echo $memfrees|awk -F',' '{OFS=","}{NF=NF;$1="?";print}'|sed 's/?,//g'`)
actives=(`echo $actives|awk -F',' '{OFS=","}{NF=NF;$1="?";print}'|sed 's/?,//g'`)
sed -i "s/memtotals/$memtotals/g" index.html
sed -i "s/memfrees/$memfrees/g" index.html
sed -i "s/actives/$actives/g" index.html
#取net指标
while read LINE
do
	if [[ "$NetReads" == "" ]]
	then	
		NetReads=`echo $LINE |awk -F',' '{print $1}'`
	else
		NetReads=$NetReads","`echo $LINE |awk -F',' '{print $1}'`
	fi
done < NETREAD
NetReads=(`echo $NetReads|awk -F',' '{OFS=","}{NF=NF;$1="?";print}'|sed 's/?,//g'`)
sed -i "s/NetReads/$NetReads/g" index.html
while read LINE
do
	if [[ "$NetWrites" == "" ]]
	then
		NetWrites=`echo ${LINE#-} |awk -F',' '{print "-"$1}'`
	else	
		NetWrites=$NetWrites","`echo ${LINE#-} |awk -F',' '{print "-"$1}'`
	fi
done < NETWRITE
NetWrites=(`echo $NetWrites|awk -F',' '{OFS=","}{NF=NF;$1="?";print}'|sed 's/?,//g'`)
sed -i "s/NetWrites/$NetWrites/g" index.html
#取disk指标
while read LINE
do
	if [[ "$DiskReads" == "" ]]
	then	
		DiskReads=`echo $LINE |awk -F',' '{print $2}'`
	else
		DiskReads=$DiskReads","`echo $LINE |awk -F',' '{print $2}'`
	fi
done < DISKREAD
DiskReads=(`echo $DiskReads|awk -F',' '{OFS=","}{NF=NF;$1="?";print}'|sed 's/?,//g'`)
sed -i "s/DiskReads/$DiskReads/g" index.html
while read LINE
do
	if [[ "$DiskWrites" == "" ]]
	then	
		DiskWrites=`echo ${LINE#-} |awk -F',' '{print "-"$2}'`
	else
		DiskWrites=$DiskWrites","`echo ${LINE#-} |awk -F',' '{print "-"$2}'`
	fi
done < DISKWRITE
DiskWrites=(`echo $DiskWrites|awk -F',' '{OFS=","}{NF=NF;$1="?";print}'|sed 's/?,//g'`)
sed -i "s/DiskWrites/$DiskWrites/g" index.html
for name in ${array[@]}  
do 
	rm -f $name
done  
rm -f NETREAD NETWRITE read write

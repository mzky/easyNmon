echo $(expr `cat version` + 1) >version
go build -ldflags "-X main.Version=0.`cat version` -X main.BuildTime=`date '+%Y-%m-%d_%H:%M:%S'` -w -s" main/easyNmon.go
./easyNmon


package common

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

func SysInfo() string {
	v, _ := mem.VirtualMemory()
	c, _ := cpu.Info()
	d, _ := disk.Usage("/")
	n, _ := host.Info()
	nv, _ := net.IOCounters(true)

	os := fmt.Sprintf("Hostname: %v OS: %v(%v) %v", n.Hostname, n.Platform, n.PlatformFamily, n.PlatformVersion)
	modelname := c[0].ModelName
	cpu := fmt.Sprintf("CPU: %v * %v cores", modelname, len(c))
	mem := fmt.Sprintf("Mem: %v GB Free: %v GB Used: %v MB", v.Total/1024/1000/1000, v.Available/1024/1000/1000, v.Used/1024/1000)
	net := fmt.Sprintf("Network: %v bytes / %v bytes", nv[0].BytesRecv, nv[0].BytesSent)
	disk := fmt.Sprintf("Disk: %v GB Free: %v GB", d.Total/1024/1024/1024, d.Free/1024/1024/1024)
	go func() {
		fmt.Println(os)
		fmt.Println(cpu)
		fmt.Println(mem)
		fmt.Println(net)
		fmt.Println(disk)
	}()
	return os + "<br/>" + cpu + "<br/>" + mem + "<br/>" + net + "<br/>" + disk
}

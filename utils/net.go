package utils

import (
	"fmt"
	"net"
)

// GetExternalIP 获取对外访问的ip
func GetExternalIP() string {
	ip := "127.0.0.1"
	netAddr, _ := net.InterfaceAddrs()
	for key, _ := range netAddr {
		networkIp, _ := netAddr[key].(*net.IPNet)
		if !networkIp.IP.IsLoopback() && networkIp.IP.To4() != nil {
			command := fmt.Sprintf("netstat -tunp|grep ESTABLISHED|grep ssh |grep %s", networkIp.IP.String())
			if addr, _ := Exec(command); addr != "" {
				return networkIp.IP.String()
			}
			command2 := fmt.Sprintf("netstat -tunp|grep ESTABLISHED|grep %s", networkIp.IP.String())
			if addr, _ := Exec(command2); addr != "" {
				ip = networkIp.IP.String()
			}
		}
	}
	return ip
}

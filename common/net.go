package common

import (
	"fmt"
	"net"
	"os/exec"
	"strings"
)

// GetExternalIP 获取对外访问的ip
func (f *Flag) GetExternalIP() {
	f.IP = "127.0.0.1"
	netAddr, _ := net.InterfaceAddrs()
	for key, _ := range netAddr {
		networkIp, _ := netAddr[key].(*net.IPNet)
		if !networkIp.IP.IsLoopback() && networkIp.IP.To4() != nil {
			command := exec.Command("bash", "-c", fmt.Sprintf("netstat -tunp | grep ESTABLISHED | grep %s", networkIp.IP))
			if output, err := command.Output(); output != nil && err == nil {
				if strings.Contains(string(output), networkIp.IP.String()) {
					f.IP = networkIp.IP.String()
				}
			}
		}
	}
}

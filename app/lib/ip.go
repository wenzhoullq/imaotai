package lib

import "net"

func CheckIP(ip string) bool {
	netIP := net.ParseIP(ip)
	ip4 := netIP.To4()
	if ip4 == nil {
		return false
	}
	return true
}

func HasLocalIP(ip string) bool {
	netIP := net.ParseIP(ip)
	if netIP.IsLoopback() {
		return true
	}

	ip4 := netIP.To4()
	return ip4[0] == 10 || // 10.0.0.0/8
		(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) || // 172.16.0.0/12
		(ip4[0] == 169 && ip4[1] == 254) || // 169.254.0.0/16
		(ip4[0] == 192 && ip4[1] == 168) // 192.168.0.0/16

}

package tcp

import (
	"fmt"
	"net"
)

func Check(ip net.IP, port int) error {
	addr := fmt.Sprintf("%s:%d", ip.String(), port)
	fmt.Println(addr)
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)

	if err != nil {
		return err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil && conn != nil {
		conn.Close()
	}

	return err
}

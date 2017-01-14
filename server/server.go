package server

import (
	"fmt"
	"net"

	"github.com/veskoy/gomas/cmd"
	"github.com/veskoy/gomas/utilities"
)

// Listen starts the Master Server.
func Listen() {
	ServerAddr, err := net.ResolveUDPAddr("udp", cmd.ServerIP+":"+cmd.ServerPort)
	utilities.ExitOnError(&err)

	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	utilities.ExitOnError(&err)

	defer ServerConn.Close()

	buffer := make([]byte, 1024)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buffer)
		fmt.Println("Received ", string(buffer[0:n]), " from ", addr)

		utilities.CheckError(&err)
	}
}

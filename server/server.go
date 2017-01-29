package server

import (
	"fmt"
	"net"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"github.com/veskoy/gomas/cmd"
	"github.com/veskoy/gomas/db"
	"github.com/veskoy/gomas/msqp"
	"github.com/veskoy/gomas/utilities"
)

// Listen starts the Master Server.
func Listen() {
	serverAddr, err := net.ResolveUDPAddr("udp", cmd.ServerIP+":"+cmd.ServerPort)
	utilities.PanicOnError(&err)

	serverConn, err := net.ListenUDP("udp", serverAddr)
	utilities.PanicOnError(&err)

	defer serverConn.Close()

	dbConn, err := db.Open()
	utilities.PanicOnError(&err)

	defer dbConn.Close()

	loop(serverConn, dbConn)
}

func loop(serverConn *net.UDPConn, dbConn *gorm.DB) {
	buffer := make([]byte, 1024)

	for {
		n, addr, err := serverConn.ReadFromUDP(buffer)
		utilities.CheckError(&err)

		if err == nil {
			go handleQuery(buffer[0:n], serverConn, addr, dbConn)
		}
	}
}

func handleQuery(query []byte, serverConn *net.UDPConn, addr *net.UDPAddr, dbConn *gorm.DB) {
	parsedQuery := msqp.ParseC2MQuery(query)
	queryType := msqp.GetQueryType(parsedQuery)

	switch queryType {
	case msqp.FirstRequestServerList:
		serverConn.WriteToUDP(msqp.ServerListQueryReply(parsedQuery, dbConn), addr)

	case msqp.ContinueRequestServerList:
		serverConn.WriteToUDP(msqp.ServerListQueryReply(parsedQuery, dbConn), addr)

	default: // msqp.Undefined
		if viper.Get("debug") == "true" {
			fmt.Println("Invalid request from IP: ", addr)
		}
	}
}

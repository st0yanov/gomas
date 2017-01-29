package msqp

import (
	"encoding/binary"
	"github.com/jinzhu/gorm"
	"github.com/veskoy/gomas/db"
	"net"
	"strconv"
	"strings"
)

// RegionCode represents the geographic area the game server is registered in
type RegionCode int

const (
	useast RegionCode = iota
	uswest
	ussouth
	europe
	asia
	australia
	middleeast
	africa
	rest = 255
)

func findSzEnd(query []byte, start int) int {
	for i := start; i < len(query); i++ {
		if query[i] == 0 {
			return i
		}
	}
	return len(query)
}

// C2MFormat represents the client to master query format
type C2MFormat struct {
	MessageType byte
	RegionCode  byte
	IPPort      string
	Filter      string
}

// ParseC2MQuery parses an incoming query and returns the parsed data in a C2MFormat struct
func ParseC2MQuery(query []byte) C2MFormat {
	messageType := byte(0x31)
	regionCode := byte(0xFF)
	ipPort := "0.0.0.0:0"
	filter := ""

	queryLength := len(query)

	if queryLength >= 1 {
		messageType = query[0]
	}

	if queryLength >= 2 {
		regionCode = query[1]
	}

	if queryLength >= 3 {
		ipEndIndex := findSzEnd(query, 2)
		ipPort = string(query[2:ipEndIndex])

		if queryLength > ipEndIndex+1 {
			filter = string(query[(ipEndIndex + 1):])
		}
	}

	result := C2MFormat{MessageType: messageType, RegionCode: regionCode, IPPort: ipPort, Filter: filter}
	return result
}

// QueryType is used to identify what kind of query has the server received
type QueryType int

const (
	// Undefined indicates that the incoming query is not supported by the MSQP
	Undefined QueryType = iota

	// FirstRequestServerList indicates that the incoming query is for
	// server list retrieval from the beginning.
	FirstRequestServerList

	// ContinueRequestServerList indicates that the incoming query is for
	// continued server list retrieval from a specified server ip:port.
	ContinueRequestServerList
)

// GetQueryType returns the type of a received query
func GetQueryType(parsedQuery C2MFormat) QueryType {
	if string(parsedQuery.MessageType) != "1" {
		return Undefined
	}

	ipPortParts := strings.Split(parsedQuery.IPPort, ":")
	if len(ipPortParts) < 2 {
		return Undefined
	}

	ip := ipPortParts[0]
	port, _ := strconv.Atoi(ipPortParts[1])

	if ip == "0.0.0.0" && port == 0 {
		return FirstRequestServerList
	}

	parsedIP := net.ParseIP(ip)
	if parsedIP.To4() != nil && port >= 0 && port <= 65535 {
		return ContinueRequestServerList
	}

	return Undefined
}

// M2CFormat represents the master to client query format
type M2CFormat struct {
	Oct1 byte
	Oct2 byte
	Oct3 byte
	Oct4 byte
	Port uint16
}

// Bytes returns the byte sequence of the M2CFormat
func (q M2CFormat) Bytes() []byte {
	var result []byte

	ip := []byte{q.Oct1, q.Oct2, q.Oct3, q.Oct4}
	port := make([]byte, 2)

	binary.BigEndian.PutUint16(port, q.Port)

	result = append(result, ip...)
	result = append(result, port...)

	return result
}

// M2CHeader is a constant byte sequence that each reply starts with
var M2CHeader = []byte{255, 255, 255, 255, 102, 10}

// M2CEnd is a constant byte sequence that indicates the end of the server list
var M2CEnd = []byte{0, 0, 0, 0, 0, 0}

// ServerListQueryReply returns a sequence of bytes containing a list of servers
func ServerListQueryReply(parsedQuery C2MFormat, dbConn *gorm.DB) []byte {
	gameServers := db.GetFilteredGameServers(dbConn, parsedQuery.Filter)
	servers := fmtM2C(gameServers)

	var response []byte
	response = append(response, M2CHeader...)

	for _, server := range servers {
		serverIP := server.Bytes()
		response = append(response, serverIP...)
	}

	response = append(response, M2CEnd...)

	return response
}

func fmtM2C(servers []db.BasicServer) []M2CFormat {
	serversM2C := []M2CFormat{}

	for _, server := range servers {
		parsedServerIP := strings.Split(server.IP, ".")

		oct1, _ := strconv.Atoi(parsedServerIP[0])
		oct2, _ := strconv.Atoi(parsedServerIP[1])
		oct3, _ := strconv.Atoi(parsedServerIP[2])
		oct4, _ := strconv.Atoi(parsedServerIP[3])
		port, _ := strconv.Atoi(server.Port)

		serverM2C := M2CFormat{
			Oct1: byte(oct1),
			Oct2: byte(oct2),
			Oct3: byte(oct3),
			Oct4: byte(oct4),
			Port: uint16(port),
		}

		serversM2C = append(serversM2C, serverM2C)
	}

	return serversM2C
}

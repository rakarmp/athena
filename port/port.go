package port

import (
	"net"
	"strconv"
	"time"

	"github.com/rakarmp/athena/handler"
)

func ScanPort(protocol, hostname string, port int) handler.ScanResult {
	result := handler.ScanResult{Port: port}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		result.State = "Closed"
		return result
	}

	defer conn.Close()
	result.State = "Open"
	return result
}

func InitialScan(hostname string) []handler.ScanResult {
	var result []handler.ScanResult

	for i := 0; i <= 1024; i++ {
		result = append(result, ScanPort("udp", hostname, i))
	}

	for i := 0; i <= 1024; i++ {
		result = append(result, ScanPort("tcp", hostname, i))
	}
	return result
}

func WideScan(hostname string) []handler.ScanResult {
	var results []handler.ScanResult

	for i := 0; i <= 49152; i++ {
		results = append(results, ScanPort("udp", hostname, i))
	}

	for i := 0; i <= 49152; i++ {
		results = append(results, ScanPort("tcp", hostname, i))
	}

	return results

}

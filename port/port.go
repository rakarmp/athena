package port

import (
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/rakarmp/athena/handler"
)

func ScanPort(protocol, hostname string, port int, wg *sync.WaitGroup, mutex *sync.Mutex, results *[]handler.ScanResult) {
	defer wg.Done()

	result := handler.ScanResult{Port: port}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		result.State = "Closed"
	} else {
		defer conn.Close()
		result.State = "Open"
	}

	mutex.Lock()
	*results = append(*results, result)
	mutex.Unlock()
}

func InitialScan(hostname string) []handler.ScanResult {
	var results []handler.ScanResult
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i <= 1024; i++ {
		wg.Add(1)
		go ScanPort("udp", hostname, i, &wg, &mutex, &results)
	}

	for i := 0; i <= 1024; i++ {
		wg.Add(1)
		go ScanPort("tcp", hostname, i, &wg, &mutex, &results)
	}

	wg.Wait()
	return results
}

func WideScan(hostname string) []handler.ScanResult {
	var results []handler.ScanResult
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i <= 49152; i++ {
		wg.Add(1)
		go ScanPort("udp", hostname, i, &wg, &mutex, &results)
	}

	for i := 0; i <= 49152; i++ {
		wg.Add(1)
		go ScanPort("tcp", hostname, i, &wg, &mutex, &results)
	}

	wg.Wait()
	return results
}

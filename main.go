package main

import (
	"fmt"

	"github.com/rakarmp/athena/port"
)

func main() {
	fmt.Println("Port Scanning")
	result := port.InitialScan("localhost")
	fmt.Println(result)

	widescanresult := port.WideScan("localhost")
	fmt.Println(widescanresult)
}

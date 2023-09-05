## Port Scanning

<samp>Simple Write Pure Go</samp></br>
<samp>For Learn Create Project</samp>

### Use

```
go get github.com/rakarmp/athena/port
```

#### Example Code

```
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
```

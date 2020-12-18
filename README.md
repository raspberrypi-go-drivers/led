# LED

[![PkgGoDev](https://pkg.go.dev/badge/github.com/bbayszczak/raspberrypi-go-drivers/led)](https://pkg.go.dev/github.com/bbayszczak/raspberrypi-go-drivers/led)

This drivers allows interact with a LED connected to a GPIO pin

## Documentation

For full documentation, please visit [![PkgGoDev](https://pkg.go.dev/badge/github.com/bbayszczak/raspberrypi-go-drivers/led)](https://pkg.go.dev/github.com/bbayszczak/raspberrypi-go-drivers/led)

## Quick start

```go
import (
	"fmt"
	"time"

	"github.com/bbayszczak/raspberrypi-go-drivers/led"
	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()

	l1 := led.NewLED(18)
	for {
		l1.Toggle()
		time.Sleep(1 * time.Second)
		state, _ := l1.GetState()
		fmt.Println(state)
	}
}
```

# LED

[![PkgGoDev](https://pkg.go.dev/badge/github.com/raspberrypi-go-drivers/led)](https://pkg.go.dev/github.com/bbayszczak/raspberrypi-go-drivers/led)
![golangci-lint](https://github.com/raspberrypi-go-drivers/led/workflows/golangci-lint/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/raspberrypi-go-drivers/led)](https://goreportcard.com/report/github.com/raspberrypi-go-drivers/led)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

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

## Raspberry Pi compatibility

This driver has has only been tested on an Raspberry Pi Zero WH using integrated bluetooth but may work well on other Raspberry Pi having integrated Bluetooth

## License

MIT License

---

Special thanks to @stianeikeland

This driver is based on his work in [stianeikeland/go-rpio](https://github.com/stianeikeland/go-rpio/)

package led_test

import (
	"fmt"
	"os"
	"time"

	"github.com/raspberrypi-go-drivers/led"
	log "github.com/sirupsen/logrus"
	"github.com/stianeikeland/go-rpio/v4"
)

// The LED will dim from 0% to 100% to 0% continuously
func Example_dimmable() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()

	l1 := led.NewLED(18)
	l1.Dimmable()
	for {
		for i := 0; i <= 100; i++ {
			if err := l1.SetBrightness(i); err != nil {
				log.WithField("error", err).Error("impossible to change LED brightness")
			}
			time.Sleep(time.Millisecond * 10)
		}
		for i := 100; i >= 0; i-- {
			if err := l1.SetBrightness(i); err != nil {
				log.WithField("error", err).Error("impossible to change LED brightness")
			}
			time.Sleep(time.Millisecond * 10)
		}
	}
}

// The LED will blink continuously and the current LED status
// will be displayed as output
func Example_toggle() {
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

func ExampleNewLED() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()

	// 18 is one of the PWM compatible GPIO pin on the Raspberry Pi
	l1 := led.NewLED(18)
	l1.On()
}

func ExampleLED_Toggle() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()

	l1 := led.NewLED(18)
	// Switch on the led if off or swicth on if off
	l1.Toggle()
}

func ExampleLED_On() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()

	l1 := led.NewLED(18)
	l1.On()
}

func ExampleLED_Off() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()

	l1 := led.NewLED(18)
	l1.Off()
}

func ExampleLED_GetState() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()

	l1 := led.NewLED(18)
	l1.On()
	state, _ := l1.GetState()
	fmt.Println(state)
	// Output: true
}

func ExampleLED_Dimmable() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()

	l1 := led.NewLED(18)
	l1.Dimmable()
	// LED is now dimmable (PWM mode is enabled)
}

func ExampleLED_NonDimmable() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()

	l1 := led.NewLED(18)
	l1.Dimmable()
	// Here you can use SetBrightness
	l1.NonDimmable()
	// LED is not dimmable anymore (PWM mode is disabled)
}

func ExampleLED_SetBrightness() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()

	l1 := led.NewLED(18)
	l1.Dimmable()
	// set LED brightness to 50% of its maximum value
	if err := l1.SetBrightness(50); err != nil {
		log.WithField("error", err).Error("impossible to change LED brightness")
	}
}

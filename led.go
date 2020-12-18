package led

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

const (
	softwarePWMFrequency int64 = 60
	timeUnit             int64 = int64(time.Second) / softwarePWMFrequency
)

// LED instance
type LED struct {
	pinID              uint8
	pin                rpio.Pin
	dimmable           bool
	brightness         int
	softwarePWMTimeOn  int64
	softwarePWMTimeOff int64
}

// NewLED creates a new LED instance
func NewLED(pinID uint8) *LED {
	led := LED{
		pinID:              pinID,
		brightness:         0,
		dimmable:           false,
		softwarePWMTimeOn:  0,
		softwarePWMTimeOff: timeUnit,
	}
	led.pin = rpio.Pin(led.pinID)
	led.pin.Mode(rpio.Output)
	return &led
}

// Toggle change the state of the LED
func (led *LED) Toggle() {
	led.pin.Toggle()
}

// On switch on the LED
func (led *LED) On() {
	led.pin.High()
}

// Off switch off the LED
func (led *LED) Off() {
	led.pin.Low()
}

// GetState return the current state of the lED
//
// returns: true for on, off for false
func (led *LED) GetState() (bool, error) {
	state := led.pin.Read()
	if state == 1 {
		return true, nil
	} else if state == 0 {
		return false, nil
	}
	return false, fmt.Errorf("unknown state '%d'", state)
}

// Dimmable make the LED dimmable
func (led *LED) Dimmable() {
	led.dimmable = true
	led.brightness = 0
	if isNativePWMPin(led.pin) {
		led.pin.Mode(rpio.Pwm)
		led.pin.Freq(10000)
	} else {
		go led.softwareBrightness()
	}
}

// NonDimmable make the dimmable LED non dimmable
func (led *LED) NonDimmable() {
	led.dimmable = false
	if isNativePWMPin(led.pin) {
		led.pin.Mode(rpio.Output)
	}
	led.brightness = -1
}

// SetBrightness set the brightness of the LED
//
//   If the LED is connected to a hardware PWM pin, SetBrightness will use hardware PWM
//   If the LED is connected to a non PWM pin, SetBrightness will use a software emulated PWM
func (led *LED) SetBrightness(percentage int) error {
	if !led.dimmable {
		return fmt.Errorf("LED is not setup as dimmable")
	}
	if percentage < 0 || percentage > 100 {
		return fmt.Errorf("percentage value must be >= 0 <= 100")
	}
	led.brightness = percentage
	if isNativePWMPin(led.pin) {
		led.pin.DutyCycle(uint32(percentage), 100)
	} else {
		led.softwarePWMTimeOn = int64(timeUnit) / 100 * int64(led.brightness)
		led.softwarePWMTimeOff = int64(timeUnit) / 100 * int64(100-led.brightness)
	}
	return nil
}

func (led *LED) softwareBrightness() {
	for led.brightness >= 0 {
		if led.brightness > 0 {
			led.On()
			time.Sleep(time.Duration(led.softwarePWMTimeOn))
		}
		if led.brightness < 100 {
			led.Off()
			time.Sleep(time.Duration(led.softwarePWMTimeOff))
		}
	}
}

func isNativePWMPin(pin rpio.Pin) bool {
	nativePWMPins := []rpio.Pin{12, 13, 18}
	for _, p := range nativePWMPins {
		if pin == p {
			return true
		}
	}
	return false
}

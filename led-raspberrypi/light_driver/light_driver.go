package lightdriver

import (
	"fmt"
	"os"

	"github.com/stianeikeland/go-rpio"
)

func TurnON(pinNumber int64) {
	pin := rpio.Pin(pinNumber)
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	// Set pin to output mode
	pin.Output()

	pin.High()
}

func TurnOff(pinNumber int64) {
	pin := rpio.Pin(pinNumber)
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	// Set pin to output mode
	pin.Output()
	pin.Low()

}

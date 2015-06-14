// example
package main

import (
	"fmt"
	"github.com/MarinX/goble"
)

func main() {

	//Locate your dev, on Windows is COM4 or equivalent
	hm, err := goble.New("/dev/ttyUSB2")
	if err != nil {
		fmt.Println(err)
		return
	}

	//Get version
	fmt.Println(
		"Software Version:",
		hm.SofwareVersion().Result,
	)

	//Last connected device
	fmt.Println(
		"Last connected device",
		hm.GetLastConnectedDeviceAddress().Param,
	)

	// Get Device name
	fmt.Println(
		"Device Name:",
		hm.GetDeviceName().Param,
	)

	//Set device name
	fmt.Println(
		"Setting device name...",
		hm.SetDeviceName("HelloWorld").Result,
	)

	//Needs to reset to changes take effect
	fmt.Println(
		"Reseting...",
		hm.Reset().Result,
	)

	//Now device name will be HelloWorld
	//Diff between result / param
	//Result is raw response from device and param is converted result
	devName := hm.GetDeviceName()
	fmt.Println(
		"Device Name with result:",
		devName.Result,
		"\nDevice Name with param:",
		devName.Param,
	)

	//Get mode by definition of BleMode
	fmt.Println(
		"Mode:",
		hm.GetMode().Param,
	)

	//Get bound mode by definition of BleBoundMode
	fmt.Println(
		"Bound mode:",
		hm.GetBoundMode().Param,
	)

	//Get PIO Pin status
	fmt.Println(
		"PIO 2 status(0-LOW, 1-HIGH):",
		hm.GetPIO(goble.PIN_2).Param,
	)

	//Set PIO Pin to HIGH
	fmt.Println(
		"Setting PIO Pin 2 to HIGH (0-LOW, 1-HIGH)",
		hm.SetPIO(goble.PIN_2, goble.HIGH).Result,
	)

	//Now check the status
	fmt.Println(
		"PIO 2 status (0-LOW, 1-HIGH):",
		hm.GetPIO(goble.PIN_2).Param,
	)

	//Close
	hm.Close()
}

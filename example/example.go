// example
package main

import (
	"fmt"
	"github.com/MarinX/goble"
)

func main() {

	//Locate your dev
	hm, err := goble.New("/dev/ttyUSB1")
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

	//Close
	hm.Close()
}

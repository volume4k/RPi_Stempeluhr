package main

import (
	"fmt"
	"github.com/fuzxxl/nfc/2.0/nfc"
)

func main(){
	devices,_ := nfc.ListDevices()
	fmt.Println(devices)

	device := nfc.Device{}
	fmt.Println(device.Connection())
}

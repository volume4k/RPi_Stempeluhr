package main

import (
	"fmt"
	"github.com/fuzxxl/nfc/2.0/nfc"
)

func main(){
	devices,_ := nfc.ListDevices()
	fmt.Println(devices)

	device := nfc.Device{}
	dev2, _ := nfc.Open(device.Connection())
	fmt.Println(dev2)
}

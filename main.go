package main

import (
	"fmt"
	"github.com/fuzxxl/nfc/2.0/nfc"
)

func main(){
	devices,_ := nfc.ListDevices()
	dev := nfc.Device{}
	fmt.Println(devices)
	_ = dev.InitiatorInit()
	fmt.Println(nfc.Version())
}

package main

import (
	"fmt"
	"github.com/fuzxxl/nfc/2.0/nfc"
)

func main(){
	devices,_ := nfc.ListDevices()
	dev := nfc.Device{}
	nfc.Open(devices[0])
	fmt.Println("bis hier")
	fmt.Println(devices)
	fmt.Print(dev)
}

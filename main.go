package main

import (
	"fmt"
	"github.com/fuzxxl/nfc/2.0/nfc"
)

func main(){
	devices,_ := nfc.ListDevices()
	dev := nfc.Device{}
	pnd, err := nfc.Open("")
	if err != nil {
		fmt.Println("Could not open device.")
	}
	fmt.Println("bis hier")
	fmt.Println(devices)
	fmt.Println(dev)
	fmt.Println(pnd)
}

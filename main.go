package main

import (
	"fmt"
	"github.com/fuzxxl/nfc/2.0/nfc"
)

var m = nfc.Modulation{Type: nfc.ISO14443a}

func main(){
	m := nfc.Modulation{Type: nfc.ISO14443a}
	devices,_ := nfc.ListDevices()
	dev := nfc.Device{}
	pnd, err := nfc.Open("")
	if err != nil {
		fmt.Println("Could not open device.")
	}
	fmt.Println("bis hier")
	fmt.Println(devices)

}

func get_card (pnd *nfc.Device) ([10]byte, error) {
	for {
		targets, err := pnd.InitiatorListPassiveTargets(m)
		if err != nil {
			return [10]byte{}, fmt.Errorf("listing available nfc targets", err)
		}

		for _, t := range targets {
			if card, ok := t.(*nfc.ISO14443aTarget); ok {
				return card.UID, nil
			}
		}
	}
}

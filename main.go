package main

import (
	"fmt"
	"github.com/fuzxxl/nfc/2.0/nfc"
	"log"
)

var m = nfc.Modulation{Type: nfc.ISO14443a}
var devstr = ""

func main(){
	//m := nfc.Modulation{Type: nfc.ISO14443a}
	//devices,_ := nfc.ListDevices()
	//dev := nfc.Device{}
	//pnd, err := nfc.Open("")
	//if err != nil {
	//	fmt.Println("Could not open device.")
	//}
	//fmt.Println("bis hier")
	//fmt.Println(devices)

	fmt.Println("using libnfc", nfc.Version())

	pnd, err := nfc.Open(devstr)
	if err != nil {
		log.Fatalf("could not open device: %v", err)
	}
	defer pnd.Close()

	if err := pnd.InitiatorInit(); err != nil {
		log.Fatalf("could not init initiator: %v", err)
	}

	fmt.Println("opened device", pnd, pnd.Connection())

	card_id, err := get_card(&pnd)
	if err != nil {
		fmt.Errorf("failed to get_card", err)
	}

	if card_id != [10]byte{} {
		fmt.Printf("card found %#X\n", card_id)
	} else {
		fmt.Printf("no card found\n")
	}
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

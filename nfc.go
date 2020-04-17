package main

import (
	"fmt"
	"github.com/fuzxxl/nfc/2.0/nfc"
	"log"
)

var m = nfc.Modulation{Type: nfc.ISO14443a}

func nfcInit() {

	fmt.Println("using libnfc", nfc.Version())

	pnd, err := nfc.Open("")
	if err != nil {
		log.Fatalf("could not open device: %v", err)
	}
	defer pnd.Close()

	if err := pnd.InitiatorInit(); err != nil {
		log.Fatalf("could not init initiator: %v", err)
	}

	fmt.Println("opened device", pnd, pnd.Connection())

	cardId, err := getCard(&pnd)
	if err != nil {
		fmt.Errorf("failed to getCard", err)
	}

	if cardId != [10]byte{} {
		fmt.Printf("card found %#X\n", cardId)
		c <- cardId
	} else {
		fmt.Printf("no card found\n")
		c <- nil
	}
}

func getCard(pnd *nfc.Device) ([10]byte, error) {
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
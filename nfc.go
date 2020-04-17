package main

import (
	"fmt"
	"periph.io/x/periph/experimental/devices/mfrc522"
	"time"
)


func readTag(){
	fmt.Println("hi")
	var auth byte = 0
	key := mfrc522.DefaultKey
	dev := mfrc522.Dev{}
	uid, err := dev.ReadCard(time.Second*15, auth, 0, 0, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(uid)
}

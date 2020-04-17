package main

import "fmt"

var c chan [10]byte

func main(){
	fmt.Println("starting up.")
	controlCircle()
}

func controlCircle(){
	// TODO: breakout document

	go nfcInit()
	handOffForDB(<- c)
}
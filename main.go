package main

import "fmt"

var c = make(chan [10]byte)

func main(){
	fmt.Println("starting up.")
	controlCircle()
}

func controlCircle(){
	// TODO: breakout document



	for  {
		go nfcInit()
		handOffForDB(<- c)
	}

}
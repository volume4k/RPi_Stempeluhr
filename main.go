package main

import "fmt"

func main(){
	fmt.Println("starting up.")
	controlCircle()
}

func controlCircle(){
	// TODO: breakout document

	c := make(chan [10]byte)

	for  {
		go nfcInit()
		handOffForDB(<- c)
	}

}
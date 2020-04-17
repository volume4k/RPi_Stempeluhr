package main

import "fmt"

func main(){
	fmt.Println(formatDSN(loadConfig()))
	readTag()
}

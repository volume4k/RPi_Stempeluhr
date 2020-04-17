package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Configuration struct {
	dbUser    string
	dbPass   string
	dbAddr string
	dbProtocol string
	dbName string
}

func loadConfig ()  {
	file, _ := ioutil.ReadFile("conf.json")
	configuration := Configuration{dbName: "db name", dbPass: "nothing kommt hier vorbei"}
	_ = json.Unmarshal(file, &configuration)
	fmt.Println(configuration)
}
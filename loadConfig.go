package main

// note: unfortunately the unmarshal/decode function doesn't work -> returns empty struct / doesn't change anything
// TODO: importing and implementing JSON config <- until then hardcode it.

type Configuration struct {
	dbUser    string
	dbPass   string
	dbAddr string
	dbProtocol string
	dbName string
}

func loadConfig () Configuration {
	configuration := Configuration{dbUser: "stempeluhr", dbPass: "Rx8723hm95Wqbnk324zx", dbAddr: "127.0.0.1:3306", dbName: "test", dbProtocol: "tcp"}
	return configuration

	// FOR TODO
	//file, _ := ioutil.ReadFile("conf.json")
	//_ = json.Unmarshal(file, &configuration)
	//fmt.Println(configuration)
}
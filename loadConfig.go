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
	configuration := Configuration{dbUser: "jonathan", dbPass: "pass", dbAddr: "hostname:port", dbName: "dbName", dbProtocol: "tcp"}
	return configuration

	// FOR TODO
	//file, _ := ioutil.ReadFile("conf.json")
	//_ = json.Unmarshal(file, &configuration)
	//fmt.Println(configuration)
}
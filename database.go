package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

func formatDSN(a Configuration) string {
	config := mysql.Config{
		User:   a.dbUser,
		Passwd: a.dbPass,
		Net:    a.dbProtocol,
		Addr:   a.dbAddr,
		DBName: a.dbName,
	}

	myDsn := config.FormatDSN()

	return myDsn
}

func none(){
	db, err := sql.Open("mysql",formatDSN(loadConfig()))
	if err != nil {
		panic(err)
	}
	fmt.Println(db)
}

func handOffForDB(u [10]byte){
// instantly hands off operation to Handler
	go DBHandler(u)
}

func DBHandler(u [10]byte){
	none()
}

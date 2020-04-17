package main

config := mysql.Config{
User:   "user",
Passwd: "password",
Net:    "tcp",
Addr:   "hostname:Port",
DBName: "dbname",
}

myDsn := config.FormatDSN()
fmt.Println(myDsn)
db, err := sql.Open("mysql",myDsn)
if err != nil {
panic(err)
}
fmt.Println(db)

package main

import "fmt"

type DataBase interface {
	writeData(s string) string
}

type MySQLDB struct {
	DBURL string
}

func (DB MySQLDB) writeData(s string) string {
	return "Data being written to MySQL DB: " + s
}

type PostgresDB struct {
	DBURL string
}

func (DB PostgresDB) writeData(s string) string {
	return "Data being written to Postgres DB: " + s
}

func storeUserData(name string, DB DataBase) string {
	return DB.writeData(name)
}

func main() {
	DB := PostgresDB{DBURL: "www.Postgres.com"}
	result := storeUserData("Shervin", DB)
	fmt.Println(result)
}

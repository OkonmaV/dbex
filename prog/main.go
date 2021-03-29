package main

import (
	"dbex"
	"log"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=tests port=9920 sslmode=disable TimeZone=Asia/Yekaterinburg"
	// dsn := "root:root@tcp(localhost:3306)/ex?charset=utf8mb4&parseTime=True&loc=Local"
	log.Println("connecting to server...")
	conn, err := dbex.NewMySqlConnection(dsn)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("connected to server")
	log.Println("select for avg test time")
	res, err := conn.GetAvgTestElapsed()
	if err != nil {
		log.Panicln(err)
	}
	log.Println("avg test time is", res)

	log.Println("select for avg test time")
	withouttests, err := conn.SelectTestersWithoutTests()
	if err != nil {
		log.Panicln(err)
	}
	log.Println("avg test time is", withouttests)

	log.Println("select for avg test time")
	tests, err := conn.SelectTestsetsByTester(1)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("avg test time is", tests)

	log.Println("Completed!")

}

package main

import (
	"dbex"
	"log"
	"time"
)

func main() {
	dsn := "root:root@tcp(localhost:3306)/ex?charset=utf8mb4&parseTime=True&loc=Local"
	log.Println("connecting to server...")
	conn, err := dbex.NewMySqlConnection(dsn)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("connected to server")
	log.Println("----------")
	log.Println("select for avg test time")
	res, err := conn.GetAvgTestElapsed() // среднее время
	if err != nil {
		log.Panicln(err)
	}
	log.Println("avg test time is ", res*time.Second)
	log.Println("----------")

	log.Println("select for testers with no tests")
	withouttests, err := conn.SelectTestersWithoutTests() // тестеры без назначенных на них тестов
	if err != nil {
		log.Panicln(err)
	}
	log.Println("testers without tests: ", withouttests)
	log.Println("----------")

	var testerid uint = 8

	log.Println("select for testset by tester ", testerid)
	testsets, err := conn.SelectTestsetsByTester(testerid) // тестсеты по айди тестера
	if err != nil {
		log.Panicln(err)
	}
	log.Println("testsets for tester", testerid, " : ", testsets)
	log.Println("----------")

	log.Println("Completed!") // при запуске тестов по описанным структурам строится отдельная схема tests, аналогичная основной
	//							тесты сами подчищают за собой тестовую таблицу, после выполнения

}

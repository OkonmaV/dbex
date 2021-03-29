package main

import (
	"dbex"
	"log"
)

func main() {

	dsn := "root:root@tcp(localhost:3306)/ex?charset=utf8mb4&parseTime=True&loc=Local"
	log.Println("connecting to server...")
	conn, err := dbex.NewMySqlConnection(dsn)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("connected to server")
	log.Println("run create testset")
	//db.Save(&Status{ID: 4, Status: "bar"})
	//fmt.Println(conn.SelectTestpointsAll())
	log.Println(conn.CreateTestset(&dbex.Testset{Name: "test1", TestPlanId: 1}))
	// d := []Status{}
	// _ = db.First(&d, 5)
	// fmt.Println(d)

	//s1 := Status{ID: 2, Status: "Ошибка"}
	log.Println("Completed!")

}

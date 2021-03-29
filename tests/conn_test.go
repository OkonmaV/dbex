package dbex_test

import (
	"dbex"
	"log"
	"os"
)

var DB *dbex.MySqlConnection

func init() {
	var err error
	if DB, err = OpenTestConnection(); err != nil {
		log.Printf("failed to connect database, got error %v", err)
		os.Exit(1)
	} else {
		sqlDB, err := DB.DB.DB()
		if err == nil {
			err = sqlDB.Ping()
		}

		if err != nil {
			log.Printf("failed to connect database, got error %v", err)
		}

		RunMigrations()
	}
}

func OpenTestConnection() (*dbex.MySqlConnection, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=tests port=9920 sslmode=disable TimeZone=Asia/Yekaterinburg"
	conn, err := dbex.NewMySqlConnection(dsn)
	if err != nil {
		return nil, err
	}
	return conn, nil
	// return dbex.NewMySqlConnection("root:root@tcp(localhost:3306)/ex?charset=utf8mb4&parseTime=True&loc=Local")
}

func RunMigrations() {
	var err error
	allModels := []interface{}{&dbex.Department{}, &dbex.Position{}, &dbex.Program{}, &dbex.Status{}, &dbex.Testcase{}, &dbex.Tester{}, &dbex.Testplan{}, &dbex.Testpoint{}, &dbex.Testset{}, &dbex.Testresult{}}

	if err = DB.DB.Migrator().DropTable(allModels...); err != nil {
		log.Printf("Failed to drop table, got error %v\n", err)
		os.Exit(1)
	}

	if err := DB.DB.Migrator().AutoMigrate(allModels...); err != nil {
		log.Printf("Failed to auto migrate, but got error %v\n", err)
		os.Exit(1)
	}

	for _, m := range allModels {
		if !DB.DB.Migrator().HasTable(m) {
			log.Printf("Failed to create table for %#v\n", m)
			os.Exit(1)
		}
	}
}

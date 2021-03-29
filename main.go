package dbex

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Conf struct {
	Db *gorm.DB
}

type Crud interface {
	Create(interface{}) error
	UpdateById(int, interface{}) error
	Delete() error
	Read() error
}

func main() {

	dsn := "root:root@tcp(127.0.0.1:3306)/ex?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	conf := &Conf{Db: db}
	//db.Save(&Status{ID: 4, Status: "bar"})
	//fmt.Println(conf.SelectTestpointsAll())
	fmt.Println(conf.CreateTestset(&Testset{Name: "test1", TestPlanId: 1}))
	// d := []Status{}
	// _ = db.First(&d, 5)
	// fmt.Println(d)

	//s1 := Status{ID: 2, Status: "Ошибка"}
	fmt.Println("Completed!")

}

func GetConf() (*Conf, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/ex?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Conf{Db: db}, nil
}

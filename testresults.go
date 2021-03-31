package dbex

import (
	"time"
)

// const format = "2006-01-02 15:04:05.00"

type Testresult struct {
	Id          uint      `gorm:"primaryKey"`
	TestPointId uint      `gorm:"column:TestPointId"`
	Testpoint   Testpoint `gorm:"foreignKey:TestPointId"`
	StatusId    uint      `gorm:"column:StatusId"`
	Status      Status    `gorm:"foreignKey:StatusId"`
	Start       time.Time `gorm:"not null"`
	Finish      time.Time `gorm:"not null"`
}

func (conn *MySqlConnection) CreateTestresult(data *Testresult) error {
	return conn.DB.Create(data).Error
}

func (conn *MySqlConnection) DeleteTestresultById(id uint) error {
	return conn.DB.Delete(&Testresult{Id: id}).Error
}

func (conn *MySqlConnection) UpdateTestresult(data *Testresult) error {
	return conn.DB.Save(data).Error
}

func (conn *MySqlConnection) SelectAllTestresults() ([]Testresult, error) {
	var res []Testresult
	err := conn.DB.Find(&res).Error
	return res, err
}

func (conn *MySqlConnection) SelectTestresultById(id uint) (Testresult, error) {
	var res Testresult
	err := conn.DB.First(&res, id).Error
	return res, err
}

func (conn *MySqlConnection) GetAvgTestElapsed() (time.Duration, error) {
	var result float64

	//smt:= conn.DB.
	err := conn.DB.Table("testresults").Select("avg(time_to_sec(timediff(finish,start)))").Scan(&result).Error //Scan(&result).Error
	if err != nil {
		return 0, err
	}
	return time.Duration(result), nil
}

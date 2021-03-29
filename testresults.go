package dbex

import "time"

// const format = "2006-01-02 15:04:05.00"

type Testresult struct {
	Id          uint `gorm:"primaryKey"`
	TestPointId uint `gorm:"foreignKey:FK_TestResults_TestPoints_TestPointId"`
	StatusId    uint `gorm:"foreignKey:FK_TestResults_Statuses"`
	Start       time.Time
	Finish      time.Time
}

func (conn *MySqlConnection) CreateTestresult(data *Testresult) {
	conn.DB.Create(data)
}

func (conn *MySqlConnection) DeleteTestresultById(id uint) {
	conn.DB.Delete(&Testresult{Id: id})
}

func (conn *MySqlConnection) UpdateTestresult(data *Testresult) {
	conn.DB.Save(data)
}

func (conn *MySqlConnection) SelectTestresultsAll() *[]Testresult {
	res := new([]Testresult)
	_ = conn.DB.Find(res)
	return res
}

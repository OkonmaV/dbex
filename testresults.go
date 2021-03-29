package dbex

import "time"

// const format = "2006-01-02 15:04:05.00"

type Testresult struct {
	Id          uint `gorm:"primaryKey"`
	TestPointId uint `gorm:"foreignKey:FK_TestResults_TestPoints_TestPointId;column:TestPointId"`
	StatusId    uint `gorm:"foreignKey:FK_TestResults_Statuses;column:StatusId"`
	Start       time.Time
	Finish      time.Time
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

func (conn *MySqlConnection) SelectTestresultsAll() (*[]Testresult, error) {
	res := new([]Testresult)
	err := conn.DB.Find(res).Error
	return res, err
}

func (conn *MySqlConnection) SelectTestresultById(id uint) (*Testresult, error) {
	res := &Testresult{}
	err := conn.DB.Where("Id = ?", id).First(res).Error
	return res, err
}

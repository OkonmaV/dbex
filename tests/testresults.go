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

func (conf *Conf) CreateTestresult(data *Testresult) {
	conf.db.Create(data)
}

func (conf *Conf) DeleteTestresultById(id uint) {
	conf.db.Delete(&Testresult{Id: id})
}

func (conf *Conf) UpdateTestresult(data *Testresult) {
	conf.db.Save(data)
}

func (conf *Conf) SelectTestresultsAll() *[]Testresult {
	res := new([]Testresult)
	_ = conf.db.Find(res)
	return res
}

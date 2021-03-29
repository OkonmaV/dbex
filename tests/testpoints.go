package dbex

type Testpoint struct {
	Id             uint `gorm:"primaryKey"`
	TestEngineerId uint `gorm:"foreignKey:FK_TestPoints_Testers_TestEngineerId"`
	TestCaseId     uint `gorm:"foreignKey:FK_TestPoints_TestCases_TestCaseId"`
	TestSetId      uint `gorm:"foreignKey:FK_TestPoints_TestSets_TestSetId"`
}

func (conf *Conf) CreateTestpoint(data *Testpoint) {
	conf.db.Create(data)
}

func (conf *Conf) DeleteTestpointById(id uint) {
	conf.db.Delete(&Testpoint{Id: id})
}

func (conf *Conf) UpdateTestpoint(data *Testpoint) {
	conf.db.Save(data)
}

func (conf *Conf) SelectTestpointsAll() *[]Testpoint {
	res := new([]Testpoint)
	_ = conf.db.Find(res)
	return res
}

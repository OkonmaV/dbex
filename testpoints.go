package dbex

type Testpoint struct {
	Id             uint `gorm:"primaryKey"`
	TestEngineerId uint `gorm:"foreignKey:FK_TestPoints_Testers_TestEngineerId"`
	TestCaseId     uint `gorm:"foreignKey:FK_TestPoints_TestCases_TestCaseId"`
	TestSetId      uint `gorm:"foreignKey:FK_TestPoints_TestSets_TestSetId"`
}

func (conn *MySqlConnection) CreateTestpoint(data *Testpoint) {
	conn.DB.Create(data)
}

func (conn *MySqlConnection) DeleteTestpointById(id uint) {
	conn.DB.Delete(&Testpoint{Id: id})
}

func (conn *MySqlConnection) UpdateTestpoint(data *Testpoint) {
	conn.DB.Save(data)
}

func (conn *MySqlConnection) SelectTestpointsAll() *[]Testpoint {
	res := new([]Testpoint)
	_ = conn.DB.Find(res)
	return res
}

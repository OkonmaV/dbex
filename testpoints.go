package dbex

type Testpoint struct {
	Id             uint `gorm:"primaryKey"`
	TestEngineerId uint `gorm:"foreignKey:FK_TestPoints_Testers_TestEngineerId;column:testengineerid"`
	TestCaseId     uint `gorm:"foreignKey:FK_TestPoints_TestCases_TestCaseId;column:testcaseid"`
	TestSetId      uint `gorm:"foreignKey:FK_TestPoints_TestSets_TestSetId;column:testsetid"`
}

func (conn *MySqlConnection) CreateTestpoint(data *Testpoint) error {
	return conn.DB.Create(data).Error
}

func (conn *MySqlConnection) DeleteTestpointById(id uint) error {
	return conn.DB.Delete(&Testpoint{Id: id}).Error
}

func (conn *MySqlConnection) UpdateTestpoint(data *Testpoint) error {
	return conn.DB.Save(data).Error
}

func (conn *MySqlConnection) SelectTestpointsAll() (*[]Testpoint, error) {
	res := new([]Testpoint)
	err := conn.DB.Find(res).Error
	return res, err
}

func (conn *MySqlConnection) SelectTestpointById(id uint) (*Testpoint, error) {
	res := &Testpoint{}
	err := conn.DB.Where("Id = ?", id).First(res).Error
	return res, err
}

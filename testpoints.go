package dbex

type Testpoint struct {
	Id             uint     `gorm:"primaryKey"`
	TestEngineerId uint     `gorm:"column:TestEngineerId"`
	Tester         Tester   `gorm:"foreignKey:TestEngineerId"`
	TestCaseId     uint     `gorm:"column:TestCaseId"`
	Testcase       Testcase `gorm:"foreignKey:TestCaseId"`
	TestSetId      uint     `gorm:"column:TestSetId"`
	Testset        Testset  `gorm:"foreignKey:TestSetId"`
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

func (conn *MySqlConnection) SelectAllTestpoints() ([]Testpoint, error) {
	var res []Testpoint
	err := conn.DB.Find(&res).Error
	return res, err
}

func (conn *MySqlConnection) SelectTestpointById(id uint) (Testpoint, error) {
	var res Testpoint
	err := conn.DB.First(&res, id).Error
	return res, err
}

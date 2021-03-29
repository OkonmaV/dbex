package dbex

type Testset struct {
	Id         uint `gorm:"primaryKey"`
	Name       string
	TestPlanId uint `gorm:"foreignKey:FK_TestSets_TestPlans_TestPlanId;column:TestPlanId"`
}

func (conn *MySqlConnection) CreateTestset(data *Testset) error {
	return conn.DB.Create(data).Error
}

func (conn *MySqlConnection) DeleteTestsetById(id uint) error {
	return conn.DB.Delete(&Testset{Id: id}).Error
}

func (conn *MySqlConnection) UpdateTestset(data *Testset) error {
	return conn.DB.Save(data).Error
}

func (conn *MySqlConnection) SelectTestsetsAll() (*[]Testset, error) {
	res := new([]Testset)
	err := conn.DB.Find(res).Error
	return res, err
}

package dbex

type Testset struct {
	Id         uint `gorm:"primaryKey"`
	Name       string
	TestPlanId uint     `gorm:"column:TestPlanId"`
	Testplan   Testplan `gorm:"foreignKey:TestPlanId"`
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

func (conn *MySqlConnection) SelectAllTestsets() ([]Testset, error) {
	var res []Testset
	err := conn.DB.Find(&res).Error
	return res, err
}

func (conn *MySqlConnection) SelectTestsetById(id uint) (Testset, error) {
	var res Testset
	err := conn.DB.First(&res, id).Error
	return res, err
}

func (conn *MySqlConnection) SelectTestsetsByTester(testerId uint) ([]Testset, error) {
	var res []Testset
	err := conn.DB.Select("testsets.id, testsets.name, testsets.testplanid").Joins("inner join testpoints on testpoints.testsetid = testsets.id").Where("testpoints.testengineerid = ?", testerId).Find(&res).Error
	return res, err
}

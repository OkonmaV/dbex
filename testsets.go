package dbex

type Testset struct {
	Id         uint `gorm:"primaryKey"`
	Name       string
	TestPlanId uint `gorm:"foreignKey:FK_TestSets_TestPlans_TestPlanId;column:TestPlanId"`
}

func (conf *Conf) CreateTestset(data *Testset) error {
	return conf.Db.Create(data).Error
}

func (conf *Conf) DeleteTestsetById(id uint) error {
	return conf.Db.Delete(&Testset{Id: id}).Error
}

func (conf *Conf) UpdateTestset(data *Testset) error {
	return conf.Db.Save(data).Error
}

func (conf *Conf) SelectTestsetsAll() (*[]Testset, error) {
	res := new([]Testset)
	err := conf.Db.Find(res).Error
	return res, err
}

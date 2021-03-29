package dbex

type Testplan struct {
	Id   uint `gorm:"primaryKey"`
	Name string
}

func (conf *Conf) CreateTestplan(data *Testplan) error {
	return conf.Db.Create(data).Error
}

func (conf *Conf) DeleteTestplanById(id uint) error {
	return conf.Db.Delete(&Testplan{Id: id}).Error
}
func (conf *Conf) UpdateTestplan(data *Testplan) error {
	return conf.Db.Save(data).Error
}
func (conf *Conf) SelectTestplansAll() (*[]Testplan, error) {
	res := new([]Testplan)
	err := conf.Db.Find(res).Error
	return res, err
}

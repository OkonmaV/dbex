package dbex

type Department struct {
	Id          uint `gorm:"primaryKey"`
	Description string
}

func (conf *Conf) CreateDepartment(data *Department) error {
	return conf.Db.Create(data).Error
}

func (conf *Conf) DeleteDepartmentById(id uint) error {
	return conf.Db.Delete(&Department{Id: id}).Error
}

func (conf *Conf) UpdateDepartment(data *Department) error {
	return conf.Db.Save(data).Error
}

func (conf *Conf) SelectDepartmentsAll() (*[]Department, error) {
	res := new([]Department)
	err := conf.Db.Find(res).Error
	return res, err
}

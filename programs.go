package dbex

type Program struct {
	Id   uint `gorm:"primaryKey"`
	Name string
}

func (conf *Conf) CreateProgram(data *Program) error {
	return conf.Db.Create(data).Error
}

func (conf *Conf) DeleteProgramById(id uint) error {
	return conf.Db.Delete(&Program{Id: id}).Error
}
func (conf *Conf) UpdateProgram(data *Program) error {
	return conf.Db.Save(data).Error
}
func (conf *Conf) SelectProgramsAll() (*[]Program, error) {
	res := new([]Program)
	err := conf.Db.Find(res).Error
	return res, err
}

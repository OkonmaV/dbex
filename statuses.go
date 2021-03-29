package dbex

type Status struct {
	Id     uint `gorm:"primaryKey"`
	Status string
}

func (conf *Conf) CreateStatus(data *Status) error {
	return conf.Db.Create(data).Error
}

func (conf *Conf) DeleteStatusById(id uint) error {
	return conf.Db.Delete(&Status{}, id).Error
}
func (conf *Conf) UpdateStatus(data *Status) error {
	return conf.Db.Save(data).Error
}

// func (conf *Conf) UpdateStatusField(data *Status, field string, status string) {
// 	conf.db.Model(data).Update(field, status)
// }
func (conf *Conf) SelectStatusesAll() (*[]Status, error) {
	res := new([]Status)
	err := conf.Db.Find(res).Error
	return res, err
}

// func (conf *Conf) SelectStatusesByField(fieldname string) []*Status {
// 	res := []*Status{}
// 	conf.db.Where(suckutils.ConcatThree()).Find(res)
// 	return res
// }

package dbex

type Position struct {
	Id          uint `gorm:"primaryKey"`
	Description string
	Code        string
}

func (conf *Conf) CreatePosition(data *Position) error {
	return conf.Db.Create(data).Error
}

func (conf *Conf) DeletePositionById(id uint) error {
	return conf.Db.Delete(&Position{Id: id}).Error
}

func (conf *Conf) UpdatePosition(data *Position) error {
	return conf.Db.Save(data).Error
}

func (conf *Conf) SelectPositionsAll() (*[]Position, error) {
	res := new([]Position)
	err := conf.Db.Find(res).Error
	return res, err
}

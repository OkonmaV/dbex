package dbex

type Programversion struct {
	Id        uint `gorm:"primaryKey"`
	Major     uint
	Minor     uint
	ProgramId uint `gorm:"foreignKey:FK_ProgramVersions_Programs_ProgramId"`
}

func (conf *Conf) CreateProgramversion(data *Programversion) {
	conf.db.Create(data)
}

func (conf *Conf) DeleteProgramversionById(id uint) {
	conf.db.Delete(&Programversion{Id: id})
}

func (conf *Conf) UpdateProgramversion(data *Programversion) {
	conf.db.Save(data)
}

func (conf *Conf) SelectProgramversionsAll() *[]Programversion {
	res := new([]Programversion)
	_ = conf.db.Find(res)
	return res
}

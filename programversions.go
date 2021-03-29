package dbex

type Programversion struct {
	Id        uint `gorm:"primaryKey"`
	Major     uint
	Minor     uint
	ProgramId uint `gorm:"foreignKey:FK_ProgramVersions_Programs_ProgramId"`
}

func (conn *MySqlConnection) CreateProgramversion(data *Programversion) {
	conn.DB.Create(data)
}

func (conn *MySqlConnection) DeleteProgramversionById(id uint) {
	conn.DB.Delete(&Programversion{Id: id})
}

func (conn *MySqlConnection) UpdateProgramversion(data *Programversion) {
	conn.DB.Save(data)
}

func (conn *MySqlConnection) SelectProgramversionsAll() *[]Programversion {
	res := new([]Programversion)
	_ = conn.DB.Find(res)
	return res
}

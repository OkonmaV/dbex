package dbex

type Programversion struct {
	Id        uint `gorm:"primaryKey"`
	Major     uint
	Minor     uint
	ProgramId uint `gorm:"foreignKey:FK_ProgramVersions_Programs_ProgramId;column:ProgramId"`
}

func (conn *MySqlConnection) CreateProgramversion(data *Programversion) error {
	return conn.DB.Create(data).Error
}

func (conn *MySqlConnection) DeleteProgramversionById(id uint) error {
	return conn.DB.Delete(&Programversion{Id: id}).Error
}

func (conn *MySqlConnection) UpdateProgramversion(data *Programversion) error {
	return conn.DB.Save(data).Error
}

func (conn *MySqlConnection) SelectProgramversionsAll() (*[]Programversion, error) {
	res := new([]Programversion)
	err := conn.DB.Find(res).Error
	return res, err
}

func (conn *MySqlConnection) SelectProgramversionById(id uint) (*Programversion, error) {
	res := &Programversion{}
	err := conn.DB.Where("Id = ?", id).First(res).Error
	return res, err
}

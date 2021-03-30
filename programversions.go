package dbex

type Programversion struct {
	Id        uint    `gorm:"primaryKey"`
	Major     uint    `gorm:"not null"`
	Minor     uint    `gorm:"not null"`
	ProgramId uint    `gorm:"column:ProgramId"`
	Program   Program `gorm:"foreignKey:ProgramId"`
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

func (conn *MySqlConnection) SelectAllProgramversions() ([]Programversion, error) {
	var res []Programversion
	err := conn.DB.Find(&res).Error
	return res, err
}

func (conn *MySqlConnection) SelectProgramversionById(id uint) (Programversion, error) {
	var res Programversion
	err := conn.DB.First(&res, id).Error
	return res, err
}

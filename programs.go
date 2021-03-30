package dbex

type Program struct {
	Id   uint `gorm:"primaryKey"`
	Name string
}

func (conn *MySqlConnection) CreateProgram(data *Program) error {
	return conn.DB.Create(data).Error
}

func (conn *MySqlConnection) DeleteProgramById(id uint) error {
	return conn.DB.Delete(&Program{Id: id}).Error
}
func (conn *MySqlConnection) UpdateProgram(data *Program) error {
	return conn.DB.Save(data).Error
}
func (conn *MySqlConnection) SelectAllPrograms() ([]Program, error) {
	var res []Program
	err := conn.DB.Find(&res).Error
	return res, err
}

func (conn *MySqlConnection) SelectProgramById(id uint) (Program, error) {
	var res Program
	err := conn.DB.First(&res, id).Error
	return res, err
}

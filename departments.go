package dbex

type Department struct {
	Id          uint `gorm:"primaryKey"`
	Description string
}

func (conn *MySqlConnection) CreateDepartment(data *Department) error {
	return conn.DB.Create(data).Error
}

func (conn *MySqlConnection) DeleteDepartmentById(id uint) error {
	return conn.DB.Delete(&Department{Id: id}).Error
}

func (conn *MySqlConnection) UpdateDepartment(data *Department) error {
	return conn.DB.Save(data).Error
}

func (conn *MySqlConnection) SelectDepartmentsAll() (*[]Department, error) {
	res := new([]Department)
	err := conn.DB.Find(res).Error
	return res, err
}

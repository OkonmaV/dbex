package dbex

type Testplan struct {
	Id   uint `gorm:"primaryKey"`
	Name string
}

func (conn *MySqlConnection) CreateTestplan(data *Testplan) error {
	return conn.DB.Create(data).Error
}

func (conn *MySqlConnection) DeleteTestplanById(id uint) error {
	return conn.DB.Delete(&Testplan{Id: id}).Error
}
func (conn *MySqlConnection) UpdateTestplan(data *Testplan) error {
	return conn.DB.Save(data).Error
}
func (conn *MySqlConnection) SelectTestplansAll() (*[]Testplan, error) {
	res := new([]Testplan)
	err := conn.DB.Find(res).Error
	return res, err
}

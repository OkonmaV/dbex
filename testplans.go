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
func (conn *MySqlConnection) SelectAllTestplans() ([]Testplan, error) {
	var res []Testplan
	err := conn.DB.Find(&res).Error
	return res, err
}

func (conn *MySqlConnection) SelectTestplanById(id uint) (Testplan, error) {
	var res Testplan
	err := conn.DB.First(&res, id).Error
	return res, err
}

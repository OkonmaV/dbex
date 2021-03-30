package dbex

type Testcase struct {
	Id             uint `gorm:"primaryKey"`
	Name           string
	Designation    uint
	Steps          string
	ForId          uint           `gorm:"not null; column:ForId"`
	Programversion Programversion `gorm:"foreignKey:ForId"`
}

func (conn *MySqlConnection) CreateTestcase(data *Testcase) error {
	return conn.DB.Create(data).Error
}

func (conn *MySqlConnection) DeleteTestcaseById(id uint) error {
	return conn.DB.Delete(&Testcase{Id: id}).Error
}

func (conn *MySqlConnection) UpdateTestcase(data *Testcase) error {
	return conn.DB.Save(data).Error
}

func (conn *MySqlConnection) SelectAllTestcases() ([]Testcase, error) {
	var res []Testcase
	err := conn.DB.Find(&res).Error
	return res, err
}

func (conn *MySqlConnection) SelectTestcaseById(id uint) (Testcase, error) {
	var res Testcase
	err := conn.DB.First(&res, id).Error
	return res, err
}

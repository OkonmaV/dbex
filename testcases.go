package dbex

type Testcase struct {
	Id          uint `gorm:"primaryKey"`
	Name        string
	Designation uint
	Steps       string
	ForId       uint `gorm:"foreignKey:FK_TestCases_ProgramVersions_FromId;column:ForId"`
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

func (conn *MySqlConnection) SelectTestcasesAll() (*[]Testcase, error) {
	res := new([]Testcase)
	err := conn.DB.Find(res).Error
	return res, err
}

func (conn *MySqlConnection) SelectTestcaseById(id uint) (*Testcase, error) {
	res := &Testcase{}
	err := conn.DB.Where("Id = ?", id).First(res).Error
	return res, err
}

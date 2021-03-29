package dbex

type Testcase struct {
	Id          uint `gorm:"primaryKey"`
	Name        string
	Designation uint
	Steps       string
	ForId       uint `gorm:"foreignKey:FK_TestCases_ProgramVersions_FromId"`
}

func (conn *MySqlConnection) CreateTestcase(data *Testcase) {
	conn.DB.Create(data)
}

func (conn *MySqlConnection) DeleteTestcaseById(id uint) {
	conn.DB.Delete(&Testcase{Id: id})
}

func (conn *MySqlConnection) UpdateTestcase(data *Testcase) {
	conn.DB.Save(data)
}

func (conn *MySqlConnection) SelectTestcasesAll() *[]Testcase {
	res := new([]Testcase)
	_ = conn.DB.Find(res)
	return res
}

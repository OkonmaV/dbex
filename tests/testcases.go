package dbex

type Testcase struct {
	Id          uint `gorm:"primaryKey"`
	Name        string
	Designation uint
	Steps       string
	ForId       uint `gorm:"foreignKey:FK_TestCases_ProgramVersions_FromId"`
}

func (conf *Conf) CreateTestcase(data *Testcase) {
	conf.db.Create(data)
}

func (conf *Conf) DeleteTestcaseById(id uint) {
	conf.db.Delete(&Testcase{Id: id})
}

func (conf *Conf) UpdateTestcase(data *Testcase) {
	conf.db.Save(data)
}

func (conf *Conf) SelectTestcasesAll() *[]Testcase {
	res := new([]Testcase)
	_ = conf.db.Find(res)
	return res
}

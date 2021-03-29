package dbex

type Tester struct {
	Id           uint `gorm:"primaryKey"`
	Name         string
	Surname      string
	PositionId   uint `gorm:"foreignKey:FK_Testers_Positions_PositionId"`
	DepartmentId uint `gorm:"foreignKey:FK_Testers_Departments_DepartmentId"`
}

func (conf *Conf) CreateTester(data *Tester) {
	conf.db.Create(data)
}

func (conf *Conf) DeleteTesterById(id uint) {
	conf.db.Delete(&Tester{Id: id})
}

func (conf *Conf) UpdateTester(data *Tester) {
	conf.db.Save(data)
}

func (conf *Conf) SelectTestersAll() *[]Tester {
	res := new([]Tester)
	_ = conf.db.Find(res)
	return res
}

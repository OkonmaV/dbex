package dbex

type Tester struct {
	Id           uint `gorm:"primaryKey"`
	Name         string
	Surname      string
	PositionId   uint `gorm:"foreignKey:FK_Testers_Positions_PositionId"`
	DepartmentId uint `gorm:"foreignKey:FK_Testers_Departments_DepartmentId"`
}

func (conn *MySqlConnection) CreateTester(data *Tester) {
	conn.DB.Create(data)
}

func (conn *MySqlConnection) DeleteTesterById(id uint) {
	conn.DB.Delete(&Tester{Id: id})
}

func (conn *MySqlConnection) UpdateTester(data *Tester) {
	conn.DB.Save(data)
}

func (conn *MySqlConnection) SelectTestersAll() *[]Tester {
	res := new([]Tester)
	_ = conn.DB.Find(res)
	return res
}

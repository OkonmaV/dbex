package dbex

type Tester struct {
	Id           uint `gorm:"primaryKey"`
	Name         string
	Surname      string
	PositionId   uint `gorm:"foreignKey:FK_Testers_Positions_PositionId;column:positionid"`
	DepartmentId uint `gorm:"foreignKey:FK_Testers_Departments_DepartmentId;column:departmentid"`
}

func (conn *MySqlConnection) CreateTester(data *Tester) error {
	return conn.DB.Create(data).Error
}

func (conn *MySqlConnection) DeleteTesterById(id uint) error {
	return conn.DB.Delete(&Tester{Id: id}).Error
}

func (conn *MySqlConnection) UpdateTester(data *Tester) error {
	return conn.DB.Save(data).Error
}

func (conn *MySqlConnection) SelectTestersAll() (*[]Tester, error) {
	res := new([]Tester)
	err := conn.DB.Find(res).Error
	return res, err
}

func (conn *MySqlConnection) SelectTesterById(id uint) (*Tester, error) {
	res := &Tester{}
	err := conn.DB.Where("Id = ?", id).First(res).Error
	return res, err
}

func (conn *MySqlConnection) SelectTestersWithoutTests() ([]Tester, error) {
	var res []Tester
	err := conn.DB.Select("testers.id, testers.name, testers.surname, testers.positionid, testers.departmentid").Joins("left join testpoints on testpoints.testengineerid = testers.id").Where("testpoints.id is NULL").Find(&res).Error
	return res, err
}

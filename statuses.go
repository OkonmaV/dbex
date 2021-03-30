package dbex

type Status struct {
	Id     uint `gorm:"primaryKey"`
	Status string
}

func (conn *MySqlConnection) CreateStatus(data *Status) error {
	return conn.DB.Create(data).Error
}

func (conn *MySqlConnection) DeleteStatusById(id uint) error {
	return conn.DB.Delete(&Status{}, id).Error
}
func (conn *MySqlConnection) UpdateStatus(data *Status) error {
	return conn.DB.Save(data).Error
}

// func (conn *MySqlConnection) UpdateStatusField(data *Status, field string, status string) {
// 	conn.DB.Model(data).Update(field, status)
// }
func (conn *MySqlConnection) SelectAllStatuses() ([]Status, error) {
	var res []Status
	err := conn.DB.Find(&res).Error
	return res, err
}

func (conn *MySqlConnection) SelectStatusById(id uint) (Status, error) {
	var res Status
	err := conn.DB.First(&res, id).Error
	return res, err
}

// func (conn *MySqlConnection) SelectStatusesByField(fieldname string) []*Status {
// 	res := []*Status{}
// 	conn.DB.Where(suckutils.ConcatThree()).Find(res)
// 	return res
// }

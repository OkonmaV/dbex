package dbex

type Position struct {
	Id          uint `gorm:"primaryKey"`
	Description string
	Code        string
}

func (conn *MySqlConnection) CreatePosition(data *Position) error {
	return conn.DB.Create(data).Error
}

func (conn *MySqlConnection) DeletePositionById(id uint) error {
	return conn.DB.Delete(&Position{Id: id}).Error
}

func (conn *MySqlConnection) UpdatePosition(data *Position) error {
	return conn.DB.Save(data).Error
}

func (conn *MySqlConnection) SelectAllPositions() ([]Position, error) {
	var res []Position
	err := conn.DB.Find(&res).Error
	return res, err
}

func (conn *MySqlConnection) SelectPositionById(id uint) (Position, error) {
	var res Position
	err := conn.DB.First(&res, id).Error
	return res, err
}

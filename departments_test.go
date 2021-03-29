package dbex_test

import (
	"dbex"
	"testing"
)

type TestDataItemDepartments struct {
	inputId  uint
	input    *dbex.Department
	isBroken bool
}

func TestCreateDepartment(t *testing.T) {

	dataItems := []TestDataItemDepartments{
		{0, &dbex.Department{Id: 1, Description: "test"}, true},
		{0, &dbex.Department{Description: "test"}, false},
	}

	conf, err := dbex.GetConf()
	if err != nil {
		t.Error("DB connection error: ", err)
		return
	}

	for _, item := range dataItems {
		err := conf.CreateDepartment(item.input)

		if item.isBroken {
			if err == nil {
				t.Error("\nFAILED: expected an error, but no error catched at Inserting ", item.input)
			} else {
				t.Log("\nPASSED: expected an error, got an error at Inserting ", item.input, "\nerror: ", err)
			}
		} else {
			if err != nil {
				t.Error("\nFAILED: non-expected error at Inserting ", item.input, "\nerror: ", err)
			} else {
				t.Log("\nPASSED: no error at Inserting ", item.input)
			}
		}
	}

}

func TestDeleteDepartmentById(t *testing.T) {

	conf, err := dbex.GetConf()
	if err != nil {
		t.Error("FAILED: DB connection error: ", err)
		return
	}

	foo := &dbex.Department{}
	err = conf.Db.Where("Description = ?", "test").Last(foo).Error
	if foo == nil && err == nil {
		_ = conf.Db.Create(&dbex.Department{Description: "test"})
		_ = conf.Db.Where("Description = ?", "test").Last(foo)
	} else if err != nil {
		t.Error("FAILED: some error : ", err)
	}

	err = conf.DeleteDepartmentById(foo.Id)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Delete by id ", foo.Id, "\nerror: ", err)
	} else {
		t.Log("\nPASSED: no error at Delete by id ", foo.Id)
	}
}

func TestUpdateDepartment(t *testing.T) {

	conf, err := dbex.GetConf()
	if err != nil {
		t.Error("FAILED: DB connection error: ", err)
		return
	}

	foo := &dbex.Department{}
	err = conf.Db.Where("Description = ?", "test").Last(foo).Error
	if foo.Id == 0 {
		_ = conf.Db.Create(&dbex.Department{Description: "test"})
		_ = conf.Db.Where("Description = ?", "test").Last(foo)
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	err = conf.UpdateDepartment(foo)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Update ", foo.Id, "\nerror: ", err)
	} else {
		t.Log("\nPASSED: no error at Update ", foo.Id)
	}
}

func TestSelectDepartmentsAll(t *testing.T) {
	conf, err := dbex.GetConf()
	if err != nil {
		t.Error("FAILED: DB connection error: ", err)
		return
	}
	foo, err := conf.SelectDepartmentsAll()

	if err != nil {
		t.Error("\nFAILED: non-expected error at SelectAll\nerror: ", err)
	} else if len(*foo) == 0 {
		t.Error("\nFAILED: returned empty slice at SelectAll")
	} else {
		t.Log("\nPASSED: returned at SelectAll:\n", foo)
	}

}

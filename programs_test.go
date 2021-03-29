package dbex_test

import (
	"dbex"
	"testing"
)

type TestDataItemPrograms struct {
	inputId  uint
	input    *dbex.Program
	isBroken bool
}

func TestCreateProgram(t *testing.T) {

	dataItems := []TestDataItemPrograms{
		{0, &dbex.Program{Id: 1, Name: "test"}, true},
		{0, &dbex.Program{Name: "test"}, false},
	}

	conf, err := dbex.GetConf()
	if err != nil {
		t.Error("DB connection error: ", err)
		return
	}

	for _, item := range dataItems {
		err := conf.CreateProgram(item.input)

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

func TestDeleteProgramById(t *testing.T) {

	conf, err := dbex.GetConf()
	if err != nil {
		t.Error("FAILED: DB connection error: ", err)
		return
	}

	foo := &dbex.Program{}
	err = conf.Db.Where("Name = ?", "test").Last(foo).Error
	if foo == nil && err == nil {
		_ = conf.Db.Create(&dbex.Program{Name: "test"})
		_ = conf.Db.Where("Name = ?", "test").Last(foo)
	} else if err != nil {
		t.Error("FAILED: some error : ", err)
	}

	err = conf.DeleteProgramById(foo.Id)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Delete by id ", foo.Id, "\nerror: ", err)
	} else {
		t.Log("\nPASSED: no error at Delete by id ", foo.Id)
	}
}

func TestUpdateProgram(t *testing.T) {

	conf, err := dbex.GetConf()
	if err != nil {
		t.Error("FAILED: DB connection error: ", err)
		return
	}

	foo := &dbex.Program{}
	err = conf.Db.Where("Name = ?", "test").Last(foo).Error
	if foo.Id == 0 {
		_ = conf.Db.Create(&dbex.Program{Name: "test"})
		_ = conf.Db.Where("Name = ?", "test").Last(foo)
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	err = conf.UpdateProgram(foo)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Update ", foo.Id, "\nerror: ", err)
	} else {
		t.Log("\nPASSED: no error at Update ", foo.Id)
	}
}

func TestSelectProgramsAll(t *testing.T) {
	conf, err := dbex.GetConf()
	if err != nil {
		t.Error("FAILED: DB connection error: ", err)
		return
	}
	foo, err := conf.SelectProgramsAll()

	if err != nil {
		t.Error("\nFAILED: non-expected error at SelectAll\nerror: ", err)
	} else if len(*foo) == 0 {
		t.Error("\nFAILED: returned empty slice at SelectAll")
	} else {
		t.Log("\nPASSED: returned at SelectAll:\n", foo)
	}

}

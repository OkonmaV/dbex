package dbex_test

import (
	"dbex"
	"testing"
)

type TestDataItemTestplans struct {
	inputId  uint
	input    *dbex.Testplan
	isBroken bool
}

func TestCreateTestplan(t *testing.T) {

	dataItems := []TestDataItemTestplans{
		{0, &dbex.Testplan{Id: 1, Name: "test"}, true},
		{0, &dbex.Testplan{Name: "test"}, false},
	}

	conf, err := dbex.GetConf()
	if err != nil {
		t.Error("DB connection error: ", err)
		return
	}

	for _, item := range dataItems {
		err := conf.CreateTestplan(item.input)

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

func TestDeleteTestplanById(t *testing.T) {

	conf, err := dbex.GetConf()
	if err != nil {
		t.Error("FAILED: DB connection error: ", err)
		return
	}

	foo := &dbex.Testplan{}
	err = conf.Db.Where("Name = ?", "test").Last(foo).Error
	if foo == nil && err == nil {
		_ = conf.Db.Create(&dbex.Testplan{Name: "test"})
		_ = conf.Db.Where("Name = ?", "test").Last(foo)
	} else if err != nil {
		t.Error("FAILED: some error : ", err)
	}

	err = conf.DeleteTestplanById(foo.Id)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Delete by id ", foo.Id, "\nerror: ", err)
	} else {
		t.Log("\nPASSED: no error at Delete by id ", foo.Id)
	}
}

func TestUpdateTestplan(t *testing.T) {

	conf, err := dbex.GetConf()
	if err != nil {
		t.Error("FAILED: DB connection error: ", err)
		return
	}

	foo := &dbex.Testplan{}
	err = conf.Db.Where("Name = ?", "test").Last(foo).Error
	if foo.Id == 0 {
		_ = conf.Db.Create(&dbex.Testplan{Name: "test"})
		_ = conf.Db.Where("Name = ?", "test").Last(foo)
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	err = conf.UpdateTestplan(foo)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Update ", foo.Id, "\nerror: ", err)
	} else {
		t.Log("\nPASSED: no error at Update ", foo.Id)
	}
}

func TestSelectTestplansAll(t *testing.T) {
	conf, err := dbex.GetConf()
	if err != nil {
		t.Error("FAILED: DB connection error: ", err)
		return
	}
	foo, err := conf.SelectTestplansAll()

	if err != nil {
		t.Error("\nFAILED: non-expected error at SelectAll\nerror: ", err)
	} else if len(*foo) == 0 {
		t.Error("\nFAILED: returned empty slice at SelectAll")
	} else {
		t.Log("\nPASSED: returned at SelectAll:\n", foo)
	}

}

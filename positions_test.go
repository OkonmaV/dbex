package dbex_test

import (
	"dbex"
	"testing"
)

type TestDataItemPositions struct {
	inputId  uint
	input    *dbex.Position
	isBroken bool
}

func TestCreatePosition(t *testing.T) {

	dataItems := []TestDataItemPositions{
		{0, &dbex.Position{Id: 1, Description: "test", Code: "test"}, true},
		{0, &dbex.Position{Description: "test", Code: "test"}, false},
	}

	conf, err := dbex.GetConf()
	if err != nil {
		t.Error("DB connection error: ", err)
		return
	}

	for _, item := range dataItems {
		err := conf.CreatePosition(item.input)

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

func TestDeletePositionById(t *testing.T) {

	conf, err := dbex.GetConf()
	if err != nil {
		t.Error("FAILED: DB connection error: ", err)
		return
	}

	foo := &dbex.Position{}
	err = conf.Db.Where("Code = ?", "test").Last(foo).Error
	if foo == nil && err == nil {
		_ = conf.Db.Create(&dbex.Position{Code: "test"})
		_ = conf.Db.Where("Code = ?", "test").Last(foo)
	} else if err != nil {
		t.Error("FAILED: some error : ", err)
	}

	err = conf.DeletePositionById(foo.Id)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Delete by id ", foo.Id, "\nerror: ", err)
	} else {
		t.Log("\nPASSED: no error at Delete by id ", foo.Id)
	}
}

func TestUpdatePosition(t *testing.T) {

	conf, err := dbex.GetConf()
	if err != nil {
		t.Error("FAILED: DB connection error: ", err)
		return
	}

	foo := &dbex.Position{}
	err = conf.Db.Where("Code = ?", "test").Last(foo).Error
	if foo.Id == 0 {
		_ = conf.Db.Create(&dbex.Position{Code: "test"})
		_ = conf.Db.Where("Code = ?", "test").Last(foo)
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	err = conf.UpdatePosition(foo)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Update ", foo.Id, "\nerror: ", err)
	} else {
		t.Log("\nPASSED: no error at Update ", foo.Id)
	}
}

func TestSelectPositionsAll(t *testing.T) {
	conf, err := dbex.GetConf()
	if err != nil {
		t.Error("FAILED: DB connection error: ", err)
		return
	}
	foo, err := conf.SelectPositionsAll()

	if err != nil {
		t.Error("\nFAILED: non-expected error at SelectAll\nerror: ", err)
	} else if len(*foo) == 0 {
		t.Error("\nFAILED: returned empty slice at SelectAll")
	} else {
		t.Log("\nPASSED: returned at SelectAll:\n", foo)
	}

}

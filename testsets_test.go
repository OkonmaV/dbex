package dbex_test

import (
	"dbex"
	"testing"
)

type TestDataItemTestsets struct {
	inputId  uint
	input    *dbex.Testset
	isBroken bool
}

func TestCreateTestset(t *testing.T) {

	dataItems := []TestDataItemTestsets{
		{0, &dbex.Testset{Id: 1, Name: "test", TestPlanId: 1}, true},
		{0, &dbex.Testset{Name: "test", TestPlanId: 1}, false},
		{0, &dbex.Testset{Name: "test", TestPlanId: 9000000}, true},
	}

	conf, err := dbex.GetConf()
	if err != nil {
		t.Error("DB connection error: ", err)
		return
	}

	for _, item := range dataItems {
		err := conf.CreateTestset(item.input)

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

func TestDeleteTestsetById(t *testing.T) {

	conf, err := dbex.GetConf()
	if err != nil {
		t.Error("FAILED: DB connection error: ", err)
		return
	}

	foo := &dbex.Testset{}
	err = conf.Db.Where("Name = ?", "test").Last(foo).Error
	if foo == nil && err == nil {
		_ = conf.Db.Create(&dbex.Testset{Name: "test"})
		_ = conf.Db.Where("Name = ?", "test").Last(foo)
	} else if err != nil {
		t.Error("FAILED: some error : ", err)
	}

	err = conf.DeleteTestsetById(foo.Id)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Delete by id ", foo.Id, "\nerror: ", err)
	} else {
		t.Log("\nPASSED: no error at Delete by id ", foo.Id)
	}
}

func TestUpdateTestset(t *testing.T) {

	conf, err := dbex.GetConf()
	if err != nil {
		t.Error("FAILED: DB connection error: ", err)
		return
	}

	foo := &dbex.Testset{}
	err = conf.Db.Where("Name = ?", "test").Last(foo).Error
	if foo.Id == 0 {
		_ = conf.Db.Create(&dbex.Testset{Name: "test", TestPlanId: 1})
		_ = conf.Db.Where("Name = ?", "test").Last(foo)
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	dataItems := []*TestDataItemTestsets{
		{0, &dbex.Testset{Id: foo.Id, Name: "test", TestPlanId: 1}, false},
		{0, &dbex.Testset{Id: foo.Id, Name: "test", TestPlanId: 0}, true},
	}

	for _, item := range dataItems {
		err := conf.UpdateTestset(item.input)

		if item.isBroken {
			if err == nil {
				t.Error("\nFAILED: expected an error, but no error catched at Update ", item.input)
			} else {
				t.Log("\nPASSED: expected an error, got an error at Update ", item.input, "\nerror: ", err)
			}
		} else {
			if err != nil {
				t.Error("\nFAILED: non-expected error at Update ", item.input, "\nerror: ", err)
			} else {
				t.Log("\nPASSED: no error at Update ", item.input)
			}
		}
	}
}

func TestSelectTestsetsAll(t *testing.T) {
	conf, err := dbex.GetConf()
	if err != nil {
		t.Error("FAILED: DB connection error: ", err)
		return
	}
	foo, err := conf.SelectTestsetsAll()

	if err != nil {
		t.Error("\nFAILED: non-expected error at SelectAll\nerror: ", err)
	} else if len(*foo) == 0 {
		t.Error("\nFAILED: returned empty slice at SelectAll")
	} else {
		t.Log("\nPASSED: returned at SelectAll:\n", foo)
	}

}

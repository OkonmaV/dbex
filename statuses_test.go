package dbex_test

import (
	"dbex"
	"testing"
)

type TestDataItemStatuses struct {
	inputId  uint
	input    *dbex.Status
	isBroken bool
}

func TestCreateStatus(t *testing.T) {

	dataItems := []TestDataItemStatuses{
		{0, &dbex.Status{Id: 1, Status: "test"}, true},
		{0, &dbex.Status{Status: "test"}, false},
	}

	conf, err := dbex.GetConf()
	if err != nil {
		t.Error("DB connection error: ", err)
		return
	}

	for _, item := range dataItems {
		err := conf.CreateStatus(item.input)

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

func TestDeleteStatusById(t *testing.T) {

	conf, err := dbex.GetConf()
	if err != nil {
		t.Error("FAILED: DB connection error: ", err)
		return
	}

	foo := &dbex.Status{}
	err = conf.Db.Where("Status = ?", "test").Last(foo).Error
	if foo.Id == 0 && err == nil {
		_ = conf.Db.Create(&dbex.Status{Status: "test"})
		_ = conf.Db.Where("Status = ?", "test").Last(foo)
	} else if err != nil {
		t.Error("FAILED: some error : ", err)
	}

	err = conf.DeleteStatusById(foo.Id)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Delete by id ", foo.Id, "\nerror: ", err)
	} else {
		t.Log("\nPASSED: no error at Delete by id ", foo.Id)
	}
	// dataItems := []TestDataItem{
	// 	{9000000, nil, true},
	// 	{foo.Id, nil, false},
	// }

	// for _, item := range dataItems {

	// 	err := conf.DeleteStatusById(item.inputId)

	// 	if item.isBroken {
	// 		if err == nil {
	// 			t.Error("\nFAILED: expected an error, but no error catched at Delete by id ", item.inputId)
	// 		} else {
	// 			t.Log("\nPASSED: expected an error, got an error at Delete by id ", item.inputId, "\nerror: ", err)
	// 		}
	// 	} else {
	// 		if err != nil {
	// 			t.Error("\nFAILED: non-expected error at Delete by id ", item.inputId, "\nerror: ", err)
	// 		} else {
	// 			t.Log("\nPASSED: no error at Delete by id ", item.inputId)
	// 		}
	// 	}
	// }
}

func TestUpdateStatus(t *testing.T) {

	conf, err := dbex.GetConf()
	if err != nil {
		t.Error("FAILED: DB connection error: ", err)
		return
	}

	foo := &dbex.Status{}
	err = conf.Db.Where("Status = ?", "test").Last(foo).Error
	if foo == nil && err == nil {
		_ = conf.Db.Create(&dbex.Status{Status: "test"})
		_ = conf.Db.Where("Status = ?", "test").First(foo)
	} else if err != nil {
		t.Error("FAILED: some error : ", err)
	}

	err = conf.UpdateStatus(foo)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Update ", foo.Id, "\nerror: ", err)
	} else {
		t.Log("\nPASSED: no error at Update ", foo.Id)
	}
}

func TestSelectStatusesAll(t *testing.T) {
	conf, err := dbex.GetConf()
	if err != nil {
		t.Error("FAILED: DB connection error: ", err)
		return
	}
	foo, err := conf.SelectStatusesAll()

	if err != nil {
		t.Error("\nFAILED: non-expected error at SelectAll\nerror: ", err)
	} else if len(*foo) == 0 {
		t.Error("\nFAILED: returned empty slice at SelectAll")
	} else {
		t.Log("\nPASSED: returned at SelectAll:\n", foo)
	}

}

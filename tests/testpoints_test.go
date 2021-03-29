package dbex_test

import (
	"dbex"
	"testing"
)

type TestDataItemTestpoint struct {
	inputId  uint
	input    *dbex.Testpoint
	isBroken bool
}

func TestCreateTestpoint(t *testing.T) {

	dataItems := []TestDataItemTestpoint{
		{0, &dbex.Testpoint{Id: 1, TestEngineerId: 1}, true},
		{0, &dbex.Testpoint{TestEngineerId: 17, TestCaseId: 1, TestSetId: 1}, false},
		{0, &dbex.Testpoint{TestEngineerId: 0, TestCaseId: 0, TestSetId: 0}, true},
	}

	conn, err := GetConf()
	if err != nil {
		t.Error("DB connection error: ", err)
		return
	}

	for _, item := range dataItems {
		err := conn.CreateTestpoint(item.input)

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

func TestDeleteTestpointById(t *testing.T) {

	conn, err := GetConf()
	if err != nil {
		t.Error("FAILED: DB connection error: ", err)
		return
	}

	foo := &dbex.Testpoint{}
	err = conn.DB.Where("TestEngineerId = ?", 17).Last(foo).Error
	if foo.Id == 0 && err == nil {
		_ = conn.DB.Create(&dbex.Testpoint{TestEngineerId: 17, TestCaseId: 1, TestSetId: 1})
		_ = conn.DB.Where("TestEngineerId = ?", 17).Last(foo)
	} else if err != nil {
		t.Error("FAILED: some error : ", err)
	}

	err = conn.DeleteTestpointById(foo.Id)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Delete by id ", foo.Id, "\nerror: ", err)
	} else {
		t.Log("\nPASSED: no error at Delete by id ", foo.Id)
	}
}

func TestUpdateTestpoint(t *testing.T) {

	conn, err := GetConf()
	if err != nil {
		t.Error("FAILED: DB connection error: ", err)
		return
	}

	foo := &dbex.Testpoint{}
	err = conn.DB.Where("TestEngineerId = ?", 17).Last(foo).Error
	if foo.Id == 0 {
		_ = conn.DB.Create(&dbex.Testpoint{TestEngineerId: 17, TestCaseId: 1, TestSetId: 1})
		_ = conn.DB.Where("TestEngineerId = ?", 17).Last(foo)
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	dataItems := []*TestDataItemTestpoint{
		{0, &dbex.Testpoint{Id: foo.Id, TestEngineerId: 1, TestCaseId: 1, TestSetId: 1}, false},
		{0, &dbex.Testpoint{Id: foo.Id, TestEngineerId: 2000}, true},
	}

	for _, item := range dataItems {
		err := conn.UpdateTestpoint(item.input)

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

func TestSelectTestpointsAll(t *testing.T) {
	conn, err := GetConf()
	if err != nil {
		t.Error("FAILED: DB connection error: ", err)
		return
	}
	foo, err := conn.SelectTestpointsAll()

	if err != nil {
		t.Error("\nFAILED: non-expected error at SelectAll\nerror: ", err)
	} else if len(*foo) == 0 {
		t.Error("\nFAILED: returned empty slice at SelectAll")
	} else {
		t.Log("\nPASSED: returned at SelectAll:\n", foo)
	}

}

func TestSelectTestpointById(t *testing.T) {
	dataItems := []TestDataItemStatuses{
		{0, nil, true},
		{1, nil, false},
	}

	conn, err := GetConf()
	if err != nil {
		t.Error("DB connection error: ", err)
		return
	}

	for _, item := range dataItems {
		foo, err := conn.SelectTestpointById(item.inputId)

		if item.isBroken {
			if err == nil {
				t.Error("\nFAILED: expected an error, but no error catched at Select by id ", item.inputId)
			} else {
				t.Log("\nPASSED: expected an error, got an error at Select by id ", foo, "\nerror: ", err)
			}
		} else {
			if err != nil {
				t.Error("\nFAILED: non-expected error at Select by id ", item.inputId, "\nerror: ", err)
			} else if foo.Id != 0 {
				t.Log("\nPASSED: no error at Select by id ", foo)
			} else {
				t.Error("\nFAILED: no error, but relust is nil at Select by id ", item.inputId)
			}
		}
	}

}

package dbex_test

import (
	"dbex"
	"testing"
)

type TestDataItemTestcase struct {
	inputId  uint
	input    *dbex.Testcase
	isBroken bool
}

func TestCreateTestcase(t *testing.T) {

	dataItems := []TestDataItemTestcase{
		{0, &dbex.Testcase{Id: 1, Name: "test", ForId: 1}, true},
		{0, &dbex.Testcase{Name: "test", ForId: 1}, false},
		{0, &dbex.Testcase{Name: "test", ForId: 0}, true},
	}

	conn, err := GetConf()
	if err != nil {
		t.Error("DB connection error: ", err)
		return
	}

	for _, item := range dataItems {
		err := conn.CreateTestcase(item.input)

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

func TestDeleteTestcaseById(t *testing.T) {

	conn, err := GetConf()
	if err != nil {
		t.Error("FAILED: DB connection error: ", err)
		return
	}

	foo := &dbex.Testcase{}
	err = conn.DB.Where("Name = ?", "test").Last(foo).Error
	if foo.Id == 0 && err == nil {
		_ = conn.DB.Create(&dbex.Testcase{Name: "test"})
		_ = conn.DB.Where("Name = ?", "test").Last(foo)
	} else if err != nil {
		t.Error("FAILED: some error : ", err)
	}

	err = conn.DeleteTestcaseById(foo.Id)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Delete by id ", foo.Id, "\nerror: ", err)
	} else {
		t.Log("\nPASSED: no error at Delete by id ", foo.Id)
	}
}

func TestUpdateTestcase(t *testing.T) {

	conn, err := GetConf()
	if err != nil {
		t.Error("FAILED: DB connection error: ", err)
		return
	}

	foo := &dbex.Testcase{}
	err = conn.DB.Where("Name = ?", "test").Last(foo).Error
	if foo.Id == 0 {
		_ = conn.DB.Create(&dbex.Testcase{Name: "test"})
		_ = conn.DB.Where("Name = ?", "test").Last(foo)
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	dataItems := []*TestDataItemTestcase{
		{0, &dbex.Testcase{Id: foo.Id, Name: "test", ForId: 1}, false},
		{0, &dbex.Testcase{Id: foo.Id, Name: "test", ForId: 0}, true},
	}

	for _, item := range dataItems {
		err := conn.UpdateTestcase(item.input)

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

func TestSelectTestcasesAll(t *testing.T) {
	conn, err := GetConf()
	if err != nil {
		t.Error("FAILED: DB connection error: ", err)
		return
	}
	foo, err := conn.SelectTestcasesAll()

	if err != nil {
		t.Error("\nFAILED: non-expected error at SelectAll\nerror: ", err)
	} else if len(*foo) == 0 {
		t.Error("\nFAILED: returned empty slice at SelectAll")
	} else {
		t.Log("\nPASSED: returned at SelectAll:\n", foo)
	}

}

func TestSelectTestcaseById(t *testing.T) {
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
		foo, err := conn.SelectTestcaseById(item.inputId)

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

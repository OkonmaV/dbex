package dbex_test

import (
	"dbex"
	"errors"
	"testing"

	"gorm.io/gorm"
)

type TestDataItemTestcases struct {
	inputId  uint
	input    *dbex.Testcase
	isBroken bool
}

var fkTestcases dbex.Programversion = dbex.Programversion{Id: 550, Major: 550, Minor: 550, ProgramId: 550, Program: dbex.Program{Id: 550, Name: "fk"}}

func TestCreateTestcase(t *testing.T) {

	dataItems := []TestDataItemTestcases{
		{10, &dbex.Testcase{Id: 100, Name: "test1", ForId: 550, Programversion: fkTestcases}, false},
		{200, &dbex.Testcase{Id: 200, Name: "test2", ForId: 1000}, true},
		{300, &dbex.Testcase{Id: 300, Name: "test3"}, true},
		{400, &dbex.Testcase{Id: 100, Name: "test4", ForId: 550, Programversion: fkTestcases}, true},
	}

	for _, item := range dataItems {
		err := DB.CreateTestcase(item.input)

		if item.isBroken {
			if err == nil {
				t.Error("\nFAILED: expected an error, but no error catched at Inserting ", item.input)
			}
		} else {
			if err != nil {
				t.Error("\nFAILED: non-expected error at Inserting ", item.input, "\nerror: ", err)
			}
		}
	}

}

func TestDeleteTestcaseById(t *testing.T) {

	bar := &dbex.Testcase{Name: "test for delete", ForId: 550, Programversion: fkTestcases}
	err := DB.CreateTestcase(bar)
	if err != nil {
		t.Error("FAILED: insert error: ", err, bar)
	}
	foo := &dbex.Testcase{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	err = DB.DeleteTestcaseById(foo.Id)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Delete by id ", foo.Id, "\nerror: ", err)
	}

	foo = &dbex.Testcase{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// passed
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	} else {
		t.Error("FAILED: record not deleted")
	}
}

func TestUpdateTestcase(t *testing.T) {

	bar := &dbex.Testcase{Name: "test for update", ForId: 550, Programversion: fkTestcases}

	// err := DB.DB.Create(&dbex.Testplan{Id: 750, Name: "fk"}).Error
	// if err != nil {
	// 	t.Error("\nFAILED: error at inserting for fk: ", err)
	// }

	err := DB.CreateTestcase(bar)
	if err != nil {
		t.Error("FAILED: insert error: ", err, bar)
	}
	foo := &dbex.Testcase{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	foo.Name = "updated"
	err = DB.UpdateTestcase(foo)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Update ", foo.Id, "\nerror: ", err)
	}

	foo = &dbex.Testcase{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	} else if foo.Name != "updated" {
		t.Error("FAILED: was not updated")
	}
}

func TestSelectAllTestcases(t *testing.T) {

	// clear table
	if err := DB.DB.Exec("delete from testcases").Error; err != nil {
		t.Error("Clear table error:", err)
	}

	// err := DB.DB.Create(&dbex.Testplan{Id: 850, Name: "fk"}).Error
	// if err != nil {
	// 	t.Error("\nFAILED: error at inserting for fk: ", err)
	// }

	dataItems := []TestDataItemTestcases{
		{0, &dbex.Testcase{Name: "test1", ForId: 550, Programversion: fkTestcases}, false},
		{0, &dbex.Testcase{Name: "test2", ForId: 550, Programversion: fkTestcases}, false},
		{0, &dbex.Testcase{Name: "test3", ForId: 550, Programversion: fkTestcases}, false},
	}

	for _, item := range dataItems {
		err := DB.CreateTestcase(item.input)

		if err != nil {
			t.Error("\nFAILED: non-expected error at Inserting ", item.input, "\nerror: ", err)
		}
	}

	foo, err := DB.SelectAllTestcases()

	if err != nil {
		t.Error("\nFAILED: non-expected error at SelectAll\nerror: ", err)
	} else if len(foo) == 0 {
		t.Error("\nFAILED: returned empty slice at SelectAll")
	}

	if len(foo) != len(dataItems) {
		t.Error("\nFAILED: returned count not match")
	}

	for i, item := range foo {
		if item.Id != dataItems[i].input.Id {
			t.Error("\nFAILED: returned id is not match ", item.Id, dataItems[i].input.Id)
		}
	}

}

func TestSelectTestcaseById(t *testing.T) {

	bar := &dbex.Testcase{Id: 500, Name: "test", ForId: 550, Programversion: fkTestcases}
	err := DB.CreateTestcase(bar)
	if err != nil {
		t.Error("FAILED: insert error: ", err)
	}

	foo := &dbex.Testcase{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	if err := DB.DB.Exec("delete from testcases").Error; err != nil {
		t.Error("Clear table error:", err)
	}
}

package dbex_test

import (
	"dbex"
	"errors"
	"testing"

	"gorm.io/gorm"
)

type TestDataItemTestplans struct {
	inputId  uint
	input    *dbex.Testplan
	isBroken bool
}

func TestCreateTestplan(t *testing.T) {

	dataItems := []TestDataItemTestplans{
		{200, &dbex.Testplan{Id: 200, Name: "test insert 1"}, false},
		{100, &dbex.Testplan{Id: 100, Name: "test insert 2"}, false},
		{200, &dbex.Testplan{Id: 200, Name: "test insert 3"}, true},
	}

	for _, item := range dataItems {
		err := DB.CreateTestplan(item.input)

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

func TestDeleteTestplanById(t *testing.T) {

	bar := &dbex.Testplan{Name: "test for delete"}
	err := DB.CreateTestplan(bar)
	if err != nil {
		t.Error("FAILED: insert error: ", err, bar)
	}

	foo := &dbex.Testplan{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	err = DB.DeleteTestplanById(foo.Id)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Delete by id ", foo.Id, "\nerror: ", err)
	}

	foo = &dbex.Testplan{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// passed
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	} else {
		t.Error("FAILED: record not deleted")
	}
}

func TestUpdateTestplan(t *testing.T) {

	bar := &dbex.Testplan{Name: "test for update"}
	err := DB.CreateTestplan(bar)
	if err != nil {
		t.Error("insert error: ", err, bar)
	}

	foo := &dbex.Testplan{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	foo.Name = "updated"
	err = DB.UpdateTestplan(foo)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Update ", foo.Id, "\nerror: ", err)
	}

	foo = &dbex.Testplan{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	} else if foo.Name != "updated" {
		t.Error("FAILED: Name do not updated")
	}
}

func TestSelectAllTestplans(t *testing.T) {

	// clear table
	if err := DB.DB.Exec("delete from testplans").Error; err != nil {
		t.Error("Clear table error:", err)
	}

	dataItems := []TestDataItemTestplans{
		{0, &dbex.Testplan{Name: "test1"}, false},
		{0, &dbex.Testplan{Name: "test2"}, false},
		{0, &dbex.Testplan{Name: "test3"}, false},
	}

	for _, item := range dataItems {
		err := DB.CreateTestplan(item.input)

		if err != nil {
			t.Error("\nFAILED: non-expected error at Inserting ", item.input, "\nerror: ", err)
		}
	}

	foo, err := DB.SelectAllTestplans()

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

func TestSelectTestplanById(t *testing.T) {
	bar := &dbex.Testplan{Id: 500, Name: "test"}
	err := DB.CreateTestplan(bar)
	if err != nil {
		t.Error("FAILED: insert error: ", err)
	}

	foo := &dbex.Testplan{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	if err := DB.DB.Exec("delete from testplans").Error; err != nil {
		t.Error("Clear table error:", err)
	}
}

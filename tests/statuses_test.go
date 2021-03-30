package dbex_test

import (
	"dbex"
	"errors"
	"testing"

	"gorm.io/gorm"
)

type TestDataItemStatuses struct {
	inputId  uint
	input    *dbex.Status
	isBroken bool
}

func TestCreateStatus(t *testing.T) {

	dataItems := []TestDataItemStatuses{
		{200, &dbex.Status{Id: 200, Status: "test insert 1"}, false},
		{100, &dbex.Status{Id: 100, Status: "test insert 2"}, false},
		{200, &dbex.Status{Id: 200, Status: "test insert 3"}, true},
	}

	for _, item := range dataItems {
		err := DB.CreateStatus(item.input)

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

func TestDeleteStatusById(t *testing.T) {

	bar := &dbex.Status{Status: "test for delete"}
	err := DB.CreateStatus(bar)
	if err != nil {
		t.Error("FAILED: insert error: ", err, bar)
	}

	foo := &dbex.Status{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	err = DB.DeleteStatusById(foo.Id)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Delete by id ", foo.Id, "\nerror: ", err)
	}

	foo = &dbex.Status{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// passed
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	} else {
		t.Error("FAILED: record not deleted")
	}
}

func TestUpdateStatus(t *testing.T) {

	bar := &dbex.Status{Status: "test for update"}
	err := DB.CreateStatus(bar)
	if err != nil {
		t.Error("insert error: ", err, bar)
	}

	foo := &dbex.Status{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	foo.Status = "updated"
	err = DB.UpdateStatus(foo)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Update ", foo.Id, "\nerror: ", err)
	}

	foo = &dbex.Status{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	} else if foo.Status != "updated" {
		t.Error("FAILED: Status do not updated")
	}
}

func TestSelectAllStatuses(t *testing.T) {

	// clear table
	if err := DB.DB.Exec("delete from statuses").Error; err != nil {
		t.Error("Clear table error:", err)
	}

	dataItems := []TestDataItemStatuses{
		{0, &dbex.Status{Status: "test1"}, false},
		{0, &dbex.Status{Status: "test2"}, false},
		{0, &dbex.Status{Status: "test3"}, false},
	}

	for _, item := range dataItems {
		err := DB.CreateStatus(item.input)

		if err != nil {
			t.Error("\nFAILED: non-expected error at Inserting ", item.input, "\nerror: ", err)
		}
	}

	foo, err := DB.SelectAllStatuses()

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

func TestSelectStatusById(t *testing.T) {
	bar := &dbex.Status{Id: 500, Status: "test"}
	err := DB.CreateStatus(bar)
	if err != nil {
		t.Error("FAILED: insert error: ", err)
	}

	foo := &dbex.Status{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	if err := DB.DB.Exec("delete from statuses").Error; err != nil {
		t.Error("Clear table error:", err)
	}
}

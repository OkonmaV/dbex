package dbex_test

import (
	"dbex"
	"errors"
	"testing"

	"gorm.io/gorm"
)

type TestDataItemPositions struct {
	inputId  uint
	input    *dbex.Position
	isBroken bool
}

func TestCreatePosition(t *testing.T) {

	dataItems := []TestDataItemPositions{
		{200, &dbex.Position{Id: 200, Description: "test insert 1"}, false},
		{100, &dbex.Position{Id: 100, Description: "test insert 2"}, false},
		{200, &dbex.Position{Id: 200, Description: "test insert 3"}, true},
	}

	for _, item := range dataItems {
		err := DB.CreatePosition(item.input)

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

func TestDeletePositionById(t *testing.T) {

	bar := &dbex.Position{Id: 400, Description: "test for delete"}
	err := DB.CreatePosition(bar)
	if err != nil {
		t.Error("FAILED: insert error: ", err, bar)
	}

	foo := &dbex.Position{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	err = DB.DeletePositionById(foo.Id)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Delete by id ", foo.Id, "\nerror: ", err)
	}

	foo = &dbex.Position{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// passed
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	} else {
		t.Error("FAILED: record not deleted")
	}
}

func TestUpdatePosition(t *testing.T) {

	bar := &dbex.Position{Description: "test for update"}
	err := DB.CreatePosition(bar)
	if err != nil {
		t.Error("insert error: ", err, bar)
	}

	foo := &dbex.Position{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	foo.Description = "updated"
	err = DB.UpdatePosition(foo)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Update ", foo.Id, "\nerror: ", err)
	}

	foo = &dbex.Position{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	} else if foo.Description != "updated" {
		t.Error("FAILED: description do not updated")
	}
}

func TestSelectAllPositions(t *testing.T) {

	// clear table
	if err := DB.DB.Exec("delete from positions").Error; err != nil {
		t.Error("Clear table error:", err)
	}

	dataItems := []TestDataItemPositions{
		{0, &dbex.Position{Description: "test1"}, false},
		{0, &dbex.Position{Description: "test2"}, false},
		{0, &dbex.Position{Description: "test3"}, false},
	}

	for _, item := range dataItems {
		err := DB.CreatePosition(item.input)

		if err != nil {
			t.Error("\nFAILED: non-expected error at Inserting ", item.input, "\nerror: ", err)
		}
	}

	foo, err := DB.SelectAllPositions()

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

func TestSelectPositionById(t *testing.T) {
	bar := &dbex.Position{Id: 500, Description: "test"}
	err := DB.CreatePosition(bar)
	if err != nil {
		t.Error("FAILED: insert error: ", err)
	}

	foo := &dbex.Position{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	if err := DB.DB.Exec("delete from positions").Error; err != nil {
		t.Error("Clear table error:", err)
	}
}

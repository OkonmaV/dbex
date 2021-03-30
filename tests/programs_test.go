package dbex_test

import (
	"dbex"
	"errors"
	"testing"

	"gorm.io/gorm"
)

type TestDataItemPrograms struct {
	inputId  uint
	input    *dbex.Program
	isBroken bool
}

func TestCreateProgram(t *testing.T) {

	dataItems := []TestDataItemPrograms{
		{200, &dbex.Program{Id: 200, Name: "test insert 1"}, false},
		{100, &dbex.Program{Id: 100, Name: "test insert 2"}, false},
		{200, &dbex.Program{Id: 200, Name: "test insert 3"}, true},
	}

	for _, item := range dataItems {
		err := DB.CreateProgram(item.input)

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

func TestDeleteProgramById(t *testing.T) {

	bar := &dbex.Program{Id: 400, Name: "test for delete"}
	err := DB.CreateProgram(bar)
	if err != nil {
		t.Error("FAILED: insert error: ", err, bar)
	}

	foo := &dbex.Program{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	err = DB.DeleteProgramById(foo.Id)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Delete by id ", foo.Id, "\nerror: ", err)
	}

	foo = &dbex.Program{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// passed
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	} else {
		t.Error("FAILED: record not deleted")
	}
}

func TestUpdateProgram(t *testing.T) {

	bar := &dbex.Program{Id: 600, Name: "test for update"}
	err := DB.CreateProgram(bar)
	if err != nil {
		t.Error("insert error: ", err, bar)
	}

	foo := &dbex.Program{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	foo.Name = "updated"
	err = DB.UpdateProgram(foo)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Update ", foo.Id, "\nerror: ", err)
	}

	foo = &dbex.Program{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	} else if foo.Name != "updated" {
		t.Error("FAILED: Name do not updated")
	}
}

func TestSelectAllPrograms(t *testing.T) {

	// clear table
	if err := DB.DB.Exec("delete from programs").Error; err != nil {
		t.Error("Clear table error:", err)
	}

	dataItems := []TestDataItemPrograms{
		{0, &dbex.Program{Name: "test1"}, false},
		{0, &dbex.Program{Name: "test2"}, false},
		{0, &dbex.Program{Name: "test3"}, false},
	}

	for _, item := range dataItems {
		err := DB.CreateProgram(item.input)

		if err != nil {
			t.Error("\nFAILED: non-expected error at Inserting ", item.input, "\nerror: ", err)
		}
	}

	foo, err := DB.SelectAllPrograms()

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

func TestSelectProgramById(t *testing.T) {
	bar := &dbex.Program{Id: 500, Name: "test"}
	err := DB.CreateProgram(bar)
	if err != nil {
		t.Error("FAILED: insert error: ", err)
	}

	foo := &dbex.Program{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	if err := DB.DB.Exec("delete from programs").Error; err != nil {
		t.Error("Clear table error:", err)
	}
}

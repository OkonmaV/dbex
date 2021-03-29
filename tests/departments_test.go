package dbex_test

import (
	"dbex"
	"errors"
	"testing"

	"gorm.io/gorm"
)

type TestDataItemDepartments struct {
	inputId  uint
	input    *dbex.Department
	isBroken bool
}

func TestCreateDepartment(t *testing.T) {

	dataItems := []TestDataItemDepartments{
		{200, &dbex.Department{Id: 200, Description: "test insert 2"}, false},
		{100, &dbex.Department{Id: 100, Description: "test insert 1"}, false},
		{200, &dbex.Department{Id: 200, Description: "test insert 2"}, true},
	}

	for _, item := range dataItems {
		err := DB.CreateDepartment(item.input)

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

func TestDeleteDepartmentById(t *testing.T) {

	depins := &dbex.Department{Description: "desc for delete"}
	err := DB.CreateDepartment(depins)
	if err != nil {
		t.Error("insert error: ", err, depins)
	}

	foo := &dbex.Department{}
	err = DB.DB.First(foo, depins.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	err = DB.DeleteDepartmentById(foo.Id)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Delete by id ", foo.Id, "\nerror: ", err)
	}

	foo = &dbex.Department{}
	err = DB.DB.First(foo, depins.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// passed
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	} else {
		t.Error("FAILED: record not deleted")
	}
}

func TestUpdateDepartment(t *testing.T) {

	depins := &dbex.Department{Description: "desc"}
	err := DB.CreateDepartment(depins)
	if err != nil {
		t.Error("insert error: ", err, depins)
	}

	foo := &dbex.Department{}
	err = DB.DB.First(foo, depins.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	foo.Description = "new desc"
	err = DB.UpdateDepartment(foo)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Update ", foo.Id, "\nerror: ", err)
	}

	foo = &dbex.Department{}
	err = DB.DB.First(foo, depins.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	} else if foo.Description != "new desc" {
		t.Error("FAILED: description do not updated")
	}
}

func TestSelectDepartmentsAll(t *testing.T) {

	// clear table
	if err := DB.DB.Exec("delete from departments").Error; err != nil {
		t.Error("Clear table error:", err)
	}

	dataItems := []TestDataItemDepartments{
		{0, &dbex.Department{Description: "test1"}, false},
		{0, &dbex.Department{Description: "test2"}, false},
		{0, &dbex.Department{Description: "test3"}, true},
	}

	for _, item := range dataItems {
		err := DB.CreateDepartment(item.input)

		if err != nil {
			t.Error("\nFAILED: non-expected error at Inserting ", item.input, "\nerror: ", err)
		}
	}

	foo, err := DB.SelectAllDepartments()

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

func TestSelectDepartmentById(t *testing.T) {
	depins := &dbex.Department{Description: "desc"}
	err := DB.CreateDepartment(depins)
	if err != nil {
		t.Error("insert error: ", err)
	}

	foo := &dbex.Department{}
	err = DB.DB.First(foo, depins.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}
}

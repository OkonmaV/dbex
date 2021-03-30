package dbex_test

import (
	"dbex"
	"errors"
	"testing"

	"gorm.io/gorm"
)

type TestDataItemTesters struct {
	inputId  uint
	input    *dbex.Tester
	isBroken bool
}

var fkTestersPositions dbex.Position = dbex.Position{Id: 550, Description: "fk"}
var fkTestersDepartments dbex.Department = dbex.Department{Id: 550, Description: "fk"}

func TestCreateTester(t *testing.T) {

	dataItems := []TestDataItemTesters{
		{100, &dbex.Tester{Id: 100, Name: "test1", PositionId: 550, Position: fkTestersPositions, DepartmentId: 550, Department: fkTestersDepartments}, false},
		{200, &dbex.Tester{Id: 300, Name: "test2"}, true},
		{300, &dbex.Tester{Id: 100, Name: "test3", PositionId: 550, Position: fkTestersPositions, DepartmentId: 550, Department: fkTestersDepartments}, true},
	}

	for _, item := range dataItems {
		err := DB.CreateTester(item.input)

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

func TestDeleteTesterById(t *testing.T) {

	bar := &dbex.Tester{Id: 400, Name: "test for delete", PositionId: 550, Position: fkTestersPositions, DepartmentId: 550, Department: fkTestersDepartments}
	err := DB.CreateTester(bar)
	if err != nil {
		t.Error("FAILED: insert error: ", err, bar)
	}
	foo := &dbex.Tester{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	err = DB.DeleteTesterById(foo.Id)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Delete by id ", foo.Id, "\nerror: ", err)
	}

	foo = &dbex.Tester{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// passed
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	} else {
		t.Error("FAILED: record not deleted")
	}
}

func TestUpdateTester(t *testing.T) {

	bar := &dbex.Tester{Name: "test for update", PositionId: 550, Position: fkTestersPositions, DepartmentId: 550, Department: fkTestersDepartments}

	// err := DB.DB.Create(&dbex.Testplan{Id: 750, Name: "fk"}).Error
	// if err != nil {
	// 	t.Error("\nFAILED: error at inserting for fk: ", err)
	// }

	err := DB.CreateTester(bar)
	if err != nil {
		t.Error("FAILED: insert error: ", err, bar)
	}
	foo := &dbex.Tester{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	foo.Name = "updated"
	err = DB.UpdateTester(foo)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Update ", foo.Id, "\nerror: ", err)
	}

	foo = &dbex.Tester{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	} else if foo.Name != "updated" {
		t.Error("FAILED: was not updated")
	}
}

func TestSelectAllTesters(t *testing.T) {

	// clear table
	if err := DB.DB.Exec("delete from testers").Error; err != nil {
		t.Error("Clear table error:", err)
	}

	dataItems := []TestDataItemTesters{
		{0, &dbex.Tester{Name: "test1", PositionId: 550, Position: fkTestersPositions, DepartmentId: 550, Department: fkTestersDepartments}, false},
		{0, &dbex.Tester{Name: "test2", PositionId: 550, Position: fkTestersPositions, DepartmentId: 550, Department: fkTestersDepartments}, false},
		{0, &dbex.Tester{Name: "test3", PositionId: 550, Position: fkTestersPositions, DepartmentId: 550, Department: fkTestersDepartments}, false},
	}

	for _, item := range dataItems {
		err := DB.CreateTester(item.input)

		if err != nil {
			t.Error("\nFAILED: non-expected error at Inserting ", item.input, "\nerror: ", err)
		}
	}

	foo, err := DB.SelectAllTesters()

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

func TestSelectTesterById(t *testing.T) {

	bar := &dbex.Tester{Id: 500, Name: "test", PositionId: 550, Position: fkTestersPositions, DepartmentId: 550, Department: fkTestersDepartments}
	err := DB.CreateTester(bar)
	if err != nil {
		t.Error("FAILED: insert error: ", err)
	}

	foo := &dbex.Tester{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	if err := DB.DB.Exec("delete from testers").Error; err != nil {
		t.Error("Clear table error:", err)
	}
}

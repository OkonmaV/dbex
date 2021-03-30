package dbex_test

import (
	"dbex"
	"errors"
	"testing"

	"gorm.io/gorm"
)

type TestDataItemTestpoints struct {
	inputId  uint
	input    *dbex.Testpoint
	isBroken bool
}

var fkTestpointsTesters dbex.Tester = dbex.Tester{Id: 550, Name: "fk", PositionId: 550, Position: dbex.Position{Id: 550, Description: "fk"}, DepartmentId: 550, Department: dbex.Department{Id: 550, Description: "fk"}}
var fkTestpointsTestcases dbex.Testcase = dbex.Testcase{Id: 550, Name: "fk", ForId: 550, Programversion: dbex.Programversion{Id: 550, Major: 550, Minor: 550, ProgramId: 550, Program: dbex.Program{Id: 550, Name: "fk"}}}
var fkTestpointsTestsets dbex.Testset = dbex.Testset{Id: 550, Name: "fk", TestPlanId: 550, Testplan: dbex.Testplan{Id: 550, Name: "fk"}}

func TestCreateTestpoint(t *testing.T) {

	dataItems := []TestDataItemTestpoints{
		{10, &dbex.Testpoint{Id: 100, TestEngineerId: 550, Tester: fkTestpointsTesters, TestCaseId: 550, Testcase: fkTestpointsTestcases, TestSetId: 550, Testset: fkTestpointsTestsets}, false},
		{200, &dbex.Testpoint{Id: 200}, true},
		{300, &dbex.Testpoint{Id: 300, TestEngineerId: 550, Tester: fkTestpointsTesters}, true},
		{400, &dbex.Testpoint{Id: 100, TestEngineerId: 550, Tester: fkTestpointsTesters, TestCaseId: 550, Testcase: fkTestpointsTestcases, TestSetId: 550, Testset: fkTestpointsTestsets}, true},
	}

	for _, item := range dataItems {
		err := DB.CreateTestpoint(item.input)

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

func TestDeleteTestpointById(t *testing.T) {

	bar := &dbex.Testpoint{TestEngineerId: 550, Tester: fkTestpointsTesters, TestCaseId: 550, Testcase: fkTestpointsTestcases, TestSetId: 550, Testset: fkTestpointsTestsets}
	err := DB.CreateTestpoint(bar)
	if err != nil {
		t.Error("FAILED: insert error: ", err, bar)
	}
	foo := &dbex.Testpoint{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	err = DB.DeleteTestpointById(foo.Id)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Delete by id ", foo.Id, "\nerror: ", err)
	}

	foo = &dbex.Testpoint{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// passed
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	} else {
		t.Error("FAILED: record not deleted")
	}
}

func TestUpdateTestpoint(t *testing.T) {

	bar := &dbex.Testpoint{TestEngineerId: 550, Tester: fkTestpointsTesters, TestCaseId: 550, Testcase: fkTestpointsTestcases, TestSetId: 550, Testset: fkTestpointsTestsets}

	err := DB.CreateTestpoint(bar)
	if err != nil {
		t.Error("FAILED: insert error: ", err, bar)
	}
	foo := &dbex.Testpoint{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	foo.TestSetId = 551
	foo.Testset = dbex.Testset{Id: 551, Name: "fk", TestPlanId: 550, Testplan: dbex.Testplan{Id: 550, Name: "fk"}}

	err = DB.UpdateTestpoint(foo)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Update ", foo.Id, "\nerror: ", err)
	}

	foo = &dbex.Testpoint{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	} else if foo.TestSetId != 551 {
		t.Error("FAILED: was not updated")
	}
}

func TestSelectAllTestpoints(t *testing.T) {

	// clear table
	if err := DB.DB.Exec("delete from testpoints").Error; err != nil {
		t.Error("FAILED: Clear table error:", err)
	}

	dataItems := []TestDataItemTestpoints{
		{0, &dbex.Testpoint{TestEngineerId: 550, Tester: fkTestpointsTesters, TestCaseId: 550, Testcase: fkTestpointsTestcases, TestSetId: 550, Testset: fkTestpointsTestsets}, false},
		{0, &dbex.Testpoint{TestEngineerId: 550, Tester: fkTestpointsTesters, TestCaseId: 550, Testcase: fkTestpointsTestcases, TestSetId: 550, Testset: fkTestpointsTestsets}, false},
		{0, &dbex.Testpoint{TestEngineerId: 550, Tester: fkTestpointsTesters, TestCaseId: 550, Testcase: fkTestpointsTestcases, TestSetId: 550, Testset: fkTestpointsTestsets}, false},
	}

	for _, item := range dataItems {
		err := DB.CreateTestpoint(item.input)

		if err != nil {
			t.Error("\nFAILED: non-expected error at Inserting ", item.input, "\nerror: ", err)
		}
	}

	foo, err := DB.SelectAllTestpoints()

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

func TestSelectTestpointById(t *testing.T) {

	bar := &dbex.Testpoint{Id: 500, TestEngineerId: 550, Tester: fkTestpointsTesters, TestCaseId: 550, Testcase: fkTestpointsTestcases, TestSetId: 550, Testset: fkTestpointsTestsets}
	err := DB.CreateTestpoint(bar)
	if err != nil {
		t.Error("FAILED: insert error: ", err)
	}

	foo := &dbex.Testpoint{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	if err := DB.DB.Exec("delete from testpoints").Error; err != nil {
		t.Error("FAILED:Clear table error:", err)
	}
}

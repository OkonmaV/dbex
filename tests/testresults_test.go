package dbex_test

import (
	"dbex"
	"errors"
	"testing"
	"time"

	"gorm.io/gorm"
)

type TestDataItemTestresults struct {
	inputId  uint
	input    *dbex.Testresult
	isBroken bool
}

var fkTestresultsTestpoints dbex.Testpoint = dbex.Testpoint{Id: 550, TestEngineerId: 550, Tester: fkTestpointsTesters, TestCaseId: 550, Testcase: fkTestpointsTestcases, TestSetId: 550, Testset: fkTestpointsTestsets}
var fkTestresultsStatuses dbex.Status = dbex.Status{Id: 550, Status: "fk"}

func TestCreateTestresult(t *testing.T) {

	dataItems := []TestDataItemTestresults{
		{100, &dbex.Testresult{Id: 100, TestPointId: 550, Testpoint: fkTestresultsTestpoints, StatusId: 550, Status: fkTestresultsStatuses, Start: time.Now(), Finish: time.Now()}, false},
		{200, &dbex.Testresult{Id: 200, TestPointId: 550, Testpoint: fkTestresultsTestpoints, StatusId: 550, Status: fkTestresultsStatuses}, true},
		{300, &dbex.Testresult{Id: 100, TestPointId: 550, Testpoint: fkTestresultsTestpoints, StatusId: 550, Status: fkTestresultsStatuses, Start: time.Now(), Finish: time.Now()}, true},
		{300, &dbex.Testresult{Id: 400, Start: time.Now(), Finish: time.Now()}, true},
	}

	for _, item := range dataItems {
		err := DB.CreateTestresult(item.input)

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

	if err := DB.DB.Exec("delete from testresults").Error; err != nil {
		t.Error("Clear table error:", err)
	}

}

func TestDeleteTestresultById(t *testing.T) {

	bar := &dbex.Testresult{TestPointId: 550, Testpoint: fkTestresultsTestpoints, StatusId: 550, Status: fkTestresultsStatuses, Start: time.Now(), Finish: time.Now()}
	err := DB.CreateTestresult(bar)
	if err != nil {
		t.Error("FAILED: insert error: ", err, bar)
	}
	foo := &dbex.Testresult{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	err = DB.DeleteTestresultById(foo.Id)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Delete by id ", foo.Id, "\nerror: ", err)
	}

	foo = &dbex.Testresult{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// passed
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	} else {
		t.Error("FAILED: record not deleted")
	}

	if err := DB.DB.Exec("delete from testresults").Error; err != nil {
		t.Error("Clear table error:", err)
	}
}

func TestUpdateTestresult(t *testing.T) {

	bar := &dbex.Testresult{TestPointId: 550, Testpoint: fkTestresultsTestpoints, StatusId: 550, Status: fkTestresultsStatuses, Start: time.Now(), Finish: time.Now()}

	err := DB.CreateTestresult(bar)
	if err != nil {
		t.Error("FAILED: insert error: ", err, bar)
	}
	foo := &dbex.Testresult{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	foo.StatusId = 551
	foo.Status = dbex.Status{Id: 551, Status: "fk"}

	err = DB.UpdateTestresult(foo)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Update ", foo.Id, "\nerror: ", err)
	}

	foo = &dbex.Testresult{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	} else if foo.StatusId != 551 {
		t.Error("FAILED: was not updated")
	}

	if err := DB.DB.Exec("delete from testresults").Error; err != nil {
		t.Error("Clear table error:", err)
	}
}

func TestSelectAllTestresults(t *testing.T) {

	// clear table
	if err := DB.DB.Exec("delete from testresults").Error; err != nil {
		t.Error("Clear table error:", err)
	}

	dataItems := []TestDataItemTestresults{
		{0, &dbex.Testresult{TestPointId: 550, Testpoint: fkTestresultsTestpoints, StatusId: 550, Status: fkTestresultsStatuses, Start: time.Now(), Finish: time.Now()}, false},
		{0, &dbex.Testresult{TestPointId: 550, Testpoint: fkTestresultsTestpoints, StatusId: 550, Status: fkTestresultsStatuses, Start: time.Now(), Finish: time.Now()}, false},
		{0, &dbex.Testresult{TestPointId: 550, Testpoint: fkTestresultsTestpoints, StatusId: 550, Status: fkTestresultsStatuses, Start: time.Now(), Finish: time.Now()}, false},
	}

	for _, item := range dataItems {
		err := DB.CreateTestresult(item.input)

		if err != nil {
			t.Error("\nFAILED: non-expected error at Inserting ", item.input, "\nerror: ", err)
		}
	}

	foo, err := DB.SelectAllTestresults()

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

	if err := DB.DB.Exec("delete from testresults").Error; err != nil {
		t.Error("Clear table error:", err)
	}

}

func TestSelectTestresultById(t *testing.T) {

	bar := &dbex.Testresult{Id: 500, TestPointId: 550, Testpoint: fkTestresultsTestpoints, StatusId: 550, Status: fkTestresultsStatuses, Start: time.Now(), Finish: time.Now()}
	err := DB.CreateTestresult(bar)
	if err != nil {
		t.Error("FAILED: insert error: ", err)
	}

	foo := &dbex.Testresult{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	if err := DB.DB.Exec("delete from testresults").Error; err != nil {
		t.Error("Clear table error:", err)
	}

	if err := DB.DB.Exec("delete from testpoints").Error; err != nil {
		t.Error("Clear table error:", err)
	}
}

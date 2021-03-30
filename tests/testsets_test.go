package dbex_test

import (
	"dbex"
	"errors"
	"testing"

	"gorm.io/gorm"
)

type TestDataItemTestsets struct {
	inputId  uint
	input    *dbex.Testset
	isBroken bool
}

func TestCreateTestset(t *testing.T) {

	dataItems := []TestDataItemTestsets{
		{200, &dbex.Testset{Id: 200, Name: "test1", TestPlanId: 600}, false},
		{100, &dbex.Testset{Id: 100, Name: "test2", TestPlanId: 1000}, true},
		{200, &dbex.Testset{Id: 200, Name: "test3"}, true},
	}

	err := DB.CreateTestplan(&dbex.Testplan{Id: 600, Name: "fk"})
	if err != nil {
		t.Error("\nFAILED: error at inserting for fk: ", err)
	}

	for _, item := range dataItems {
		err := DB.CreateTestset(item.input)

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

func TestDeleteTestsetById(t *testing.T) {

	bar := &dbex.Testset{Name: "test for delete"}
	err := DB.CreateTestset(bar)
	if err != nil {
		t.Error("FAILED: insert error: ", err, bar)
	}
	foo := &dbex.Testset{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	err = DB.DeleteTestsetById(foo.Id)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Delete by id ", foo.Id, "\nerror: ", err)
	}

	foo = &dbex.Testset{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// passed
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	} else {
		t.Error("FAILED: record not deleted")
	}
}

func TestUpdateTestset(t *testing.T) {

	bar := &dbex.Testset{Name: "test for update", TestPlanId: 700}

	err := DB.CreateTestplan(&dbex.Testplan{Id: 700, Name: "fk"})
	if err != nil {
		t.Error("\nFAILED: error at inserting for fk: ", err)
	}

	err = DB.CreateTestset(bar)
	if err != nil {
		t.Error("FAILED: insert error: ", err, bar)
	}
	foo := &dbex.Testset{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	foo.Name = "updated"
	err = DB.UpdateTestset(foo)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Update ", foo.Id, "\nerror: ", err)
	}

	foo = &dbex.Testset{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	} else if foo.Name != "updated" {
		t.Error("FAILED: was not updated")
	}
}

func TestSelectAllTestsets(t *testing.T) {

	// clear table
	if err := DB.DB.Exec("delete from testsets").Error; err != nil {
		t.Error("Clear table error:", err)
	}

	dataItems := []TestDataItemTestsets{
		{0, &dbex.Testset{Name: "test1"}, false},
		{0, &dbex.Testset{Name: "test2"}, false},
		{0, &dbex.Testset{Name: "test3"}, false},
	}

	for _, item := range dataItems {
		err := DB.CreateTestset(item.input)

		if err != nil {
			t.Error("\nFAILED: non-expected error at Inserting ", item.input, "\nerror: ", err)
		}
	}

	foo, err := DB.SelectAllTestsets()

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

func TestSelectTestsetById(t *testing.T) {

	bar := &dbex.Testset{Id: 500, Name: "test"}
	err := DB.CreateTestset(bar)
	if err != nil {
		t.Error("FAILED: insert error: ", err)
	}

	foo := &dbex.Testset{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}
}

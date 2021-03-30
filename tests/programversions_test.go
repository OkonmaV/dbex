package dbex_test

import (
	"dbex"
	"errors"
	"testing"

	"gorm.io/gorm"
)

type TestDataItemProgramversions struct {
	inputId  uint
	input    *dbex.Programversion
	isBroken bool
}

var fkProgramversions dbex.Program = dbex.Program{Id: 550, Name: "fk"}

func TestCreateProgramversion(t *testing.T) {

	dataItems := []TestDataItemProgramversions{
		{100, &dbex.Programversion{Id: 100, Major: 10, Minor: 10, ProgramId: 550, Program: fkProgramversions}, false},
		{200, &dbex.Programversion{Id: 200, Major: 20, Minor: 20}, true},
		{300, &dbex.Programversion{Id: 300, Major: 30, Minor: 30, ProgramId: 1000}, true},
		{400, &dbex.Programversion{Id: 100, Major: 10, Minor: 10, ProgramId: 550, Program: fkProgramversions}, true},
	}

	// err := DB.DB.Create(&dbex.Testplan{Id: 650, Name: "fk"}).Error
	// if err != nil {
	// 	t.Error("\nFAILED: error at inserting for fk: ", err)
	// }

	for _, item := range dataItems {
		err := DB.CreateProgramversion(item.input)

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

func TestDeleteProgramversionById(t *testing.T) {

	// err := DB.DB.Create(&dbex.Testplan{Id: 550, Name: "fk"}).Error
	// if err != nil {
	// 	t.Error("\nFAILED: error at inserting for fk: ", err)
	// }

	bar := &dbex.Programversion{Major: 10, Minor: 10, ProgramId: 550, Program: fkProgramversions}
	err := DB.CreateProgramversion(bar)
	if err != nil {
		t.Error("FAILED: insert error: ", err, bar)
	}
	foo := &dbex.Programversion{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	err = DB.DeleteProgramversionById(foo.Id)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Delete by id ", foo.Id, "\nerror: ", err)
	}

	foo = &dbex.Programversion{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// passed
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	} else {
		t.Error("FAILED: record not deleted")
	}
}

func TestUpdateProgramversion(t *testing.T) {

	bar := &dbex.Programversion{Major: 10, Minor: 10, ProgramId: 550, Program: fkProgramversions}

	// err := DB.DB.Create(&dbex.Testplan{Id: 750, Name: "fk"}).Error
	// if err != nil {
	// 	t.Error("\nFAILED: error at inserting for fk: ", err)
	// }

	err := DB.CreateProgramversion(bar)
	if err != nil {
		t.Error("FAILED: insert error: ", err, bar)
	}
	foo := &dbex.Programversion{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	foo.Major = 20
	err = DB.UpdateProgramversion(foo)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Update ", foo.Id, "\nerror: ", err)
	}

	foo = &dbex.Programversion{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	} else if foo.Major != 20 {
		t.Error("FAILED: was not updated")
	}
}

func TestSelectAllProgramversions(t *testing.T) {

	// clear table
	if err := DB.DB.Exec("delete from programversions").Error; err != nil {
		t.Error("Clear table error:", err)
	}

	// err := DB.DB.Create(&dbex.Testplan{Id: 850, Name: "fk"}).Error
	// if err != nil {
	// 	t.Error("\nFAILED: error at inserting for fk: ", err)
	// }

	dataItems := []TestDataItemProgramversions{
		{0, &dbex.Programversion{Major: 15, Minor: 15, ProgramId: 550, Program: fkProgramversions}, false},
		{0, &dbex.Programversion{Major: 25, Minor: 25, ProgramId: 550, Program: fkProgramversions}, false},
		{0, &dbex.Programversion{Major: 35, Minor: 35, ProgramId: 550, Program: fkProgramversions}, false},
	}

	for _, item := range dataItems {
		err := DB.CreateProgramversion(item.input)

		if err != nil {
			t.Error("\nFAILED: non-expected error at Inserting ", item.input, "\nerror: ", err)
		}
	}

	foo, err := DB.SelectAllProgramversions()

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

func TestSelectProgramversionById(t *testing.T) {

	bar := &dbex.Programversion{Id: 500, Major: 10, Minor: 10, ProgramId: 550, Program: fkProgramversions}
	err := DB.CreateProgramversion(bar)
	if err != nil {
		t.Error("FAILED: insert error: ", err)
	}

	foo := &dbex.Programversion{}
	err = DB.DB.First(foo, bar.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("FAILED: not found record")
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	if err := DB.DB.Exec("delete from programversions").Error; err != nil {
		t.Error("Clear table error:", err)
	}
}

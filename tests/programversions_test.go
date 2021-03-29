package dbex_test

import (
	"dbex"
	"testing"
)

type TestDataItemProgramversion struct {
	inputId  uint
	input    *dbex.Programversion
	isBroken bool
}

func TestCreateProgramversion(t *testing.T) {

	dataItems := []TestDataItemProgramversion{
		{0, &dbex.Programversion{Id: 1, Major: 1, Minor: 1, ProgramId: 5}, true},
		{0, &dbex.Programversion{Major: 1, Minor: 1, ProgramId: 5}, false},
		{0, &dbex.Programversion{Major: 1, Minor: 1, ProgramId: 9000000}, true},
		{0, &dbex.Programversion{Major: 1, ProgramId: 5}, false},
		{0, &dbex.Programversion{Minor: 1, ProgramId: 5}, false},
	}

	conn, err := GetConf()
	if err != nil {
		t.Error("DB connection error: ", err)
		return
	}

	for _, item := range dataItems {
		err := conn.CreateProgramversion(item.input)

		if item.isBroken {
			if err == nil {
				t.Error("\nFAILED: expected an error, but no error catched at Inserting ", item.input)
			} else {
				t.Log("\nPASSED: expected an error, got an error at Inserting ", item.input, "\nerror: ", err)
			}
		} else {
			if err != nil {
				t.Error("\nFAILED: non-expected error at Inserting ", item.input, "\nerror: ", err)
			} else {
				t.Log("\nPASSED: no error at Inserting ", item.input)
			}
		}
	}

}

func TestDeleteProgramversionById(t *testing.T) {

	conn, err := GetConf()
	if err != nil {
		t.Error("FAILED: DB connection error: ", err)
		return
	}

	foo := &dbex.Programversion{}
	err = conn.DB.Where("ProgramId = ?", 5).Last(foo).Error
	if foo.Id == 0 && err == nil {
		_ = conn.DB.Create(&dbex.Programversion{Major: 1, Minor: 1, ProgramId: 1})
		_ = conn.DB.Where("ProgramId = ?", 5).Last(foo)
	} else if err != nil {
		t.Error("FAILED: some error : ", err)
	}

	err = conn.DeleteProgramversionById(foo.Id)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Delete by id ", foo.Id, "\nerror: ", err)
	} else {
		t.Log("\nPASSED: no error at Delete by id ", foo.Id)
	}
}

func TestUpdateProgramversion(t *testing.T) {

	conn, err := GetConf()
	if err != nil {
		t.Error("FAILED: DB connection error: ", err)
		return
	}

	foo := &dbex.Programversion{}
	err = conn.DB.Where("ProgramId = ?", 5).Last(foo).Error
	if foo.Id == 0 && err == nil {
		_ = conn.DB.Create(&dbex.Programversion{Major: 1, Minor: 1, ProgramId: 1})
		_ = conn.DB.Where("ProgramId = ?", 5).Last(foo)
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	dataItems := []*TestDataItemProgramversion{
		{0, &dbex.Programversion{Id: foo.Id, Major: 1, Minor: 1, ProgramId: 5}, false},
		{0, &dbex.Programversion{Id: foo.Id, Major: 1, Minor: 1, ProgramId: 0}, true},
	}

	for _, item := range dataItems {
		err := conn.UpdateProgramversion(item.input)

		if item.isBroken {
			if err == nil {
				t.Error("\nFAILED: expected an error, but no error catched at Update ", item.input)
			} else {
				t.Log("\nPASSED: expected an error, got an error at Update ", item.input, "\nerror: ", err)
			}
		} else {
			if err != nil {
				t.Error("\nFAILED: non-expected error at Update ", item.input, "\nerror: ", err)
			} else {
				t.Log("\nPASSED: no error at Update ", item.input)
			}
		}
	}
}

func TestSelectProgramversionsAll(t *testing.T) {
	conn, err := GetConf()
	if err != nil {
		t.Error("FAILED: DB connection error: ", err)
		return
	}
	foo, err := conn.SelectProgramversionsAll()

	if err != nil {
		t.Error("\nFAILED: non-expected error at SelectAll\nerror: ", err)
	} else if len(*foo) == 0 {
		t.Error("\nFAILED: returned empty slice at SelectAll")
	} else {
		t.Log("\nPASSED: returned at SelectAll:\n", foo)
	}

}

func TestSelectProgramversionById(t *testing.T) {
	dataItems := []TestDataItemStatuses{
		{0, nil, true},
		{1, nil, false},
	}

	conn, err := GetConf()
	if err != nil {
		t.Error("DB connection error: ", err)
		return
	}

	for _, item := range dataItems {
		foo, err := conn.SelectProgramversionById(item.inputId)

		if item.isBroken {
			if err == nil {
				t.Error("\nFAILED: expected an error, but no error catched at Select by id ", item.inputId)
			} else {
				t.Log("\nPASSED: expected an error, got an error at Select by id ", foo, "\nerror: ", err)
			}
		} else {
			if err != nil {
				t.Error("\nFAILED: non-expected error at Select by id ", item.inputId, "\nerror: ", err)
			} else if foo.Id != 0 {
				t.Log("\nPASSED: no error at Select by id ", foo)
			} else {
				t.Error("\nFAILED: no error, but relust is nil at Select by id ", item.inputId)
			}
		}
	}

}

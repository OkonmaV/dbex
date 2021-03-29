package dbex_test

import (
	"dbex"
	"testing"
	"time"
)

type TestDataItemTestresult struct {
	inputId  uint
	input    *dbex.Testresult
	isBroken bool
}

func TestCreateTestresult(t *testing.T) {

	dataItems := []TestDataItemTestresult{
		{0, &dbex.Testresult{Id: 1, TestPointId: 1, StatusId: 4}, true},
		{0, &dbex.Testresult{TestPointId: 26, StatusId: 4, Start: time.Now(), Finish: time.Now()}, false},
		{0, &dbex.Testresult{TestPointId: 0, StatusId: 0}, true},
	}

	conn, err := GetConf()
	if err != nil {
		t.Error("DB connection error: ", err)
		return
	}

	for _, item := range dataItems {
		err := conn.CreateTestresult(item.input)

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

func TestDeleteTestresultById(t *testing.T) {

	conn, err := GetConf()
	if err != nil {
		t.Error("FAILED: DB connection error: ", err)
		return
	}

	foo := &dbex.Testresult{}
	err = conn.DB.Where("StatusId = ?", 4).Last(foo).Error
	if foo.Id == 0 && err == nil {
		_ = conn.DB.Create(&dbex.Testresult{StatusId: 4, TestPointId: 1, Start: time.Now(), Finish: time.Now()})
		_ = conn.DB.Where("StatusId = ?", 4).Last(foo)
	} else if err != nil {
		t.Error("FAILED: some error : ", err)
	}

	err = conn.DeleteTestresultById(foo.Id)

	if err != nil {
		t.Error("\nFAILED: non-expected error at Delete by id ", foo.Id, "\nerror: ", err)
	} else {
		t.Log("\nPASSED: no error at Delete by id ", foo.Id)
	}
}

func TestUpdateTestresult(t *testing.T) {

	conn, err := GetConf()
	if err != nil {
		t.Error("FAILED: DB connection error: ", err)
		return
	}

	foo := &dbex.Testresult{}
	err = conn.DB.Where("StatusId = ?", 4).Last(foo).Error
	if foo.Id == 0 {
		_ = conn.DB.Create(&dbex.Testresult{StatusId: 4, TestPointId: 1, Start: time.Now(), Finish: time.Now()})
		_ = conn.DB.Where("StatusId = ?", 4).Last(foo)
	} else if err != nil {
		t.Error("FAILED: some error : ", err, "\nfoo: ", foo)
	}

	dataItems := []*TestDataItemTestresult{
		{0, &dbex.Testresult{Id: foo.Id, StatusId: 4, TestPointId: 1, Start: time.Now(), Finish: time.Now()}, false},
		{0, &dbex.Testresult{Id: foo.Id, StatusId: 4, TestPointId: 1}, true},
	}

	for _, item := range dataItems {
		err := conn.UpdateTestresult(item.input)

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

func TestSelectTestresultsAll(t *testing.T) {
	conn, err := GetConf()
	if err != nil {
		t.Error("FAILED: DB connection error: ", err)
		return
	}
	foo, err := conn.SelectTestresultsAll()

	if err != nil {
		t.Error("\nFAILED: non-expected error at SelectAll\nerror: ", err)
	} else if len(*foo) == 0 {
		t.Error("\nFAILED: returned empty slice at SelectAll")
	} else {
		t.Log("\nPASSED: returned at SelectAll:\n", foo)
	}

}

func TestSelectTestresultById(t *testing.T) {
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
		foo, err := conn.SelectTestresultById(item.inputId)

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

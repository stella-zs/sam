package dbmsasg

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestSet(t *testing.T) {
	testcases := []struct {
		desc     string
		c        Car
		expected bool
	}{
		{"setting valid data", Car{1, "audi", "r8", "petrol"}, true},
		//{"setting valid data", car{2, "benz", "mrceries", "diesel"}, true},
		//{"setting valid data", car{3, "tata", "nano", "petrol"}, true},
		//{"setting invalid data", Car{1, "", "", ""}, true},
		//{"setting invalid data", Car{-1, "", "", ""}, true},
		//{"not setting any data", Car{}, true},
	}
	var s Store
	//var err error
	//e := driver.Mysql_config{"stella", "125", "localhost", "3306", "sam"}
	db, mock, err := sqlmock.New()
	s.Db = db
	//fmt.Println("helooooooooooooooooooo")
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer db.Close()
	for _, val := range testcases {
		mock.ExpectExec("insert into Car values").
			WithArgs(val.c.Id, val.c.Name, val.c.Model, val.c.Enginetype).
			WillReturnResult(sqlmock.NewResult(1, 1)).
			WillReturnError(err)
		output := s.Set(val.c)
		if output != val.expected {
			t.Errorf("testcasefailed got %v expected %v", output, val.expected)
		}
	}
}
func TestGet(t *testing.T) {
	testcases := []struct {
		desc     string
		id       int
		expected Car
	}{
		//{"when valid id is given", 1, car{1, "audi", "r8", "petrol"}},
		//{"when valid id is given", 2, car{2, "benz", "merceries", "diesel"}},
		{"when valid id is given", 1, Car{1, "audi", "r8", "petrol"}},
		{"when non-valid id is given", -3, Car{0, "", "", ""}},
		{"when not given anything", 0, Car{0, "", "", ""}},
	}
	var s Store
	//var err error
	//e := driver.Mysql_config{"stella", "125", "localhost", "3306", "sam"}
	db, mock, err := sqlmock.New()
	s.Db = db
	//fmt.Println("helooooooooooooooooooo")
	if err != nil {
		t.Fatalf("%v", err)

	}
	defer db.Close()
	for i, val := range testcases {
		rows := sqlmock.NewRows([]string{"Id", "Name", "Model", "EngineType"}).
			AddRow(val.expected.Id, val.expected.Name, val.expected.Model, val.expected.Enginetype)
		mock.ExpectQuery("select (.+) from Car where id=?").
			WithArgs(val.id).WillReturnRows(rows).WillReturnError(err)
		output := s.Get(val.id)
		if output != val.expected {
			t.Errorf("testcase failed %v got %v expected %v", i, output, val.expected)
		}
	}
}
func TestDelete(t *testing.T) {
	testcases := []struct {
		desc     string
		id       int
		expected bool
	}{
		{"deleting from table", 1, true},
		//{"deleting non existing from table", 5, false},
		//{"deleting without any id specified", 0, false},
	}
	var s Store
	//var err error
	//e := driver.Mysql_config{"stella", "125", "localhost", "3306", "sam"}
	db, mock, err := sqlmock.New()
	s.Db = db
	//fmt.Println("helooooooooooooooooooo")
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer db.Close()

	for _, val := range testcases {
		mock.ExpectExec("delete from Car where id=?").
			WithArgs(val.id).WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(err)
		output := s.Delete(val.id)
		if output != val.expected {
			t.Errorf("testcasefailed expected %v got %v", val.expected, output)
		}
	}
}

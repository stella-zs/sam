package main

import (
	"dbms/dbmsasg"
	"dbms/driver"
	"fmt"
	_ "fmt"
)

func main() {
	var s dbmsasg.Store
	a := driver.Mysql_config{"stella", "125", "localhost", "3306", "sam"}
	s.Db, s.Err = driver.Connection(a)
	if s.Err != nil {
		panic(s.Err)
	}
	var c dbmsasg.Car
	c = dbmsasg.Car{1, "audi", "r8", "petrol"}
	setop := s.Set(c)
	fmt.Println(setop)
	getop := s.Get(c.Id)
	fmt.Println(getop)
	delop := s.Delete(c.Id)
	fmt.Println(delop)

}

/*
		insert, err1 := db.Exec("insert into sample values('ef',46,'ef@gmail.com')")
		insert2, err5 := db.Exec("insert into sample values('stellar',545,'sr@gmail.com')")
		if err1 != nil {
			panic(err1)
		}
		fmt.Println("inserted", insert)
		if err5 != nil {
			panic(err5)
		}
		fmt.Println("inserted1", insert2)
		rows, err2 := db.Query("SELECT * FROM sam.sample")
		if err != nil {
			panic(err.Error())
		}
		type emp struct {
			Name  string
			ID    int
			Email string
		}

		for rows.Next() {
			var e emp
			err2 = rows.Scan(&e.Name, &e.ID, &e.Email)
			if err2 != nil {
				panic(err2.Error())
			}
			//fmt.Printf("%d %s %s\n", e.Name, e.ID, e.Email)
			fmt.Println(e.Name, e.ID, e.Email)
		}
		altering, err := db.Query("ALTER TABLE sample MODIFY COLUMN name varchar(30)")
		if err != nil {
			panic(err.Error())
		}
		//fmt.Printf("%d %s %s\n", e.Name, e.ID, e.Email)
		fmt.Println("altered", altering)
		del, err6 := db.Query("delete from sample where name='stellar'")
		if err6 != nil {
			panic(err6)
		}
		fmt.Println("deleted", del)

	}

*/

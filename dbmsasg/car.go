package dbmsasg

import (
	"database/sql"
	"fmt"
)

type Store struct {
	Db  *sql.DB
	Err error
}

type Car struct {
	Id         int
	Name       string
	Model      string
	Enginetype string
}

func (s Store) Get(id int) (c Car) {
	//var c Car
	if id > 0 {
		//a := fmt.Sprintf("select * from car where id=%v", id)
		se, err := s.Db.Query("select * from car where id=?", id)
		if err != nil {
			fmt.Errorf("%v", err)
		}
		se.Next()
		se.Scan(&c.Id, &c.Name, &c.Model, &c.Enginetype)
		se.Close()
		return c
	}
	c.Id = 0
	c.Name = ""
	c.Model = ""
	c.Enginetype = ""
	return c
}
func (s Store) Set(c Car) bool {
	if c.Id > 0 {
		//x := fmt.Sprintf("insert into car values %v,%v,%v,%v", c.id, c.name, c.model, c.enginetype)
		res, err := s.Db.Exec("insert ignore into car values(?,?,?,?)", c.Id, c.Name, c.Model, c.Enginetype)
		rows, err := res.RowsAffected()
		//fmt.Println("welcomekfjefjek")
		if err != nil {
			fmt.Printf("%v", err)
		}
		if rows == 0 {

			return false
		}
		return true
	}
	return false
}

func (s Store) Delete(id int) bool {
	//var s Store
	//d := fmt.Sprintf("DELETE FROM car WHERE id=%v", id)
	del, err := s.Db.Exec("DELETE FROM car WHERE id=?", id)
	if err != nil {
		fmt.Printf("%v", err)
	}

	x, err := del.RowsAffected()
	if x == 0 {
		return false
	}

	return true
}

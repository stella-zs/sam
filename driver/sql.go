package driver

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Mysql_config struct {
	User     string
	Password string
	Host     string
	Port     string
	Db_name  string
}

func Connection(a Mysql_config) (*sql.DB, error) {
	con_string := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", a.User, a.Password, a.Host, a.Port, a.Db_name)
	//sql.Open("my_sql")
	db, err := sql.Open("mysql", con_string)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	db.Exec("create table if not exists car(id int primary key,name varchar(255),model varchar(255),enginetype varchar(255));")
	//return db, err
	//defer db.Close()
	return db, err
}

/*
func Connstr(config Mysql_config) string {
	connStr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", config.User, config.Password, config.Host, config.Port, config.Db_name)
	return connStr
}
func Connection() (*sql.DB, error) {
	a := Mysql_config{"stella", "125", "localhost", "3306", "sam"}
	db, err := sql.Open("mysql", Connstr(a))

	return db, err
}
*/

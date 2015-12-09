package models
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
//mysql
	mysql_user = ""
	mysql_pass = ""
	mysql_host = ""

	mysql_port = ""
	database_name = ""
)

var (
	db  *sqlx.DB
	err error
)

func reconnect() {
	if db != nil && db.Ping() == nil {
		return
	}
	db, err = sqlx.Open("mysql", mysql_user + ":" + mysql_pass + "@tcp(" + mysql_host + ":" + mysql_port + ")/" + database_name)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)

}
func GetConn() *sqlx.DB {
	reconnect()
	return db
}

type Response struct {
	Status int `json:"status"`
	Msg    string `json:"msg"`
	Data   interface{} `json:data`
}
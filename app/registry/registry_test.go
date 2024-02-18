package registry

import (
	"os"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	conf = &mysql.Config{
		Addr:   "127.0.0.1:3306",
		User:   "root",
		Passwd: "root",
		DBName: "mydb",
	}
	db = sqlx.MustConnect("mysql", conf.FormatDSN())
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

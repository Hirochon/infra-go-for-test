package mysqlconnection

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLClient(mysqlUser, mysqlPassword, mysqlHost, mysqlDatabase, mysqlExtraProperties string) (*sql.DB, error) {
	// %2FはURLエンコードされた"/"を表す。また%をエスケープするために%%となっている。
	mysqlClient, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local%s", mysqlUser, mysqlPassword, mysqlHost, mysqlDatabase, mysqlExtraProperties))
	if err != nil {
		fmt.Println(err)
	}
	// https://tutuz-tech.hatenablog.com/entry/2020/03/24/170159
	mysqlClient.SetMaxOpenConns(25)
	mysqlClient.SetMaxIdleConns(25)
	mysqlClient.SetConnMaxLifetime(5 * time.Minute)
	return mysqlClient, nil
}

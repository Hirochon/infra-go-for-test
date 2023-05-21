package main

import (
	"os"

	"github.com/Hirochon/infra-go-for-test/internal/infrastructure/externalconnection/mysqlconnection"
)

func main() {
	mysqlClient, err := mysqlconnection.NewMySQLClient(os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DATABASE"), os.Getenv("MYSQL_EXTRA_PROPERTIES"))
	if err != nil {
		panic(err)
	}
	defer mysqlClient.Close()
}

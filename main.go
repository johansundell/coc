// coc project main.go
package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var mysqlUser, mysqlPass, mysqlDb, mysqlHost string
var urlClan = "https://api.clashofclans.com/v1/clans/%s"
var urlMembers = "https://api.clashofclans.com/v1/clans/%s/members"
var myKey, myClanTag string
var basePath string

func init() {
	myKey = os.Getenv("COC_KEY")
	myClanTag = os.Getenv("COC_CLANTAG")

	mysqlDb = "cocsniffer"
	mysqlHost = os.Getenv("MYSQL_COC_HOST")
	mysqlUser = os.Getenv("MYSQL_USER")
	mysqlPass = os.Getenv("MYSQL_PASS")

	basePath = os.Getenv("WWW_BASE_PATH")

}

func main() {
	db, _ = sql.Open("mysql", mysqlUser+":"+mysqlPass+"@tcp("+mysqlHost+":3306)/"+mysqlDb)
	defer db.Close()

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

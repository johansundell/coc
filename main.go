// coc project main.go
package main

import (
	"log"
	"net/http"
	"os"
)

var urlClan = "https://api.clashofclans.com/v1/clans/%s"
var urlMembers = "https://api.clashofclans.com/v1/clans/%s/members"
var myKey, myClanTag string

func init() {
	myKey = os.Getenv("COC_KEY")
	myClanTag = os.Getenv("COC_CLANTAG")
}

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

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
	/*clan, err := getClanInfo(myClanTag)
	if err != nil {
		panic(err)
	}
	fmt.Println(clan.Name, clan.BadgeUrls.Large)*/
	/*members, err := getMemberInfo(myClanTag)
	if err != nil {
		panic(err)
	}*/
	//fmt.Println(members.Items[0])

	//http.Handle("bower_components", http.FileServer(http.Dir("/tmp")))
	//http.HandleFunc("/", handleIndexPage)
	//log.Fatal(http.ListenAndServe(":8080", nil))

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}

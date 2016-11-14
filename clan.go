package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"github.com/johansundell/cocapi"
)

func init() {
	routes = append(routes, Route{"members", "GET", "/members", getMembers})
}

func getMembers(w http.ResponseWriter, req *http.Request) {
	sortDir := req.FormValue("sort")

	clan, err := cocapi.GetClanInfo(myClanTag)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		fmt.Println(err)
		return
	}

	switch sortDir {
	case "rate":
		sort.Sort(cocapi.DonationRatio(clan.MemberList))
		break
	case "role":
		sort.Sort(cocapi.Roles(clan.MemberList))
		break
	}

	b, err := json.Marshal(clan.MemberList)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		fmt.Println(err)
		return
	}
	fmt.Fprint(w, string(b))
}

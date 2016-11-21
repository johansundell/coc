package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"github.com/johansundell/cocapi"
)

func init() {
	routes = append(routes, Route{"members", "GET", "/members", handleGetMembers})
}

func handleGetMembers(w http.ResponseWriter, req *http.Request) {
	sortDir := req.FormValue("sort")

	clan, err := getMembers(myClanTag, sortDir)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		fmt.Println(err)
		return
	}

	b, err := json.Marshal(clan.MemberList)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		fmt.Println(err)
		return
	}
	fmt.Fprint(w, string(b))
}

func getMembers(clanTag, sortDir string) (cocapi.ClanInfo, error) {
	//clan, err := cocapi.GetClanInfo(myClanTag)
	clan, err := cocClient.GetClanInfo(myClanTag)
	if err != nil {
		return clan, err
	}

	switch sortDir {
	case "rate":
		sort.Sort(cocapi.DonationRatio(clan.MemberList))
		break
	case "role":
		sort.Sort(cocapi.Roles(clan.MemberList))
		break
	}
	return clan, nil
}

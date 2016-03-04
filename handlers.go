package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

type page struct {
	Title       string
	Name        string
	Description string
	MembersJson string
	Image       string
}

func handlePages(w http.ResponseWriter, req *http.Request) {
	t, err := template.New("clan-info.html").Delims("*{{", "}}*").ParseFiles("tmpl/clan-info.html")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		fmt.Println(err)
		return
	}

	p := page{}
	clan, err := getClanInfo(myClanTag)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		fmt.Println(err)
		return
	}
	//fmt.Println(clan.MemberList)
	b, err := json.Marshal(clan.MemberList)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		fmt.Println(err)
		return
	}
	p.Name = clan.Name
	p.Description = clan.Description
	p.MembersJson = string(b)
	p.Image = clan.BadgeUrls.Large

	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		fmt.Println(err)
	}
}

func handleIndexPage(w http.ResponseWriter, req *http.Request) {
	t, err := template.New("index.html").Delims("*{{", "}}*").ParseFiles("pages/index.html")

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		fmt.Println(err)
		return
	}
	/*clan, err := getClanInfo(myClanTag)
	if err != nil {
		panic(err)
	}*/
	p := page{}
	p.Title = "COC Playground"
	//p.Name = clan.Name
	//p.Description = clan.Description
	//b, err := json.Marshal(clan.MemberList)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//p.MembersJson = string(b)
	t.Delims("*{{", "}}*")
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		fmt.Println(err)
	}
}

func test() string {
	return "hello"
}

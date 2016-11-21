package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/johansundell/cocapi"
)

type page struct {
	Title       string
	Name        string
	Description string
	MembersJson string
	Image       string
}

type alert struct {
	page
	OldMembersJson string
}

func handlePages(w http.ResponseWriter, req *http.Request) {
	t, err := template.New("clan-info.html").Delims("*{{", "}}*").ParseFiles(basePath + "tmpl/clan-info.html")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		fmt.Println(err)
		return
	}

	p := page{}
	clan, err := getMembers(myClanTag, "rank")
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
	p.Name = clan.Name
	p.Description = clan.Description
	p.MembersJson = string(b)
	p.Image = clan.BadgeUrls.Small

	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		fmt.Println(err)
	}
}

func handleIndexPage(w http.ResponseWriter, req *http.Request) {
	t, err := template.New("index.html").Delims("*{{", "}}*").ParseFiles(basePath + "pages/index.html")

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		fmt.Println(err, basePath+"pages/index.html")
		return
	}

	p := page{}
	p.Title = "COC Playground"

	t.Delims("*{{", "}}*")
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		fmt.Println(err)
	}
}

func getAlerts(w http.ResponseWriter, req *http.Request) {
	t, err := template.New("alert.html").Delims("*{{", "}}*").ParseFiles(basePath + "pages/alert.html")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		fmt.Println(err, basePath+"pages/alert.html")
		return
	}
	p := page{}
	p.Title = "COC Errors"

	t.Delims("*{{", "}}*")
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		fmt.Println(err)
	}
}

func handleAlerts(w http.ResponseWriter, req *http.Request) {
	t, err := template.New("clan-errors.html").Delims("*{{", "}}*").ParseFiles(basePath + "tmpl/clan-errors.html")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		fmt.Println(err)
		return
	}

	p := alert{}

	var players = make([]cocapi.Player, 0)
	rows, err := db.Query("SELECT tag FROM members WHERE active = 1 AND exited > 0")
	if err != nil {
		log.Println(err)
		return
	}
	for rows.Next() {
		var tag string
		rows.Scan(&tag)
		//p, err := cocapi.GetPlayerInfo(tag)
		p, err := cocClient.GetPlayerInfo(tag)
		if err == nil {
			players = append(players, p)
		}

	}

	b, err := json.Marshal(players)
	p.OldMembersJson = string(b)
	rows, err = db.Query("SELECT GROUP_CONCAT(name) as usernames, tag, count(*) AS c FROM members GROUP BY tag HAVING c > 1")
	if err != nil {
		log.Println(err)
	}
	players = make([]cocapi.Player, 0)
	for rows.Next() {
		var tag string
		var usernames string
		var count int
		err = rows.Scan(&usernames, &tag, &count)
		if err != nil {
			log.Println(err)
		}
		log.Println("Found:", tag)
		//p, err := cocapi.GetPlayerInfo(tag)
		p, err := cocClient.GetPlayerInfo(tag)
		p.Name = usernames
		if err == nil {
			players = append(players, p)
		}
	}
	log.Println(players)
	b, err = json.Marshal(players)
	p.MembersJson = string(b)

	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		fmt.Println(err)
	}
}

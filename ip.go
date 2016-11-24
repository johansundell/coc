package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type IpInfo struct {
	Ip string `json:"ip"`
}

func (ip *IpInfo) Clean() {
	if strings.Contains(ip.Ip, ":") {
		ip.Ip = ip.Ip[:strings.Index(ip.Ip, ":")]
	}
}

func init() {
	routes = append(routes, Route{"ip", "GET", "/ip", handleIpReq})
}

func handleIpReq(w http.ResponseWriter, req *http.Request) {
	ip := IpInfo{req.RemoteAddr}
	ip.Clean()
	b, err := json.Marshal(ip)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		fmt.Println(err)
		return
	}
	fmt.Fprint(w, string(b))
}

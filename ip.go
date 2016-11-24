package main

import (
	"fmt"
	"net/http"
)

func init() {
	routes = append(routes, Route{"ip", "GET", "/ip", handleIpReq})
}

func handleIpReq(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, req.RemoteAddr)
}

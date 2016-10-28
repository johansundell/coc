package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handleIndexPage,
	},
	Route{
		"clan-info",
		"GET",
		"/tmpl/clan-info.html",
		handlePages,
	},
	Route{
		"Alerts",
		"GET",
		"/alert",
		getAlerts,
	},
	Route{
		"clan-errors",
		"GET",
		"/tmpl/clan-errors.html",
		handleAlerts,
	},
}

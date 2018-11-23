package router

import "net/http"

type Route struct {
	Name        string
	Pattern     string
	Method      string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

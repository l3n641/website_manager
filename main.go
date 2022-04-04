package main

import (
	"website_manager/api"
	"website_manager/conf"
)

func main() {
	conf.Init()
	r := api.LoadRouter()
	r.Run()
}

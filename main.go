package main

import (
	"monitorclient/routes"
	"monitorclient/settings"
)

func main() {
	settings.AppSettings, _ = settings.ReadSettings()
	routes.Init()
}

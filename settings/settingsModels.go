package settings

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Settings struct {
	AppParams    Params       `json:"app"`
	ServerParams ServerParams `json:"serverParams"`
	Services     []Service    `json:"services"`
}

type Params struct {
	Active  bool   `json:"active"`
	PortRun string `json:"portRun"`
	Log     bool   `json:"log"`
}

type ServerParams struct {
	Cpu    bool `json:"cpu"`
	Memory bool `json:"memory"`
	Disc   bool `json:"disc"`
}
type Service struct {
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

var AppSettings Settings

// ReadSettings to init app settings
func ReadSettings() (Settings, error) {
	var appParams Settings
	doc, err := ioutil.ReadFile("./settings-dev.json")
	if err != nil {
		if appParams.AppParams.Log {
			log.Println("ReadFile_error: ", err.Error())
		}
		panic(err.Error())
	}
	err = json.Unmarshal(doc, &appParams)
	if err != nil {
		if appParams.AppParams.Log {
			log.Println("Unmarshal error: ", err.Error())
		}
		panic(err.Error())
	}
	return appParams, err
}

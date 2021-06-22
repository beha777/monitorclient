package settings

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Settings struct {
	AppParams    Params        `json:"app"`
	BotParams    BotSettings   `json:"botParams"`
	Hosts        []Host        `json:"hosts"`
	ServerParams []ServerParam `json:"serverParams"`
	Services     []Service     `json:"services"`
}

type Params struct {
	Active  bool    `json:"active"`
	PortRun string  `json:"portRun"`
	Log     bool    `json:"log"`
	Version float64 `json:"version"`
}

type BotSettings struct {
	Url      string `json:"url"`
	Login    string `json:"login"`
	Password string `json:"password"`
	UrlId    string `json:"urlId"`
	Token    string `json:"token"`
	ChatID   string `json:"chat_id"`
	Owners   string `json:"owners"`
}

type Host struct {
	Name               string `json:"name"`
	CheckPeriod        uint   `json:"check_period"`
	NotificationPeriod uint   `json:"notification_period"`
	Active             bool   `json:"active"`
}

type Service struct {
	Name               string `json:"name"`
	State              string `json:"state"`
	CheckPeriod        uint   `json:"check_period"`
	NotificationPeriod uint   `json:"notification_period"`
	Active             bool   `json:"active"`
}

type ServerParam struct {
	Name               string  `json:"name"`
	Condition          string  `json:"condition"`
	Limit              float64 `json:"limit"`
	CheckPeriod        uint    `json:"check_period"`
	NotificationPeriod uint    `json:"notification_period"`
	Active             bool    `json:"active"`
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

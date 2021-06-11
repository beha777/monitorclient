package monitorclient

import "monitorclient/settings"

func main() {
	settings.AppSettings, _ = settings.ReadSettings()



}

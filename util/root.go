package util

import (
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		Fatal(err)
	}
}

func CheckErrorWithMsg(err error, msg string) {
	if err != nil {
		Fatal(msg + " - " + err.Error())
	}
}

func ReportStatus(configPath string) {
	url := viper.GetString(configPath)

	if url == "" || !strings.HasPrefix(url, "http") {
		viper.Set(configPath, "your_url_here")
		Fatal("Invalid url! Please manually set the push url in the config first!")
	}

	_, err := http.Get(url)
	CheckErrorWithMsg(err, "Error contacting remote server, is the push url valid?")
}
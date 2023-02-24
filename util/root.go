package util

import (
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

var (
	ViperInstance *viper.Viper
)

func InitConfig(viperInstance *viper.Viper) {
	ViperInstance = viperInstance
}

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
	url := ViperInstance.GetString(configPath)

	if url == "" || !strings.HasPrefix(url, "http") {
		ViperInstance.Set(configPath, "your_url_here")
		err := ViperInstance.WriteConfig()
		CheckErrorWithMsg(err, "Couldn't write config")
		Fatal("Invalid url for " + configPath + "! Please manually set the push url in the config first!")
		return
	}

	_, err := http.Get(url)
	CheckErrorWithMsg(err, "Error contacting remote server, is the push url valid?")
}

func ArrayMap[T, U any](data []T, f func(T) U) []U {

	res := make([]U, 0, len(data))

	for _, e := range data {
		res = append(res, f(e))
	}

	return res
}

func SetNodeUrlIfEmpty(configItem string) {
	if ViperInstance.GetString(configItem) == "" {
		ViperInstance.Set(configItem, "your_url_here")
		err := ViperInstance.WriteConfig()
		CheckErrorWithMsg(err, "Couldn't write config")
	}
}

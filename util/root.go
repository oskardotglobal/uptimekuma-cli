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

func ReportStatus(viperInstance *viper.Viper, configPath string) {
	url := viperInstance.GetString(configPath)

	if url == "" || !strings.HasPrefix(url, "http") {
		viperInstance.Set(configPath, "your_url_here")
		Fatal("Invalid url! Please manually set the push url in the config first!")
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
	if viper.GetString(configItem) == "" {
		viper.Set(configItem, "your_url_here")
	}
}

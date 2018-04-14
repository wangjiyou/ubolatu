package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	GetGroups           string
	SentResult          string
	Beat                int
	UserName            string
	Password            string
	IPAddress           string
	DBName              string
	PortAddress         string
	AlarmUrl            string
	StatusChangeRatio   float32
	ReportedRatioPeriod int
	HttpTimeout         int
	SecretKey           string
	AccessKey           string
}

var _globalConfig Config

func GlobalConfig() *Config {
	return &_globalConfig
}

func ConfigFile(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &_globalConfig)

}

func LoadConfigFile(configFileName string) error {
	var filenames []string
	if configFileName != "" {
		filenames = append(filenames, configFileName)
	} else {
		filenames = []string{"config/config.json"}
	}
	var err error
	for _, filename := range filenames {
		err = ConfigFile(filename)
		if err == nil {
			return nil
		}
	}
	return err
}

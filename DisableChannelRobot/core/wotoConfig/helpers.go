package wotoConfig

import (
	"encoding/json"
	"io/ioutil"
)

func ParseConfig(configFile string) (*GenshinBotConfig, error) {
	if ConfigSettings != nil {
		return ConfigSettings, nil
	}

	s := &GenshinBotConfig{}

	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, s)
	if err != nil {
		return nil, err
	}

	ConfigSettings = s

	return ConfigSettings, nil
}

func LoadConfig() (*GenshinBotConfig, error) {
	return ParseConfig("config.json")
}

func GetCmdPrefixes() []rune {
	return []rune{'/', '!'}
}

func GetBotToken() string {
	if ConfigSettings != nil {
		return ConfigSettings.BotToken
	}
	return ""
}

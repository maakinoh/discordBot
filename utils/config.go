package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/maakinoh/discordBot/plugin"
)

var instance *Configuration

type Configuration struct {
	CommandPrefix string
	Plugins       []plugin.Plugin
}

func createBotDirAndFiles() Configuration {

	os.Mkdir(".bot", os.ModePerm)
	os.Mkdir(".bot/plugins", os.ModePerm)
	cfg := Configuration{CommandPrefix: "!", Plugins: make([]plugin.Plugin, 0)}
	json, _ := json.Marshal(cfg)
	ioutil.WriteFile(".bot/config.json", json, os.ModePerm)
	return cfg
}

// GetConfigurationInstance return the current config instance.
// If the instance is nil it will read the config files from the .bot dir.
// If on the other hand this dir neither exist it will create a new one.
func GetConfigurationInstance() *Configuration {
	if instance == nil {
		dir, _ := os.Getwd()
		if _, err := os.Stat(dir + "/.bot"); os.IsNotExist(err) {
			fmt.Println("Bot dir not exist")
			c := createBotDirAndFiles()
			instance = &c
		} else {
			cfgFile, _ := ioutil.ReadFile(".bot/config.json")

			cfg := Configuration{}
			json.Unmarshal(cfgFile, &cfg)
			instance = &cfg
		}
	}
	return instance
}

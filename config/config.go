package config

import (
	"onlinemall/until"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

type datasourceConfig struct {
	Dialect         string
	Dburl           string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int
	LogDir          string
	LogFile         string
}

var DatasourceConfig datasourceConfig

type serverConfig struct {
	Environment string
	Mod			string
	LogDir      string
	LogFile     string
	Address     string
}

var ServerConfig serverConfig

func initConfig() {
	var configData map[string]interface{}

	bytes, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		fmt.Println("init config: ", err.Error())
		os.Exit(-1)
	}

	reg := regexp.MustCompile(`/\*.*\*/`)
	configStr := reg.ReplaceAllString(string(bytes[:]), "")
	bytes = []byte(configStr)

	if err := json.Unmarshal(bytes, &configData); err != nil {
		fmt.Println("invalid config: ", err.Error())
		os.Exit(-1)
	}

	//init datasource config
	util.SetStructByJSON(&DatasourceConfig, configData["datasource"].(map[string]interface{}))

	//init server config
	util.SetStructByJSON(&ServerConfig, configData["server"].(map[string]interface{}))
}

func init() {
	initConfig()
}

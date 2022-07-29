package main

import (
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
)

type SoftwareConfig struct {
	Port int

	DatabaseType     string
	DatabaseHost     string
	DatabasePort     int
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string

	LogDirectory string
}

type GlobalVar struct {
	Config   SoftwareConfig
	Logger   *zap.SugaredLogger
	DateBase *gorm.DB
}

// Global 全局变量
var Global GlobalVar

func InitGlobalVar(configFile string) {
	config, err := getConfigFromYaml(ConfigFile)
	if err != nil {
		log.Fatalln(err)
		return
	}

	logger, err := getInitLogger(config.LogDirectory)
	if err != nil {
		log.Println(err)
	}

	db, err := openDateBase(config)
	if err != nil {
		logger.Error("打开数据库失败！", err)
	}

	Global.Config = config
	Global.Logger = logger
	Global.DateBase = db
}

// getConfigFromYaml 从配置文件中读取软件配置文件
func getConfigFromYaml(configFile string) (config SoftwareConfig, err error) {

	//config.Port = 8060
	//config.DatabaseType = "postgres"
	//config.DatabaseHost = "127.0.0.1"
	//config.DatabasePort = 5432
	//config.DatabaseUser = "openpg"
	//config.DatabasePassword = "openpgpwd"
	//config.DatabaseName = "my"
	//config.LogDirectory = "./log"

	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return
	}

	return

}

func GetConfig() SoftwareConfig {
	return Global.Config
}

func GetLogger() *zap.SugaredLogger {

	return Global.Logger
}

func GetDateBase() *gorm.DB {

	return Global.DateBase
}

//func SaveSoftwareConfig() (err error) {
//
//	yamlBytes, err := yaml.Marshal(Global.Config)
//	if err != nil {
//		return
//	}
//
//	err = saveFile(SoftwareConfigFile, yamlBytes)
//	if err != nil {
//		return
//	}
//
//	return
//}

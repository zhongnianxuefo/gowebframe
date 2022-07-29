package main

// SoftwareConfigFile 配件文件
const ConfigFile = "config.yaml"

func main() {

	InitGlobalVar(ConfigFile)

	log := GetLogger()
	db := GetDateBase()

	err := InitDataBase(db)
	if err != nil {
		log.Error(err)
	}

	InitGin()

	return
}

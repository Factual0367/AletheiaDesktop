package config

import "log"

func UpdateDownloadPath(newDownloadPath string) {
	currentConfigContent, configReadErr := ReadConfigFile()
	if configReadErr != nil {
		log.Fatalln(configReadErr.Error())
	}
	currentConfigContent["downloadPath"] = newDownloadPath
	writeConfigErr := WriteConfigFile(currentConfigContent)
	if writeConfigErr != nil {
		log.Fatalln(writeConfigErr.Error())
	}
}

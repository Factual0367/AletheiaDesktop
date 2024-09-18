package config

import "log"

func UpdateDownloadPath(newDownloadPath string) {
	// should move if there are any books in the
	// old download path and ask
	currentConfigContent, configReadErr := ReadConfigFile()
	if configReadErr != nil {
		log.Fatalln(configReadErr.Error())
	}
	currentConfigContent["downloadLocation"] = newDownloadPath
	writeConfigErr := WriteConfigFile(currentConfigContent)
	if writeConfigErr != nil {
		log.Fatalln(writeConfigErr.Error())
	}
}

package adbt

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// BackLog defines a log unit
type BackLog struct {
	Identifier string `json:"identifier"`
	Success bool `json:"success"`
	Log string `json:"log"`
	Time string `json:"time"`
}

type logs struct {
	BackLogs []BackLog `json:"logs"`
}

func getLogFilePath() string {
	createFolderIfNotExist("adbt")
	createFolderIfNotExist(filepath.Join("adbt", "log"))
	logFilePath := filepath.Join("adbt", "log", "log.json")
	if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
		createFileIfNotExist(logFilePath)
	}
	return logFilePath
}

func (b BackLog) writeLog() {
	logs := readLog()
	logs.BackLogs = append(logs.BackLogs, b)
	fileContent, err := json.Marshal(logs)
	if err != nil {
		log.Panic(err)
	}
	_ = ioutil.WriteFile(getLogFilePath(), fileContent,0644)
}

func readLog() logs {
	var backlogs logs
	jsonFile, err := os.Open(getLogFilePath())
	if err != nil {
		log.Fatal(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &backlogs)
	return backlogs
}
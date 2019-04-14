package adbt

import (
	"bytes"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type backupScheduler struct {
	Identifier  string `toml:"identifier" json:"identifier"`
	URI         string `toml:"uri" json:"uri"`
	Periodicity string `toml:"periodicity" json:"periodicity"`
	Time        string `toml:"time" json:"time"`
	Database    string `toml:"database" json:"database"`
	Name        string `toml:"name" json:"name"`
}

type adbtConfig struct {
	Jobs []backupScheduler `toml:"scheduler"`
}

func getDefaultConfig() adbtConfig {
	var config = adbtConfig{}
	return config
}

func getConfigFilePath() string {
	createFolderIfNotExist("data")
	createFolderIfNotExist(filepath.Join("data", "config"))
	adbtConfigToml := filepath.Join("data", "config", "config.toml")
	if _, err := os.Stat(adbtConfigToml); os.IsNotExist(err) {
		createFileIfNotExist(adbtConfigToml)
		// config file not exists, write default to it
		writeConfig(getDefaultConfig())
	}
	return adbtConfigToml
}

func readConfig() adbtConfig {
	var config adbtConfig
	configFile := getConfigFilePath()
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		return config
	}
	return config
}

func writeConfig(config adbtConfig) {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(config); err != nil {
		log.Fatal(err)
	}
	configFile := getConfigFilePath()
	err := ioutil.WriteFile(configFile, buf.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func addJobConfig(s backupScheduler) adbtConfig {
	s.Identifier = uuid.New().String()
	config := readConfig()
	config.Jobs = append(config.Jobs, s)
	writeConfig(config)
	return readConfig()
}

func removeJobConfig(jobIdentifier string) adbtConfig {
	config := readConfig()
	var removeIdx int
	for idx, job := range config.Jobs {
		if job.Identifier == jobIdentifier {
			removeIdx = idx
		}
	}
	config.Jobs = append(config.Jobs[:removeIdx], config.Jobs[removeIdx+1:]...)
	writeConfig(config)
	return config
}

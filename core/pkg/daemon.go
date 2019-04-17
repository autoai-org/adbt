package adbt

import (
	"fmt"
	"github.com/kardianos/service"
	"log"
	"os"
	"os/user"
	"runtime"
)


type sol struct {
}

func (s *sol) Start(srv service.Service) error {
	go NewServer(DaemonPort)
	return nil
}

func (s *sol) Stop(srv service.Service) error {
	return nil
}

func getRunUser() string {
	currentUser, _ := user.Current()
	if currentUser.Username == "root" && runtime.GOOS != "windows" {
		return os.Getenv("SUDO_USER")
	}
	return currentUser.Username
}

func getADBTDConfig() *service.Config {
	realUsername := getRunUser()
	currentWDirectory, _ := os.Getwd()
	fmt.Println(currentWDirectory)
	srvConf := &service.Config{
		Name:        "adbtd",
		DisplayName: "ADBT Daemon",
		Description: "Automatic Database Backup Toolkits[Daemon]",
		Arguments:   []string{"serve"},
		UserName:    realUsername,
		WorkingDirectory: currentWDirectory,
	}
	return srvConf
}

// InstallService install the background daemon service into the system
func InstallService() {
	srvConfig := getADBTDConfig()
	dae := &sol{}
	s, err := service.New(dae, srvConfig)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Install()
	if err != nil {
		log.Fatal(err)
	}
	err = s.Start()
	if err != nil {
		log.Fatal(err)
	}
}

// UninstallService remove the background daemon service from system
func UninstallService() {
	srvConfig := getADBTDConfig()
	dae := &sol{}
	s, err := service.New(dae, srvConfig)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Stop()
	if err != nil {
		log.Fatal(err)
	}
	err = s.Uninstall()
	if err != nil {
		log.Fatal(err)
	}
}

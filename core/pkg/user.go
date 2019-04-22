package adbt

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin/json"
	"github.com/levigross/grequests"
	"log"
	"os"
	"strconv"
	"time"
)

// login stores uid and username into config file
func Login(username string, uid string) {
	config:= readConfig()
	config.UID = uid
	config.Username = username
	writeConfig(config)
	Sync()
}

// Sync uploads config and log into leancloud
func Sync() error {
	config := readConfig()
	if config.Username == "" || config.UID == "" {
		fmt.Errorf("Cannot sync because you did not login yet.")
		return errors.New("Cannot sync because you did not login yet.")
	}
	logs := readLog()
	configJsonString, err := json.Marshal(config)
	logsJsonString, err := json.Marshal(logs)
	if err != nil {
		fmt.Errorf("Failed to parse Config or Log file")
		fmt.Printf(err.Error())
	}
	var params = make(map[string]string)
	params["config"] = string(configJsonString)
	params["log"] = string(logsJsonString)
	params["username"] = config.Username
	params["uid"] = config.UID
	lcAdd(params)
	return nil
}

func lcAdd(params map[string]string) *grequests.Response {
	url := "https://em7amivq.api.lncld.net/1.1/classes/content"
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	sign := fmt.Sprintf("%x", md5.Sum([]byte(timestamp+os.Getenv("LC_KEY"))))
	finalSign := sign + "," + timestamp
	resp, err := grequests.Post(url, &grequests.RequestOptions{
		JSON:    params,
		Headers: map[string]string{
			"X-LC-Id": os.Getenv("LC_ID"),
			"X-LC-Sign": finalSign,
			"Content-Type":"application/json",
	},
	})
	if err != nil {
		log.Fatal(err)
	}
	if !resp.Ok {
		log.Fatal("Bad Response from Daemon")
	}
	fmt.Println("Successfully Uploaded config/log file.")
	return resp
}
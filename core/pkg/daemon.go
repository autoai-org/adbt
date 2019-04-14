package adbt

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const DaemonPort = "10591"

var sched Scheduler

type deleteSchedulerRequest struct {
	Identifier string `json:"identifier"`
}

func createSchedulerHandler(c *gin.Context) {
	var createSchedulerRequest backupScheduler
	c.BindJSON(&createSchedulerRequest)
	addJobConfig(createSchedulerRequest)
	config := readConfig()
	sched = loadScheduler(config)
	c.JSON(http.StatusOK, sched.jobs)
}

func getAllJobs(c *gin.Context) {
	c.JSON(http.StatusOK, sched.jobs)
}

func removeJob(c *gin.Context) {
	var deleteReq deleteSchedulerRequest
	c.BindJSON(&deleteReq)
	config := removeJobConfig(deleteReq.Identifier)
	sched = loadScheduler(config)
	c.JSON(http.StatusOK, sched.jobs)
}

// Creates a Server
func NewServer(port string) {
	sched = loadScheduler(readConfig())
	// read config and load
	r := gin.Default()
	r.Use(gin.Recovery())
	r.POST("/jobs", createSchedulerHandler)
	r.GET("/jobs", getAllJobs)
	r.DELETE("/jobs", removeJob)
	r.Run("0.0.0.0:" + port)
}

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
	config := readConfig()
	jobExists := false
	for _, job := range config.Jobs {
		if job.Name == createSchedulerRequest.Name {
			jobExists = true
		}
	}
	if jobExists {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":400,
			"status":"FAILED",
			"info": "Name already exists",
		})
	} else {
		addJobConfig(createSchedulerRequest)
		config = readConfig()
		sched = loadScheduler(config)
		c.JSON(http.StatusOK, sched.jobs)
	}
}

func modifySchedulerHandler(c * gin.Context) {
	identifier := c.Param("identifier")
	var modifySchedulerRequest backupScheduler
	c.BindJSON(&modifySchedulerRequest)
	config := readConfig()
	jobExists := false
	var requestedId int
	for idx, job := range config.Jobs {
		if job.Identifier == identifier {
			jobExists = true
			requestedId = idx
		}
	}
	if jobExists {
		config.Jobs[requestedId] = modifySchedulerRequest
		writeConfig(config)
		c.JSON(http.StatusOK, config.Jobs[requestedId])
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":404,
			"status":"FAILED",
			"info": "Job Not Found",
		})
	}

}

func getAllJobs(c *gin.Context) {
	c.JSON(http.StatusOK, sched.jobs)
}

func getJobDetail(c *gin.Context) {
	identifier := c.Param("identifier")
	config := readConfig()
	var requestedJob backupScheduler
	var hasFound bool = false
	for _, job := range config.Jobs {
		if (job.Identifier == identifier) {
			requestedJob = job
			hasFound = true
		}
	}
	if hasFound {
		c.JSON(http.StatusOK, requestedJob)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"code":404,
			"status":"FAILED",
			"info": "Requested Job Not Found",
		})
	}
}

func removeJob(c *gin.Context) {
	var deleteReq deleteSchedulerRequest
	c.BindJSON(&deleteReq)
	config := removeJobConfig(deleteReq.Identifier)
	sched = loadScheduler(config)
	c.JSON(http.StatusOK, sched.jobs)
}

func getStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":200,
		"status":"OK",
		"info": "Server is Running",
	})
}

func backupNow(c *gin.Context) {
	var createSchedulerRequest backupScheduler
	c.BindJSON(&createSchedulerRequest)
	if createSchedulerRequest.Database == "MongoDB" {
		m := MongoDB{
			URI: createSchedulerRequest.URI,
			Name: createSchedulerRequest.Name,
		}
		isFinished := m.Backup("")
		if isFinished {
			c.JSON(http.StatusOK, gin.H{
				"code":200,
				"status":"OK",
				"info": "Backup Finished",
			})
		} else {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"code":500,
				"status":"FAILED",
				"info": "Cannot Finish Backup",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":400,
			"status":"FAILED",
			"info": "Unsupported Database",
		})
	}
}

func testConnString(c *gin.Context) {
	var createSchedulerRequest backupScheduler
	c.BindJSON(&createSchedulerRequest)
	if createSchedulerRequest.Database == "MongoDB" {
		m := MongoDB{
			URI: createSchedulerRequest.URI,
			Name: createSchedulerRequest.Name,
		}
		isConnectable := m.test()
		if isConnectable {
			c.JSON(http.StatusOK, gin.H{
				"code":200,
				"status":"OK",
				"info": "Database Connected",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":400,
				"status":"FAILED",
				"info": "Cannot Connect to Database",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":400,
			"status":"FAILED",
			"info": "Unsupported Database",
		})
	}
}

// beforeResponse handles cors settings
func beforeResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		c.Writer.Header().Set("adbt-version", "0.1")
		if c.Writer.Header().Get("Access-Control-Allow-Origin") == "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		if c.Request.Method == "OPTIONS" {
			c.Writer.WriteHeader(http.StatusOK)
		}
	}
}

// Creates a Server
func NewServer(port string) {
	sched = loadScheduler(readConfig())
	// read config and load
	r := gin.Default()
	r.Use(beforeResponse())
	r.Use(gin.Recovery())
	r.GET("/job/:identifier", getJobDetail)
	r.PATCH("/job/:identifier", modifySchedulerHandler)
	r.POST("/jobs", createSchedulerHandler)
	r.POST("/jobs/now", backupNow)
	r.GET("/jobs", getAllJobs)
	r.DELETE("/jobs", removeJob)
	r.POST("/databases/test", testConnString)
	r.GET("/status", getStatus)
	r.Run("0.0.0.0:" + port)
}

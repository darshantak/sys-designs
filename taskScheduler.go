package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type Task struct {
	TaskId   string      `json:"task_id"`
	ExecTime int         `json:"exec_time"`
	Payload  interface{} `json:"payload"`
	Timer    *time.Timer `json:"-"`
}

var taskStore = make(map[string]*Task)
var taskMu sync.Mutex

func addTask(task Task) {
	taskMu.Lock()
	defer taskMu.Unlock()

	delay := time.Until(time.Unix(int64(task.ExecTime), 0))
	if delay < 0 {
		log.Printf("Task %s has an invalid execution time", task.TaskId)
		return
	}

	task.Timer = time.AfterFunc(delay, func() {
		execTask(task.TaskId)
	})
	taskStore[task.TaskId] = &task
	log.Printf("Executing task %s: %+v", task.TaskId, task.Payload)
}

func execTask(taskId string) {
	taskMu.Lock()
	defer taskMu.Unlock()

	if _, exists := taskStore[taskId]; !exists {
		return
	}
	delete(taskStore, taskId)
}
func scheduleTaskHandler(c *gin.Context) {
	var req Task
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "request body invalid"})
		return
	}

	if time.Now().Unix() > int64(req.ExecTime) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "please provide future time"})
		return
	}

	addTask(req)
	c.JSON(http.StatusOK, gin.H{"message": "task scheduled"})

}

func getTasks(c *gin.Context) {
	c.JSON(http.StatusOK, taskStore)
}
func deleteTask(c *gin.Context) {
	taskId := c.Param("id")
	delete(taskStore, taskId)
	c.JSON(http.StatusOK, gin.H{"message": "task deleted"})

}
func TaskMain() {
	r := gin.Default()
	r.POST("/task", scheduleTaskHandler)
	r.GET("/task", getTasks)
	r.DELETE("/task/:id", deleteTask)
	r.Run(":8181")
}

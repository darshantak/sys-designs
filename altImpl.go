package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type Priority string

var mu sync.Mutex

const (
	LOW    Priority = "LOW"
	MEDIUM Priority = "MEDIUM"
	HIGH   Priority = "HIGH"
)

type TaskManager struct {
	Id       string
	TaskList map[string]Task
	mu       sync.Mutex
}

type Task struct {
	Name           string
	Status         bool
	Prioritization Priority
	Deadline       time.Time
}

func NewTaskManager(id string) *TaskManager {
	return &TaskManager{
		Id:       id,
		TaskList: make(map[string]Task),
	}
}

func (tm *TaskManager) createTask(name string, status bool, priority string, mu *sync.Mutex) bool {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := tm.TaskList[name]; exists {
		fmt.Printf("%s already exists", name)
		return false
	}
	newTask := Task{
		Name:           name,
		Status:         status,
		Prioritization: Priority(priority),
		Deadline:       time.Now().Add(24 * time.Hour)}
	tm.TaskList[name] = newTask
	return true
}

func (tm *TaskManager) updateTask(name string, paramKey string, paramValue interface{}) bool {
	if _, exists := tm.TaskList[name]; !exists {
		fmt.Printf("%s doesn't exist", name)
		return false
	}
	task := tm.TaskList[name]
	task.Status = true
	tm.TaskList[name] = task
	switch strings.ToLower(paramKey) {
	case "name":
		task.Name = paramValue.(string)
	case "status":
		task.Status = paramValue.(bool)
	case "priority":
		v := paramValue.(string)
		task.Prioritization = Priority(v)
	case "deadline":
		task.Deadline = paramValue.(time.Time)
	}

	return true
}

func (tm *TaskManager) deleteTask(name string) bool {
	if _, exists := tm.TaskList[name]; !exists {
		fmt.Printf("%s doesn't exist", name)
		return false
	}
	delete(tm.TaskList, name)
	return true
}

func TaskManagerCLI() {
	scanner := bufio.NewScanner(os.Stdin)
	LocalNewTM := NewTaskManager("TM123")

	for {
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		if strings.ToLower(input) == "exit" {
			break
		}
		parts := strings.Split(input, " ")
		if len(parts) == 0 {
			continue
		}
		command, args := parts[0], parts[1:]
		switch command {
		case "add":
			if len(args) != 2 {
				fmt.Println("usage: add [Task Name]")
				continue
			} else {
				flag := LocalNewTM.createTask(args[0], false, args[1], &mu)
				fmt.Println(flag)
			}
		case "update":
			if len(args) != 3 {
				fmt.Println("usage: update [Task Name]")
				continue
			} else {
				flag := LocalNewTM.updateTask(args[0], args[1], args[2])
				fmt.Println(flag)
			}
		case "delete":
			if len(args) != 1 {
				fmt.Println("usage: delete [Task Name]")
				continue
			} else {
				flag := LocalNewTM.deleteTask(args[0])
				fmt.Println(flag)
			}
		}
	}

}

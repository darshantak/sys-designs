package cmd

import (
	"fmt"
)

type TaskManager struct {
	Id       string
	TaskList map[string]Task
}

type Task struct {
	Name   string
	Status bool
}

const stateFile = "tmState.json"

func NewTaskManager(id string) *TaskManager {
	return &TaskManager{
		Id:       id,
		TaskList: make(map[string]Task),
	}
}

func (tm *TaskManager) CreateTask(name string, status bool) bool {
	if _, exists := tm.TaskList[name]; exists {
		fmt.Printf("%s already exists", name)
		return false
	}
	newTask := Task{
		Name:   name,
		Status: false,
	}
	tm.TaskList[name] = newTask
	return true
}

func (tm *TaskManager) UpdateTask(name string) bool {
	if _, exists := tm.TaskList[name]; !exists {
		fmt.Printf("%s doesn't exist", name)
		return false
	}
	task := tm.TaskList[name]
	task.Status = true
	tm.TaskList[name] = task
	return true
}

func (tm *TaskManager) DeleteTask(name string) bool {
	if _, exists := tm.TaskList[name]; !exists {
		fmt.Printf("%s doesn't exist", name)
		return false
	}
	delete(tm.TaskList, name)
	return true
}

var NewTM = NewTaskManager("TM123")

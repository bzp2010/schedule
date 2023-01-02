package models

import (
	"database/sql"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// TaskType is the type used to describe the type of task
type TaskType uint8

const (
	// TaskTypeShell is a task for starting external process via shell
	TaskTypeShell TaskType = iota
	// TaskTypeWebhook is a task type that triggers an external webhook
	TaskTypeWebhook
)

// Task define the data structure of a timed task
type Task struct {
	// gorm base model
	gorm.Model

	// Name of task
	Name string

	// Type indicates the type of this task
	Type TaskType

	// Configuration
	Configuration datatypes.JSON

	// Rules of task
	Rules []TaskRule

	// LastRunningAt indicates the last running time of the current task
	LastRunningAt sql.NullTime

	// LastRunningTime indicates the time consumed for the last run of the current task
	LastRunningTime int64

	// Status indicates the current task status
	Status Status
}

// TaskConfigurationShell define the data structure of shell task configuration
type TaskConfigurationShell struct {
	Command string `json:"command"`
	Timeout time.Duration
}

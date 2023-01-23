package models

import (
	"database/sql"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// TaskType is the type used to describe the type of task
type TaskType string

const (
	// TaskTypeShell is a task for starting external process via shell
	TaskTypeShell TaskType = "SHELL"
	// TaskTypeWebhook is a task type that triggers an external webhook
	TaskTypeWebhook TaskType = "WEBHOOK"
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
	Command string        `json:"command"`
	Timeout time.Duration `json:"timeout"`
}

func (tcs TaskConfigurationShell) IsTaskConfiguration() {}

// TaskConfigurationWebhook define the data structure of webhook task configuration
type TaskConfigurationWebhook struct {
	URL    string `json:"url"`
	Method string `json:"method"`
}

func (tcw TaskConfigurationWebhook) IsTaskConfiguration() {}

// Implement GraphQL's models.Model interface to simplify data assembly,
// the data corresponding to models.Model is gorm.Model

func (t Task) IsModel() {}

func (t Task) GetID() int64 {
	return int64(t.ID)
}

func (t Task) GetCreatedAt() int64 {
	return t.CreatedAt.UnixMilli()
}

func (t Task) GetUpdatedAt() int64 {
	return t.UpdatedAt.UnixMilli()
}

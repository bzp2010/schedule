package models

import (
	"database/sql"

	"gorm.io/gorm"
)

// TaskRule define the data structure of task running rule
type TaskRule struct {
	// gorm base model
	gorm.Model

	// TaskID indicates the task ID associated with the rule
	TaskID uint

	// Task indicates the task associated with the Job
	Task Task

	// Description of task rule
	Description sql.NullString

	// Rule of task
	Rule string

	// LastRunningAt indicates the last running time of the current task rule
	LastRunningAt sql.NullTime

	// LastRunningTime indicates the time consumed for the last run of the current task rule
	LastRunningTime int64

	// Status indicates the current task rule status
	Status Status
}

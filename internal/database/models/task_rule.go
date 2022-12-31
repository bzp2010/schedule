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

	// Description of task rule
	Description sql.NullString

	// Rule of task
	Rule string

	// Status indicates the current task rule status
	Status Status
}

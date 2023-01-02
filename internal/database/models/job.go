package models

import (
	"time"

	"gorm.io/gorm"
)

type JobFlag uint8

const (
	// JobFlagTimeout indicates that the task execution exceeds the time limit
	JobFlagTimeout JobFlag = 1 << iota
)

// Job define the data structure of a task execution
type Job struct {
	// gorm base model
	gorm.Model

	// TaskID indicates the task ID associated with the Job
	TaskID uint

	// Task indicates the task associated with the Job
	Task Task

	// TaskRuleID indicates the task rule ID associated with the Job
	TaskRuleID uint

	// TaskRule indicates the task rule associated with the Job
	TaskRule TaskRule

	// Stdout stores the standard output of the Job
	Stdout string

	// Stderr stores the standard error output of the Job
	Stderr string

	// StartAt stores the start time of the Job
	StartAt time.Time

	// StopAt stores the stop time of the Job
	StopAt time.Time

	// Flags store multiple flag values to indicate Job status
	Flags uint64
}

package models

import (
	"time"

	"gorm.io/gorm"
)

// JobFlag is a unique type of flag in Job, which is used as bitmask
type JobFlag uint64

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

// AddFlag helps the caller to add flags to the Job record
func (job *Job) AddFlag(flag JobFlag) {
	job.Flags |= uint64(flag)
}

// HasFlag helps the caller to check the flags in the Job record
func (job *Job) HasFlag(flag JobFlag) bool {
	return job.Flags&uint64(flag) != 0
}

// RemoveFlag helps the caller to remove the flags from the Job record
func (job *Job) RemoveFlag(flag JobFlag) {
	job.Flags &= ^uint64(flag)
}

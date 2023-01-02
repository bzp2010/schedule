package task

import (
	"database/sql"
	"fmt"

	"github.com/bzp2010/schedule/internal/database"
	"github.com/bzp2010/schedule/internal/database/models"
)

// Task is the base field for each task
type Task struct {
	TaskID     uint
	TaskRuleID uint
}

// SaveJob is used to store task execution records
func (t *Task) SaveJob(job models.Job) {
	var (
		db          = database.GetDatabase()
		runningTime = job.StopAt.Sub(job.StartAt).Milliseconds()
	)

	// create job
	db.Create(&job)

	// update task last running
	db.Model(models.Task{}).
		Where("id = ?", job.TaskID).
		Updates(models.Task{
			LastRunningAt: sql.NullTime{
				Time:  job.StartAt,
				Valid: true,
			},
			LastRunningTime: runningTime,
		})

	// update task rule last running
	db.Model(models.TaskRule{}).
		Where("id = ?", job.TaskRuleID).
		Updates(models.TaskRule{
			LastRunningAt: sql.NullTime{
				Time:  job.StartAt,
				Valid: true,
			},
			LastRunningTime: runningTime,
		})
}

// PrintJob is used to print a task execution record
func (t *Task) PrintJob(job models.Job) {
	fmt.Println("[START] ", job.StartAt)
	fmt.Println("[STDOUT] ", job.Stdout)
	fmt.Println("[STDERR] ", job.Stderr)
	fmt.Println("[STOP] ", job.StopAt)
	fmt.Println("[COST] ", job.StopAt.Sub(job.StartAt).Milliseconds(), "ms")
	fmt.Println("----------------------------------------")
}

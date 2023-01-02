package scheduler

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/bzp2010/schedule/internal/database"
	"github.com/bzp2010/schedule/internal/database/models"
)

type jobRecord struct {
	taskID     uint
	taskRuleID uint
	start      time.Time
	stop       time.Time
	timeout    bool
	stdout     string
	stderr     string
}

func newJobRecord(taskID, taskRuleID uint) *jobRecord {
	return &jobRecord{
		start:      time.Now(),
		taskID:     taskID,
		taskRuleID: taskRuleID,
	}
}

func (jr *jobRecord) Save() {
	jr.stop = time.Now()

	// generate job flag
	var flags uint64 = 0
	if jr.timeout {
		flags |= uint64(models.JobFlagTimeout)
	}

	var (
		db          = database.GetDatabase()
		runningTime = jr.stop.Sub(jr.start).Milliseconds()
	)

	// create job
	db.Create(models.Job{
		TaskID:     jr.taskID,
		TaskRuleID: jr.taskRuleID,
		StartAt:    jr.start,
		StopAt:     jr.stop,
		Stdout:     jr.stdout,
		Stderr:     jr.stderr,
		Flags:      flags,
	})

	// update task last running
	db.Model(models.Task{}).
		Where("id = ?", jr.taskID).
		Updates(models.Task{
			LastRunningAt: sql.NullTime{
				Time:  jr.start,
				Valid: true,
			},
			LastRunningTime: runningTime,
		})

	// update task rule last running
	db.Model(models.TaskRule{}).
		Where("id = ?", jr.taskRuleID).
		Updates(models.TaskRule{
			LastRunningAt: sql.NullTime{
				Time:  jr.start,
				Valid: true,
			},
			LastRunningTime: runningTime,
		})
}

func (jr *jobRecord) printLog() {
	fmt.Println("[START] ", jr.start)
	fmt.Println("[STDOUT] ", jr.stdout)
	fmt.Println("[STDERR] ", jr.stderr)
	fmt.Println("[STOP] ", jr.stop)
	fmt.Println("[COST] ", jr.stop.Sub(jr.start).Milliseconds(), "ms")
	fmt.Println("----------------------------------------")
}

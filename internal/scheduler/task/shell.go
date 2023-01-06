package task

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"syscall"
	"time"

	"github.com/bzp2010/schedule/internal/database/models"
	"github.com/reugn/go-quartz/quartz"
)

// ShellJob is a task used to execute external command lines
type ShellJob struct {
	Task

	Command   string
	Timeout   time.Duration
	Result    string
	JobStatus quartz.JobStatus
}

// NewShellJob is used to create ShellJob instances
func NewShellJob(task Task, command string, timeout time.Duration) *ShellJob {
	return &ShellJob{
		Task:      task,
		Command:   command,
		Timeout:   timeout,
		Result:    "",
		JobStatus: quartz.NA,
	}
}

// Execute is called by a Scheduler when the Trigger associated with this job fires.
func (sj *ShellJob) Execute() {
	job := models.Job{TaskID: sj.TaskID, TaskRuleID: sj.TaskRuleID, StartAt: time.Now()}

	ctx := context.Background()
	if sj.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), sj.Timeout)
		defer cancel()
	}

	cmd := exec.CommandContext(ctx, "sh", "-c", sj.Command)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				if status.Signaled() && status.Signal() == syscall.SIGKILL {
					// task timeout, the process is killed
					job.AddFlag(models.JobFlagTimeout)
				}
			}
		}
		sj.JobStatus = quartz.FAILURE
		sj.Result = err.Error()
	} else {
		sj.JobStatus = quartz.OK
	}

	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	sj.Result = outStr

	job.Stdout = outStr
	job.Stderr = errStr
	job.StopAt = time.Now()
	sj.SaveJob(job)
}

// Description returns the description of the Job.
func (sj *ShellJob) Description() string {
	return fmt.Sprintf("ShellJob: %s", sj.Command)
}

// Key returns the unique key for the Job.
func (sj *ShellJob) Key() int {
	return quartz.HashCode(sj.Description())
}

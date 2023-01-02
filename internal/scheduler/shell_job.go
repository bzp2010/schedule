package scheduler

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"syscall"
	"time"

	"github.com/reugn/go-quartz/quartz"
)

type shellJob struct {
	Command   string
	Timeout   time.Duration
	Result    string
	JobStatus quartz.JobStatus

	taskID     uint
	taskRuleID uint
}

func newShellJob(command string, timeout time.Duration) *shellJob {
	return &shellJob{
		Command:   command,
		Timeout:   timeout,
		Result:    "",
		JobStatus: quartz.NA,
	}
}

// Execute is called by a Scheduler when the Trigger associated with this job fires.
func (sj *shellJob) Execute() {
	record := newJobRecord(sj.taskID, sj.taskRuleID)

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
					record.timeout = true
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

	record.stdout = outStr
	record.stderr = errStr
	record.Save()
}

// Description returns the description of the Job.
func (sj *shellJob) Description() string {
	return fmt.Sprintf("ShellJob: %s", sj.Command)
}

// Key returns the unique key for the Job.
func (sj *shellJob) Key() int {
	return quartz.HashCode(sj.Description())
}

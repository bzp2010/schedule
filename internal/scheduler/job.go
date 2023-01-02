package scheduler

import (
	"encoding/json"
	"errors"

	"github.com/bzp2010/schedule/internal/database/models"
	"github.com/reugn/go-quartz/quartz"
)

func generateJob(task models.Task, rule models.TaskRule) (quartz.Job, error) {
	switch task.Type {
	case models.TaskTypeShell:
		cfg := models.TaskConfigurationShell{}
		err := json.Unmarshal(task.Configuration, &cfg)
		if err != nil {
			return nil, err
		}
		job := newShellJob(cfg.Command, cfg.Timeout)
		job.taskID = task.ID
		job.taskRuleID = rule.ID
		return job, nil
	}
	return nil, errors.New("unsupported task type")
}

func generateTrigger(rule models.TaskRule) (quartz.Trigger, error) {
	return quartz.NewCronTrigger(rule.Rule)
}

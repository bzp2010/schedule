package scheduler

import (
	"encoding/json"
	"errors"

	"github.com/bzp2010/schedule/internal/database/models"
	builtinTask "github.com/bzp2010/schedule/internal/scheduler/task"
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
		return builtinTask.NewShellJob(builtinTask.Task{
			TaskID:     task.ID,
			TaskRuleID: rule.ID,
		}, cfg.Command, cfg.Timeout), nil
	}
	return nil, errors.New("unsupported task type")
}

func generateTrigger(rule models.TaskRule) (quartz.Trigger, error) {
	return quartz.NewCronTrigger(rule.Rule)
}

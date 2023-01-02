package scheduler

import (
	"github.com/bzp2010/schedule/internal/database"
	"github.com/bzp2010/schedule/internal/database/models"
	"github.com/bzp2010/schedule/internal/log"
	"github.com/pkg/errors"
	"github.com/reugn/go-quartz/quartz"
)

// Scheduler for timed tasks
type Scheduler struct {
	scheduler *quartz.StdScheduler
}

var (
	scheduler *Scheduler
)

// GetScheduler gets the `Scheduler` singleton
func GetScheduler() *Scheduler {
	return scheduler
}

// SetupScheduler creates scheduler instance
func SetupScheduler() error {
	// create and start quartz std scheduler
	qs := quartz.NewStdScheduler()
	qs.Start()

	scheduler = &Scheduler{
		scheduler: qs,
	}

	// initializing tasks for scheduler
	scheduler.loadTask()

	return nil
}

// LoadTask List
func (s *Scheduler) loadTask() error {
	var (
		db    = database.GetDatabase()
		tasks []models.Task
	)

	// fetch tasks from database
	result := db.Preload("Rules").Find(&tasks)
	if err := result.Error; err != nil {
		log.GetLogger().Errorw("failed to load task from database", "error", err)
		return errors.Wrap(err, "failed to load task from database")
	}

	// add all tasks to gron scheduler
	for idxTask := range tasks {
		task := tasks[idxTask]
		if task.Status == models.StatusDisabled {
			continue
		}
		for idxRule := range task.Rules {
			rule := task.Rules[idxRule]

			job, err := generateJob(task, rule)
			if err != nil {
				log.GetLogger().Errorw("failed to generate job", "error", err)
				return err
			}

			trigger, err := generateTrigger(rule)
			if err != nil {
				log.GetLogger().Errorw("failed to generate trigger", "error", err)
				return err
			}

			s.scheduler.ScheduleJob(job, trigger)
		}
	}

	return nil
}

// ReloadTask of scheduler
func (s *Scheduler) ReloadTask() {
	s.scheduler.Clear()
	s.loadTask()
}

// Stop scheduler
func (s *Scheduler) Stop() {
	s.scheduler.Stop()
}

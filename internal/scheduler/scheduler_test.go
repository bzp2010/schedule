package scheduler_test

import (
	"database/sql"
	"encoding/json"
	"testing"
	"time"

	"github.com/bzp2010/schedule/internal/config"
	"github.com/bzp2010/schedule/internal/database"
	"github.com/bzp2010/schedule/internal/database/models"
	"github.com/bzp2010/schedule/internal/log"
	"github.com/stretchr/testify/assert"
	"gorm.io/datatypes"
)

func TestCreateATask(t *testing.T) {
	// config
	cfg := config.NewDefaultConfig()

	assert.NoError(t, config.SetupConfig(&cfg, "../../config/config.yaml"), "failed to setup config")

	cfg.DSN = "sqlite://../../data.db"
	assert.NoError(t, log.SetupLogger(cfg), "failed to setup logger")
	assert.NoError(t, database.SetupDatabase(cfg), "failed to setup database")

	assert.NoError(t, database.Migrate(database.GetDatabase()))

	d, err := json.Marshal(models.TaskConfigurationShell{
		Command: "ls -la /; sleep 1; ls -la /",
		Timeout: 500 * time.Millisecond,
	})
	assert.NoError(t, err)

	task := models.Task{
		Name:          "test_task",
		Type:          models.TaskTypeShell,
		Configuration: datatypes.JSON(d),
		Rules: []models.TaskRule{
			{
				Description: sql.NullString{
					String: "test_task_rule_1",
					Valid:  true,
				},
				Rule:   "0/2 * * * * *",
				Status: models.StatusEnabled,
			},
			/* {
				Description: sql.NullString{
					String: "test_task_rule_2",
					Valid:  true,
				},
				Rule:   "0/5 * * * * *",
				Status: models.StatusEnabled,
			}, */
		},
		Status: models.StatusEnabled,
	}

	result := database.GetDatabase().Create(&task)

	assert.NoError(t, result.Error)
	assert.Equal(t, int64(1), result.RowsAffected)
}

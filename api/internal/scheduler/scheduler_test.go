package scheduler_test

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/bzp2010/schedule/internal/config"
	"github.com/bzp2010/schedule/internal/database"
	"github.com/bzp2010/schedule/internal/database/models"
	"github.com/bzp2010/schedule/internal/log"
	"github.com/stretchr/testify/assert"
	"gorm.io/datatypes"
)

func setup(t *testing.T) {
	// config
	cfg := config.NewDefaultConfig()

	assert.NoError(t, config.SetupConfig(&cfg, "../../config/config.yaml"), "failed to setup config")

	cfg.DSN = "sqlite://../../data.db"
	assert.NoError(t, log.SetupLogger(cfg), "failed to setup logger")
	assert.NoError(t, database.SetupDatabase(cfg), "failed to setup database")

	assert.NoError(t, database.Migrate(database.GetDatabase()))
}

func TestCreateShellTask(t *testing.T) {
	setup(t)

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
			{
				Description: sql.NullString{
					String: "test_task_rule_2",
					Valid:  true,
				},
				Rule:   "0/5 * * * * *",
				Status: models.StatusEnabled,
			},
		},
		Status: models.StatusEnabled,
	}

	result := database.GetDatabase().Create(&task)
	assert.NoError(t, result.Error)
	assert.Equal(t, int64(1), result.RowsAffected)

}

func TestCreateWebhookTask(t *testing.T) {
	setup(t)

	d, err := json.Marshal(models.TaskConfigurationWebhook{
		URL:    "https://1.1.1.1",
		Method: http.MethodGet,
	})
	assert.NoError(t, err)

	task := models.Task{
		Name:          "test_task_2",
		Type:          models.TaskTypeWebhook,
		Configuration: datatypes.JSON(d),
		Rules: []models.TaskRule{
			{
				Description: sql.NullString{
					String: "test_webhook_1",
					Valid:  true,
				},
				Rule:   "0/5 * * * * *",
				Status: models.StatusEnabled,
			},
		},
		Status: models.StatusEnabled,
	}

	result := database.GetDatabase().Create(&task)
	assert.NoError(t, result.Error)
	assert.Equal(t, int64(1), result.RowsAffected)
}

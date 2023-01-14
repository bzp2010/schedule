package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"fmt"

	"github.com/bzp2010/schedule/internal/database"
	"github.com/bzp2010/schedule/internal/database/models"
	"github.com/bzp2010/schedule/internal/handler/graphql/generated"
)

// Task is the resolver for the task field.
func (r *queryResolver) Task(ctx context.Context, id int64) (*models.Task, error) {
	var task models.Task
	result := database.GetDatabase().Where("id = ?", id).Find(&task)
	if err := result.Error; err != nil {
		return nil, err
	}
	if result.RowsAffected <= 0 {
		return nil, fmt.Errorf("task does not exist: id %d", id)
	}
	return &task, nil
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context, limit *int, offset *int) ([]*models.Task, error) {
	var tasks []*models.Task
	result := database.GetDatabase().Offset(*offset).Limit(*limit).Find(&tasks)
	if err := result.Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

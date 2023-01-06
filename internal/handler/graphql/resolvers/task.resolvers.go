package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"

	"github.com/bzp2010/schedule/internal/database/models"
	"github.com/bzp2010/schedule/internal/handler/graphql/generated"
	"github.com/icza/gog"
)

// ID is the resolver for the id field.
func (r *taskResolver) ID(ctx context.Context, obj *models.Task) (*int64, error) {
	return gog.Ptr(int64(obj.ID)), nil
}

// Type is the resolver for the type field.
func (r *taskResolver) Type(ctx context.Context, obj *models.Task) (*models.TaskType, error) {
	return gog.Ptr(obj.Type), nil
}

// Configuration is the resolver for the configuration field.
func (r *taskResolver) Configuration(ctx context.Context, obj *models.Task) (*string, error) {
	return gog.Ptr(obj.Configuration.String()), nil
}

// LastRunningAt is the resolver for the last_running_at field.
func (r *taskResolver) LastRunningAt(ctx context.Context, obj *models.Task) (*int64, error) {
	return gog.Ptr(
		gog.If(
			obj.LastRunningAt.Valid,
			obj.LastRunningAt.Time.UnixMilli(),
			int64(0),
		),
	), nil
}

// CreatedAt is the resolver for the created_at field.
func (r *taskResolver) CreatedAt(ctx context.Context, obj *models.Task) (*int64, error) {
	return gog.Ptr(obj.CreatedAt.UnixMilli()), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *taskResolver) UpdatedAt(ctx context.Context, obj *models.Task) (*int64, error) {
	return gog.Ptr(obj.UpdatedAt.UnixMilli()), nil
}

// Task returns generated.TaskResolver implementation.
func (r *Resolver) Task() generated.TaskResolver { return &taskResolver{r} }

type taskResolver struct{ *Resolver }

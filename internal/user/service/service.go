package service

import (
	"context"

	"go-service/internal/user/model"
)

type UserService interface {
	All(ctx context.Context) ([]model.User, error)
	Load(ctx context.Context, id model.UserId) (*model.User, error)
	Create(ctx context.Context, user *model.User) (int64, error)
	Update(ctx context.Context, user *model.User) (int64, error)
	Patch(ctx context.Context, user map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id model.UserId) (int64, error)
	Search(ctx context.Context, filter *model.UserFilter, limit int64, offset int64) ([]model.User, int64, error)
}

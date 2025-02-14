package service

import (
	"context"

	"go-service/internal/user/model"
	"go-service/internal/user/repository"
)

func NewUserService(repository repository.UserRepository) *UserUseCase {
	return &UserUseCase{repository: repository}
}

type UserUseCase struct {
	repository repository.UserRepository
}

func (s *UserUseCase) All(ctx context.Context) ([]model.User, error) {
	return s.repository.All(ctx)
}
func (s *UserUseCase) Load(ctx context.Context, id model.UserId) (*model.User, error) {
	return s.repository.Load(ctx, id)
}
func (s *UserUseCase) Create(ctx context.Context, user *model.User) (int64, error) {
	return s.repository.Create(ctx, user)
}
func (s *UserUseCase) Update(ctx context.Context, user *model.User) (int64, error) {
	return s.repository.Update(ctx, user)
}
func (s *UserUseCase) Patch(ctx context.Context, user map[string]interface{}) (int64, error) {
	return s.repository.Patch(ctx, user)
}
func (s *UserUseCase) Delete(ctx context.Context, id model.UserId) (int64, error) {
	return s.repository.Delete(ctx, id)
}
func (s *UserUseCase) Search(ctx context.Context, filter *model.UserFilter, limit int64, offset int64) ([]model.User, int64, error) {
	return s.repository.Search(ctx, filter, limit, offset)
}

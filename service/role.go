package service

import (
	"context"
	"github.com/chincharovpc/goarch/model"
	"github.com/chincharovpc/goarch/repository"
)

type RoleService interface {
	FindByID(ctx context.Context, id int) (*model.Role, error)
	FindByName(ctx context.Context, name string) (*model.Role, error)
	Create(ctx context.Context, input CreateRoleInput) (*model.Role, error)
}

type Role struct {
	roleRepository repository.RoleRepository
}

func NewRoleService(roleRepository repository.RoleRepository) *Role {
	return &Role{roleRepository: roleRepository}
}

func (s *Role) FindByID(ctx context.Context, id int) (*model.Role, error) {
	return s.roleRepository.FindByID(ctx, id)
}

func (s *Role) FindByName(ctx context.Context, name string) (*model.Role, error) {
	return s.roleRepository.FindByName(ctx, name)
}

type CreateRoleInput struct {
	Name        string
	DisplayName string
	Description string
}

func (s *Role) Create(ctx context.Context, input CreateRoleInput) (*model.Role, error) {
	return s.roleRepository.Create(ctx, &model.Role{
		Name:        input.Name,
		DisplayName: input.DisplayName,
		Description: input.Description,
	})
}

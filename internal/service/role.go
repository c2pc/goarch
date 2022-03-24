package service

import (
	"context"
	"github.com/chincharovpc/goarch/internal/core"
	"github.com/chincharovpc/goarch/internal/repository"
)

type IRoleService interface {
	FindByID(ctx context.Context, id int) (*core.Role, error)
	FindByName(ctx context.Context, name string) (*core.Role, error)
	Create(ctx context.Context, input CreateRoleInput) (*core.Role, error)
}

type Role struct {
	roleRepository repository.IRoleRepository
}

func NewRoleService(roleRepository repository.IRoleRepository) *Role {
	return &Role{roleRepository: roleRepository}
}

func (s *Role) FindByID(ctx context.Context, id int) (*core.Role, error) {
	return s.roleRepository.FindByID(ctx, id)
}

func (s *Role) FindByName(ctx context.Context, name string) (*core.Role, error) {
	return s.roleRepository.FindByName(ctx, name)
}

type CreateRoleInput struct {
	Name        string
	DisplayName string
	Description string
}

func (s *Role) Create(ctx context.Context, input CreateRoleInput) (*core.Role, error) {
	return s.roleRepository.Create(ctx, &repository.CreateRoleInput{
		Name:        input.Name,
		DisplayName: input.DisplayName,
		Description: input.Description,
	})
}

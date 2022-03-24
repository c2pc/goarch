package repository

import (
	"context"
	"github.com/chincharovpc/goarch/internal/core"
	"github.com/chincharovpc/goarch/pkg/apperr"
	"github.com/go-pg/pg/v10"
	"go.uber.org/zap"
	"net/http"
)

type IRoleRepository interface {
	FindByID(ctx context.Context, id int) (*core.Role, error)
	FindByName(ctx context.Context, name string) (*core.Role, error)
	Create(ctx context.Context, m *CreateRoleInput) (*core.Role, error)
}

type RoleRepository struct {
	db     *pg.DB
	logger *zap.Logger
}

func NewRoleRepository(db *pg.DB, logger *zap.Logger) *RoleRepository {
	return &RoleRepository{
		db:     db,
		logger: logger,
	}
}

func (r *RoleRepository) FindByID(ctx context.Context, id int) (*core.Role, error) {
	role := new(core.Role)

	sql := `SELECT * FROM ` + core.RoleTable + ` WHERE id = ?`
	_, err := r.db.QueryOneContext(ctx, role, sql, id)
	if err != nil {
		r.logger.Warn("RoleRepository Error", zap.String("Error:", err.Error()))
		return nil, apperr.NotFound
	}

	return role, nil
}

func (r *RoleRepository) FindByName(ctx context.Context, name string) (*core.Role, error) {
	role := new(core.Role)

	sql := `SELECT * FROM ` + core.RoleTable + ` WHERE name = ?`
	_, err := r.db.QueryOneContext(ctx, role, sql, name)
	if err != nil {
		r.logger.Warn("RoleRepository Error", zap.String("Error:", err.Error()))
		return nil, apperr.NotFound
	}

	return role, nil
}

type CreateRoleInput struct {
	Name        string
	DisplayName string
	Description string
}

func (r *RoleRepository) Create(ctx context.Context, m *CreateRoleInput) (*core.Role, error) {
	role := &core.Role{
		Name:        m.Name,
		DisplayName: m.DisplayName,
		Description: m.Description,
	}

	sql := `SELECT id FROM ` + core.RoleTable + ` WHERE name = ?`
	res, err := r.db.QueryContext(ctx, role, sql, m.Name)
	if err != nil {
		r.logger.Warn("RoleRepository Error", zap.String("Error:", err.Error()))
		return nil, apperr.DB
	}
	if res.RowsReturned() != 0 {
		return nil, apperr.New(http.StatusBadRequest, "Role already exists.")
	}

	_, err = r.db.WithContext(ctx).Model(m).Insert()
	if err != nil {
		r.logger.Warn("RoleRepository Error", zap.String("Error:", err.Error()))
		return nil, apperr.DB
	}

	return role, nil
}

package repository

import (
	"context"
	"github.com/chincharovpc/goarch/apperr"
	"github.com/chincharovpc/goarch/model"
	"github.com/go-pg/pg/v10"
	"go.uber.org/zap"
	"net/http"
)

type RoleRepository interface {
	FindByID(ctx context.Context, id int) (*model.Role, error)
	FindByName(ctx context.Context, name string) (*model.Role, error)
	Create(ctx context.Context, role *model.Role) (*model.Role, error)
}

type Role struct {
	db     *pg.DB
	logger *zap.Logger
	model  *model.Role
}

func NewRoleRepository(db *pg.DB, logger *zap.Logger) *Role {
	return &Role{
		db:     db,
		logger: logger,
	}
}

func (r *Role) FindByID(ctx context.Context, id int) (*model.Role, error) {
	role := new(model.Role)

	sql := `SELECT * FROM roles WHERE id = ?`
	_, err := r.db.QueryOneContext(ctx, role, sql, id)
	if err != nil {
		r.logger.Warn("RoleRepository Error", zap.String("Error:", err.Error()))
		return nil, apperr.NotFound
	}

	return role, nil
}

func (r *Role) FindByName(ctx context.Context, name string) (*model.Role, error) {
	role := new(model.Role)

	sql := `SELECT * FROM roles WHERE name = ?`
	_, err := r.db.QueryOneContext(ctx, role, sql, name)
	if err != nil {
		r.logger.Warn("RoleRepository Error", zap.String("Error:", err.Error()))
		return nil, apperr.NotFound
	}

	return role, nil
}

func (r *Role) Create(ctx context.Context, m *model.Role) (*model.Role, error) {
	role := new(model.Role)

	sql := `SELECT id FROM roles WHERE name = ?`
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

	return m, nil
}

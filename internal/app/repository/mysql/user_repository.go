package mysql

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lecterkn/yumemiban_backend/internal/app/entity"
	"github.com/lecterkn/yumemiban_backend/internal/app/port"
)

type UserRepositoryImpl struct {
	database *sqlx.DB
}

func NewUserRepositoryImpl(database *sqlx.DB) port.UserRepository {
	return &UserRepositoryImpl{
		database,
	}
}

func (r *UserRepositoryImpl) Create(ctx context.Context, userEntity *entity.UserEntity) error {
	query := `
        INSERT INTO users(id, name, password, created_at, updated_at)
        VALUES(:id, :name, :password, :createdAt, :updatedAt)
    `
	return RunInTx(ctx, r.database, func(tx *sqlx.Tx) error {
		_, err := tx.NamedExec(query, map[string]any{
			"id":        userEntity.Id[:],
			"name":      userEntity.Name,
			"password":  userEntity.Password,
			"createdAt": userEntity.CreatedAt,
			"updatedAt": userEntity.UpdatedAt,
		})
		return err
	})
}

// TODO
func (r *UserRepositoryImpl) FindById(ctx context.Context, id uuid.UUID) (*entity.UserEntity, error) {
	return nil, nil
}

// TODO
func (r *UserRepositoryImpl) FindByName(ctx context.Context, name string) (*entity.UserEntity, error) {
	return nil, nil
}

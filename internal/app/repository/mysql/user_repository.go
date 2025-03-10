package mysql

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lecterkn/yumemiban_backend/internal/app/entity"
	"github.com/lecterkn/yumemiban_backend/internal/app/port"
	"github.com/lecterkn/yumemiban_backend/internal/app/repository/mysql/model"
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

func (r *UserRepositoryImpl) FindById(ctx context.Context, id uuid.UUID) (*entity.UserEntity, error) {
	query := `
        SELECT id, name, password, created_at, updated_at
        FROM users
        WHERE id = ?
        LIMIT 1
    `
	userEntity := entity.UserEntity{}
	err := RunInTx(ctx, r.database, func(tx *sqlx.Tx) error {
		userModel := model.UserModel{}
		err := tx.Get(&userModel, query, id[:])
		if err != nil {
			return err
		}
		userEntityPtr, err := r.toEntity(&userModel)
		if err != nil {
			return err
		}
		userEntity = *userEntityPtr
		return nil
	})
	return &userEntity, err
}

func (r *UserRepositoryImpl) FindByName(ctx context.Context, name string) (*entity.UserEntity, error) {
	query := `
        SELECT id, name, password, created_at, updated_at
        FROM users
        WHERE name = ?
        LIMIT 1
    `
	userEntity := entity.UserEntity{}
	err := RunInTx(ctx, r.database, func(tx *sqlx.Tx) error {
		userModel := model.UserModel{}
		err := tx.Select(&userModel, query, name)
		if err != nil {
			return err
		}
		userEntityPtr, err := r.toEntity(&userModel)
		if err != nil {
			return err
		}
		userEntity = *userEntityPtr
		return nil
	})
	return &userEntity, err
}

func (r *UserRepositoryImpl) toEntity(userModel *model.UserModel) (*entity.UserEntity, error) {
	id, err := uuid.FromBytes(userModel.Id)
	if err != nil {
		return nil, err
	}
	return &entity.UserEntity{
		Id:        id,
		Name:      userModel.Name,
		Password:  userModel.Password,
		CreatedAt: userModel.CreatedAt,
		UpdatedAt: userModel.UpdatedAt,
	}, nil
}

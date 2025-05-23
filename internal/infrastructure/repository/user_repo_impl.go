package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	domainErrors "github.com/lyonnee/go-template/internal/domain/errors"

	"github.com/lyonnee/go-template/internal/domain/entity"
	"github.com/lyonnee/go-template/internal/domain/repository"
	"github.com/lyonnee/go-template/internal/infrastructure/repository/model"
	"github.com/lyonnee/go-template/pkg/persistence"
)

// UserRepositoryImpl 用户存储库实现
type UserRepositoryImpl struct {
	executer persistence.Executer
}

// NewUserRepository 创建一个新的用户存储库实例
func NewUserRepository() repository.UserRepository {
	return &UserRepositoryImpl{
		executer: nil, // 初始化时没有执行器，需要通过 WithExecuter 设置
	}
}

// WithExecuter 设置特定的执行器，返回一个新的存储库实例
func (r *UserRepositoryImpl) WithExecuter(executer persistence.Executer) repository.UserRepository {
	return &UserRepositoryImpl{
		executer: executer,
	}
}

// 获取当前执行器，如果未设置则返回错误
func (r *UserRepositoryImpl) getExecuter() (persistence.Executer, error) {
	if r.executer == nil {
		return nil, errors.New("executer not set, use WithExecuter() to set an executer")
	}
	return r.executer, nil
}

// FindById 根据ID查找用户
func (r *UserRepositoryImpl) FindById(ctx context.Context, userId int64) (*entity.User, error) {
	executer, err := r.getExecuter()
	if err != nil {
		return nil, err
	}

	query := `
		SELECT id, created_at, updated_at, username, password, email, phone, is_deleted, deleted_at 
		FROM users 
		WHERE id = $1 AND is_deleted = false
	`

	var userModel model.UserModel
	err = executer.QueryRowxContext(ctx, query, userId).Scan(
		&userModel.ID,
		&userModel.CreatedAt,
		&userModel.UpdatedAt,
		&userModel.Username,
		&userModel.Password,
		&userModel.Email,
		&userModel.Phone,
		&userModel.IsDeleted,
		&userModel.DeletedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainErrors.ErrUserNotFound
		}
		return nil, err
	}

	return r.modelToEntity(&userModel), nil
}

// Create 创建新用户
func (r *UserRepositoryImpl) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	executer, err := r.getExecuter()
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, domainErrors.ErrInvalidUserInput
	}

	// 检查用户名、邮箱和手机号是否已存在
	exists, err := r.checkUserFieldsExist(ctx, user)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("user with these details already exists")
	}

	now := time.Now().Unix()
	query := `
		INSERT INTO users (created_at, updated_at, username, password, email, phone, is_deleted, deleted_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`

	var id int64
	err = executer.QueryRowxContext(ctx, query,
		now,
		now,
		user.Username,
		user.Password,
		user.Email,
		user.Phone,
		false,
		0,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	user.ID = id
	user.CreatedAt = now
	user.UpdatedAt = now
	user.IsDeleted = false
	user.DeletedAt = 0

	return user, nil
}

// Update 更新用户信息
func (r *UserRepositoryImpl) Update(ctx context.Context, user *entity.User) error {
	executer, err := r.getExecuter()
	if err != nil {
		return err
	}

	if user == nil || user.ID == 0 {
		return domainErrors.ErrInvalidUserInput
	}

	// 检查用户是否存在
	_, err = r.FindById(ctx, user.ID)
	if err != nil {
		return err
	}

	now := time.Now().Unix()
	query := `
		UPDATE users 
		SET updated_at = $1, username = $2, password = $3, email = $4, phone = $5 
		WHERE id = $6
	`

	_, err = executer.ExecContext(ctx, query,
		now,
		user.Username,
		user.Password,
		user.Email,
		user.Phone,
		user.ID,
	)

	if err != nil {
		return err
	}

	user.UpdatedAt = now
	return nil
}

// Delete 删除用户（软删除）
func (r *UserRepositoryImpl) Delete(ctx context.Context, userId int64) error {
	executer, err := r.getExecuter()
	if err != nil {
		return err
	}

	now := time.Now().Unix()
	query := `
		UPDATE users 
		SET is_deleted = true, deleted_at = $1 
		WHERE id = $2
	`

	result, err := executer.ExecContext(ctx, query, now, userId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return domainErrors.ErrUserNotFound
	}

	return nil
}

// FindByUsername 根据用户名查找用户
func (r *UserRepositoryImpl) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	executer, err := r.getExecuter()
	if err != nil {
		return nil, err
	}

	query := `
		SELECT id, created_at, updated_at, username, password, email, phone, is_deleted, deleted_at 
		FROM users 
		WHERE username = $1 AND is_deleted = false
	`

	var userModel model.UserModel
	err = executer.QueryRowxContext(ctx, query, username).Scan(
		&userModel.ID,
		&userModel.CreatedAt,
		&userModel.UpdatedAt,
		&userModel.Username,
		&userModel.Password,
		&userModel.Email,
		&userModel.Phone,
		&userModel.IsDeleted,
		&userModel.DeletedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainErrors.ErrUserNotFound
		}
		return nil, err
	}

	return r.modelToEntity(&userModel), nil
}

// FindByEmail 根据邮箱查找用户
func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	executer, err := r.getExecuter()
	if err != nil {
		return nil, err
	}

	query := `
		SELECT id, created_at, updated_at, username, password, email, phone, is_deleted, deleted_at 
		FROM users 
		WHERE email = $1 AND is_deleted = false
	`

	var userModel model.UserModel
	err = executer.QueryRowxContext(ctx, query, email).Scan(
		&userModel.ID,
		&userModel.CreatedAt,
		&userModel.UpdatedAt,
		&userModel.Username,
		&userModel.Password,
		&userModel.Email,
		&userModel.Phone,
		&userModel.IsDeleted,
		&userModel.DeletedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainErrors.ErrUserNotFound
		}
		return nil, err
	}

	return r.modelToEntity(&userModel), nil
}

// FindByPhone 根据手机号查找用户
func (r *UserRepositoryImpl) FindByPhone(ctx context.Context, phone string) (*entity.User, error) {
	executer, err := r.getExecuter()
	if err != nil {
		return nil, err
	}

	query := `
		SELECT id, created_at, updated_at, username, password, email, phone, is_deleted, deleted_at 
		FROM users 
		WHERE phone = $1 AND is_deleted = false
	`

	var userModel model.UserModel
	err = executer.QueryRowxContext(ctx, query, phone).Scan(
		&userModel.ID,
		&userModel.CreatedAt,
		&userModel.UpdatedAt,
		&userModel.Username,
		&userModel.Password,
		&userModel.Email,
		&userModel.Phone,
		&userModel.IsDeleted,
		&userModel.DeletedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainErrors.ErrUserNotFound
		}
		return nil, err
	}

	return r.modelToEntity(&userModel), nil
}

// UpdateUsername 更新用户名
func (r *UserRepositoryImpl) UpdateUsername(ctx context.Context, user *entity.User) error {
	executer, err := r.getExecuter()
	if err != nil {
		return err
	}

	if user == nil || user.ID == 0 || user.Username == "" {
		return domainErrors.ErrInvalidUserInput
	}

	// 检查用户名是否已被使用
	existingUser, err := r.FindByUsername(ctx, user.Username)
	if err != nil && !errors.Is(err, domainErrors.ErrUserNotFound) {
		return err
	}
	if existingUser != nil && existingUser.ID != user.ID {
		return domainErrors.ErrUsernameTaken
	}

	now := time.Now().Unix()
	query := `
		UPDATE users 
		SET updated_at = $1, username = $2 
		WHERE id = $3
	`

	result, err := executer.ExecContext(ctx, query, now, user.Username, user.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return domainErrors.ErrUserNotFound
	}

	user.UpdatedAt = now
	return nil
}

// UpdatePassword 更新密码
func (r *UserRepositoryImpl) UpdatePassword(ctx context.Context, user *entity.User) error {
	executer, err := r.getExecuter()
	if err != nil {
		return err
	}

	if user == nil || user.ID == 0 || user.Password == "" {
		return domainErrors.ErrInvalidUserInput
	}

	now := time.Now().Unix()
	query := `
		UPDATE users 
		SET updated_at = $1, password = $2 
		WHERE id = $3
	`

	result, err := executer.ExecContext(ctx, query, now, user.Password, user.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return domainErrors.ErrUserNotFound
	}

	user.UpdatedAt = now
	return nil
}

// UpdateEmail 更新邮箱
func (r *UserRepositoryImpl) UpdateEmail(ctx context.Context, user *entity.User) error {
	executer, err := r.getExecuter()
	if err != nil {
		return err
	}

	if user == nil || user.ID == 0 || user.Email == "" {
		return domainErrors.ErrInvalidUserInput
	}

	// 检查邮箱是否已被使用
	existingUser, err := r.FindByEmail(ctx, user.Email)
	if err != nil && !errors.Is(err, domainErrors.ErrUserNotFound) {
		return err
	}
	if existingUser != nil && existingUser.ID != user.ID {
		return domainErrors.ErrEmailTaken
	}

	now := time.Now().Unix()
	query := `
		UPDATE users 
		SET updated_at = $1, email = $2 
		WHERE id = $3
	`

	result, err := executer.ExecContext(ctx, query, now, user.Email, user.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return domainErrors.ErrUserNotFound
	}

	user.UpdatedAt = now
	return nil
}

// UpdatePhone 更新手机号
func (r *UserRepositoryImpl) UpdatePhone(ctx context.Context, user *entity.User) error {
	executer, err := r.getExecuter()
	if err != nil {
		return err
	}

	if user == nil || user.ID == 0 || user.Phone == "" {
		return domainErrors.ErrInvalidUserInput
	}

	// 检查手机号是否已被使用
	existingUser, err := r.FindByPhone(ctx, user.Phone)
	if err != nil && !errors.Is(err, domainErrors.ErrUserNotFound) {
		return err
	}
	if existingUser != nil && existingUser.ID != user.ID {
		return domainErrors.ErrPhoneTaken
	}

	now := time.Now().Unix()
	query := `
		UPDATE users 
		SET updated_at = $1, phone = $2 
		WHERE id = $3
	`

	result, err := executer.ExecContext(ctx, query, now, user.Phone, user.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return domainErrors.ErrUserNotFound
	}

	user.UpdatedAt = now
	return nil
}

// 检查用户字段是否存在
func (r *UserRepositoryImpl) checkUserFieldsExist(ctx context.Context, user *entity.User) (bool, error) {
	if user.Username != "" {
		u, err := r.FindByUsername(ctx, user.Username)
		if err != nil && !errors.Is(err, domainErrors.ErrUserNotFound) {
			return false, err
		}
		if u != nil {
			return true, domainErrors.ErrUsernameTaken
		}
	}

	if user.Email != "" {
		u, err := r.FindByEmail(ctx, user.Email)
		if err != nil && !errors.Is(err, domainErrors.ErrUserNotFound) {
			return false, err
		}
		if u != nil {
			return true, domainErrors.ErrEmailTaken
		}
	}

	if user.Phone != "" {
		u, err := r.FindByPhone(ctx, user.Phone)
		if err != nil && !errors.Is(err, domainErrors.ErrUserNotFound) {
			return false, err
		}
		if u != nil {
			return true, domainErrors.ErrPhoneTaken
		}
	}

	return false, nil
}

func (r *UserRepositoryImpl) entityToModel(user *entity.User) *model.UserModel {
	return &model.UserModel{
		SoftDelete_BaseModel: model.SoftDelete_BaseModel{
			BaseModel: model.BaseModel{
				ID:        user.ID,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			},
			IsDeleted: user.IsDeleted,
			DeletedAt: user.DeletedAt,
		},
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Phone:    user.Phone,
	}
}

func (r *UserRepositoryImpl) modelToEntity(userModel *model.UserModel) *entity.User {
	return &entity.User{
		ID:        userModel.ID,
		CreatedAt: userModel.CreatedAt,
		UpdatedAt: userModel.UpdatedAt,
		Username:  userModel.Username,
		Password:  userModel.Password,
		Email:     userModel.Email,
		Phone:     userModel.Phone,
		IsDeleted: userModel.IsDeleted,
		DeletedAt: userModel.DeletedAt,
	}
}

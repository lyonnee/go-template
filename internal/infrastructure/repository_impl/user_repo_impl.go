package repository_impl

import (
	"context"
	"database/sql"
	"errors"
	"time"

	domainErrors "github.com/lyonnee/go-template/internal/domain/errors"

	"github.com/lyonnee/go-template/internal/domain/entity"
	"github.com/lyonnee/go-template/internal/domain/repository"
	"github.com/lyonnee/go-template/internal/infrastructure/di"
	"github.com/lyonnee/go-template/internal/infrastructure/log"
	"github.com/lyonnee/go-template/internal/infrastructure/repository_impl/model"
)

// UserRepositoryImpl 用户存储库实现
type UserRepositoryImpl struct {
	executor repository.Executor
	logger   log.Logger
}

// NewUserRepository 创建一个新的用户存储库实例
func NewUserRepository() (repository.UserRepository, error) {
	return &UserRepositoryImpl{
		executor: nil, // 初始化时没有执行器，需要通过 WithExecutor 设置
		logger:   di.GetService[log.Logger](),
	}, nil
}

// WithExecutor 设置特定的执行器，返回一个新的存储库实例
func (r *UserRepositoryImpl) WithExecutor(executor repository.Executor) repository.UserRepository {
	return &UserRepositoryImpl{
		executor: executor,
		logger:   r.logger,
	}
}

// 获取当前执行器，如果未设置则返回错误
func (r *UserRepositoryImpl) getExecutor() (repository.Executor, error) {
	if r.executor == nil {
		return nil, errors.New("executor not set, use WithExecutor() to set an executor")
	}
	return r.executor, nil
}

// FindById 根据ID查找用户
func (r *UserRepositoryImpl) FindById(ctx context.Context, userId int64) (*entity.User, error) {
	r.logger.DebugKV("Finding user by ID", "userId", userId)

	executor, err := r.getExecutor()
	if err != nil {
		r.logger.ErrorKV("Failed to get executor", "error", err)
		return nil, err
	}

	query := `
		SELECT id, created_at, updated_at, username, pwd_secret, email, phone, deleted_at 
		FROM users 
		WHERE id = $1 AND deleted_at = 0
	`

	var userModel model.UserModel
	err = executor.QueryRowxContext(ctx, query, userId).Scan(
		&userModel.ID,
		&userModel.CreatedAt,
		&userModel.UpdatedAt,
		&userModel.Username,
		&userModel.PwdSecret,
		&userModel.Email,
		&userModel.Phone,
		&userModel.DeletedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			r.logger.DebugKV("User not found", "userId", userId)
			return nil, domainErrors.ErrUserNotFound
		}
		r.logger.ErrorKV("Failed to find user by ID", "userId", userId, "error", err)
		return nil, err
	}

	r.logger.DebugKV("User found successfully", "userId", userId, "username", userModel.Username)
	return r.modelToEntity(&userModel), nil
}

// Create 创建新用户
func (r *UserRepositoryImpl) Create(ctx context.Context, user *entity.User) error {
	r.logger.InfoKV("Creating new user", "username", user.Username, "email", user.Email)

	executor, err := r.getExecutor()
	if err != nil {
		r.logger.ErrorKV("Failed to get executor", "error", err)
		return err
	}

	if user == nil {
		r.logger.Error("Invalid user input: user is nil")
		return domainErrors.ErrInvalidUserInput
	}

	// 检查用户名、邮箱和手机号是否已存在
	exists, err := r.checkUserFieldsExist(ctx, user)
	if err != nil {
		r.logger.ErrorKV("Failed to check user fields existence",
			"username", user.Username,
			"email", user.Email,
			"error", err)
		return err
	}
	if exists {
		r.logger.WarnKV("User with these details already exists",
			"username", user.Username,
			"email", user.Email)
		return errors.New("user with these details already exists")
	}

	now := time.Now().Unix()
	query := `
		INSERT INTO users (created_at, updated_at, username, pwd_secret, email, phone, deleted_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`

	var id int64
	err = executor.QueryRowxContext(ctx, query,
		now,
		now,
		user.Username,
		user.PwdSecret,
		user.Email,
		user.Phone,
		0,
	).Scan(&id)

	if err != nil {
		r.logger.ErrorKV("Failed to create user",
			"username", user.Username,
			"email", user.Email,
			"error", err)
		return err
	}

	user.ID = id
	user.CreatedAt = now
	user.UpdatedAt = now
	user.DeletedAt = 0

	r.logger.InfoKV("User created successfully",
		"userId", id,
		"username", user.Username)

	return nil
}

// Update 更新用户信息
func (r *UserRepositoryImpl) Update(ctx context.Context, user *entity.User) error {
	executor, err := r.getExecutor()
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
		SET updated_at = $1, username = $2, pwd_secret = $3, email = $4, phone = $5 
		WHERE id = $6
	`

	_, err = executor.ExecContext(ctx, query,
		now,
		user.Username,
		user.PwdSecret,
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
	executor, err := r.getExecutor()
	if err != nil {
		return err
	}

	now := time.Now().Unix()
	query := `
		UPDATE users 
		SET deleted_at = $1 
		WHERE id = $2
	`

	result, err := executor.ExecContext(ctx, query, now, userId)
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
	executor, err := r.getExecutor()
	if err != nil {
		return nil, err
	}

	query := `
		SELECT id, created_at, updated_at, username, pwd_secret, email, phone, deleted_at 
		FROM users 
		WHERE username = $1 AND deleted_at = 0
	`

	var userModel model.UserModel
	// err = executor.QueryRowxContext(ctx, query, username).Scan(
	// 	&userModel.ID,
	// 	&userModel.CreatedAt,
	// 	&userModel.UpdatedAt,
	// 	&userModel.Username,
	// 	&userModel.PwdSecret,
	// 	&userModel.Email,
	// 	&userModel.Phone,
	// 	&userModel.DeletedAt,
	// )
	err = executor.QueryRowxContext(ctx, query, username).StructScan(&userModel)

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
	executor, err := r.getExecutor()
	if err != nil {
		return nil, err
	}

	query := `
		SELECT id, created_at, updated_at, username, pwd_secret, email, phone, deleted_at 
		FROM users 
		WHERE email = $1 AND deleted_at = 0
	`

	var userModel model.UserModel
	err = executor.QueryRowxContext(ctx, query, email).Scan(
		&userModel.ID,
		&userModel.CreatedAt,
		&userModel.UpdatedAt,
		&userModel.Username,
		&userModel.PwdSecret,
		&userModel.Email,
		&userModel.Phone,
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
	executor, err := r.getExecutor()
	if err != nil {
		return nil, err
	}

	query := `
		SELECT id, created_at, updated_at, username, pwd_secret, email, phone, deleted_at 
		FROM users 
		WHERE phone = $1 AND deleted_at = 0
	`

	var userModel model.UserModel
	err = executor.QueryRowxContext(ctx, query, phone).Scan(
		&userModel.ID,
		&userModel.CreatedAt,
		&userModel.UpdatedAt,
		&userModel.Username,
		&userModel.PwdSecret,
		&userModel.Email,
		&userModel.Phone,
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
	r.logger.InfoKV("Updating username",
		"userId", user.ID,
		"newUsername", user.Username)

	executor, err := r.getExecutor()
	if err != nil {
		r.logger.ErrorKV("Failed to get executor", "error", err)
		return err
	}

	if user == nil || user.ID == 0 || user.Username == "" {
		r.logger.ErrorKV("Invalid user input for username update",
			"userId", user.ID,
			"username", user.Username)
		return domainErrors.ErrInvalidUserInput
	}

	// 检查用户名是否已被使用
	existingUser, err := r.FindByUsername(ctx, user.Username)
	if err != nil && !errors.Is(err, domainErrors.ErrUserNotFound) {
		r.logger.ErrorKV("Failed to check username availability",
			"username", user.Username,
			"error", err)
		return err
	}
	if existingUser != nil && existingUser.ID != user.ID {
		r.logger.WarnKV("Username already taken",
			"username", user.Username,
			"existingUserId", existingUser.ID)
		return domainErrors.ErrUsernameTaken
	}

	now := time.Now().Unix()
	query := `
		UPDATE users 
		SET updated_at = $1, username = $2 
		WHERE id = $3
	`

	result, err := executor.ExecContext(ctx, query, now, user.Username, user.ID)
	if err != nil {
		r.logger.ErrorKV("Failed to update username",
			"userId", user.ID,
			"username", user.Username,
			"error", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.logger.ErrorKV("Failed to get affected rows", "error", err)
		return err
	}

	if rowsAffected == 0 {
		r.logger.WarnKV("No rows affected during username update", "userId", user.ID)
		return domainErrors.ErrUserNotFound
	}

	user.UpdatedAt = now
	r.logger.InfoKV("Username updated successfully",
		"userId", user.ID,
		"newUsername", user.Username)

	return nil
}

// UpdatePwdSecret 更新密码
func (r *UserRepositoryImpl) UpdatePwdSecret(ctx context.Context, user *entity.User) error {
	executor, err := r.getExecutor()
	if err != nil {
		return err
	}

	if user == nil || user.ID == 0 || user.PwdSecret == "" {
		return domainErrors.ErrInvalidUserInput
	}

	now := time.Now().Unix()
	query := `
		UPDATE users 
		SET updated_at = $1, pwd_secret = $2 
		WHERE id = $3
	`

	result, err := executor.ExecContext(ctx, query, now, user.PwdSecret, user.ID)
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
	executor, err := r.getExecutor()
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

	result, err := executor.ExecContext(ctx, query, now, user.Email, user.ID)
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
	executor, err := r.getExecutor()
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

	result, err := executor.ExecContext(ctx, query, now, user.Phone, user.ID)
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
			DeletedAt: user.DeletedAt,
		},
		Username:  user.Username,
		PwdSecret: user.PwdSecret,
		Email:     user.Email,
		Phone:     user.Phone,
	}
}

func (r *UserRepositoryImpl) modelToEntity(userModel *model.UserModel) *entity.User {
	return &entity.User{
		ID:        userModel.ID,
		CreatedAt: userModel.CreatedAt,
		UpdatedAt: userModel.UpdatedAt,
		Username:  userModel.Username,
		PwdSecret: userModel.PwdSecret,
		Email:     userModel.Email,
		Phone:     userModel.Phone,
		DeletedAt: userModel.DeletedAt,
	}
}

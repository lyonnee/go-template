package repository_impl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/lyonnee/go-template/internal/domain/entity"
	domainErrors "github.com/lyonnee/go-template/internal/domain/errors"
	"github.com/lyonnee/go-template/internal/domain/repository"
	"github.com/lyonnee/go-template/internal/infrastructure/database"
	"github.com/lyonnee/go-template/internal/infrastructure/repository_impl/model"
	"github.com/lyonnee/go-template/pkg/di"
	"github.com/lyonnee/go-template/pkg/log"
	"go.uber.org/zap"
)

// 保证对接口实现
var _ repository.UserRepository = (*UserRepositoryImpl)(nil)

// UserRepositoryImpl 用户存储库实现
type UserRepositoryImpl struct {
	logger *log.Logger
}

func init() {
	err := di.AddSingletonImpl[repository.UserRepository, *UserRepositoryImpl](NewUserRepository)
	if err != nil {
		panic(err)
	}
}

// NewUserRepository 创建一个新的用户存储库实例
func NewUserRepository() (*UserRepositoryImpl, error) {
	repo := &UserRepositoryImpl{
		logger: di.Get[*log.Logger](),
	}

	return repo, nil
}

// FindById 根据ID查找用户
func (r *UserRepositoryImpl) FindById(ctx context.Context, userId uint64) (*entity.User, error) {
	r.logger.Debug("Finding user by ID", zap.Uint64("userId", userId))

	dbExecutor, err := database.GetDBExecutor(ctx)
	if err != nil {
		r.logger.Error("Failed to get DBExecutor", zap.Error(err))
		return nil, err
	}

	query := `
		SELECT id, created_at, updated_at, username, pwd_secret, email, phone, deleted_at 
		FROM users 
		WHERE id = $1 AND deleted_at = 0
	`

	var userModel model.UserModel
	err = dbExecutor.QueryRowxContext(ctx, query, userId).Scan(
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
			r.logger.Debug("User not found", zap.Uint64("userId", userId))
			return nil, domainErrors.ErrUserNotFound
		}
		r.logger.Error("Failed to find user by ID", zap.Uint64("userId", userId), zap.Error(err))
		return nil, err
	}

	r.logger.Debug("User found successfully", zap.Uint64("userId", userId), zap.String("username", userModel.Username))
	return r.modelToEntity(&userModel), nil
}

// Create 创建新用户
func (r *UserRepositoryImpl) Create(ctx context.Context, user *entity.User) error {
	r.logger.Debug("Creating new user", zap.String("username", user.Username), zap.String("email", user.Email))

	if user == nil {
		r.logger.Error("Invalid user input: user is nil")
		return domainErrors.ErrInvalidUserInput
	}

	dbExecutor, err := database.GetDBExecutor(ctx)
	if err != nil {
		r.logger.Error("Failed to get DBExecutor", zap.Error(err))
		return err
	}

	now := time.Now().Unix()
	query := `
		INSERT INTO users (created_at, updated_at, username, pwd_secret, email, phone) 
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	var id uint64
	err = dbExecutor.QueryRowxContext(ctx, query,
		now,
		now,
		user.Username,
		user.PwdSecret,
		user.Email,
		user.Phone,
	).Scan(&id)

	if err != nil {
		r.logger.Error("Failed to create user",
			zap.String("username", user.Username),
			zap.String("email", user.Email),
			zap.Error(err))
		return err
	}

	user.ID = id
	user.CreatedAt = now
	user.UpdatedAt = now
	user.DeletedAt = 0

	r.logger.Info("User created successfully",
		zap.Uint64("userId", id),
		zap.String("username", user.Username))

	return nil
}

// Update 更新用户信息
func (r *UserRepositoryImpl) Update(ctx context.Context, user *entity.User) error {
	dbExecutor, err := database.GetDBExecutor(ctx)
	if err != nil {
		r.logger.Error("Failed to get DBExecutor", zap.Error(err))
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

	_, err = dbExecutor.ExecContext(ctx, query,
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
func (r *UserRepositoryImpl) Delete(ctx context.Context, userId uint64) error {
	dbExecutor, err := database.GetDBExecutor(ctx)
	if err != nil {
		r.logger.Error("Failed to get DBExecutor", zap.Error(err))
		return err
	}

	now := time.Now().Unix()
	query := `
		UPDATE users 
		SET deleted_at = $1 
		WHERE id = $2
	`

	result, err := dbExecutor.ExecContext(ctx, query, now, userId)
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
	dbExecutor, err := database.GetDBExecutor(ctx)
	if err != nil {
		r.logger.Error("Failed to get DBExecutor", zap.Error(err))
		return nil, err
	}

	query := `
		SELECT id, created_at, updated_at, username, pwd_secret, email, phone, deleted_at 
		FROM users 
		WHERE username = $1 AND deleted_at = 0
	`

	var userModel model.UserModel

	err = dbExecutor.QueryRowxContext(ctx, query, username).StructScan(&userModel)
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
	dbExecutor, err := database.GetDBExecutor(ctx)
	if err != nil {
		r.logger.Error("Failed to get DBExecutor", zap.Error(err))
		return nil, err
	}

	query := `
		SELECT id, created_at, updated_at, username, pwd_secret, email, phone, deleted_at 
		FROM users 
		WHERE email = $1 AND deleted_at = 0
	`

	var userModel model.UserModel
	err = dbExecutor.QueryRowxContext(ctx, query, email).Scan(
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
	dbExecutor, err := database.GetDBExecutor(ctx)
	if err != nil {
		r.logger.Error("Failed to get DBExecutor", zap.Error(err))
		return nil, err
	}

	query := `
		SELECT id, created_at, updated_at, username, pwd_secret, email, phone, deleted_at 
		FROM users 
		WHERE phone = $1 AND deleted_at = 0
	`

	var userModel model.UserModel
	if err = dbExecutor.QueryRowxContext(ctx, query, phone).Scan(
		&userModel.ID,
		&userModel.CreatedAt,
		&userModel.UpdatedAt,
		&userModel.Username,
		&userModel.PwdSecret,
		&userModel.Email,
		&userModel.Phone,
		&userModel.DeletedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainErrors.ErrUserNotFound
		}
		return nil, err
	}

	return r.modelToEntity(&userModel), nil
}

// UpdateUsername 更新用户名
func (r *UserRepositoryImpl) UpdateUsername(ctx context.Context, user *entity.User) error {
	r.logger.Debug("Updating username",
		zap.Uint64("userId", user.ID),
		zap.String("newUsername", user.Username))

	dbExecutor, err := database.GetDBExecutor(ctx)
	if err != nil {
		r.logger.Error("Failed to get DBExecutor", zap.Error(err))
		return err
	}

	if user == nil || user.ID == 0 || user.Username == "" {
		r.logger.Error("Invalid user input for username update",
			zap.Uint64("userId", user.ID),
			zap.String("username", user.Username))
		return domainErrors.ErrInvalidUserInput
	}

	// 检查用户名是否已被使用
	existingUser, err := r.FindByUsername(ctx, user.Username)
	if err != nil && !errors.Is(err, domainErrors.ErrUserNotFound) {
		r.logger.Error("Failed to check username availability",
			zap.String("username", user.Username),
			zap.Error(err))
		return err
	}
	if existingUser != nil && existingUser.ID != user.ID {
		r.logger.Warn("Username already taken",
			zap.String("username", user.Username),
			zap.Uint64("existingUserId", existingUser.ID))
		return domainErrors.ErrUsernameTaken
	}

	now := time.Now().Unix()
	query := `
		UPDATE users 
		SET updated_at = $1, username = $2 
		WHERE id = $3
	`

	result, err := dbExecutor.ExecContext(ctx, query, now, user.Username, user.ID)
	if err != nil {
		r.logger.Error("Failed to update username",
			zap.Uint64("userId", user.ID),
			zap.String("username", user.Username),
			zap.Error(err))
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.logger.Error("Failed to get affected rows", zap.Error(err))
		return err
	}

	if rowsAffected == 0 {
		r.logger.Warn("No rows affected during username update",
			zap.Uint64("userId", user.ID))
		return domainErrors.ErrUserNotFound
	}

	user.UpdatedAt = now
	r.logger.Info("Username updated successfully",
		zap.Uint64("userId", user.ID),
		zap.String("newUsername", user.Username))

	return nil
}

// UpdatePwdSecret 更新密码
func (r *UserRepositoryImpl) UpdatePwdSecret(ctx context.Context, user *entity.User) error {
	dbExecutor, err := database.GetDBExecutor(ctx)
	if err != nil {
		r.logger.Error("Failed to get DBExecutor", zap.Error(err))
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

	result, err := dbExecutor.ExecContext(ctx, query, now, user.PwdSecret, user.ID)
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
	dbExecutor, err := database.GetDBExecutor(ctx)
	if err != nil {
		r.logger.Error("Failed to get DBExecutor", zap.Error(err))
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

	result, err := dbExecutor.ExecContext(ctx, query, now, user.Email, user.ID)
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
	dbExecutor, err := database.GetDBExecutor(ctx)
	if err != nil {
		r.logger.Error("Failed to get DBExecutor", zap.Error(err))
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

	result, err := dbExecutor.ExecContext(ctx, query, now, user.Phone, user.ID)
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
func (r *UserRepositoryImpl) CheckUserFieldsExist(ctx context.Context, username, email, phone string) (bool, error) {
	dbExecutor, err := database.GetDBExecutor(ctx)
	if err != nil {
		r.logger.Error("Failed to get DBExecutor", zap.Error(err))
		return false, err
	}

	// 构建动态查询条件
	var conditions []string
	var args []interface{}
	argIndex := 1

	if username != "" {
		conditions = append(conditions, "username = $"+fmt.Sprintf("%d", argIndex))
		args = append(args, username)
		argIndex++
	}

	if email != "" {
		conditions = append(conditions, "email = $"+fmt.Sprintf("%d", argIndex))
		args = append(args, email)
		argIndex++
	}

	if phone != "" {
		conditions = append(conditions, "phone = $"+fmt.Sprintf("%d", argIndex))
		args = append(args, phone)
		argIndex++
	}

	// 如果没有需要检查的字段，直接返回
	if len(conditions) == 0 {
		return false, nil
	}

	// 一次查询检查所有字段
	query := fmt.Sprintf(`
		SELECT username, email, phone 
		FROM users 
		WHERE (%s) AND deleted_at = 0
		LIMIT 1
	`, strings.Join(conditions, " OR "))

	err = dbExecutor.QueryRowxContext(ctx, query, args...).Scan(&username, &email, &phone)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// 没有找到重复记录，表示字段可用
			return false, nil
		}
		r.logger.Error("Failed to check user fields existence", zap.Error(err))
		return false, err
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

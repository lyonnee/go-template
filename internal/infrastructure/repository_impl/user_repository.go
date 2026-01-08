package repository_impl

import (
	"context"
	"errors"
	"time"

	"github.com/lyonnee/go-template/internal/domain/entity"
	domainErrors "github.com/lyonnee/go-template/internal/domain/errors"
	"github.com/lyonnee/go-template/internal/domain/repository"
	"github.com/lyonnee/go-template/internal/infrastructure/repository_impl/model"
	"github.com/lyonnee/go-template/pkg/di"
	"gorm.io/gorm"
)

// 保证对接口实现
var _ repository.UserRepository = (*UserRepositoryImpl)(nil)
var _ di.Repository = (*UserRepositoryImpl)(nil)

// UserRepositoryImpl 用户存储库实现
type UserRepositoryImpl struct {
	BaseRepository
}

func init() {
	err := di.AddTransientImpl[repository.UserRepository, *UserRepositoryImpl](NewUserRepository)
	if err != nil {
		panic(err)
	}
}

// NewUserRepository 创建一个新的用户存储库实例
func NewUserRepository() (*UserRepositoryImpl, error) {
	repo := &UserRepositoryImpl{}

	return repo, nil
}

// FindById 根据ID查找用户
func (r *UserRepositoryImpl) FindById(ctx context.Context, userId uint64) (*entity.User, error) {
	var userModel model.UserModel
	db, err := r.DB()
	if err != nil {
		return nil, err
	}

	if err := db.WithContext(ctx).Where("id = ? AND deleted_at = 0", userId).First(&userModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrUserNotFound
		}
		return nil, err
	}

	return r.modelToEntity(&userModel), nil
}

// Create 创建新用户
func (r *UserRepositoryImpl) Create(ctx context.Context, user *entity.User) error {
	if user == nil {
		return domainErrors.ErrInvalidUserInput
	}

	db, err := r.DB()
	if err != nil {
		return err
	}

	userModel := r.entityToModel(user)
	if err := db.WithContext(ctx).Create(userModel).Error; err != nil {
		return err
	}

	user.ID = userModel.ID

	return nil
}

// Update 更新用户信息
func (r *UserRepositoryImpl) Update(ctx context.Context, user *entity.User) error {
	if user == nil || user.ID == 0 {
		return domainErrors.ErrInvalidUserInput
	}

	db, err := r.DB()
	if err != nil {
		return err
	}

	userModel := r.entityToModel(user)
	result := db.WithContext(ctx).Model(&model.UserModel{}).Where("id = ? AND deleted_at = 0", user.ID).Updates(userModel)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domainErrors.ErrUserNotFound
	}

	return nil
}

// Delete 删除用户（软删除）
func (r *UserRepositoryImpl) Delete(ctx context.Context, userId uint64) error {
	db, err := r.DB()
	if err != nil {
		return err
	}

	now := time.Now().Unix()
	result := db.WithContext(ctx).Model(&model.UserModel{}).Where("id = ?", userId).Update("deleted_at", now)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domainErrors.ErrUserNotFound
	}

	return nil
}

// FindByUsername 根据用户名查找用户
func (r *UserRepositoryImpl) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	db, err := r.DB()
	if err != nil {
		return nil, err
	}

	var userModel model.UserModel
	if err := db.WithContext(ctx).Where("username = ? AND deleted_at = 0", username).First(&userModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrUserNotFound
		}
		return nil, err
	}

	return r.modelToEntity(&userModel), nil
}

// FindByEmail 根据邮箱查找用户
func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	db, err := r.DB()
	if err != nil {
		return nil, err
	}
	var userModel model.UserModel
	if err := db.WithContext(ctx).Where("email = ? AND deleted_at = 0", email).First(&userModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrUserNotFound
		}
		return nil, err
	}

	return r.modelToEntity(&userModel), nil
}

// FindByPhone 根据手机号查找用户
func (r *UserRepositoryImpl) FindByPhone(ctx context.Context, phone string) (*entity.User, error) {
	db, err := r.DB()
	if err != nil {
		return nil, err
	}

	var userModel model.UserModel
	if err := db.WithContext(ctx).Where("phone = ? AND deleted_at = 0", phone).First(&userModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrUserNotFound
		}
		return nil, err
	}

	return r.modelToEntity(&userModel), nil
}

// UpdateUsername 更新用户名
func (r *UserRepositoryImpl) UpdateUsername(ctx context.Context, user *entity.User) error {
	db, err := r.DB()
	if err != nil {
		return err
	}

	if user == nil || user.ID == 0 || user.Username == "" {
		return domainErrors.ErrInvalidUserInput
	}

	result := db.WithContext(ctx).Model(&model.UserModel{}).Where("id = ?", user.ID).Updates(map[string]interface{}{
		"updated_at": user.UpdatedAt,
		"username":   user.Username,
	})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domainErrors.ErrUserNotFound
	}

	return nil
}

// UpdatePwdSecret 更新密码
func (r *UserRepositoryImpl) UpdatePwdSecret(ctx context.Context, user *entity.User) error {
	db, err := r.DB()
	if err != nil {
		return err
	}

	if user == nil || user.ID == 0 || user.PwdSecret == "" {
		return domainErrors.ErrInvalidUserInput
	}

	result := db.WithContext(ctx).Model(&model.UserModel{}).Where("id = ?", user.ID).Updates(map[string]interface{}{
		"updated_at": user.UpdatedAt,
		"pwd_secret": user.PwdSecret,
	})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domainErrors.ErrUserNotFound
	}

	return nil
}

// UpdateEmail 更新邮箱
func (r *UserRepositoryImpl) UpdateEmail(ctx context.Context, user *entity.User) error {
	db, err := r.DB()
	if err != nil {
		return err
	}

	if user == nil || user.ID == 0 || user.Email == "" {
		return domainErrors.ErrInvalidUserInput
	}

	result := db.WithContext(ctx).Model(&model.UserModel{}).Where("id = ?", user.ID).Updates(map[string]interface{}{
		"updated_at": user.UpdatedAt,
		"email":      user.Email,
	})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domainErrors.ErrUserNotFound
	}

	return nil
}

// UpdatePhone 更新手机号
func (r *UserRepositoryImpl) UpdatePhone(ctx context.Context, user *entity.User) error {
	db, err := r.DB()
	if err != nil {
		return err
	}

	if user == nil || user.ID == 0 || user.Phone == "" {
		return domainErrors.ErrInvalidUserInput
	}

	result := db.WithContext(ctx).Model(&model.UserModel{}).Where("id = ?", user.ID).Updates(map[string]interface{}{
		"updated_at": user.UpdatedAt,
		"phone":      user.Phone,
	})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domainErrors.ErrUserNotFound
	}

	return nil
}

// 检查用户字段是否存在
func (r *UserRepositoryImpl) CheckUserFieldsExist(ctx context.Context, username, email, phone string) (bool, error) {
	db, err := r.DB()
	if err != nil {
		return false, err
	}

	db = db.WithContext(ctx).Model(&model.UserModel{}).Where("deleted_at = 0")

	if username != "" {
		db = db.Or("username = ?", username)
	}
	if email != "" {
		db = db.Or("email = ?", email)
	}
	if phone != "" {
		db = db.Or("phone = ?", phone)
	}

	var count int64
	if err := db.Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
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

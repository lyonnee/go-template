package repository

import (
	"context"

	"github.com/lyonnee/go-template/internal/domain/entity"
)

// UserRepository 用户存储库接口
// 存储库方法可能返回的错误：
// - errors.ErrUserNotFound：用户不存在
// - errors.ErrUsernameTaken：用户名已被占用
// - errors.ErrEmailTaken：邮箱已被占用
// - errors.ErrPhoneTaken：手机号已被占用
// - errors.ErrInvalidUserInput：无效的用户输入
type UserRepository interface {
	// 基本的CRUD操作
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, userId int64) error

	// 查询操作
	FindById(ctx context.Context, userId int64) (*entity.User, error)
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindByPhone(ctx context.Context, phone string) (*entity.User, error)

	// 特定更新操作
	UpdateUsername(ctx context.Context, user *entity.User) error
	UpdatePwdSecret(ctx context.Context, user *entity.User) error
	UpdateEmail(ctx context.Context, user *entity.User) error
	UpdatePhone(ctx context.Context, user *entity.User) error
}

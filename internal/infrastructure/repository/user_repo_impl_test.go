package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/lyonnee/go-template/internal/domain/entity"
	"github.com/lyonnee/go-template/internal/infrastructure/repository"
)

func TestUserRepositoryImpl_Create(t *testing.T) {
	// 设置模拟DB
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// 设置测试数据
	username := "testuser"
	password := "password123"
	email := "test@example.com"
	phone := "13800138000"
	userID := int64(1)

	// 设置SQL期望
	mock.ExpectQuery("INSERT INTO users").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), username, password, email, phone, false, 0).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(userID))

	// 创建存储库并设置执行器
	userRepo := repository.NewUserRepository().WithExecuter(sqlxDB)

	// 创建测试用户实体
	user := &entity.User{
		Username: username,
		Password: password,
		Email:    email,
		Phone:    phone,
	}

	// 执行测试
	ctx := context.Background()
	createdUser, err := userRepo.Create(ctx, user)

	// 断言结果
	require.NoError(t, err)
	assert.NotNil(t, createdUser)
	assert.Equal(t, userID, createdUser.ID)
	assert.Equal(t, username, createdUser.Username)
	assert.Equal(t, password, createdUser.Password)
	assert.Equal(t, email, createdUser.Email)
	assert.Equal(t, phone, createdUser.Phone)
	assert.False(t, createdUser.IsDeleted)

	// 验证所有期望都已满足
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepositoryImpl_FindById(t *testing.T) {
	// 设置模拟DB
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// 设置测试数据
	userID := int64(1)
	now := time.Now().Unix()
	username := "testuser"
	password := "password123"
	email := "test@example.com"
	phone := "13800138000"

	// 设置SQL期望
	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "username", "password", "email", "phone", "is_deleted", "deleted_at"}).
		AddRow(userID, now, now, username, password, email, phone, false, 0)

	mock.ExpectQuery("SELECT (.+) FROM users").
		WithArgs(userID).
		WillReturnRows(rows)

	// 创建存储库并设置执行器
	userRepo := repository.NewUserRepository().WithExecuter(sqlxDB)

	// 执行测试
	ctx := context.Background()
	foundUser, err := userRepo.FindById(ctx, userID)

	// 断言结果
	require.NoError(t, err)
	assert.NotNil(t, foundUser)
	assert.Equal(t, userID, foundUser.ID)
	assert.Equal(t, username, foundUser.Username)
	assert.Equal(t, password, foundUser.Password)
	assert.Equal(t, email, foundUser.Email)
	assert.Equal(t, phone, foundUser.Phone)
	assert.False(t, foundUser.IsDeleted)

	// 验证所有期望都已满足
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepositoryImpl_NoExecuter(t *testing.T) {
	// 创建存储库但不设置执行器
	userRepo := repository.NewUserRepository()

	// 执行测试
	ctx := context.Background()
	_, err := userRepo.FindById(ctx, 1)

	// 断言结果 - 应该返回错误
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "executer not set")
}

func TestUserRepositoryImpl_WithTransaction(t *testing.T) {
	// 设置模拟DB
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// 设置事务期望
	mock.ExpectBegin()

	// 开始事务
	tx, err := sqlxDB.Beginx()
	require.NoError(t, err)

	// 设置创建用户的期望
	mock.ExpectQuery("INSERT INTO users").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "txuser", "password", "tx@example.com", "13900001111", false, 0).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))

	// 设置查找用户的期望
	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "username", "password", "email", "phone", "is_deleted", "deleted_at"}).
		AddRow(2, time.Now().Unix(), time.Now().Unix(), "txuser", "password", "tx@example.com", "13900001111", false, 0)

	mock.ExpectQuery("SELECT (.+) FROM users").
		WithArgs(int64(2)).
		WillReturnRows(rows)

	// 设置提交期望
	mock.ExpectCommit()

	// 创建存储库并设置事务执行器
	userRepo := repository.NewUserRepository().WithExecuter(tx)

	// 执行测试 - 创建用户
	ctx := context.Background()
	user := &entity.User{
		Username: "txuser",
		Password: "password",
		Email:    "tx@example.com",
		Phone:    "13900001111",
	}

	createdUser, err := userRepo.Create(ctx, user)
	require.NoError(t, err)
	assert.Equal(t, int64(2), createdUser.ID)

	// 查找用户
	foundUser, err := userRepo.FindById(ctx, createdUser.ID)
	require.NoError(t, err)
	assert.Equal(t, "txuser", foundUser.Username)

	// 提交事务
	err = tx.Commit()
	require.NoError(t, err)

	// 验证所有期望都已满足
	assert.NoError(t, mock.ExpectationsWereMet())
}

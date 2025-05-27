-- 创建用户表
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,

    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    deleted_at BIGINT DEFAULT 0,

    username VARCHAR(255) NOT NULL UNIQUE,
    pwd_secret VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE,
    phone VARCHAR(50) UNIQUE
);

-- 创建索引
CREATE INDEX IF NOT EXISTS idx_users_username ON users (username);
CREATE INDEX IF NOT EXISTS idx_users_email ON users (email);
CREATE INDEX IF NOT EXISTS idx_users_phone ON users (phone);
CREATE INDEX IF NOT EXISTS idx_users_is_deleted ON users (is_deleted);
CREATE INDEX IF NOT EXISTS idx_users_is_deleted_created_at ON users (is_deleted, created_at);

COMMENT ON TABLE users IS '用户表';
COMMENT ON COLUMN users.id IS '用户ID，自增主键';
COMMENT ON COLUMN users.created_at IS '用户创建时间戳（UTC时区）';
COMMENT ON COLUMN users.updated_at IS '用户信息更新时间戳（UTC时区）';
COMMENT ON COLUMN users.is_deleted IS '软删除标记（TRUE=已删除）';
COMMENT ON COLUMN users.deleted_at IS '删除时间戳（UTC时区）';
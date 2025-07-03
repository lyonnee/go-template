-- 创建用户表
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,

    deleted_at BIGINT NOT NULL DEFAULT 0,

    username VARCHAR(255) NOT NULL UNIQUE,
    pwd_secret VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE DEFAULT '',
    phone VARCHAR(50) NOT NULL UNIQUE DEFAULT ''
);

-- 创建索引
CREATE INDEX IF NOT EXISTS idx_users_username ON users (username);
CREATE INDEX IF NOT EXISTS idx_users_email ON users (email);
CREATE INDEX IF NOT EXISTS idx_users_phone ON users (phone);

COMMENT ON TABLE users IS '用户表';
COMMENT ON COLUMN users.id IS '用户ID，自增主键';
COMMENT ON COLUMN users.created_at IS '用户创建时间戳（UTC时区）';
COMMENT ON COLUMN users.updated_at IS '用户信息更新时间戳（UTC时区）';
COMMENT ON COLUMN users.deleted_at IS '删除时间戳（UTC时区）';
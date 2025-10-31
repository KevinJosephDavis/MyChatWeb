USE chatroom;

-- 用户表
CREATE TABLE users (
    id BIGINT AUTO_INCREMENT PRIMARY KEY, --内部ID
    account_id BIGINT UNIQUE NOT NULL, --对外展示ID
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(100),
    avatar VARCHAR(255) DEFAULT '',
    status TINYINT DEFAULT 1 COMMENT '1-正常 2-禁用',
    last_login_at TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 好友关系表
CREATE TABLE friends (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    friend_id BIGINT NOT NULL,
    remark VARCHAR(50) COMMENT '好友备注',
    status TINYINT DEFAULT 1 COMMENT '1-好友 2-拉黑 3-已删除',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY uk_user_friend (user_id, friend_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (friend_id) REFERENCES users(id) ON DELETE CASCADE
);

-- 群组表
CREATE TABLE `groups` (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    avatar VARCHAR(255) DEFAULT '',
    owner_id BIGINT NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (owner_id) REFERENCES users(id)
);

-- 群成员表
CREATE TABLE group_members (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    group_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    role TINYINT DEFAULT 1 COMMENT '1-成员 2-管理员 3-群主',
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY uk_group_user (group_id, user_id),
    FOREIGN KEY (group_id) REFERENCES `groups`(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- 消息表（单聊和群聊消息）
CREATE TABLE messages (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    sender_id BIGINT NOT NULL,
    receiver_id BIGINT COMMENT '单聊接收者ID',
    group_id BIGINT COMMENT '群聊ID',
    message_type TINYINT NOT NULL COMMENT '1-文本 2-图片 3-文件 4-语音 5-视频',
    content TEXT NOT NULL,
    file_url VARCHAR(500) COMMENT '文件URL',
    file_size BIGINT COMMENT '文件大小',
    file_name VARCHAR(255) COMMENT '原始文件名',
    is_read TINYINT DEFAULT 0 COMMENT '0-未读 1-已读',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    KEY idx_sender (sender_id),
    KEY idx_receiver (receiver_id),
    KEY idx_group (group_id),
    KEY idx_created (created_at),
    FOREIGN KEY (sender_id) REFERENCES users(id),
    FOREIGN KEY (receiver_id) REFERENCES users(id),
    FOREIGN KEY (group_id) REFERENCES `groups`(id)
);
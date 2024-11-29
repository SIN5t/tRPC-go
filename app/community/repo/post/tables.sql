Create TABLE if not exists `post`(
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
    topic_id BIGINT NOT NULL,                   -- 关联的 Topic ID
    create_at DATETIME NOT NULL,                -- 创建时间，假设使用 DATETIME 类型
    content TEXT NOT NULL,                      -- 帖子的内容，假设使用 TEXT 类型
    FOREIGN KEY (topic_id) REFERENCES topics(id), -- 假设存在一个 topics 表与之关联
    PRIMARY KEY (`id`),
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户账户表';


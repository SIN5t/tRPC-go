create database if not exists `yicwu`
       use `yicwu`
create table if not exists `topic`(
    `id` int(11) NOT NULL  AUTO_INCREMENT COMMENT '主键 ID',
    PRIMARY KEY (`id`),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
);

INSERT INTO topics (title, description) VALUES ('Second Topic', 'This is the description of the second topic');

INSERT INTO topics (title, description)
VALUES
    ('Title 1', 'Description 1'),
    ('Title 2', 'Description 2'),
    ('Title 3', 'Description 3'),
    ('Title 4', 'Description 4'),
    ('Title 5', 'Description 5');
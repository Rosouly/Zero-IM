CREATE TABLE `user`
(
    `id`            char(32)     NOT NULL COMMENT '主键id',
    `username`      char(31)     NOT NULL UNIQUE COMMENT '用户名',
    `password`      char(64)     NOT NULL COMMENT '密码',
    `nickname`      varchar(31)  NOT NULL COMMENT '昵称',
    `sign`          varchar(63)  NOT NULL DEFAULT '' COMMENT '个性签名',
    `avatar`        varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
    `province`      varchar(63)  NOT NULL DEFAULT '' COMMENT '省',
    `city`          varchar(63)  NOT NULL DEFAULT '' COMMENT '市',
    `district`      varchar(63)  NOT NULL DEFAULT '' COMMENT '区',
    `birthday`      char(10)     NOT NULL DEFAULT '' COMMENT '生日',
    `register_time` char(19)     NOT NULL COMMENT '注册时间',
    `is_male`       boolean      NOT NULL DEFAULT 1 COMMENT '是否是男性',
    PRIMARY KEY (`id`),
    INDEX (`username`),
    INDEX (`province`),
    INDEX (`city`),
    INDEX (`district`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';
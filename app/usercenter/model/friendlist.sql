CREATE TABLE IF NOT EXISTS `friendlist` (
    `user_id` char(32) NOT NULL COMMENT '主键 用户id',
    `self_id` char(32) NOT NULL COMMENT '自己id,分表键',
    `created_at` int(10) NOT NULL COMMENT '加好友时间 秒级时间戳',
    `remark` varchar(255) DEFAULT NULL COMMENT '备注',
    PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
USE tin;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for T_USER
-- ----------------------------

CREATE TABLE IF NOT EXISTS `t_user` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `nickname` varchar(100) DEFAULT NULL COMMENT '昵称',
  `real_name` varchar(100) DEFAULT NULL COMMENT '实名',
  `mail` varchar(100) DEFAULT NULL COMMENT '邮箱',
  `phone` varchar(20) DEFAULT NULL COMMENT '',
  `gender` tinyint(1) DEFAULT '1' COMMENT '性别：0.女 1.男',
  `dob` datetime DEFAULT NULL COMMENT '生日',
  `avatar` varchar(100) DEFAULT NULL COMMENT '头像uri',
  `create_user` bigint(20) DEFAULT NULL COMMENT '创建人',
  `update_user` bigint(20) DEFAULT NULL COMMENT '修改人',
  `create_time` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) DEFAULT NULL COMMENT '修改时间',
  `delete_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

USE tin;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for T_AUTH
-- ----------------------------

CREATE TABLE IF NOT EXISTS `t_auth` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `provider` varchar(32) DEFAULT 'password' COMMENT '认证方式',
  `uid` varchar(64) NOT NULL DEFAULT '""' COMMENT 'auth uid',
  `password` varchar(64) DEFAULT NULL COMMENT 'auth password',
  `create_user` bigint(20) DEFAULT NULL COMMENT '创建人',
  `update_user` bigint(20) DEFAULT NULL COMMENT '修改人',
  `create_time` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) DEFAULT NULL COMMENT '修改时间',
  `delete_time` datetime(3) DEFAULT NULL,
  `sign_logs` mediumtext COLLATE utf8mb4_general_ci,
  PRIMARY KEY (`id`),
  KEY `USER_ID` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
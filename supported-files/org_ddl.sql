USE tin;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
-- ----------------------------
-- Table structure for T_ORG
-- ----------------------------

CREATE TABLE IF NOT EXISTS `t_org` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `uid` varchar(50) NOT NULL COMMENT 'org unique code',
  `join_code` varchar(100) DEFAULT NULL COMMENT 'org join code',
  `nickname` varchar(100) DEFAULT NULL COMMENT '昵称',
  `real_name` varchar(100) DEFAULT NULL COMMENT '实名',
  `verified` tinyint(1) DEFAULT '0' COMMENT '是否认证：0.否 1.是',
  `activated` tinyint(1) DEFAULT '0' COMMENT '是否激活：0.否 1.是',
  `create_user` bigint(20) DEFAULT NULL COMMENT '创建人',
  `update_user` bigint(20) DEFAULT NULL COMMENT '修改人',
  `create_time` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) DEFAULT NULL COMMENT '修改时间',
  `delete_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for T_ORG_USER
-- ----------------------------

CREATE TABLE IF NOT EXISTS `t_org_user` (
  `id` bigint(20) NOT NULL,
  `org_id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `role_code` varchar(500) NOT NULL,
  `create_user` bigint(20) DEFAULT NULL COMMENT '创建人',
  `update_user` bigint(20) DEFAULT NULL COMMENT '修改人',
  `create_time` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) DEFAULT NULL COMMENT '修改时间',
  `delete_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `USER_ID` (`user_id`),
  KEY `ORG_ID` (`org_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for T_ORG_APP
-- ----------------------------

CREATE TABLE IF NOT EXISTS `t_org_app` (
  `id` bigint(20) NOT NULL,
  `app_uid` varchar(50) DEFAULT NULL,
  `package_uid` varchar(50) DEFAULT NULL,
  `user_id` bigint(20) DEFAULT NULL COMMENT '订阅人',
  `org_id` bigint(20) DEFAULT NULL,
  `create_user` bigint(20) DEFAULT NULL COMMENT '创建人',
  `update_user` bigint(20) DEFAULT NULL COMMENT '修改人',
  `create_time` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) DEFAULT NULL COMMENT '修改时间',
  `delete_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `USER_ID` (`user_id`),
  KEY `ORG_ID`(`org_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
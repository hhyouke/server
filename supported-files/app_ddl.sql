USE tin;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for T_APP
-- ----------------------------

CREATE TABLE IF NOT EXISTS `t_app` (
  `id` bigint(20) NOT NULL,
  `uid` varchar(50) DEFAULT NULL,
  `name` varchar(100) DEFAULT NULL,
  `dev_team` varchar(100) DEFAULT NULL COMMENT 'develop team name',
  `release_status` tinyint(1) DEFAULT '0' COMMENT 'release status: 0.n 1.y',
  `release_version` varchar(20) DEFAULT NULL,
  `create_user` bigint(20) DEFAULT NULL COMMENT '创建人',
  `update_user` bigint(20) DEFAULT NULL COMMENT '修改人',
  `create_time` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for T_APP_PACKAGE
-- ----------------------------

CREATE TABLE IF NOT EXISTS `t_app_package` (
  `id` bigint(20) NOT NULL,
  `app_id` bigint(20) NOT NULL,
  `uid` varchar(50) DEFAULT NULL,
  `name` varchar(100) DEFAULT NULL,
  `foc` tinyint(1) DEFAULT '0' COMMENT 'free of charge: 0.n 1.y',
  `create_user` bigint(20) DEFAULT NULL COMMENT '创建人',
  `update_user` bigint(20) DEFAULT NULL COMMENT '修改人',
  `create_time` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `APP_ID` (`app_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for T_ROLE
-- ----------------------------

CREATE TABLE IF NOT EXISTS `t_role` (
  `id` bigint(20) NOT NULL,
  `code` varchar(50) DEFAULT NULL,
  `name` varchar(100) DEFAULT NULL,
  `create_user` bigint(20) DEFAULT NULL COMMENT '创建人',
  `update_user` bigint(20) DEFAULT NULL COMMENT '修改人',
  `create_time` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
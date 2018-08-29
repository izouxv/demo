CREATE TABLE `tbl_domain` (
  `did` bigint(20) NOT NULL,
  `domainName` varchar(128) NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01',
  `update_time` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01',
  PRIMARY KEY (`did`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tbl_domain
-- ----------------------------
INSERT INTO `tbl_domain` VALUES ('923456632830690001', '蓝涟域', '2017-11-29 19:59:29', '2017-11-29 19:59:29');
INSERT INTO `tbl_domain` VALUES ('923456632830690002', '元安域', '1970-01-01 08:00:01', '1970-01-01 08:00:01');
INSERT INTO `tbl_domain` VALUES ('923456632830690003', '鹏联域', '1970-01-01 08:00:01', '1970-01-01 08:00:01');

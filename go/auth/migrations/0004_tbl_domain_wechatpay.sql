CREATE TABLE `tbl_domain_wechatpay` (
  `did` bigint(20) NOT NULL,
  `app_id` varchar(255) NOT NULL,
  `mch_id` varchar(255) NOT NULL,
  `key` varchar(255) NOT NULL,
  `app_secret` varchar(255) NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01',
  `update_time` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01',
  PRIMARY KEY (`did`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tbl_domain_wechatpay
-- ----------------------------

CREATE TABLE `tbl_policy` (
  `policy_id` int(11) NOT NULL AUTO_INCREMENT,
  `policy_sid` int(11) NOT NULL COMMENT '计费策略所属的服务id',
  `policy_name` varchar(255) NOT NULL COMMENT '计费策略名',
  `policy_type` int(11) NOT NULL COMMENT '计费策略类型',
  `policy_cycle` int(11) NOT NULL COMMENT '计费策略计算周期',
  `policy_fee_type` int(11) NOT NULL,
  `policy_unit_price` decimal(20,2) NOT NULL,
  `policy_unit_type` int(11) NOT NULL,
  `policy_unit_count` int(11) NOT NULL,
  `create_time` timestamp NULL DEFAULT '1970-01-01 08:00:01',
  `update_time` timestamp NULL DEFAULT '1970-01-01 08:00:01',
  PRIMARY KEY (`policy_id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tbl_policy
-- ----------------------------
INSERT INTO `tbl_policy` VALUES ('1', '1', 'test1', '2', '2', '1', '2.10', '1', '100', '2018-03-18 16:10:44', '2018-03-18 17:11:04');
INSERT INTO `tbl_policy` VALUES ('2', '2', 'test2', '2', '2', '1', '2.10', '1', '100', '2018-03-18 16:10:44', '2018-03-18 17:11:04');

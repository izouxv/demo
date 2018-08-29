CREATE TABLE `tbl_action_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `action_username` varchar(255) NOT NULL,
  `action_time` int(11) NOT NULL,
  `action_type` int(11) NOT NULL,
  `action_name` varchar(255) NOT NULL,
  `action_object` varchar(255) NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1762 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tbl_action_log
-- ----------------------------

CREATE TABLE `tbl_module` (
  `module_id` int(11) NOT NULL AUTO_INCREMENT,
  `module_name` varchar(255) DEFAULT NULL,
  `module_did` bigint(20) DEFAULT NULL,
  `module_tid` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`module_id`)
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tbl_module
-- ----------------------------
INSERT INTO `tbl_module` VALUES ('1', '服务', '923456632830690002', null);
INSERT INTO `tbl_module` VALUES ('2', '用户', '923456632830690002', null);
INSERT INTO `tbl_module` VALUES ('3', '角色', '923456632830690002', null);
INSERT INTO `tbl_module` VALUES ('4', '日志', '923456632830690002', null);
INSERT INTO `tbl_module` VALUES ('5', '设置', '923456632830690002', null);
INSERT INTO `tbl_module` VALUES ('6', '密码', '923456632830690002', null);
INSERT INTO `tbl_module` VALUES ('7', '文件库', '923456632830690002', null);
INSERT INTO `tbl_module` VALUES ('8', '保留规则', '923456632830690002', null);
INSERT INTO `tbl_module` VALUES ('9', '财务总览', '923456632830690002', null);
INSERT INTO `tbl_module` VALUES ('10', '消费纵览', '923456632830690002', null);
INSERT INTO `tbl_module` VALUES ('11', '服务', null, '935840632830697474');
INSERT INTO `tbl_module` VALUES ('12', '用户', null, '935840632830697474');
INSERT INTO `tbl_module` VALUES ('13', '角色', null, '935840632830697474');
INSERT INTO `tbl_module` VALUES ('14', '日志', null, '935840632830697474');
INSERT INTO `tbl_module` VALUES ('15', '设置', null, '935840632830697474');
INSERT INTO `tbl_module` VALUES ('16', '密码', null, '935840632830697474');
INSERT INTO `tbl_module` VALUES ('17', '文件库', null, '935840632830697474');
INSERT INTO `tbl_module` VALUES ('18', '保留规则', null, '935840632830697474');
INSERT INTO `tbl_module` VALUES ('19', '财务总览', null, '935840632830697474');
INSERT INTO `tbl_module` VALUES ('20', '消费纵览', null, '935840632830697474');
INSERT INTO `tbl_module` VALUES ('21', '服务', '0', '961777349764845568');
INSERT INTO `tbl_module` VALUES ('22', '用户', '0', '961777349764845568');
INSERT INTO `tbl_module` VALUES ('23', '角色', '0', '961777349764845568');
INSERT INTO `tbl_module` VALUES ('24', '日志', '0', '961777349764845568');
INSERT INTO `tbl_module` VALUES ('25', '设置', '0', '961777349764845568');
INSERT INTO `tbl_module` VALUES ('26', '密码', '0', '961777349764845568');
INSERT INTO `tbl_module` VALUES ('27', '文件库', '0', '961777349764845568');
INSERT INTO `tbl_module` VALUES ('28', '保留规则', '0', '961777349764845568');
INSERT INTO `tbl_module` VALUES ('29', '财务总览', '0', '961777349764845568');
INSERT INTO `tbl_module` VALUES ('30', '消费总览', '0', '961777349764845568');
INSERT INTO `tbl_module` VALUES ('31', '租户管理', '923456632830690002', '0');
INSERT INTO `tbl_module` VALUES ('32', '设备管理服务', '0', '935840632830697474');
INSERT INTO `tbl_module` VALUES ('33', 'KMS服务', '0', '935840632830697474');
INSERT INTO `tbl_module` VALUES ('34', '设备管理服务', '0', '961777349764845568');
INSERT INTO `tbl_module` VALUES ('35', 'KMS服务', '0', '961777349764845568');

/*
Navicat MySQL Data Transfer

Source Server         : 1.6
Source Server Version : 50719
Source Host           : 192.168.1.6:3306
Source Database       : auth

Target Server Type    : MYSQL
Target Server Version : 50719
File Encoding         : 65001

Date: 2018-04-13 16:50:10
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tbl_role
-- ----------------------------
DROP TABLE IF EXISTS `tbl_role`;
CREATE TABLE `tbl_role` (
  `rid` int(11) NOT NULL AUTO_INCREMENT,
  `roleName` varchar(64) CHARACTER SET utf8 NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `tid` bigint(20) DEFAULT '0',
  `create_time` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01',
  `update_time` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01',
  `did` bigint(20) DEFAULT '0',
  PRIMARY KEY (`rid`)
) ENGINE=InnoDB AUTO_INCREMENT=131 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tbl_role
-- ----------------------------
INSERT INTO `tbl_role` VALUES ('1', '超级管理员', '超级管理员', '0', '2017-11-24 08:00:01', '2018-02-01 11:59:01', '923456632830690002');
INSERT INTO `tbl_role` VALUES ('2', '超级管理员', '超级管理员', '935840632830697474', '2017-11-24 08:00:01', '2018-02-05 10:18:20', '0');
INSERT INTO `tbl_role` VALUES ('3', '超级管理员', '超级管理员', '961777349764845568', '2017-11-24 08:00:01', '2018-02-01 11:59:01', '0');

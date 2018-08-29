/*
Navicat MySQL Data Transfer

Source Server         : 1.6
Source Server Version : 50719
Source Host           : 192.168.1.6:3306
Source Database       : auth

Target Server Type    : MYSQL
Target Server Version : 50719
File Encoding         : 65001

Date: 2018-04-13 16:50:45
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tbl_tenant
-- ----------------------------
DROP TABLE IF EXISTS `tbl_tenant`;
CREATE TABLE `tbl_tenant` (
  `tid` bigint(20) NOT NULL AUTO_INCREMENT,
  `tenantName` varchar(128) CHARACTER SET utf8 NOT NULL,
  `pid` bigint(20) DEFAULT '1',
  `create_time` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01',
  `update_time` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01',
  `did` bigint(20) NOT NULL,
  `tenantURL` varchar(255) CHARACTER SET utf8 NOT NULL,
  `tenantState` tinyint(4) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `contacts` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  `phone` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  `icon` varchar(255) CHARACTER SET utf8 DEFAULT '',
  `logo` varchar(255) CHARACTER SET utf8 DEFAULT '',
  PRIMARY KEY (`tid`),
  KEY `tid` (`tid`)
) ENGINE=InnoDB AUTO_INCREMENT=971693127364513793 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tbl_tenant
-- ----------------------------
INSERT INTO `tbl_tenant` VALUES ('935840632830697474', 'iot', '0', '2017-11-29 19:59:29', '2017-11-29 19:59:29', '923456632830690002', 'iot.cotxnetworks.com', '1', '租户描述', '小王', 'unicom@163.com', '18888888888', 'http://file.radacat.com:88/v1.0/file/', 'http://file.radacat.com:88/v1.0/file/c4aeaeea412385c050813e7abcc8ba6f');
INSERT INTO `tbl_tenant` VALUES ('961777349764845568', 'demo', '0', '2018-02-09 09:42:44', '2018-02-09 09:42:44', '923456632830690002', 'demo.cotxnetworks.com', '1', '', 'demo', 'demo@cotxnetworks.com', '', 'http://file.radacat.com:88/v1.0/file/', 'http://file.radacat.com:88/v1.0/file/c4aeaeea412385c050813e7abcc8ba6f');

/*
Navicat MySQL Data Transfer

Source Server         : 1.6
Source Server Version : 50719
Source Host           : 192.168.1.6:3306
Source Database       : auth

Target Server Type    : MYSQL
Target Server Version : 50719
File Encoding         : 65001

Date: 2018-04-13 16:51:20
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tbl_user_domain_role
-- ----------------------------
DROP TABLE IF EXISTS `tbl_user_domain_role`;
CREATE TABLE `tbl_user_domain_role` (
  `uid` bigint(20) NOT NULL,
  `did` bigint(20) NOT NULL,
  `rid` int(11) NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01',
  `update_time` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01',
  PRIMARY KEY (`uid`,`did`,`rid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tbl_user_domain_role
-- ----------------------------
INSERT INTO `tbl_user_domain_role` VALUES ('935840632809725001', '923456632830690002', '0', '1970-01-01 08:00:01', '1970-01-01 08:00:01');
INSERT INTO `tbl_user_domain_role` VALUES ('935840632809725001', '923456632830690002', '1', '1970-01-01 08:00:01', '1970-01-01 08:00:01');
INSERT INTO `tbl_user_domain_role` VALUES ('935840632809725004', '923456632830690002', '0', '1970-01-01 08:00:01', '1970-01-01 08:00:01');
INSERT INTO `tbl_user_domain_role` VALUES ('935840632809725004', '923456632830690002', '1', '1970-01-01 08:00:01', '1970-01-01 08:00:01');
INSERT INTO `tbl_user_domain_role` VALUES ('935840632809725911', '923456632830690002', '0', '1970-01-01 08:00:01', '1970-01-01 08:00:01');
INSERT INTO `tbl_user_domain_role` VALUES ('935840632809725911', '923456632830690002', '1', '1970-01-01 08:00:01', '1970-01-01 08:00:01');
INSERT INTO `tbl_user_domain_role` VALUES ('959376014343938043', '923456632830690002', '0', '1970-01-01 08:00:01', '1970-01-01 08:00:01');
INSERT INTO `tbl_user_domain_role` VALUES ('959376014343938043', '923456632830690002', '1', '1970-01-01 08:00:01', '1970-01-01 08:00:01');
INSERT INTO `tbl_user_domain_role` VALUES ('959376631343923643', '923456632830690002', '0', '1970-01-01 08:00:01', '1970-01-01 08:00:01');
INSERT INTO `tbl_user_domain_role` VALUES ('959376631343923643', '923456632830690002', '1', '1970-01-01 08:00:01', '1970-01-01 08:00:01');

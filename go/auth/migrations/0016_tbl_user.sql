/*
Navicat MySQL Data Transfer

Source Server         : 1.6
Source Server Version : 50719
Source Host           : 192.168.1.6:3306
Source Database       : auth

Target Server Type    : MYSQL
Target Server Version : 50719
File Encoding         : 65001

Date: 2018-04-13 16:51:07
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tbl_user
-- ----------------------------
DROP TABLE IF EXISTS `tbl_user`;
CREATE TABLE `tbl_user` (
  `uid` bigint(20) NOT NULL,
  `username` varchar(64) CHARACTER SET utf8 NOT NULL,
  `password` char(64) CHARACTER SET utf8 DEFAULT NULL,
  `salt` char(6) CHARACTER SET utf8 DEFAULT NULL,
  `nickname` varchar(128) CHARACTER SET utf8 NOT NULL,
  `email` varchar(128) CHARACTER SET utf8 DEFAULT NULL,
  `phone` varchar(20) CHARACTER SET utf8 DEFAULT NULL,
  `state` int(11) NOT NULL DEFAULT '0' COMMENT '0异常；1被邀请未激活；2不可用；3可用',
  `create_time` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01',
  `update_time` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01',
  `password_strength` int(11) DEFAULT '1' COMMENT '1.弱 2.中 3.强',
  `token` varchar(255) NOT NULL DEFAULT '',
  `tid` bigint(20) DEFAULT '0',
  `did` bigint(20) DEFAULT '0',
  PRIMARY KEY (`uid`),
  UNIQUE KEY `tbl_account_username_uindex` (`username`,`tid`,`did`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tbl_user
-- ----------------------------
INSERT INTO `tbl_user` VALUES ('935840632809725001', 'admin@cotxnetworks.com', '9e960bbba458f239a3f90df3dd7b3cdf', 'FMQYYQ', 'admin', '', '', '3', '1970-01-01 08:00:01', '1970-01-01 08:00:01', '1', '', null, '923456632830690002');
INSERT INTO `tbl_user` VALUES ('935840632809725002', 'admin@cotxnetworks.com', '3edb60e76749188d440c41377164c3ac', 'SHBOXA', 'admin', '', '', '3', '1970-01-01 08:00:01', '1970-01-01 08:00:01', '1', '', '935840632830697474', null);
INSERT INTO `tbl_user` VALUES ('935840632809725003', 'admin@cotxnetworks.com', '3edb60e76749188d440c41377164c3ac', 'SHBOXA', 'admin', '', '', '3', '1970-01-01 08:00:01', '1970-01-01 08:00:01', '1', '', '961777349764845568', null);
INSERT INTO `tbl_user` VALUES ('935840632809725004', 'demo@cotxnetworks.com', '9e960bbba458f239a3f90df3dd7b3cdf', 'FMQYYQ', 'demo', '', '', '3', '1970-01-01 08:00:01', '1970-01-01 08:00:01', '1', '', null, '923456632830690002');
INSERT INTO `tbl_user` VALUES ('935840632809725005', 'demo@cotxnetworks.com', '9e960bbba458f239a3f90df3dd7b3cdf', 'FMQYYQ', 'demo', '', '', '3', '1970-01-01 08:00:01', '1970-01-01 08:00:01', '1', '', '935840632830697474', null);
INSERT INTO `tbl_user` VALUES ('935840632809725006', 'demo@cotxnetworks.com', '9e960bbba458f239a3f90df3dd7b3cdf', 'FMQYYQ', 'demo', '', '', '3', '1970-01-01 08:00:01', '1970-01-01 08:00:01', '1', '', '961777349764845568', null);
INSERT INTO `tbl_user` VALUES ('935840632809725911', 'quyf@radacat.com', 'ae3eb66418a07921902e56da2f106994', 'TPUZVP', 'quyf', '', '', '3', '2018-01-23 11:34:27', '2018-01-23 11:34:27', '1', '', null, '923456632830690002');
INSERT INTO `tbl_user` VALUES ('935840632809725922', 'quyf@radacat.com', 'ae3eb66418a07921902e56da2f106994', 'TPUZVP', 'quyf', '', '', '3', '2018-01-23 11:34:27', '2018-01-23 11:34:27', '1', '', '935840632830697474', null);
INSERT INTO `tbl_user` VALUES ('935840632809725933', 'quyf@radacat.com', 'ae3eb66418a07921902e56da2f106994', 'TPUZVP', 'quyf', '', '', '3', '2018-01-23 11:34:27', '2018-01-23 11:34:27', '1', '', '961777349764845568', null);
INSERT INTO `tbl_user` VALUES ('959376014343938043', 'yujj@cotxnetworks.com', '197ac4ca7255d33cb98630d01d3393fa', 'MEJDGU', 'yujj', null, null, '3', '2018-02-02 18:40:41', '2018-02-02 18:40:41', '1', 'GBJf03wp7ROkW2OAR9ipH507ND5Mt6od', null, '923456632830690002');
INSERT INTO `tbl_user` VALUES ('959376014343938048', 'yujj@cotxnetworks.com', 'c61e03a6cf3e6719fd220654108ae437', 'DNRHRQ', 'yujj', null, null, '3', '2018-02-02 18:40:41', '2018-02-02 18:40:41', '1', 'GBJf03wp7ROkW2OAR9ipH507ND5Mt6od', '935840632830697474', null);
INSERT INTO `tbl_user` VALUES ('959376014343938453', 'yujj@cotxnetworks.com', 'c61e03a6cf3e6719fd220654108ae437', 'DNRHRQ', 'yujj', null, null, '3', '2018-02-02 18:40:41', '2018-02-02 18:40:41', '1', 'GBJf03wp7ROkW2OAR9ipH507ND5Mt6od', '961777349764845568', null);
INSERT INTO `tbl_user` VALUES ('959376631343923643', 'hehh@cotxnetworks.com', 'c61e03a6cf3e6719fd220654108ae437', 'DNRHRQ', 'hehh', null, null, '3', '2018-02-02 18:40:41', '2018-02-02 18:40:41', '1', 'GBJf03wp7ROkW2OAR9ipH507ND5Mt6od', null, '923456632830690002');
INSERT INTO `tbl_user` VALUES ('959376631343936453', 'hehh@cotxnetworks.com', '197ac4ca7255d33cb98630d01d3393fa', 'MEJDGU', 'hehh', null, null, '3', '2018-02-02 18:40:41', '2018-02-02 18:40:41', '1', 'GBJf03wp7ROkW2OAR9ipH507ND5Mt6od', '935840632830697474', '0');
INSERT INTO `tbl_user` VALUES ('959376631343995254', 'hehh@cotxnetworks.com', 'c61e03a6cf3e6719fd220654108ae437', 'DNRHRQ', 'hehh', null, null, '3', '2018-02-02 18:40:41', '2018-02-02 18:40:41', '1', 'GBJf03wp7ROkW2OAR9ipH507ND5Mt6od', '961777349764845568', '0');

/*
Navicat MySQL Data Transfer

Source Server         : 1.6
Source Server Version : 50719
Source Host           : 192.168.1.6:3306
Source Database       : auth

Target Server Type    : MYSQL
Target Server Version : 50719
File Encoding         : 65001

Date: 2018-04-13 16:51:32
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tbl_user_tenant_role
-- ----------------------------
DROP TABLE IF EXISTS `tbl_user_tenant_role`;
CREATE TABLE `tbl_user_tenant_role` (
  `uid` bigint(20) NOT NULL,
  `tid` bigint(20) NOT NULL,
  `rid` int(11) NOT NULL,
  `isDefault` tinyint(4) DEFAULT NULL,
  `create_time` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01',
  `update_time` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01',
  PRIMARY KEY (`uid`,`tid`,`rid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tbl_user_tenant_role
-- ----------------------------
INSERT INTO `tbl_user_tenant_role` VALUES ('935840632809725002', '935840632830697474', '0', '1', '1970-01-01 08:00:01', '1970-01-01 08:00:01');
INSERT INTO `tbl_user_tenant_role` VALUES ('935840632809725002', '935840632830697474', '2', '1', '1970-01-01 08:00:01', '1970-01-01 08:00:01');
INSERT INTO `tbl_user_tenant_role` VALUES ('935840632809725003', '961777349764845568', '0', '1', '1970-01-01 08:00:01', '1970-01-01 08:00:01');
INSERT INTO `tbl_user_tenant_role` VALUES ('935840632809725003', '961777349764845568', '3', '1', '1970-01-01 08:00:01', '1970-01-01 08:00:01');
INSERT INTO `tbl_user_tenant_role` VALUES ('935840632809725005', '935840632830697474', '0', '1', '1970-01-01 08:00:01', '1970-01-01 08:00:01');
INSERT INTO `tbl_user_tenant_role` VALUES ('935840632809725005', '935840632830697474', '2', '1', '1970-01-01 08:00:01', '1970-01-01 08:00:01');
INSERT INTO `tbl_user_tenant_role` VALUES ('935840632809725006', '961777349764845568', '0', '1', '1970-01-01 08:00:01', '1970-01-01 08:00:01');
INSERT INTO `tbl_user_tenant_role` VALUES ('935840632809725006', '961777349764845568', '3', '1', '1970-01-01 08:00:01', '1970-01-01 08:00:01');
INSERT INTO `tbl_user_tenant_role` VALUES ('935840632809725922', '935840632830697474', '0', '1', '1970-01-01 08:00:01', '1970-01-01 08:00:01');
INSERT INTO `tbl_user_tenant_role` VALUES ('935840632809725922', '935840632830697474', '2', '1', '1970-01-01 08:00:01', '1970-01-01 08:00:01');
INSERT INTO `tbl_user_tenant_role` VALUES ('935840632809725933', '961777349764845568', '0', '1', '1970-01-01 08:00:01', '1970-01-01 08:00:01');
INSERT INTO `tbl_user_tenant_role` VALUES ('935840632809725933', '961777349764845568', '3', '1', '1970-01-01 08:00:01', '1970-01-01 08:00:01');
INSERT INTO `tbl_user_tenant_role` VALUES ('959376014343938048', '935840632830697474', '0', '1', '1970-01-01 08:00:01', '1970-01-01 08:00:01');
INSERT INTO `tbl_user_tenant_role` VALUES ('959376014343938048', '935840632830697474', '2', '1', '1970-01-01 08:00:01', '1970-01-01 08:00:01');
INSERT INTO `tbl_user_tenant_role` VALUES ('959376014343938453', '961777349764845568', '0', '1', '1970-01-01 08:00:01', '1970-01-01 08:00:01');
INSERT INTO `tbl_user_tenant_role` VALUES ('959376014343938453', '961777349764845568', '3', '1', '1970-01-01 08:00:01', '1970-01-01 08:00:01');
INSERT INTO `tbl_user_tenant_role` VALUES ('959376631343936453', '935840632830697474', '0', '1', '1970-01-01 08:00:01', '1970-01-01 08:00:01');
INSERT INTO `tbl_user_tenant_role` VALUES ('959376631343936453', '935840632830697474', '2', '1', '1970-01-01 08:00:01', '1970-01-01 08:00:01');
INSERT INTO `tbl_user_tenant_role` VALUES ('959376631343995254', '961777349764845568', '0', '1', '1970-01-01 08:00:01', '1970-01-01 08:00:01');
INSERT INTO `tbl_user_tenant_role` VALUES ('959376631343995254', '961777349764845568', '3', '1', '1970-01-01 08:00:01', '1970-01-01 08:00:01');

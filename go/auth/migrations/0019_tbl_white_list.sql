/*
Navicat MySQL Data Transfer

Source Server         : 1.6
Source Server Version : 50719
Source Host           : 192.168.1.6:3306
Source Database       : auth

Target Server Type    : MYSQL
Target Server Version : 50719
File Encoding         : 65001

Date: 2018-04-13 16:51:55
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tbl_white_list
-- ----------------------------
DROP TABLE IF EXISTS `tbl_white_list`;
CREATE TABLE `tbl_white_list` (
  `id` int(11) NOT NULL,
  `permissionName` varchar(255) DEFAULT NULL,
  `url` varchar(255) NOT NULL,
  `opt` varchar(255) NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' ON UPDATE CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tbl_white_list
-- ----------------------------
INSERT INTO `tbl_white_list` VALUES ('1', '登录域', '/v1.1/tokens/domain/([0-9])+', 'POST', '2018-01-31 11:06:32', '2018-01-31 11:06:32');
INSERT INTO `tbl_white_list` VALUES ('2', '登录租户', '/v1.1/tokens', 'POST', '2018-02-01 17:50:09', '2018-02-01 17:50:09');
INSERT INTO `tbl_white_list` VALUES ('3', '获取用户信息', '/v1.1/tokens', 'GET', '2018-01-31 11:05:33', '2018-01-31 11:05:33');
INSERT INTO `tbl_white_list` VALUES ('4', '登出', '/v1.1/tokens', 'DELETE', '2018-01-31 11:05:41', '2018-01-31 11:05:41');
INSERT INTO `tbl_white_list` VALUES ('5', '修改用户信息', '/v1.1/accounts/([0-9])+', 'PUT', '2018-02-01 17:50:45', '2018-02-01 17:50:45');
INSERT INTO `tbl_white_list` VALUES ('6', '修改密码', '/v1.1/accounts/([0-9])+/password', 'PUT', '2018-02-01 17:50:48', '2018-02-01 17:50:48');
INSERT INTO `tbl_white_list` VALUES ('7', '忘记密码', '/v1.1/message/password', 'POST', '2018-02-01 18:38:40', '2018-02-01 18:38:40');
INSERT INTO `tbl_white_list` VALUES ('8', '重置密码', '/v1.1/message/password/[0-9a-zA-Z]+', 'POST', '2018-02-01 18:39:55', '2018-02-01 18:39:55');
INSERT INTO `tbl_white_list` VALUES ('33', '通过URL获取租户ID', '/v1.1/logos/tenants', 'POST', '2018-03-08 18:22:01', '2018-03-08 18:22:01');
INSERT INTO `tbl_white_list` VALUES ('15', 'TEST同意进入租户', '/v1.1/tenants', 'PUT', '2018-02-01 15:14:02', '2018-02-01 15:14:02');
INSERT INTO `tbl_white_list` VALUES ('28', 'TEST同意进入域', '/v1.1/domains', 'PUT', '2018-02-01 15:14:02', '2018-02-01 15:14:02');
INSERT INTO `tbl_white_list` VALUES ('52', '通过tid获取did', '/v1.1/message/tenants/([0-9])+', 'GET', '2018-04-09 09:57:12', '2018-04-09 09:57:12');

/*
Navicat MySQL Data Transfer

Source Server         : 1.6
Source Server Version : 50719
Source Host           : 192.168.1.6:3306
Source Database       : auth

Target Server Type    : MYSQL
Target Server Version : 50719
File Encoding         : 65001

Date: 2018-04-13 16:50:34
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tbl_service
-- ----------------------------
DROP TABLE IF EXISTS `tbl_service`;
CREATE TABLE `tbl_service` (
  `service_id` int(11) NOT NULL AUTO_INCREMENT,
  `service_name` varchar(255) NOT NULL DEFAULT '',
  `service_key` varchar(255) NOT NULL DEFAULT '',
  `service_url` varchar(255) NOT NULL DEFAULT '',
  `service_type` int(11) NOT NULL COMMENT '1.本地服务 2.外部服务 3.系统服务',
  `service_tid` bigint(20) NOT NULL,
  `service_state` int(11) NOT NULL DEFAULT '1' COMMENT '1.启用 2.停用',
  `service_description` varchar(255) DEFAULT '',
  `update_time` timestamp NULL DEFAULT '1970-01-01 08:00:01' ON UPDATE CURRENT_TIMESTAMP,
  `create_time` timestamp NULL DEFAULT '1970-01-01 08:00:01' ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`service_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tbl_service
-- ----------------------------
INSERT INTO `tbl_service` VALUES ('1', '设备管理', '', '', '3', '935840632830697474', '1', '物联网设备接入，数据解析服务', '2018-03-21 15:28:11', '2018-03-21 15:28:11');
INSERT INTO `tbl_service` VALUES ('2', 'KMS服务', '', '', '3', '935840632830697474', '1', '物联网应用防火墙，防火墙对物联网数据进行规则匹配，访问控制，以及数据转发工作。', '2018-03-21 15:28:18', '2018-03-21 15:28:18');

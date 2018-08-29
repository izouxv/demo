/*
Navicat MySQL Data Transfer

Source Server         : 1.6
Source Server Version : 50719
Source Host           : 192.168.1.6:3306
Source Database       : auth

Target Server Type    : MYSQL
Target Server Version : 50719
File Encoding         : 65001

Date: 2018-04-13 16:50:56
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tbl_tenant_service_policy
-- ----------------------------
DROP TABLE IF EXISTS `tbl_tenant_service_policy`;
CREATE TABLE `tbl_tenant_service_policy` (
  `sid` int(11) NOT NULL,
  `tid` bigint(20) NOT NULL,
  `pid` int(11) NOT NULL,
  `start_time` timestamp NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`sid`,`tid`,`pid`),
  KEY `tid` (`tid`),
  KEY `pid` (`pid`),
  CONSTRAINT `pid` FOREIGN KEY (`pid`) REFERENCES `tbl_policy` (`policy_id`),
  CONSTRAINT `sid` FOREIGN KEY (`sid`) REFERENCES `tbl_service` (`service_id`),
  CONSTRAINT `tid` FOREIGN KEY (`tid`) REFERENCES `tbl_tenant` (`tid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tbl_tenant_service_policy
-- ----------------------------
INSERT INTO `tbl_tenant_service_policy` VALUES ('1', '935840632830697474', '1', '0000-00-00 00:00:00');
INSERT INTO `tbl_tenant_service_policy` VALUES ('2', '935840632830697474', '2', '0000-00-00 00:00:00');

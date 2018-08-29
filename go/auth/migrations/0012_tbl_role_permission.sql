/*
Navicat MySQL Data Transfer

Source Server         : 1.6
Source Server Version : 50719
Source Host           : 192.168.1.6:3306
Source Database       : auth

Target Server Type    : MYSQL
Target Server Version : 50719
File Encoding         : 65001

Date: 2018-04-13 16:50:23
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tbl_role_permission
-- ----------------------------
DROP TABLE IF EXISTS `tbl_role_permission`;
CREATE TABLE `tbl_role_permission` (
  `rid` int(11) NOT NULL,
  `per_id` int(11) NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`rid`,`per_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tbl_role_permission
-- ----------------------------
INSERT INTO `tbl_role_permission` VALUES ('1', '41', '2018-04-09 16:19:50', '2018-04-09 16:19:50');
INSERT INTO `tbl_role_permission` VALUES ('1', '42', '2018-04-09 16:19:51', '2018-04-09 16:19:51');
INSERT INTO `tbl_role_permission` VALUES ('1', '43', '2018-04-09 16:19:52', '2018-04-09 16:19:52');
INSERT INTO `tbl_role_permission` VALUES ('1', '44', '2018-04-09 16:19:52', '2018-04-09 16:19:52');
INSERT INTO `tbl_role_permission` VALUES ('1', '45', '2018-04-09 16:19:53', '2018-04-09 16:19:53');
INSERT INTO `tbl_role_permission` VALUES ('1', '46', '2018-04-09 16:19:54', '2018-04-09 16:19:54');
INSERT INTO `tbl_role_permission` VALUES ('1', '47', '2018-04-09 16:19:54', '2018-04-09 16:19:54');
INSERT INTO `tbl_role_permission` VALUES ('1', '48', '2018-04-09 16:19:55', '2018-04-09 16:19:55');
INSERT INTO `tbl_role_permission` VALUES ('1', '49', '2018-04-09 16:19:56', '2018-04-09 16:19:56');
INSERT INTO `tbl_role_permission` VALUES ('1', '50', '2018-04-09 16:20:00', '2018-04-09 16:20:00');
INSERT INTO `tbl_role_permission` VALUES ('1', '51', '2018-04-09 16:20:01', '2018-04-09 16:20:01');
INSERT INTO `tbl_role_permission` VALUES ('1', '52', '2018-04-09 16:20:02', '2018-04-09 16:20:02');
INSERT INTO `tbl_role_permission` VALUES ('1', '53', '2018-04-09 16:20:03', '2018-04-09 16:20:03');
INSERT INTO `tbl_role_permission` VALUES ('1', '54', '2018-04-09 16:20:08', '2018-04-09 16:20:08');
INSERT INTO `tbl_role_permission` VALUES ('1', '55', '2018-04-09 16:20:08', '2018-04-09 16:20:08');
INSERT INTO `tbl_role_permission` VALUES ('1', '56', '2018-04-09 16:20:10', '2018-04-09 16:20:10');
INSERT INTO `tbl_role_permission` VALUES ('1', '57', '2018-04-09 16:20:11', '2018-04-09 16:20:11');
INSERT INTO `tbl_role_permission` VALUES ('1', '58', '2018-04-09 16:20:13', '2018-04-09 16:20:13');
INSERT INTO `tbl_role_permission` VALUES ('1', '59', '2018-04-09 16:20:15', '2018-04-09 16:20:15');
INSERT INTO `tbl_role_permission` VALUES ('1', '60', '2018-04-09 16:20:16', '2018-04-09 16:20:16');
INSERT INTO `tbl_role_permission` VALUES ('1', '61', '2018-04-09 16:20:18', '2018-04-09 16:20:18');
INSERT INTO `tbl_role_permission` VALUES ('1', '62', '2018-04-09 16:20:19', '2018-04-09 16:20:19');
INSERT INTO `tbl_role_permission` VALUES ('1', '63', '2018-04-09 16:20:20', '2018-04-09 16:20:20');
INSERT INTO `tbl_role_permission` VALUES ('1', '64', '2018-04-09 16:20:21', '2018-04-09 16:20:21');
INSERT INTO `tbl_role_permission` VALUES ('1', '65', '2018-04-09 16:20:23', '2018-04-09 16:20:23');
INSERT INTO `tbl_role_permission` VALUES ('1', '66', '2018-04-09 16:20:24', '2018-04-09 16:20:24');
INSERT INTO `tbl_role_permission` VALUES ('1', '67', '2018-04-09 16:20:26', '2018-04-09 16:20:26');
INSERT INTO `tbl_role_permission` VALUES ('1', '68', '2018-04-09 16:20:28', '2018-04-09 16:20:28');
INSERT INTO `tbl_role_permission` VALUES ('1', '69', '2018-04-09 16:20:30', '2018-04-09 16:20:30');
INSERT INTO `tbl_role_permission` VALUES ('1', '70', '2018-04-09 16:20:32', '2018-04-09 16:20:32');
INSERT INTO `tbl_role_permission` VALUES ('1', '71', '2018-04-09 16:20:34', '2018-04-09 16:20:34');
INSERT INTO `tbl_role_permission` VALUES ('1', '72', '2018-04-09 16:20:35', '2018-04-09 16:20:35');
INSERT INTO `tbl_role_permission` VALUES ('1', '73', '2018-04-09 16:20:36', '2018-04-09 16:20:36');
INSERT INTO `tbl_role_permission` VALUES ('1', '74', '2018-04-09 16:20:37', '2018-04-09 16:20:37');
INSERT INTO `tbl_role_permission` VALUES ('1', '75', '2018-04-09 16:20:38', '2018-04-09 16:20:38');
INSERT INTO `tbl_role_permission` VALUES ('1', '76', '2018-04-09 16:20:42', '2018-04-09 16:20:42');
INSERT INTO `tbl_role_permission` VALUES ('1', '77', '2018-04-09 16:20:44', '2018-04-09 16:20:44');
INSERT INTO `tbl_role_permission` VALUES ('1', '78', '2018-04-09 16:20:46', '2018-04-09 16:20:46');
INSERT INTO `tbl_role_permission` VALUES ('1', '79', '2018-04-09 16:20:48', '2018-04-09 16:20:48');
INSERT INTO `tbl_role_permission` VALUES ('1', '80', '2018-04-09 16:20:50', '2018-04-09 16:20:50');
INSERT INTO `tbl_role_permission` VALUES ('1', '89', '2018-04-09 16:19:50', '2018-04-09 16:19:50');
INSERT INTO `tbl_role_permission` VALUES ('1', '90', '2018-04-09 16:19:51', '2018-04-09 16:19:51');
INSERT INTO `tbl_role_permission` VALUES ('1', '91', '2018-04-09 16:19:52', '2018-04-09 16:19:52');
INSERT INTO `tbl_role_permission` VALUES ('1', '92', '2018-04-09 16:19:52', '2018-04-09 16:19:52');
INSERT INTO `tbl_role_permission` VALUES ('2', '1', '2018-04-09 16:19:50', '2018-04-09 16:19:50');
INSERT INTO `tbl_role_permission` VALUES ('2', '2', '2018-04-09 16:19:51', '2018-04-09 16:19:51');
INSERT INTO `tbl_role_permission` VALUES ('2', '3', '2018-04-09 16:19:52', '2018-04-09 16:19:52');
INSERT INTO `tbl_role_permission` VALUES ('2', '4', '2018-04-09 16:19:52', '2018-04-09 16:19:52');
INSERT INTO `tbl_role_permission` VALUES ('2', '5', '2018-04-09 16:19:53', '2018-04-09 16:19:53');
INSERT INTO `tbl_role_permission` VALUES ('2', '6', '2018-04-09 16:19:54', '2018-04-09 16:19:54');
INSERT INTO `tbl_role_permission` VALUES ('2', '7', '2018-04-09 16:19:54', '2018-04-09 16:19:54');
INSERT INTO `tbl_role_permission` VALUES ('2', '8', '2018-04-09 16:19:55', '2018-04-09 16:19:55');
INSERT INTO `tbl_role_permission` VALUES ('2', '9', '2018-04-09 16:19:56', '2018-04-09 16:19:56');
INSERT INTO `tbl_role_permission` VALUES ('2', '10', '2018-04-09 16:20:00', '2018-04-09 16:20:00');
INSERT INTO `tbl_role_permission` VALUES ('2', '11', '2018-04-09 16:20:01', '2018-04-09 16:20:01');
INSERT INTO `tbl_role_permission` VALUES ('2', '12', '2018-04-09 16:20:02', '2018-04-09 16:20:02');
INSERT INTO `tbl_role_permission` VALUES ('2', '13', '2018-04-09 16:20:03', '2018-04-09 16:20:03');
INSERT INTO `tbl_role_permission` VALUES ('2', '14', '2018-04-09 16:19:50', '2018-04-09 16:19:50');
INSERT INTO `tbl_role_permission` VALUES ('2', '15', '2018-04-09 16:20:08', '2018-04-09 16:20:08');
INSERT INTO `tbl_role_permission` VALUES ('2', '16', '2018-04-09 16:20:10', '2018-04-09 16:20:10');
INSERT INTO `tbl_role_permission` VALUES ('2', '17', '2018-04-09 16:20:11', '2018-04-09 16:20:11');
INSERT INTO `tbl_role_permission` VALUES ('2', '18', '2018-04-09 16:20:13', '2018-04-09 16:20:13');
INSERT INTO `tbl_role_permission` VALUES ('2', '19', '2018-04-09 16:20:15', '2018-04-09 16:20:15');
INSERT INTO `tbl_role_permission` VALUES ('2', '20', '2018-04-09 16:20:16', '2018-04-09 16:20:16');
INSERT INTO `tbl_role_permission` VALUES ('2', '21', '2018-04-09 16:20:18', '2018-04-09 16:20:18');
INSERT INTO `tbl_role_permission` VALUES ('2', '22', '2018-04-09 16:20:19', '2018-04-09 16:20:19');
INSERT INTO `tbl_role_permission` VALUES ('2', '23', '2018-04-09 16:20:20', '2018-04-09 16:20:20');
INSERT INTO `tbl_role_permission` VALUES ('2', '24', '2018-04-09 16:20:21', '2018-04-09 16:20:21');
INSERT INTO `tbl_role_permission` VALUES ('2', '25', '2018-04-09 16:20:23', '2018-04-09 16:20:23');
INSERT INTO `tbl_role_permission` VALUES ('2', '26', '2018-04-09 16:20:24', '2018-04-09 16:20:24');
INSERT INTO `tbl_role_permission` VALUES ('2', '27', '2018-04-09 16:20:26', '2018-04-09 16:20:26');
INSERT INTO `tbl_role_permission` VALUES ('2', '28', '2018-04-09 16:20:28', '2018-04-09 16:20:28');
INSERT INTO `tbl_role_permission` VALUES ('2', '29', '2018-04-09 16:20:30', '2018-04-09 16:20:30');
INSERT INTO `tbl_role_permission` VALUES ('2', '30', '2018-04-09 16:20:32', '2018-04-09 16:20:32');
INSERT INTO `tbl_role_permission` VALUES ('2', '31', '2018-04-09 16:20:34', '2018-04-09 16:20:34');
INSERT INTO `tbl_role_permission` VALUES ('2', '32', '2018-04-09 16:20:35', '2018-04-09 16:20:35');
INSERT INTO `tbl_role_permission` VALUES ('2', '33', '2018-04-09 16:20:36', '2018-04-09 16:20:36');
INSERT INTO `tbl_role_permission` VALUES ('2', '34', '2018-04-09 16:20:37', '2018-04-09 16:20:37');
INSERT INTO `tbl_role_permission` VALUES ('2', '35', '2018-04-09 16:20:38', '2018-04-09 16:20:38');
INSERT INTO `tbl_role_permission` VALUES ('2', '36', '2018-04-09 16:20:42', '2018-04-09 16:20:42');
INSERT INTO `tbl_role_permission` VALUES ('2', '37', '2018-04-09 16:20:44', '2018-04-09 16:20:44');
INSERT INTO `tbl_role_permission` VALUES ('2', '38', '2018-04-09 16:20:46', '2018-04-09 16:20:46');
INSERT INTO `tbl_role_permission` VALUES ('2', '39', '2018-04-09 16:20:48', '2018-04-09 16:20:48');
INSERT INTO `tbl_role_permission` VALUES ('2', '40', '2018-04-09 16:20:50', '2018-04-09 16:20:50');
INSERT INTO `tbl_role_permission` VALUES ('2', '81', '2018-04-09 16:19:50', '2018-04-09 16:19:50');
INSERT INTO `tbl_role_permission` VALUES ('2', '82', '2018-04-09 16:19:51', '2018-04-09 16:19:51');
INSERT INTO `tbl_role_permission` VALUES ('2', '83', '2018-04-09 16:19:52', '2018-04-09 16:19:52');
INSERT INTO `tbl_role_permission` VALUES ('2', '84', '2018-04-09 16:19:52', '2018-04-09 16:19:52');
INSERT INTO `tbl_role_permission` VALUES ('2', '85', '2018-04-09 16:19:50', '2018-04-09 16:19:50');
INSERT INTO `tbl_role_permission` VALUES ('2', '86', '2018-04-09 16:19:51', '2018-04-09 16:19:51');
INSERT INTO `tbl_role_permission` VALUES ('2', '87', '2018-04-09 16:19:52', '2018-04-09 16:19:52');
INSERT INTO `tbl_role_permission` VALUES ('2', '88', '2018-04-09 16:19:52', '2018-04-09 16:19:52');
INSERT INTO `tbl_role_permission` VALUES ('3', '1', '2018-04-09 16:19:50', '2018-04-09 16:19:50');
INSERT INTO `tbl_role_permission` VALUES ('3', '2', '2018-04-09 16:19:51', '2018-04-09 16:19:51');
INSERT INTO `tbl_role_permission` VALUES ('3', '3', '2018-04-09 16:19:52', '2018-04-09 16:19:52');
INSERT INTO `tbl_role_permission` VALUES ('3', '4', '2018-04-09 16:19:52', '2018-04-09 16:19:52');
INSERT INTO `tbl_role_permission` VALUES ('3', '5', '2018-04-09 16:19:53', '2018-04-09 16:19:53');
INSERT INTO `tbl_role_permission` VALUES ('3', '6', '2018-04-09 16:19:54', '2018-04-09 16:19:54');
INSERT INTO `tbl_role_permission` VALUES ('3', '7', '2018-04-09 16:19:54', '2018-04-09 16:19:54');
INSERT INTO `tbl_role_permission` VALUES ('3', '8', '2018-04-09 16:19:55', '2018-04-09 16:19:55');
INSERT INTO `tbl_role_permission` VALUES ('3', '9', '2018-04-09 16:19:56', '2018-04-09 16:19:56');
INSERT INTO `tbl_role_permission` VALUES ('3', '10', '2018-04-09 16:20:00', '2018-04-09 16:20:00');
INSERT INTO `tbl_role_permission` VALUES ('3', '11', '2018-04-09 16:20:01', '2018-04-09 16:20:01');
INSERT INTO `tbl_role_permission` VALUES ('3', '12', '2018-04-09 16:20:02', '2018-04-09 16:20:02');
INSERT INTO `tbl_role_permission` VALUES ('3', '13', '2018-04-09 16:20:03', '2018-04-09 16:20:03');
INSERT INTO `tbl_role_permission` VALUES ('3', '14', '2018-04-09 16:19:50', '2018-04-09 16:19:50');
INSERT INTO `tbl_role_permission` VALUES ('3', '15', '2018-04-09 16:20:08', '2018-04-09 16:20:08');
INSERT INTO `tbl_role_permission` VALUES ('3', '16', '2018-04-09 16:20:10', '2018-04-09 16:20:10');
INSERT INTO `tbl_role_permission` VALUES ('3', '17', '2018-04-09 16:20:11', '2018-04-09 16:20:11');
INSERT INTO `tbl_role_permission` VALUES ('3', '18', '2018-04-09 16:20:13', '2018-04-09 16:20:13');
INSERT INTO `tbl_role_permission` VALUES ('3', '19', '2018-04-09 16:20:15', '2018-04-09 16:20:15');
INSERT INTO `tbl_role_permission` VALUES ('3', '20', '2018-04-09 16:20:16', '2018-04-09 16:20:16');
INSERT INTO `tbl_role_permission` VALUES ('3', '21', '2018-04-09 16:20:18', '2018-04-09 16:20:18');
INSERT INTO `tbl_role_permission` VALUES ('3', '22', '2018-04-09 16:20:19', '2018-04-09 16:20:19');
INSERT INTO `tbl_role_permission` VALUES ('3', '23', '2018-04-09 16:20:20', '2018-04-09 16:20:20');
INSERT INTO `tbl_role_permission` VALUES ('3', '24', '2018-04-09 16:20:21', '2018-04-09 16:20:21');
INSERT INTO `tbl_role_permission` VALUES ('3', '25', '2018-04-09 16:20:23', '2018-04-09 16:20:23');
INSERT INTO `tbl_role_permission` VALUES ('3', '26', '2018-04-09 16:20:24', '2018-04-09 16:20:24');
INSERT INTO `tbl_role_permission` VALUES ('3', '27', '2018-04-09 16:20:26', '2018-04-09 16:20:26');
INSERT INTO `tbl_role_permission` VALUES ('3', '28', '2018-04-09 16:20:28', '2018-04-09 16:20:28');
INSERT INTO `tbl_role_permission` VALUES ('3', '29', '2018-04-09 16:20:30', '2018-04-09 16:20:30');
INSERT INTO `tbl_role_permission` VALUES ('3', '30', '2018-04-09 16:20:32', '2018-04-09 16:20:32');
INSERT INTO `tbl_role_permission` VALUES ('3', '31', '2018-04-09 16:20:34', '2018-04-09 16:20:34');
INSERT INTO `tbl_role_permission` VALUES ('3', '32', '2018-04-09 16:20:35', '2018-04-09 16:20:35');
INSERT INTO `tbl_role_permission` VALUES ('3', '33', '2018-04-09 16:20:36', '2018-04-09 16:20:36');
INSERT INTO `tbl_role_permission` VALUES ('3', '34', '2018-04-09 16:20:37', '2018-04-09 16:20:37');
INSERT INTO `tbl_role_permission` VALUES ('3', '35', '2018-04-09 16:20:38', '2018-04-09 16:20:38');
INSERT INTO `tbl_role_permission` VALUES ('3', '36', '2018-04-09 16:20:42', '2018-04-09 16:20:42');
INSERT INTO `tbl_role_permission` VALUES ('3', '37', '2018-04-09 16:20:44', '2018-04-09 16:20:44');
INSERT INTO `tbl_role_permission` VALUES ('3', '38', '2018-04-09 16:20:46', '2018-04-09 16:20:46');
INSERT INTO `tbl_role_permission` VALUES ('3', '39', '2018-04-09 16:20:48', '2018-04-09 16:20:48');
INSERT INTO `tbl_role_permission` VALUES ('3', '40', '2018-04-09 16:20:50', '2018-04-09 16:20:50');
INSERT INTO `tbl_role_permission` VALUES ('3', '81', '2018-04-09 16:19:50', '2018-04-09 16:19:50');
INSERT INTO `tbl_role_permission` VALUES ('3', '82', '2018-04-09 16:19:51', '2018-04-09 16:19:51');
INSERT INTO `tbl_role_permission` VALUES ('3', '83', '2018-04-09 16:19:52', '2018-04-09 16:19:52');
INSERT INTO `tbl_role_permission` VALUES ('3', '84', '2018-04-09 16:19:52', '2018-04-09 16:19:52');
INSERT INTO `tbl_role_permission` VALUES ('3', '85', '2018-04-09 16:19:50', '2018-04-09 16:19:50');
INSERT INTO `tbl_role_permission` VALUES ('3', '86', '2018-04-09 16:19:51', '2018-04-09 16:19:51');
INSERT INTO `tbl_role_permission` VALUES ('3', '87', '2018-04-09 16:19:52', '2018-04-09 16:19:52');
INSERT INTO `tbl_role_permission` VALUES ('3', '88', '2018-04-09 16:19:52', '2018-04-09 16:19:52');

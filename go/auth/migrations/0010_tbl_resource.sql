CREATE TABLE `tbl_resource` (
  `res_id` int(11) NOT NULL AUTO_INCREMENT,
  `res_name` varchar(255) DEFAULT NULL,
  `res_url` varchar(255) DEFAULT NULL,
  `res_opt` varchar(255) DEFAULT NULL,
  `res_project` varchar(255) DEFAULT '',
  PRIMARY KEY (`res_id`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tbl_resource
-- ----------------------------
INSERT INTO `tbl_resource` VALUES ('1', '新建群组', '/v1.0/tenants/([0-9])+/groups', 'POST', 'dms');
INSERT INTO `tbl_resource` VALUES ('2', '获取群组列表', '/v1.0/tenants/([0-9])+/groups', 'GET', 'dms');
INSERT INTO `tbl_resource` VALUES ('3', '删除群组', '/v1.0/tenants/([0-9])+/groups/([0-9])+', 'DELETE', 'dms');
INSERT INTO `tbl_resource` VALUES ('4', '修改群组信息', '/v1.0/tenants/([0-9])+/groups/([0-9])+', 'PUT', 'dms');
INSERT INTO `tbl_resource` VALUES ('5', '查询指定群组信息', '/v1.0/tenants/([0-9])+/groups/([0-9])+', 'GET', 'dms');
INSERT INTO `tbl_resource` VALUES ('6', '新建资产', '/v1.0/tenants/([0-9])+/assets', 'POST', 'dms');
INSERT INTO `tbl_resource` VALUES ('7', '获取资产列表', '/v1.0/tenants/([0-9])+/assets', 'GET', 'dms');
INSERT INTO `tbl_resource` VALUES ('8', '删除资产', '/v1.0/tenants/([0-9])+/assets/([0-9])+', 'DELETE', 'dms');
INSERT INTO `tbl_resource` VALUES ('9', '修改资产', '/v1.0/tenants/([0-9])+/assets/([0-9])+', 'PUT', 'dms');
INSERT INTO `tbl_resource` VALUES ('10', '查询指定资产信息', '/v1.0/tenants/([0-9])+/assets/([0-9])+', 'GET', 'dms');
INSERT INTO `tbl_resource` VALUES ('11', '批量绑定资产', '/v1.0/tenants/([0-9])+/groups/([0-9])+/assets', 'POST', 'dms');
INSERT INTO `tbl_resource` VALUES ('12', '批量解绑资产', '/v1.0/tenants/([0-9])+/groups/([0-9])+/assets', 'DELETE', 'dms');
INSERT INTO `tbl_resource` VALUES ('13', '获取绑定（未绑定）资产', '/v1.0/tenants/([0-9])+/groups/([0-9])+/assets', 'GET', 'dms');
INSERT INTO `tbl_resource` VALUES ('14', '新建数据点', '/v1.0/tenants/([0-9])+/datapoints', 'POST', 'dms');
INSERT INTO `tbl_resource` VALUES ('15', '获取数据点列表', '/v1.0/tenants/([0-9])+/datapoints', 'GET', 'dms');
INSERT INTO `tbl_resource` VALUES ('16', '删除数据点', '/v1.0/tenants/([0-9])+/datapoints/([0-9])+', 'DELETE', 'dms');
INSERT INTO `tbl_resource` VALUES ('17', '修改数据点', '/v1.0/tenants/([0-9])+/datapoints/([0-9])+', 'PUT', 'dms');
INSERT INTO `tbl_resource` VALUES ('18', '获取指定数据点', '/v1.0/tenants/([0-9])+/datapoints/([0-9])+', 'GET', 'dms');
INSERT INTO `tbl_resource` VALUES ('19', '获取群组下数据点', '/v1.0/tenants/([0-9])+/groups/([0-9])+/datapoints', 'GET', 'dms');
INSERT INTO `tbl_resource` VALUES ('20', '获取资产下数据点', '/v1.0/tenants/([0-9])+/assets/([0-9])+/datapoints', 'GET', 'dms');
INSERT INTO `tbl_resource` VALUES ('21', '群组下新建资产', '/v1.0/tenants/([0-9])+/groups/([0-9])+/assets/([0-9])+', 'POST', 'dms');
INSERT INTO `tbl_resource` VALUES ('22', '删除群组下资产', '/v1.0/tenants/([0-9])+/groups/([0-9])+/assets/([0-9])+', 'DELETE', 'dms');
INSERT INTO `tbl_resource` VALUES ('23', '获取appKey', '/v1.0/tenants/([0-9])+/appkey', 'GET', 'dms');
INSERT INTO `tbl_resource` VALUES ('24', '新建一条事件告警', 'v1.0/tenants/([0-9])+/rules', 'POST', 'dms');
INSERT INTO `tbl_resource` VALUES ('25', '获取事件告警列表', 'v1.0/tenants/([0-9])+/rules', 'GET', 'dms');
INSERT INTO `tbl_resource` VALUES ('26', '删除一条事件告警', 'v1.0/tenants/([0-9])+/rules/([0-9])+', 'DELETE', 'dms');
INSERT INTO `tbl_resource` VALUES ('27', '更新一条事件告警', 'v1.0/tenants/([0-9])+/rules/([0-9])+\r', 'POST', 'dms');
INSERT INTO `tbl_resource` VALUES ('28', '获取一条事件告警', 'v1.0/tenants/([0-9])+/rules/([0-9])+', 'GET', 'dms');
INSERT INTO `tbl_resource` VALUES ('29', '基于资产id获取事件告警', 'v1.0/tenants/([0-9])+/assets/([0-9])+/rules', 'GET', 'dms');
INSERT INTO `tbl_resource` VALUES ('30', '获取数据记录列表', 'v1.0/tenants/([0-9])+/records', 'GET', 'dms');
INSERT INTO `tbl_resource` VALUES ('31', '删除一条数据记录', 'v1.0/tenants/([0-9])+/records/([0-9])+', 'DELETE', 'dms');
INSERT INTO `tbl_resource` VALUES ('32', '获取单条数据记录', 'v1.0/tenants/([0-9])+/records/([0-9])+', 'GET', 'dms');
INSERT INTO `tbl_resource` VALUES ('33', '基于组id和起止时间获取数据记录', 'v1.0/tenants/([0-9])+/groups/([0-9])+/records', 'GET', 'dms');
INSERT INTO `tbl_resource` VALUES ('34', '基于资产id和起止时间获取记录', 'v1.0/tenants/([0-9])+/assets/([0-9])+/records', 'GET', 'dms');
INSERT INTO `tbl_resource` VALUES ('35', '基于资产id获取数字孪生数据', 'v1.0/tenants/([0-9])+/twins/([0-9])+', 'GET', 'dms');
INSERT INTO `tbl_resource` VALUES ('36', '基于资产id更新数字孪生数据', 'v1.0/tenants/([0-9])+/twins/([0-9])+', 'PUT', 'dms');
INSERT INTO `tbl_resource` VALUES ('37', '添加版本', '/v1.1/tenant/([0-9])+/radacat/versions', 'POST', 'domain');
INSERT INTO `tbl_resource` VALUES ('38', '修改版本', '/v1.1/tenant/([0-9])+/radacat/versions/([0-9])+', 'PUT', 'domain');
INSERT INTO `tbl_resource` VALUES ('39', '获取全部版本', '/v1.1/tenant/([0-9])+/radacat/versions', 'GET', 'domain');
INSERT INTO `tbl_resource` VALUES ('40', '删除版本', '/v1.1/tenant/([0-9])+/radacat/versions/([0-9])+', 'DELETE', 'domain');
INSERT INTO `tbl_resource` VALUES ('41', '添加工单', '/v1.1/tenant/([0-9])+/feedbacks', 'POST', 'domain');
INSERT INTO `tbl_resource` VALUES ('42', '基于租户获取工单', '/v1.1/tenant/([0-9])+/feedbacks', 'GET', 'domain');
INSERT INTO `tbl_resource` VALUES ('43', '基于租户和工单id获取工单', '/v1.1/tenant/([0-9])+/feedbacks/([0-9])+', 'GET', 'domain');
INSERT INTO `tbl_resource` VALUES ('44', '查询app用户', '/v1.1/tenant/([0-9])+/accounts', 'GET', 'domain');
INSERT INTO `tbl_resource` VALUES ('45', '基于uid查询APP用户', '/v1.1/tenant/([0-9])+/accounts/([0-9])+', 'GET', 'domain');
INSERT INTO `tbl_resource` VALUES ('46', '基于用户名查询APP用户', '/v1.1/tenant/([0-9])+/accounts', 'GET', 'domain');
INSERT INTO `tbl_resource` VALUES ('47', '添加广告', '/v1.1/tenant/([0-9])+/advertisements', 'POST', 'domain');
INSERT INTO `tbl_resource` VALUES ('48', '修改广告', '/v1.1/tenant/([0-9])+/advertisements/([0-9])+', 'PUT', 'domain');
INSERT INTO `tbl_resource` VALUES ('49', '获取广告', '/v1.1/tenant/([0-9])+/advertisements', 'GET', 'domain');
INSERT INTO `tbl_resource` VALUES ('50', '删除广告', '/v1.1/tenant/([0-9])+/advertisements/([0-9])+', 'DELETE', 'domain');
INSERT INTO `tbl_resource` VALUES ('51', '修改用户状态', '/v1.1/accounts/([0-9])+/state', 'PUT', 'auth_user');
INSERT INTO `tbl_resource` VALUES ('52', '添加租户', '/v1.1/domains/([0-9])+/tenants', 'POST', 'auth_tenant');
INSERT INTO `tbl_resource` VALUES ('53', '添加子租户', '/v1.1/domains/([0-9])+/tenants/tid', 'POST', 'auth_tenant');
INSERT INTO `tbl_resource` VALUES ('54', '修改租户信息', '/v1.1/domains/([0-9])+/tenants/([0-9])+', 'PUT', 'auth_tenant');
INSERT INTO `tbl_resource` VALUES ('55', '修改租户状态', '/v1.1/domains/([0-9])+/tenants/([0-9])+/states', 'PUT', 'auth_tenant');
INSERT INTO `tbl_resource` VALUES ('56', '获取域下租户列表', '/v1.1/domains/([0-9])+/tenants', 'GET', 'auth_tenant');
INSERT INTO `tbl_resource` VALUES ('57', '删除指定租户', '/v1.1/domains/([0-9])+/tenants/([0-9])+', 'DELETE', 'auth_tenant');
INSERT INTO `tbl_resource` VALUES ('58', '获取租户权限模块列表', '/v1.1/tenants/([0-9])+/modules', 'GET', 'auth_tenant_role');
INSERT INTO `tbl_resource` VALUES ('59', '添加租户角色', '/v1.1/tenants/([0-9])+/roles', 'POST', 'auth_tenant_role');
INSERT INTO `tbl_resource` VALUES ('60', '修改租户角色', '/v1.1/tenants/([0-9])+/roles', 'PUT', 'auth_tenant_role');
INSERT INTO `tbl_resource` VALUES ('61', '获取租户角色列表', '/v1.1/tenants/([0-9])+/roles', 'GET', 'auth_tenant_role');
INSERT INTO `tbl_resource` VALUES ('62', '删除租户角色', '/v1.1/tenants/([0-9])+/roles/([0-9])+', 'DELETE', 'auth_tenant_role');
INSERT INTO `tbl_resource` VALUES ('63', '通过rid查询租户角色', '/v1.1/tenants/([0-9])+/roles/([0-9])+', 'GET', 'auth_tenant_role');
INSERT INTO `tbl_resource` VALUES ('64', '获取域权限模块列表', '/v1.1/domains/([0-9])+/modules', 'GET', 'auth_domain_role');
INSERT INTO `tbl_resource` VALUES ('65', '添加域角色', '/v1.1/domains/([0-9])+/roles', 'POST', 'auth_domain_role');
INSERT INTO `tbl_resource` VALUES ('66', '修改域角色信息', '/v1.1/domains/([0-9])+/roles', 'PUT', 'auth_domain_role');
INSERT INTO `tbl_resource` VALUES ('67', '获取域角色列表', '/v1.1/domains/([0-9])+/roles', 'GET', 'auth_domain_role');
INSERT INTO `tbl_resource` VALUES ('68', '删除域角色', '/v1.1/domains/([0-9])+/roles/([0-9])+', 'DELETE', 'auth_domain_role');
INSERT INTO `tbl_resource` VALUES ('69', '通过rid获取域角色信息', '/v1.1/domains/([0-9])+/roles/([0-9])+', 'GET', 'auth_domain_role');
INSERT INTO `tbl_resource` VALUES ('70', '邀请用户进入租户', '/v1.1/tenants/([0-9])+/users', 'POST', 'auth_tenant_user');
INSERT INTO `tbl_resource` VALUES ('71', '修改租户用户角色', '/v1.1/tenants/([0-9])+/users/([0-9])+/roles', 'PUT', 'auth_tenant_user');
INSERT INTO `tbl_resource` VALUES ('72', '修改租户用户状态', '/v1.1/tenants/([0-9])+/users/([0-9])+/states', 'PUT', 'auth_tenant_user');
INSERT INTO `tbl_resource` VALUES ('73', '获取租户用户列表', '/v1.1/tenants/([0-9])+/users', 'GET', 'auth_tenant_user');
INSERT INTO `tbl_resource` VALUES ('74', '删除租户用户', '/v1.1/tenants/([0-9])+/users/([0-9])+', 'DELETE', 'auth_tenant_user');
INSERT INTO `tbl_resource` VALUES ('75', '邀请用户进入域', '/v1.1/domains/([0-9])+/users', 'POST', 'auth_domain_user');
INSERT INTO `tbl_resource` VALUES ('76', '修改域用户角色', '/v1.1/domains/([0-9])+/users/([0-9])+/roles', 'PUT', 'auth_domain_user');
INSERT INTO `tbl_resource` VALUES ('77', '修改域用户状态', '/v1.1/domains/([0-9])+/users/([0-9])+/states', 'PUT', 'auth_domain_user');
INSERT INTO `tbl_resource` VALUES ('78', '获取域用户列表', '/v1.1/domains/([0-9])+/users', 'GET', 'auth_domain_user');
INSERT INTO `tbl_resource` VALUES ('79', '删除域用户', '/v1.1/domains/([0-9])+/users/([0-9])+', 'DELETE', 'auth_domain_user');
INSERT INTO `tbl_resource` VALUES ('80', '添加服务', '/v1.1/tenants/([0-9])+/services', 'POST', 'auth_service');
INSERT INTO `tbl_resource` VALUES ('81', '删除服务', '/v1.1/tenants/([0-9])+/services/([0-9])+', 'DELETE', 'auth_service');
INSERT INTO `tbl_resource` VALUES ('82', '修改服务', '/v1.1/tenants/([0-9])+/services/([0-9])+', 'PUT', 'auth_service');
INSERT INTO `tbl_resource` VALUES ('83', '通过Sid获取服务', '/v1.1/tenants/([0-9])+/services/([0-9])+', 'GET', 'auth_service');
INSERT INTO `tbl_resource` VALUES ('84', '通过Tid获取服务', '/v1.1/tenants/([0-9])+/services', 'GET', 'auth_service');
INSERT INTO `tbl_resource` VALUES ('85', '添加计费策略', '/v1.1/domains/([0-9])+/services/([0-9])+/policy', 'POST', 'auth_policy');
INSERT INTO `tbl_resource` VALUES ('86', '通过pid删除计费策略', '/v1.1/domains/([0-9])+/services/([0-9])+/policy/([0-9])+', 'DELETE', 'auth_policy');
INSERT INTO `tbl_resource` VALUES ('87', '通过sid删除计费策略', '/v1.1/domains/([0-9])+/services/([0-9])+/policy', 'DELETE', 'auth_policy');
INSERT INTO `tbl_resource` VALUES ('88', '修改计费策略', '/v1.1/domains/([0-9])+/services/([0-9])+/policy/([0-9])+', 'PUT', 'auth_policy');
INSERT INTO `tbl_resource` VALUES ('89', '通过Pid获取计费策略', '/v1.1/domains/([0-9])+/services/([0-9])+/policy/([0-9])+', 'GET', 'auth_policy');
INSERT INTO `tbl_resource` VALUES ('90', '通过Sid获取计费策略', '/v1.1/domains/([0-9])+/services/([0-9])+/policy', 'GET', 'auth_policy');
INSERT INTO `tbl_resource` VALUES ('91', '添加支付宝信息', '/v1.1/domains/([0-9])+/alipay', 'POST', 'auth_payment');
INSERT INTO `tbl_resource` VALUES ('92', '删除支付宝信息', '/v1.1/domains/([0-9])+/alipay', 'DELETE', 'auth_payment');
INSERT INTO `tbl_resource` VALUES ('93', '修改支付宝信息', '/v1.1/domains/([0-9])+/alipay', 'PUT', 'auth_payment');
INSERT INTO `tbl_resource` VALUES ('94', '查询支付宝信息', '/v1.1/domains/([0-9])+/alipay', 'GET', 'auth_payment');
INSERT INTO `tbl_resource` VALUES ('95', '添加微信支付信息', '/v1.1/domains/([0-9])+/wechatpay', 'POST', 'auth_payment');
INSERT INTO `tbl_resource` VALUES ('96', '删除微信支付信息', '/v1.1/domains/([0-9])+/wechatpay', 'DELETE', 'auth_payment');
INSERT INTO `tbl_resource` VALUES ('97', '修改微信支付信息', '/v1.1/domains/([0-9])+/wechatpay', 'PUT', 'auth_payment');
INSERT INTO `tbl_resource` VALUES ('98', '查询微信支付信息', '/v1.1/domains/([0-9])+/wechatpay', 'GET', 'auth_payment');
INSERT INTO `tbl_resource` VALUES ('99', '	查询操作日志', '/v1.1/actionlogs', 'GET', 'auth_actionlog');

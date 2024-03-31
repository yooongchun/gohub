# ************************************************************
# Antares - SQL Client
# Version 0.7.23-beta.0
# 
# https://antares-sql.app/
# https://github.com/antares-sql/antares
# 
# Host: 127.0.0.1 (MySQL Community Server - GPL 8.0.33)
# Database: gohub
# Generation time: 2024-03-30T21:12:24+08:00
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table sys_login_log
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_login_log`;

CREATE TABLE `sys_login_log` (
  `info_id` bigint NOT NULL AUTO_INCREMENT COMMENT '访问ID',
  `login_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '登录账号',
  `ipaddr` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '登录IP地址',
  `login_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '登录地点',
  `browser` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '浏览器类型',
  `os` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '操作系统',
  `status` tinyint DEFAULT '0' COMMENT '登录状态（0成功 1失败）',
  `msg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '提示消息',
  `login_time` datetime DEFAULT NULL COMMENT '登录时间',
  `module` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '登录模块',
  PRIMARY KEY (`info_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=COMPACT COMMENT='系统访问记录';

LOCK TABLES `sys_login_log` WRITE;
/*!40000 ALTER TABLE `sys_login_log` DISABLE KEYS */;

INSERT INTO `sys_login_log` (`info_id`, `login_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`, `module`) VALUES
	(1, "demo", "::1", "内网IP", "Chrome", "Windows 10", 1, "登录成功", "2023-01-19 10:17:18", "系统后台"),
	(2, "demo", "::1", "内网IP", "Chrome", "Windows 10", 1, "登录成功", "2024-03-29 23:43:36", "系统后台");

/*!40000 ALTER TABLE `sys_login_log` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_user`;

CREATE TABLE `sys_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
  `user_nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
  `user_password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录密码;cmf_password加密',
  `user_salt` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '加密盐',
  `user_status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '用户状态;0:禁用,1:正常,2:未验证',
  `user_email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户登录邮箱',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户头像',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `is_admin` tinyint NOT NULL DEFAULT '1' COMMENT '是否后台管理员 1 是  0   否',
  `describe` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '描述信息',
  `last_login_ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '最后登录ip',
  `last_login_time` datetime DEFAULT NULL COMMENT '最后登录时间',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `user_login` (`user_name`,`deleted_at`) USING BTREE,
  UNIQUE KEY `mobile` (`mobile`,`deleted_at`) USING BTREE,
  KEY `user_nickname` (`user_nickname`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=COMPACT COMMENT='用户表';

LOCK TABLES `sys_user` WRITE;
/*!40000 ALTER TABLE `sys_user` DISABLE KEYS */;

INSERT INTO `sys_user` (`id`, `user_name`, `mobile`, `user_nickname`, `user_password`, `user_salt`, `user_status`, `user_email`, `avatar`, `remark`, `is_admin`, `describe`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, "admin", "13578342363", "超级管理员", "c567ae329f9929b518759d3bea13f492", "f9aZTAa8yz", 1, "yxh669@qq.com", "https://yxh-1301841944.cos.ap-chongqing.myqcloud.com/gfast/2021-07-19/ccwpeuqz1i2s769hua.jpeg", "", 1, "asdasfdsaf大发放打发士大夫发按时", "::1", "2022-10-26 03:01:52", "2021-06-22 17:58:00", "2022-11-03 15:44:38", NULL),
	(2, "lisi", "13699885599", "李四-管理员", "542a6e44dbac171f260fc4a032cd5522", "dlqVVBTADg", 1, "yxh@qq.com", "upload_file/2022-11-04/co3e5ljknns8jhlp8s.jpg", "备注", 1, "", "::1", "2022-11-04 09:54:56", "2021-06-22 17:58:00", "2022-11-04 17:54:56", NULL),
	(3, "zs", "16399669855", "张三", "41e3778c20338f4d7d6cc886fd3b2a52", "redoHIj524", 1, "zs@qq.com", "https://yxh-1301841944.cos.ap-chongqing.myqcloud.com/gfast/2021-08-02/cd8nif79egjg9kbkgk.jpeg", "", 0, "", "::1", "2022-04-28 10:01:47", "2021-06-22 17:58:00", "2022-04-28 10:01:47", NULL),
	(4, "wangwu", "13758596696", "王五", "542a6e44dbac171f260fc4a032cd5522", "dlqVVBTADg", 1, "qlgl@qq.com", "", "", 0, "", "127.0.0.1", NULL, "2021-06-22 17:58:00", "2022-11-03 15:44:20", NULL),
	(5, "maliu", "13845696696", "马六", "542a6e44dbac171f260fc4a032cd5522", "dlqVVBTADg", 1, "123@qq.com", "", "", 0, "", "::1", "2022-03-30 10:50:39", "2021-06-22 17:58:00", "2022-11-03 15:44:10", NULL);
/*!40000 ALTER TABLE `sys_user` ENABLE KEYS */;
UNLOCK TABLES;


DROP TABLE IF EXISTS `sys_operate_log`;

CREATE TABLE `sys_operate_log` (
    `operate_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '日志主键',
    `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '模块标题',
    `business_type` int DEFAULT '0' COMMENT '业务类型（0其它 1新增 2修改 3删除）',
    `method` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '方法名称',
    `request_method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '请求方式',
    `operator_type` int DEFAULT '0' COMMENT '操作类别（0其它 1后台用户 2手机端用户）',
    `operate_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '操作人员',
    `operate_url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '请求URL',
    `operate_ip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '主机地址',
    `operate_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '操作地点',
    `operate_param` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '请求参数',
    `error_msg` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '错误消息',
    `operate_time` datetime DEFAULT NULL COMMENT '操作时间',
    PRIMARY KEY (`operate_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=57 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=COMPACT COMMENT='操作日志记录';

LOCK TABLES `sys_operate_log` WRITE;
/*!40000 ALTER TABLE `sys_operate_log` DISABLE KEYS */;

INSERT INTO `sys_operate_log` (`operate_id`, `title`, `business_type`, `method`, `request_method`, `operator_type`, `operate_name`, `operate_url`, `operate_ip`, `operate_location`, `operate_param`, `error_msg`, `operate_time`) VALUES
    (1, "", 0, "/api/v1/system/dict/data/getDictData", "GET", 1, "demo", "/api/v1/system/dict/data/getDictData?dictType=sys_oper_log_type&defaultValue=", "::1", "内网IP", "{\"defaultValue\":\"\",\"dictType\":\"sys_oper_log_type\"}", "", "2023-01-19 10:10:49"),
    (2, "操作日志", 0, "/api/v1/system/operLog/list", "GET", 1, "demo", "/api/v1/system/operLog/list?pageNum=1&pageSize=10", "::1", "内网IP", "{\"pageNum\":\"1\",\"pageSize\":\"10\"}", "", "2023-01-19 10:10:49"),
    (3, "操作日志", 0, "/api/v1/system/operLog/list", "GET", 1, "demo", "/api/v1/system/operLog/list?pageNum=1&pageSize=10", "::1", "内网IP", "{\"pageNum\":\"1\",\"pageSize\":\"10\"}", "", "2023-01-19 10:11:04"),
    (4, "在线用户", 0, "/api/v1/system/online/list", "GET", 1, "demo", "/api/v1/system/online/list?ipaddr=&userName=&pageNum=1&pageSize=10", "::1", "内网IP", "{\"ipaddr\":\"\",\"pageNum\":\"1\",\"pageSize\":\"10\",\"userName\":\"\"}", "", "2023-01-19 10:16:55"),
    (5, "", 0, "/api/v1/system/dict/data/getDictData", "GET", 1, "demo", "/api/v1/system/dict/data/getDictData?dictType=sys_oper_log_type&defaultValue=", "::1", "内网IP", "{\"defaultValue\":\"\",\"dictType\":\"sys_oper_log_type\"}", "", "2023-01-19 10:16:57");

/*!40000 ALTER TABLE `sys_operate_log` ENABLE KEYS */;
UNLOCK TABLES;


/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

# Dump completed on 2024-03-30T21:12:24+08:00

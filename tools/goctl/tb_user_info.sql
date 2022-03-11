/*
 Navicat Premium Data Transfer

 Source Server         : LZ_TEST
 Source Server Type    : MySQL
 Source Server Version : 50730
 Source Host           : 39.108.112.202:3306
 Source Schema         : st_agent

 Target Server Type    : MySQL
 Target Server Version : 50730
 File Encoding         : 65001

 Date: 11/03/2022 15:00:47
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tb_user_info
-- ----------------------------
DROP TABLE IF EXISTS `tb_user_info`;
CREATE TABLE `tb_user_info` (
  `id` varchar(32) NOT NULL COMMENT '主键',
  `nickName` varchar(64) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT 'nickName',
  `sex` int(2) NOT NULL DEFAULT '0' COMMENT '性别:0-未知 1-男性 2-女性',
  `country` varchar(64) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '国家',
  `province` varchar(64) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '省份',
  `city` varchar(64) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '城市',
  `avatar` varchar(1024) DEFAULT NULL COMMENT '头像',
  `createBy` varchar(32) NOT NULL COMMENT '创建人',
  `createTime` datetime NOT NULL COMMENT '创建时间',
  `updateBy` varchar(32) DEFAULT NULL COMMENT '修改人',
  `updateTime` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `tb_user_info_nickName_index` (`nickName`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;

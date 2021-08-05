/*
 Navicat Premium Data Transfer

 Source Server         : todo
 Source Server Type    : MySQL
 Source Server Version : 80018
 Source Host           : rm-bp1lmwte7c7oqb1b7so.mysql.rds.aliyuncs.com:3306
 Source Schema         : todo

 Target Server Type    : MySQL
 Target Server Version : 80018
 File Encoding         : 65001

 Date: 02/08/2021 22:58:03
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tasks
-- ----------------------------
DROP TABLE IF EXISTS `tasks`;
CREATE TABLE `tasks`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `text` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `ddl` timestamp(0) NULL DEFAULT NULL,
  `finished` tinyint(4) NOT NULL DEFAULT 0,
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `updated_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 75 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;

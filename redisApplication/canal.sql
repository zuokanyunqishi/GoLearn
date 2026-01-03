/*
 Navicat Premium Data Transfer

 Source Server         : 神舟mysql
 Source Server Type    : MySQL
 Source Server Version : 80033 (8.0.33)
 Source Host           : 192.168.0.103:3306
 Source Schema         : canal_redis

 Target Server Type    : MySQL
 Target Server Version : 80033 (8.0.33)
 File Encoding         : 65001

 Date: 28/05/2023 20:15:10
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;


drop database if exists canal_redis;
create database canal_redis;
-- ----------------------------
-- Table structure for baseCreditInfo
-- ----------------------------
DROP TABLE IF EXISTS `baseCreditInfo`;
CREATE TABLE `baseCreditInfo`
(
    `id`         int UNSIGNED                                                  NOT NULL AUTO_INCREMENT,
    `bankName`   varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '银行名字',
    `billDate`   tinyint                                                       NOT NULL DEFAULT 0 COMMENT '账单日',
    `dueDate`    tinyint                                                       NOT NULL DEFAULT 0 COMMENT '还款日',
    `cardNo`     varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL DEFAULT '' COMMENT '卡号',
    `score`      int                                                           NOT NULL DEFAULT 0 COMMENT '积分',
    `cardType`   tinyint                                                       NULL     DEFAULT 0 COMMENT '1:光大2:招商3:民生4:中信',
    `createTime` timestamp                                                     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updateTIme` timestamp                                                     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 7
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT = '银行卡信息'
  ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;

INSERT INTO `canal_redis`.`baseCreditInfo` (`id`, `bankName`, `billDate`, `dueDate`, `cardNo`, `score`, `cardType`, `createTime`, `updateTIme`) VALUES (3, '民生银行', 25, 0, '622607777749669', 3, 3, '2019-04-08 22:03:10', '2023-05-28 19:45:35');
INSERT INTO `canal_redis`.`baseCreditInfo` (`id`, `bankName`, `billDate`, `dueDate`, `cardNo`, `score`, `cardType`, `createTime`, `updateTIme`) VALUES (4, '中信银行', 19, 0, '6228888589415', 4, 4, '2019-04-08 22:04:15', '2023-05-28 19:45:43');
INSERT INTO `canal_redis`.`baseCreditInfo` (`id`, `bankName`, `billDate`, `dueDate`, `cardNo`, `score`, `cardType`, `createTime`, `updateTIme`) VALUES (5, '浦发银行', 14, 0, '7888888888888', 5, 5, '2023-05-28 19:14:39', '2023-05-28 19:45:50');
INSERT INTO `canal_redis`.`baseCreditInfo` (`id`, `bankName`, `billDate`, `dueDate`, `cardNo`, `score`, `cardType`, `createTime`, `updateTIme`) VALUES (6, '中原银行', 13, 0, '12777777777777', 6, 6, '2023-05-28 19:50:38', '2023-05-28 19:50:52');


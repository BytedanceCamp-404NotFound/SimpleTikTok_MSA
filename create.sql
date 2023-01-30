/*
 Navicat Premium Data Transfer
 Source Server         : localhost-Mysql5.7
 Source Server Type    : MySQL
 Source Server Version : 50738
 Source Host           : localhost:3306
 Source Schema         : SimpleTikTok
 Target Server Type    : MySQL
 Target Server Version : 50738
 File Encoding         : 65001
 Date: 19/01/2023 10:18:48
*/
CREATE DATABASE IF NOT EXISTS SimpleTikTok DEFAULT CHARACTER SET utf8mb4 COLLATE  utf8mb4_general_ci;

use SimpleTikTok;
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for favorite_list
-- ----------------------------
DROP TABLE IF EXISTS `favorite_list`;
CREATE TABLE `favorite_list` (
  `record_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '记录编号,自增',
  `favorite_video_id` int(11) DEFAULT NULL COMMENT '被点赞视频id',
  `favorite_user_id` int(11) DEFAULT NULL COMMENT '点赞账号id',
  `record_time` datetime DEFAULT NULL COMMENT '记录时间',
  PRIMARY KEY (`record_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for follow_and_follower_list
-- ----------------------------
DROP TABLE IF EXISTS `follow_and_follower_list`;
CREATE TABLE `follow_and_follower_list` (
  `record_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '记录编号，自增',
  `user_id` int(11) NOT NULL COMMENT '被关注的账号id',
  `follower_id` int(11) NOT NULL COMMENT '粉丝账号id',
  `record_time` datetime DEFAULT NULL COMMENT '关注时间',
  PRIMARY KEY (`record_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for user_info
-- ----------------------------
DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info` (
  `user_id` int(11) NOT NULL COMMENT '账号id',
  `user_nick_name` varchar(255) NOT NULL COMMENT '用户昵称',
  `follow_count` int(11) DEFAULT '0' COMMENT '关注量',
  `follower_count` int(11) DEFAULT '0' COMMENT '粉丝量',
  `update_time` datetime DEFAULT NULL COMMENT '资料更新时间',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for user_login
-- ----------------------------
DROP TABLE IF EXISTS `user_login`;
CREATE TABLE `user_login` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户ID,自增',
  `user_name` varchar(255) NOT NULL COMMENT '用户账户',
  `user_pwd` varchar(255) NOT NULL COMMENT '用户密码',
  `register_date` datetime DEFAULT NULL COMMENT '注册时间',
  PRIMARY KEY (`user_id`,`user_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for video_info
-- ----------------------------
DROP TABLE IF EXISTS `video_info`;
CREATE TABLE `video_info` (
  `video_id` int(11) NOT NULL COMMENT '视频唯一标识',
  `video_title` varchar(255) DEFAULT NULL COMMENT '视频标题',
  `author_id` int(255) NOT NULL COMMENT '视频作者账号id',
  `cover_url` varchar(255) DEFAULT NULL COMMENT '视频封面地址',
  `play_url` varchar(255) DEFAULT NULL COMMENT '视频播放地址',
  `favorite_count` int(11) DEFAULT '0' COMMENT '视频的点赞总数',
  `comment_count` int(11) DEFAULT '0' COMMENT '视频的评论总数',
  `update_time` datetime DEFAULT NULL COMMENT '上传时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`video_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
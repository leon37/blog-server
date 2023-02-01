-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        10.11.1-MariaDB - mariadb.org binary distribution
-- 服务器操作系统:                      Win64
-- HeidiSQL 版本:                  11.3.0.6295
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT = @@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS = 0 */;
/*!40101 SET @OLD_SQL_MODE = @@SQL_MODE, SQL_MODE = 'NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES = @@SQL_NOTES, SQL_NOTES = 0 */;


-- 导出 blog_server 的数据库结构
CREATE DATABASE IF NOT EXISTS `blog_server` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */;
USE `blog_server`;

-- 导出  表 blog_server.blog_article 结构
CREATE TABLE IF NOT EXISTS `blog_article`
(
    `id`              int(10) unsigned NOT NULL AUTO_INCREMENT,
    `title`           varchar(100)        DEFAULT '' COMMENT '文章标题',
    `desc`            varchar(255)        DEFAULT '' COMMENT '文章简述',
    `cover_image_url` varchar(255)        DEFAULT '' COMMENT '封面图片地址',
    `content`         longtext            DEFAULT NULL COMMENT '文章内容',
    `created_on`      int(10) unsigned    DEFAULT 0 COMMENT '新建时间',
    `created_by`      varchar(100)        DEFAULT '' COMMENT '创建人',
    `modified_on`     int(10) unsigned    DEFAULT 0 COMMENT '修改时间',
    `modified_by`     varchar(100)        DEFAULT '' COMMENT '修改人',
    `deleted_on`      int(10) unsigned    DEFAULT 0 COMMENT '删除时间',
    `is_del`          tinyint(3) unsigned DEFAULT 0 COMMENT '是否删除： 0为未删除、1为已删除',
    `state`           tinyint(3) unsigned DEFAULT 1 COMMENT '状态：0为禁用、1为启用',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 3
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT ='文章管理';

-- 数据导出被取消选择。

-- 导出  表 blog_server.blog_article_tag 结构
CREATE TABLE IF NOT EXISTS `blog_article_tag`
(
    `id`          int(10) unsigned NOT NULL AUTO_INCREMENT,
    `article_id`  int(11)          NOT NULL COMMENT '文章ID',
    `tag_id`      int(10) unsigned NOT NULL DEFAULT 0 COMMENT '标签ID',
    `created_on`  int(10) unsigned          DEFAULT 0 COMMENT '创建时间',
    `created_by`  varchar(100)              DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned          DEFAULT 0 COMMENT '修改时间',
    `modified_by` varchar(100)              DEFAULT '' COMMENT '修改人',
    `deleted_on`  int(10) unsigned          DEFAULT 0 COMMENT '删除时间',
    `is_del`      tinyint(3) unsigned       DEFAULT 0 COMMENT '是否删除： 0为未删除、1为已删除',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT ='文章标签关联';

-- 数据导出被取消选择。

-- 导出  表 blog_server.blog_tag 结构
CREATE TABLE IF NOT EXISTS `blog_tag`
(
    `id`          int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name`        varchar(100)        DEFAULT '' COMMENT '标签名称',
    `created_on`  int(10) unsigned    DEFAULT 0 COMMENT '创建时间',
    `created_by`  varchar(100)        DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned    DEFAULT 0 COMMENT '修改时间',
    `modified_by` varchar(100)        DEFAULT '' COMMENT '修改人',
    `deleted_on`  int(10) unsigned    DEFAULT 0 COMMENT '删除时间',
    `is_del`      tinyint(3) unsigned DEFAULT 0 COMMENT '是否删除 0为未删除、1为已删除',
    `state`       tinyint(3) unsigned DEFAULT 1 COMMENT '状态 0为禁用、1为启用',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 9
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT ='标签管理';

-- 数据导出被取消选择。

/*!40101 SET SQL_MODE = IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS = IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT = @OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES = IFNULL(@OLD_SQL_NOTES, 1) */;


CREATE TABLE `blog_auth`
(
    `id`          int(10) unsigned NOT NULL AUTO_INCREMENT,
    `app_key`     varchar(20)         DEFAULT '' COMMENT 'Key',
    `app_secret`  varchar(50)         DEFAULT '' COMMENT 'Secret',
    `created_on`  int(10) unsigned    DEFAULT 0 COMMENT '创建时间',
    `created_by`  varchar(100)        DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned    DEFAULT 0 COMMENT '修改时间',
    `modified_by` varchar(100)        DEFAULT '' COMMENT '修改人',
    `deleted_on`  int(10) unsigned    DEFAULT 0 COMMENT '删除时间',
    `is_del`      tinyint(3) unsigned DEFAULT 0 COMMENT '是否删除 0为未删除、1为已删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='认证管理'
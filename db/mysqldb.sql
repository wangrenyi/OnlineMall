### 创建 database `onlinemall`
CREATE DATABASE `onlinemall` DEFAULT CHARACTER SET utf8;

### 创建用户,权限 table
CREATE TABLE IF NOT EXISTS `online_mall_shops`
(
    `ID`               varchar(32)      NOT NULL,
    `userId`           int              NOT NULL,
    `shopsName`        varchar(255)     NOT NULL,
    `shopsRate`        tinyint(4),
    `businessCategory` varchar(255),
    `qualification`    varchar(255),
    `remarks`          varchar(255),
    `createUser`       varchar(50),
    `createTime`       datetime         NULL     DEFAULT CURRENT_TIMESTAMP,
    `updateUser`       varchar(50),
    `updateTime`       datetime         NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `version`          int(10) unsigned NOT NULL DEFAULT '1',
    PRIMARY KEY (`ID`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;


CREATE TABLE IF NOT EXISTS `mst_user_info`
(
    `id`          int(11) unsigned NOT NULL AUTO_INCREMENT,
    `loginName`   varchar(50)      NOT NULL,
    `password`    varchar(50)      NOT NULL,
    `displayName` varchar(50),
    `email`       varchar(50),
    `phoneNumber` varchar(50),
    `type`        varchar(10)      NOT NULL DEFAULT 'normal', ##普通用户或者管理员用户
    `enabled`     bit(1)           NOT NULL DEFAULT b'1',     ##标识用户是否有效
    `remarks`     varchar(255),
    `createUser`  varchar(50),
    `createTime`  datetime         NULL     DEFAULT CURRENT_TIMESTAMP,
    `updateUser`  varchar(50),
    `updateTime`  datetime         NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `version`     int(10) unsigned NOT NULL DEFAULT '1',
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_loginName` (`loginName`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

### 创建商品详情
CREATE TABLE IF NOT EXISTS `online_mall_goods`
(
    `ID`         varchar(32)      NOT NULL,
    `shopsId`    varchar(32)      NOT NULL,
    `goodsName`  varchar(255)     NOT NULL,
    `category`   varchar(255)     NOT NULL,
    `status`     bit              NOT NULL DEFAULT b'1',
    `remarks`    varchar(255),
    `createUser` varchar(50),
    `createTime` datetime         NULL     DEFAULT CURRENT_TIMESTAMP,
    `updateUser` varchar(50),
    `updateTime` datetime         NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `version`    int(10) unsigned NOT NULL DEFAULT '1',
    PRIMARY KEY (`ID`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `online_mall_goods_detail`
(
    `ID`            int(11) unsigned NOT NULL AUTO_INCREMENT,
    `goodsId`       varchar(32)      NOT NULL,
    `origin`        varchar(255),             ##产地
    `price`         int              NOT NULL,##价格,单位/分
    `freight`       int,                      ##运费,单位/分
    `originalPrice` int,                      ##原价,单位/分
    `specification` varchar(50)      NOT NULL,##规格,尺寸
    `remarks`       varchar(255),
    `createUser`    varchar(50),
    `createTime`    datetime         NULL     DEFAULT CURRENT_TIMESTAMP,
    `updateUser`    varchar(50),
    `updateTime`    datetime         NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `version`       int(10) unsigned NOT NULL DEFAULT '1',
    PRIMARY KEY (`ID`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

###  商品图,文件信息
CREATE TABLE IF NOT EXISTS `online_mall_file`
(
    `id`               varchar(32)      NOT NULL,
    `reqId`            varchar(255)     NOT NULL,
    `reqType`          varchar(255)     NOT NULL,
    `fileOriginalName` varchar(255)     NOT NULL,
    `fileUniqueName`   varchar(255)     NOT NULL,
    `subdirectory`     varchar(255)              DEFAULT NULL,
    `createUser`       varchar(50)               DEFAULT NULL,
    `createTime`       datetime                  DEFAULT NULL,
    `updateUser`       varchar(50)               DEFAULT NULL,
    `updateTime`       datetime                  DEFAULT NULL,
    `version`          int(10) unsigned NOT NULL DEFAULT '1',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;
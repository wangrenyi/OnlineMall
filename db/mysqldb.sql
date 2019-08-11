### 创建 database `onlinemall`
CREATE DATABASE `onlinemall` DEFAULT CHARACTER SET utf8;

### 创建用户,权限 table
CREATE TABLE IF NOT EXISTS `online_mall_shops`
(
    `id`                varchar(32)      NOT NULL,
    `user_id`           int              NOT NULL,
    `shops_name`        varchar(255)     NOT NULL,
    `shops_rate`        tinyint(4),
    `business_category` varchar(255),
    `qualification`     varchar(255),
    `status`            tinyint(4)       NOT NULL DEFAULT 1,
    `remarks`           varchar(255),
    `create_user`       varchar(50),
    `create_time`       datetime         NULL     DEFAULT CURRENT_TIMESTAMP,
    `update_user`       varchar(50),
    `update_time`       datetime         NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `version`           int(10) unsigned NOT NULL DEFAULT '1',
    PRIMARY KEY (`id`),
    KEY `index_user_id` (`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;


CREATE TABLE IF NOT EXISTS `mst_user_info`
(
    `id`           int(11) unsigned NOT NULL AUTO_INCREMENT,
    `login_name`   varchar(50)      NOT NULL,
    `password`     varchar(50)      NOT NULL,
    `display_name` varchar(50),
    `email`        varchar(50),
    `phone_number` varchar(50),
    `type`         varchar(10)      NOT NULL DEFAULT 'normal', ##普通用户或者管理员用户
    `enabled`      tinyint(4)       NOT NULL DEFAULT 1,        ##标识用户是否有效
    `remarks`      varchar(255),
    `create_user`  varchar(50),
    `create_time`  datetime         NULL     DEFAULT CURRENT_TIMESTAMP,
    `update_user`  varchar(50),
    `update_time`  datetime         NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `version`      int(10) unsigned NOT NULL DEFAULT '1',
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_loginName` (`login_name`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

### 创建商品详情
CREATE TABLE IF NOT EXISTS `online_mall_goods`
(
    `id`          varchar(32)      NOT NULL,
    `shops_id`    varchar(32)      NOT NULL,
    `goods_name`  varchar(255)     NOT NULL,
    `category`    varchar(255)     NOT NULL,
    `status`      tinyint(4)       NOT NULL DEFAULT 1,
    `remarks`     varchar(255),
    `create_user` varchar(50),
    `create_time` datetime         NULL     DEFAULT CURRENT_TIMESTAMP,
    `update_user` varchar(50),
    `update_time` datetime         NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `version`     int(10) unsigned NOT NULL DEFAULT '1',
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_shops_id` (`shops_id`)
) ENGINE = InnoDB  DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `online_mall_goods_detail`
(
    `id`             int(11) unsigned NOT NULL AUTO_INCREMENT,
    `goods_id`       varchar(32)      NOT NULL,
    `origin`         varchar(255),             ##产地
    `price`          int              NOT NULL,##价格,单位/分
    `freight`        int,                      ##运费,单位/分
    `original_price` int,                      ##原价,单位/分
    `specification`  varchar(50)      NOT NULL,##规格,尺寸
    `remarks`        varchar(255),
    `create_user`    varchar(50),
    `create_time`    datetime         NULL     DEFAULT CURRENT_TIMESTAMP,
    `update_user`    varchar(50),
    `update_time`    datetime         NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `version`        int(10) unsigned NOT NULL DEFAULT '1',
    PRIMARY KEY (`id`),
    KEY `index_goods_id` (`goods_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

###  商品图,文件信息
CREATE TABLE IF NOT EXISTS `online_mall_file`
(
    `id`                 varchar(32)      NOT NULL,
    `req_id`             varchar(255)     NOT NULL,
    `req_type`           varchar(255)     NOT NULL,
    `file_original_name` varchar(255)     NOT NULL,
    `file_unique_name`   varchar(255)     NOT NULL,
    `subdirectory`       varchar(255)              DEFAULT NULL,
    `create_user`        varchar(50)               DEFAULT NULL,
    `create_time`        datetime                  DEFAULT NULL,
    `update_user`        varchar(50)               DEFAULT NULL,
    `update_time`        datetime                  DEFAULT NULL,
    `version`            int(10) unsigned NOT NULL DEFAULT '1',
    PRIMARY KEY (`id`),
    KEY `index_req_id` (`req_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

### 用户收件信息
CREATE TABLE IF NOT EXISTS `online_mall_address_info`
(
    `id`             int(11) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`        int(11)          NOT NULL,
    `consignee_id`   int(11)          NOT NULL COMMENT '收货人id',
    `consignee_name` varchar(255)     NOT NULL COMMENT '收货人name',
    `phone_number`   varchar(32)      NOT NULL comment '收货人手机号码',
    `address`        varchar(255)     NOT NULL COMMENT '收件地址',
    `address_detail` varchar(255) COMMENT '详细地址',
    `default`        tinyint(4)       NOT NULL DEFAULT 0 COMMENT '默认地址',
    `status`         tinyint(4)       NOT NULL DEFAULT 1,
    `create_user`    varchar(50)               DEFAULT NULL,
    `create_time`    datetime                  DEFAULT CURRENT_TIMESTAMP,
    `update_user`    varchar(50)               DEFAULT NULL,
    `update_time`    datetime                  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `version`        int(10) unsigned NOT NULL DEFAULT '1',
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_user_id` (`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;


### 购物车shopping cart
CREATE TABLE IF NOT EXISTS `online_mall_shopping_cart`
(
    `id`          varchar(32)      NOT NULL,
    `user_id`     int(11)          NOT NULL,
    `create_user` varchar(50)               DEFAULT NULL,
    `create_time` datetime                  DEFAULT CURRENT_TIMESTAMP,
    `update_user` varchar(50)               DEFAULT NULL,
    `update_time` datetime                  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `version`     int(10) unsigned NOT NULL DEFAULT '1',
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_shoping_cart_user_id` (`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

### 购物车商品关系表
CREATE TABLE IF NOT EXISTS `online_mall_shopping_cart_goods`
(
    `id`               varchar(32)      NOT NULL,
    `shopping_cart_id` varchar(32)      NOT NULL,
    `goods_id`         varchar(32)      NOT NULL,
    `goods_name`       varchar(255)     NOT NULL,
    `quantity`         int(11)          NOT NULL DEFAULT '1' COMMENT '商品数量',
    `status`           tinyint(4)       NOT NULL DEFAULT '1',
    `create_user`      varchar(50)               DEFAULT NULL,
    `create_time`      datetime                  DEFAULT CURRENT_TIMESTAMP,
    `update_user`      varchar(50)               DEFAULT NULL,
    `update_time`      datetime                  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `version`          int(10) unsigned NOT NULL DEFAULT '1',
    PRIMARY KEY (`id`),
    KEY `index_shopping_cart_id` (`shopping_cart_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

### 订单detail
CREATE TABLE `online_mall_shops_orders`
(
    `id`              varchar(32)      NOT NULL,
    `orders_number`   varchar(255)     NOT NULL COMMENT '订单编号',
    `user_id`         int(11)          NOT NULL COMMENT '用户id',
    `address_id`      int(11)          NOT NULL COMMENT '收获信息id',
    `goods_id`        varchar(32)      NOT NULL,
    `goods_name`      varchar(255)     NOT NULL,
    `quantity`        int(11)          NOT NULL DEFAULT 1 COMMENT '商品数量',
    `payWay`          varchar(255)     NOT NULL COMMENT '支付方式',
    `remarks`         varchar(255)              DEFAULT '备注',
    `status`          tinyint(4)       NOT NULL DEFAULT '1' COMMENT '订单状态',
    `create_user`     varchar(50)               DEFAULT NULL,
    `create_time`     datetime                  DEFAULT CURRENT_TIMESTAMP,
    `update_user`     varchar(50)               DEFAULT NULL,
    `update_time`     datetime                  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `version`         int(10) unsigned NOT NULL DEFAULT '1',
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_orders_number` (`orders_number`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;


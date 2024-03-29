* 建库语句

  ```sql
CREATE DATABASE `imaotai`  DEFAULT CHARACTER SET utf8mb4 ;
  ```

* 建表语句
  
``` sql
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态',
  `deleted` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `md5` varchar(100) DEFAULT '' COMMENT 'md5',
  `device_id` varchar(100) DEFAULT '' COMMENT '设备名称',
  `token` varchar(400) DEFAULT NULL,
  `lat` float DEFAULT NULL,
  `lng` float DEFAULT NULL,
  `city_name` varchar(100) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `source` tinyint(4) DEFAULT NULL,
  `user_name` varchar(20) DEFAULT NULL,
  `province_name` varchar(50) DEFAULT NULL,
  `district_name` varchar(50) DEFAULT NULL,
  `cookie` varchar(400) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4;
```

``` sql
CREATE TABLE `record` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `deleted` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `user_id` int(11) DEFAULT '0',
  `user_name` varchar(20) DEFAULT '',
  `item_id` int(11) DEFAULT '0',
  `item_name` varchar(100) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

```sql
CREATE TABLE `item` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态',
  `deleted` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `item_code` varchar(20) DEFAULT '' COMMENT '酒编码',
  `title` varchar(20) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;
```

```sql
CREATE TABLE `shop` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态',
  `deleted` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `address` varchar(100) DEFAULT NULL,
  `city_name` varchar(20) DEFAULT NULL,
  `district_name` varchar(20) DEFAULT NULL,
  `full_address` varchar(100) DEFAULT NULL,
  `lng` float DEFAULT '0',
  `name` varchar(100) DEFAULT NULL,
  `province_name` varchar(20) DEFAULT NULL,
  `shop_id` varchar(20) DEFAULT NULL,
  `lat` float DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2861 DEFAULT CHARSET=utf8mb4
```

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `device`;
CREATE TABLE `device` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `device_name` varchar(50) DEFAULT '',
  `os` varchar(100) DEFAULT '',
  `machine_id` varchar(100) DEFAULT '',
  `created_on` int(10) NULL DEFAULT NULL,
  `modified_on` int(10) NULL DEFAULT NULL,
  `deleted_on` int(10) NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
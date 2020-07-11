-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '',
  `password` varchar(100) DEFAULT '',
  `is_admin` tinyint(1) NOT NULL DEFAULT 0,
  `created_on` int(10) NULL DEFAULT NULL,
  `modified_on` int(10) NULL DEFAULT NULL,
  `deleted_on` int(10) NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `user` (`id`, `username`, `password`, `is_admin`, `created_on`, `modified_on`, `deleted_on`) VALUES ('1', 'admin', '$2a$10$xIG496.l6ZfJQBOixKmxnO.Nx/Z1OsJYx2KQ1wPwnMp', 1, '1594460461', '1594460461', '0');
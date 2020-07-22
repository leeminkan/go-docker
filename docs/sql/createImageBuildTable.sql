-- ----------------------------
-- Table structure for image_build
-- ----------------------------
DROP TABLE IF EXISTS `image_build`;
CREATE TABLE `image_build` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `repo_name` varchar(50) UNIQUE DEFAULT '',
  `image_id` varchar(100) DEFAULT '',
  `user_id` int(10) unsigned NOT NULL,
  `status` enum('on progress','done', 'fail') NOT NULL DEFAULT 'on progress',
  `old_repo_name` varchar(50) DEFAULT '',
  `created_on` int(10) NULL DEFAULT NULL,
  `modified_on` int(10) NULL DEFAULT NULL,
  `deleted_on` int(10) NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
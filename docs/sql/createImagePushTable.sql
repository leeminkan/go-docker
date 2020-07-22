-- ----------------------------
-- Table structure for image_push
-- ----------------------------
DROP TABLE IF EXISTS `image_push`;
CREATE TABLE `image_push` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `repo_name` varchar(50) UNIQUE DEFAULT '',
  `user_id` int(10) unsigned NOT NULL,
  `digest` varchar(100) DEFAULT '',
  `status` enum('on progress','done', 'fail') NOT NULL DEFAULT 'on progress',
  `old_repo_name` varchar(50) DEFAULT '',
  `created_on` int(10) NULL DEFAULT NULL,
  `modified_on` int(10) NULL DEFAULT NULL,
  `deleted_on` int(10) NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
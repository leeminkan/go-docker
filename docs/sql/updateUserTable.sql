ALTER TABLE `user` ADD `x_registry_auth` varchar(500) DEFAULT NULL;
ALTER TABLE `user` ADD `is_login_docker_hub` tinyint(1) NOT NULL DEFAULT 0;
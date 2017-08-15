CREATE TABLE `short` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `origin_url` varchar(512) DEFAULT NULL,
  `last_modify_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
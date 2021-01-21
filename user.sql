use test;
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
                        `id` bigint(20) NOT NULL AUTO_INCREMENT,
                        `user_name` varchar(20) NOT NULL DEFAULT '',
                        `password` varchar(20) NOT NULL DEFAULT '0',
                        PRIMARY KEY (`id`),
                        KEY `id` (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;
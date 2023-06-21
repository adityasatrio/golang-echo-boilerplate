CREATE TABLE `miraj` (
                         `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                         `owner` varchar (255) NOT NULL,
                         `balance` bigint NOT NULL,
                         `currency` varchar (255) NOT NULL,
                         `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         PRIMARY KEY (`id`)
);
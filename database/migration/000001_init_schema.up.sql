CREATE TABLE `accounts` (
                            `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                            `owner` varchar (255) NOT NULL,
                            `balance` bigint NOT NULL,
                            `currency` varchar (255) NOT NULL,
                            `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            PRIMARY KEY (`id`)
);

CREATE TABLE `entries` (
                           `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                           `account_id` bigint NOT NULL,
                           `amount` bigint NOT NULL,
                           `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                           PRIMARY KEY (`id`)
);

CREATE TABLE `transfers` (
                             `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                             `from_account_id` bigint NOT NULL,
                             `to_account_id` bigint NOT NULL,
                             `amount` bigint NOT NULL,
                             `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                             PRIMARY KEY (`id`)
);
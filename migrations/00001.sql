CREATE TABLE `shorten`
(
    `id`         int unsigned                                                 NOT NULL AUTO_INCREMENT,
    `data`       text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci        NOT NULL,
    `hash`       varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `lock`       tinyint unsigned                                             NOT NULL,
    `created_at` datetime                                                     NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `hash` (`hash`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;
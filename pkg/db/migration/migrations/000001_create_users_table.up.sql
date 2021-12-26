CREATE TABLE IF NOT EXISTS users (
    id int(20) NOT NULL AUTO_INCREMENT,
    name varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    created_at timestamp NOT NULL default current_timestamp(),
    updated_at timestamp NOT NULL default current_timestamp() ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY name_index (name)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
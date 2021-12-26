CREATE TABLE IF NOT EXISTS todos (
    id int(20) NOT NULL AUTO_INCREMENT,
    todo varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    finished bool NOT NULL default false,
    user_id int(10) NOT NULL,
    created_at timestamp NOT NULL default current_timestamp(),
    updated_at timestamp NOT NULL default current_timestamp() ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY todos_user_id_foreign (user_id),
    CONSTRAINT `todos_user_id_foreign` FOREIGN KEY (user_id) REFERENCES users (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
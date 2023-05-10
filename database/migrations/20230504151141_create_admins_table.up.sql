CREATE TABLE admins (
    id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    created_by INT(10) UNSIGNED NULL DEFAULT NULL,
    updated_by INT(10) UNSIGNED NULL DEFAULT NULL,

    PRIMARY KEY (id),
    UNIQUE INDEX admins_email_unique (email),

    CONSTRAINT admins_created_by_foreign FOREIGN KEY (created_by) REFERENCES admins (id) ON DELETE SET NULL,
    CONSTRAINT admins_updated_by_foreign FOREIGN KEY (updated_by) REFERENCES admins (id) ON DELETE SET NULL

) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
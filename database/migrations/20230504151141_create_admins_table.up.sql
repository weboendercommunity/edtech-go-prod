CREATE TABLE admins (
    id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    created_by NULL DEFAULT NULL,
    updated_by NULL DEFAULT NULL,
    PRIMARY KEY (id),
    UNIQUE INDEX admins_email_unique (email)
    INDEX admins_email (email),
    INDEX admins_created_by (created_by),
    
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
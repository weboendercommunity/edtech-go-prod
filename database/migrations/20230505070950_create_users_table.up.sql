CREATE TABLE users(
    id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,

    code_verified VARCHAR(255) NULL DEFAULT NULL,
    email_verified_at TIMESTAMP NULL DEFAULT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL,

    created_by INT(10) UNSIGNED NULL DEFAULT NULL,
    updated_by INT(10) UNSIGNED NULL DEFAULT NULL,

    PRIMARY KEY (id),
    UNIQUE INDEX users_email_unique (email),
    INDEX users_name_index (name),
    INDEX users_email_index (email),
    INDEX users_created_by_index (created_by),
    INDEX users_updated_by_index (updated_by),

    CONSTRAINT users_created_by_foreign FOREIGN KEY (created_by) REFERENCES admins (id) ON DELETE SET NULL,
    CONSTRAINT users_updated_by_foreign FOREIGN KEY (updated_by) REFERENCES admins (id) ON DELETE SET NULL
)
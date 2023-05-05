CREATE TABLE product_categories(
    id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    created_by INT(10) UNSIGNED NULL DEFAULT NULL,
    updated_by INT(10) UNSIGNED NULL DEFAULT NULL,

    PRIMARY KEY (id),
    INDEX product_categories_name_index (name),

    CONSTRAINT product_categories_created_by_foreign FOREIGN KEY (created_by) REFERENCES admins (id) ON DELETE SET NULL,
    CONSTRAINT product_categories_updated_by_foreign FOREIGN KEY (updated_by) REFERENCES admins (id) ON DELETE SET NULL
)
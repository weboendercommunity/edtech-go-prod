CREATE TABLE carts(
    id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    product_id INT(10) UNSIGNED NULL,
    user_id INT(10) UNSIGNED NULL,

    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL,

    created_by INT(10) UNSIGNED NULL,
    updated_by INT(10) UNSIGNED NULL,

    PRIMARY KEY (id),
    CONSTRAINT carts_product_id_foreign FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE SET NULL,
    CONSTRAINT carts_user_id_foreign FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE SET NULL,
    CONSTRAINT carts_created_by_foreign FOREIGN KEY (created_by) REFERENCES admins (id) ON DELETE SET NULL,
    CONSTRAINT carts_updated_by_foreign FOREIGN KEY (updated_by) REFERENCES admins (id) ON DELETE SET NULL
)
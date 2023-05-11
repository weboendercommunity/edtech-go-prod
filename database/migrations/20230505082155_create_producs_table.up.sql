CREATE TABLE products (
    id INT(10) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NULL DEFAULT NULL,
    price INT(10) UNSIGNED NOT NULL,

    product_category_id INT(10) UNSIGNED NOT NULL,

    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL,

    created_by INT(10) UNSIGNED NOT NULL,
    updated_by INT(10) UNSIGNED NULL,

    CONSTRAINT products_product_category_id_foreign FOREIGN KEY (product_category_id) REFERENCES product_categories(id),
    CONSTRAINT products_created_by_foreign FOREIGN KEY (created_by) REFERENCES admins(id),
    CONSTRAINT products_updated_by_foreign FOREIGN KEY (updated_by) REFERENCES admins(id)
)
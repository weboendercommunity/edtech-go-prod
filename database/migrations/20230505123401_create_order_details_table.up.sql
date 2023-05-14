CREATE TABLE order_details(
    id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    price INT(10) UNSIGNED NOT NULL,
    product_id INT(10) UNSIGNED NULL,
    order_id INT(10) UNSIGNED NOT NULL,

    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL,

    created_by INT(10) UNSIGNED NULL,
    updated_by INT(10) UNSIGNED NULL,

    PRIMARY KEY (id),
    CONSTRAINT order_details_product_id_foreign FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE SET NULL,
    CONSTRAINT order_details_order_id_foreign FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE,\
    CONSTRAINT order_details_created_by_foreign FOREIGN KEY (created_by) REFERENCES users (id) ON DELETE SET NULL,
    CONSTRAINT order_details_updated_by_foreign FOREIGN KEY (updated_by) REFERENCES users (id) ON DELETE SET NULL


)
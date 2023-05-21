CREATE TABLE orders(
    id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    checkout_link VARCHAR(255) NOT NULL,
    price INT(10) UNSIGNED NOT NULL,
    total_price INT(10) UNSIGNED NOT NULL,
    external_id VARCHAR(255) NOT NULL,
    status ENUM('pending', 'completed','settled','canceled') NOT NULL DEFAULT 'pending',

    discount_id INT(10) UNSIGNED NULL DEFAULT NULL,
    user_id INT(10) UNSIGNED NOT NULL,

    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL,

    created_by INT(10) UNSIGNED NULL DEFAULT NULL,
    updated_by INT(10) UNSIGNED NULL DEFAULT NULL,

    PRIMARY KEY (id),
    INDEX orders_checkout_link_index (checkout_link),
    INDEX orders_external_id_index (external_id),

    CONSTRAINT orders_discount_id_foreign FOREIGN KEY (discount_id) REFERENCES discounts (id) ON DELETE SET NULL,
    CONSTRAINT orders_user_id_foreign FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT orders_created_by_foreign FOREIGN KEY (created_by) REFERENCES admins (id) ON DELETE SET NULL,
    CONSTRAINT orders_updated_by_foreign FOREIGN KEY (updated_by) REFERENCES admins (id) ON DELETE SET NULL
)
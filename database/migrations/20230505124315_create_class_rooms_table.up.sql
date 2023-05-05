CREATE TABLE class_rooms(
    id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    product_id INT(10) UNSIGNED NOT NULL,
    user_id INT(10) UNSIGNED NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,

    created_by INT(10) UNSIGNED NULL,
    updated_by INT(10) UNSIGNED NULL,

    PRIMARY KEY (id),
    CONSTRAINT class_rooms_product_id_foreign FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE,
    CONSTRAINT class_rooms_user_id_foreign FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    
    CONSTRAINT class_rooms_created_by_foreign FOREIGN KEY (created_by) REFERENCES admins (id) ON DELETE SET NULL,
    CONSTRAINT class_rooms_updated_by_foreign FOREIGN KEY (updated_by) REFERENCES admins (id) ON DELETE SET NULL
)
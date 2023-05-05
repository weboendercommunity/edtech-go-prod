CREATE TABLE discounts(
    id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(255) NOT NULL,
    quantity INT(10) UNSIGNED NOT NULL,
    remaining_quantity INT(10) UNSIGNED NOT NULL,
    type ENUM('percentage', 'fixed') NOT NULL,

    start_date TIMESTAMP NULL,
    end_date TIMESTAMP NULL,

    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL,

    created_by INT(10) UNSIGNED NULL DEFAULT NULL,
    updated_by INT(10) UNSIGNED NULL DEFAULT NULL,

    PRIMARY KEY (id),
    INDEX discounts_name_index (name),
    INDEX discounts_code_index (code),

    CONSTRAINT discounts_created_by_foreign FOREIGN KEY (created_by) REFERENCES admins (id) ON DELETE SET NULL,
    CONSTRAINT discounts_updated_by_foreign FOREIGN KEY (updated_by) REFERENCES admins (id) ON DELETE SET NULL
)
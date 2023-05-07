CREATE TABLE forgot_passwords (
    id INT(10) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id INT(10) UNSIGNED NOT NULL,
    expired_at TIMESTAMP NULL,
    valid BOOLEAN NOT NULL DEFAULT 1,
    verification_code VARCHAR(255) NOT NULL,
    
    CONSTRAINT forgot_passwords_user_id_foreign FOREIGN KEY (user_id) REFERENCES users (id)
)
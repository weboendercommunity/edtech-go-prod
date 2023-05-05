CREATE TABLE oauth_access_tokens(
    id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    access_token VARCHAR(100) NOT NULL,
    user_id INT(10) UNSIGNED NOT NULL,
    scope VARCHAR(255) NULL,
    oauth_client_id INT(10) UNSIGNED NOT NULL,
    expired_at TIMESTAMP NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,

    PRIMARY KEY (id),
    CONSTRAINT oauth_access_tokens_oauth_client_id_foreign FOREIGN KEY (oauth_client_id) REFERENCES oauth_clients (id) ON DELETE CASCADE
)
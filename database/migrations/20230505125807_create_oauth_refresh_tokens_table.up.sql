CREATE TABLE oauth_refresh_tokens(
    id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    user_id INT(10) UNSIGNED NOT NULL,
    access_token VARCHAR(100) NOT NULL,
    oauth_access_token_id INT(10) UNSIGNED NOT NULL,
    expired_at TIMESTAMP NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,

    PRIMARY KEY (id),
    CONSTRAINT oauth_refresh_tokens_oauth_access_token_id_foreign FOREIGN KEY (oauth_access_token_id) REFERENCES oauth_access_tokens (id) ON DELETE CASCADE
)
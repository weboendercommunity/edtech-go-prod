CREATE TABLE oauth_clients(
    id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    client_id VARCHAR(255) NOT NULL,
    client_secret VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    redirect VARCHAR(255) NULL,
    description TEXT NULL,
    scope VARCHAR(255) NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,

    PRIMARY KEY (id)
)
-- create databases
CREATE DATABASE IF NOT EXISTS `keycloak`;
CREATE DATABASE IF NOT EXISTS `chat`;

-- create root user and grant rights
-- CREATE USER 'root'@'localhost' IDENTIFIED BY 'local';
-- GRANT ALL PRIVILEGES ON *.* TO 'root'@'%';

CREATE USER 'keycloak'@'keycloak' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON `keycloak`.* TO 'keycloak'@'keycloak';

CREATE USER 'chat'@'localhost' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON `chat`.* TO 'chat'@'localhost';
FLUSH PRIVILEGES;
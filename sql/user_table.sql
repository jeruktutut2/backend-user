CREATE TABLE IF NOT EXISTS user(
	id varchar(36) PRIMARY KEY,
	username varchar(255) NOT NULL,
	password varchar(255) NOT NULL,
	UNIQUE KEY(username)
)ENGINE=INNODB;

INSERT INTO user(id, username, password) VALUES(uuid(), "admin", "$2a$04$bwRt0F41n/MsvPuKgPG3ZuQsK7A78Fy0CwxD0MII8QjcTFgT9iKgW");
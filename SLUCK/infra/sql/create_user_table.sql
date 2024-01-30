BEGIN;

SHOW DATABASES;

SELECT DATABASE();

USE sluck;

SELECT DATABASE();

SHOW TABLES;

CREATE TABLE user (id int auto_increment, name varchar(255), email varchar(255), age int, created_time datetime, updated_time datetime, index(id));

SHOW TABLES;

INSERT INTO user (name, email) values ('Tanaka', 'hoge@gmail.com');

SELECT * FROM user;

COMMIT;

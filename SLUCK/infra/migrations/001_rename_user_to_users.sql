BEGIN;

SHOW DATABASES;

SELECT DATABASE();

USE sluck;

RENAME TABLE user TO users;

SELECT * FROM users;

COMMIT;
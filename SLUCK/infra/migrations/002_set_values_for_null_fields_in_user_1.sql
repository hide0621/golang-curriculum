BEGIN;

SHOW DATABASES;

SELECT DATABASE();

USE sluck;

SELECT DATABASE();

UPDATE users 
SET age = 25, 
    created_time = NOW(), 
    updated_time = NOW() 
WHERE id = 1;

SELECT * FROM users;

COMMIT;
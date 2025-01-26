ALTER TABLE `users` ADD `username` varchar(100) NOT NULL AFTER `id`;

ALTER TABLE users
ADD CONSTRAINT UNIQUE unique_username (username);
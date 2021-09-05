CREATE DATABASE funsiba; 
USE funsiba; 

CREATE TABLE  `avatar`(
	avatar_id INT AUTO_INCREMENT PRIMARY KEY,
    url VARCHAR(255) NOT NULL, 
	created_at DATETIME NOT NULL DEFAULT NOW(),
    updated_at DATETIME, 
	deleted_at DATETIME
); 

CREATE TABLE `user`(
	user_id INT AUTO_INCREMENT PRIMARY KEY, 
	username VARCHAR(20) NOT NULL,
	avatar_id INT NOT NULL, 
    score INT NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT NOW(),
    updated_at DATETIME, 
	deleted_at DATETIME, 
    CONSTRAINT FOREIGN KEY(avatar_id) REFERENCES avatar(avatar_id)
);

CREATE TABLE `subject`(
	subject_id INT AUTO_INCREMENT PRIMARY KEY, 
    name VARCHAR(50) NOT NULL, 
    color VARCHAR(50),
	created_at DATETIME NOT NULL DEFAULT NOW(),
    updated_at DATETIME, 
	deleted_at DATETIME
); 